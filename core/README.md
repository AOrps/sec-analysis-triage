# Core


## Development

### TODOs
- [ ] Templates
- [ ] GraphQL API
- [ ] Dockerize
- [ ] Kubernetes
- [ ] Redis
- [ ] Research good db for SARIFs
- [ ] Routing


### Setting up `gopls` with Emacs LSP (Eglot)

1. Install latest stable version of `gopls`
   - `go install golang.org/x/tools/gopls@latest`
1. Double check `gopls` exists
   - `ls $(go env GOPATH)/bin`
- Manual gopls running (at localhost:56902)
   - `$(go env GOPATH)/bin/gopls service -listen 127.0.0.1:56902`
- Point eglot at gopls
  - Using the path at `go env GOPATH`, do `<GOPATH>/bin/gopls`
