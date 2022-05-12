#!/bin/bash
CURRENT_DIR=$(pwd)
echo $CURRENT_DIR
for x in $(find ${CURRENT_DIR}/delivery_protos/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/delivery_protos -I /usr/local/include --go_out=plugins=grpc:${CURRENT_DIR} ${x}/*.proto
done