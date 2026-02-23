# Codebase Concerns

**Analysis Date:** 2026-02-22

## Error Handling Gaps

**Ignored error in status fetch:**
- Issue: `pkg.RunCommand` error is explicitly ignored with `_` in `status.go:39`
- Files: `ui/system/subcommands/status.go`
- Impact: Command execution failures silently populate the output box with potentially empty or malformed status data. User cannot distinguish between "container not responding" and "parsing error"
- Fix approach: Capture and propagate errors through the status message; display error state in UI when command fails

**Silent failure in start/stop operations:**
- Issue: `exec.Command().Run()` error is captured but only shown in progress bar result (success/failure emoji). No details about what failed
- Files: `ui/system/subcommands/start.go:77`, `ui/system/subcommands/stop.go:77`
- Impact: If `container` CLI is missing, installation fails, or permissions denied, user only sees a failure emoji with no actionable information
- Fix approach: Capture stderr separately and display error details in the View output; log error to stderr before quitting

**No error bubbling on startup:**
- Issue: `main.go:18-20` catches generic error from `p.Run()` but doesn't distinguish between terminal setup failure, model failure, or command crash
- Files: `main.go`
- Impact: Exit code is always 1 on any error; cannot distinguish fatal from transient failures
- Fix approach: Use structured errors with exit codes reflecting severity

## Untested Code

**Zero test coverage:**
- What's not tested: All core functionality - model state transitions, command execution, parsing, rendering
- Files: `ui/system/model.go`, `ui/system/subcommands/status.go`, `ui/system/subcommands/start.go`, `ui/system/subcommands/stop.go`, `pkg/exec.go`
- Risk: Changes to state update logic, refactored field parsing, or render logic can break silently. No regression detection on container CLI output format changes
- Priority: High - Status parsing (line 41-49 in status.go) is fragile and has no validation tests

## Fragile Areas

**Hardcoded status field parsing:**
- Issue: Status output parsing relies on exact field ordering and format from `container system status` command
- Files: `ui/system/subcommands/status.go:40-49`
- Impact: If `container` CLI output format changes (different spacing, field order, new fields), parsing breaks and displays empty values or crashes
- Why fragile: Uses `strings.Fields()` which splits on any whitespace; assumes index [0] is key and [1:] is value; silently skips malformed lines
- Safe modification: Add integration tests with sample `container system status` outputs; add defensive checks for field count before indexing; log parse errors for debugging

**Fixed-width field display without validation:**
- Issue: Field name key hardcoded in UpdateLoop (line 57-59) without checking if key exists in map
- Files: `ui/system/subcommands/status.go`
- Impact: If a new field is added to map initialization (line 20-29) but not added to the display loop, it silently won't show
- Safe modification: Iterate over map keys instead of hardcoding; validate all initialized keys are displayed

**Menu cursor bounds checking relies on exact comparison:**
- Issue: Cursor navigation uses exact comparison `m.Cursor < len(m.Subcommands)-1` without defensive checks
- Files: `ui/system/model.go:49`
- Impact: If subcommands list becomes empty or nil, cursor position becomes unvalidated. Accessing `m.Subcommands[m.Cursor]` on line 53 could panic
- Safe modification: Validate subcommands list length on Init; guard cursor access with bounds check

**Fake progress bar duration mismatch:**
- Issue: Start and stop operations show fake progress (increments 0.02 every 50ms = ~2.5 seconds to completion) but actual command execution time is unknown
- Files: `ui/system/subcommands/start.go:34-36`, `ui/system/subcommands/stop.go:34-36`
- Impact: If actual command takes 10 seconds but progress bar completes in 2.5, user is confused by visual disconnect. Conversely, if command finishes instantly, progress bar still shows "in progress" until ticker catches up
- Improvement path: Sync progress bar to actual command completion; show indeterminate spinner instead of fake progress; display elapsed time

## Dependency on External CLI

**Hard dependency on `container` CLI tool:**
- Risk: Application is completely non-functional if `apple/container` binary is not installed or not in PATH
- Files: `pkg/exec.go`, `ui/system/subcommands/status.go:39`, `ui/system/subcommands/start.go:77`, `ui/system/subcommands/stop.go:77`
- Impact: No graceful degradation; startup doesn't validate availability; user gets confusing "command not found" error buried in command output
- Scaling path: Add startup validation check in `main.go`; display helpful error if `container` CLI is missing; consider vendoring or providing fallback mock mode for testing

## Inconsistent Error Message Format

**No unified error handling:**
- Issue: Start/stop show only emoji + progress bar, Status shows empty values with "<not running>" fallback, main shows generic "Error running TUI"
- Files: `ui/system/subcommands/start.go:62-65`, `ui/system/subcommands/stop.go:62-65`, `ui/system/subcommands/status.go:62-64`, `main.go:19`
- Impact: Inconsistent user experience; some errors are silent (status), some show emoji, some terminate process
- Improvement path: Implement centralized error display component; standardize error message format across all models

## State Management Issues

**ActiveView model not properly initialized on startup:**
- Issue: `NewSystemModel` defaults to `StatusModel()` without calling `Init()` to fetch data on first render
- Files: `ui/system/model.go:21`
- Impact: First render shows "Fetching container system status..." but this doesn't kick off the async fetch. The fetch only starts when user manually selects Status again
- Fix approach: Call `Init()` during model construction or explicitly trigger on first Update

**Width/Height not set until WindowSizeMsg:**
- Issue: `m.Width` and `m.Height` default to 0 in Model struct
- Files: `ui/system/model.go:13-14`, `ui/system/view.go:30`
- Impact: If render happens before WindowSizeMsg (before terminal resize), `m.Width / 4` produces 0, output box has negative width
- Fix approach: Initialize with default values or guard division operations

**No command cancellation on view switch:**
- Issue: If user switches from Start to another command mid-execution, the start command continues running in background
- Files: `ui/system/model.go:52-63`
- Impact: Multiple `container system start` commands could execute simultaneously if user rapidly switches views
- Improvement path: Add context cancellation to command execution; store running command reference; cancel before launching new view

## Security Considerations

**No input validation on CLI command construction:**
- Risk: While not directly user-controlled in current code, command names are hardcoded. However, if subcommand list becomes user-configurable, could allow arbitrary command injection
- Files: `pkg/exec.go:6`, `ui/system/subcommands/status.go:39`, `ui/system/subcommands/start.go:77`, `ui/system/subcommands/stop.go:77`
- Current mitigation: Hardcoded subcommand strings prevent injection attacks currently
- Recommendations: If making subcommands dynamic, use allowlist validation and avoid shell interpolation

**No environment variable sanitization:**
- Risk: `exec.Command` inherits parent process environment. If PATH is compromised or container CLI is shadowed in working directory, could execute wrong binary
- Files: `pkg/exec.go`
- Current mitigation: None
- Recommendations: Explicitly set safe PATH; use absolute binary path if available; validate binary signature on sensitive operations

## Missing Critical Features

**No graceful shutdown handling:**
- Problem: Only quit on `q` or `ctrl+c` keypress. If underlying system is unstable, no recovery mechanism
- Blocks: Cannot handle container system crashes gracefully; no reconnect logic

**No command timeout:**
- Problem: `exec.Command().Run()` has no timeout. If container CLI hangs, TUI hangs
- Files: `ui/system/subcommands/start.go:77`, `ui/system/subcommands/stop.go:77`, `ui/system/subcommands/status.go:39`
- Impact: TUI becomes unresponsive; user must force-quit terminal
- Priority: High - Essential for reliability

**No persistence of executed commands:**
- Problem: No history of start/stop operations executed; no audit trail
- Blocks: Users cannot review what operations were performed or when

**No status refresh on demand:**
- Problem: Status is only fetched on explicit selection. No auto-refresh or refresh button
- Impact: User must repeatedly click status to see updates

## Performance Bottlenecks

**Synchronous rendering with potential long operations:**
- Problem: Each View render rechecks `ActiveView` type and status field, which could be expensive during frequent re-renders
- Files: `ui/system/view.go:11-17`
- Cause: Type assertion on every frame; could be optimized with cached type or explicit status propagation
- Improvement path: Cache the last status value; use separate flag instead of type checking

**Unbounded field parsing without size limits:**
- Problem: Status output parsing uses `strings.Split("\n")` without checking for suspiciously large responses
- Files: `ui/system/subcommands/status.go:40-50`
- Cause: No validation on output size; potential for DoS if `container` CLI returns massive output
- Improvement path: Limit max lines parsed; validate line count; add timeout to command execution

## Test Coverage Gaps

**Status model parsing logic untested:**
- What's not tested: Field extraction from command output, empty field handling, fallback to "<not running>"
- Files: `ui/system/subcommands/status.go`
- Risk: Regressions in parsing could go unnoticed for multiple releases
- Priority: High

**Model state transitions untested:**
- What's not tested: Cursor navigation bounds, view switching, message routing between parent and child models
- Files: `ui/system/model.go`
- Risk: Panic on invalid cursor position; commands could be queued incorrectly; update routing could break
- Priority: High

**Exec wrapper untested:**
- What's not tested: Error handling from Command.Run(), output capture, combined output format
- Files: `pkg/exec.go`
- Risk: Silent failures if Go's exec package behavior changes; no validation of captured output
- Priority: Medium

---

*Concerns audit: 2026-02-22*
