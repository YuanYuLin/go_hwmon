package hwmon

import "common"
import "utils"
import "config"
import "net/http"
import "ops_log"
import "io/ioutil"

func GetDeviceMaxTemp(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }

    data := utils.ConvertBytesToDeviceInfo(b)
    obj := utils.PullObjDeviceMaxTemp(data.Entity, data.Instant)
    responseWithJsonV1(w, http.StatusOK, obj)
}
func SetDeviceMaxTemp(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }
    data := utils.ConvertBytesToDeviceInfo(b)
    val := utils.ToFloat(data.Value)
    obj := utils.PushObjDeviceMaxTemp(data.Entity, data.Instant, val)
    responseWithJsonV1(w, http.StatusOK, obj)
}

func GetDeviceAbsTemp(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }

    data := utils.ConvertBytesToDeviceInfo(b)
    obj := utils.PullObjDeviceAbsTemp(data.Entity, data.Instant)
    responseWithJsonV1(w, http.StatusOK, obj)
}
func GetDeviceRelTemp(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }

    data := utils.ConvertBytesToDeviceInfo(b)
    obj := utils.PullObjDeviceRelTemp(data.Entity, data.Instant)
    responseWithJsonV1(w, http.StatusOK, obj)
}

func SetDeviceAbsTemp(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }
    data := utils.ConvertBytesToDeviceInfo(b)
    val := utils.ToFloat(data.Value)
    obj := utils.PushObjDeviceAbsTemp(data.Entity, data.Instant, val)
    responseWithJsonV1(w, http.StatusOK, obj)
}
func SetDeviceRelTemp(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }
    data := utils.ConvertBytesToDeviceInfo(b)
    val := utils.ToFloat(data.Value)
    obj := utils.PushObjDeviceRelTemp(data.Entity, data.Instant, val)
    responseWithJsonV1(w, http.StatusOK, obj)
}

func GetDeviceAveragePower(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }

    data := utils.ConvertBytesToDeviceInfo(b)
    obj := utils.PullObjDeviceAveragePower(data.Entity, data.Instant)
    responseWithJsonV1(w, http.StatusOK, obj)
}

func SetDeviceAveragePower(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }
    data := utils.ConvertBytesToDeviceInfo(b)
    val := utils.ToFloat(data.Value)
    obj := utils.PushObjDeviceAveragePower(data.Entity, data.Instant, val)
    responseWithJsonV1(w, http.StatusOK, obj)
}

func GetDeviceMaxPower(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }

    data := utils.ConvertBytesToDeviceInfo(b)
    obj := utils.PullObjDeviceMaxPower(data.Entity, data.Instant)
    responseWithJsonV1(w, http.StatusOK, obj)
}

func SetDeviceMaxPower(w http.ResponseWriter, r* http.Request) {
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }
    data := utils.ConvertBytesToDeviceInfo(b)
    val := utils.ToFloat(data.Value)
    obj := utils.PushObjDeviceMaxPower(data.Entity, data.Instant, val)
    responseWithJsonV1(w, http.StatusOK, obj)
}

func GetMapAllDeviceFan(w http.ResponseWriter, r* http.Request) {
    obj := utils.PullObjListDeviceFanMap()
    responseWithJsonV1(w, http.StatusOK, obj)
}
func GetMapAllFanDutyOut(w http.ResponseWriter, r* http.Request) {
    obj := utils.PullObjListDeviceFanDutyOutput()
    responseWithJsonV1(w, http.StatusOK, obj)
}
func GetMapAllDevicesExpectFanDuty(w http.ResponseWriter, r* http.Request) {
    obj := utils.PullObjListDevicesExpectFanDuty()
    responseWithJsonV1(w, http.StatusOK, obj)
}

func ExitMain(w http.ResponseWriter, r* http.Request) {
    _, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }
    value := common.ValueRequest_t{Value: "Exit Main"}
    data := common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_REQUEST, Value:value }
    res_msg := utils.TalkToHwmon(config.EXIT_APPLICATION, data)
    //obj := ConvertBytesToDeviceInfo(res_msg.Data)
    obj := res_msg.Data
    responseWithJsonV1(w, http.StatusOK, obj)
}

func EnableOutOfBandInterface(w http.ResponseWriter, r* http.Request) {
    _, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }
    value := common.ValueRequest_t{Value: "Enable out of band interface"}
    data := common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_REQUEST, Value:value }
    res_msg := utils.TalkToHwmon(config.ENABLE_OUTOFBAND_INTERFACE, data)
    //obj := ConvertBytesToDeviceInfo(res_msg.Data)
    obj := res_msg.Data
    responseWithJsonV1(w, http.StatusOK, obj)
}

func DisableOutOfBandInterface(w http.ResponseWriter, r* http.Request) {
    _, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        ops_log.Debug(0x1, "Set Error %s", err)
    }
    value := common.ValueRequest_t{Value: "Disable out of band interface"}
    data := common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_REQUEST, Value:value }
    res_msg := utils.TalkToHwmon(config.DISABLE_OUTOFBAND_INTERFACE, data)
    //obj := ConvertBytesToDeviceInfo(res_msg.Data)
    obj := res_msg.Data
    responseWithJsonV1(w, http.StatusOK, obj)
}

var rest_api_list = []rest_api_t {
    /*
     *
     */
    {"/api/v1/hwmon/get/device/abstemp",	GetDeviceAbsTemp},
    {"/api/v1/hwmon/set/device/abstemp",	SetDeviceAbsTemp},
    {"/api/v1/hwmon/get/device/maxtemp",	GetDeviceMaxTemp},
    {"/api/v1/hwmon/set/device/maxtemp",	SetDeviceMaxTemp},

    {"/api/v1/hwmon/get/device/reltemp",	GetDeviceRelTemp},
    {"/api/v1/hwmon/set/device/reltemp",	SetDeviceRelTemp},

    /*
     *
     */
    {"/api/v1/hwmon/get/device/averagepower",	GetDeviceAveragePower},
    {"/api/v1/hwmon/set/device/averagepower",	SetDeviceAveragePower},
    {"/api/v1/hwmon/get/device/maxpower",	GetDeviceMaxPower},
    {"/api/v1/hwmon/set/device/maxpower",	SetDeviceMaxPower},

    /*
     *
     */
    {"/api/v1/hwmon/get/map/alldevicefan",	GetMapAllDeviceFan},
    {"/api/v1/hwmon/get/map/allfandutyout",	GetMapAllFanDutyOut},
    {"/api/v1/hwmon/get/map/allexpectduty",	GetMapAllDevicesExpectFanDuty},

    /*
     *
     */
    {"/api/v1/hwmon/exit/main",			ExitMain},
    {"/api/v1/hwmon/enable/out/interface",	EnableOutOfBandInterface},
    {"/api/v1/hwmon/disable/out/interface",	DisableOutOfBandInterface},
    /*
     *
     */
    {"/", PageIndex},
    {"/debug.html", PageDebug},
}

