package hwmon

import "encoding/json"
import "mailbox"
//import "fmt"
/*
 *
 */
const ENTITY_CPU	int	= 1
const ENTITY_AMB	int	= 2
const ENTITY_AIC	int	= 3
/*
 *
 */
const GET_DEVICE_LIST_MAXPOWER		string = "get_device_list_maxpower"
const GET_DEVICE_MAXPOWER		string = "get_device_maxpower"
const SET_DEVICE_MAXPOWER		string = "set_device_maxpower"

const GET_DEVICE_LIST_AVERAGEPOWER	string = "get_device_list_averagepower"
const GET_DEVICE_AVERAGEPOWER		string = "get_device_averagepower"
const SET_DEVICE_AVERAGEPOWER		string = "set_device_averagepower"

const GET_DEVICE_LIST_TEMPERATURE	string = "get_device_list_temperature"
const GET_DEVICE_TEMPERATURE		string = "get_device_temperature"
const SET_DEVICE_TEMPERATURE		string = "set_device_temperature"

const GET_EXPECT_FAN_DUTY		string = "get_expect_fan_duty"
const SET_EXPECT_FAN_DUTY		string = "set_expect_fan_duty"
const GET_ALL_EXPECT_FAN_DUTY		string = "get_all_expect_fan_duty"

const EXIT_APPLICATION			string = "exit_application"
/*
 *
 */
const TYPE_TEMPERATURE		string = "temperature"
const TYPE_AVERAGEPOWER		string = "averagepower"
const TYPE_MAXPOWER		string = "maxpower"
// Used in "EXPECT FAN DUTY"
const TYPE_FANDUTY_TEMPERATURE	string = "fdt"
const TYPE_FANDUTY_POWER	string = "fdp"
// Used in "SET" response packet
const TYPE_RSP_OK		string = "rsp_ok"
const TYPE_RSP_ERROR		string = "rsp_error"
//
const TYPE_REQ_EXIT		string = "req_exit"
const TYPE_RSP_EXIT		string = "res_exit"

type DeviceInfo_t struct {
    Entity      int             `json:"entity"`
    Instant     int             `json:"instant"`
    ValueType   string          `json:"valuetype"`
    Value       interface{}     `json:"value"`
}

func ConvertBytesToDeviceInfoList(bytes []byte) (map[string]DeviceInfo_t) {
	var list map[string]DeviceInfo_t
	json.Unmarshal(bytes, &list)
	return list
}

func ConvertDeviceInfoListToBytes(list map[string]DeviceInfo_t)([]byte) {
	bytes, _ := json.Marshal(list)
	return bytes
}

func ConvertBytesToDeviceInfo(bytes []byte) (DeviceInfo_t) {
	var data DeviceInfo_t
	json.Unmarshal(bytes, &data)
	return data
}

func ConvertDeviceInfoToBytes(data DeviceInfo_t) ([]byte) {
	bytes, _ := json.Marshal(data)
	return bytes
}

func TalkToDao(fn string, obj DeviceInfo_t) (mailbox.Msg_t) {
	req_bytes := ConvertDeviceInfoToBytes(obj)

	mb_src := mailbox.CreateMailboxTempBytes()
	mb_dst := mailbox.CreateMailboxDao()
	req_msg := mailbox.Msg_t { Function:fn, ChannelSrc:mb_src.Channel, ChannelDst:mb_dst.Channel, Data: req_bytes }

	mb_msghndlr := mailbox.CreateMailboxMsgHndlr()
	mb_msghndlr.Channel <-req_msg

	res_obj := <-mb_src.Channel
	return res_obj
}

func TalkToHwmon(fn string, obj DeviceInfo_t) (mailbox.Msg_t) {
	req_bytes := ConvertDeviceInfoToBytes(obj)

	mb_src := mailbox.CreateMailboxTempBytes()
	mb_dst := mailbox.CreateMailboxHwmon()
	req_msg := mailbox.Msg_t { Function:fn, ChannelSrc:mb_src.Channel, ChannelDst:mb_dst.Channel, Data: req_bytes }

	mb_msghndlr := mailbox.CreateMailboxMsgHndlr()
	mb_msghndlr.Channel <-req_msg

	res_obj := <-mb_src.Channel
	return res_obj
}

func TalkToRest(fn string, obj DeviceInfo_t) (mailbox.Msg_t) {
	req_bytes := ConvertDeviceInfoToBytes(obj)

	mb_src := mailbox.CreateMailboxTempBytes()
	mb_dst := mailbox.CreateMailboxRest()
	req_msg := mailbox.Msg_t { Function:fn, ChannelSrc:mb_src.Channel, ChannelDst:mb_dst.Channel, Data: req_bytes }

	mb_msghndlr := mailbox.CreateMailboxMsgHndlr()
	mb_msghndlr.Channel <-req_msg

	res_obj := <-mb_src.Channel
	return res_obj
}

func TalkToMsghndlr(fn string, obj DeviceInfo_t) (mailbox.Msg_t) {
	req_bytes := ConvertDeviceInfoToBytes(obj)

	mb_src := mailbox.CreateMailboxTempBytes()
	req_msg := mailbox.Msg_t { Function:fn, ChannelSrc:mb_src.Channel, ChannelDst:nil, Data: req_bytes }

	mb_msghndlr := mailbox.CreateMailboxMsgHndlr()
	mb_msghndlr.Channel <-req_msg

	res_obj := <-mb_src.Channel
	return res_obj
}

