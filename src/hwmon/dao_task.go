package hwmon

import "mailbox"
import "encoding/json"
import "strconv"
import "strings"
import "fmt"

type TaskDao struct {
	db_temperature		map[string][]byte
	db_expectfanduty	map[string][]byte
	db_averagepower		map[string][]byte
	db_maxpower		map[string][]byte
}

func getKey(entity int, instant int, valuetype string) (string) {
	return valuetype + "_" + strconv.Itoa(entity) + "_" + strconv.Itoa(instant)
}

func getKeyValueTypeEntity(entity int, valuetype string) (string) {
	return valuetype + "_" + strconv.Itoa(entity) + "_"
}

func getKeyEntity(entity int) (string) {
	return "_" + strconv.Itoa(entity) + "_"
}

func get_record(msg mailbox.Msg_t, getkey func (int, int, string)(string), db map[string][]byte) (mailbox.Msg_t){
	var res_msg mailbox.Msg_t
	obj := ConvertBytesToDeviceInfo(msg.Data)
	key := getkey(obj.Entity, obj.Instant, obj.ValueType)
	bytes, ok := db[key]
	if ok {
		res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, bytes)
	} else {
		data := DeviceInfo_t { Entity:obj.Entity, Instant:obj.Instant, ValueType:TYPE_RSP_ERROR, Value:"Not found" }
		res_bytes := ConvertDeviceInfoToBytes(data)
		res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, res_bytes)
	}
	return res_msg
}

func get_records(msg mailbox.Msg_t, getkey func (int)(string), db map[string][]byte) (mailbox.Msg_t){
	list := make(map[string]DeviceInfo_t)
	obj := ConvertBytesToDeviceInfo(msg.Data)
	keypart := getkey(obj.Entity)
	for key, bytes := range db {
		if strings.Contains(key, keypart) {
			info := ConvertBytesToDeviceInfo(bytes)
			list[key]=info
		}
	}
	bytes_list := ConvertDeviceInfoListToBytes(list)
	res_msg := mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, bytes_list)
	return res_msg
}

func set_record(msg mailbox.Msg_t, getkey func (int, int, string)(string), db map[string][]byte) (mailbox.Msg_t){
	var res_msg mailbox.Msg_t
	obj := ConvertBytesToDeviceInfo(msg.Data)
	key := getkey(obj.Entity, obj.Instant, obj.ValueType)
	db[key] = msg.Data
	data := DeviceInfo_t { Entity:obj.Entity, Instant:obj.Instant, ValueType:TYPE_RSP_OK, Value:"updated data" }
	res_bytes := ConvertDeviceInfoToBytes(data)
	res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, res_bytes)
	return res_msg
}

func (o* TaskDao)Run() {
	o.db_temperature	= make(map[string][]byte)
	o.db_expectfanduty	= make(map[string][]byte)
	o.db_averagepower	= make(map[string][]byte)
	o.db_maxpower		= make(map[string][]byte)
	mb_dao := mailbox.CreateMailboxDao()
	var res_msg mailbox.Msg_t
	isBreakTask := false
	for {
		msg :=<-mb_dao.Channel

		switch msg.Function {
		case GET_DEVICE_LIST_MAXPOWER:
			res_msg = get_records(msg, getKeyEntity, o.db_maxpower)
		case GET_DEVICE_MAXPOWER:
			res_msg = get_record(msg, getKey, o.db_maxpower)
		case SET_DEVICE_MAXPOWER:
			res_msg = set_record(msg, getKey, o.db_maxpower)

		case GET_DEVICE_LIST_AVERAGEPOWER:
			res_msg = get_records(msg, getKeyEntity, o.db_averagepower)
		case GET_DEVICE_AVERAGEPOWER:
			res_msg = get_record(msg, getKey, o.db_averagepower)
		case SET_DEVICE_AVERAGEPOWER:
			res_msg = set_record(msg, getKey, o.db_averagepower)

		case GET_DEVICE_LIST_TEMPERATURE:
			res_msg = get_records(msg, getKeyEntity, o.db_temperature)
		case GET_DEVICE_TEMPERATURE:
			res_msg = get_record(msg, getKey, o.db_temperature)
		case SET_DEVICE_TEMPERATURE:
			res_msg = set_record(msg, getKey, o.db_temperature)

		case GET_EXPECT_FAN_DUTY:
			res_msg = get_record(msg, getKey, o.db_expectfanduty)
		case SET_EXPECT_FAN_DUTY:
			res_msg = set_record(msg, getKey, o.db_expectfanduty)
		case GET_ALL_EXPECT_FAN_DUTY:
			res_bytes, _ := json.Marshal(o.db_expectfanduty)
			res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, res_bytes)

		case EXIT_APPLICATION:
			isBreakTask = true
			data := DeviceInfo_t { Entity:0, Instant:0, ValueType:TYPE_RSP_EXIT, Value:"Stop task" }
			res_bytes := ConvertDeviceInfoToBytes(data)
			res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, res_bytes)
		default :
			obj := ConvertBytesToDeviceInfo(msg.Data)
			data := DeviceInfo_t { Entity:obj.Entity, Instant:obj.Instant, ValueType:TYPE_RSP_ERROR, Value:"Not Support Operation" }
			res_bytes := ConvertDeviceInfoToBytes(data)
			res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, res_bytes)
		}
		msg.ChannelDst <- res_msg
		if isBreakTask {
			break
		}
	}
	fmt.Println("Exit TaskDao")
}
