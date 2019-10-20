#!/bin/bash

#curl -d $DATA -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/gettemperature?
DATA='{"entity":0, "instant":0}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/api/v1/hwmon/exit/main
echo ""

