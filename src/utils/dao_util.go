package utils

import "common"
import "config"
import "fmt"

func createKeyByTEI(vtype int32, entity int32, instant int32) (string) {
	key := fmt.Sprintf("_VT:%d_E:%d_I:%d_", vtype, entity, instant)
	return key
}

func createKeyByEIV(entity int32, instant int32, value interface{}) (string) {
	key := fmt.Sprintf("_E:%d_I:%d_V:%d_", entity, instant, ToInt(value))
	return key
}

func createKeyByEI(entity int32, instant int32) (string) {
	key := fmt.Sprintf("_E:%d_I:%d_", entity, instant)
	return key
}

func createKeyByE(entity int32) (string){
	key := fmt.Sprintf("_E:%d_", entity)
	return key
}

// GET record
func PullObj(key string) (interface{}) {
	data := common.DeviceInfo_t { Entity:-1, Instant:-1, ValueType:config.TYPE_OBJECT, Value: ""}
	data.Key = key
	msg := TalkToDao(config.GET_OBJ_BY_KEY, data)
	if IsResponse(msg.Data) {
		return nil
	} else {
		return msg.Data
	}
}

// SET record
func PushObj(key string, val interface{}) (bool){
	data := common.DeviceInfo_t { Entity:-1, Instant:-1, ValueType:config.TYPE_OBJECT, Value:val }
	data.Key = key
	msg := TalkToDao(config.SET_OBJ_BY_KEY, data)
	rsp, ok := (msg.Data).(common.DeviceInfo_t)
	if ok {
		if rsp.Value == 0 {
			return true
		} 
	}
	return false
}

// GET DeviceMaxTemp
func PullObjDeviceMaxTemp(entity int32, instant int32) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_TEMPERATURE, Value:0 }
	data.Key = createKeyByTEI(data.ValueType, data.Entity, data.Instant)
	msg := TalkToDao(config.GET_DEVICE_MAXTEMP, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// SET DeviceMaxTemp
func PushObjDeviceMaxTemp(entity int32, instant int32, value float32) (common.DeviceInfo_t){
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_TEMPERATURE, Value:value }
	data.Key = createKeyByTEI(data.ValueType, data.Entity, data.Instant)
	msg := TalkToDao(config.SET_DEVICE_MAXTEMP, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// GET DeviceAbsTemp
func PullObjDeviceAbsTemp(entity int32, instant int32) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_TEMPERATURE, Value:0 }
	data.Key = createKeyByTEI(data.ValueType, data.Entity, data.Instant)
	msg := TalkToDao(config.GET_DEVICE_ABSTEMP, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// SET DeviceAbsTemp
func PushObjDeviceAbsTemp(entity int32, instant int32, value float32) (common.DeviceInfo_t){
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_TEMPERATURE, Value:value }
	data.Key = createKeyByTEI(data.ValueType, data.Entity, data.Instant)
	msg := TalkToDao(config.SET_DEVICE_ABSTEMP, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// Get DeviceAbsTemp list
func PullObjListDeviceAbsTemp(entity int32) (map[string]common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:-1, ValueType:config.TYPE_TEMPERATURE, Value:0 }
	data.Key = createKeyByE(data.Entity)
	msg := TalkToDao(config.GET_DEVICE_LIST_ABSTEMP, data)
	obj := (msg.Data).(map[string]common.DeviceInfo_t)
	return obj
}
// GET DeviceRelTemp
func PullObjDeviceRelTemp(entity int32, instant int32) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_TEMPERATURE, Value:0 }
	data.Key = createKeyByTEI(data.ValueType, data.Entity, data.Instant)
	msg := TalkToDao(config.GET_DEVICE_RELTEMP, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// SET DeviceRelTemp
func PushObjDeviceRelTemp(entity int32, instant int32, value float32) (common.DeviceInfo_t){
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_TEMPERATURE, Value:value }
	data.Key = createKeyByTEI(data.ValueType, data.Entity, data.Instant)
	msg := TalkToDao(config.SET_DEVICE_RELTEMP, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// GET record list
func PullObjListDeviceRelTemp(entity int32) (map[string]common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:-1, ValueType:config.TYPE_TEMPERATURE, Value:0 }
	data.Key = createKeyByE(data.Entity)
	msg := TalkToDao(config.GET_DEVICE_LIST_RELTEMP, data)
	obj := (msg.Data).(map[string]common.DeviceInfo_t)
	return obj
}

// GET DeviceAveragePower
func PullObjDeviceAveragePower(entity int32, instant int32) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_AVERAGEPOWER, Value:0 }
	data.Key = createKeyByTEI(data.ValueType, data.Entity, data.Instant)
	msg := TalkToDao(config.GET_DEVICE_AVERAGEPOWER, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// SET DeviceAveragePower
func PushObjDeviceAveragePower(entity int32, instant int32, value float32) (common.DeviceInfo_t){
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_AVERAGEPOWER, Value:value }
	data.Key = createKeyByTEI(data.ValueType, data.Entity, data.Instant)
	msg := TalkToDao(config.SET_DEVICE_AVERAGEPOWER, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// GET DeviceAveragePower list
func PullObjListDeviceAveragePower(entity int32) (map[string]common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:-1, ValueType:config.TYPE_AVERAGEPOWER, Value:0 }
	data.Key = createKeyByE(data.Entity)
	msg := TalkToDao(config.GET_DEVICE_LIST_AVERAGEPOWER, data)
	obj := (msg.Data).(map[string]common.DeviceInfo_t)
	return obj
}

// SET ExpectFanDuty
func PushObjExpectFanDuty(key string, entity int32, instant int32, value float32) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_FANDUTY, Value:value }
	data.Key = key
	msg := TalkToDao(config.SET_EXPECT_FAN_DUTY, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// GET record list
func PullObjListDeviceFanMap()(map[string]common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:-1, Instant:-1, ValueType:config.TYPE_DEVICEFANMAP, Value:0 }
	data.Key = ""
	msg := TalkToDao(config.GET_ALL_DEVICE_FAN_MAP, data)
	obj := (msg.Data).(map[string]common.DeviceInfo_t)
	return obj
}
// GET record list
func PullObjDeviceFanMap(entity int32, instant int32) (map[string]common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_DEVICEFANMAP, Value:0 }
	data.Key = createKeyByEI(data.Entity, data.Instant)
	msg := TalkToDao(config.GET_DEVICE_FAN_MAP, data)
	obj := (msg.Data).(map[string]common.DeviceInfo_t)
	return obj
}
// SET DeviceFanMap
func PushObjDeviceFanMap(entity int32, instant int32, fan_index int32) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_DEVICEFANMAP, Value:fan_index }
	data.Key = createKeyByEIV(data.Entity, data.Instant, data.Value)
	msg := TalkToDao(config.SET_DEVICE_FAN_MAP, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// GET record
func PullObjDeviceFanDutyOutput(entity int32, instant int32) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_FANDUTY, Value:0 }
	data.Key = createKeyByEI(data.Entity, data.Instant)
	msg := TalkToDao(config.GET_DEVICE_FAN_DUTY_OUTPUT, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// SET record
func PushObjDeviceFanDutyOutput(entity int32, instant int32, duty float32) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_FANDUTY, Value:duty }
	data.Key = createKeyByEI(data.Entity, data.Instant)
	msg := TalkToDao(config.SET_DEVICE_FAN_DUTY_OUTPUT, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// SET record
func InitObjDeviceFanDutyOutput(entity int32, instant int32, duty float32) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_INITFANDUTY, Value:duty }
	data.Key = createKeyByEI(data.Entity, data.Instant)
	msg := TalkToDao(config.SET_DEVICE_FAN_DUTY_OUTPUT, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// GET record list
func PullObjListDeviceFanDutyOutput() (map[string]common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:config.ENTITY_FAN, Instant:-1, ValueType:config.TYPE_FANDUTY, Value:0 }
	data.Key = ""
	msg := TalkToDao(config.GET_ALL_FAN_DUTY_OUTPUT, data)
	obj := (msg.Data).(map[string]common.DeviceInfo_t)
	return obj
}
/*
func PushObjExpectFanDutyByTemperature(entity int, instant int, key string, value float32) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, Key:key, ValueType:config.TYPE_FANDUTY_TEMPERATURE, Value:value }
	msg := TalkToDao(config.SET_EXPECT_FAN_DUTY, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}

func PushObjExpectFanDutyByPower(entity int, instant int, value float32) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_FANDUTY_POWER, Value:value }
	msg := TalkToDao(config.SET_EXPECT_FAN_DUTY, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
*/
// GET record
func PullObjDeviceMaxPower(entity int32, instant int32) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_MAXPOWER, Value:0 }
	data.Key = createKeyByTEI(data.ValueType, data.Entity, data.Instant)
	msg := TalkToDao(config.GET_DEVICE_MAXPOWER, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// GET record list
func PullObjListDeviceMaxPower(entity int32) (map[string]common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:-1, ValueType:config.TYPE_MAXPOWER, Value:0 }
	data.Key = createKeyByE(data.Entity)
	msg := TalkToDao(config.GET_DEVICE_LIST_MAXPOWER, data)
	obj := (msg.Data).(map[string]common.DeviceInfo_t)
	return obj
}
// SET record
func PushObjDeviceMaxPower(entity int32, instant int32, value float32) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_MAXPOWER, Value:value }
	data.Key = createKeyByTEI(data.ValueType, data.Entity, data.Instant)
	msg := TalkToDao(config.SET_DEVICE_MAXPOWER, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// GET record list
func PullObjListDevicesExpectFanDuty() (map[string]common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity: -1, Instant: -1, ValueType: config.TYPE_FANDUTY, Value: 0 }
	data.Key = ""
	msg := TalkToDao(config.GET_ALL_EXPECT_FAN_DUTY, data)
	obj_map := (msg.Data).(map[string]common.DeviceInfo_t)
	return obj_map
}

