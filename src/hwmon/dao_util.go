package hwmon

import "encoding/json"

// GET
func PullObjDeviceTemperature(entity int, instant int) (DeviceInfo_t) {
	data := DeviceInfo_t { Entity:entity, Instant:instant, ValueType:TYPE_TEMPERATURE, Value:0 }
	msg := TalkToDao(GET_DEVICE_TEMPERATURE, data)
	obj := ConvertBytesToDeviceInfo(msg.Data)
	return obj
}
func PullObjListDeviceTemperature(entity int) (map[string]DeviceInfo_t) {
	data := DeviceInfo_t { Entity:entity, Instant:-1, ValueType:TYPE_TEMPERATURE, Value:0 }
	msg := TalkToDao(GET_DEVICE_LIST_TEMPERATURE, data)
	obj := ConvertBytesToDeviceInfoList(msg.Data)
	return obj
}
// SET
func PushObjDeviceTemperature(entity int, instant int, value float64) (DeviceInfo_t){
	data := DeviceInfo_t { Entity:entity, Instant:instant, ValueType:TYPE_TEMPERATURE, Value:value }
	msg := TalkToDao(SET_DEVICE_TEMPERATURE, data)
	obj := ConvertBytesToDeviceInfo(msg.Data)
	return obj
}
// GET
func PullObjDeviceAveragePower(entity int, instant int) (DeviceInfo_t) {
	data := DeviceInfo_t { Entity:entity, Instant:instant, ValueType:TYPE_AVERAGEPOWER, Value:0 }
	msg := TalkToDao(GET_DEVICE_AVERAGEPOWER, data)
	obj := ConvertBytesToDeviceInfo(msg.Data)
	return obj
}
func PullObjListDeviceAveragePower(entity int) (map[string]DeviceInfo_t) {
	data := DeviceInfo_t { Entity:entity, Instant:-1, ValueType:TYPE_AVERAGEPOWER, Value:0 }
	msg := TalkToDao(GET_DEVICE_LIST_AVERAGEPOWER, data)
	obj := ConvertBytesToDeviceInfoList(msg.Data)
	return obj
}
// SET
func PushObjDeviceAveragePower(entity int, instant int, value float64) (DeviceInfo_t){
	data := DeviceInfo_t { Entity:entity, Instant:instant, ValueType:TYPE_AVERAGEPOWER, Value:value }
	msg := TalkToDao(SET_DEVICE_AVERAGEPOWER, data)
	obj := ConvertBytesToDeviceInfo(msg.Data)
	return obj
}
// SET
func PushObjExpectFanDutyByTemperature(entity int, instant int, value float64) (DeviceInfo_t) {
	data := DeviceInfo_t { Entity:entity, Instant:instant, ValueType:TYPE_FANDUTY_TEMPERATURE, Value:value }
	msg := TalkToDao(SET_EXPECT_FAN_DUTY, data)
	obj := ConvertBytesToDeviceInfo(msg.Data)
	return obj
}

func PushObjExpectFanDutyByPower(entity int, instant int, value float64) (DeviceInfo_t) {
	data := DeviceInfo_t { Entity:entity, Instant:instant, ValueType:TYPE_FANDUTY_POWER, Value:value }
	msg := TalkToDao(SET_EXPECT_FAN_DUTY, data)
	obj := ConvertBytesToDeviceInfo(msg.Data)
	return obj
}
//GET
func PullObjDeviceMaxPower(entity int, instant int) (DeviceInfo_t) {
	data := DeviceInfo_t { Entity:entity, Instant:instant, ValueType:TYPE_MAXPOWER, Value:0 }
	msg := TalkToDao(GET_DEVICE_MAXPOWER, data)
	obj := ConvertBytesToDeviceInfo(msg.Data)
	return obj
}
func PullObjListDeviceMaxPower(entity int) (map[string]DeviceInfo_t) {
	data := DeviceInfo_t { Entity:entity, Instant:-1, ValueType:TYPE_MAXPOWER, Value:0 }
	msg := TalkToDao(GET_DEVICE_LIST_MAXPOWER, data)
	obj := ConvertBytesToDeviceInfoList(msg.Data)
	return obj
}
//SET
func PushObjDeviceMaxPower(entity int, instant int, value float64) (DeviceInfo_t) {
	data := DeviceInfo_t { Entity:entity, Instant:instant, ValueType:TYPE_MAXPOWER, Value:value }
	msg := TalkToDao(SET_DEVICE_MAXPOWER, data)
	obj := ConvertBytesToDeviceInfo(msg.Data)
	return obj
}
//
func GetMapExpectFanDuty() (map[string]DeviceInfo_t) {
	dev_map := make(map[string]DeviceInfo_t)
	obj := DeviceInfo_t { Entity: 0, Instant: 0, ValueType: "string", Value: "temp_" }
	msg := TalkToDao(GET_ALL_EXPECT_FAN_DUTY, obj)
	var obj_map map[string][]byte
	json.Unmarshal(msg.Data, &obj_map)
	for key, value := range obj_map {
		var data DeviceInfo_t
		json.Unmarshal(value, &data)
		dev_map[key]=data
	}
	return dev_map
}

