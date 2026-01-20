#!/bin/bash

if [[ $# -ne 1 ]]; then
	exit 1
fi
 
if [[ "$1" = *[!0-9]* ]]; then
	exit 1
fi

value=$1

if [[ $value -gt 100 ]]; then
	
	for i in {1..100}
	do
		echo "This is loop number $i"
	done

fi

