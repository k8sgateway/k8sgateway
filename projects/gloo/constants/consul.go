package constants

const (
	EndpointMetadataMatchTrue  = "1"
	EndpointMetadataMatchFalse = "0"

	// We use these prefixes to avoid shadowing in case a data center name is the same as a tag name
	TagKeyPrefix        = "tag_"
	DataCenterKeyPrefix = "dc_"
)
