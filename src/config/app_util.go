package config

var IN_SERVICE_PORT = "localhost:8088"
var OUT_SERVICE_PORT = "0.0.0.0:8080"
/*
 *
 */
const ENTITY_CPU	int	= 1
const ENTITY_AMB	int	= 2
const ENTITY_AIC	int	= 3
const ENTITY_DIMM	int	= 4
/*
 *
 */
const GET_DEVICE_LIST_MAXPOWER		string = "g_dev_list_maxpower"
const GET_DEVICE_MAXPOWER		string = "g_dev_maxpower"
const SET_DEVICE_MAXPOWER		string = "s_dev_maxpower"

const GET_DEVICE_LIST_AVERAGEPOWER	string = "g_dev_list_averagepower"
const GET_DEVICE_AVERAGEPOWER		string = "g_dev_averagepower"
const SET_DEVICE_AVERAGEPOWER		string = "s_dev_averagepower"

const GET_DEVICE_MAXTEMP		string = "g_dev_maxtemp"
const SET_DEVICE_MAXTEMP		string = "s_dev_maxtemp"
const GET_DEVICE_LIST_ABSTEMP		string = "g_dev_list_abstemp"
const GET_DEVICE_ABSTEMP		string = "g_dev_abstemp"
const SET_DEVICE_ABSTEMP		string = "s_dev_abstemp"
const GET_DEVICE_LIST_RELTEMP		string = "g_dev_list_reltemp"
const GET_DEVICE_RELTEMP		string = "g_dev_reltemp"
const SET_DEVICE_RELTEMP		string = "s_dev_reltemp"


const GET_EXPECT_FAN_DUTY		string = "g_expect_fan_duty"
const SET_EXPECT_FAN_DUTY		string = "s_expect_fan_duty"
const GET_ALL_EXPECT_FAN_DUTY		string = "g_all_expect_fan_duty"

const GET_OBJ_BY_KEY			string = "g_o_b_k"
const SET_OBJ_BY_KEY			string = "s_o_b_k"

const EXIT_APPLICATION			string = "exit_app"
const ENABLE_OUTOFBAND_INTERFACE	string = "enable_out_ifc"
const DISABLE_OUTOFBAND_INTERFACE	string = "disable_out_ifc"
/*
 *
 */
const TYPE_OBJECT		string = "object"
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
const TYPE_REQ_CMD		string = "req_cmd"
