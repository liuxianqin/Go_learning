
### 下载go


### 配置环境变量


### 安装插件


国内无法安装VS code 插件：

`go env -w GO111MODULE=on`
`go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct`

安装完插件后无法build `go build`报`go: cannot find main module; see 'go help modules'`:

把环境改回去
`go env -w GO111MODULE=auto`