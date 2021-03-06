// SPDX-License-Identifier: AGPL-3.0-only

package os

import (
	"fmt"
	"path/filepath"

	"github.com/brainupdaters/drlm-common/pkg/os/client"
)

// CmdPkgInstallBinary installs a binary in a system
func (os OS) CmdPkgInstallBinary(c client.Client, usr, name string, b []byte) error {
	switch os {
	case Linux:
		home, err := os.CmdFSHome(c, usr)
		if err != nil {
			return fmt.Errorf("error installing the binary: %v", err)
		}

		binDir := filepath.Join(home, ".bin")

		exists, err := c.Exists(binDir)
		if err != nil {
			return fmt.Errorf("error installing the binary: %v", err)
		}
		if !exists {
			if err = c.MkdirAll(binDir, 0755); err != nil {
				return fmt.Errorf("error creating the bin directory: %v", err)
			}

			if err = c.Append(filepath.Join(home, ".profile"), []byte(`
# Generated by DRLM
export PATH="$PATH:$HOME/.bin"
`)); err != nil {
				return fmt.Errorf("error adding the bin dir to the PATH: %v", err)
			}
		}

		binPath := filepath.Join(binDir, name)
		if err := c.Write(binPath, b); err != nil {
			return fmt.Errorf("error writting the binary: %v", err)
		}

		if err := c.Chmod(binPath, 0755); err != nil {
			return fmt.Errorf("error making the binary executable: %v", err)
		}

		return nil

	default:
		return ErrUnsupportedOS
	}
}

// CmdPkgWriteConfig writes the configuration of a DRLM program
func (os OS) CmdPkgWriteConfig(c client.Client, usr, name string, b []byte) error {
	switch os {
	case Linux:
		home, err := os.CmdFSHome(c, usr)
		if err != nil {
			return fmt.Errorf("error installing the binary: %v", err)
		}

		cfgDir := filepath.Join(home, ".config", "drlm")
		exists, err := c.Exists(cfgDir)
		if err != nil {
			return fmt.Errorf("error writting the configuration: %v", err)
		}
		if !exists {
			if err = c.MkdirAll(cfgDir, 0755); err != nil {
				return fmt.Errorf("error creating the configuration directory: %v", err)
			}
		}

		cfgPath := filepath.Join(cfgDir, name)
		if err := c.Write(cfgPath, b); err != nil {
			return fmt.Errorf("error writting the configuration file: %v", err)
		}

		return nil

	default:
		return ErrUnsupportedOS
	}
}
