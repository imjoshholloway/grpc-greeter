// +build tools

package tools

// Track our build tool dependencies
// See: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
// & https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md for more information

import (
	// binaries that we need to be consistent from a dev perspective
	_ "github.com/bitnami/kubecfg"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/tetratelabs/istio-tools/grpc-transcoder"
	_ "github.com/uber/prototool/cmd/prototool"
)
