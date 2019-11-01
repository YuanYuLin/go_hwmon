package hwmon

import "common"

func PreSetRecord(key string, devinfo common.DeviceInfo_t) (bool, common.DeviceInfo_t) {
	return true, devinfo
}
