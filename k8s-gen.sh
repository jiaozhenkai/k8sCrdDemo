
# 代码生成的工作目录，也就是我们的项目路径
ROOT_PACKAGE="k8sCrdDemo"
# API Group
CUSTOM_RESOURCE_NAME="hero"
# API Version
CUSTOM_RESOURCE_VERSION="v1"

# 安装k8s.io/code-generator
go get -u k8s.io/code-generator/...
cd $GOPATH/src/k8s.io/code-generator

# 执行代码自动生成，其中pkg/client是生成目标目录，pkg/apis是类型定义目录
./generate-groups.sh all "k8sCrdDemo/pkg/client" "k8sCrdDemo/pkg/apis" "hero:v1"
