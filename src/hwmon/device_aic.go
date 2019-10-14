package hwmon

/*
 *
 */
type DeviceAic struct {
	Entity int
}

func (o *DeviceAic)GetListTemperature() (map[string]DeviceInfo_t) {
	obj := PullObjListDeviceTemperature(o.Entity)
	return obj
}

func (o *DeviceAic)GetTemperature(instant int) (DeviceInfo_t) {
	obj := PullObjDeviceTemperature(o.Entity, instant)
	return obj
}

func (o *DeviceAic)SetTemperature(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjDeviceTemperature(o.Entity, instant, value)
	return obj
}

func (o *DeviceAic)GetListAveragePower() (map[string]DeviceInfo_t) {
	obj := PullObjListDeviceAveragePower(o.Entity)
	return obj
}

func (o *DeviceAic)GetAveragePower(instant int) (DeviceInfo_t) {
	obj := PullObjDeviceAveragePower(o.Entity, instant)
	return obj
}

func (o *DeviceAic)SetAveragePower(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjDeviceAveragePower(o.Entity, instant, value)
	return obj
}

func (o *DeviceAic)SetExpectFanDutyByTemperature(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjExpectFanDutyByTemperature(o.Entity, instant, value)
	return obj
}

func (o *DeviceAic)SetExpectFanDutyByPower(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjExpectFanDutyByPower(o.Entity, instant, value)
	return obj
}

func (o *DeviceAic)GetMaxPower(instant int) (DeviceInfo_t) {
	obj := PullObjDeviceMaxPower(o.Entity, instant)
	return obj
}

func (o *DeviceAic)GetListMaxPower()(map[string]DeviceInfo_t) {
	list := PullObjListDeviceMaxPower(o.Entity)
	return list
}
