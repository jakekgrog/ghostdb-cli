package structures

type CacheObject struct {
	Key   string `json:"Key"`
	Value interface{} `json:"Value"`
	TTL   int64 `json:"TTL,string"`
}

func NewEmptyCacheObject() CacheObject {
	return CacheObject {
		Key: "",
		Value: nil,
		TTL: -1,
	}
}
