package factory

//import "fmt"
import "utils"
import "algorithm"

func CreateAlgorithmPid(key string) (algorithm.Pid_t) {
	var obj algorithm.Pid_t
	db_obj := utils.PullObj(key)
	if db_obj != nil {
		obj = db_obj.(algorithm.Pid_t)
	} else {
		obj.Kp = 3
		obj.Ki = 2
		obj.Kd = 1

		obj.ClampMin = 0
		obj.ClampMax = 100
		obj.TargetTemperature = 60
		obj.DBKey = key
	}
	return obj
}

func CreateAlgorithmFanTable(key string) (algorithm.LookupTable_t) {
	var obj algorithm.LookupTable_t
	db_obj := utils.PullObj(key)
	if db_obj != nil {
		obj = db_obj.(algorithm.LookupTable_t)
	} else {
		obj.DBKey = key
	}
	return obj
}

func SaveAlgorithmToDB(key string, algorithm interface{}) {
	utils.PushObj(key, algorithm)
}

