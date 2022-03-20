#!/bin/sh
echo \"`cat /dev/urandom | tr -dc 'a-z
A-Z0-9' | fold -w 36 | head -n 1 | sort | uniq`\" | 
sed -e '1iSALT=' | tr -d '\n' > salt.env 