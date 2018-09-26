set env

set $GOPATH

```
go get -u github.com/astaxie/beego
go get -u github.com/beego/bee
```

```
cd $GOPATH/src
cp -R {{$git}}/ToyProject/tyg03485/todo .
bee run todo
```

