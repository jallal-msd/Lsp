package message

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`

	//Error
	//Result
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
