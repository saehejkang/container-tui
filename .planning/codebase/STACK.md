# Technology Stack

**Analysis Date:** 2026-02-22

## Languages

**Primary:**
- Go 1.25 - Entire codebase (main application)

## Runtime

**Environment:**
- Go 1.25 - Native binary compilation and execution

**Package Manager:**
- Go Modules (go.mod/go.sum) - Dependency management
- Lockfile: Present (go.sum with resolved versions)

## Frameworks

**Core:**
- Bubble Tea v1.3.10 - Terminal UI framework (Model-Update-View pattern)
- Lip Gloss v1.1.0 - Terminal styling and layout rendering

**Utilities (transitive dependencies):**
- charmbracelet/x/ansi v0.11.6 - ANSI escape sequence handling
- charmbracelet/colorprofile v0.4.2 - Color profile detection
- charmbracelet/x/cellbuf v0.0.15 - Terminal cell buffer
- charmbracelet/x/term v0.2.2 - Terminal utilities
- muesli/termenv v0.16.0 - Terminal environment detection

## Key Dependencies

**Critical:**
- github.com/charmbracelet/bubbletea v1.3.10 - TUI event loop and message passing model
- github.com/charmbracelet/lipgloss v1.1.0 - Layout composition and styling (header, footer, menu, output box)

**Text Rendering & Display:**
- clipperhouse/displaywidth v0.11.0 - Unicode character width calculation
- clipperhouse/stringish v0.1.1 - String manipulation utilities
- mattn/go-runewidth v0.0.20 - Rune width for terminal display
- rivo/uniseg v0.4.7 - Unicode segmentation
- lucasb-eyer/go-colorful v1.3.0 - Color manipulation
- xo/terminfo v0.0.0-20220910002029-abceb7e1c41e - Terminal info parsing

**System Integration:**
- golang.org/x/sys v0.41.0 - Low-level system calls for terminal control
- mattn/go-isatty v0.0.20 - TTY detection
- erikgeiser/coninput v0.0.0-20211004153227-1c3628e74d0f - Console input handling
- muesli/cancelreader v0.2.2 - Cancellable input reading

**Platform Support:**
- aymanbagabas/go-osc52/v2 v2.0.1 - OSC 52 escape sequence support (clipboard)
- golang.org/x/text v0.34.0 - Unicode and internationalization support

## Configuration

**Environment:**
- No `.env` or configuration files required
- Runtime configuration passed via command-line
- Container CLI binary must be in system PATH

**Build:**
- `go.mod` - Module declaration (module: container-tui, Go version: 1.25)
- `go.sum` - Locked dependency versions for reproducible builds
- Standard Go build toolchain

## Platform Requirements

**Development:**
- Go 1.25+ toolchain
- `apple/container` CLI installed and available in PATH
- Unix-like environment (macOS, Linux) - relies on terminal control sequences

**Production:**
- Deployment target: macOS (Apple Silicon or Intel), Linux, or any system with Go support
- `apple/container` CLI system dependency (required at runtime)
- Terminal with ANSI escape sequence support

---

*Stack analysis: 2026-02-22*
