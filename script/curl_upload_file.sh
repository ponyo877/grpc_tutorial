#!/bin/bash

USER_ID=PONYO877
FILE=../assets/img/test_auto.jpg
COLOR_CODE="#000001"
HOST_NAME=localhost:8080

curl -X POST -F userId=$USER_ID -F file=@$FILE $HOST_NAME/save_image_unary/01G0V9KFBTV5NR95M5Z2Q7444Q | jq .
# curl -X POST -F userId=$USER_ID -F file=@$FILE $HOST_NAME/save_image/01G0V9KFBTV5NR95M5Z2Q7444Q | jq .
curl -X GET $HOST_NAME/get_image_url/01G0V9KFBTV5NR95M5Z2Q7444Q | jq .
curl -X POST -F userId=$USER_ID -F colorCode=$COLOR_CODE $HOST_NAME/save_color_code/01G0V9KFBTV5NR95M5Z2Q7444Q | jq .
curl -X GET $HOST_NAME/get_color_code/01G0V9KFBTV5NR95M5Z2Q7444Q | jq .