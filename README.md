# dontdisturb

Small program to enable "Do not disturb" mode on **OSX** for X minutes,
since OSX only allows for setting this for the rest of the day.

## Install
# Use a Release

Grab the latest release from [here](https://github.com/embano1/dontdisturb/releases).

# Build from Source

`go get -u github.com/embano1/dontdisturb`

## Run (default: 15min)

make sure $GOPATH/bin is in your path, then: `dontdisturb`

## Option -t [minutes]

Enable "Do not disturb" for 10 minutes: `dontdisturb -t 10`
