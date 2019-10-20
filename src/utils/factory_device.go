package utils

import "config"

func CreateDeviceCpu() (DeviceCpu) {
	obj := DeviceCpu{ Entity:config.ENTITY_CPU }
	return obj
}

func CreateDeviceAmb() (DeviceAmb) {
	obj := DeviceAmb{ Entity:config.ENTITY_AMB }
	return obj
}

func CreateDeviceAic() (DeviceAic) {
	obj := DeviceAic{ Entity:config.ENTITY_AIC }
	return obj
}

func CreateDeviceDimm() (DeviceDimm) {
	obj := DeviceDimm{ Entity:config.ENTITY_DIMM }
	return obj
}
