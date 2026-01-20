#!/bin/bash

value="$1"


if [[ $value -gt 100 ]]; then
	
	for i in {1..5}
	do
		echo "This is loop number $i"
	done

fi
