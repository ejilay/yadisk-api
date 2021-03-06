package yadisk

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
)

// ShareResponse struct is returned by the API for upload request.
type UploadResponse struct {
	HRef      string `json:"href"`
	Method    string `json:"method"`
	Templated bool   `json:"templated"`
}

// Upload will put specified data to Yandex.Disk.
func (a *API) Upload(data io.Reader, remotePath string, overwrite bool) error {
	ur, err := a.UploadRequest(remotePath, overwrite)
	if err != nil {
		return err
	}

	if err := PerformUpload(ur.HRef, data, a.HTTPClient); err != nil {
		return err
	}

	return nil
}

// UploadRequest will make an upload request and return a URL to upload data to.
func (a *API) UploadRequest(remotePath string, overwrite bool) (UploadResponse, error) {
	values := url.Values{}
	values.Add("path", remotePath)
	values.Add("overwrite", strconv.FormatBool(overwrite))

	req, err := a.scopedRequest("GET", "/v1/disk/resources/upload?"+values.Encode(), nil)
	if err != nil {
		return UploadResponse{}, err
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return UploadResponse{}, err
	}
	if err := CheckAPIError(resp); err != nil {
		return UploadResponse{}, err
	}

	defer resp.Body.Close()
	ur, err := ParseUploadResponse(resp.Body)
	if err != nil {
		return UploadResponse{}, err
	}

	return ur, nil
}

// ParseUploadResponse tries to read and parse ShareResponse struct.
func ParseUploadResponse(data io.Reader) (UploadResponse, error) {
	dec := json.NewDecoder(data)
	var ur UploadResponse

	if err := dec.Decode(&ur); err == io.EOF {
		// ok
	} else if err != nil {
		return ur, err
	}

	// TODO: check if there is any trash data after JSON and crash if there is.

	return ur, nil
}
