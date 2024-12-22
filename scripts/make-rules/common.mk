#
# These variables should not need tweaking.
#

# ==============================================================================
# Includes

# include the common make file
ifeq ($(origin ACEX_ROOT),undefined)
ACEX_ROOT :=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
endif

# ==============================================================================
# Build options
#
ACEX_SRC_PATH :=github.com/mindmatterlab/acex

ifeq ($(origin OUTPUT_DIR),undefined)
OUTPUT_DIR := $(ACEX_ROOT)/_output
$(shell mkdir -p $(OUTPUT_DIR))
endif

# The OS must be linux when building docker images
PLATFORMS ?= linux_amd64 linux_arm64
# The OS can be linux/windows/darwin when building binaries
# PLATFORMS ?= darwin_amd64 windows_amd64 linux_amd64 linux_arm64

# Set a specific PLATFORM
ifeq ($(origin PLATFORM), undefined)
	ifeq ($(origin GOOS), undefined)
		GOOS := $(shell go env GOOS)
	endif
	ifeq ($(origin GOARCH), undefined)
		GOARCH := $(shell go env GOARCH)
	endif
	PLATFORM := $(GOOS)_$(GOARCH)
	# Use linux as the default OS when building images
	IMAGE_PLAT := linux_$(GOARCH)
else
	GOOS := $(word 1, $(subst _, ,$(PLATFORM)))
	GOARCH := $(word 2, $(subst _, ,$(PLATFORM)))
	IMAGE_PLAT := $(PLATFORM)
endif

MANIFESTS_DIR=$(ACEX_ROOT)/manifests
SCRIPTS_DIR=$(ACEX_ROOT)/scripts
APIROOT ?= $(ACEX_ROOT)/pkg/api
