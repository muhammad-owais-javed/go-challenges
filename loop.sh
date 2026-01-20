#!/bin/bash

value="$1"

if [[ $# -ne 1 ]]; then
	exit 1
fi
 
if [[ $value -gt 100 ]]; then
	
	for i in {1..100}
	do
		echo "This is loop number $i"
	done

fi

