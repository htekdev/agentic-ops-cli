# agentic-ops-cli

Command-line interface for the [agentic-ops](https://github.com/htekdev/agentic-ops) local agent workflow engine.

## Overview

`agentic-ops-cli` is a Go CLI that discovers, validates, and executes local workflows for AI agents. It provides a GitHub Actions-like experience for defining agent governance gates.

## Installation

### Using Go

```bash
go install github.com/htekdev/agentic-ops-cli/cmd/agentic-ops@latest
```

### Download Binary

Download pre-built binaries from the [Releases](https://github.com/htekdev/agentic-ops-cli/releases) page.

## Usage

```bash
# Discover workflows in the current directory
agentic-ops discover

# Validate workflow files
agentic-ops validate

# Run workflows for an event (used by hooks)
agentic-ops run --event '{"hook":{"type":"preToolUse","tool":{"name":"edit"}}}'

# Run a specific workflow
agentic-ops run --workflow my-workflow --event '{"file":{"path":"src/main.ts"}}'
```

## Workflow Syntax

Workflows are defined in `.github/agent-workflows/*.yml`:

```yaml
name: Block Sensitive Files
description: Prevent edits to sensitive files

on:
  tool:
    name: edit
    args:
      path: '**/*.env*'

blocking: true

steps:
  - name: Deny edit
    run: |
      echo "Cannot edit sensitive files"
      exit 1
```

## Event Types

| Trigger | Description |
|---------|-------------|
| `hooks` | Match by hook type (preToolUse, postToolUse) |
| `tool` | Match specific tools with argument patterns |
| `tools` | Match multiple tool configurations |
| `file` | Match file creation/edit events |
| `commit` | Match git commit events |
| `push` | Match git push events |

## Expression Engine

Supports `${{ }}` expressions with GitHub Actions parity:

```yaml
steps:
  - name: Conditional step
    if: ${{ endsWith(event.file.path, '.ts') }}
    run: echo "TypeScript file: ${{ event.file.path }}"
```

### Built-in Functions

- `contains(search, item)` - Check if string/array contains item
- `startsWith(str, value)` - String starts with value
- `endsWith(str, value)` - String ends with value
- `format(str, ...args)` - String formatting
- `join(array, sep)` - Join array to string
- `toJSON(value)` - Convert to JSON string
- `fromJSON(str)` - Parse JSON string
- `always()` - Always true
- `success()` - Previous steps succeeded
- `failure()` - Previous step failed

## Development

```bash
# Build
go build -o bin/agentic-ops ./cmd/agentic-ops

# Test
go test ./... -v

# Test with coverage
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Related Projects

- [agentic-ops](https://github.com/htekdev/agentic-ops) - Copilot plugin that uses this CLI

## License

MIT
