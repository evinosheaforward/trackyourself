#!/bin/bash

docker-compose down

docker build ./recorder -t trackyourself:recorder
