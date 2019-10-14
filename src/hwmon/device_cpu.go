package hwmon

/*
 *
 */
type DeviceCpu struct {
	Entity int
}

func (o *DeviceCpu)GetListTemperature() (map[string]DeviceInfo_t) {
	obj := PullObjListDeviceTemperature(o.Entity)
	return obj
}

func (o *DeviceCpu)GetTemperature(instant int) (DeviceInfo_t) {
	obj := PullObjDeviceTemperature(o.Entity, instant)
	return obj
}

func (o *DeviceCpu)SetTemperature(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjDeviceTemperature(o.Entity, instant, value)
	return obj
}

func (o *DeviceCpu)GetListAveragePower() (map[string]DeviceInfo_t) {
	obj := PullObjListDeviceAveragePower(o.Entity)
	return obj
}

func (o *DeviceCpu)GetAveragePower(instant int) (DeviceInfo_t) {
	obj := PullObjDeviceAveragePower(o.Entity, instant)
	return obj
}

func (o *DeviceCpu)SetAveragePower(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjDeviceAveragePower(o.Entity, instant, value)
	return obj
}

func (o *DeviceCpu)SetExpectFanDutyByTemperature(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjExpectFanDutyByTemperature(o.Entity, instant, value)
	return obj
}

func (o *DeviceCpu)SetExpectFanDutyByPower(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjExpectFanDutyByPower(o.Entity, instant, value)
	return obj
}

func (o *DeviceCpu)GetMaxPower(instant int) (DeviceInfo_t) {
	obj := PullObjDeviceMaxPower(o.Entity, instant)
	return obj
}

func (o *DeviceCpu)GetListMaxPower()(map[string]DeviceInfo_t) {
	list := PullObjListDeviceMaxPower(o.Entity)
	return list
}
