# dontdisturb
Small program to enable "Do not disturb" mode on **OSX** for &lt;int> minutes, since OSX only allows for setting this for the rest of the day.

# install
go get github.com/embano1/dontdisturb

# run (default: 15min)
make sure $GOPATH/bin is in your path, then:  
dontdisturb

# Option -t \<min\>
enable "Do not disturb" for 10 minutes  
dontdisturb -t 10
