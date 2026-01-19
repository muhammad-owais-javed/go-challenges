#!/bin/bash

tF=$( ls -R | wc -w)
total=$(((tF-2)*5))

printf '\t\vTotal files * 5: %d\v\n' "$total"

