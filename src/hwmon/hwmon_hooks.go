package hwmon

import "common"

/*
 * Pre Hook -> DB GET -> Post Hook
 */
func PreGetRecord(dev common.DeviceInfo_t) {
}

func PostGetRecord(isGet bool, devinfo common.DeviceInfo_t) (common.DeviceInfo_t) {
	return devinfo
}

/*
 * Pre Hook -true-> DB SET -> Post Hook
 *     |			^
 *     false	-->	-->	|
 */
func PreSetRecord(devinfo common.DeviceInfo_t) (bool, common.DeviceInfo_t) {
	return true, devinfo
}

func PostSetRecord(isSet bool, devinfo common.DeviceInfo_t)  {
}
