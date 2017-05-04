#!/bin/bash
env CGO_ENABLED=0 go build
docker build -t docker/reverse .
docker run -it --rm docker/reverse

