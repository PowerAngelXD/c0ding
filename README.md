# c0ding 程序员交流社区
### 该项目目前主要为学习作品，暂无实际意义

## Build - Dev
本项目使用NuxtUI + Golang + Postgresql数据库

### NuxtUI
cd 到 `frontend` 下：
```bash
cd frontend
```
本项目使用的工具为 `pnpm`，因此请按照下面的方式安装依赖：
```bash
pnpm install
```
### Golang
先切换工作目录到 `backend` 下：
```bash
cd backend
```
之后，运行下列命令下载所需依赖，并tidy：
```bash
go mod download

go mod tidy
```

### Postgresql 配置
以 `Arch linux` 为例：
```bash
sudo pacman -Syu postgresql
```
即可