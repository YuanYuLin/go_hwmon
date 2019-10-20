package utils

import "common"
/*
 *
 */
type DeviceCpu struct {
	Entity int
}

func (o *DeviceCpu)GetListAbsTemp() (map[string]common.DeviceInfo_t) {
	obj := PullObjListDeviceAbsTemp(o.Entity)
	return obj
}

func (o *DeviceCpu)GetAbsTemp(instant int) (common.DeviceInfo_t) {
	obj := PullObjDeviceAbsTemp(o.Entity, instant)
	return obj
}

func (o *DeviceCpu)SetAbsTemp(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjDeviceAbsTemp(o.Entity, instant, value)
	return obj
}

func (o *DeviceCpu)GetListRelTemp() (map[string]common.DeviceInfo_t) {
	obj := PullObjListDeviceRelTemp(o.Entity)
	return obj
}

func (o *DeviceCpu)GetRelTemp(instant int) (common.DeviceInfo_t) {
	obj := PullObjDeviceRelTemp(o.Entity, instant)
	return obj
}

func (o *DeviceCpu)SetRelTemp(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjDeviceRelTemp(o.Entity, instant, value)
	return obj
}

func (o *DeviceCpu)GetListAveragePower() (map[string]common.DeviceInfo_t) {
	obj := PullObjListDeviceAveragePower(o.Entity)
	return obj
}

func (o *DeviceCpu)GetAveragePower(instant int) (common.DeviceInfo_t) {
	obj := PullObjDeviceAveragePower(o.Entity, instant)
	return obj
}

func (o *DeviceCpu)SetAveragePower(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjDeviceAveragePower(o.Entity, instant, value)
	return obj
}

func (o *DeviceCpu)SetExpectFanDutyByTemperature(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjExpectFanDutyByTemperature(o.Entity, instant, value)
	return obj
}

func (o *DeviceCpu)SetExpectFanDutyByPower(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjExpectFanDutyByPower(o.Entity, instant, value)
	return obj
}

func (o *DeviceCpu)GetMaxPower(instant int) (common.DeviceInfo_t) {
	obj := PullObjDeviceMaxPower(o.Entity, instant)
	return obj
}

func (o *DeviceCpu)GetListMaxPower()(map[string]common.DeviceInfo_t) {
	list := PullObjListDeviceMaxPower(o.Entity)
	return list
}
