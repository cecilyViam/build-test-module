build:
	GOOS=${VIAM_BUILD_OS} GOARCH=${VIAM_BUILD_ARCH} go build -o bin/build-test-module .

.PHONY: build
