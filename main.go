package main

import "os"
import "config"
import "common"
import "utils"
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
    var data common.DeviceInfo_t
    for {
        msg :=<-mb_hwmon.Channel
	msg_func := msg.Function
        switch msg_func {
        case config.EXIT_APPLICATION:
            data = common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_RSP_OK, Value:"Stop task" }
	case config.DISABLE_OUTOFBAND_INTERFACE:
            data = common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_RSP_OK, Value:"Disable interface" }
	case config.ENABLE_OUTOFBAND_INTERFACE:
            data = common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_RSP_OK, Value:"Enable interface" }
        default:
            data = common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_RSP_ERROR, Value:"Command Not Found" }
	}
        var res_msg common.Msg_t
        //res_bytes := hwmon.ConvertDeviceInfoToBytes(data)
	res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, data)
        msg.ChannelDst <-res_msg

	switch msg_func {
        case config.EXIT_APPLICATION:
            isBreakTask = true
	    data = common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_REQ_CMD, Value:0 }
	    res_msg = utils.TalkToRest(config.EXIT_APPLICATION, data)
	    //fmt.Println(res_msg)
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
	    res_msg = utils.TalkToDao(config.EXIT_APPLICATION, data)
	    //fmt.Println(res_msg)
	    res_msg = utils.TalkToMsghndlr(config.EXIT_APPLICATION, data)
	    //fmt.Println(res_msg)
	case config.DISABLE_OUTOFBAND_INTERFACE:
	    data = common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_REQ_CMD, Value:0 }
	    res_msg = utils.TalkToRest(config.DISABLE_OUTOFBAND_INTERFACE, data)
	    //fmt.Println(res_msg)
	case config.ENABLE_OUTOFBAND_INTERFACE:
	    data = common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_REQ_CMD, Value:0 }
	    res_msg = utils.TalkToRest(config.ENABLE_OUTOFBAND_INTERFACE, data)
	    //fmt.Println(res_msg)
        default:
	}

        //mb_hwmon.Channel <-msg
        if isBreakTask {
            break
        }
    }
    fmt.Println("Modules end...")
}

