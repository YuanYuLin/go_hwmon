package device

import "common"
import "utils"
/*
 *
 */
type Amb_t struct {
	Entity int32
}

func (o *Amb_t)GetListAbsTemp() (map[string]common.DeviceInfo_t) {
	obj := utils.PullObjListDeviceAbsTemp(o.Entity)
	return obj
}

func (o *Amb_t)GetAbsTemp(instant int32) (common.DeviceInfo_t) {
	obj := utils.PullObjDeviceAbsTemp(o.Entity, instant)
	return obj
}

func (o *Amb_t)SetAbsTemp(instant int32, value float32) (common.DeviceInfo_t) {
	obj := utils.PushObjDeviceAbsTemp(o.Entity, instant, value)
	return obj
}

func (o *Amb_t)SetExpectFanDuty(key string, instant int32, value float32) (common.DeviceInfo_t) {
	obj := utils.PushObjExpectFanDuty(key, o.Entity, instant, value)
	return obj
}
/*
func (o *Amb_t)SetExpectFanDutyByTemperature(instant int32, value float32) (common.DeviceInfo_t) {
	obj := utils.PushObjExpectFanDutyByTemperature(o.Entity, instant, value)
	return obj
}

func (o *Amb_t)SetExpectFanDutyByPower(instant int32, value float32) (common.DeviceInfo_t) {
	obj := utils.PushObjExpectFanDutyByPower(o.Entity, instant, value)
	return obj
}
*/
