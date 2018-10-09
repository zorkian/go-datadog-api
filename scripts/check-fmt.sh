#!/usr/bin/env bash

diff -u <(echo -n) <(gofmt -s -d $(find . -name '*.go' | grep -v vendor))

# See if contents have changed and error if they have
if [[ $? != 0 ]] ; then
    echo "Did you run 'make fmt' before committing?"
    exit 1
fi
