#!/usr/bin/python2.7

import sys
import utils

def request_list(hostname, out_format):
    entity = utils.ENTITY_CPU
    instant = 1

    json = '{"entity":%d, "instant":%d, "value":{"maxtdp":%d, "cores":%d, "id":%d}}' % (entity, instant, 165, 64, 0x8086)
    url='http://%s/api/v1/hwmon/set/device/cpu/info' % hostname
    utils.response_output(out_format, utils.http_request(url, json))

    instant = 2
    json = '{"entity":%d, "instant":%d, "value":{"maxtdp":%d, "cores":%d, "id":%d}}' % (entity, instant, 165, 64, 0x8086)
    url='http://%s/api/v1/hwmon/set/device/cpu/info' % hostname
    utils.response_output(out_format, utils.http_request(url, json))

if __name__ == '__main__':
    if len(sys.argv) < 2:
        utils.help_usage()

    hostname=sys.argv[1]

    request_list(hostname, 'json')

