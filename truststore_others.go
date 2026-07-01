//go:build !linux && !darwin && !windows && !freebsd
// +build !linux,!darwin,!windows,!freebsd

package truststore

import "crypto/x509"

var (
	// NSSProfiles are the directories that may contain a Firefox profiles.ini file.
	NSSProfiles []string

	// CertutilInstallHelp is the command to add NSS support.
	CertutilInstallHelp = ""
)

func installPlatform(string, *x509.Certificate) error {
	return ErrTrustNotSupported
}

func uninstallPlatform(string, *x509.Certificate) error {
	return ErrTrustNotSupported
}
