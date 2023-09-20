# Core


## Development (Core)

### Objective
- Clone to the OG omega-triage-portal 
- Learn new technologies

### Use Case
1. View SARIF File, Online
1. View SARIF File, Offline
1. Triage Github Packages

### Requirements (TODOs)
- [ ] Templates
  - [ ] Usage of HTMX
  - [ ] [PURECSS Frontend Framework](https://purecss.io/)
- [ ] GraphQL API
- [ ] Dockerize Core
- [ ] Routing

### Security Requirements
- Ensure no leakage of sensitive files
- Pentest GraphQL Endpoint

### Plus
- [ ] Gitignore Emacs Backups on the `/assets` FileServer 


---

## Development (Infra)
- This is a bonus, if it needs to go on a server

### Objective
- To setup a kubernetes deployment that hosts the core and all it's external components for consumption

### Use Case
- TBD

### Diagram
- TBD

### Requirements (TODOs) 
- [ ] Kubernetes Deployments and Services
- [ ] Research good db for SARIFs
- [ ] Redis for Caching

### Security Requirements
- Double-check Configs for no gap
- Flesh this out more

--- 

### Notes
#### **Setting up `gopls` with Emacs LSP (Eglot)**

1. Install latest stable version of `gopls`
   - `go install golang.org/x/tools/gopls@latest`
1. Double check `gopls` exists
   - `ls $(go env GOPATH)/bin`
- Manual gopls running (at localhost:56902)
   - `$(go env GOPATH)/bin/gopls service -listen 127.0.0.1:56902`
- Point eglot at gopls
  - Using the path at `go env GOPATH`, do `<GOPATH>/bin/gopls`
