# Bookkeeper

Bookkeeper is a command-line application for managing your personal collection of books. It provides an easy way to add, view, and organize your library, storing all information in a local database on your computer.

**Version:** dev.0.0.0

---

## Features

- Add books to your personal library
- View and search your collection
- Local storage using SQLite for simplicity and portability
- Clean, prompt-driven interface for easy data entry

*More features coming soon as development continues!*

---

## Installing Bookkeeper

Pre-built executables are available for each release.  
**To install Bookkeeper:**

1. Go to the [Releases page](https://github.com/SirVoly/bookkeeper/releases).
2. Download the executable for your operating system:
    - For Windows: `bookkeeper-windows-amd64.exe`
    - For macOS: `bookkeeper-darwin-amd64` or `bookkeeper-darwin-arm64`
    - For Linux: `bookkeeper-linux-amd64`
3. Place the executable somewhere in your PATH or run it directly.

**Example usage:**
```
./bookkeeper
```
You’ll be greeted with an interactive prompt to manage your books.

---

## Versioning

This project uses [Semantic Versioning](https://semver.org/) in the format `MAJOR.MINOR.PATCH`:

- **MAJOR:** Breaking changes to data structures or logic.
- **MINOR:** Additions or non-breaking changes.
- **PATCH:** Bugfixes or small tweaks.

Current version: **dev.0.0.0** (development phase)

---

## Technologies Used

- **Go:** Main programming language
- **SQLite:** Embedded database for local storage
- **Cobra:** For structuring CLI commands
- **Survey:** For interactive prompt-driven input
- **Goose:** For database migrations
- **SQLC:** For type-safe database access

---

## Using Source Code

If you prefer to build Bookkeeper yourself:

1. **Clone the repository**
    ```
    git clone https://github.com/YOUR_USERNAME/bookkeeper.git
    cd bookkeeper
    ```

2. **Build the project**
    ```
    go build
    ```
    This creates an executable file in the current directory, named `bookkeeper` (or `bookkeeper.exe` on Windows).

3. **(Optional) Run tests**
    ```
    go test ./...
    ```

4. **Run the program**
    ```
    ./bookkeeper
    ```

---

## License

This project is licensed under the [MIT License](LICENSE).

---
