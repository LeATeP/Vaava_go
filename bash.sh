#!/bin/bash

for i in {0..10}; do
    go run . &> /dev/null &
    sleep .1
done