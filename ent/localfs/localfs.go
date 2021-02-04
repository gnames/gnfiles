package localfs

import "time"

type MetaData struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Version
	PastVersions []Version `json:"pastVersions"`
}

type Version struct {
	ID       string    `json:"cid"`
	SHA      string    `json:"sha"`
	DateTime time.Time `json:"dateTime"`
	Size     int64     `json:"size"`
}
