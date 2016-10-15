#!/bin/bash

ROOT=$(dirname "${BASH_SOURCE}")/..

cd $ROOT

godep go build -o app 

cd - > /dev/null
