#!/bin/bash

OBJ_PATH="/Users/hnijad/Desktop/lab/TestObjects"
BASE_URL="http://127.0.0.1:8080/data"

i=10

for file in "$OBJ_PATH"/*; do
  ((i++))
  if [ -f "$file" ]; then
    id="1,$i"
    #echo "Processing file: $(basename $file) id = $id"
    echo "$BASE_URL/$id"
    curl -F "file=@$file" "$BASE_URL/$id" &
  fi
done