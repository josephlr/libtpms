package platform

// // Include other header directories
// #cgo CFLAGS: -I../../include/libtpms -I.. -I../../src -I../../src/tpm2 -I../../src/tpm2/crypto -I../../src/tpm2/crypto/openssl
// // Include specific files
// #cgo CFLAGS: -include ../config.h
// // Default flags used by libtpms build system
// #cgo CFLAGS: -g -O2
// // Enable a bunch of warnings (had to remove -Wmissing-prototypes and add -Wno-unused-parameter)
// #cgo CFLAGS: -Wall -Wextra -Wreturn-type -Wsign-compare -Wself-assign -Wno-unused-parameter -fstack-protector-strong
// // Flags to find OpenSSL installation on macOS (default Homebrew location)
// #cgo darwin CFLAGS: -I/usr/local/opt/openssl/include
// #cgo darwin LDFLAGS: -L/usr/local/opt/openssl/lib
// // Flags to find OpenSSL installation on Windows (default install location)
// #cgo windows CFLAGS: -I"C:/Program Files/OpenSSL-Win64/include"
// #cgo windows LDFLAGS: -L"C:/Program Files/OpenSSL-Win64/lib"
// // Link against OpenSSL
// #cgo LDFLAGS: -lcrypto
//
// #include "tpm_library.h"
import "C"

func GetVersion() uint32 {
	return uint32(C.TPMLIB_GetVersion())
}
