package common

// Msg of Channel communication
type Msg_t struct {
    Function    string
    ChannelSrc  chan    Msg_t
    ChannelDst  chan    Msg_t
    Data        interface{}
}
// Msg of Rest communication
type JsonMsg_t struct {
    Status	int32		`json:"status"`
    Version	int32		`json:"version"`
    Data	interface{}	`json:"data"`
}

// Data of Channel/Rest communication
type DeviceInfo_t struct {
    Entity      int32		`json:"entity"`
    Instant     int32		`json:"instant"`
    ValueType   int32		`json:"valuetype"`
    Value       interface{}	`json:"value"`
    Key		string		`json:"key"`
}

