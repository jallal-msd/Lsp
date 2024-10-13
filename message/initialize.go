package message

type InitializeRequest struct {
	Request
	Params InitialzeParams `json:"params"`
}
type InitialzeParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
}
type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
type InitializeResponse struct {
	Response
	Result InitialzeResult `json:"result"`
}
type InitialzeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct {
	TextDocumentSync int `json:"textDocumentSync"`
}
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitialzeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync: 1,
			},
			ServerInfo: ServerInfo{
				Name:    "LSPTEST",
				Version: "0.0.1-beta.final",
			},
		},
	}
}
