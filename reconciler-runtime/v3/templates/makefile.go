package templates

import "sigs.k8s.io/kubebuilder/v3/pkg/machinery"

var _ machinery.Template = &Makefile{}

type Makefile struct {
	machinery.TemplateMixin
	machinery.ComponentConfigMixin

	// TODO mbrauer - IMAGE_REGISTRY
	// TODO mbrauer - IMAGE_REGISTRY_PROJECT
	// TODO mbrauer - CONTROLLER_IMAGE
	// TODO mbrauer - PACKAGE_IMAGE
	// TODO mbrauer - PACKAGE_REPOSITORY_IMAGE
}

func (m *Makefile) SetTemplateDefaults() error {
	if m.Path == "" {
		m.Path = "Makefile"
	}

	m.TemplateBody = makefileTemplate
	m.IfExistsAction = machinery.Error

	return nil
}

const makefileTemplate = `
# ----------------
# This is the Makefile. It exposes all the project's affordances.
#
# If something is automated, you will find it here.
# If something should be automated, put it here.
#
# The convention for Make target names generally is "<noun>(-<verb>)"-ish,
# e.g. packaging, cluster-create or git-ensure-clean.
#
# Learn more about Makefiles:
#  * Quick reference: https://www.gnu.org/software/make/manual/html_node/Quick-Reference.html#Quick-Reference
#  * The Makefile tutorial: https://makefiletutorial.com/
#  * Make documentation: https://www.gnu.org/software/make/manual/html_node/index.html#Top
#  * Functions, variables and directives: https://www.gnu.org/software/make/manual/html_node/Name-Index.html
#
# Happy making!
# ----------------

.DEFAULT_GOAL := help

# Use Bash as the shell.
# See https://www.gnu.org/software/make/manual/html_node/Choosing-the-Shell.html
SHELL := bash

# Run each Make recipe as one single shell session.
# See https://www.gnu.org/software/make/manual/html_node/One-Shell.html#One-Shell
.ONESHELL:

# Run safe shell commands.
# See https://www.gnu.org/software/make/manual/html_node/Choosing-the-Shell.html
.SHELLFLAGS := -eu -o pipefail -c

# Remove targets if their recipes fail. This avoids corrupted or improperly built targets.
# See https://www.gnu.org/software/make/manual/html_node/Errors.html#Errors
.DELETE_ON_ERROR:
# Caveat: this only works for regular files and not for directories.
# See http://savannah.gnu.org/bugs/?func=detailitem&item_id=16372

MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

GINKGO := go run github.com/onsi/ginkgo/v2/ginkgo
GINKGO_FLAGS := --keep-going -v -r
ifeq ($(CI),true)
# On CI we run Ginkgo with all the recommended flags
# See https://onsi.github.io/ginkgo/#recommended-continuous-integration-configuration
GINKGO_FLAGS := $(GINKGO_FLAGS) --randomize-all --randomize-suites --fail-on-pending --race --trace --json-report=report.json
endif

.PHONY: help
help: ## Describe all make targets (default)
	@./Makefile_help.awk $(MAKEFILE_LIST)

.PHONY: test
test: ## Run unit tests
	$(GINKGO) $(GINKGO_FLAGS) --skip-package e2e

.PHONY: e2e
e2e: ## Run E2E tests
	$(GINKGO) $(GINKGO_FLAGS) e2e
`
