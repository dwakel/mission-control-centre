package models

type Metadata struct {
	Id                   int32  `json:"id"`
	Owner                string `json:"owner"`
	ConfigurationManager string `json:"configurationManager"`
	ApplicationId        int32  `json:"applicationId"`
}
