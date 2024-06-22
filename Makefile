# 定义项目名称
PROJECT_NAME := protohub.api

GOCTL := goctl

# 设置go模块环境变量
export GO111MODULE := on

# 格式化API定义文件
format:
	$(GOCTL) api format --dir ./

# 生成Go API代码
gen-go-api:
	$(GOCTL) api go --api $(PROJECT_NAME) --dir ./ --style=go_zero


# 默认目标
all: format gen-go-api

api: format gen-go-api

.PHONY: format gen-go-api