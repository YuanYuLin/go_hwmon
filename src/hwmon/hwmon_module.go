package hwmon

import "fmt"
import "time"

func funcAmbOpenloop() {
	dev_amb := CreateDeviceAmb()
	list_temp := dev_amb.GetListTemperature()
	for _, temp := range list_temp {
		if temp.ValueType == TYPE_RSP_ERROR {
			continue
		}
		val := temp.Value.(float64)
		//fmt.Printf("AMB Temperature %f\n", val)
		if val < 25 {
			dev_amb.SetExpectFanDutyByTemperature(temp.Instant, 20)
		} else if (val >= 25) && (val < 35) {
			dev_amb.SetExpectFanDutyByTemperature(temp.Instant, 40)
		} else {
			dev_amb.SetExpectFanDutyByTemperature(temp.Instant, 60)
		}
	}
}

func funcCpuPid() {
	dev_cpu := CreateDeviceCpu()
	list_temp := dev_cpu.GetListTemperature()
	for _, temp := range list_temp {
		if temp.ValueType == TYPE_RSP_ERROR {
			continue
		}
		val := temp.Value.(float64)
		//fmt.Printf("CPU Temperature %f\n", val)
		if val < 25 {
			dev_cpu.SetExpectFanDutyByTemperature(temp.Instant, 10)
		} else if (val >= 25) && (val < 35) {
			dev_cpu.SetExpectFanDutyByTemperature(temp.Instant, 60)
		} else {
			dev_cpu.SetExpectFanDutyByTemperature(temp.Instant, 80)
		}
	}
}

func funcCpuPowerband() {
	dev_cpu := CreateDeviceCpu()
	list_mp := dev_cpu.GetListMaxPower()
	list_ap := dev_cpu.GetListAveragePower()
	var val_ap float64
	var val_mp float64
	var val float64
	isFound := false

	for _, ap := range list_ap {
		isFound = false
		if ap.ValueType == TYPE_RSP_ERROR {
			continue
		}
		for _, mp := range list_mp {
			if mp.ValueType == TYPE_RSP_ERROR {
				continue
			}
			if (ap.Entity == mp.Entity) && (ap.Instant == mp.Instant) {
				isFound = true
				val_mp = mp.Value.(float64)
				break
			}
		}
		if isFound {
			val_ap = ap.Value.(float64)
			val = val_ap/val_mp
			//fmt.Printf("CPU Power %f/%f=%f\n", val_ap, val_mp, val_ap/val_mp)
			if val < 0.3 {
				dev_cpu.SetExpectFanDutyByPower(ap.Instant, 15)
			} else if (val >= 0.3) && (val < 0.6) {
				dev_cpu.SetExpectFanDutyByPower(ap.Instant, 30)
			} else if (val >= 0.6) && (val < 0.9) {
				dev_cpu.SetExpectFanDutyByPower(ap.Instant, 70)
			} else {
				dev_cpu.SetExpectFanDutyByPower(ap.Instant, 100)
			}
		}
	}
}

func funcFanMap() {
	map_expect_fan_duty := GetMapExpectFanDuty()
	maxobj := DeviceInfo_t{Entity:0, Instant:0, ValueType:"", Value:0.0}
	var maxval float64
	var val float64
	for _, obj := range map_expect_fan_duty {
		//fmt.Printf("[%s]EID:%d, INST:%d, DUTY:%f\n", val.ValueType, val.Entity, val.Instant, val.Value)
		maxval = maxobj.Value.(float64)
		val = obj.Value.(float64)
		if val > maxval {
			maxobj = obj
		}
	}
	maxval = maxobj.Value.(float64)
	if maxval != 0 {
		fmt.Printf("[%s]EID:%d, INST:%d, DUTY:%f\n", maxobj.ValueType, maxobj.Entity, maxobj.Instant, maxobj.Value)
	}
}

const FUNC_STAT_INIT		= 0
const FUNC_STAT_RUNNING		= 1
const FUNC_STAT_EXIT		= 99
type TaskInfo struct {
	Index		int
	Name		string
	FunctionExit	bool
	FunctionStatus	int
	Function	func()
}

func (o *TaskInfo)RunTask() {
	o.FunctionStatus = FUNC_STAT_RUNNING
	for {
		if o.FunctionExit {
			break
		}
		time.Sleep(1 * time.Second)
		o.Function()
	}
	o.FunctionStatus = FUNC_STAT_EXIT
	fmt.Println("Exit " + o.Name)
}

func GetModules() ([]TaskInfo){
	tasks := []TaskInfo {
		{Index:1, Name:"AMB_Openloop",	Function:funcAmbOpenloop,	FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
		{Index:2, Name:"CPU_PID",	Function:funcCpuPid,		FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
		{Index:3, Name:"FAN_Map",	Function:funcFanMap,		FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
		{Index:4, Name:"CPU_Powerband",	Function:funcCpuPowerband,	FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
	}
	return tasks
}

