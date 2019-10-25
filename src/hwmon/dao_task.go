package hwmon

import "common"
import "mailbox"
//import "encoding/json"
import "strconv"
import "strings"
import "config"
import "fmt"
//import "reflect"

type TaskDao struct {
	db_maxtemp		map[string]common.DeviceInfo_t
	db_abstemp		map[string]common.DeviceInfo_t
	db_reltemp		map[string]common.DeviceInfo_t
	db_averagepower		map[string]common.DeviceInfo_t
	db_maxpower		map[string]common.DeviceInfo_t
	db_obj			map[string]common.DeviceInfo_t
	db_expectfanduty	map[string]common.DeviceInfo_t
	db_device_fan_map	map[string]common.DeviceInfo_t
	db_fan_output		map[string]common.DeviceInfo_t
}

func getKeyObj(dev common.DeviceInfo_t) (string) {
	return dev.Key
}

func getKey(obj common.DeviceInfo_t) (string) {
	return obj.ValueType + "_" + strconv.Itoa(obj.Entity) + "_" + strconv.Itoa(obj.Instant)
}

func getKeyValueTypeEntity(obj common.DeviceInfo_t) (string) {
	return obj.ValueType + "_" + strconv.Itoa(obj.Entity) + "_"
}

func getKeyEntity(obj common.DeviceInfo_t) (string){
	return "_" + strconv.Itoa(obj.Entity) + "_"
}

func getKeyEntityInstant(obj common.DeviceInfo_t) (string) {
	return "_" + strconv.Itoa(obj.Entity) + "_" + strconv.Itoa(obj.Instant) + "_"
}

func getKeyEntityInstantValue(obj common.DeviceInfo_t) (string) {
	key := "_" + strconv.Itoa(obj.Entity) + "_" + strconv.Itoa(obj.Instant)
	val := obj.Value
	switch v:=val.(type) {
	case string:
		key = key + "_" + val.(string)
	case int:
		key = fmt.Sprintf("%s_%d", key, val.(int))
	case float64:
		key = fmt.Sprintf("%s_%f", key, val.(float64))
	case float32:
		key = fmt.Sprintf("%s_%f", key, val.(float32))
	default:
		fmt.Print("Value Type : ")
		fmt.Println(v)
	}
	return key
}

func get_record(msg common.Msg_t, getmapkey func (common.DeviceInfo_t)(string), db map[string]common.DeviceInfo_t) (common.Msg_t){
	var res_msg common.Msg_t
	obj := (msg.Data).(common.DeviceInfo_t)
	key := getmapkey(obj)
	dev, ok := db[key]
	if ok {
		res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, dev)
	} else {
		data := common.DeviceInfo_t { Entity:obj.Entity, Instant:obj.Instant, ValueType:config.TYPE_RSP_ERROR, Value:"Not found" }
		res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, data)
	}
	return res_msg
}

func get_records(msg common.Msg_t, getmapkey func (common.DeviceInfo_t)(string), db map[string]common.DeviceInfo_t) (common.Msg_t){
	list := make(map[string]common.DeviceInfo_t)
	obj := (msg.Data).(common.DeviceInfo_t)
	keypart := ""
	if getmapkey != nil {
		keypart = getmapkey(obj)
	}
	for key, dev := range db {
		if strings.Contains(key, keypart) {
			list[key]=dev
		}
	}
	res_msg := mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, list)
	return res_msg
}

func set_record(msg common.Msg_t, getmapkey func (common.DeviceInfo_t)(string), db map[string]common.DeviceInfo_t) (common.Msg_t){
	var res_msg common.Msg_t
	obj := (msg.Data).(common.DeviceInfo_t)
	key := getmapkey(obj)
	db[key] = (msg.Data).(common.DeviceInfo_t)
	data := common.DeviceInfo_t { Entity:obj.Entity, Instant:obj.Instant, ValueType:config.TYPE_RSP_OK, Value:"updated data" }
	res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, data)
	return res_msg
}

func (o* TaskDao)Run() {
	o.db_abstemp		= make(map[string]common.DeviceInfo_t)
	o.db_reltemp		= make(map[string]common.DeviceInfo_t)
	o.db_maxtemp		= make(map[string]common.DeviceInfo_t)
	o.db_averagepower	= make(map[string]common.DeviceInfo_t)
	o.db_maxpower		= make(map[string]common.DeviceInfo_t)
	o.db_obj		= make(map[string]common.DeviceInfo_t)
	o.db_expectfanduty	= make(map[string]common.DeviceInfo_t)
	o.db_device_fan_map	= make(map[string]common.DeviceInfo_t)
	o.db_fan_output		= make(map[string]common.DeviceInfo_t)
	mb_dao := mailbox.CreateMailboxDao()
	var res_msg common.Msg_t
	isBreakTask := false
	for {
		msg :=<-mb_dao.Channel

		switch msg.Function {
		case config.GET_DEVICE_LIST_MAXPOWER:
			res_msg = get_records(msg, getKeyEntity, o.db_maxpower)
		case config.GET_DEVICE_MAXPOWER:
			res_msg = get_record(msg, getKey, o.db_maxpower)
		case config.SET_DEVICE_MAXPOWER:
			res_msg = set_record(msg, getKey, o.db_maxpower)

		case config.GET_DEVICE_LIST_AVERAGEPOWER:
			res_msg = get_records(msg, getKeyEntity, o.db_averagepower)
		case config.GET_DEVICE_AVERAGEPOWER:
			res_msg = get_record(msg, getKey, o.db_averagepower)
		case config.SET_DEVICE_AVERAGEPOWER:
			res_msg = set_record(msg, getKey, o.db_averagepower)

		case config.GET_DEVICE_MAXTEMP:
			res_msg = get_record(msg, getKey, o.db_maxtemp)
		case config.SET_DEVICE_MAXTEMP:
			res_msg = set_record(msg, getKey, o.db_maxtemp)
		case config.GET_DEVICE_LIST_ABSTEMP:
			res_msg = get_records(msg, getKeyEntity, o.db_abstemp)
		case config.GET_DEVICE_ABSTEMP:
			res_msg = get_record(msg, getKey, o.db_abstemp)
		case config.SET_DEVICE_ABSTEMP:
			res_msg = set_record(msg, getKey, o.db_abstemp)
		case config.GET_DEVICE_LIST_RELTEMP:
			res_msg = get_records(msg, getKeyEntity, o.db_reltemp)
		case config.GET_DEVICE_RELTEMP:
			res_msg = get_record(msg, getKey, o.db_reltemp)
		case config.SET_DEVICE_RELTEMP:
			res_msg = set_record(msg, getKey, o.db_reltemp)


		case config.GET_EXPECT_FAN_DUTY:
			res_msg = get_record(msg, getKeyObj, o.db_expectfanduty)
		case config.SET_EXPECT_FAN_DUTY:
			res_msg = set_record(msg, getKeyObj, o.db_expectfanduty)
		case config.GET_ALL_EXPECT_FAN_DUTY:
			res_msg = get_records(msg, nil, o.db_expectfanduty)

		case config.SET_DEVICE_FAN_DUTY_OUTPUT:
			res_msg = set_record(msg, getKeyEntityInstant, o.db_fan_output)
		case config.GET_DEVICE_FAN_DUTY_OUTPUT:
			res_msg = get_record(msg, getKeyEntityInstant, o.db_fan_output)
		case config.GET_ALL_FAN_DUTY_OUTPUT:
			res_msg = get_records(msg, nil, o.db_fan_output)

		case config.GET_DEVICE_FAN_MAP:
			res_msg = get_records(msg, getKeyEntityInstant, o.db_device_fan_map)
		case config.SET_DEVICE_FAN_MAP:
			res_msg = set_record(msg, getKeyEntityInstantValue, o.db_device_fan_map)
		case config.GET_ALL_DEVICE_FAN_MAP:
			res_msg = get_records(msg, nil, o.db_device_fan_map)

		case config.GET_OBJ_BY_KEY:
			res_msg = get_record(msg, getKeyObj, o.db_obj)
		case config.SET_OBJ_BY_KEY:
			res_msg = set_record(msg, getKeyObj, o.db_obj)

		case config.EXIT_APPLICATION:
			isBreakTask = true
			data := common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_RSP_OK, Value:"Stop task" }
			res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, data)
		default :
			obj := (msg.Data).(common.DeviceInfo_t)
			data := common.DeviceInfo_t { Entity:obj.Entity, Instant:obj.Instant, ValueType:config.TYPE_RSP_ERROR, Value:"Not Support Operation" }
			res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, data)
		}
		msg.ChannelDst <- res_msg
		if isBreakTask {
			break
		}
	}
	fmt.Println("Exit TaskDao")
}
