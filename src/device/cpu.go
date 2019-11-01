package device

import "common"
import "utils"
/*
 *
 */
type Cpu_t struct {
	Entity int32
}
/*
func (o *Cpu_t)GetListAbsTemp() (map[string]common.DeviceInfo_t) {
	obj := utils.PullObjListDeviceAbsTemp(o.Entity)
	return obj
}

func (o *Cpu_t)GetAbsTemp(instant int32) (common.DeviceInfo_t) {
	obj := utils.PullObjDeviceAbsTemp(o.Entity, instant)
	return obj
}

func (o *Cpu_t)SetAbsTemp(instant int32, value float32) (common.DeviceInfo_t) {
	obj := utils.PushObjDeviceAbsTemp(o.Entity, instant, value)
	return obj
}
*/
func (o *Cpu_t)GetListRelTemp() (map[string]common.DeviceInfo_t) {
	obj := utils.PullObjListDeviceRelTemp(o.Entity)
	return obj
}

func (o *Cpu_t)GetRelTemp(instant int32) (common.DeviceInfo_t) {
	obj := utils.PullObjDeviceRelTemp(o.Entity, instant)
	return obj
}

func (o *Cpu_t)SetRelTemp(instant int32, value float32) (common.DeviceInfo_t) {
	obj := utils.PushObjDeviceRelTemp(o.Entity, instant, value)
	return obj
}

func (o *Cpu_t)GetListAveragePower() (map[string]common.DeviceInfo_t) {
	obj := utils.PullObjListDeviceAveragePower(o.Entity)
	return obj
}

func (o *Cpu_t)GetAveragePower(instant int32) (common.DeviceInfo_t) {
	obj := utils.PullObjDeviceAveragePower(o.Entity, instant)
	return obj
}

func (o *Cpu_t)SetAveragePower(instant int32, value float32) (common.DeviceInfo_t) {
	obj := utils.PushObjDeviceAveragePower(o.Entity, instant, value)
	return obj
}

func (o *Cpu_t)SetExpectFanDuty(key string, instant int32, value float32) (common.DeviceInfo_t) {
	obj := utils.PushObjExpectFanDuty(key, o.Entity, instant, value)
	return obj
}
/*
func (o *Cpu_t)SetExpectFanDutyByTemperature(instant int32, key string, value float32) (common.DeviceInfo_t) {
	obj := utils.PushObjExpectFanDutyByTemperature(o.Entity, instant, key, value)
	return obj
}

func (o *Cpu_t)SetExpectFanDutyByPower(instant int32, value float32) (common.DeviceInfo_t) {
	obj := utils.PushObjExpectFanDutyByPower(o.Entity, instant, value)
	return obj
}
*/
func (o *Cpu_t)GetMaxPower(instant int32) (common.DeviceInfo_t) {
	obj := utils.PullObjDeviceMaxPower(o.Entity, instant)
	return obj
}

func (o *Cpu_t)GetListMaxPower()(map[string]common.DeviceInfo_t) {
	list := utils.PullObjListDeviceMaxPower(o.Entity)
	return list
}
