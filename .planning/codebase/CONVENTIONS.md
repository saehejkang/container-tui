# Coding Conventions

**Analysis Date:** 2026-02-22

## Naming Patterns

**Files:**
- Package-scoped files: lowercase with underscores (e.g., `exec.go`)
- Model files: `model.go` or `<modelname>.go` (e.g., `status.go`, `start.go`)
- View/rendering files: `view.go` or descriptor names (e.g., `header.go`, `footer.go`, `styles.go`)

**Functions:**
- Exported: PascalCase (e.g., `NewStatusModel()`, `RunCommand()`, `RenderSystem()`)
- Private: camelCase (e.g., `runStartCommand()`, `tickProgress()`)
- Constructor pattern: `New<TypeName>()` (e.g., `NewStatusModel()`, `NewStartModel()`)
- Renderer functions: `Render<Component>()` (e.g., `RenderSystem()`, `RenderHeaderWithStatus()`, `RenderFooter()`)

**Variables:**
- Struct fields: PascalCase for exported (e.g., `Model`, `Fields`, `Output`), camelCase for private (e.g., `progress`, `done`, `err`)
- Local variables: camelCase (e.g., `menuLines`, `outputWidth`, `statusModel`)
- Constants: PascalCase (e.g., `StatusModel`, `MenuStyle`, `CursorStyle`)
- Message types (custom Msg structs): PascalCase suffixed with `Msg` (e.g., `StatusMsg`, `startFinishedMsg`, `tickMsg`)

**Types:**
- Structs: PascalCase (e.g., `StatusModel`, `StartModel`, `StopModel`)
- Private message types: camelCase with `Msg` suffix (e.g., `startFinishedMsg`, `stopFinishedMsg`, `tickMsg`, `tickStopMsg`)
- Exported types follow export rule: `StatusMsg` for exported messages

## Code Style

**Formatting:**
- Go standard fmt (gofmt) - code follows Go idioms
- Tabs for indentation (Go standard)
- Single space between logical sections
- No visible linting configuration file - uses Go conventions

**Linting:**
- No linter config detected (relies on Go idioms)
- Follows Go best practices: error handling, package organization, naming

**Import Organization:**
- Standard library imports first (e.g., `fmt`, `strings`, `os/exec`, `time`)
- Blank line
- External/third-party imports (e.g., `tea`, `lipgloss`)
- Blank line
- Local package imports (e.g., `container-tui/pkg`, `container-tui/ui/...`)

See examples:
- `ui/system/model.go`: Standard lib first (none), then Bubble Tea, then local imports
- `ui/system/subcommands/status.go`: `fmt`, `strings` → `tea` → `container-tui/pkg`
- `ui/system/view.go`: `container-tui/ui/components`, `container-tui/ui/system/subcommands` → `lipgloss`

## Error Handling

**Patterns:**
- Return tuple with error as second value: `func (name string, args ...string) (string, error)`
- Errors are often ignored in background commands (via `tea.Cmd` functions) - errors are stored in model state and displayed in UI
- Errors from command execution checked and displayed: `if m.err != nil { return "❌ ... Failed\n" + bar }`
- Main entry point: explicit error check with stderr output and os.Exit(1)

Example from `main.go`:
```go
if _, err := p.Run(); err != nil {
    fmt.Fprintf(os.Stderr, "Error running TUI: %v\n", err)
    os.Exit(1)
}
```

Example from `ui/system/subcommands/status.go`:
```go
out, _ := pkg.RunCommand("container", "system", "status")
```

Pattern: Blanking error (underscore) when error is non-critical or handled by upstream caller.

## Logging

**Framework:** `fmt` package only - no logging framework

**Patterns:**
- Stderr for errors (via `fmt.Fprintf(os.Stderr, ...)`)
- No logging within model updates or renders - status displayed via UI text rendering
- Output from commands captured and rendered as strings in views

## Comments

**When to Comment:**
- No comments observed in the codebase - code is self-documenting through clear naming and structure
- Bubble Tea pattern is implicit in model structure (Init, Update, View methods)

**JSDoc/TSDoc:**
- Go uses package-level comments above exported types/functions in production codebases, but this codebase has none
- Consider adding package comments and exported function comments for public APIs if library usage grows

## Function Design

**Size:**
- Functions are small and focused (8-80 lines, most under 50)
- Largest: `RenderSystem()` (46 lines) - justified by layout composition

**Parameters:**
- Constructor-style: no parameters for `NewStatusModel()`, `NewStartModel()`, etc. - state initialized empty
- Receiver methods use pointer receivers for `tea.Model` interface implementation: `(m *StatusModel)`
- Simple parameters: `func RunCommand(name string, args ...string)` uses variadic for flexibility
- Styling functions accept simple types: `func RenderHeaderWithStatus(title string, isRunning bool)`

**Return Values:**
- Message-passing pattern: tea.Cmd returns tea.Msg
- Constructor returns pointer to model: `*StatusModel`
- Helper functions return formatted strings: `string`
- Command functions return `tea.Cmd` wrapping the work: `return func() tea.Msg { ... }`

## Module Design

**Exports:**
- Clear public API per package:
  - `main` imports `system.NewSystemModel`
  - `ui/system` exports `Model`, `NewSystemModel()`, and `RenderSystem()`
  - `ui/system/subcommands` exports model types and constructors
  - `ui/components` exports `Style` variables and `Render*` functions
  - `pkg` exports `RunCommand()`

**Barrel Files:**
- No barrel files (index files) used
- Direct imports from specific modules

## Interfaces

**Bubble Tea Interface Implementation:**
- All models implement `tea.Model` interface: `Init() tea.Cmd`, `Update(tea.Msg) (tea.Model, tea.Cmd)`, `View() string`
- Follows Bubble Tea architecture exactly as documented in CLAUDE.md

**Example from `ui/system/subcommands/status.go`:**
```go
type StatusModel struct {
    Fields map[string]string
    Output []string
}

func (m *StatusModel) Init() tea.Cmd { ... }
func (m *StatusModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { ... }
func (m *StatusModel) View() string { ... }
```

---

*Convention analysis: 2026-02-22*
