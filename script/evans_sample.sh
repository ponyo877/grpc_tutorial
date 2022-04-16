#!/bin/bash

cat <<EOF | evans --proto ../tutorialpb/tutorial.proto cli call grpc_tutorial.TurtorialService.PrintStr
{
    "message": "ponyo"
}
EOF