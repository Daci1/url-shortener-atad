# Project Overview

This repository contains a **full‑stack application** with a Go backend, an Angular frontend, and a PostgreSQL database.

* **Frontend**: Angular 19 (Node.js 24)
* **Backend**: Go 1.24 using the Echo framework
* **Database**: PostgreSQL (managed via Docker Compose)

The backend and database live in the same project and are designed to be run together during development.

---

## Repository Structure

````text
.
├── backend/        # Go backend (Echo)
│   ├── envrc.nix   # direnv configuration (use flake)
│   ├── docker-compose.yml
│   └── ...
├── frontend/       # Angular 19 frontend
│   ├── envrc.nix   # direnv configuration (use flake)
│   ├── flake.nix   # Nix flake for frontend
│   └── ...
└── README.md
````

---

## Prerequisites (Non‑Nix)

If you are **not using Nix/NixOS** (I really recommend that you at least give it a look, I personally love it), make sure you have the following installed:

* **Node.js 24**
* **Angular CLI** (compatible with Angular 19)
* **Go 1.24**
* **Docker** and **Docker Compose**
* **Make** (optional)

---

## Running the Database

The PostgreSQL database is defined **inside the backend project** and is managed with Docker Compose.

From the `backend` directory:

```bash
docker compose up
```

This will start a PostgreSQL instance required by the backend.

---

## Running the Backend

The backend is written in Go and uses the Echo framework.

From the backend directory (or project root, depending on your Makefile):

```bash
make
```

or explicitly:

```bash
make run
```

or directly:
```bash
go run ./cmd/url-shortener
```

The backend expects the PostgreSQL database (via Docker Compose) to be running.

---

## Running the Frontend

The frontend is an Angular 19 application.

From the `frontend` directory:

```bash
ng serve
```

This will start the development server, typically available at `http://localhost:4200`.

---

## Nix / NixOS Development Environment

Both the **backend** and **frontend** have their **own `flake.nix` files**, each defining a dedicated development environment.

A root-level `.envrc` file is included and configured with:

```bash
use flake
```

When entering either the `backend/` or `frontend/` directory, the corresponding flake will be used.

### Using direnv (Recommended)

If you use **direnv**, the setup is automatic.

1. Enable direnv integration in your shell
2. Run (once per directory):

```bash
direnv allow
```

Each project will automatically provide the correct versions of Go or Node.js and all required tooling.

---

## NixOS Without direnv

If you are on **NixOS** but do **not** use `direnv`, you can still use the flake manually.

### Enter the Development Shell

From the project root:

```bash
nix develop
```

This will drop you into a shell with:

* Go 1.24
* Node.js 24
* All required build and runtime dependencies

Once inside the shell, you can run:

* `docker compose up` for the database
* `make` / `make run` for the backend
* `ng serve` for the frontend

---

## Final submission checklist

| User Stories                                                           | Done |
|------------------------------------------------------------------------|------|
| I can customize my short link                                          | [x]  |
| I can view statistics: total clicks, unique visitors (only backend)    | [x]  |
| I can generate a QR code for my short link                             | [x]  |
| I can set an expiration date for my links                              | [ ]  |


| Technical Requirements                                                 | Done |
|------------------------------------------------------------------------|------|
| RESTful API with endpoints for shortening, redirecting                 | [x]  |
| Short code generation (6-8 characters, alphanumeric)                   | [x]  |
| Collision detection and handling                                       | [x]  |
| Rate limiting (e.g., 10 requests/minute per IP)                        | [ ]  |
| Web dashboard showing all links and statistics                         | [ ]  |
| Database for URLs and click events                                     | [x]  |
