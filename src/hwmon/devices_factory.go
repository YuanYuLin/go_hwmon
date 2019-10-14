package hwmon

func CreateDeviceCpu() (DeviceCpu) {
	obj := DeviceCpu{ Entity:ENTITY_CPU }
	return obj
}

func CreateDeviceAmb() (DeviceAmb) {
	obj := DeviceAmb{ Entity:ENTITY_AMB }
	return obj
}

func CreateDeviceAic() (DeviceAic) {
	obj := DeviceAic{ Entity:ENTITY_AIC }
	return obj
}

func CreateDeviceDimm() (DeviceDimm) {
	obj := DeviceDimm{ Entity:ENTITY_DIMM }
	return obj
}
