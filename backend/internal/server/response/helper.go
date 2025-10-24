package response

func New(resourceType string, attributes interface{}) APIResponse {
	return APIResponse{
		Data: DataResponse{
			Type:       resourceType,
			Attributes: attributes,
		},
	}
}
