import requests
import pprint
import json

ENTITY_CPU  = 1
ENTITY_AMB  = 2
ENTITY_AIC  = 3
ENTITY_DIMM = 4
ENTITY_FAN  = 5

TYPE_OBJECT	    = 0x1
TYPE_TEMPERATURE    = 0x2
TYPE_AVERAGEPOWER   = 0x3
TYPE_MAXPOWER	    = 0x4
TYPE_INITFANDUTY    = 0x5
TYPE_FANDUTY	    = 0x6
TYPE_DEVICEFANMAP   = 0x7

TYPE_RSP_OK	    = 0x0
TYPE_RSP_ERROR	    = 0xF0
TYPE_CMD	    = 0x80

def http_request(url, payload):
    hds = {'content-type':'application/json; charset=utf-8', 'user-agent':'iopc-app'}
    rsp=requests.post(url, headers=hds, data=payload)
    return rsp

def response_output(out_format, rsp):
    if rsp.status_code == 200 :
        pprint.pprint(rsp.json())

def help_usage():
    print "hwmon.py <hostname>"
    sys.exit(1)

