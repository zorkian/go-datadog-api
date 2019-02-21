#!/usr/bin/env bash

if ! diff -u <(echo -n) <(gofmt -s -d "$(find . -name '*.go' | grep -v vendor)")
then
    # See if contents have changed and error if they have
    echo "Did you run 'make fmt' before committing?"
    exit 1
fi
