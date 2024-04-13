#!/bin/bash

set -ex

mkdir -p ~/.trl
mkdir -p ~/.local/bin
cp ./src/conf/settings.yaml ~/.trl/
go build -C src -o ~/.local/bin
