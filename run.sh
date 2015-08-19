#!/bin/bash

exec 1> >(logger -t $(basename $0)) 2>&1

./godemo
