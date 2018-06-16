package main

import (
	"fmt"
	"unsafe"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NetAuth/NetAuth/pkg/client"
)

/*
#cgo LDFLAGS: -lpam -fPIC
#define PAM_SM_AUTH
#include <security/pam_appl.h>
#include <stdlib.h>

char *get_service(pam_handle_t *pamh);
char *get_user(pam_handle_t *pamh);
char *get_secret(pam_handle_t *pamh);
*/
import "C"

//export pam_sm_authenticate
func pam_sm_authenticate(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	cService := C.get_service(pamh)
	if cService == nil {
		return C.PAM_SYSTEM_ERR
	}
	defer C.free(unsafe.Pointer(cService))

	nacl, err := client.New(nil)
	if err != nil {
		// Couldn't get a client
		return C.PAM_AUTHTOK_ERR
	}
	nacl.SetServiceID(C.GoString(cService))

	cUsername := C.get_user(pamh)
	if cUsername == nil {
		_, err := nacl.EntityInfo(C.GoString(cUsername))
		if status.Code(err) == codes.NotFound {
			return C.PAM_USER_UNKNOWN
		}
		// Something went wrong trying to run the info query
		return C.PAM_AUTHTOK_ERR

	}
	defer C.free(unsafe.Pointer(cUsername))

	cSecret := C.get_secret(pamh)
	if cSecret == nil {
		fmt.Println("cSecret was nil")
		return C.PAM_CRED_INSUFFICIENT
	}
	defer C.free(unsafe.Pointer(cSecret))

	// Run the authentication
	result, err := nacl.Authenticate(C.GoString(cUsername), C.GoString(cSecret))
	if status.Code(err) != codes.OK || !result.GetSuccess() {
		return C.PAM_AUTH_ERR
	}
	fmt.Println(result)
	return C.PAM_SUCCESS
}

// This doesn't do anything but the compiler needs to see a "main"
// symbol in order to proceed.
func main() {}
