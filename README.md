pam_netauth
===========

pam_netauth is, as you might have guessed, a PAM service module that
implements the NetAuth protocol.  This module can be installed
wherever your system installs modules for PAM and used for
authentication.  No other services are currently implemented.

This module is heavily inspired by the certificate based PAM project
from Uber.  Their code provided an excellent toe-hold to figure out
how to interface between PAM and Golang.
