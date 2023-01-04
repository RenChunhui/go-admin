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

DB_NAME=$(dasel -f configs/debug.yaml "Database.Name")
DB_USERNAME=$(dasel -f configs/debug.yaml "Database.Username")
DB_PASSWORD=$(dasel -f configs/debug.yaml "Database.Password")

# 创建数据库
Q1="CREATE DATABASE IF NOT EXISTS $DB_NAME;"
Q2="CREATE USER '$DB_USERNAME'@'localhost' IDENTIFIED BY '$DB_PASSWORD'"
Q3="GRANT ALL ON *.* TO '$DB_USERNAME'@'localhost';"
Q4="FLUSH PRIVILEGES;"
SQL="${Q1}${Q2}${Q3}${Q4}"

mysql -u root -e "$SQL"
