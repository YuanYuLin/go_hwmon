#!/bin/bash

#curl -d $DATA -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/gettemperature?
DATA='{"entity":1, "instant":1, "value":50}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/set/device/averagepower
echo ""

DATA='{"entity":1, "instant":2, "value":40}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/set/device/averagepower
echo ""

DATA='{"entity":1, "instant":1, "value":165}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/set/device/maxpower
echo ""

DATA='{"entity":1, "instant":2, "value":165}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/set/device/maxpower
echo ""

DATA='{"entity":1, "instant":1}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/get/device/averagepower
echo ""

DATA='{"entity":1, "instant":2}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/get/device/averagepower
echo ""

DATA='{"entity":1, "instant":1}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/get/device/maxpower
echo ""

DATA='{"entity":1, "instant":2}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/get/device/maxpower
echo ""

