package hwmon

import "common"
import "mailbox"
import "strings"
import "config"
import "fmt"

type TaskDao struct {
	db_maxtemp			map[string]common.DeviceInfo_t
	db_abstemp			map[string]common.DeviceInfo_t
	db_reltemp			map[string]common.DeviceInfo_t
	db_averagepower		map[string]common.DeviceInfo_t
	db_maxpower			map[string]common.DeviceInfo_t
	db_obj				map[string]common.DeviceInfo_t
	db_expectfanduty	map[string]common.DeviceInfo_t
	db_device_fan_map	map[string]common.DeviceInfo_t
	db_fan_output		map[string]common.DeviceInfo_t
}

func get_hdr(data interface{}) (int32, int32, string) {
	var entity int32
	var instant int32
	var key string
	switch data.(type) {
		case common.DeviceInfo_t:
		req := data.(common.DeviceInfo_t)
		key = req.Key
		entity = req.Entity
		instant = req.Instant
	}
	return entity, instant, key
}

func get_record(msg common.Msg_t, db map[string]common.DeviceInfo_t) (common.Msg_t) {
	var res_msg common.Msg_t
	PreGetRecord((msg.Data).(common.DeviceInfo_t))
	entity, instant, key := get_hdr(msg.Data)
	data, ok := db[key]
	if !ok {
		value := common.ValueResponse_t { Value: config.REQUEST_ERROR_NOT_FOUND }
		data = common.DeviceInfo_t { Entity:entity, Instant:instant, Key: key, ValueType:config.TYPE_REQUEST_ERROR, Value:value }
	}
	data = PostGetRecord(ok, data)
	res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, data)

	return res_msg
}
func set_record(msg common.Msg_t, db map[string]common.DeviceInfo_t) (common.Msg_t){
	var res_msg common.Msg_t
	ok, dev_info := PreSetRecord((msg.Data).(common.DeviceInfo_t))
	entity, instant, key := get_hdr(dev_info)
	data := dev_info
	if ok {
		db[key] = dev_info//(msg.Data).(common.DeviceInfo_t)
	} else {
		value := common.ValueResponse_t { Value: config.REQUEST_ERROR_NOT_SET }
		data = common.DeviceInfo_t { Entity:entity, Instant:instant, Key:key, ValueType:config.TYPE_REQUEST_ERROR, Value:value }
	}
	PostSetRecord(ok, dev_info)

	res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, data)
	return res_msg
}

func get_records(msg common.Msg_t, db map[string]common.DeviceInfo_t) (common.Msg_t){
	list := make(map[string]common.DeviceInfo_t)
	_, _, keypart := get_hdr(msg.Data)
	for key, dev := range db {
		if strings.Contains(key, keypart) {
			list[key]=dev
		}
	}
	res_msg := mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, list)
	return res_msg
}

func (o* TaskDao)Run() {
	o.db_maxtemp		= make(map[string]common.DeviceInfo_t)
	o.db_abstemp		= make(map[string]common.DeviceInfo_t)
	o.db_reltemp		= make(map[string]common.DeviceInfo_t)
	o.db_averagepower	= make(map[string]common.DeviceInfo_t)
	o.db_maxpower		= make(map[string]common.DeviceInfo_t)
	o.db_obj			= make(map[string]common.DeviceInfo_t)
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
			res_msg = get_records(msg, o.db_maxpower)
		case config.GET_DEVICE_MAXPOWER:
			res_msg = get_record(msg, o.db_maxpower)
		case config.SET_DEVICE_MAXPOWER:
			res_msg = set_record(msg, o.db_maxpower)

		case config.GET_DEVICE_LIST_AVERAGEPOWER:
			res_msg = get_records(msg, o.db_averagepower)
		case config.GET_DEVICE_AVERAGEPOWER:
			res_msg = get_record(msg, o.db_averagepower)
		case config.SET_DEVICE_AVERAGEPOWER:
			res_msg = set_record(msg, o.db_averagepower)

		case config.GET_DEVICE_MAXTEMP:
			res_msg = get_record(msg, o.db_maxtemp)
		case config.SET_DEVICE_MAXTEMP:
			res_msg = set_record(msg, o.db_maxtemp)
		case config.GET_DEVICE_LIST_ABSTEMP:
			res_msg = get_records(msg, o.db_abstemp)
		case config.GET_DEVICE_ABSTEMP:
			res_msg = get_record(msg, o.db_abstemp)
		case config.SET_DEVICE_ABSTEMP:
			res_msg = set_record(msg, o.db_abstemp)
		case config.GET_DEVICE_LIST_RELTEMP:
			res_msg = get_records(msg, o.db_reltemp)
		case config.GET_DEVICE_RELTEMP:
			res_msg = get_record(msg, o.db_reltemp)
		case config.SET_DEVICE_RELTEMP:
			res_msg = set_record(msg, o.db_reltemp)

		case config.GET_EXPECT_FAN_DUTY:
			res_msg = get_record(msg, o.db_expectfanduty)
		case config.SET_EXPECT_FAN_DUTY:
			res_msg = set_record(msg, o.db_expectfanduty)
		case config.GET_ALL_EXPECT_FAN_DUTY:
			res_msg = get_records(msg, o.db_expectfanduty)

		case config.SET_DEVICE_FAN_DUTY_OUTPUT:
			res_msg = set_record(msg, o.db_fan_output)
		case config.GET_DEVICE_FAN_DUTY_OUTPUT:
			res_msg = get_record(msg, o.db_fan_output)
		case config.GET_ALL_FAN_DUTY_OUTPUT:
			res_msg = get_records(msg, o.db_fan_output)

		case config.GET_DEVICE_FAN_MAP:
			res_msg = get_records(msg, o.db_device_fan_map)
		case config.SET_DEVICE_FAN_MAP:
			res_msg = set_record(msg, o.db_device_fan_map)
		case config.GET_ALL_DEVICE_FAN_MAP:
			res_msg = get_records(msg, o.db_device_fan_map)

		case config.GET_OBJ_BY_KEY:
			res_msg = get_record(msg, o.db_obj)
		case config.SET_OBJ_BY_KEY:
			res_msg = set_record(msg, o.db_obj)

		case config.EXIT_APPLICATION:
			isBreakTask = true
			value := common.ValueResponse_t { Value : config.REQUEST_OK }
			data := common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_REQUEST_OK, Value:value }
			res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, data)
		default :
			entity, instant, key := get_hdr(msg.Data)
			value := common.ValueResponse_t { Value : config.REQUEST_ERROR_NOT_FOUND }
			data := common.DeviceInfo_t { Entity:entity, Instant:instant, Key: key, ValueType:config.TYPE_REQUEST_ERROR, Value:value }
			res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, data)
		}
		msg.ChannelDst <- res_msg
		if isBreakTask {
			break
		}
	}
	fmt.Println("Exit TaskDao")
}
