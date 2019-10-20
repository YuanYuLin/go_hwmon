#!/bin/bash

SRV_HOST="127.0.0.1:8088"
# CPU Relative Temperature
DATA='{"entity":0, "instant":0}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/disable/out/interface
echo ""

