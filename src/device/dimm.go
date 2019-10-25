package device

import "common"
import "utils"
/*
 *
 */
type Dimm_t struct {
	Entity int
}

func (o *Dimm_t)GetListAbsTemp() (map[string]common.DeviceInfo_t) {
	obj := utils.PullObjListDeviceAbsTemp(o.Entity)
	return obj
}

func (o *Dimm_t)GetAbsTemp(instant int) (common.DeviceInfo_t) {
	obj := utils.PullObjDeviceAbsTemp(o.Entity, instant)
	return obj
}

func (o *Dimm_t)SetAbsTemp(instant int, value float64) (common.DeviceInfo_t) {
	obj := utils.PushObjDeviceAbsTemp(o.Entity, instant, value)
	return obj
}

func (o *Dimm_t)GetListAveragePower() (map[string]common.DeviceInfo_t) {
	obj := utils.PullObjListDeviceAveragePower(o.Entity)
	return obj
}

func (o *Dimm_t)GetAveragePower(instant int) (common.DeviceInfo_t) {
	obj := utils.PullObjDeviceAveragePower(o.Entity, instant)
	return obj
}

func (o *Dimm_t)SetAveragePower(instant int, value float64) (common.DeviceInfo_t) {
	obj := utils.PushObjDeviceAveragePower(o.Entity, instant, value)
	return obj
}

func (o *Dimm_t)SetExpectFanDuty(key string, instant int, value float64) (common.DeviceInfo_t) {
	obj := utils.PushObjExpectFanDuty(key, o.Entity, instant, value)
	return obj
}
/*
func (o *Dimm_t)SetExpectFanDutyByTemperature(instant int, value float64) (common.DeviceInfo_t) {
	obj := utils.PushObjExpectFanDutyByTemperature(o.Entity, instant, value)
	return obj
}

func (o *Dimm_t)SetExpectFanDutyByPower(instant int, value float64) (common.DeviceInfo_t) {
	obj := utils.PushObjExpectFanDutyByPower(o.Entity, instant, value)
	return obj
}
*/
func (o *Dimm_t)GetMaxPower(instant int) (common.DeviceInfo_t) {
	obj := utils.PullObjDeviceMaxPower(o.Entity, instant)
	return obj
}

func (o *Dimm_t)GetListMaxPower()(map[string]common.DeviceInfo_t) {
	list := utils.PullObjListDeviceMaxPower(o.Entity)
	return list
}
