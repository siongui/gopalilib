# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
ifndef TRAVIS
	export GOROOT=$(realpath ../paligo/go)
	export GOPATH=$(realpath ../paligo)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

PALILIB=$(GOPATH)/src/github.com/siongui/gopalilib/lib
PALIUTIL=$(GOPATH)/src/github.com/siongui/gopalilib/util
DATA_REPO_DIR=$(CURDIR)/data
VFSDIR=$(GOPATH)/src/github.com/siongui/gopaliwordvfs


test_lib: test_url
	@echo "\033[92mTesting common library for online/offline dictionary/tipitaka ...\033[0m"
	@cd lib; go test -v dictionary.go filter.go json.go url.go url_test.go

test_symlink: fmt
	@echo "\033[92mTesting making Pāli Dictionary symlinks for GitHub Pages...\033[0m"
	@cd dicutil; go test -v symlink.go symlink_test.go path_test.go

test_vfsbuild: fmt
	@echo "\033[92mBuilding virtual file system of Pāli dictionary words ...\033[0m"
	#@[ -d $(VFSDIR) ] || mkdir -p $(VFSDIR)
	@cd dicutil; go test -v vfsbuild.go vfsbuild_test.go path_test.go -args -pkgdir=$(VFSDIR)
	@cd dicutil; go test -v vfs_test.go path_test.go

# test_triebuild must run before test_vfsbuild. Or re-run test_wordparser
test_triebuild: fmt
	@echo "\033[92mTesting building succinct trie ...\033[0m"
	@cd dicutil; go test -v triebuild.go triebuild_test.go path_test.go

test_wordparser: fmt
	@echo "\033[92mTesting parse CSV of dictionary words ...\033[0m"
	#@[ -d /tmp/paliwords/ ] || mkdir /tmp/paliwords/
	@cd dicutil; go test -v wordparser.go wordparser_test.go lib.go path_test.go

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


generate:
	@echo "\033[92mlib/: go generate ...\033[0m"
	@cd lib; go generate


clone_pali_data:
	@echo "\033[92mClone Pāli data Repo ...\033[0m"
	@[ -d $(DATA_REPO_DIR) ] || git clone  --depth 1 https://github.com/siongui/data.git $(DATA_REPO_DIR)

install: install_palilib install_gojianfan lib_succinct_trie install_goef

install_gojianfan:
	@echo "\033[92mInstalling Go Chinese conversion package ...\033[0m"
	go get -u github.com/siongui/gojianfan

#install_gocc:
#	@echo "\033[92mInstalling Golang version OpenCC package ...\033[0m"
#	go get -u github.com/liuzl/gocc

install_goef:
	@echo "\033[92mInstalling Go file embedder ...\033[0m"
	go get -u github.com/siongui/goef

install_palilib:
ifdef TRAVIS
	go get -u github.com/siongui/gopalilib/lib
	go get -u github.com/siongui/gopalilib/util
else
	@echo "\033[92mInstall ${PALILIB} locally ...\033[0m"
	@rm -rf ${PALILIB}
	@mkdir -p ${PALILIB}
	@cp -r lib/*.go ${PALILIB}/
	@echo "\033[92mInstall ${PALIUTIL} locally ...\033[0m"
	@rm -rf ${PALIUTIL}
	@mkdir -p ${PALIUTIL}
	@cp -r util/*.go ${PALIUTIL}/
endif

install_stringer:
	@echo "\033[92mInstalling golang.org/x/tools/cmd/stringer ...\033[0m"
	go get -u golang.org/x/tools/cmd/stringer

lib_succinct_trie:
	@echo "\033[92mInstalling Go Succinct Trie library ...\033[0m"
	go get -u github.com/siongui/go-succinct-data-structure-trie

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt lib/*.go
	@go fmt dicutil/*.go
	@go fmt i18n/*.go
	@go fmt util/*.go

clean:
	rm -rf pkg/ src/
