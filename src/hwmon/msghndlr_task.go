package hwmon

import "mailbox"
import "fmt"

type TaskMsgHndlr struct {
}

func (o* TaskMsgHndlr)Run() {
	mb_msg := mailbox.CreateMailboxMsgHndlr()
	isBreakTask := false
	for {
		msg := <-mb_msg.Channel
		req_msg := mailbox.Msg_t { Function:msg.Function, ChannelSrc:msg.ChannelDst, ChannelDst:msg.ChannelSrc, Data:msg.Data }
		if msg.ChannelDst == nil {
			var res_msg mailbox.Msg_t
			switch msg.Function {
			case EXIT_APPLICATION:
				isBreakTask = true
				data := DeviceInfo_t { Entity:0, Instant:0, ValueType:TYPE_RSP_EXIT, Value:"Stop task" }
				res_bytes := ConvertDeviceInfoToBytes(data)
				res_msg = mailbox.WrapMsg(msg.Function, msg.ChannelSrc, msg.ChannelDst, res_bytes)
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
