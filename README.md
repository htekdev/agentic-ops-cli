# ⚠️ DEPRECATED - This repository has been renamed

> **This project has been rebranded to `hookflow` and moved to a new location.**

## New Repository

| Purpose | New Location |
|---------|--------------|
| **CLI** | [htekdev/hookflow](https://github.com/htekdev/hookflow) |

## Migration

```bash
# Install new CLI via npm
npm install -g hookflow

# Or via go install
go install github.com/htekdev/hookflow/cmd/hookflow@latest
```

## What Changed?

- **Name**: `agentic-ops` → `hookflow`
- **Workflow Directory**: `.github/agent-workflows/` → `.github/hooks/`
- **Module Path**: `github.com/htekdev/agentic-ops-cli` → `github.com/htekdev/hookflow`

The syntax and functionality remain the same.

---

<details>
<summary>📜 Original README (archived)</summary>

# Agentic-Ops CLI

The command-line interface for Agentic-Ops - local agent workflow governance.

## Installation

```bash
# npm
npm install -g agentic-ops

# go install
go install github.com/htekdev/agentic-ops-cli/cmd/agentic-ops@latest
```

## Commands

- `agentic-ops discover` - Find workflow files
- `agentic-ops validate` - Validate workflow YAML
- `agentic-ops run` - Execute workflows for events

## License

MIT

</details>