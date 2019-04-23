#!/bin/bash

SCRIPT_DIR=$(cd $(dirname $0) && pwd)


docker-compose -f docker-compose_dev.yml up -d || exit
