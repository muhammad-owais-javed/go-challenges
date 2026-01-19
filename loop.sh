#!/bin/bash

value="$@"


if [[ $value -gt 100 ]]; then
	
	for i in {1..100}
	do
		echo "This is loop number $i"
	done

fi
