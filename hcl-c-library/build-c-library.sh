#!/bin/bash
rm *.a *.h
go build -buildmode=c-archive -o hcl-c-library.a hcl-c-library.go
