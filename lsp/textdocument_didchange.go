package lsp

type DidChangeTextDocumentNotification struct {
	Notification
	Params DidChangeTextDocumentParams `json:"params"`
}
 
type DidChangeTextDocumentParams struct {
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`
	ConentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}
