package hwmon

import "fmt"
import "time"
import "common"
import "utils"
import "config"
import "factory"
import "sort"

func funcAmbOpenloop(args map[string]interface{}) {
	dev_amb := factory.CreateDeviceAmb()
	list_abstemp := dev_amb.GetListAbsTemp()
	for _, abstemp := range list_abstemp {
		if abstemp.ValueType == config.TYPE_RSP_ERROR {
			continue
		}
		instant := abstemp.Instant
		absval := utils.ToFloat(abstemp.Value)
		key := fmt.Sprintf("FanTable_AMB_%d", instant)
		alg_obj := factory.CreateAlgorithmFanTable(key)
		pwm := alg_obj.Compute(absval)
		factory.SaveAlgorithmToDB(key, alg_obj)
		dev_amb.SetExpectFanDuty(key, instant, pwm)
	}
}

func funcAicThreshold(args map[string]interface{}) {
	dev_aic := factory.CreateDeviceAic()
	list_abstemp := dev_aic.GetListAbsTemp()
	for _, abstemp := range list_abstemp {
		if abstemp.ValueType == config.TYPE_RSP_ERROR {
			continue
		}
		instant := abstemp.Instant
		absval := utils.ToFloat(abstemp.Value)
		key := fmt.Sprintf("FanTable_AIC_%d", instant)
		alg_obj := factory.CreateAlgorithmFanTable(key)
		pwm := alg_obj.Compute(absval)
		factory.SaveAlgorithmToDB(key, alg_obj)
		dev_aic.SetExpectFanDuty(key, instant, pwm)
	}
}

func funcDimmThreshold(args map[string]interface{}) {
	dev_dimm := factory.CreateDeviceDimm()
	list_abstemp := dev_dimm.GetListAbsTemp()
	for _, abstemp := range list_abstemp {
		if abstemp.ValueType == config.TYPE_RSP_ERROR {
			continue
		}
		instant := abstemp.Instant
		absval := utils.ToFloat(abstemp.Value)
		key := fmt.Sprintf("FanTable_DIMM_%d", instant)
		alg_obj := factory.CreateAlgorithmFanTable(key)
		pwm := alg_obj.Compute(absval)
		factory.SaveAlgorithmToDB(key, alg_obj)
		dev_dimm.SetExpectFanDuty(key, instant, pwm)
	}
}

func funcCpuPid(args map[string]interface{}) {
	dev_cpu := factory.CreateDeviceCpu()
	list_reltemp := dev_cpu.GetListRelTemp()
	for _, reltemp := range list_reltemp {
		if reltemp.ValueType == config.TYPE_RSP_ERROR {
			continue
		}
		instant := reltemp.Instant
		relval := utils.ToFloat(reltemp.Value)
		key := fmt.Sprintf("PID_CPU_%d", instant)
		alg_obj := factory.CreateAlgorithmPid(key)
		pwm := alg_obj.Compute(relval)
		factory.SaveAlgorithmToDB(key, alg_obj)
		dev_cpu.SetExpectFanDuty(key, instant, pwm)
	}
}

func funcCpuThreshold(args map[string]interface{}) {
	dev_cpu := factory.CreateDeviceCpu()
	list_reltemp := dev_cpu.GetListRelTemp()
	for _, reltemp := range list_reltemp {
		if reltemp.ValueType == config.TYPE_RSP_ERROR {
			continue
		}
		instant := reltemp.Instant
		relval := utils.ToFloat(reltemp.Value)
		key := fmt.Sprintf("FanTable_CPU_%d", instant)
		alg_obj := factory.CreateAlgorithmFanTable(key)
		pwm := alg_obj.Compute(relval)
		factory.SaveAlgorithmToDB(key, alg_obj)
		dev_cpu.SetExpectFanDuty(key, instant, pwm)
	}
}

func funcCpuPowerband(args map[string]interface{}) {
	dev_cpu := factory.CreateDeviceCpu()
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
		key := fmt.Sprintf("PB_CPU_%d", instant)
		//fmt.Printf("CPU Power %f/%f=%f\n", val_ap, val_mp, val_ap/val_mp)
		if val < 0.3 {
			dev_cpu.SetExpectFanDuty(key, instant, 15)
		} else if (val >= 0.3) && (val < 0.6) {
			dev_cpu.SetExpectFanDuty(key, instant, 30)
		} else if (val >= 0.6) && (val < 0.9) {
			dev_cpu.SetExpectFanDuty(key, instant, 70)
		} else {
			dev_cpu.SetExpectFanDuty(key, instant, 100)
		}
	}
}

func setupFansAndDevices() {
	dev_fan := factory.CreateDeviceFan()

	var defDuty float32
	defDuty = 60.0
	//Set Default Fans Duty
	fanMap := []common.DeviceInfo_t {
		{ Entity: config.ENTITY_FAN, Instant: config.FAN_INSTANT1, Value: defDuty },
		{ Entity: config.ENTITY_FAN, Instant: config.FAN_INSTANT2, Value: defDuty },
		{ Entity: config.ENTITY_FAN, Instant: config.FAN_INSTANT3, Value: defDuty },
	}
	for _, obj := range fanMap {
		fan_duty := utils.ToFloat(obj.Value)
		dev_fan.InitDutyOutput(obj.Instant, fan_duty)
	}

	//Set Devices and Fans Map
	deviceMapFan := []common.DeviceInfo_t {
		{ Entity: config.ENTITY_CPU,	Instant: config.DEV_INSTANT1,	Value: config.FAN_INSTANT1 },
		{ Entity: config.ENTITY_DIMM,	Instant: config.DEV_INSTANT1,	Value: config.FAN_INSTANT1 },
		{ Entity: config.ENTITY_AIC,	Instant: config.DEV_INSTANT1,	Value: config.FAN_INSTANT2 },
		{ Entity: config.ENTITY_CPU,	Instant: config.DEV_INSTANT2,	Value: config.FAN_INSTANT3 },
		{ Entity: config.ENTITY_DIMM,	Instant: config.DEV_INSTANT2,	Value: config.FAN_INSTANT3 },
		{ Entity: config.ENTITY_AMB,	Instant: config.DEV_INSTANT1,	Value: config.FAN_INSTANT1 },
		{ Entity: config.ENTITY_AMB,	Instant: config.DEV_INSTANT1,	Value: config.FAN_INSTANT2 },
		{ Entity: config.ENTITY_AMB,	Instant: config.DEV_INSTANT1,	Value: config.FAN_INSTANT3 },
	}
	for _, obj := range deviceMapFan {
		fan_instant := utils.ToInt(obj.Value)
		dev_fan.SetDeviceMap(obj.Entity, obj.Instant, fan_instant)
		/*
		list := dev_fan.GetDeviceMap(obj.Entity, obj.Instant)
		for _, fm := range list {
			if fm.ValueType == config.TYPE_RSP_ERROR { // Not Found
				fan_instant := utils.ToInt(obj.Value)
				dev_fan.SetDeviceMap(obj.Entity, obj.Instant, fan_instant)
			}
		}
		*/
	}
}

func getSortedKeys(list map[string]common.DeviceInfo_t) ([]string) {
	keys := make([]string, 0, len(list))
	for key := range list {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func funcFanMap(args map[string]interface{}) {
	dev_fan := factory.CreateDeviceFan()
	map_expect_fan_duty := dev_fan.GetAllDevicesExpectFanDuty()
	fmt.Printf("==BEGIN==\n")
	keys1 := getSortedKeys(map_expect_fan_duty)
	for _, key := range keys1 {
		obj := map_expect_fan_duty[key]
		fmt.Printf("[%d]EID:%d, INST:%d, DUTY:%f[%s]\n", obj.ValueType, obj.Entity, obj.Instant, obj.Value, obj.Key)
		list := dev_fan.GetDeviceMap(obj.Entity, obj.Instant)
		for _, fm := range list {
			fmt.Printf("\tmap to fan : %d\n", fm.Value)
			fanInstant := utils.ToInt(-1)
			if fm.ValueType != config.TYPE_RSP_ERROR {
				fanInstant = utils.ToInt(fm.Value)
			}
			out_duty := utils.ToFloat(obj.Value)
			dev := dev_fan.GetDutyOutput(fanInstant)
			db_duty := utils.ToFloat(dev.Value)
			if (dev.ValueType == config.TYPE_RSP_ERROR) || (dev.ValueType == config.TYPE_INITFANDUTY) {
				dev_fan.SetDutyOutput(fanInstant, out_duty)
			} else {
				if db_duty < out_duty {
					dev_fan.SetDutyOutput(fanInstant, out_duty)
				}
			}
		}
	}
	fmt.Printf("==Result==\n")
	fan_list := dev_fan.GetAllDutyOutput()
	keys2 := getSortedKeys(fan_list)
	for _, key := range keys2 {
		obj := fan_list[key]
		fmt.Printf("Fan[%d]=%f[%s-%d]\n", obj.Instant, obj.Value, key, obj.ValueType)
	}
	fmt.Printf("==END==\n")
}

const FUNC_STAT_INIT		= 0
const FUNC_STAT_RUNNING		= 1
const FUNC_STAT_EXIT		= 99
type TaskInfo struct {
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
	setupFansAndDevices()

	tasks := []TaskInfo {
		{Name:"AMB_Openloop",	Function:funcAmbOpenloop,	FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
		{Name:"CPU_Threshold",	Function:funcCpuThreshold,	FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
		{Name:"FAN_Map",	Function:funcFanMap,		FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
		{Name:"CPU_Powerband",	Function:funcCpuPowerband,	FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
		{Name:"CPU_PID",	Function:funcCpuPid,		FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
		{Name:"DIMM_Threshold",	Function:funcDimmThreshold,	FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
		{Name:"AIC_Threshold",	Function:funcAicThreshold,	FunctionExit:false,	FunctionStatus:FUNC_STAT_INIT},
	}
	return tasks
}

