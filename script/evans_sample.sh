#!/bin/bash

cat <<EOF | evans --proto ../tutorialpb/tutorial.proto cli call grpc_tutorial.TurtorialService.PrintStr | jq .
{ 
    "message": "ponyo"
}
EOF

cat <<EOF | evans --proto ../tutorialpb/tutorial.proto cli call grpc_tutorial.TurtorialService.GetImageUrl | jq .
{
    "place_id": "01G0V9KFBTV5NR95M5Z2Q7444Q"
}
EOF

cat <<EOF | evans --proto ../tutorialpb/tutorial.proto cli call grpc_tutorial.TurtorialService.SaveColorCode | jq .
{
    "place_id":   "01G0V9KFBTV5NR95M5Z2Q7444Q",
    "user_id":    "PONYO877",
    "color_code": "#000000"
}
EOF

cat <<EOF | evans --proto ../tutorialpb/tutorial.proto cli call grpc_tutorial.TurtorialService.GetColorCode | jq .
{
    "place_id":   "01G0V9KFBTV5NR95M5Z2Q7444Q"
}
EOF