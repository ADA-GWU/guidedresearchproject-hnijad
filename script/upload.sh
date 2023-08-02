#!/bin/bash

OBJ_PATH="/Users/hnijad/Desktop/lab/read-test/"
BASE_URL="http://127.0.0.1:8081/data"

i=0

for file in "$OBJ_PATH"/*; do
  ((i++))
  if [ -f "$file" ]; then
    id="16,$i"
    #echo "Processing file: $(basename $file) id = $id"
    #echo "$BASE_URL/$id"
    curl -s -F "file=@$file" "$BASE_URL/$id" > /dev/null
  fi
done