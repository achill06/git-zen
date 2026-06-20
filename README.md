# git-zen
![demo](demo.gif)

> A blazing-fast fuzzy git branch switcher with live GitHub PR status in your terminal.

## Prerequisites
- Go 1.24+
- `gh` CLI installed and authenticated (`gh auth login`)

## Features (v1.0)

* **Fuzzy Search:** Instantly filter through massive lists of local branches as you type.
* **Async GitHub Integration:** Fetches live Pull Request statuses in the background without freezing the UI thread.
* **Clean Unicode Aesthetics:** Uses minimalist colored status indicators (`● Open`, `● Merged`, `● Closed`) instead of heavy text blocks or emojis to ensure a native terminal look.
* **Shell Autocompletion**: bash, zsh, fish, and PowerShell supported via git-zen completion <shell>
* **Graceful Degradation:** Functions flawlessly as a local fuzzy branch switcher even if you are offline or unauthenticated with GitHub.

---

## Installation

### macOS & Linux (Recommended)

Install via Homebrew. This method automatically configures shell autocompletions for you.

```bash
brew install achill06/tap/git-zen
```

### Windows (Recommended)

Install via Scoop:

```powershell
scoop bucket add git-zen https://github.com/achill06/scoop-bucket.git
scoop install git-zen
```

### Alternative Methods

**Manual Binary Download:** Head over to the [GitHub Releases](https://github.com/achill06/git-zen/releases) page to download the pre-compiled binary for your OS and move the executable into a folder in your `PATH`. Run `git-zen completion <your-shell>` to set up autocompletion manually.

**Install via Go:**

```bash
go install github.com/achill06/git-zen@latest
```

> **Note:** Ensure `~/go/bin` is in your system `PATH`. Add `export PATH=$PATH:~/go/bin` to your `~/.bashrc` or `~/.zshrc`, then run `source ~/.bashrc`.

---

## Usage

Navigate to any local git repository and run:

```bash
git zen switch
```

> **Tip:** Try typing `git-zen sw` and pressing `<TAB>` for autocompletion.

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