# container-tui

## What This Is

A visual interface for Apple's `container` CLI. The primary product is a polished Go/Bubble Tea TUI inspired by Lazydocker — panels, colors, live container status, and interactive management. A local web-based UI dashboard is planned as a follow-on milestone for richer management from the browser.

## Core Value

A developer running Apple containers should be able to see what's running and control it without memorizing CLI commands — all from a single, visually clear interface.

## Requirements

### Validated

- ✓ Bubble Tea MVU TUI with left menu + right output pane — existing
- ✓ `container system status` subcommand (polls CLI, parses output fields) — existing
- ✓ `container system start` subcommand (with animated progress bar) — existing
- ✓ `container system stop` subcommand (with animated progress bar) — existing
- ✓ Lip Gloss styling (header, footer, menu, output box) — existing

### Active

- [ ] Lazydocker-style panel layout — multiple panes, clear visual hierarchy
- [ ] Color-coded container state indicators (running, stopped, error)
- [ ] Resource usage bars (CPU, memory where available from CLI)
- [ ] Real-time auto-refreshing container list
- [ ] Inline log viewing — select container, tail logs in output pane
- [ ] Additional container commands (exec, inspect, delete/remove)
- [ ] Keyboard shortcut help overlay
- [ ] Local web UI with container dashboard (auto-refresh, lifecycle management, logs, inspection)

### Out of Scope

- Native macOS app (not a GUI app — stays terminal + browser) — avoids Apple developer account requirements
- Direct daemon socket API (CLI shell-out only) — Apple's container CLI is the interface layer
- Multi-host / remote container management — local machine only
- Authentication / multi-user — single-developer local tool

## Context

The codebase already has a working Bubble Tea TUI shell. Subcommand models are self-contained (`Init/Update/View`), so adding new commands is straightforward. The current layout is functional but minimal — no colors beyond basic Lip Gloss defaults, single-column output, no panel splitting.

The Apple `container` CLI is the only interface available — there is no documented REST or socket API. All data comes from parsing CLI output.

## Constraints

- **Tech stack:** Go + Bubble Tea + Lip Gloss — extend, don't replace
- **API:** Shell out to `container system <subcommand>` — no direct daemon access
- **Platform:** macOS only (Apple Silicon primarily) — Apple's container CLI is macOS-exclusive
- **Web UI:** Local server only — no cloud, no auth required

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| TUI before web UI | User's primary context is terminal; polish what exists first | — Pending |
| Lazydocker as visual reference | Panels + colors + keyboard-driven — proven UX pattern for container management | — Pending |
| Shell CLI for all data | Only available interface; simplifies auth/discovery | — Pending |

---
*Last updated: 2026-02-22 after initialization*
