# GHP — Good Hygiene Practices

> **Note:** this repository contains an older, standalone GHP CLI implementation.
> The maintained framework and current releases live at
> [GHP-GoLang-Framework/ghp](https://github.com/GHP-GoLang-Framework/ghp).

A Go CLI for managing GHP web projects: development server, build and
documentation from a template.

## Commands

| Command | Description |
| --- | --- |
| `ghp dev` | Start the development server |
| `ghp build` | Build the current project |
| `ghp help` | Show help |

## Usage

```bash
go run ./cmd/ghp help
```

The base project uses a template located at [`docs/template.ghp`](./docs/template.ghp).
