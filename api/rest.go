package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	ErrInvalidStatusCode = "the expected status code did not come back from the api"

	ContentTypeJson = "application/json"
)

type RestClient struct {
	client http.Client
}

func NewRestClient() *RestClient {
	return &RestClient{
		client: http.Client{},
	}
}

type RestArgs struct {
	Url         string
	StatusCode  int
	ContentType string
	Body        interface{}
	//Model       interface{}
}

func (c RestClient) Get(ctx context.Context, Args RestArgs) ([]byte, error) {
	var res []byte
	var r *http.Response

	r, err := c.request(ctx, http.MethodGet, Args)
	if err != nil {
		return res, err
	}
	defer r.Body.Close()

	res, err = io.ReadAll(r.Body)
	if err != nil {
		return res, err
	}

	if r.StatusCode != Args.StatusCode {
		return res, fmt.Errorf(string(res))
	}

	return res, nil
}

func (c RestClient) Post(ctx context.Context, Args RestArgs) (*http.Response, error) {
	var r *http.Response

	r, err := c.request(ctx, http.MethodPost, Args)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (c RestClient) Delete(ctx context.Context, Args RestArgs) (*http.Response, error) {
	var r *http.Response

	r, err := c.request(ctx, http.MethodDelete, Args)
	if err != nil {
		return r, err
	}

	return r, nil
}

// This handles the request flow and is the main logic loop for talking to the API.
func (c RestClient) request(ctx context.Context, method string, args RestArgs) (*http.Response, error) {
	var r *http.Response

	// replace spaces with url safe values
	// I have not figured out the url package yet.
	if strings.Contains(args.Url, " ") {
		u := strings.Replace(args.Url, " ", "%20", -1)
		args.Url = u
	}

	req, err := c.generateRequest(ctx, args, method)
	if err != nil {
		return r, err
	}

	r, err = c.client.Do(req)
	if err != nil {
		return r, err
	}

	err = c.checkResponse(r.StatusCode, args.StatusCode)
	if err != nil {
		return r, err
	}

	return r, nil
}

// This generates the http.Request object based in the information given.
func (c RestClient) generateRequest(ctx context.Context, Args RestArgs, method string) (*http.Request, error) {
	var req *http.Request

	body, err := c.body(Args.Body)
	if err != nil {
		return req, err
	}

	req, err = http.NewRequestWithContext(ctx, method, Args.Url, body)
	if err != nil {
		return req, err
	}

	// Add headers here later
	if Args.ContentType != "" {
		req.Header.Add("Content-Type", Args.ContentType)
	}

	return req, nil
}

// This will create a body based on the struct that was passed.
// If the body is nil, a blank reader is returned.
// This does assume that the struct contains json tags.
func (c RestClient) body(b interface{}) (io.Reader, error) {
	var r io.Reader

	if b == nil {
		return r, nil
	}

	j, err := json.Marshal(b)
	if err != nil {
		return r, err
	}

	r = bytes.NewBuffer(j)

	return r, nil
}

// This runs checks against the response that comes back.
func (c RestClient) checkResponse(received int, expected int) error {
	if received != expected {
		return errors.New(ErrInvalidStatusCode)
	}

	return nil
}
