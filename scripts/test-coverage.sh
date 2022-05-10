#!/bin/bash
echo 'Running tests'
mkdir -p .test-coverage
docker-compose exec cli \
    go test -coverprofile=.test-coverage/c.out -coverpkg=./...  ./... \
    && go tool cover -html=.test-coverage/c.out -o ./.test-coverage/test-coverage.html

if [[ $? -ne 0 ]]; then 
    # likely the docker wasn't running and some output will be generate by the docker-compose 
    # command, so we won't output anything
    exit 1
fi

if [[ `command -v open` != '' ]]; then
    open .test-coverage/test-coverage.html
else
    echo ''
    echo 'Test coverage report in HTML format generated in .test-coverage/test-coverage.html'
fi