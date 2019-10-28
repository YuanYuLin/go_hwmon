package device

import "common"
import "utils"
/*
 *
 */
type Fan_t struct {
	Entity int32
}

func (o *Fan_t)GetDutyOutput(instant int32) (common.DeviceInfo_t) {
	obj := utils.PullObjDeviceFanDutyOutput(o.Entity, instant)
	return obj
}

func (o *Fan_t)SetDutyOutput(instant int32, duty float32) (common.DeviceInfo_t) {
	obj := utils.PushObjDeviceFanDutyOutput(o.Entity, instant, duty)
	return obj
}

func (o *Fan_t)InitDutyOutput(instant int32, duty float32) (common.DeviceInfo_t) {
	obj := utils.InitObjDeviceFanDutyOutput(o.Entity, instant, duty)
	return obj
}

func (o *Fan_t)GetAllDutyOutput() (map[string]common.DeviceInfo_t) {
	obj_list := utils.PullObjListDeviceFanDutyOutput()
	return obj_list
}

func (o *Fan_t)SetDeviceMap(dev_entity int32, dev_instant int32, fan_instant int32) (common.DeviceInfo_t) {
	obj := utils.PushObjDeviceFanMap(dev_entity, dev_instant, fan_instant)
	return obj
}

func (o *Fan_t)GetDeviceMap(dev_entity int32, dev_instant int32) (map[string]common.DeviceInfo_t) {
	obj := utils.PullObjDeviceFanMap(dev_entity, dev_instant)
	return obj
}

func (o *Fan_t)GetAllDevicesExpectFanDuty() (map[string]common.DeviceInfo_t) {
	obj_list := utils.PullObjListDevicesExpectFanDuty()
	return obj_list
}
