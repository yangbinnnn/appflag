## appflag

简单的子命令创建工具

示例: [bin/main.go](bin/main.go)

演示:
```
$git clone https://github.com/yangbinnnn/appflag.git
$cd appflag/bin
$go run .
Usage of app:
  -do
        just do it
  -hello
        hello cmd set
  -migrate
        migrate cmd set

# 查看某个子命令
$ go run . -hello
Usage of hello:
  -hi
        hi to name
  -say
        say something

# 带参数子命令
$ go run . -hello -hi -name yangbin   
hi: yangbin

# 多级子命令嵌套
$ go run . -migrate -deepcmd -sleep -second 3
sleep...
sleep done.
```
