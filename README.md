# MCMD - Command Line Utility

MCMD is a powerful command-line utility that enhances your terminal experience with quick navigation, environment management, and system commands.

## Installation

1. Clone the repository
2. Run `go build -o mcmd`
3. Add the binary to your PATH
4. Add the following line to your `.bashrc` or `.zshrc`:
   ```bash
   eval "$(mcmd goto --init-shell)"
   ```
5. Reload your shell configuration:
   ```bash
   source ~/.bashrc  # or source ~/.zshrc
   ```

## Features

### 1. Quick Directory Navigation
Navigate to your frequently used directories with shortcuts:
```bash
goto work      # Navigate to /data/work
goto go        # Navigate to /data/work/go_work
goto python    # Navigate to /data/work/python_work
goto conda     # Navigate to /data/work/python_work/pytorch_work
goto llm       # Navigate to /data/work/llm_work
goto ebpf      # Navigate to /data/work/githubwork/ebpf
goto mcmd      # Navigate to /data/work/go_work/ws/mcmd
```

### 2. Quick Commands
Execute common tasks with simple commands:
```bash
mcmd-goto -r pytorch_env  # Activate pytorch environment
mcmd-goto -r base        # Activate base environment
mcmd-goto -r edit        # Edit config file
mcmd-goto -r zsh         # Reload zsh configuration
```

### 3. System Commands
Quick access to system information:
```bash
mcmd-goto sys   # Show system info (uname -a)
mcmd-goto mem   # Show memory info (free -h)
mcmd-goto net   # Show network ports (netstat -tuln)
```

### 4. Custom Commands
Run any custom shell command:
```bash
mcmd-goto 'your custom command'
```

## Configuration

All shortcuts and commands are configured in `cmd/config.go`. You can customize:
- Directory shortcuts in `DirectoryMapping`
- Quick commands in `OtherCommandMapping`
- System commands in `SysCommandMapping`

## Usage

View all available commands and shortcuts:
```bash
mcmd goto
```

For detailed help on any command:
```bash
mcmd goto --help
```

## Shell Integration

The shell integration enables the following aliases:
- `goto` for directory navigation
- `mcmd-goto` for all features
- `activate` for environment activation

Make sure to add the initialization command to your shell configuration:
```bash
eval "$(mcmd goto --init-shell)"
```
