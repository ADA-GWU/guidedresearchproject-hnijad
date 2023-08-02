#!/bin/bash

OBJ_PATH="/Users/hnijad/Desktop/lab/read-test/"
BASE_URL="http://127.0.0.1:8081/data"

for ((i=0; i<10000; i++))
do
    id="16,$i"
    curl -s "$BASE_URL/$id" >> /dev/null
done