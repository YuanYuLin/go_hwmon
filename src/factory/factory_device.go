package factory

import "config"
import "device"

func CreateDeviceCpu() (device.Cpu_t) {
	obj := device.Cpu_t{ Entity:config.ENTITY_PROCESSOR }
	return obj
}

func CreateDeviceAmb() (device.Amb_t) {
	obj := device.Amb_t{ Entity:config.ENTITY_EXTERNAL_ENVIROMENT }
	return obj
}

func CreateDeviceAic() (device.Aic_t) {
	obj := device.Aic_t{ Entity:config.ENTITY_ADD_IN_CARD }
	return obj
}

func CreateDeviceDimm() (device.Dimm_t) {
	obj := device.Dimm_t{ Entity:config.ENTITY_MEMORY_DEVICE }
	return obj
}

func CreateDeviceFan() (device.Fan_t) {
	obj := device.Fan_t{ Entity:config.ENTITY_FAN_COOLING }
	return obj
}
