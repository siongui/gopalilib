# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

PALILIB=src/github.com/siongui/gopalilib/lib
PALIUTIL=src/github.com/siongui/gopalilib/util

test_bookparser: fmt
	@echo "\033[92mTesting parse CSV of dictionary books ...\033[0m"
	@cd dicutil; go test -v bookparser.go bookparser_test.go path_test.go

test_util: fmt
	@echo "\033[92mTest utility func ...\033[0m"
	@cd util; go test -v

test_twpo2cn: fmt
	@echo "\033[92mCreating zh_CN PO from zh_TW PO ...\033[0m"
	@cd i18n; go test -v twpo2cn.go twpo2cn_test.go

test_url: fmt
	@echo "\033[92mTesting Url ...\033[0m"
	@cd lib; go test -v url.go url_test.go

test_html: fmt
	@echo "\033[92mTesting Pāli Dictionary HTML ...\033[0m"
	@cd dicutil; go test -v html.go html_test.go path_test.go

test_symlink: fmt
	@echo "\033[92mTesting making Pāli Dictionary symlinks for GitHub Pages...\033[0m"
	@cd dicutil; go test -v symlink.go symlink_test.go

generate:
	@echo "\033[92mlib/: go generate ...\033[0m"
	@cd lib; go generate

install: install_palilib install_gotm install_gojianfan

install_gojianfan:
	@echo "\033[92mInstalling Go Chinese conversion package ...\033[0m"
	go get -u github.com/siongui/gojianfan

install_gotm:
	@echo "\033[92mInstalling Go template manager ...\033[0m"
	go get -u github.com/siongui/gotm

install_palilib:
	@echo "\033[92mInstall ${PALILIB} locally ...\033[0m"
	@rm -rf ${PALILIB}
	@mkdir -p ${PALILIB}
	@cp -r lib/*.go ${PALILIB}/
	@echo "\033[92mInstall ${PALIUTIL} locally ...\033[0m"
	@rm -rf ${PALIUTIL}
	@mkdir -p ${PALIUTIL}
	@cp -r util/*.go ${PALIUTIL}/

install_stringer:
	@echo "\033[92mInstalling golang.org/x/tools/cmd/stringer ...\033[0m"
	go get -u golang.org/x/tools/cmd/stringer

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt lib/*.go
	@go fmt dicutil/*.go
	@go fmt i18n/*.go
	@go fmt util/*.go

clean:
	rm -rf pkg/ src/
