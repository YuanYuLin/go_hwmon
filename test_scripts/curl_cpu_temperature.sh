#!/bin/bash

SRV_HOST="127.0.0.1:8088"
# CPU Relative Temperature
DATA='{"entity":1, "instant":1, "value":-50.0}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/set/device/reltemp
echo ""
DATA='{"entity":1, "instant":1}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/get/device/reltemp
echo ""

DATA='{"entity":1, "instant":2, "value":-50.12}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/set/device/reltemp
echo ""
DATA='{"entity":1, "instant":2}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/get/device/reltemp
echo ""

# CPU MAX Temperature
DATA='{"entity":1, "instant":1, "value":100.0}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/set/device/maxtemp
echo ""
DATA='{"entity":1, "instant":1}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/get/device/maxtemp
echo ""


DATA='{"entity":1, "instant":2, "value":100.0}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/set/device/maxtemp
echo ""
DATA='{"entity":1, "instant":2}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/get/device/maxtemp
echo ""

# CPU Absolute Temperature
DATA='{"entity":1, "instant":1, "value":55.12}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/set/device/abstemp
echo ""
DATA='{"entity":1, "instant":1}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/get/device/abstemp
echo ""

DATA='{"entity":1, "instant":2, "value":53.2}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/set/device/abstemp
echo ""
DATA='{"entity":1, "instant":2}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/get/device/abstemp
echo ""

