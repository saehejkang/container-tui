# Requirements: container-tui

**Defined:** 2026-02-22
**Core Value:** A developer running Apple containers should see what's running and control it without memorizing CLI commands — from a single, visually clear interface.

## v1 Requirements

### TUI Layout & Visuals

- [x] **TUI-01**: User sees a multi-pane layout — container list on left, detail/output pane on right
- [x] **TUI-02**: Container list shows color-coded status indicators (running=green, stopped=gray, error=red)
- [x] **TUI-03**: Container list shows resource usage bars (CPU, memory) where data is available from CLI
- [ ] **TUI-04**: User can press `?` to open a keyboard shortcut help overlay

### TUI Container Management

- [ ] **TUI-05**: User can select a container and view its details (config, env vars, mounts, network info)

## v2 Requirements

### TUI Container Management (Deferred)

- **TUI-06**: User can start/stop containers directly from the container list
- **TUI-07**: User can tail container logs inline in the output pane
- **TUI-08**: User can delete/remove a stopped container

### Web UI

- **WEB-01**: Local web server serves a container dashboard
- **WEB-02**: Dashboard shows all containers with auto-refresh (polling every few seconds)
- **WEB-03**: User can start, stop, and delete containers from the web UI
- **WEB-04**: User can view live container logs in the browser
- **WEB-05**: User can inspect container details (config, env vars, mounts) in the browser

## Out of Scope

| Feature | Reason |
|---------|--------|
| Native macOS app (non-terminal) | Avoids Apple developer account; terminal + browser covers all use cases |
| Direct daemon socket/API | Apple's container CLI is the only available interface |
| Remote/multi-host management | Local developer tool only |
| Authentication / multi-user | Single developer, local machine |
| Real-time WebSocket streaming (web) | Auto-refresh polling sufficient for v1 web UI |

## Traceability

| Requirement | Phase | Status |
|-------------|-------|--------|
| TUI-01 | Phase 1 | In Progress (01-01 data model complete) |
| TUI-02 | Phase 1 | Complete |
| TUI-03 | Phase 1 | Complete |
| TUI-04 | Phase 1 | Pending |
| TUI-05 | Phase 2 | Pending |

**Coverage:**
- v1 requirements: 5 total
- Mapped to phases: 5
- Unmapped: 0 ✓

---
*Requirements defined: 2026-02-22*
*Last updated: 2026-02-22 after roadmap creation — traceability updated to reflect 2-phase structure*
