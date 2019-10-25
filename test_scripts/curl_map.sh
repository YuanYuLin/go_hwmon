#!/bin/bash

SRV_HOST="127.0.0.1:8088"
DATA='{"entity":-1, "instant":-1}'

curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/get/map/alldevicefan | json_pp
echo ""

curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/get/map/allfandutyout | json_pp
echo ""

curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/get/map/allexpectduty | json_pp
echo ""
