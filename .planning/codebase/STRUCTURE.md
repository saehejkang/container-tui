# Codebase Structure

**Analysis Date:** 2026-02-22

## Directory Layout

```
container-tui/
├── main.go                              # Program entry point
├── go.mod                               # Go module definition
├── go.sum                               # Go dependency lock file
├── pkg/
│   └── exec.go                          # Shell command execution wrapper
└── ui/
    ├── components/
    │   ├── header.go                    # Header and status indicator rendering
    │   ├── footer.go                    # Footer with instructions rendering
    │   └── styles.go                    # Reusable Lip Gloss style definitions
    └── system/
        ├── model.go                     # Root system model (menu + view switching)
        ├── view.go                      # Layout rendering (header + menu + output + footer)
        └── subcommands/
            ├── status.go                # Status query model
            ├── start.go                 # Start operation with progress model
            └── stop.go                  # Stop operation with progress model
```

## Directory Purposes

**Project Root:**
- Purpose: Program entry point and Go module configuration
- Contains: `main.go`, module files, documentation
- Key files: `main.go`

**pkg/:**
- Purpose: Shared utility packages, internal cross-cutting concerns
- Contains: Command execution wrapper
- Key files: `exec.go`

**ui/:**
- Purpose: All terminal UI code using Bubble Tea and Lip Gloss
- Contains: Layout components and state models
- Key files: None - dispatches to subdirectories

**ui/components/:**
- Purpose: Reusable UI styling and rendering helpers
- Contains: Lip Gloss style definitions, header/footer rendering functions
- Key files: `styles.go` (centralized style definitions)

**ui/system/:**
- Purpose: Root model and layout orchestration
- Contains: Main `Model` coordinating subcommand views and menu
- Key files: `model.go` (state), `view.go` (rendering)

**ui/system/subcommands/:**
- Purpose: Feature-specific models and business logic
- Contains: Each subcommand as independent `tea.Model` implementation
- Key files: `status.go`, `start.go`, `stop.go`

## Key File Locations

**Entry Points:**
- `main.go`: Creates subcommand list, instantiates `system.Model`, runs Bubble Tea program

**Configuration:**
- `go.mod`: Declares Go version and dependencies (Bubble Tea, Lip Gloss)
- `go.sum`: Lock file for exact dependency versions

**Core Logic:**
- `ui/system/model.go`: Routes messages, switches active view, manages cursor
- `ui/system/view.go`: Composes layout by calling component renderers
- `ui/system/subcommands/*.go`: Implements feature operations

**Utilities:**
- `pkg/exec.go`: Thin wrapper around `os/exec` for command execution

**Shared Patterns:**
- `ui/components/styles.go`: Centralized Lip Gloss style objects
- `ui/components/header.go`: Status indicator rendering with color coding
- `ui/components/footer.go`: Instructions and keybinding hints

## Naming Conventions

**Files:**
- Lowercase with underscores for multi-word names: `header.go`, `subcommands/status.go`
- Subcommand models named after feature: `status.go`, `start.go`, `stop.go`

**Types (Structs):**
- PascalCase suffixed with `Model` for state types: `StatusModel`, `StartModel`, `StopModel`
- PascalCase suffixed with `Msg` for message types: `StatusMsg`, `startFinishedMsg`, `tickMsg`
- PascalCase suffixed with `Style` for Lip Gloss styles: `MenuStyle`, `HeaderStyle`

**Functions:**
- `New*` constructors: `NewSystemModel`, `NewStatusModel`, `NewStartModel`
- `Render*` for UI functions: `RenderSystem`, `RenderHeaderWithStatus`, `RenderFooter`
- `run*` for command execution: `runStartCommand`, `runStopCommand`
- `tick*` for animation: `tickProgress`, `tickProgressStop`

**Variables:**
- camelCase for local variables: `menuLines`, `output`, `barWidth`
- UPPER_CASE for unexported package-level constants (none currently)
- Exported styles use PascalCase: `MenuStyle`, `CursorStyle`, `OutputBoxStyle`

**Packages:**
- Lowercase, no underscores: `main`, `pkg`, `system`, `subcommands`, `components`

## Where to Add New Code

**New Subcommand Feature:**
1. Primary code: `ui/system/subcommands/<feature>.go` - implement `tea.Model` with `Init`, `Update`, `View`
2. Add feature name string to `subcommandsList` in `main.go`
3. Add `case "<feature>"` switch branch in `system.Model.Update` (in `model.go`) to instantiate new model
4. Tests: No test framework configured - would go in `<feature>_test.go` in same directory

**New Shared UI Component:**
1. Rendering function: `ui/components/<component>.go` - create function like `RenderHeaderWithStatus`
2. Styles: Add to `ui/components/styles.go` using Lip Gloss builder API
3. Usage: Import and call from `ui/system/view.go` in layout composition

**New Shared Utility:**
1. Utility code: `pkg/<utility>.go`
2. Import path: `container-tui/pkg` as shown in `status.go`

**New Model in System Package:**
1. File: `ui/system/<model>.go` if distinct from subcommands
2. Must implement `tea.Model` interface
3. Delegate message handling to `system.Model.Update` for integration

## Special Directories

**Other Directories (not in source tree):**
- `.git/`: Version control metadata
- `.planning/codebase/`: Generated analysis documents (this file's location)

## Import Patterns

**Within Package:**
```go
import (
    "container-tui/ui/system/subcommands"  // Peer packages
    "container-tui/pkg"                     // Utility packages
    tea "github.com/charmbracelet/bubbletea" // External with alias
)
```

**Module Path:**
- All imports rooted at `container-tui/` (defined in go.mod)
- No relative imports (Go style)

---

*Structure analysis: 2026-02-22*
