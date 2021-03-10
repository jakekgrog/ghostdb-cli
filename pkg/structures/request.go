package structures

type CacheRequest struct {
	Gobj CacheObject `json:"Gobj"`
}

func NewEmptyRequest() CacheRequest {
	return CacheRequest{
		Gobj: NewEmptyCacheObject(),
	}
}
