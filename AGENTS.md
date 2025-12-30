# Project objectives
- A CLI for remote management of Yamaha audio devices on the local network.

# Implementation specifics
- Single binary executable written in Go.
- Use of the Cobra CLI library.
- Use of `dns-sd` to detect compatible devices on the local network.
- Context-efficiency for LLM agent use.

# NON-goals (we **DO NOT** need these)
- Support for non-Yamaha devices.
- Support for non-macOS environments.
- Backwards compatibility.
- End-user extensibility.
- External configuration files.
- Unit testing fixtures.

# Reference materials
- Yamaha Extended Control API
    - [Official Specification PDFs (Basic)](reference/YXC_API_Spec_Basic)
    - [Official Specification PDFs (Advanced)](reference/YXC_API_Spec_Advanced)
    - [Unofficial (LLM Generated) OpenAPI Specification](reference/yamaha-extended-control-api.yaml)

# Environment
- `go` for developing Go applications.
- `cdp` for analyzing websites with Chrome DevTools Protocol (CDP).
- `git` for fetching remote software repositories.
- `curl` for fetching content over HTTP.
- `.png`/`.jpeg`/`.gif`/`.webp` LLM perception enabled.
- `.pdf` LLM perception enabled.
- Unrestricted internet access enabled.

# Codebase style
- High density
- Flat organization
- Minimalistic
- Explicitly wired
- Internally consistent
- Artisanal

# Additional constraints
- No code comments.