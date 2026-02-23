---
phase: 01-panel-layout-visuals
plan: "03"
subsystem: ui
tags: [go, bubbletea, model-rearchitect, container-list, auto-refresh, keyboard-nav]

# Dependency graph
requires:
  - "pkg.Container struct (01-01)"
  - "pkg.FetchContainers tea.Cmd (01-01)"
provides:
  - "Rearchitected system.Model with Containers, SelectedIndex, ShowHelp, Width, Height, Loading, Error"
  - "Auto-refresh via 5-second tea.Tick cycle"
  - "Keyboard navigation: up/k, down/j, ?, esc, q/ctrl+c"
  - "WindowSizeMsg guard: layout math only executes with non-zero dimensions"
affects: [01-panel-layout-visuals]

# Tech tracking
tech-stack:
  added: []
  patterns: [tea.Tick-auto-refresh, WindowSizeMsg-guard, help-toggle, selected-index-clamp]

key-files:
  created: []
  modified:
    - ui/system/model.go
    - ui/system/view.go
    - main.go

key-decisions:
  - "Removed subcommands/ActiveView architecture entirely — model now owns container list directly"
  - "5-second tea.Tick refresh re-issues FetchContainers and re-schedules next tick"
  - "WindowSizeMsg guarded (Width > 0 && Height > 0) to prevent layout math with zero dimensions"
  - "SelectedIndex clamped on FetchContainersMsg to handle list shrink between refreshes"

requirements-completed: [TUI-01]

# Metrics
duration: 2min
completed: 2026-02-23
---

# Phase 01 Plan 03: Model Rearchitect Summary

**Root system.Model rearchitected from subcommand-menu owner to container-list owner with FetchContainers on Init, 5-second auto-refresh tick, cursor navigation, and ShowHelp toggle**

## Performance

- **Duration:** 2 min
- **Started:** 2026-02-23T05:12:56Z
- **Completed:** 2026-02-23T05:14:04Z
- **Tasks:** 2
- **Files modified:** 3

## Accomplishments

- Replaced Model struct fields (Cursor, Subcommands, ActiveView) with (Containers, SelectedIndex, ShowHelp, Width, Height, Loading, Error)
- NewSystemModel() constructor now takes no arguments and sets Loading: true
- Init() returns tea.Batch of pkg.FetchContainers() and tickRefresh() (5-second tea.Tick)
- Update() handles FetchContainersMsg (set containers, clamp SelectedIndex), refreshTickMsg (re-fetch + re-schedule), WindowSizeMsg (guarded), and KeyMsg (?, esc, up/k, down/j, q/ctrl+c)
- main.go updated: removed subcommandsList variable, calls NewSystemModel() with no arguments
- view.go auto-fixed to use new Model fields (Rule 3 blocking deviation)

## Task Commits

Each task was committed atomically:

1. **Task 1: Rearchitect Model struct and Update handler** - `7323715` (feat)
2. **Task 2: Update main.go for new constructor** - `e24e0b0` (feat)

**Plan metadata:** (docs commit follows)

## Files Created/Modified

- `ui/system/model.go` - Complete rearchitect: new Model struct, Init/Update/View, tickRefresh
- `ui/system/view.go` - Auto-fixed to remove old field references (ActiveView, Subcommands, Cursor)
- `main.go` - Removed subcommandsList, calls NewSystemModel() with no args

## Decisions Made

- Removed the entire subcommands/ActiveView delegation architecture — container list is now first-class model state
- 5-second tick uses tea.Tick pattern: each refreshTickMsg re-issues FetchContainers and re-schedules the next tick (self-perpetuating refresh cycle)
- WindowSizeMsg guard prevents layout math from executing before first terminal dimensions arrive
- SelectedIndex clamp on FetchContainersMsg handles edge case where running containers decrease between refreshes

## Deviations from Plan

### Auto-fixed Issues

**1. [Rule 3 - Blocking] Updated view.go to use new Model fields**
- **Found during:** Task 1 (build verification of ui/system package)
- **Issue:** view.go referenced m.ActiveView, m.Subcommands, and m.Cursor — all removed from the new Model struct. Also imported ui/system/subcommands package which is no longer relevant to the system view layer
- **Fix:** Rewrote view.go to render container list using m.Containers, m.SelectedIndex, m.Loading, m.Error, m.Width, m.Height — consistent with the new Model shape
- **Files modified:** ui/system/view.go
- **Committed in:** 7323715 (Task 1 commit)

---

**Total deviations:** 1 auto-fixed (1 blocking — view.go used old Model fields)
**Impact on plan:** Necessary for compilation. view.go update is expected work for Phase 1 panel layout; bringing it into this task keeps the package coherent. No scope creep.

## Verification Results

- `go build ./ui/system/` — PASS (after Task 1)
- `go build ./...` — PASS (after Task 2)
- `go vet ./...` — PASS
- Model struct fields: Containers, SelectedIndex, ShowHelp, Width, Height, Loading, Error — all present
- model.go line count: 92 (>= 80 minimum)
- Init() contains pkg.FetchContainers() and tickRefresh() — confirmed
- ShowHelp toggle on '?' — confirmed
- WindowSizeMsg guard (Width > 0 && Height > 0) — confirmed
- SelectedIndex navigation with bounds — confirmed

## Self-Check: PASSED

- FOUND: ui/system/model.go
- FOUND: ui/system/view.go
- FOUND: main.go
- FOUND: .planning/phases/01-panel-layout-visuals/01-03-SUMMARY.md
- FOUND: commit 7323715
- FOUND: commit e24e0b0
