package hwmon

import "fmt"
import "time"
import "common"
import "utils"
import "config"

func funcAmbOpenloop(args map[string]interface{}) {
	dev_amb := utils.CreateDeviceAmb()
	list_abstemp := dev_amb.GetListAbsTemp()
	for _, abstemp := range list_abstemp {
		if abstemp.ValueType == config.TYPE_RSP_ERROR {
			continue
		}
		instant := abstemp.Instant
		absval := utils.ToFloat(abstemp.Value)
		key := fmt.Sprintf("FanTable_AMB_%d", instant)
		alg_obj := utils.CreateAlgorithmFanTable(key)
		pwm := alg_obj.Compute(absval)
		utils.SaveAlgorithmToDB(alg_obj.DBKey, alg_obj)
		dev_amb.SetExpectFanDutyByTemperature(instant, pwm)
	}
}

func funcCpuPid(args map[string]interface{}) {
	dev_cpu := utils.CreateDeviceCpu()
	list_reltemp := dev_cpu.GetListRelTemp()
	for _, reltemp := range list_reltemp {
		if reltemp.ValueType == config.TYPE_RSP_ERROR {
			continue
		}
		instant := reltemp.Instant
		relval := utils.ToFloat(reltemp.Value)
		key := fmt.Sprintf("PID_CPU_%d", instant)
		alg_obj := utils.CreateAlgorithmPid(key)
		pwm := alg_obj.Compute(relval)
		utils.SaveAlgorithmToDB(alg_obj.DBKey, alg_obj)
		dev_cpu.SetExpectFanDutyByTemperature(instant, pwm)
	}
}

func funcCpuThreshold(args map[string]interface{}) {
	dev_cpu := utils.CreateDeviceCpu()
	list_reltemp := dev_cpu.GetListRelTemp()
	for _, reltemp := range list_reltemp {
		if reltemp.ValueType == config.TYPE_RSP_ERROR {
			continue
		}
		instant := reltemp.Instant
		relval := utils.ToFloat(reltemp.Value)
		key := fmt.Sprintf("FanTable_CPU_%d", instant)
		alg_obj := utils.CreateAlgorithmFanTable(key)
		pwm := alg_obj.Compute(relval)
		utils.SaveAlgorithmToDB(alg_obj.DBKey, alg_obj)
		dev_cpu.SetExpectFanDutyByTemperature(instant, pwm)
	}
}

func funcCpuPowerband(args map[string]interface{}) {
	dev_cpu := utils.CreateDeviceCpu()
	list_ap := dev_cpu.GetListAveragePower()

	for _, ap := range list_ap {
		if ap.ValueType == config.TYPE_RSP_ERROR {
			continue
		}
		mp := dev_cpu.GetMaxPower(ap.Instant)
		if mp.ValueType == config.TYPE_RSP_ERROR {
			continue
		}
		val_ap := utils.ToFloat(ap.Value)
		val_mp := utils.ToFloat(mp.Value)

		val := val_ap/val_mp
		instant := ap.Instant
		//fmt.Printf("CPU Power %f/%f=%f\n", val_ap, val_mp, val_ap/val_mp)
		if val < 0.3 {
			dev_cpu.SetExpectFanDutyByPower(instant, 15)
		} else if (val >= 0.3) && (val < 0.6) {
			dev_cpu.SetExpectFanDutyByPower(instant, 30)
		} else if (val >= 0.6) && (val < 0.9) {
			dev_cpu.SetExpectFanDutyByPower(instant, 70)
		} else {
			dev_cpu.SetExpectFanDutyByPower(instant, 100)
		}
	}
}

func funcFanMap(args map[string]interface{}) {
	map_expect_fan_duty := utils.GetMapExpectFanDuty()
	maxobj := common.DeviceInfo_t{Entity:0, Instant:0, ValueType:"", Value:0.0}
	var maxval float64
	var val float64
	for _, obj := range map_expect_fan_duty {
		fmt.Printf("[%s]EID:%d, INST:%d, DUTY:%f\n", obj.ValueType, obj.Entity, obj.Instant, obj.Value)
		maxval = utils.ToFloat(maxobj.Value)
		val = utils.ToFloat(obj.Value)
		if val > maxval {
			maxobj = obj
		}
	}
	maxval = utils.ToFloat(maxobj.Value)
	if maxval != 0 {
		//fmt.Printf("[%s]EID:%d, INST:%d, DUTY:%f\n", maxobj.ValueType, maxobj.Entity, maxobj.Instant, maxobj.Value)
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
	Function	func(map[string]interface{})
	FunctionArgs	map[string]interface{}
}

func (o *TaskInfo)RunTask() {
	o.FunctionStatus = FUNC_STAT_RUNNING
	o.FunctionArgs = make(map[string]interface{})
	for {
		if o.FunctionExit {
			break
		}
		time.Sleep(1 * time.Second)
		o.Function(o.FunctionArgs)
	}
	o.FunctionStatus = FUNC_STAT_EXIT
	fmt.Println("Exit " + o.Name)
}

func GetModules() ([]TaskInfo){
	tasks := []TaskInfo {
		{Index:1, Name:"AMB_Openloop",	Function:funcAmbOpenloop,	FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
		{Index:2, Name:"CPU_Threshold",	Function:funcCpuThreshold,	FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
		{Index:3, Name:"FAN_Map",	Function:funcFanMap,		FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
		{Index:4, Name:"CPU_Powerband",	Function:funcCpuPowerband,	FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
		{Index:5, Name:"CPU_PID",	Function:funcCpuPid,		FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
	}
	return tasks
}

