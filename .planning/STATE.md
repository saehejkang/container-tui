# Project State

## Project Reference

See: .planning/PROJECT.md (updated 2026-02-22)

**Core value:** A developer running Apple containers should see what's running and control it without memorizing CLI commands — from a single, visually clear interface.
**Current focus:** Phase 1 - Panel Layout & Visuals

## Current Position

Phase: 1 of 2 (Panel Layout & Visuals)
Plan: 3 of 4 in current phase
Status: In progress
Last activity: 2026-02-23 — Completed 01-03: Model rearchitect (container list, auto-refresh, keyboard nav)

Progress: [███░░░░░░░] 30%

## Performance Metrics

**Velocity:**
- Total plans completed: 3
- Average duration: 5 min
- Total execution time: 0.2 hours

**By Phase:**

| Phase | Plans | Total | Avg/Plan |
|-------|-------|-------|----------|
| 01-panel-layout-visuals | 3/4 | 16 min | 5 min |

**Recent Trend:**
- Last 5 plans: 7 min, 7 min, 2 min
- Trend: accelerating

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
- [01-03]: Removed subcommands/ActiveView architecture — model now owns container list directly
- [01-03]: 5-second tea.Tick self-perpetuating refresh cycle (each tick re-schedules next)
- [01-03]: WindowSizeMsg guarded (Width > 0 && Height > 0) to prevent zero-dimension layout math
- [01-03]: SelectedIndex clamped on FetchContainersMsg to handle list shrink between refreshes

### Pending Todos

None.

### Blockers/Concerns

- [Codebase]: view.go is now a stub renderer — Phase 1 panel layout work (01-02 context) still needs full two-panel layout composition
- [Resolved 01-01]: Go 1.20.4 was installed but project requires Go 1.25+; upgraded to 1.26.0 via brew
- [Resolved 01-03]: Width/Height defaulted to 0 before first WindowSizeMsg — guarded in WindowSizeMsg handler
- [Resolved 01-03]: ActiveView not initialized on startup — eliminated by removing ActiveView; containers now load automatically via Init()

## Session Continuity

Last session: 2026-02-23
Stopped at: Completed 01-03-PLAN.md (Model rearchitect)
Resume file: None
