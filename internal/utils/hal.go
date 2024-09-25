package utils

type HALResponse struct {
	Embedded interface{}            `json:"_embedded,omitempty"`
	Links    map[string]interface{} `json:"_links"`
}

func NewHALResponse(data interface{}) HALResponse {
	return HALResponse{
		Embedded: data,
		Links:    make(map[string]interface{}),
	}
}

func (h *HALResponse) AddLink(rel string, href interface{}) {
	h.Links[rel] = href
}
