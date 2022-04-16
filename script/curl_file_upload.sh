#!/bin/bash

FILE=../assets/img/test_auto.jpg
HOST_NAME=localhost:8080/upload

curl -X POST -F file=@$FILE $HOST_NAME