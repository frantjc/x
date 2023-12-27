package xurl

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func Open(name string) (io.ReadCloser, error) {
	return OpenContext(context.TODO(), name)
}

func OpenContext(ctx context.Context, name string) (io.ReadCloser, error) {
	u, err := url.Parse(name)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "", "file":
		return os.Open(u.Path)
	case "http", "https":
		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodGet,
			u.String(),
			nil,
		)
		if err != nil {
			return nil, err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		return res.Body, nil
	}

	return nil, fmt.Errorf("unopenable url: %s", name)
}
