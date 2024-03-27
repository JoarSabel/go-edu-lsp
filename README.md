# Test build lsp

> with go

Neovim script to load the LSP we build, has to load pretty much last, I think

```lua
local client = vim.lsp.start_client {
  name = "lsp-test-project",
  cmd = { "path/to/executable" },
  -- on_attach = require("vim.lsp").on_attach,
}
if not client then
  vim.notify "You done did ducked it"
  return
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = "markdown",
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end
})

```

PHP
Here is some text with VS Code in it.
