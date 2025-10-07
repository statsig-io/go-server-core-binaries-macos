//go:build darwin && x86_64

package go_server_core_binaries_macos

import (
	_ "embed"
)

//go:embed libstatsig_ffi.dylib
var binaryData []byte

//go:embed libstatsig_ffi.dylib.sig
var signatureData []byte

func GetBinaryData() []byte {
	return binaryData
}

func GetSignatureData() []byte {
	return signatureData
}
