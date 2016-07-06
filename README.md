# tripper
tripper recomend of spot and can share with others

##requirement

```
go 1.6
mongo 3.2
git 2.7.2
```

## install

```
1.copy the project to your go work path
    mkdir -p $GOPATH/src/github.com/showntop
    cp {tripper} -rf $GOPATH/src/github.com/showntop/
    cd $GOPATH/src/github.com/showntop/tripper
    go get

    OR

2. use go get install
    go get github.com/showntop/tripper
    cd $GOPATH/src/github.com/showntop/tripper
    go get
```

## config
```
the config info in the tripper/config directory
in the config go modify the http port and database option

```

## run 

```
1. createdb tripper or the name you specified in the config file
1. go run main.go or go build && ./tripper
2. then test it
    curl -XGET "http://localhost:9000/users"
    .........
```