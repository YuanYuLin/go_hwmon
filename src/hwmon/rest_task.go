package hwmon

import "net/http"
import "context"
import "mailbox"
import "fmt"

type TaskRest struct {
	Folder	string
}

type rest_api_function_t func(w http.ResponseWriter, r *http.Request)

type rest_api_t struct {
    Url		string
    Function	rest_api_function_t
}

var api_url_prefix = "/api/{api_version}"

var rest_api_list = []rest_api_t {
    {"/api/v1/hwmon/get/device/temperature",	GetDeviceTemperature},
    {"/api/v1/hwmon/set/device/temperature",	SetDeviceTemperature},

    {"/api/v1/hwmon/get/device/averagepower",	GetDeviceAveragePower},
    {"/api/v1/hwmon/set/device/averagepower",	SetDeviceAveragePower},

    {"/api/v1/hwmon/get/device/maxpower",	GetDeviceMaxPower},
    {"/api/v1/hwmon/set/device/maxpower",	SetDeviceMaxPower},

    {"/api/v1/hwmon/exit/main",			ExitMain},
}

var SERVICE_PORT = "localhost:8080"

func (o* TaskRest)SetFolder(folder string) {
	o.Folder = folder
}

func (o* TaskRest)Run() {
	mux := http.NewServeMux()
	srv  := http.Server{Addr: SERVICE_PORT, Handler: mux}

	for _, rest := range rest_api_list {
		mux.HandleFunc(rest.Url, rest.Function)
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		}
	}()
	mb_rest := mailbox.CreateMailboxRest()
	var res_msg mailbox.Msg_t
	isBreakTask := false
	for {
		msg :=<-mb_rest.Channel
		switch msg.Function {
		case EXIT_APPLICATION:
			srv.Shutdown(context.Background())
			isBreakTask = true
			data := DeviceInfo_t { Entity:0, Instant:0, ValueType:TYPE_RSP_EXIT, Value:"Stop task" }
			res_bytes := ConvertDeviceInfoToBytes(data)
			res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, res_bytes)
		}
		msg.ChannelDst <- res_msg
		if isBreakTask {
			break
		}
	}
	fmt.Println("Exit TaskRest")
}
