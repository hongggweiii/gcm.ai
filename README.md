# gcm.ai

**gcm.ai** is a CLI tool designed to write high-quality commit messages. It analyzes your staged changes and suggests context-aware, professional commit messages in seconds.

## Features

- **AI-Powered Analysis**: Uses AI to understand the *intent* behind your code changes.
- **Conventional Commits**: Supports the `type(scope): description` standard out of the box.
- **Interactive UI**: Built with a polished terminal interface using `huh`.
- **Customizable**: Toggle between single-line and multi-line (subject/body) formats.
- **Stateless & Fast**: Minimal overhead; just run it and commit.

---

## Installation

### Prerequisites
- **Go 1.21+** installed on your system.
- A **Gemini API Key** (Get one for free at [Google AI Studio](https://aistudio.google.com/)).
- **Git** initialized in your project.

### Global Installation
To use `gcm.ai` from any directory on your machine, install it via the Go toolchain:

**1. Clone the repository**
```bash
git clone [https://github.com/hongggweiii/gcm.ai.git](https://github.com/hongggweiii/gcm.ai.git)
cd gcm.ai
```

**2. Configure API keys**
Create an `.env` file in the root of the `gcm.ai` folder and add your key:
```bash
GEMINI_API_KEY=...
```

**3. Setting up `.env` file**
Add an .env file to the root of gcm.ai
```bash
GEMINI_API_KEY=your_api_key_here
```

**3. Install the binary**
Run the following command inside the cloned repository to build and install the tool to your `$GOPATH/bin`
```bash
go install .
```

**4. Update your system PATH**
Ensure Go's bin folder is in your system's PATH. Add this line to your ~/.zshrc (or ~/.bashrc):
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```
Run this line below to restart your terminal
```bash
source ~/.zshrc
```

**5. Run the tool**
You can now run the project from any Git repository on your machine by simply typing
```bash
gcm.ai
```

## Roadmap (Future Implementations)
Upcoming features include:

- **Multi-Provider Support:** Expanding beyond Gemini, new models from:
  - OpenAI
  - DeepSeek
  - Anthropic (Claude)
  - Local LLMs (via Ollama)
- **Granular Model Selection:** Allowing users to specify exactly which model version they want to use within their chosen provider (e.g., toggling between `gpt-4o`, `gpt-5`, or `gemini-3-pro` via a configuration file).
- **Global Config File:** A `~/.gcm-config.json` setup to permanently save your preferred AI provider, model choice, and commit style so you don't have to select it every time.