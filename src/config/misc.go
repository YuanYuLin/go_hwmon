package config

var IN_SERVICE_PORT = "localhost:8088"
var OUT_SERVICE_PORT = "0.0.0.0:8080"
/*
 *
 */
const TYPE_OBJECT		string = "object"
const TYPE_TEMPERATURE		string = "temperature"
const TYPE_AVERAGEPOWER		string = "averagepower"
const TYPE_MAXPOWER		string = "maxpower"
// Used in "EXPECT FAN DUTY"
const TYPE_INITFANDUTY		string = "init_fanduty"
const TYPE_FANDUTY		string = "fanduty"
const TYPE_DEVICEFANMAP		string = "d_f_map"
// Used in "SET" response packet
const TYPE_RSP_OK		string = "rsp_ok"
const TYPE_RSP_ERROR		string = "rsp_error"
//
const TYPE_REQ_CMD		string = "req_cmd"
