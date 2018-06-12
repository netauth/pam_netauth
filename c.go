package main

/*
#define PAM_SM_AUTH
#include <security/pam_appl.h>
#include <security/pam_ext.h>
#include <string.h>
#include <stdio.h>

char *get_user(pam_handle_t *pamh) {
  if (!pamh) {
    return NULL;
  }
  int err = 0;
  const char *user;
  if ((err = pam_get_item(pamh, PAM_USER, (const void**)&user)) != PAM_SUCCESS) {
    return NULL;
  }
  return strdup(user);
}

char *get_service(pam_handle_t *pamh) {
  if (!pamh) {
    return NULL;
  }
  int err = 0;
  const char *service;
  if ((err = pam_get_item(pamh, PAM_SERVICE, (const void**)&service)) != PAM_SUCCESS) {
    return NULL;
  }
  return strdup(service);
}

char *get_secret(pam_handle_t *pamh) {
  if (!pamh) {
    return NULL;
  }
  int err = 0;
  const char *secret;
  if ((err = pam_get_authtok(pamh, PAM_AUTHTOK, (const char **)&secret, NULL)) != PAM_SUCCESS) {
    return NULL;
  }
  return strdup(secret);
}
*/
import "C"
