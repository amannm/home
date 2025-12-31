#!/usr/bin/env bash
set -eu

go build
chmod +x yxc
mv yxc ~/bin/yxc