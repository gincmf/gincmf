PROJECT="gincmf"
BINARY="gincmf"

default:
    echo ${PROJECT}
    @go env -w GO111MODULE=on
    @go env -w GOPROXY=https://goproxy.cn
    @go mod tidy

install: default
    @go build

clean:
    @if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: default install clean
