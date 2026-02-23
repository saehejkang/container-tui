# Roadmap: container-tui

## Overview

Transform the existing functional-but-minimal Bubble Tea shell into a polished, Lazydocker-style TUI. Phase 1 rebuilds the layout into a real multi-pane interface with color-coded container status, resource bars, and a keyboard help overlay. Phase 2 adds container selection and detail inspection — the interactive capability that makes the tool useful for day-to-day container management.

## Phases

**Phase Numbering:**
- Integer phases (1, 2, 3): Planned milestone work
- Decimal phases (2.1, 2.2): Urgent insertions (marked with INSERTED)

Decimal phases appear between their surrounding integers in numeric order.

- [ ] **Phase 1: Panel Layout & Visuals** - Rebuild the TUI into a Lazydocker-style multi-pane layout with live container status, color coding, resource bars, and a keyboard help overlay
- [ ] **Phase 2: Container Selection & Detail View** - Enable navigating the container list and viewing full container details in the right pane

## Phase Details

### Phase 1: Panel Layout & Visuals
**Goal**: Users see a polished, information-dense TUI with live container status at a glance
**Depends on**: Nothing (first phase)
**Requirements**: TUI-01, TUI-02, TUI-03, TUI-04
**Success Criteria** (what must be TRUE):
  1. The TUI displays a persistent left pane listing all containers and a right pane for detail/output — visible simultaneously without switching menus
  2. Each container in the list shows a color-coded status indicator: green for running, gray for stopped, red for error
  3. Each container in the list shows CPU and memory usage bars when that data is available from the CLI
  4. Pressing `?` opens a keyboard shortcut help overlay that lists all available keybindings, and pressing `?` or `Esc` dismisses it
**Plans**: 4 plans

Plans:
- [ ] 01-01-PLAN.md — Container data model and FetchContainers CLI command
- [ ] 01-02-PLAN.md — Container row rendering: status colors and resource bars
- [ ] 01-03-PLAN.md — Root model rearchitecture: container list state, navigation, auto-refresh
- [ ] 01-04-PLAN.md — Layout assembly: two-pane view, help overlay, visual verification

### Phase 2: Container Selection & Detail View
**Goal**: Users can navigate the container list and inspect any container's full configuration details
**Depends on**: Phase 1
**Requirements**: TUI-05
**Success Criteria** (what must be TRUE):
  1. User can move up and down through the container list using arrow keys or `j`/`k`
  2. Selecting a container (pressing Enter or navigating to it) populates the right pane with that container's details: config, env vars, mounts, and network info
  3. The detail view updates when the user selects a different container without restarting or re-entering the TUI
**Plans**: TBD

## Progress

**Execution Order:**
Phases execute in numeric order: 1 → 2

| Phase | Plans Complete | Status | Completed |
|-------|----------------|--------|-----------|
| 1. Panel Layout & Visuals | 0/4 | Not started | - |
| 2. Container Selection & Detail View | 0/TBD | Not started | - |
