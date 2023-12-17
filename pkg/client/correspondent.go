package client

import (
	"context"
	"time"
)

type Correspondent struct {
	ID                 int64             `json:"id"`
	Slug               string            `json:"slug"`
	Name               string            `json:"name"`
	Match              string            `json:"match"`
	MatchingAlgorithm  MatchingAlgorithm `json:"matching_algorithm"`
	IsInsensitive      bool              `json:"is_insensitive"`
	DocumentCount      int64             `json:"document_count"`
	LastCorrespondence *time.Time        `json:"last_correspondence"`
}

func (c *Client) correspondentCrudOpts() crudOptions {
	return crudOptions{
		base:       "api/correspondents/",
		newRequest: c.newRequest,
		getID: func(v any) int64 {
			return v.(Correspondent).ID
		},
		setPage: func(opts any, page *PageToken) {
			opts.(*ListCorrespondentsOptions).Page = page
		},
	}
}

type ListCorrespondentsOptions struct {
	ListOptions

	Ordering OrderingSpec   `url:"ordering"`
	Owner    IntFilterSpec  `url:"owner"`
	Name     CharFilterSpec `url:"name"`
}

func (c *Client) ListCorrespondents(ctx context.Context, opts ListCorrespondentsOptions) ([]Correspondent, *Response, error) {
	return crudList[Correspondent](ctx, c.correspondentCrudOpts(), opts)
}

// ListAllCorrespondents iterates over all correspondents matching the filters
// specified in opts, invoking handler for each.
func (c *Client) ListAllCorrespondents(ctx context.Context, opts ListCorrespondentsOptions, handler func(context.Context, Correspondent) error) error {
	return crudListAll[Correspondent](ctx, c.correspondentCrudOpts(), opts, handler)
}

func (c *Client) GetCorrespondent(ctx context.Context, id int64) (*Correspondent, *Response, error) {
	return crudGet[Correspondent](ctx, c.correspondentCrudOpts(), id)
}

func (c *Client) CreateCorrespondent(ctx context.Context, data *Correspondent) (*Correspondent, *Response, error) {
	return crudCreate[Correspondent](ctx, c.correspondentCrudOpts(), data)
}

func (c *Client) UpdateCorrespondent(ctx context.Context, id int64, data *Correspondent) (*Correspondent, *Response, error) {
	return crudUpdate[Correspondent](ctx, c.correspondentCrudOpts(), id, data)
}

func (c *Client) DeleteCorrespondent(ctx context.Context, id int64) (*Response, error) {
	return crudDelete[Correspondent](ctx, c.correspondentCrudOpts(), id)
}
