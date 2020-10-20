#!/usr/bin/env bash

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT="$DIR/../.."

echo "" > $ROOT/coverage.txt

for d in $(go list $ROOT/... | grep -v vendor); do
    go test -race -coverprofile=$ROOT/profile.out -covermode=atomic $d
    if [ -f $ROOT/profile.out ]; then
        cat $ROOT/profile.out >> $ROOT/coverage.txt
        rm $ROOT/profile.out
    fi
done
