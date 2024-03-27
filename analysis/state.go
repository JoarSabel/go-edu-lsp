package analysis

import (
	"fmt"
	"lsp-test-project/lsp"
	"strings"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{
		Documents: make(map[string]string),
	}
}

func getDiagnosticsForFile(text string) []lsp.Diagnostic{
	diagnostics := []lsp.Diagnostic{}
	for i, line := range strings.Split(text, "\n") {
		if strings.Contains(line, "VS Code") {
			idx := strings.Index(line, "VS Code")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range: LineRange(i, idx, idx+len("VS Code")),
				Severity: 1,
				Source:   "Common sense",
				Message:  "Please refrain from such foul things.",
			})
		}
		if strings.Contains(line, "PHP") {
			idx := strings.Index(line, "PHP")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range: LineRange(i, idx, idx+len("PHP")),
				Severity: 4,
				Source:   "Common sense",
				Message:  "Wonderful stuff, keep it up!",
			})
		}
	}
	return diagnostics
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text)
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	// irl would lookup like type analysis

	document := s.Documents[uri]

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, Characters: %d", uri, len(document)),
		},
	}
}

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	// irl would lookup defintion

	// document := s.Documents[uri]

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}

func (s *State) TextDocumentCodeAction(id int, uri string) lsp.TextDocumentCodeActionResponse {
	text := s.Documents[uri]

	actions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "NeoVim",
				},
			}
			actions = append(actions, lsp.CodeAction{
				Title: "Replace VS C*de with a _superior_ editor",
				Edit:  &lsp.WorkspaceEdit{Changes: replaceChange},
			})

			censorChange := map[string][]lsp.TextEdit{}
			censorChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "VS C*ode",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Censor VS C*de",
				Edit:  &lsp.WorkspaceEdit{Changes: censorChange},
			})
		}
	}
	return lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}
}

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {
	// text := s.Documents[uri]

	items := []lsp.CompletionItem{
		{
			Label:         "Neovim",
			Detail:        "The best editor *tips fedora*",
			Documentation: "You see using a mouse if an utterly futile endeavor, you should use the keyboard for everything. Why? BECAUSE IT MAKES YOU LOOK COOL",
		},
	}
	return lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}
}

func LineRange(row, idx, i int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      row,
			Character: idx,
		},
		End: lsp.Position{
			Line:      row,
			Character: i,
		},
	}
}
