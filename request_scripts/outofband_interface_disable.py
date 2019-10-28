#!/usr/bin/python2.7

import sys
import utils

def request_list(hostname, out_format):
    json = '{"entity":%d, "instant":%d}' % (-1, -1)
    url='http://%s/api/v1/hwmon/disable/out/interface' % hostname
    utils.response_output(out_format, utils.http_request(url, json))

if __name__ == '__main__':
    if len(sys.argv) < 2:
        utils.help_usage()

    hostname=sys.argv[1]

    request_list(hostname, 'json')

