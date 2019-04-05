#!/bin/sh

for dir in `pwd`/*/
do
    gofmt -s -l ${dir}. \
        && go vet ${dir} \
        && golint -set_exit_status $(go list ${dir}) \
        && go test ${dir}. -race -coverprofile=coverage.txt -covermode=atomic
done
