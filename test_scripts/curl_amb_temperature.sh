#!/bin/bash

SRV_HOST="127.0.0.1:8088"
# AMB MAX Temperature
DATA='{"entity":2, "instant":1, "value":100.0}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/set/device/maxtemp
echo ""
DATA='{"entity":2, "instant":1}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/get/device/maxtemp
echo ""

# AMB Absolute Temperature
DATA='{"entity":2, "instant":1, "value":25.12}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/set/device/abstemp
echo ""
DATA='{"entity":2, "instant":1}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/get/device/abstemp
echo ""

