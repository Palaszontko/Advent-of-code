# Advent of Code

This repository contains my solutions for [Advent of Code](https://adventofcode.com/) challenges. It is organized by day, and each day contains a Go program to solve the corresponding puzzle. The setup includes a simple `justfile` that automates creating new days, running solutions, and removing completed challenges.

## Features

- **Automatic Setup**: Easily set up a new day with the necessary files and input data.
- **Go-based Solutions**: Each day contains a Go program that solves the problem.
- **Utility Functions**: Common utilities are stored in `cmd/utils/utils.go`.
- **Daily Run Command**: Easily run each day’s solution with a simple command.

## Setup

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/Palaszontko/advent-of-code.git
   cd advent-of-code
   ```

2. Install the necessary dependencies for Go.

3. Install Just.

   ```bash
   brew install just
   ```

4. Set up your session cookie for Advent of Code. You can obtain it by logging into your Advent of Code account and copying the value from your browser.

   ```bash
   export SESSION="your_session_cookie_here"
   ```

## Commands

- **Create a new day**: To create a new day’s directory with a Go program template and input file, run:

  ```bash
  just new-day 2024 1
  ```

  Replace `2024` with the year and `1` with the day number you want to create. This will:

  - Create a new folder in `cmd/{YEAR}/day_{DAY}`
  - Copy a Go template into the new directory
  - Fetch the input data for the day from the Advent of Code website

- **Run the solution**: To run the Go program for a specific day, use:

  ```bash
  just run 2024 1
  ```

  Replace `2024` with the year and `1` with the day number you want to run.

- **Remove a day**: If you want to remove the day’s folder and files, run:

  ```bash
  just remove 2024 1
  ```

  Replace `2024` with the year and `1` with the day number you want to remove.

## File Structure

```
advent-of-code/
├── cmd/
│   ├── 2023/
│   │   └── day_1/
│   │   │   ├── input.txt
│   │   │   └── main.go
│   │   └── ...
│   ├── 2024/
│   │   ├── day_1/
│   │   │   ├── input.txt
│   │   │   └── main.go
│   │   └── ...
│   └── utils/
│       ├── utils.go
│       └── directedGraph.go
├── go.mod
├── justfile
└── template/
    └── main.go
```

- `cmd/{year}/day_{n}/`: Contains the Go solution for day `{n}` and the input file for that day.
- `cmd/utils/`: Contains utility functions used in the solutions (`utils.go`, `directedGraph.go`).
- `go.mod`: Go module file, handling dependencies.
- `justfile`: File containing task automation commands with Just.
- `template/`: Contains a template Go program for new days.

## Rust

A parallel Rust setup is available for any year, living in a Cargo workspace under `rust/`. The Go setup remains usable for any year as well — pick whichever you prefer per-day.

Layout:

```
rust/
├── Cargo.toml            # workspace root
├── utils/                # shared aoc-utils crate (read_file)
├── {year}/day_{n}/       # per-day binary crate
│   ├── Cargo.toml
│   ├── src/main.rs
│   └── input.txt
└── template/             # used by `just new-day-rs`
```

Requirements: `rustup` / `cargo` (Rust 1.85+ for edition 2024).

Commands (work for any year):

- **Create a new day**: `just new-day-rs {YEAR} {DAY}` — scaffolds `rust/{YEAR}/day_{DAY}/`, renders the template, and fetches the input (needs `SESSION`).
- **Run**: `just run-rs {YEAR} {DAY}` — runs `cargo run -p aoc-{YEAR}-day-{DAY}`.
- **Fetch input only**: `just input-rs {YEAR} {DAY}`.
- **Remove**: `just remove-rs {YEAR} {DAY}`.

Inside a day's `main.rs` load the input with `include_str!` (path is relative to the source file, so it works regardless of cwd):

```rust
let input: &str = include_str!("../input.txt");
```
