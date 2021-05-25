APP = waiter

IMAGE ?= otaviof/$(APP)
IMAGE_TAG ?= $(IMAGE):latest

GO_FLAGS ?= -v -mod=vendor
GO_TEST_FLAGS ?= -cover

ARGS ?=

default: build

.PHONY: $(APP)
$(APP):
	go build $(GO_FLAGS) .

build: $(APP)

.PHONY: clean
clean:
	rm -f $(APP) > /dev/null

# .PHONY: test
# test:
# 	go test $(GO_FLAGS) $(GO_TEST_FLAGS) .

run:
	go run $(GO_FLAGS) . $(ARGS)

install: clean build
	CGO_ENABLED="0" go install $(GO_FLAGS) -ldflags '-extldflags "-static"' .

image:
	docker build --tag="$(IMAGE_TAG)" .