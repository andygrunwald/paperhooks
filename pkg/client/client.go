package client

import (
	"context"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

// Options for constructing a Paperless client.
type Options struct {
	// Paperless URL. May include a path.
	BaseURL string

	// API authentication.
	Auth AuthMechanism

	// Enable debug mode with many details logged.
	DebugMode bool

	// Logger for writing log messages. If debug mode is enabled and no logger
	// is configured all messages are written to standard library's default
	// logger (log.Default()).
	Logger Logger

	// HTTP headers to set on all requests.
	Header http.Header

	// Override the default HTTP transport.
	transport http.RoundTripper
}

type Client struct {
	logger Logger
	r      *resty.Client
}

// New creates a new client instance.
func New(opts Options) *Client {
	if opts.Logger == nil {
		if opts.DebugMode {
			opts.Logger = &wrappedStdLogger{log.Default()}
		} else {
			opts.Logger = &discardLogger{}
		}
	}

	r := resty.New().
		SetDebug(opts.DebugMode).
		SetLogger(&prefixLogger{
			wrapped: opts.Logger,
			prefix:  "Resty: ",
		}).
		SetDisableWarn(true).
		SetBaseURL(opts.BaseURL).
		SetHeader("Accept", "application/json; version=2").
		SetRedirectPolicy(resty.NoRedirectPolicy())

	if opts.Auth != nil {
		opts.Auth.authenticate(r)
	}

	if opts.transport != nil {
		r.SetTransport(opts.transport)
	}

	if len(opts.Header) > 0 {
		r.SetPreRequestHook(func(_ *resty.Client, req *http.Request) error {
			for name, values := range opts.Header {
				req.Header[http.CanonicalHeaderKey(name)] = values
			}

			return nil
		})
	}

	return &Client{
		logger: opts.Logger,
		r:      r,
	}
}

func (c *Client) newRequest(ctx context.Context) *resty.Request {
	return c.r.R().
		SetContext(ctx).
		SetError(requestError{}).
		ExpectContentType("application/json")
}

type Response struct {
	*http.Response

	// Token for fetching next page in paginated result sets.
	NextPage *PageToken

	// Token for fetching previous page in paginated result sets.
	PrevPage *PageToken
}

func wrapResponse(r *resty.Response) *Response {
	return &Response{Response: r.RawResponse}
}
