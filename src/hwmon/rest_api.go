package hwmon

import "encoding/json"
import "net/http"
import "ops_log"
import "io/ioutil"
//import "fmt"

type json_msg_t struct {
    Status              int     `json:"status"`
    Version             int     `json:"version"`
    Data                interface{} `json:"data"`
}

func responseWithJsonV1(w http.ResponseWriter, code int,  data interface{}) {
    json_msg := json_msg_t { Status:1, Version:1, Data:data }
    response, _ := json.Marshal(json_msg)
    ops_log.Debug(0x01, "Response : %s", string(response))
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func GetDeviceTemperature(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }

    data := ConvertBytesToDeviceInfo(b)
    obj := PullObjDeviceTemperature(data.Entity, data.Instant)
    responseWithJsonV1(w, http.StatusOK, obj)
}

func SetDeviceTemperature(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }
    data := ConvertBytesToDeviceInfo(b)
    val := data.Value.(float64)
    obj := PushObjDeviceTemperature(data.Entity, data.Instant, val)
    responseWithJsonV1(w, http.StatusOK, obj)
}

func GetDeviceAveragePower(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }

    data := ConvertBytesToDeviceInfo(b)
    obj := PullObjDeviceAveragePower(data.Entity, data.Instant)
    responseWithJsonV1(w, http.StatusOK, obj)
}

func SetDeviceAveragePower(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }
    data := ConvertBytesToDeviceInfo(b)
    val := data.Value.(float64)
    obj := PushObjDeviceAveragePower(data.Entity, data.Instant, val)
    responseWithJsonV1(w, http.StatusOK, obj)
}

func GetDeviceMaxPower(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }

    data := ConvertBytesToDeviceInfo(b)
    obj := PullObjDeviceMaxPower(data.Entity, data.Instant)
    responseWithJsonV1(w, http.StatusOK, obj)
}

func SetDeviceMaxPower(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }
    data := ConvertBytesToDeviceInfo(b)
    val := data.Value.(float64)
    obj := PushObjDeviceMaxPower(data.Entity, data.Instant, val)
    responseWithJsonV1(w, http.StatusOK, obj)
}

func ExitMain(w http.ResponseWriter, r* http.Request) {
    _, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }
    data := DeviceInfo_t { Entity:0, Instant:0, ValueType:TYPE_REQ_EXIT, Value:0 }
    res_msg := TalkToHwmon(EXIT_APPLICATION, data)
    obj := ConvertBytesToDeviceInfo(res_msg.Data)
    responseWithJsonV1(w, http.StatusOK, obj)
}
