package device

import "common"
import "utils"
/*
 *
 */
type Aic_t struct {
	Entity int
}

func (o *Aic_t)GetListAbsTemp() (map[string]common.DeviceInfo_t) {
	obj := utils.PullObjListDeviceAbsTemp(o.Entity)
	return obj
}

func (o *Aic_t)GetAbsTemp(instant int) (common.DeviceInfo_t) {
	obj := utils.PullObjDeviceAbsTemp(o.Entity, instant)
	return obj
}

func (o *Aic_t)SetAbsTemp(instant int, value float64) (common.DeviceInfo_t) {
	obj := utils.PushObjDeviceAbsTemp(o.Entity, instant, value)
	return obj
}

func (o *Aic_t)GetListAveragePower() (map[string]common.DeviceInfo_t) {
	obj := utils.PullObjListDeviceAveragePower(o.Entity)
	return obj
}

func (o *Aic_t)GetAveragePower(instant int) (common.DeviceInfo_t) {
	obj := utils.PullObjDeviceAveragePower(o.Entity, instant)
	return obj
}

func (o *Aic_t)SetAveragePower(instant int, value float64) (common.DeviceInfo_t) {
	obj := utils.PushObjDeviceAveragePower(o.Entity, instant, value)
	return obj
}

func (o *Aic_t)SetExpectFanDuty(key string, instant int, value float64) (common.DeviceInfo_t) {
	obj := utils.PushObjExpectFanDuty(key, o.Entity, instant, value)
	return obj
}
/*
func (o *Aic_t)SetExpectFanDutyByTemperature(instant int, value float64) (common.DeviceInfo_t) {
	obj := utils.PushObjExpectFanDutyByTemperature(o.Entity, instant, value)
	return obj
}

func (o *Aic_t)SetExpectFanDutyByPower(instant int, value float64) (common.DeviceInfo_t) {
	obj := utils.PushObjExpectFanDutyByPower(o.Entity, instant, value)
	return obj
}
*/
func (o *Aic_t)GetMaxPower(instant int) (common.DeviceInfo_t) {
	obj := utils.PullObjDeviceMaxPower(o.Entity, instant)
	return obj
}

func (o *Aic_t)GetListMaxPower()(map[string]common.DeviceInfo_t) {
	list := utils.PullObjListDeviceMaxPower(o.Entity)
	return list
}
