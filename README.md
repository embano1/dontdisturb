# dontdisturb

Small program to enable "Do not disturb" mode on **OSX** for X minutes,
since OSX only allows for setting this for the rest of the day.

## Install

`go get github.com/embano1/dontdisturb`

## Run (default: 15min)

make sure $GOPATH/bin is in your path, then: `dontdisturb`

## Option -t [minutes]

Enable "Do not disturb" for 10 minutes: `dontdisturb -t 10`
