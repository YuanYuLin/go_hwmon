package utils

//import "fmt"

func CreateAlgorithmPid(key string) (AlgorithmPid_t) {
	var obj AlgorithmPid_t
	db_obj := PullObj(key)
	if db_obj != nil {
		obj = db_obj.(AlgorithmPid_t)
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

func CreateAlgorithmFanTable(key string) (AlgorithmFanTable_t) {
	var obj AlgorithmFanTable_t
	db_obj := PullObj(key)
	if db_obj != nil {
		obj = db_obj.(AlgorithmFanTable_t)
	} else {
		obj.DBKey = key
	}
	return obj
}

func SaveAlgorithmToDB(key string, algorithm interface{}) {
	PushObj(key, algorithm)
}

