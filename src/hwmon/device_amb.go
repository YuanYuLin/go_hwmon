package hwmon

/*
 *
 */
type DeviceAmb struct {
	Entity int
}

func (o *DeviceAmb)GetListTemperature() (map[string]DeviceInfo_t) {
	obj := PullObjListDeviceTemperature(o.Entity)
	return obj
}

func (o *DeviceAmb)GetTemperature(instant int) (DeviceInfo_t) {
	obj := PullObjDeviceTemperature(o.Entity, instant)
	return obj
}

func (o *DeviceAmb)SetTemperature(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjDeviceTemperature(o.Entity, instant, value)
	return obj
}

func (o *DeviceAmb)SetExpectFanDutyByTemperature(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjExpectFanDutyByTemperature(o.Entity, instant, value)
	return obj
}

func (o *DeviceAmb)SetExpectFanDutyByPower(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjExpectFanDutyByPower(o.Entity, instant, value)
	return obj
}

