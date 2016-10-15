#!/bin/bash

ROOT=$(dirname "${BASH_SOURCE}")/..

cd $ROOT

godep go build -o app 
docker build -t gaocegege/hackys-writer .
docker push gaocegege/hackys-writer
cd - > /dev/null
