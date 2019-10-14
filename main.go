package main

import "os"
import "hwmon"
import "fmt"
import "mailbox"

/*
 * Tasks:
 *  - rest Interface
 *  - DAO
 *  - Fan control Algorisim
 * [www rest interface] -> [dispatcher] -> 
 *    request               DataStore
 */
func main() {
    if len(os.Args) <= 1 {
        fmt.Println("program <www_dir>")
        os.Exit(-1)
    }
    www_path := os.Args[1]

    fmt.Println("Creating & Running TaskRest")
    task_rest := new (hwmon.TaskRest)
    task_rest.SetFolder(www_path)
    go task_rest.Run()

    fmt.Println("Creating & Running TaskDao")
    task_dao := new (hwmon.TaskDao)
    go task_dao.Run()

    fmt.Println("Creating & Running TaskMsgHndlr")
    task_msghndlr := new (hwmon.TaskMsgHndlr)
    go task_msghndlr.Run()

    fmt.Println("Modules starting...")

    tasks := hwmon.GetModules()
    for index := range tasks {
        go tasks[index].RunTask()
    }
    mb_hwmon := mailbox.CreateMailboxHwmon()
    isBreakTask := false
    for {
        msg :=<-mb_hwmon.Channel
        switch msg.Function {
        case hwmon.EXIT_APPLICATION:

            isBreakTask = true
	    var res_msg mailbox.Msg_t
            data := hwmon.DeviceInfo_t { Entity:0, Instant:0, ValueType:hwmon.TYPE_RSP_EXIT, Value:"Stop task" }
            res_bytes := hwmon.ConvertDeviceInfoToBytes(data)
            res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, res_bytes)
            msg.ChannelDst <-res_msg

	    data = hwmon.DeviceInfo_t { Entity:0, Instant:0, ValueType:hwmon.TYPE_REQ_EXIT, Value:0 }
	    res_msg = hwmon.TalkToRest(hwmon.EXIT_APPLICATION, data)
	    fmt.Println(res_msg)
            for {
                isFuncExit := true
                for index := range tasks {
                    tasks[index].FunctionExit = true
		    if tasks[index].FunctionStatus != hwmon.FUNC_STAT_EXIT {
                        isFuncExit = false
		    }
                }
                if isFuncExit {
                    break
                }
            }
	    res_msg = hwmon.TalkToDao(hwmon.EXIT_APPLICATION, data)
	    fmt.Println(res_msg)
	    res_msg = hwmon.TalkToMsghndlr(hwmon.EXIT_APPLICATION, data)
	    fmt.Println(res_msg)
        default:
	}
        //mb_hwmon.Channel <-msg
        if isBreakTask {
            break
        }
    }
    fmt.Println("Modules end...")
}

