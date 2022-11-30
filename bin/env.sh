#!/bin/sh
# Main class name, must be set
export APP_MAIN_CLASS=server-codebase
if [ -z $ENV ]; then
    export ENV="prod"
fi
