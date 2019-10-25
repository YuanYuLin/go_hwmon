package factory

import "config"
import "device"

func CreateDeviceCpu() (device.Cpu_t) {
	obj := device.Cpu_t{ Entity:config.ENTITY_CPU }
	return obj
}

func CreateDeviceAmb() (device.Amb_t) {
	obj := device.Amb_t{ Entity:config.ENTITY_AMB }
	return obj
}

func CreateDeviceAic() (device.Aic_t) {
	obj := device.Aic_t{ Entity:config.ENTITY_AIC }
	return obj
}

func CreateDeviceDimm() (device.Dimm_t) {
	obj := device.Dimm_t{ Entity:config.ENTITY_DIMM }
	return obj
}

func CreateDeviceFan() (device.Fan_t) {
	obj := device.Fan_t{ Entity:config.ENTITY_FAN }
	return obj
}
