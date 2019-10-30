package hwmon

import "mailbox"
import "fmt"
import "common"
import "config"

type TaskMsgHndlr struct {
}

func (o* TaskMsgHndlr)Run() {
	mb_msg := mailbox.CreateMailboxMsgHndlr()
	isBreakTask := false
	for {
		msg := <-mb_msg.Channel
		req_msg := common.Msg_t { Function:msg.Function, ChannelSrc:msg.ChannelDst, ChannelDst:msg.ChannelSrc, Data:msg.Data }
		if msg.ChannelDst == nil {
			var res_msg common.Msg_t
			switch msg.Function {
			case config.EXIT_APPLICATION:
				isBreakTask = true
				value := common.ValueResponse_t { Value: config.RESPONSE_OK }
				data := common.DeviceInfo_t { Entity:0, Instant:0, ValueType:config.TYPE_RESPONSE, Value:value }
				//res_bytes := ConvertDeviceInfoToBytes(data)
				res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, data)
			default:
			}
			msg.ChannelSrc <-res_msg
		} else {
			msg.ChannelDst <-req_msg
		}
		if isBreakTask {
			break
		}
	}
	fmt.Println("Exit TaskMsgHndlr")
}
