package response

type PageInfoResponse struct {
	Data  interface{}
	Total int64 `json:"total"`
}
