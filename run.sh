#!/bin/bash

BASEDIR=$(dirname "$0")

fswatch "$BASEDIR" | while read file
do
    filename=$(basename -- "$file")
    extension="${filename##*.}"
    filename="${filename%.*}"

    if [ "$extension" = "go" ]
    then
        echo "=============================================\n\n"
        time go run ${file}
    fi
done

