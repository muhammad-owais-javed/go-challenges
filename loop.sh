#!/bin/bash

n=$1

if (( $n >100 )); then
	n=100
fi

for i in $(seq 1 $n);
do 
 printf "This is loop number $i\n"
done
