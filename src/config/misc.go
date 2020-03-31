package config

var IN_SERVICE_PORT = "localhost:8088"
var OUT_SERVICE_PORT = "0.0.0.0:8080"
/*
 * GET -> 
 *   Successed	: TYPE_xxxx
 *   Failed	: TYPE_REQUEST_ERROR
 * SET ->
 *   Successed	: TYPE_REQUEST_OK
 *   Failed	: TYPE_REQUEST_ERROR
 * COMMAND ->
 *   Successed	: TYPE_REQUEST_OK
 *   Failed	: TYPE_REQUEST_ERROR
 */
const TYPE_OBJECT		int32 = 0x1
const TYPE_TEMPERATURE		int32 = 0x2
const TYPE_AVERAGEPOWER		int32 = 0x3
const TYPE_MAXPOWER		int32 = 0x4
// Used in "EXPECT FAN DUTY"
const TYPE_INITFANDUTY		int32 = 0x5
const TYPE_FANDUTY		int32 = 0x6
const TYPE_DEVICEFANMAP		int32 = 0x7
const TYPE_CPUINFO		int32 = 0x8
// Used in "SET" response packet
const TYPE_REQUEST		int32 = 0x80

const TYPE_REQUEST_OK		int32 = 0x90
const REQUEST_OK		int32 = 0x0

const TYPE_REQUEST_ERROR	int32 = 0xF0
const REQUEST_ERROR_NOT_FOUND	int32 = 0x1
const REQUEST_ERROR_NOT_SET	int32 = 0x2

/*
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
*/
