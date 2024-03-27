package lsp

type InitializeRequest struct {
	Request
	Params IntializeRequestParams `json:"params"`
}

type IntializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// and a lot mote stuff
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type IntializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities *ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo          `json:"serverInfo"`
}

type ServerCapabilities struct {
	TextDocumentSync   int  `json:"textDocumentSync"`
	HoverProvider      bool `json:"hoverProvider"`
	DeifinitonProvider bool `json:"definitionProvider"`
	CodeActionProvider bool `json:"codeActionProvider"`
	CompletionProvider map[string]any `json:"completionProvider"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) IntializeResponse {
	return IntializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: &ServerCapabilities{
				TextDocumentSync:   1,
				HoverProvider:      true,
				DeifinitonProvider: true,
				CodeActionProvider: true,
				CompletionProvider: map[string]any{},
			},
			ServerInfo: ServerInfo{
				Name:    "lsp-test-project",
				Version: "rc-0.0.1-pre-alpha-beta-rc-0.1",
			},
		},
	}
}
