#!/bin/sh

go mod init github.com/renchunhui/go-admin
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go mod tidy

# 安装 mysql
if ! command -v mysql >/dev/null 2>&1; then
  brew install mysql
  echo ${GREEN}✔${RESET} "mysql"
fi

# YAML 解析工具
if ! command -v dasel >/dev/null 2>&1; then
	brew install dasel
	echo ${GREEN}✔${RESET} "dasel"
fi

DB_NAME=$(dasel 'select' -f configs/debug.yaml "database.name")
DB_USERNAME=$(dasel 'select' -f configs/debug.yaml "database.username")
DB_PASSWORD=$(dasel 'select' -f configs/debug.yaml "database.password")
