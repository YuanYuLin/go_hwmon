package utils

import "encoding/json"
import "common"
import "mailbox"

func IsResponse(data interface{})(bool) {
	_, ok := data.(common.DeviceInfo_t)
	return ok
}

func ToResponse(data interface{})(common.DeviceInfo_t) {
	res := data.(common.DeviceInfo_t)
	return res
}

func GetHeaders(data interface{})(int32, int32, string) {
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

func ToFloat(val interface{}) (float32) {
	var result float32
	result = -1
	switch val.(type) {
	case float64:
		result = float32(val.(float64))
	case float32:
		result = val.(float32)
	}
	return result
}

func ToInt(val interface{}) (int32) {
	var result int32
	result = -1
	switch val.(type) {
	case int64:
		result = int32(val.(int64))
	case int32:
		result = val.(int32)
	}
	return result
}

func ConvertBytesToDeviceInfo(bytes []byte) (common.DeviceInfo_t) {
	var data common.DeviceInfo_t
	json.Unmarshal(bytes, &data)
	return data
}

func TalkToDao(fn string, obj interface{}) (common.Msg_t) {
	mb_src := mailbox.CreateMailboxTemporary()
	mb_dst := mailbox.CreateMailboxDao()
	req_msg := common.Msg_t { Function:fn, ChannelSrc:mb_src.Channel, ChannelDst:mb_dst.Channel, Data: obj }

	mb_msghndlr := mailbox.CreateMailboxMsgHndlr()
	mb_msghndlr.Channel <-req_msg

	res_obj := <-mb_src.Channel
	return res_obj
}

func TalkToHwmon(fn string, obj interface{}) (common.Msg_t) {
	mb_src := mailbox.CreateMailboxTemporary()
	mb_dst := mailbox.CreateMailboxHwmon()
	req_msg := common.Msg_t { Function:fn, ChannelSrc:mb_src.Channel, ChannelDst:mb_dst.Channel, Data: obj }

	mb_msghndlr := mailbox.CreateMailboxMsgHndlr()
	mb_msghndlr.Channel <-req_msg

	res_obj := <-mb_src.Channel
	return res_obj
}

func TalkToRest(fn string, obj interface{}) (common.Msg_t) {
	mb_src := mailbox.CreateMailboxTemporary()
	mb_dst := mailbox.CreateMailboxRest()
	req_msg := common.Msg_t { Function:fn, ChannelSrc:mb_src.Channel, ChannelDst:mb_dst.Channel, Data: obj }

	mb_msghndlr := mailbox.CreateMailboxMsgHndlr()
	mb_msghndlr.Channel <-req_msg

	res_obj := <-mb_src.Channel
	return res_obj
}

func TalkToMsghndlr(fn string, obj interface{}) (common.Msg_t) {
	mb_src := mailbox.CreateMailboxTemporary()
	req_msg := common.Msg_t { Function:fn, ChannelSrc:mb_src.Channel, ChannelDst:nil, Data: obj }

	mb_msghndlr := mailbox.CreateMailboxMsgHndlr()
	mb_msghndlr.Channel <-req_msg

	res_obj := <-mb_src.Channel
	return res_obj
}

