# External Integrations

**Analysis Date:** 2026-02-22

## APIs & External Services

**Container System Management:**
- `apple/container` CLI - Manages container lifecycle (start, stop, status operations)
  - SDK/Client: None (shell execution via `os/exec`)
  - Auth: None (CLI authentication handled by `apple/container` itself)
  - Integration method: Direct subprocess execution of `container system <subcommand>`

## Data Storage

**Databases:**
- None - Application is stateless

**File Storage:**
- Local filesystem only - No persistent data storage
- No remote storage integrations

**Caching:**
- None - Real-time data fetched on each subcommand execution

## Authentication & Identity

**Auth Provider:**
- None - Application has no authentication layer
- Delegates auth to `apple/container` CLI (auth handled by underlying system)

## Monitoring & Observability

**Error Tracking:**
- None - Application writes errors to stderr only

**Logs:**
- Console output to stdout/stderr - No log aggregation or persistence
- Status model (`ui/system/subcommands/status.go`) displays command output in real-time

## CI/CD & Deployment

**Hosting:**
- Not applicable - Standalone binary application
- Deployment: Users build locally with `go build` or via release binaries

**CI Pipeline:**
- Not detected - No GitHub Actions, GitLab CI, or other CI configuration found

## Environment Configuration

**Required env vars:**
- None explicitly required
- System PATH must include `apple/container` CLI

**Secrets location:**
- Not applicable - No secrets management in codebase

## Webhooks & Callbacks

**Incoming:**
- None - Application is not a server

**Outgoing:**
- None - Application only executes local CLI commands

## Command Execution Flow

The application communicates exclusively with the local `apple/container` CLI:

**Status Command Flow:**
1. User selects "status" from menu
2. `StatusModel.Init()` calls `pkg.RunCommand("container", "system", "status")`
3. Combined stdout+stderr captured from subprocess
4. Output parsed into key-value fields and displayed in UI

**Start Command Flow:**
1. User selects "start" from menu
2. `StartModel.Init()` executes `os/exec.Command("container", "system", "start")`
3. Animated progress bar (0-95%) displayed while command runs
4. On completion, status updated (success/failure)

**Stop Command Flow:**
1. User selects "stop" from menu
2. `StopModel.Init()` executes `os/exec.Command("container", "system", "stop")`
3. Animated progress bar (0-95%) displayed while command runs
4. On completion, status updated (success/failure)

---

*Integration audit: 2026-02-22*
