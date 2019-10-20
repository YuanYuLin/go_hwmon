package utils

import "common"
/*
 *
 */
type DeviceAmb struct {
	Entity int
}

func (o *DeviceAmb)GetListAbsTemp() (map[string]common.DeviceInfo_t) {
	obj := PullObjListDeviceAbsTemp(o.Entity)
	return obj
}

func (o *DeviceAmb)GetAbsTemp(instant int) (common.DeviceInfo_t) {
	obj := PullObjDeviceAbsTemp(o.Entity, instant)
	return obj
}

func (o *DeviceAmb)SetAbsTemp(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjDeviceAbsTemp(o.Entity, instant, value)
	return obj
}

func (o *DeviceAmb)SetExpectFanDutyByTemperature(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjExpectFanDutyByTemperature(o.Entity, instant, value)
	return obj
}

func (o *DeviceAmb)SetExpectFanDutyByPower(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjExpectFanDutyByPower(o.Entity, instant, value)
	return obj
}

