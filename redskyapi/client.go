/*
Copyright 2019 GramLabs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package redskyapi

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client interface {
	URL(endpoint string) *url.URL
	Do(context.Context, *http.Request) (*http.Response, []byte, error)
}

// NewClient returns a new client for accessing Red Sky APIs; the supplied context is used for authentication/authorization
// requests and the supplied transport (which may be nil in the case of the default transport) is used for all requests made
// to the API server.
func NewClient(cfg Config, ctx context.Context, transport http.RoundTripper) (Client, error) {
	var err error

	hc := &httpClient{config: cfg}
	hc.client.Timeout = 10 * time.Second

	// Configure the OAuth2 transport
	hc.client.Transport, err = cfg.Authorize(ctx, transport)
	if err != nil {
		return nil, err
	}

	// Make sure that we can ignore the error from ExperimentsURL
	_, err = cfg.ExperimentsURL("")
	if err != nil {
		return nil, err
	}

	return hc, nil
}

type httpClient struct {
	config Config
	client http.Client
}

func (c *httpClient) URL(ep string) *url.URL {
	u, _ := c.config.ExperimentsURL(ep)
	return u
}

func (c *httpClient) Do(ctx context.Context, req *http.Request) (*http.Response, []byte, error) {
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var body []byte
	done := make(chan struct{})
	go func() {
		body, err = ioutil.ReadAll(resp.Body)
		close(done)
	}()

	select {
	case <-ctx.Done():
		<-done
		err = resp.Body.Close()
		if err == nil {
			err = ctx.Err()
		}
	case <-done:
	}

	return resp, body, err
}
