#!/bin/bash

tF=$(ls -1 | wc -l)
total=$((tF*5))

printf 'Total files * 5: %d\n' "$total"

