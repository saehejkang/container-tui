# Architecture

**Analysis Date:** 2026-02-22

## Pattern Overview

**Overall:** Model-Update-View (MVU) via Bubble Tea framework

**Key Characteristics:**
- Hierarchical model composition with a root system model delegating to view-specific submodels
- Event-driven command system for async operations (CLI execution, progress ticking)
- Message-based communication between models
- Declarative rendering using Lip Gloss layout composition

## Layers

**Presentation Layer (UI):**
- Purpose: Render the terminal UI layout and handle terminal resize events
- Location: `ui/system/view.go`, `ui/components/`
- Contains: Layout rendering functions, reusable styles, header/footer/menu composition
- Depends on: Component styles, Lip Gloss library
- Used by: `system.Model.View()` to produce terminal output

**State Management Layer:**
- Purpose: Manage application state, route messages, and coordinate model transitions
- Location: `ui/system/model.go`
- Contains: `system.Model` with cursor position, active view state, terminal dimensions
- Depends on: Bubble Tea interface (`tea.Model`), subcommand models
- Used by: `main.go` as root model, delegates Update/View to active subcommand

**Subcommand Layer:**
- Purpose: Implement feature-specific logic and handle CLI operations
- Location: `ui/system/subcommands/`
- Contains: `StatusModel`, `StartModel`, `StopModel` - each implementing `tea.Model`
- Depends on: Bubble Tea, `pkg.RunCommand` (status only), `os/exec` (start/stop)
- Used by: `system.Model` switching to appropriate model on Enter keypress

**Component Layer:**
- Purpose: Provide reusable UI patterns and styling
- Location: `ui/components/`
- Contains: `RenderHeaderWithStatus`, `RenderFooter`, predefined Lip Gloss styles
- Depends on: Lip Gloss library
- Used by: `view.go` for layout composition

**Command Execution Layer:**
- Purpose: Wrap low-level system command execution
- Location: `pkg/exec.go`
- Contains: `RunCommand` function for combined stdout+stderr capture
- Depends on: `os/exec` standard library
- Used by: `StatusModel.Init()` for parsing container system status

## Data Flow

**Status Query Flow:**

1. User selects "status" from menu and presses Enter
2. `system.Model.Update` instantiates `StatusModel` and calls `Init()`
3. `StatusModel.Init()` returns `tea.Cmd` that calls `pkg.RunCommand("container", "system", "status")`
4. Result is parsed into `Fields` map, wrapped as `StatusMsg`, returned to `Update`
5. `StatusModel.Update` formats fields into `Output` string slice
6. `StatusModel.View()` joins output lines and returns rendered string
7. `RenderSystem` displays output in right-side box with menu on left

**Start/Stop Operation Flow:**

1. User selects "start" or "stop" and presses Enter
2. `system.Model.Update` instantiates `StartModel` or `StopModel` and calls `Init()`
3. `Init()` returns `tea.Batch` of two commands:
   - `runStartCommand()`: executes `container system start` in background
   - `tickProgress()`: emits `tickMsg` every 50ms for animation
4. Each `tickMsg` increments progress bar until command completes
5. Command result wrapped as `startFinishedMsg` or `stopFinishedMsg` with error status
6. `Update` sets final progress (1.0) and error state
7. `View()` renders progress bar and success/failure message

**State Management:**
- Root model (`system.Model`) owns cursor position and subcommand list
- On Enter, root model switches `ActiveView` to new subcommand model
- All non-keyboard messages delegate to `ActiveView.Update()`
- Terminal resize updates root model's `Width`/`Height`, then delegates to active view

## Key Abstractions

**tea.Model Interface:**
- Purpose: Standard Bubble Tea contract for all stateful components
- Examples: `system.Model`, `StatusModel`, `StartModel`, `StopModel`
- Pattern: Each implements `Init() tea.Cmd`, `Update(msg tea.Msg) (tea.Model, tea.Cmd)`, `View() string`

**tea.Cmd:**
- Purpose: Represent side effects (async operations, IO) as values
- Examples: Status query, start/stop execution, progress ticking
- Pattern: Function that returns `tea.Msg` result; Bubble Tea schedules execution

**tea.Msg:**
- Purpose: Type-safe event messages passed from commands back to models
- Examples: `StatusMsg` (map[string]string), `startFinishedMsg`, `tickMsg`
- Pattern: Empty structs for tick events, typed maps for data payloads

**Lip Gloss Styles:**
- Purpose: Encapsulate terminal styling (colors, padding, borders) for reuse
- Examples: `MenuStyle`, `HeaderStyle`, `OutputBoxStyle` in `ui/components/styles.go`
- Pattern: Predefined variables using Lip Gloss builder API

## Entry Points

**main.go:**
- Location: `/Users/kevingeorge/Documents/projects/container-tui/container-tui/main.go`
- Triggers: Program start
- Responsibilities: Define subcommand list, create `system.Model`, run Bubble Tea program

**system.Model.Update:**
- Location: `ui/system/model.go`
- Triggers: Every message (keypresses, commands, ticks)
- Responsibilities: Handle navigation (up/down arrows), quit (q/ctrl+c), view switching (Enter)

**Subcommand Init:**
- Location: `ui/system/subcommands/*.go`
- Triggers: When model instantiated and `Init()` called
- Responsibilities: Launch background CLI command, schedule ticks for animation

## Error Handling

**Strategy:** Capture stderr in command execution; store error state in model; display in view

**Patterns:**
- `StatusModel`: Ignores errors from `pkg.RunCommand`, shows `<not running>` for missing fields
- `StartModel`/`StopModel`: Capture error from `exec.Command.Run()`, store in model, render failure message with check mark or X emoji
- No error recovery; user must re-enter command via menu navigation

## Cross-Cutting Concerns

**Logging:** Not implemented - all output via terminal UI rendering

**Validation:** CLI commands assume `container` binary available in PATH; no pre-validation

**Authentication:** None - delegates to `container` CLI which handles auth

---

*Architecture analysis: 2026-02-22*
