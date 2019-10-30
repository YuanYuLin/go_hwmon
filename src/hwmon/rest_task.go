package hwmon

import "common"
import "encoding/json"
import "net/http"
//import "time"
import "context"
import "mailbox"
import "fmt"
//import "ops_log"
import "config"

type TaskRest struct {
	Folder	string
}

type rest_api_function_t func(w http.ResponseWriter, r *http.Request)

type rest_api_t struct {
    Url		string
    Function	rest_api_function_t
}

func responseWithJsonV1(w http.ResponseWriter, code int,  data interface{}) {
    json_msg := common.JsonMsg_t { Status:1, Version:1, Data:data }
    response, _ := json.Marshal(json_msg)
    //ops_log.Debug(0x01, "Response : %s", string(response))
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func (o* TaskRest)SetFolder(folder string) {
	o.Folder = folder
}

func (o* TaskRest)Run() {
	mux := http.NewServeMux()
	for _, rest := range rest_api_list {
		mux.HandleFunc(rest.Url, rest.Function)
	}

	srv  := http.Server{ Addr: config.IN_SERVICE_PORT, Handler: mux }
	is_started_srv := true // default enabled
	is_kill_srv := false

	//var srv2 http.Server
	srv2 := http.Server{ Addr: config.OUT_SERVICE_PORT, Handler: mux }
	is_started_srv2 := false
	is_kill_srv2 := false

	mb_rest := mailbox.CreateMailboxRest()
	var res_msg common.Msg_t
	isBreakTask := false

	if is_started_srv {
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				fmt.Println(err)
			}
			fmt.Printf("Stop Listen And Serve\n")
		}()
	}

	var data common.DeviceInfo_t
	for {
		msg :=<-mb_rest.Channel
		switch msg.Function {
		case config.EXIT_APPLICATION:
			if is_started_srv {
				is_started_srv = false
				is_kill_srv = true
			}
			if is_started_srv2 {
				is_started_srv2 = false
				is_kill_srv2 = true
			}
			isBreakTask = true
			data = common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_RESPONSE, Value:"Stop task" }
		case config.ENABLE_OUTOFBAND_INTERFACE:
			if !is_started_srv2 {
				go func() {
					if err := srv2.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						fmt.Println(err)
					}
					fmt.Printf("Stop Listen And Serve2\n")
				}()
			}
			is_started_srv2 = true
			is_kill_srv2 = false
			value := common.ValueResponse_t { Value: config.RESPONSE_OK }
			data = common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_RESPONSE, Value:value }
		case config.DISABLE_OUTOFBAND_INTERFACE:
			if is_started_srv2 {
				is_started_srv2 = false
				is_kill_srv2 = true
				fmt.Printf("Disable out of band interface\n")
			}
			value := common.ValueResponse_t { Value: config.RESPONSE_OK }
			data = common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_RESPONSE, Value:value }
		default:
			value := common.ValueResponse_t { Value: config.RESPONSE_NOT_FOUND }
			data = common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_RESPONSE, Value:value }
		}
		//res_bytes := ConvertDeviceInfoToBytes(data)
		res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, data)
		msg.ChannelDst <- res_msg
		if is_kill_srv {
			srv.Shutdown(context.Background())
		}
		if is_kill_srv2 {
			fmt.Printf("closing srv2\n")
			srv2.Shutdown(context.Background())
			is_kill_srv2 = false
			fmt.Printf("closed srv2\n")
		}
		if isBreakTask {
			break
		}
	}
	fmt.Println("Exit TaskRest")
}
