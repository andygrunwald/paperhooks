package client

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"unicode"
)

func readFile(name string) (string, error) {
	content, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}

	return string(bytes.TrimRightFunc(content, unicode.IsSpace)), nil
}

// Flags contains attributes to construct a Paperless client instance. The
// separate "kpflag" package implements bindings for
// [github.com/alecthomas/kingpin/v2].
type Flags struct {
	// Whether to enable verbose log messages.
	DebugMode bool

	// HTTP(S) URL for Paperless.
	BaseURL string

	// Authenticate via token.
	AuthToken string

	// Read the authentication token from a file.
	AuthTokenFile string

	// Authenticate via HTTP basic authentication (username and password).
	AuthUsername string
	AuthPassword string

	// Read the password from a file.
	AuthPasswordFile string

	// HTTP headers to set on all requests.
	Header http.Header
}

// This function makes no attempt to deconflict different authentication
// options. Tokens from a files are preferred.
func (f *Flags) buildAuth() (AuthMechanism, error) {
	var err error

	token := f.AuthToken

	if f.AuthTokenFile != "" {
		token, err = readFile(f.AuthTokenFile)
		if err != nil {
			return nil, fmt.Errorf("reading authentication token failed: %w", err)
		}
	}

	if token != "" {
		return &TokenAuth{token}, nil
	}

	if f.AuthUsername != "" {
		password := f.AuthPassword

		if f.AuthPasswordFile != "" {
			password, err = readFile(f.AuthPasswordFile)
			if err != nil {
				return nil, fmt.Errorf("reading password failed: %w", err)
			}
		}

		return &UsernamePasswordAuth{
			Username: f.AuthUsername,
			Password: password,
		}, nil
	}

	return nil, nil
}

// BuildOptions returns the client options derived from flags.
func (f *Flags) BuildOptions() (*Options, error) {
	if f.BaseURL == "" {
		return nil, errors.New("Paperless URL is not specified")
	}

	opts := &Options{
		BaseURL:   f.BaseURL,
		DebugMode: f.DebugMode,
		Header:    http.Header{},
	}

	for name, values := range f.Header {
		name = http.CanonicalHeaderKey(name)
		for _, value := range values {
			opts.Header.Add(name, value)
		}
	}

	if auth, err := f.buildAuth(); err != nil {
		return nil, err
	} else {
		opts.Auth = auth
	}

	return opts, nil
}

// Build returns a fully configured Paperless client derived from flags.
func (f *Flags) Build() (*Client, error) {
	opts, err := f.BuildOptions()
	if err != nil {
		return nil, err
	}

	return New(*opts), nil
}
