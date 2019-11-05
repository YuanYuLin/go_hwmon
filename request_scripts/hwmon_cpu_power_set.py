#!/usr/bin/python2.7

import sys
import utils

def request_list(hostname, out_format):
    entity = utils.ENTITY_PROCESSOR
    instant = 1

    json = '{"entity":%d, "instant":%d, "value":%f}' % (entity, instant, 165.0)
    url='http://%s/api/v1/hwmon/set/device/maxpower' % hostname
    utils.response_output(out_format, utils.http_request(url, json))

    json = '{"entity":%d, "instant":%d, "value":%f}' % (entity, instant, 65.0)
    url='http://%s/api/v1/hwmon/set/device/averagepower' % hostname
    utils.response_output(out_format, utils.http_request(url, json))

if __name__ == '__main__':
    if len(sys.argv) < 2:
        utils.help_usage()

    hostname=sys.argv[1]

    request_list(hostname, 'json')

