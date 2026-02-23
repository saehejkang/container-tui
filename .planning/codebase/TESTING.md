# Testing Patterns

**Analysis Date:** 2026-02-22

## Test Framework

**Runner:**
- Go built-in testing (no external test framework configured)
- Standard `go test` command

**Assertion Library:**
- None - Go testing uses simple comparisons and conditionals

**Run Commands:**
```bash
go test ./...         # Run all tests
go test -v ./...      # Verbose output
go test -cover ./...  # With coverage
go test -race ./...   # Race detection
```

**Current Status:**
- No test files found in codebase (`*_test.go` files: 0)
- No testing configuration files (no `go.test.yml`, `testing.toml`, etc.)

## Test File Organization

**Location:**
- Not implemented - no test files currently exist

**Recommended Pattern:**
- Co-located tests: `<filename>_test.go` in same package as source file
- Example structure (not currently present):
  - `ui/system/model.go` → `ui/system/model_test.go`
  - `pkg/exec.go` → `pkg/exec_test.go`
  - `ui/components/header.go` → `ui/components/header_test.go`

**Naming:**
- Test function prefix: `Test<FunctionName>`
- Example: `func TestNewSystemModel(t *testing.T)`, `func TestRenderSystem(t *testing.T)`
- Benchmark prefix: `Benchmark<FunctionName>`
- Sub-tests: `t.Run("case name", func(t *testing.T) { ... })`

## Test Structure

**Suite Organization:**
- Go testing uses table-driven tests as common pattern:

```go
func TestRenderSystem(t *testing.T) {
    tests := []struct {
        name    string
        input   *system.Model
        wantErr bool
    }{
        {
            name: "renders menu with cursor",
            input: &system.Model{
                Cursor: 0,
                Subcommands: []string{"start", "stop", "status"},
            },
            wantErr: false,
        },
        // additional cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // arrange
            // act
            result := RenderSystem(tt.input)
            // assert
            if result == "" {
                t.Errorf("expected non-empty output")
            }
        })
    }
}
```

**Patterns:**
- Arrange-Act-Assert: Set up test data, execute function, verify results
- Parallel testing: `t.Parallel()` for independent tests
- Helper functions: `func helper(t *testing.T) type` pattern for test utilities

## Mocking

**Framework:**
- No mocking framework installed or configured (no testify/mock, github.com/golang/mock, etc.)

**Recommended Pattern for pkg.RunCommand:**
- Interface extraction for `os/exec` package if testing needed:

```go
// in pkg/exec.go
type CommandRunner interface {
    Run(name string, args ...string) (string, error)
}

// in tests
type MockCommandRunner struct {
    RunFunc func(string, ...string) (string, error)
}

func (m *MockCommandRunner) Run(name string, args ...string) (string, error) {
    return m.RunFunc(name, args...)
}
```

**What to Mock:**
- External command execution (`container` CLI calls)
- File I/O if added
- Network calls if integrations added

**What NOT to Mock:**
- Bubble Tea framework behavior - use actual tea.Model interface
- Lipgloss rendering - test actual output or use snapshot comparison
- Simple data transformations in StatusModel.Update()

## Fixtures and Factories

**Test Data:**
- Not currently used - no test fixtures or factories

**Recommended Pattern:**
```go
func testSystemModel() *system.Model {
    return &system.Model{
        Cursor: 0,
        Subcommands: []string{"status", "start", "stop"},
        ActiveView: subcommands.NewStatusModel(),
        Width: 120,
        Height: 30,
    }
}

func testStatusModel() *subcommands.StatusModel {
    return &subcommands.StatusModel{
        Fields: map[string]string{
            "status": "running",
            "appRoot": "/app",
        },
        Output: []string{},
    }
}
```

**Location:**
- `pkg/testutil/` or `testhelpers.go` files in test packages
- Or: `_test.go` file with setup functions in each package

## Coverage

**Requirements:**
- Not enforced - no coverage targets set
- No code coverage tooling configured

**View Coverage:**
```bash
go test -cover ./...
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```

**Current Status:**
- 0% coverage (no tests written)
- Critical paths without tests:
  - `ui/system/model.go`: Model.Update() - cursor movement, view switching
  - `ui/system/subcommands/status.go`: StatusModel.Init(), Update() - command output parsing
  - `ui/system/view.go`: RenderSystem() - layout composition
  - `pkg/exec.go`: RunCommand() - external command execution

## Test Types

**Unit Tests:**
- Scope: Individual functions and methods
- Approach: Input/output verification for pure functions, state mutation verification for model methods
- Examples to write:
  - `pkg.RunCommand()` with mock exec
  - `StatusModel.Update()` with StatusMsg input
  - Rendering functions: `RenderSystem()`, `RenderHeaderWithStatus()` output format
  - Cursor movement in `system.Model.Update()` with key messages

**Integration Tests:**
- Scope: Multiple models/packages working together
- Approach: Minimal - could test full flow through system.Model → subcommand → external cmd
- Considerations: Would require either mocking CLI or running against live container CLI
- Placement: `ui/system/integration_test.go` or similar

**E2E Tests:**
- Framework: Not applicable - Bubble Tea TUI is interactive
- Alternative: Manual testing or snapshot/recording-based tests with bubbletea/testing package features

## Common Patterns

**Async Testing:**
- Bubble Tea models use `tea.Cmd` for async work
- Testing pattern: send command result as message and verify model state changes

```go
func TestStatusModelAsync(t *testing.T) {
    m := subcommands.NewStatusModel()

    // Simulate Init() returning a cmd that sends StatusMsg
    _, cmd := m.Update(nil)

    // The cmd is a function that returns tea.Msg
    // Execute it and capture the message
    msg := cmd()

    // Now send that message back through Update
    m2, _ := m.Update(msg)

    // Verify state changed appropriately
    if m2.(*subcommands.StatusModel).Output == nil {
        t.Errorf("expected Output to be populated")
    }
}
```

**Error Testing:**
```go
func TestRunCommandFailure(t *testing.T) {
    // Assuming RunCommand is refactored to be testable
    output, err := pkg.RunCommand("nonexistent", "command")

    if err == nil {
        t.Errorf("expected error for nonexistent command")
    }

    if output != "" {
        t.Errorf("expected empty output on error")
    }
}

func TestStartModelErrorDisplay(t *testing.T) {
    m := subcommands.NewStartModel()

    // Send error message
    msg := startFinishedMsg{err: errors.New("container failed to start")}
    m2, _ := m.Update(msg)

    view := m2.(*subcommands.StartModel).View()
    if !strings.Contains(view, "❌") {
        t.Errorf("expected error indicator in view")
    }
}
```

## Golden/Snapshot Testing

**Recommended for this codebase:**
- Use `go-cmp` or similar for comparing rendered output
- Store baseline rendered strings as test fixtures

```go
func TestRenderSystemLayout(t *testing.T) {
    m := testSystemModel()
    output := system.RenderSystem(m)

    want := `... multiline expected output ...`
    if output != want {
        t.Errorf("output mismatch:\ngot:\n%s\nwant:\n%s", output, want)
    }
}
```

---

*Testing analysis: 2026-02-22*
