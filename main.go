package main

import (
	"fmt"
	"unsafe"
)

/*
#cgo LDFLAGS: -lpam -fPIC
#define PAM_SM_AUTH
#include <security/pam_appl.h>
#include <stdlib.h>

char *get_user(pam_handle_t *pamh);
char *get_secret(pam_handle_t *pamh);
*/
import "C"

//export pam_sm_authenticate
func pam_sm_authenticate(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	fmt.Println("In netauth code")
	cUsername := C.get_user(pamh)
	if cUsername == nil {
		return C.PAM_USER_UNKNOWN
	}
	defer C.free(unsafe.Pointer(cUsername))

	cSecret := C.get_secret(pamh)
	if cSecret == nil {
		fmt.Println("cSecret was nil")
		return C.PAM_CRED_INSUFFICIENT
	}
	defer C.free(unsafe.Pointer(cSecret))

	fmt.Printf("Authenticating as %s:%s\n", C.GoString(cUsername), C.GoString(cSecret))
	return C.PAM_SUCCESS
}

// This doesn't do anything but the compiler needs to see a "main"
// symbol in order to proceed.
func main() {}
