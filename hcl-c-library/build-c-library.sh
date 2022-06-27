#!/bin/bash
rm *.a *.h
go build -buildmode=c-archive -o hcl-c-library.a hcl-c-library.go
cp hcl-c-library.h ./C_example/.
cp hcl-c-library.a ./C_example/.
