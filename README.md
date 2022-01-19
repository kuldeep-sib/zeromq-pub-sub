# zeromq-pub-sub
zeromq POC

## install

make sure `libzmq` is installed on your system

using homebrew

```
brew install zeromq
```

check installed version
```
$ pkg-config --modversion libzmq
```

CGO should be enabled for go compiler

```
$ go env CGO_ENABLED
```
## build 

build the go project

```
mkdir ./build
go build -v  -o ./build  ./cmd/...
```

install dependencies for node project
```
yarn install
```

## run

run `proxy` , `pub` (publisher 1) and `pub2` (publisher 2) in any order

```
$ ./proxy
$ ./pub
$ ./pub2
```

run the subscribers
```
$ node sub.js
$ node sub2.js
```

you can run them in any order, messages may be lost but it's not a problem for our use case