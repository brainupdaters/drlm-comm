// SPDX-License-Identifier: AGPL-3.0-only

package os

import (
	"errors"
	"fmt"
	"strings"

	"github.com/brainupdaters/drlm-common/pkg/os/client"
)

// OS is an Operative System
type OS int

const (
	// Unknown is a not known OS
	Unknown OS = iota
	// Linux is a Linux OS
	Linux
	// Windows is a Microsoft Windows OS
	Windows
	// Darwin is a macOS OS
	Darwin
	// AIX is an IBM AIX OS
	AIX
	// Dragonfly is a Dragonfly BSD OS
	Dragonfly
	// FreeBSD is a FreeBSD OS
	FreeBSD
	// NetBSD is a NetBSD OS
	NetBSD
	// OpenBSD is an OpenBSD OS
	OpenBSD
	// Plan9 is a Plan9 OS
	Plan9
	// Solaris is a Solaris OS
	Solaris
)

// ErrUnsupportedOS is an error that gets returned when the command is not supported in the OS
var ErrUnsupportedOS = errors.New("os not supported yet")

// IsUnix returns whether the OS is an Unix-like OS
func (os OS) IsUnix() bool {
	switch os {
	case Linux, Darwin, Dragonfly, FreeBSD, NetBSD, OpenBSD:
		return true
	default:
		return false
	}
}

// DetectOS detects the OS the Client is running
func DetectOS(c client.Client) (OS, error) {
	out, err := c.Exec("uname", "-s")
	if err != nil {
		// NOT UNIX
		return Unknown, fmt.Errorf("error getting the OS: %v", err)
	}

	return FromString(strings.TrimSpace(string(out))), nil
}

// DetectVersion returns the OS version
func (os OS) DetectVersion(c client.Client) (string, error) {
	switch {
	case os.IsUnix():
		out, err := c.Exec("uname", "-r")
		if err != nil {
			return "", fmt.Errorf("error detecting the OS version: %v", err)
		}

		return strings.TrimSpace(string(out)), nil

	default:
		return "unknown", ErrUnsupportedOS
	}
}

// FromString returns an OS based on a string
func FromString(s string) OS {
	switch strings.ToLower(s) {
	case "linux":
		return Linux

	case "windows":
		return Windows

	case "darwin":
		return Darwin

	case "aix":
		return AIX

	case "dragonfly":
		return Dragonfly

	case "freebsd":
		return FreeBSD

	case "netbsd":
		return NetBSD

	case "openbsd":
		return OpenBSD

	case "plan9":
		return Plan9

	case "solaris":
		return Solaris

	default:
		return Unknown
	}
}

// DetectDistro returns the OS distro and distro version (or the OS equivalent)
func (os OS) DetectDistro(c client.Client) (string, string, error) {
	distro := "unknown"
	version := "unknown"

	switch os {
	case Linux:
		out, err := c.ReadFile("/etc/os-release")
		if err != nil {
			return distro, version, fmt.Errorf("error detecting the linux distro: %v", err)
		}

		for _, l := range strings.Split(string(out), "\n") {
			if strings.HasPrefix(l, "ID=") {
				distro = strings.TrimSpace(strings.Split(l, "=")[1])
			}

			if strings.HasPrefix(l, "VERSION_ID=") {
				version = strings.TrimSpace(strings.Split(l, "=")[1])
			}
		}

		return distro, version, nil

	default:
		return "unknown", "unknown", ErrUnsupportedOS
	}
}
