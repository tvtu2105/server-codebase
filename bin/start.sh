#!/bin/sh

this="${BASH_SOURCE-$0}"
while [ -h "$this" ]; do
  ls=$(ls -ld "$this")
  link=$(expr "$ls" : '.*-> \(.*\)$')
  if expr "$link" : '.*/.*' >/dev/null; then
    this="$link"
  else
    this=$(dirname "$this")/"$link"
  fi
done

# convert relative path to absolute path
bin=$(dirname "$this")
script=$(basename "$this")
bin=$(
  cd "$bin" >/dev/null
  pwd
)
this="$bin/$script"
ENV=$1
if [ -z $ENV ]; then
    export ENV="prod"
fi
if [ -z "$APP_HOME" ]; then
  export APP_HOME=$(dirname "$this")/..
fi
APP_CONF_DIR="${APP_CONF_DIR:-$APP_HOME/conf}"

if [ -f "$APP_CONF_DIR/env.sh" ]; then
  . "$APP_CONF_DIR/env.sh"
fi

if [ "$APP_CLASSPATH" != "" ]; then
  CLASSPATH=${CLASSPATH}:${APP_CLASSPATH}
fi

if [ "$APP_LOG_DIR" = "" ]; then
  APP_LOG_DIR="$APP_HOME/logs"
fi

if [ "$APP_LOGFILE" = "" ]; then
  APP_LOGFILE='app.log'
fi

if [ "$APP_DIST" = "" ]; then
  export APP_DIST="$APP_HOME/dist"
fi

APP_OPTS="$APP_OPTS -Dapp.log.dir=$APP_LOG_DIR"
APP_OPTS="$APP_OPTS -Dapp.log.file=$APP_LOGFILE"

source $bin/env.sh
if [ -z ${APP_MAIN_CLASS} ]; then
  echo "You must set app main class name in conf/env.sh"
  exit 1
fi

if [ ! -z ${APP_NAME} ]; then
  APP_OPTS="$APP_OPTS -Dproc_$APP_NAME"
fi

export CLASSPATH

chmod +x $APP_DIST/$APP_MAIN_CLASS
echo $APP_DIST/$APP_MAIN_CLASS $ENV
exec $APP_DIST/$APP_MAIN_CLASS $ENV
