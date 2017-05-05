#!/bin/bash
#env CGO_ENABLED=0 go build
docker build -t docker/reverse .
docker run -it --rm docker/reverse
docker tag docker/reverse:latest "$AWS_ACCOUNT.dkr.ecr.us-east-1.amazonaws.com/docker/reverse:latest"
docker push "$AWS_ACCOUNT.dkr.ecr.us-east-1.amazonaws.com/docker/reverse:latest"

