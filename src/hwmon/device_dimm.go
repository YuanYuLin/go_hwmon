package hwmon

/*
 *
 */
type DeviceDimm struct {
	Entity int
}

func (o *DeviceDimm)GetListTemperature() (map[string]DeviceInfo_t) {
	obj := PullObjListDeviceTemperature(o.Entity)
	return obj
}

func (o *DeviceDimm)GetTemperature(instant int) (DeviceInfo_t) {
	obj := PullObjDeviceTemperature(o.Entity, instant)
	return obj
}

func (o *DeviceDimm)SetTemperature(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjDeviceTemperature(o.Entity, instant, value)
	return obj
}

func (o *DeviceDimm)GetListAveragePower() (map[string]DeviceInfo_t) {
	obj := PullObjListDeviceAveragePower(o.Entity)
	return obj
}

func (o *DeviceDimm)GetAveragePower(instant int) (DeviceInfo_t) {
	obj := PullObjDeviceAveragePower(o.Entity, instant)
	return obj
}

func (o *DeviceDimm)SetAveragePower(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjDeviceAveragePower(o.Entity, instant, value)
	return obj
}

func (o *DeviceDimm)SetExpectFanDutyByTemperature(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjExpectFanDutyByTemperature(o.Entity, instant, value)
	return obj
}

func (o *DeviceDimm)SetExpectFanDutyByPower(instant int, value float64) (DeviceInfo_t) {
	obj := PushObjExpectFanDutyByPower(o.Entity, instant, value)
	return obj
}

func (o *DeviceDimm)GetMaxPower(instant int) (DeviceInfo_t) {
	obj := PullObjDeviceMaxPower(o.Entity, instant)
	return obj
}

func (o *DeviceDimm)GetListMaxPower()(map[string]DeviceInfo_t) {
	list := PullObjListDeviceMaxPower(o.Entity)
	return list
}
