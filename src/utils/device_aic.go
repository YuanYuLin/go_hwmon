package utils

import "common"
/*
 *
 */
type DeviceAic struct {
	Entity int
}

func (o *DeviceAic)GetListAbsTemp() (map[string]common.DeviceInfo_t) {
	obj := PullObjListDeviceAbsTemp(o.Entity)
	return obj
}

func (o *DeviceAic)GetAbsTemp(instant int) (common.DeviceInfo_t) {
	obj := PullObjDeviceAbsTemp(o.Entity, instant)
	return obj
}

func (o *DeviceAic)SetAbsTemp(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjDeviceAbsTemp(o.Entity, instant, value)
	return obj
}

func (o *DeviceAic)GetListAveragePower() (map[string]common.DeviceInfo_t) {
	obj := PullObjListDeviceAveragePower(o.Entity)
	return obj
}

func (o *DeviceAic)GetAveragePower(instant int) (common.DeviceInfo_t) {
	obj := PullObjDeviceAveragePower(o.Entity, instant)
	return obj
}

func (o *DeviceAic)SetAveragePower(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjDeviceAveragePower(o.Entity, instant, value)
	return obj
}

func (o *DeviceAic)SetExpectFanDutyByTemperature(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjExpectFanDutyByTemperature(o.Entity, instant, value)
	return obj
}

func (o *DeviceAic)SetExpectFanDutyByPower(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjExpectFanDutyByPower(o.Entity, instant, value)
	return obj
}

func (o *DeviceAic)GetMaxPower(instant int) (common.DeviceInfo_t) {
	obj := PullObjDeviceMaxPower(o.Entity, instant)
	return obj
}

func (o *DeviceAic)GetListMaxPower()(map[string]common.DeviceInfo_t) {
	list := PullObjListDeviceMaxPower(o.Entity)
	return list
}
