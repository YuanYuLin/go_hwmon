#!/bin/bash

#curl -d $DATA -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/gettemperature?
DATA='{"entity":1, "instant":1, "value":33.10}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/set/device/temperature
echo ""

DATA='{"entity":1, "instant":2, "value":50.12}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/set/device/temperature
echo ""

DATA='{"entity":2, "instant":1, "value":55.12}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/set/device/temperature
echo ""

DATA='{"entity":3, "instant":1, "value":52.12}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/set/device/temperature
echo ""

DATA='{"entity":1, "instant":1}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/get/device/temperature
echo ""

DATA='{"entity":1, "instant":2}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/get/device/temperature
echo ""

DATA='{"entity":2, "instant":1}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/get/device/temperature
echo ""

DATA='{"entity":3, "instant":1}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/get/device/temperature
echo ""
