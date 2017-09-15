#!/bin/bash
GOOS=darwin
GOARCH=amd64

./buildmaster.sh
./buildcell.sh
#./buildmock.sh
