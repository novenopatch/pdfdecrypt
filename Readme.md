
## PDFDecrypt
**pdfdecrypt** started as a simple personal project: I wanted to decrypt my bank statements on my Mac without typing the password every time. One Friday evening, tired of repeating tedious steps, I wrote a small Go script to automate the process.

Now, **pdfdecrypt** is a lightweight, cross-platform CLI tool that:

- Decrypts password-protected PDFs using embedded `qpdf` binaries
- Works on **macOS (ARM64)**, **Linux (AMD64)**, and **Windows (AMD64)**
- Requires **no additional dependencies**
- Is fast, autonomous, and easy to use
- Suitable for anyone who needs to **batch-decrypt PDFs safely and efficiently**

It’s ideal for personal use or for quickly processing multiple PDF files without repetitive manual steps.

---

## Features

* Decrypt all PDFs in a folder or a single file
* Fully autonomous: embedded `qpdf` binary for supported platforms
* Logging to `decrypt.log` with success/error and duration
* Cross-platform support: macOS ARM64, Linux AMD64, Windows AMD64
* Command-line options for flexibility
* Optional overwrite and dry-run modes

---

## Installation

1. Clone the repository:

```bash
git clone <repo-url>
cd PdfProject
```

2. Build the Go executable:

```bash
go build -o pdfdecrypt main.go
```

3. Optionally, move the executable to `/usr/local/bin` (for mac osx )for global access:

```bash
sudo mv pdfdecrypt /usr/local/bin/pdfdecrypt
```

---

## Usage

```bash
./pdfdecrypt [options]
```

### Options

| Flag               | Description                                                         |
| ------------------ | ------------------------------------------------------------------- |
| `--help`           | Show help information                                               |
| `--version`        | Show script version                                                 |
| `--password=PASS`  | PDF password (can also be set in `.env` as `PDF_PASSWORD`)          |
| `--file=FILE`      | Decrypt a specific PDF file                                         |
| `--input-dir=DIR`  | Folder containing PDFs to decrypt (default: current folder)         |
| `--output-dir=DIR` | Folder where decrypted PDFs are saved (default: `dercipts`)         |
| `--overwrite`      | Overwrite original PDF files                                        |
| `--dry-run`        | List files that would be decrypted without actually processing them |

---

### Examples

1. Decrypt all PDFs in the current folder using the password from `.env`:

```bash
./pdfdecrypt
```

2. Decrypt a specific PDF with a password:

```bash
./pdfdecrypt --file document.pdf --password 1234
```

3. Decrypt all PDFs in a folder and save to a custom output directory:

```bash
./pdfdecrypt --input-dir ./pdfs --output-dir ./decrypted --password 1234
```

4. List all PDFs that would be decrypted without actually decrypting:

```bash
./pdfdecrypt --dry-run
```

---

## Environment Variables (`.env`)

You can create a `.env` file in the project folder to provide default settings:

```env
PDF_PASSWORD=960003
PDF_SRC_DIR=/path/to/pdfs
PDF_OUT_DIR=/path/to/decrypted
```

---

## Logging

All decryption actions are logged in `decrypt.log`:

```
2025/08/16 23:05:12 SUCCES: document1.pdf → dercipts/document1.pdf (duration: 0.25s)
2025/08/16 23:05:15 ERREUR: document2.pdf → dercipts/document2.pdf : exit status 1 (duration: 0.20s)
```

---