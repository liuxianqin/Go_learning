
### 下载go


### 配置环境变量


### 安装插件


国内无法安装VS code 插件：

`go env -w GO111MODULE=on`
`go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct`

安装完插件后无法build `go build`报`go: cannot find main module; see 'go help modules'`:

把环境改回去
`go env -w GO111MODULE=auto`



### 常用的GO 命令

`go build` 编译Go程序

`go build -o name` 指定可执行文件名字

`go run main.go` 直接运行go文件

`go install`  把可执行文件拷贝到GOPATH，能在机器中直接使用



### 一个go文件的注意事项

如果是可执行文件，必须声明main包`package main`，main()函数作为入口，没有参数，也没有返回值\

变量声明：

1. var name string
2. var name = "HHH"
3. name := "HHH"  （函数内部专用）




