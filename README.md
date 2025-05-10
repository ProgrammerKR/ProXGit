# ProXGit

**ProXGit** is a lightweight, fast, and self-hosted Git service inspired by Gitea. It aims to provide a simple, easy-to-deploy alternative for managing Git repositories locally or on private servers.

---

## Features

- Fast and minimal self-hosted Git service  
- Cross-platform support (Linux, Android via Termux, macOS, Windows)  
- Simple installation and setup  
- Web UI for repositories, issues, pull requests, and user management  
- RESTful API support  
- Written in Go (single binary)  

---

## Installation

1. **Download the binary** (or build from source)  
2. **Run the server**:
   ```bash
   ./proxgit web
   ```
3. **Visit in browser**:  
   `http://localhost:3000` (default)

---

## Build From Source

To build manually:

```bash
make build
```

Requires:
- Go (version from `go.mod`)  
- Node.js (optional, for frontend)  

---

## Folder Structure

```
/ProXGit
  ├── custom/
  ├── data/
  ├── log/
  ├── proxgit (binary)
  └── app.ini (config)
```

---

## Configuration

Edit the `app.ini` file inside the `custom/conf/` folder to adjust your:

- Database  
- Repository path  
- Server port  
- Log settings  

---

## Usage

Start the server:

```bash
./proxgit web
```

Use `localhost:3000` or your configured domain to access the UI.

---

## Contribution

Fork → Patch → Pull Request  
Please read the `CONTRIBUTING.md` before sending PRs.

---

## License

MIT License.  
See `LICENSE` for full text.
