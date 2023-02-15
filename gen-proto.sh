#!/usr/bin/env bash

SERVICE_PATH="proto/$1"

if [ -d "$SERVICE_PATH/proto" ]
then
    echo "Directory $SERVICE_PATH/proto existed..."
else
    mkdir $SERVICE_PATH/proto
fi

if [ -d "$SERVICE_PATH/pb" ]
then
    echo "Directory $SERVICE_PATH/pb existed"
else
    mkdir $SERVICE_PATH/pb
fi

if [ -d "$SERVICE_PATH/docs" ]
then
    echo "Directory $SERVICE_PATH/docs existed"
else
    mkdir $SERVICE_PATH/docs
fi

make generate SERVICE_PATH=$SERVICE_PATH
