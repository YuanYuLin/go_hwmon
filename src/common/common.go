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
    Status              int     `json:"status"`
    Version             int     `json:"version"`
    Data                interface{} `json:"data"`
}

// Data of Channel/Rest communication
type DeviceInfo_t struct {
    Entity      int             `json:"entity"`
    Instant     int             `json:"instant"`
    Key		string		`json:"key"`
    ValueType   string          `json:"valuetype"`
    Value       interface{}     `json:"value"`
}

