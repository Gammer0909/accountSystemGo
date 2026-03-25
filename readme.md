# accountSystemGo

A lightweight account system backend for an example website, written in Go. It provides user sign-up and sign-in API endpoints, serves a static frontend, and persists user data to a CSV file.

---

## Features

- **User Registration** — `POST /api/signup` to create a new account
- **User Authentication** — `POST /api/signin` to log in with existing credentials
- **Static File Serving** — serves frontend assets from the `./static` directory
- **CSV-based Persistence** — user data is stored in `data.csv` (no database required)
- **Graceful Shutdown** — handles `SIGINT`/`SIGTERM` with a 10-second shutdown window
- **Password Security** — uses `golang.org/x/crypto` for secure password hashing

---

## Project Structure

```
accountSystemGo/
├── data/           # Data-related utilities
├── security/       # Password hashing and auth helpers
├── server/         # HTTP handler logic (SignUp, SignIn)
├── static/         # Frontend HTML/CSS/JS assets
├── main.go         # Entry point
├── go.mod
└── go.sum
```

---

## Prerequisites

- [Go 1.23+](https://golang.org/dl/)

---

## Getting Started

**1. Clone the repository**

```bash
git clone https://github.com/Gammer0909/accountSystemGo.git
cd accountSystemGo
```

**2. Create the data file**

The server expects a `data.csv` file to exist in the project root before starting:

```bash
touch data.csv
```

**3. Install dependencies**

```bash
go mod download
```

**4. Run the server**

```bash
go run main.go
```

The server will start on [http://localhost:8080](http://localhost:8080).

---

## API Reference

### `POST /api/signup`

Register a new user account.

**Request body (JSON):**
```json
{
  "username": "yourname",
  "password": "yourpassword"
}
```

### `POST /api/signin`

Authenticate an existing user.

**Request body (JSON):**
```json
{
  "username": "yourname",
  "password": "yourpassword"
}
```

---

## Dependencies

| Package | Purpose |
|---|---|
| [`gorilla/mux`](https://github.com/gorilla/mux) | HTTP request routing |
| [`gocarina/gocsv`](https://github.com/gocarina/gocsv) | CSV serialization for user storage |
| [`golang.org/x/crypto`](https://pkg.go.dev/golang.org/x/crypto) | Secure password hashing |

---

## License

This project is unlicensed and intended for educational/example purposes.
