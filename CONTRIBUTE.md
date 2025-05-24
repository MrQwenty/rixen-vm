# 🤝 Contributing to Rixen

First off, thanks for your interest in contributing to **Rixen**! 🚀  
Whether you're fixing bugs, adding features, improving documentation, or just trying it out, you're welcome here.

---

## 📦 What is Rixen?

**Rixen (`rx`)** is a CLI tool for developers that simplifies creating and managing virtual machines on macOS using QEMU. It’s built with Go and designed with a dev-first mindset, supporting isolated workspaces, shared folders, and OS auto-setup.

---

## 🛠️ Local Setup

1. **Clone the repo**

```bash
git clone https://github.com/valance-dev/rixen.git
cd rixen
```

2. **Build the CLI**

```bash
go build -o rx
./rx --help
```

3. **Run tests (coming soon)**

```bash
go test ./...
```

---

## 💡 How to Contribute

### 🐛 Found a Bug?
- Check [Issues](https://github.com/valance-dev/rixen/issues) to see if it’s known
- If not, open a new issue with:
  - Reproduction steps
  - OS / CPU details
  - Expected vs actual behavior

### ✨ Want to Add a Feature?
- Open an issue or discussion before starting
- Follow the current CLI patterns (modular with Cobra)
- Keep OS-agnostic logic where possible

### 🧹 Improving Docs?
- Docs live in `/docs` or in `README.md`, `CONTRIBUTE.md`, etc.
- All contributions welcome: clarity, examples, translations

---

## 📂 Code Structure

```
/cmd            → Cobra CLI commands
/internal       → Business logic (VM creation, ISO handling, etc.)
/internal/config → Shared types like VMConfig
```

---

## ✅ Commit Style

- Use clear, concise commit messages:
  ```
  fix: handle missing ISO path in create
  feat: add Windows provider fallback
  docs: improve installation guide
  ```

---

## 📜 License

By contributing, you agree that your contributions will be licensed under the [MIT License](LICENSE).

---

Thanks for helping make Rixen better!  
Feel free to [open an issue](https://github.com/valance-dev/rixen/issues) or ping us with questions.

— MrQwenty
