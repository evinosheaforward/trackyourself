#!/bin/bash

MYPWD="$(pwd)"

docker run -it --env-file=".env" -p 9299:9299 -v $MYPWD/data:/data trackyourself:recorder
