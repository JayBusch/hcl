#!/bin/bash

clang -o c_example c_example.c -L. -l:hcl-c-library.a -lpthread

