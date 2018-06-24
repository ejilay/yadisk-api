package yadisk

import (
	"encoding/json"
	"io"
	"net/url"
	"fmt"
)

// ShareResponse struct is returned by the API for upload request.
type ShareResponse struct {
	HRef      string `json:"href"`
	Method    string `json:"method"`
	Templated bool   `json:"templated"`
}

// Share will put specified data to Yandex.Disk.
func (a *API) Share(remotePath string) (string, error) {
	ur, err := a.ShareRequest(remotePath)
	if err != nil {
		return ur, err
	}

	return ur, nil
}

// ShareRequest will make an share request and return a URL to upload data to.
func (a *API) ShareRequest(remotePath string) (string, error) {
	values := url.Values{}
	values.Add("path", remotePath)

	req, err := a.scopedRequest("PUT", "/v1/disk/resources/publish?"+values.Encode(), nil)
	if err != nil {
		return "", err
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	if err := CheckAPIError(resp); err != nil {
		return "", err
	}

	defer resp.Body.Close()
	ur, err := a.GetMeta(remotePath)
	if err != nil {
		return "", err
	}
	if ur.PublicURL == nil {
		return "", fmt.Errorf("unable to publish %q", remotePath)
	}
	return *ur.PublicURL, nil
}

// ParseUploadResponse tries to read and parse ShareResponse struct.
func ParseShareResponse(data io.Reader) (ShareResponse, error) {
	dec := json.NewDecoder(data)
	var ur ShareResponse

	if err := dec.Decode(&ur); err == io.EOF {
		// ok
	} else if err != nil {
		return ur, err
	}

	// TODO: check if there is any trash data after JSON and crash if there is.

	return ur, nil
}
