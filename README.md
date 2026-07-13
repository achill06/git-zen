# git-zen

Fuzzy git branch switcher with live GitHub PR status in your terminal.

![platform](<https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-blue>) ![go](https://img.shields.io/badge/Go-1.24%2B-00ADD8?logo=go&logoColor=white) ![license](https://img.shields.io/badge/license-MIT-green)

![demo](demo.gif)

## Features

* **Fuzzy Search:** Instantly filter through massive lists of local branches as you type.
* **Async GitHub Integration:** Fetches live Pull Request statuses in the background without freezing the UI thread.
* **Clean Unicode Aesthetics:** Uses minimalist colored status indicators (`● Open`, `● Merged`, `● Closed`) instead of heavy text blocks or emojis to ensure a native terminal look.
* **Shell Autocompletion:** bash, zsh, fish, and PowerShell supported via `git-zen completion <shell>`
* **Graceful Degradation:** Functions flawlessly as a local fuzzy branch switcher even if you are offline or unauthenticated with GitHub.

---

## Requirements

* Go 1.24+
* `gh` CLI installed and authenticated (`gh auth login`)
* macOS, Linux, or Windows

---

## Installation

### macOS & Linux (Recommended)

Install via Homebrew. Shell autocompletions are configured automatically — no extra steps needed.

```bash
brew install achill06/tap/git-zen
```

### Windows (Recommended)

Install via Scoop, then add completions to your PowerShell profile:

```powershell
scoop bucket add git-zen https://github.com/achill06/scoop-bucket.git
scoop install git-zen

# Add completions to your PowerShell profile:
echo 'git-zen completion powershell | Out-String | Invoke-Expression' >> $PROFILE
```

### Install via Go

```bash
go install github.com/achill06/git-zen@latest
```

> **Note:** Ensure `~/go/bin` is in your system `PATH`. Add `export PATH=$PATH:~/go/bin` to your `~/.bashrc` or `~/.zshrc`, then run `source ~/.bashrc`.

Then add completions for your shell:

```bash
# bash
echo 'source <(git-zen completion bash)' >> ~/.bashrc
source ~/.bashrc

# zsh
echo 'source <(git-zen completion zsh)' >> ~/.zshrc
source ~/.zshrc

# fish
git-zen completion fish > ~/.config/fish/completions/git-zen.fish
```

### Manual Binary Download

Head over to the [GitHub Releases](https://github.com/achill06/git-zen/releases) page, download the pre-compiled binary for your OS, and move it into a folder in your `PATH`.

Then add completions for your shell:

```bash
# bash
echo 'source <(git-zen completion bash)' >> ~/.bashrc
source ~/.bashrc

# zsh
echo 'source <(git-zen completion zsh)' >> ~/.zshrc
source ~/.zshrc

# fish
git-zen completion fish > ~/.config/fish/completions/git-zen.fish

# PowerShell
echo 'git-zen completion powershell | Out-String | Invoke-Expression' >> $PROFILE
```

---

## Upgrading

**Homebrew:**

```bash
brew upgrade git-zen
```

**Scoop:**

```powershell
scoop update git-zen
```

**Go:**

```bash
go install github.com/achill06/git-zen@latest
```

---

## Usage

### CLI Commands

| Command                        | Description                                                     |
| ------------------------------ | --------------------------------------------------------------- |
| `git-zen switch`             | Launch the interactive fuzzy branch switcher (TUI)              |
| `git-zen completion <shell>` | Generate shell autocompletion script (bash/zsh/fish/PowerShell) |
| `git-zen version`            | Show version information                                        |

Navigate to any local git repository and run:

```bash
git-zen switch
```

> **Tip:** With completion set up, type `git-zen sw<TAB>` and it will autocomplete to `git-zen switch`.

### Keyboard Controls

| Key                                 | Action                                                                                    |
| ----------------------------------- | ----------------------------------------------------------------------------------------- |
| Arrow Keys /`Ctrl+J` / `Ctrl+K` | Navigate up and down through the filtered branch list                                     |
| Alphanumeric Keys                   | Start typing at any time to instantly fuzzy-search and filter your branches               |
| `Enter`                           | Confirm your selection and immediately perform a`git checkout` onto the selected branch |
| `Escape` / `Ctrl+C`             | Safely exit the interface without changing your current branch status                     |

---

## Tech Stack

| Component        | Library                                                 |
| ---------------- | ------------------------------------------------------- |
| Language         | Go (Golang)                                             |
| CLI Engine       | [Cobra](https://github.com/spf13/cobra)                  |
| TUI Architecture | [Bubble Tea](https://github.com/charmbracelet/bubbletea) |
| Styling & Layout | [Lipgloss](https://github.com/charmbracelet/lipgloss)    |
| Fuzzy Filtering  | [Fuzzy](https://github.com/sahilm/fuzzy)                 |

---

## Roadmap

* **v1.1: AI Commit Summarizer** (`git-zen log`) — Automatically convert cryptic commit histories into polished markdown changelogs using local LLMs.
* **v2.0: Background Undo Daemon** — A safety net mechanism allowing developers to instantly reverse accidental git operations.

---

## License

MIT
