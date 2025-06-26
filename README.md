# Lode - SSH Session Manager

![Demo of Lode in action](demo.gif)

A terminal user interface (TUI) for managing and connecting to SSH hosts. Built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea).

## Quick Start

**Requirements:**
- Go 1.24 or later ([Download Go](https://go.dev/dl/))

**Install the latest release:**
```sh
go install github.com/claykom/lode/cmd/lode@latest
```

**Run:**
```sh
lode
```

## Build from Source

1. **Clone the repository:**
   ```sh
   git clone https://github.com/claykom/lode
   cd lode
   ```
2. **Download dependencies:**
   ```sh
   go mod download
   ```
3. **Build the application:**
   ```sh
   go build -o lode ./cmd/lode
   # On Windows, you may want:
   go build -o lode.exe ./cmd/lode
   ```
4. **Run the application:**
   ```sh
   ./lode    # or lode.exe on Windows
   ```

## Features

- Parse and read SSH config files
- Interactive, filterable list of SSH hosts
- Quick connection to selected hosts
- Beautiful terminal UI with keyboard controls

## Controls

- `↑/↓`: Navigate through hosts
- `/`: Filter hosts
- `enter`: Connect to selected host
- `q`: Quit
- `ctrl+c`: Exit

## Development

**Install to your Go bin directory:**
```sh
go install ./cmd/lode
```

## Project Structure
- `cmd/lode/` — Main entry point
- `internal/` — Internal packages (TUI, SSH logic)

## License

MIT 