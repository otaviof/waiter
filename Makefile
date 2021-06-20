APP = waiter

CONTAINER_MANAGER ?= docker
IMAGE_PREFIX ?= otaviof
IMAGE_TAG ?= latest
IMAGE ?= $(IMAGE_PREFIX)/$(APP):$(IMAGE_TAG)

GO_FLAGS ?= -v -mod=vendor
GO_TEST_FLAGS ?= -cover -race

ARGS ?= --help

default: build

# compiles the applicatoin binary
.PHONY: $(APP)
$(APP):
	go build $(GO_FLAGS) .

# build target points to the applicaton binary
build: $(APP)

# removes the application binary
.PHONY: clean
clean:
	rm -f $(APP) > /dev/null || true

# executes all project tests
.PHONY: test
test:
	go test $(GO_FLAGS) $(GO_TEST_FLAGS) .

# executes "go run" using ARGS as the flags passed to the application
run:
	go run $(GO_FLAGS) . $(ARGS)

# installs the application with "go install"
install: clean build
	CGO_ENABLED="0" go install $(GO_FLAGS) -ldflags '-extldflags "-static"' .

# builds the container-image using pre-defined container manager
image:
	$(CONTAINER_MANAGER) build --tag="$(IMAGE)" .