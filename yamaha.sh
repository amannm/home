#!/usr/bin/env bash
set -eux


# this blocks, read all devices and identify the instance name
dns-sd -B _http._tcp local

# this blocks, use the instance name identified above to get the hostname
dns-sd -L "Some Instance Name" _http._tcp local
# ... can be reached at Some-Instance-Name.local

# verify connectivity
curl http://Some-Instance-Name.local/YamahaExtendedControl/v1/main/getStatus

# curl -LJO https://community.symcon.de/uploads/short-url/vRXaJXAn6vI2DSQYMHF0aqLbdir.pdf
# split PDF into pages