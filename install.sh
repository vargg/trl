#!/bin/bash

set -ex

mkdir -p ~/.trl
cp ./src/conf/settings.yaml ~/.trl/
go build -C src -o ~/.local/bin
