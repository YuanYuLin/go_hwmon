package utils

import "common"
/*
 *
 */
type DeviceDimm struct {
	Entity int
}

func (o *DeviceDimm)GetListAbsTemp() (map[string]common.DeviceInfo_t) {
	obj := PullObjListDeviceAbsTemp(o.Entity)
	return obj
}

func (o *DeviceDimm)GetAbsTemp(instant int) (common.DeviceInfo_t) {
	obj := PullObjDeviceAbsTemp(o.Entity, instant)
	return obj
}

func (o *DeviceDimm)SetAbsTemp(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjDeviceAbsTemp(o.Entity, instant, value)
	return obj
}

func (o *DeviceDimm)GetListAveragePower() (map[string]common.DeviceInfo_t) {
	obj := PullObjListDeviceAveragePower(o.Entity)
	return obj
}

func (o *DeviceDimm)GetAveragePower(instant int) (common.DeviceInfo_t) {
	obj := PullObjDeviceAveragePower(o.Entity, instant)
	return obj
}

func (o *DeviceDimm)SetAveragePower(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjDeviceAveragePower(o.Entity, instant, value)
	return obj
}

func (o *DeviceDimm)SetExpectFanDutyByTemperature(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjExpectFanDutyByTemperature(o.Entity, instant, value)
	return obj
}

func (o *DeviceDimm)SetExpectFanDutyByPower(instant int, value float64) (common.DeviceInfo_t) {
	obj := PushObjExpectFanDutyByPower(o.Entity, instant, value)
	return obj
}

func (o *DeviceDimm)GetMaxPower(instant int) (common.DeviceInfo_t) {
	obj := PullObjDeviceMaxPower(o.Entity, instant)
	return obj
}

func (o *DeviceDimm)GetListMaxPower()(map[string]common.DeviceInfo_t) {
	list := PullObjListDeviceMaxPower(o.Entity)
	return list
}
