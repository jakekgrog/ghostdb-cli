package structures

type CacheRequest struct {
	Gobj CacheObject `json:"Gobj"`
}

func NewEmptyRequest() CacheRequest {
	return CacheRequest{
		Gobj: NewEmptyCacheObject(),
	}
}

func NewStoreRequest(key, value string) CacheRequest {
	return CacheRequest{
		Gobj: NewKVCacheObject(key, value),
	}
}