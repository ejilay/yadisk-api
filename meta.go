package yadisk

import (
	"net/url"
	"io"
	"encoding/json"
)

// GetMeta will make an request and return a META
func (a *API) GetMeta(remotePath string) (Resource, error) {

	values := url.Values{}
	values.Add("path", remotePath)

	req, err := a.scopedRequest("GET", "/v1/disk/resources?"+values.Encode(), nil)
	if err != nil {
	return Resource{}, err
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
	return Resource{}, err
	}
	if err := CheckAPIError(resp); err != nil {
	return Resource{}, err
	}

	defer resp.Body.Close()
	ur, err := ParseMetaResponse(resp.Body)
	if err != nil {
	return Resource{}, err
	}

	return ur, nil

}

// ParseMetaResponse tries to read and parse Resource struct.
func ParseMetaResponse(data io.Reader) (Resource, error) {
	dec := json.NewDecoder(data)
	var ur Resource

	if err := dec.Decode(&ur); err == io.EOF {
		// ok
	} else if err != nil {
		return ur, err
	}

	// TODO: check if there is any trash data after JSON and crash if there is.

	return ur, nil
}

