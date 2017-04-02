# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

PALILIB=src/github.com/siongui/gopalilib/lib

test_url: fmt
	@echo "\033[92mTesting Url ...\033[0m"
	@cd lib; go test -v url.go url_test.go

test_html: fmt
	@echo "\033[92mTesting Pāli Dictionary HTML ...\033[0m"
	@cd dicutil; go test -v html.go html_test.go

test_symlink: fmt
	@echo "\033[92mTesting making Pāli Dictionary symlinks for GitHub Pages...\033[0m"
	@cd dicutil; go test -v symlink.go symlink_test.go

generate:
	@echo "\033[92mlib/: go generate ...\033[0m"
	@cd lib; go generate

install: install_palilib
	@echo "\033[92mInstalling Go template utility ...\033[0m"
	go get -u github.com/siongui/gotemplateutil

install_palilib:
	@echo "\033[92mInstall Pali lib locally ...\033[0m"
	@mkdir -p ${PALILIB}
	@cp -r lib/*.go ${PALILIB}/

install_stringer:
	@echo "\033[92mInstalling golang.org/x/tools/cmd/stringer ...\033[0m"
	go get -u golang.org/x/tools/cmd/stringer

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt lib/*.go
	@go fmt dicutil/*.go
