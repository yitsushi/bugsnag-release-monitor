export GOPROXY ?= https://proxy.golang.org
export GOSUMDB ?= sum.golang.org

GIT_VERSION      = $(shell git describe --always --abbrev=7 --dirty)

CGO         ?= 0
BINARIES    ?= bugsnag-release-monitor

TARGET_ARCH_LOCAL := amd64
ifneq ($(OS),Windows_NT)
LOCAL_ARCH = $(shell uname -m)
ifeq ($(LOCAL_ARCH),x86_64)
	TARGET_ARCH_LOCAL=amd64
else ifeq ($(shell echo $(LOCAL_ARCH) | head -c 5),armv8)
	TARGET_ARCH_LOCAL=arm64
else ifeq ($(shell echo $(LOCAL_ARCH) | head -c 4),armv)
	TARGET_ARCH_LOCAL=arm
else
	TARGET_ARCH_LOCAL=amd64
endif
endif

export GOARCH ?= $(TARGET_ARCH_LOCAL)

ifeq ($(OS),Windows_NT)
	TARGET_OS_LOCAL ?= windows
else
LOCAL_OS := $(shell uname)
ifeq ($(LOCAL_OS),Linux)
   TARGET_OS_LOCAL = linux
else ifeq ($(LOCAL_OS),Darwin)
   TARGET_OS_LOCAL = darwin
else
   TARGET_OS_LOCAL ?= windows
endif
endif
export GOOS ?= $(TARGET_OS_LOCAL)

ifeq ($(GOOS),windows)
BINARY_EXT_LOCAL:=.exe
GOLANGCI_LINT:=golangci-lint.exe
export ARCHIVE_EXT = .zip
else
BINARY_EXT_LOCAL:=
GOLANGCI_LINT:=golangci-lint
export ARCHIVE_EXT = .tar.gz
endif

export BINARY_EXT ?= $(BINARY_EXT_LOCAL)

OUT_DIR := ./dist
BASE_PACKAGE_NAME := https://github.com/yitsushi/bugsnag-release-monitor

DEFAULT_LDFLAGS:=\
  -X $(BASE_PACKAGE_NAME)/pkg/version.Build=$(GIT_VERSION)

ifeq ($(origin DEBUG), undefined)
  BUILDTYPE_DIR:=release
  LDFLAGS:="$(DEFAULT_LDFLAGS) -s -w"
else ifeq ($(DEBUG),0)
  BUILDTYPE_DIR:=release
  LDFLAGS:="$(DEFAULT_LDFLAGS) -s -w"
else
  BUILDTYPE_DIR:=debug
  GCFLAGS:=-gcflags="all=-N -l"
  LDFLAGS:="$(DEFAULT_LDFLAGS)"
  $(info Build with debugger information)
endif

DAPR_OUT_DIR := $(OUT_DIR)/$(GOOS)_$(GOARCH)/$(BUILDTYPE_DIR)
DAPR_LINUX_OUT_DIR := $(OUT_DIR)/linux_$(GOARCH)/$(BUILDTYPE_DIR)

.PHONY: build
DAPR_BINS:=$(foreach ITEM,$(BINARIES),$(DAPR_OUT_DIR)/$(ITEM)$(BINARY_EXT))
build: $(DAPR_BINS)

ifeq ($(GOOS),windows)
define genBinariesForTarget
.PHONY: $(5)/$(1)
$(5)/$(1):
	set CGO_ENABLED=$(CGO)
	set GOOS=$(3)
	set GOARCH=$(4)
	go build $(GCFLAGS) -ldflags=$(LDFLAGS) \
	-o $(5)/$(1) \
	$(2)/main.go
endef
else
define genBinariesForTarget
.PHONY: $(5)/$(1)
$(5)/$(1):
	CGO_ENABLED=$(CGO) GOOS=$(3) GOARCH=$(4) go build $(GCFLAGS) -ldflags=$(LDFLAGS) \
	-o $(5)/$(1) \
	$(2)/main.go;
endef
endif

$(foreach ITEM,$(BINARIES),$(eval $(call genBinariesForTarget,$(ITEM)$(BINARY_EXT),./cmd/$(ITEM),$(GOOS),$(GOARCH),$(DAPR_OUT_DIR))))

BUILD_LINUX_BINS:=$(foreach ITEM,$(BINARIES),$(DAPR_LINUX_OUT_DIR)/$(ITEM))
build-linux: $(BUILD_LINUX_BINS)

ifneq ($(GOOS), linux)
$(foreach ITEM,$(BINARIES),$(eval $(call genBinariesForTarget,$(ITEM),./cmd/$(ITEM),linux,$(GOARCH),$(DAPR_LINUX_OUT_DIR))))
endif

ARCHIVE_OUT_DIR ?= $(DAPR_OUT_DIR)
ARCHIVE_FILE_EXTS:=$(foreach ITEM,$(BINARIES),archive-$(ITEM)_$(GIT_VERSION)$(ARCHIVE_EXT))

archive: $(ARCHIVE_FILE_EXTS)

define genArchiveBinary
ifeq ($(GOOS),windows)
archive-$(1)_$(GIT_VERSION).zip:
	7z.exe a -tzip "$(2)\\$(1)_$(GIT_VERSION)_$(GOOS)_$(GOARCH)$(ARCHIVE_EXT)" "$(DAPR_OUT_DIR)\\$(1)$(BINARY_EXT)"
else
archive-$(1)_$(GIT_VERSION).tar.gz:
	tar czf "$(2)/$(1)_$(GIT_VERSION)_$(GOOS)_$(GOARCH)$(ARCHIVE_EXT)" -C "$(DAPR_OUT_DIR)" "$(1)$(BINARY_EXT)"
endif
endef

$(foreach ITEM,$(BINARIES),$(eval $(call genArchiveBinary,$(ITEM),$(ARCHIVE_OUT_DIR))))

release: build archive

.PHONY: test
test:
	go test ./pkg/...

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run --enable-all
