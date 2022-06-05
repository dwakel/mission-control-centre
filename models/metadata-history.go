package models

import "time"

type MetadataHistory struct {
	Id                   int32     `json:"id"`
	MetadataId           int32     `json:"metadataId"`
	Owner                string    `json:"owner"`
	ConfigurationManager string    `json:"configurationManager"`
	ApplicationId        int32     `json:"applicationId"`
	Version              int32     `json:"version"`
	Updated              time.Time `json:"updated"`
}
