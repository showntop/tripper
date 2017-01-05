# tripper
==========

shares your inner world with others

##requirement

```
go 1.7
postgresql 9.5.3
git 2.7.2
```

## install

```
1.copy the project to your go work path
    mkdir -p $GOPATH/src/github.com/showntop
    cp {journey} -rf $GOPATH/src/github.com/showntop/
    cd $GOPATH/src/github.com/showntop/journey
    go get

    OR

2. use go get install
    go get github.com/showntop/journey
    cd $GOPATH/src/github.com/showntop/journey
    go get
```

## config
```
the config info in the journey/config directory
in the config go modify the http port and database option

```

## run 

```
1. createdb journey or the name you specified in the config file
1. go run main.go or go build && ./journey
2. then test it
    curl -XGET "http://localhost:9000/apps"
    .........
```
