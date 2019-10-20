package utils

import "encoding/json"
import "common"
import "mailbox"

func ToFloat(val interface{}) (float64) {
	return val.(float64)
}

func ConvertBytesToDeviceInfo(bytes []byte) (common.DeviceInfo_t) {
	var data common.DeviceInfo_t
	json.Unmarshal(bytes, &data)
	return data
}

func TalkToDao(fn string, obj common.DeviceInfo_t) (common.Msg_t) {
	mb_src := mailbox.CreateMailboxTemporary()
	mb_dst := mailbox.CreateMailboxDao()
	req_msg := common.Msg_t { Function:fn, ChannelSrc:mb_src.Channel, ChannelDst:mb_dst.Channel, Data: obj }

	mb_msghndlr := mailbox.CreateMailboxMsgHndlr()
	mb_msghndlr.Channel <-req_msg

	res_obj := <-mb_src.Channel
	return res_obj
}

func TalkToHwmon(fn string, obj common.DeviceInfo_t) (common.Msg_t) {
	mb_src := mailbox.CreateMailboxTemporary()
	mb_dst := mailbox.CreateMailboxHwmon()
	req_msg := common.Msg_t { Function:fn, ChannelSrc:mb_src.Channel, ChannelDst:mb_dst.Channel, Data: obj }

	mb_msghndlr := mailbox.CreateMailboxMsgHndlr()
	mb_msghndlr.Channel <-req_msg

	res_obj := <-mb_src.Channel
	return res_obj
}

func TalkToRest(fn string, obj common.DeviceInfo_t) (common.Msg_t) {
	mb_src := mailbox.CreateMailboxTemporary()
	mb_dst := mailbox.CreateMailboxRest()
	req_msg := common.Msg_t { Function:fn, ChannelSrc:mb_src.Channel, ChannelDst:mb_dst.Channel, Data: obj }

	mb_msghndlr := mailbox.CreateMailboxMsgHndlr()
	mb_msghndlr.Channel <-req_msg

	res_obj := <-mb_src.Channel
	return res_obj
}

func TalkToMsghndlr(fn string, obj common.DeviceInfo_t) (common.Msg_t) {
	mb_src := mailbox.CreateMailboxTemporary()
	req_msg := common.Msg_t { Function:fn, ChannelSrc:mb_src.Channel, ChannelDst:nil, Data: obj }

	mb_msghndlr := mailbox.CreateMailboxMsgHndlr()
	mb_msghndlr.Channel <-req_msg

	res_obj := <-mb_src.Channel
	return res_obj
}

