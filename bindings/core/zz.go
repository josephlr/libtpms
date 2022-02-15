package core

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
// #include "Tpm.h"
// #include "_TPM_Init_fp.h"
import "C"

func Check() bool {
	C._TPM_Init()
	C._plat__RunCommand()

	C._TPM_Hash_Start()
	C._TPM_Hash_Data()
	C._TPM_Hash_End()

	C.TPM_Manufacture()
	C.TPM_TearDown()

	// C.VolatileState_Marshal()
	// C.VolatileState_Unmarshal()
	// C.PERSISTENT_ALL_Marshal()
	// C.PERSISTENT_ALL_Unmarshal()

	return C.g_inFailureMode
}
