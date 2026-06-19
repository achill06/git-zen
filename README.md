# git-zen

An open-source Go CLI extension suite offering a blazing-fast, interactive terminal user interface (TUI) for git branch management. It replaces standard, unstyled git outputs with an asynchronous dashboard that integrates with the GitHub API to show live PR statuses.

## Features (v1.0)

* **Fuzzy Search:** Instantly filter through massive lists of local branches as you type.
* **Async GitHub Integration:** Fetches live Pull Request statuses in the background without freezing the UI thread.
* **Clean Unicode Aesthetics:** Uses minimalist colored status indicators (`● Open`, `● Merged`, `● Closed`) instead of heavy text blocks or emojis to ensure a native terminal look.
* **Graceful Degradation:** Functions flawlessly as a local fuzzy branch switcher even if you are offline or unauthenticated with GitHub.
* **Native Extensibility:** Extends the native `git` engine—once installed, it integrates smoothly into your daily workflow as a native command.

---

## Installation

### 1. Download Binaries (Recommended)

Head over to the [GitHub Releases](https://github.com/achill06/git-zen/releases) page to download the pre-compiled binary for your operating system (macOS, Linux, or Windows). Move the executable into a folder in your `PATH` (like `/usr/local/bin`).

### 2. Alternative: Install via Go

If you have Go installed, you can build from source:

```bash
go install github.com/achill06/git-zen@latest
```

> **Note:** Ensure `~/go/bin` is in your system `PATH`. Add `export PATH=$PATH:~/go/bin` to your `~/.bashrc` or `~/.zshrc`, then run `source ~/.bashrc`.

---

## Usage

Navigate to any local git repository on your computer and execute:

```bash
git zen switch
```

### Keyboard Controls

| Key | Action |
|-----|--------|
| Arrow Keys / `Ctrl+J` / `Ctrl+K` | Navigate up and down through the filtered branch list |
| Alphanumeric Keys | Start typing at any time to instantly fuzzy-search and filter your branches |
| `Enter` | Confirm your selection and immediately perform a `git checkout` onto the selected branch |
| `Escape` / `Ctrl+C` | Safely exit the interface without changing your current branch status |

---

## Tech Stack

| Component | Library |
|-----------|---------|
| Language | Go (Golang) |
| CLI Engine | [Cobra](https://github.com/spf13/cobra) |
| TUI Architecture | [Bubble Tea](https://github.com/charmbracelet/bubbletea) |
| Styling & Layout | [Lipgloss](https://github.com/charmbracelet/lipgloss) |
| Fuzzy Filtering | [Fuzzy](https://github.com/sahilm/fuzzy) |

---

## Roadmap

* **v1.1: AI Commit Summarizer** (`git zen log`) — Automatically convert cryptic commit histories into polished markdown changelogs using local LLMs.
* **v2.0: Background Undo Daemon** — A safety net mechanism allowing developers to instantly reverse accidental git operations.