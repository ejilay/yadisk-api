package yadisk

import "time"

type Resource struct {
	// Key of a published resource.
	// It is included in the response only if
	// the specified file or folder is published.
	PublicKey *string `json:"public_key"`

	// Resources directly contained in the folder.
	// It is included in the response only
	// when requesting metainformation about a folder.
	Embedded *ResourceList `json:"_embedded"`

	// Resource name
	Name string `json:"name"`

	// The date and time of the resource was created.
	// JSON data is in ISO 8601 format.
	Created time.Time `json:"created"`

	// Object with the user defined attributes
	CustomProperties *map[string]string `json:"custom_properties"`

	// Link to a published resource.
	// It is included in the response only if
	// the specified file or folder is published.
	PublicURL *string `json:"public_url"`

	// Path to the resource before it was moved to the Trash.
	// Included in the response only for a request
	// for metainformation about a resource in the Trash.
	OriginPath *string `json:"origin_path"`

	// The date and time the resource was modified.
	// JSON data is in ISO 8601 format.
	Modified time.Time `json:"modified"`

	// Full path to the resource on Disk.
	// In metainformation for a published folder, paths are relative
	// to the folder itself. For published files, the value
	// of the key is always "/". For a resource located in the Trash,
	// this attribute may have the unique ID appended to it
	// (for example, trash:/foo_1408546879).
	// Using this ID, the resource can be differentiated from other
	// deleted resources with the same name.
	Path string `json:"path"`

	// MD5 hash of the file.
	MD5 string `json:"md5"`

	// sha256 hash of the file.
	SHA256 string `json:"sha256"`

	// Revision number of the resource.
	Revision uint `json:"revision"`

	// Unique resource id.
	ResourceID string `json:"resource_id"`

	// Resource type:
	// * "dir" - folder
	// * "file" - file
	Type string `json:"type"`

	// The media type of the file (e.g. "image").
	MediaType string `json:"media_type"`

	// The MIME type of the file (e.g. image/jpeg).
	MimeType string `json:"mime_type"`

	// File size.
	Size uint `json:"size"`
}

// ResourceList is a list of resources contained in the folder.
// Contains Resource objects and list properties.
// https://tech.yandex.com/disk/api/reference/response-objects-docpage/#resourcelist
type ResourceList struct {
	// The field used for sorting the list.
	Sort string `json:"sort"`

	// The key of a published folder that contains resources from this list.
	// It is included in the response only when requesting
	// metainformation about a public folder.
	PublicKey string `json:"public_key"`

	// Array of resources contained in the folder.
	Items []Resource `json:"items"`

	// The path to the folder whose contents are described
	// in this ResourceList object.
	// For a public folder, the value of the attribute is always "/".
	Path string `json:"path"`

	// The maximum number of items in the items array; set in the request.
	Limit uint `json:"limit"`

	// How much to offset the beginning of the
	// list from the first resource in the folder.
	Offset uint `json:"offset"`

	// The total number of resources in the folder.
	Total uint `json:"total"`
}
