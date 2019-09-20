package ssh

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/brainupdaters/drlm-common/pkg/fs"

	"github.com/spf13/afero"
	stdSSH "golang.org/x/crypto/ssh"
)

// NewSessionWithKey creates a new session using a private key as authentication
func NewSessionWithKey(host string, port int, user string, keysPath string, hostKeys []string) (*Session, error) {
	b, err := afero.ReadFile(fs.FS, filepath.Join(keysPath, "id_rsa"))
	if err != nil {
		return &Session{}, fmt.Errorf("error reading the private key: %v", err)
	}

	privKey, err := stdSSH.ParsePrivateKey(b)
	if err != nil {
		return &Session{}, fmt.Errorf("error parsing the private key: %v", err)
	}

	return newSession(host, port, user, []stdSSH.AuthMethod{stdSSH.PublicKeys(privKey)}, hostKeys)
}

// NewSessionWithPassword creates a new session using a password as authentication
func NewSessionWithPassword(host string, port int, user, pass string, hostKeys []string) (*Session, error) {
	return newSession(host, port, user, []stdSSH.AuthMethod{stdSSH.Password(pass)}, hostKeys)
}

// newSession creates a new SSH session
func newSession(host string, port int, user string, auth []stdSSH.AuthMethod, hostKeys []string) (*Session, error) {
	var hk []stdSSH.PublicKey
	for _, k := range hostKeys {
		_, _, h, _, _, err := stdSSH.ParseKnownHosts([]byte(k))
		if err != nil {
			return &Session{}, fmt.Errorf("error parsing the host public key: %v", err)
		}

		hk = append(hk, h)
	}

	sshCfg := &stdSSH.ClientConfig{
		User:            user,
		HostKeyCallback: MultipleFixedHostKeys(hk),
		Auth:            auth}

	conn, err := stdSSH.Dial("tcp", fmt.Sprintf("%s:%d", host, port), sshCfg)
	if err != nil {
		return &Session{}, fmt.Errorf("error dialing the host: %v", err)
	}

	sshSess, err := conn.NewSession()
	if err != nil {
		return &Session{}, fmt.Errorf("error creating the SSH session: %v", err)
	}

	modes := stdSSH.TerminalModes{
		stdSSH.ECHO:          0,
		stdSSH.TTY_OP_ISPEED: 14400,
		stdSSH.TTY_OP_OSPEED: 14400,
	}

	if err := sshSess.RequestPty("xterm", 80, 40, modes); err != nil {
		return &Session{}, fmt.Errorf("error requiesting the PTY: %v", err)
	}

	return &Session{sshSess}, nil
}

// Session is a wrapper of ssh.Session to make calls easier
type Session struct {
	s *stdSSH.Session
}

// Close closes the session
func (s *Session) Close() error {
	return s.s.Close()
}

// Exec executes a command and gets the output
func (s *Session) Exec(cmd string) ([]byte, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	s.s.Stdout = &stdout
	s.s.Stderr = &stderr

	if err := s.s.Run(cmd); err != nil {
		return stdout.Bytes(), fmt.Errorf("%v: %s", err, stderr.String())
	}

	return stdout.Bytes(), nil
}