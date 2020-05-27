#!/bin/sh

rm -r ./gen
mkdir -p ./gen/avro
gogen-avro ./gen/avro schema.avro
go test

