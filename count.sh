#!/bin/bash

tF=$( find . | wc -l)
total=$(((tF)*5))

printf '\t\vTotal files * 5: %d\v\n' "$total"

