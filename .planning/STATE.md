# Project State

## Project Reference

See: .planning/PROJECT.md (updated 2026-02-22)

**Core value:** A developer running Apple containers should see what's running and control it without memorizing CLI commands — from a single, visually clear interface.
**Current focus:** Phase 1 - Panel Layout & Visuals

## Current Position

Phase: 1 of 2 (Panel Layout & Visuals)
Plan: 1 of 4 in current phase
Status: In progress
Last activity: 2026-02-23 — Completed 01-01: Container data model (Container struct + FetchContainers cmd)

Progress: [█░░░░░░░░░] 10%

## Performance Metrics

**Velocity:**
- Total plans completed: 1
- Average duration: 7 min
- Total execution time: 0.1 hours

**By Phase:**

| Phase | Plans | Total | Avg/Plan |
|-------|-------|-------|----------|
| 01-panel-layout-visuals | 1/4 | 7 min | 7 min |

**Recent Trend:**
- Last 5 plans: 7 min
- Trend: -

*Updated after each plan completion*

## Accumulated Context

### Decisions

Decisions are logged in PROJECT.md Key Decisions table.
Recent decisions affecting current work:

- [Init]: TUI before web UI — polish existing terminal experience first
- [Init]: Lazydocker as visual reference — panels + colors + keyboard-driven UX
- [Init]: Shell CLI for all data — only available interface from apple/container
- [01-01]: Used pkg.RunCommand wrapper consistent with StatusModel pattern
- [01-01]: Status normalized to three canonical values: running/error/stopped
- [01-01]: CPU/memory parsed best-effort; zero-value on missing or malformed fields

### Pending Todos

None.

### Blockers/Concerns

- [Codebase]: Current layout is single-column output; Phase 1 requires rearchitecting view.go and the system model's layout composition
- [Codebase]: Width/Height default to 0 before first WindowSizeMsg — guard layout math before rebuilding panels
- [Codebase]: ActiveView not initialized on startup — status data doesn't load until user manually selects it; address during Phase 1 auto-refresh work
- [Resolved 01-01]: Go 1.20.4 was installed but project requires Go 1.25+; upgraded to 1.26.0 via brew

## Session Continuity

Last session: 2026-02-23
Stopped at: Completed 01-01-PLAN.md (Container data model)
Resume file: None
