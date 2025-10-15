//go:build amd64

package go_server_core_binaries_macos

import (
	_ "embed"
)

//go:embed macos_x86_64.dylib
var binaryData []byte

//go:embed macos_x86_64.dylib.sig
var signatureData []byte

func GetBinaryData() []byte {
	return binaryData
}

func GetSignatureData() []byte {
	return signatureData
}
