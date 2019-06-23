#!/bin/sh
#
# Set GOPATH to script dir.
# Usage: source setgo.sh

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export GOPATH=$DIR
