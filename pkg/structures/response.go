package structures

type CacheResponse struct {
	// GhostDB Cache Object
	Gobj    CacheObject `json:"Gobj"`
	// Status of the result - 0 or 1 - fail or succeed
	Status  int32 `json:"Status"`
	// Message contains any textual data that is to 
	// be sent back
	Message string `json:"Message"`
	// Error message returned if something went wrong
	// during command execution
	Error   string `json:"Error"`
}
