#!/bin/bash

# 检查参数个数
if [ $# -eq 0 ]; then
  echo "没有传递任何参数"
  exit 1
fi


# 检查 go 命令是否存在于 PATH 环境变量中
if ! command -v go &>/dev/null; then
  echo "错误：go 命令未在 PATH 中找到。请安装或将其添加到 PATH 中。"
  exit 1
fi



# 检查 kitex 命令是否存在于 PATH 环境变量中
if ! command -v kitex &>/dev/null; then
  echo "错误：kitex 命令未在 PATH 中找到，尝试安装......"
  go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
fi

# 再次检查 kitex 命令是否存在于 PATH 环境变量中
if ! command -v kitex &>/dev/null; then
  echo "错误：kitex 命令未在 PATH 中找到，看起来安装失败了。请手动安装。"
  exit 1
fi

kitex -module "douyin" -I idl/ idl/"$1".thrift

mkdir -p rpc/"$1"
cd rpc/"$1" && kitex -module "douyin" -service "$1" -use douyin/kitex_gen/ -I ../../idl/ ../../idl/"$1".thrift

go mod tidy