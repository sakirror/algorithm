#!/bin/bash

for i in $(seq 10); do
    number="$number $(echo $RANDOM | cut -c 1-3)"
done

echo "[$number]"

if [ $# -ne 0 ]; then
    $1 $number
else
    for src in $(ls *.c | grep -v main); do
        name=$(echo $src | cut -d '.' -f 1)
        echo $name
        ./bin/$name $number
    done
fi
