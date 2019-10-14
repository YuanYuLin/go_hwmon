package mailbox

import "sync"

type Msg_t struct {
    Function    string
    ChannelSrc  chan    Msg_t
    ChannelDst  chan    Msg_t
    Data        []byte
}

/*
 *
 */
type MailboxRest struct{
	Channel chan Msg_t
}
var mb_rest *MailboxRest
var sync_rest sync.Once

func CreateMailboxRest() (*MailboxRest) {
	sync_rest.Do(func() {
		mb_rest = &MailboxRest{ Channel: make(chan Msg_t) }
	})
	return mb_rest
}

/*
 *
 */
type MailboxDao struct{
	Channel chan Msg_t
}
var mb_dao *MailboxDao
var sync_dao sync.Once

func CreateMailboxDao() (*MailboxDao) {
	sync_dao.Do(func() {
		mb_dao = &MailboxDao{ Channel: make(chan Msg_t) }
	})
	return mb_dao
}

/*
 *
 */
type MailboxMsgHndlr struct{
	Channel chan Msg_t
}
var mb_msghndlr *MailboxMsgHndlr
var sync_msghndlr sync.Once

func CreateMailboxMsgHndlr() (*MailboxMsgHndlr) {
	sync_msghndlr.Do(func() {
		mb_msghndlr = &MailboxMsgHndlr{ Channel: make(chan Msg_t) }
	})
	return mb_msghndlr
}

/*
 *
 */
type MailboxHwmon struct{
	Channel chan Msg_t
}
var mb_hwmon *MailboxHwmon
var sync_hwmon sync.Once

func CreateMailboxHwmon() (*MailboxHwmon) {
	sync_hwmon.Do(func() {
		mb_hwmon = &MailboxHwmon{ Channel: make(chan Msg_t) }
	})
	return mb_hwmon
}

/*
 *
 */
type MailboxTemp struct {
	Channel chan Msg_t
}
func CreateMailboxTempBytes() (*MailboxTemp) {
	mb_temp := new(MailboxTemp)
	mb_temp.Channel = make(chan Msg_t)
	return mb_temp
}

/*
 *
 */
func WrapMsg(fn string, chsrc chan Msg_t, chdst chan Msg_t, data []byte) (Msg_t) {
	msg := Msg_t { Function:fn, ChannelSrc:chsrc, ChannelDst:chdst, Data:data }
	return msg
}
