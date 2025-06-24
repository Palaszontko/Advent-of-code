# Advent of Code

This repository contains my solutions for [Advent of Code](https://adventofcode.com/) challenges. It is organized by day, and each day contains a program to solve the corresponding puzzle. The setup includes a simple `justfile` that automates creating new days, running solutions, and removing completed challenges.

## Features

- **Automatic Setup**: Easily set up a new day with the necessary files and input data.
- **Multi-language Support**: Solutions can be written in Go or Java.
- **Utility Functions**: Common utilities are stored in `cmd/utils/utils.go`.
- **Daily Run Command**: Easily run each day's solution with a simple command.

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

- **Create a new day (Go)**: To create a new day's directory with a Go program template and input file, run:

  ```bash
  just new-day-go 2024 1
  ```

- **Create a new day (Java)**: To create a new day's directory with a Java program template and input file, run:

  ```bash
  just new-day-java 2024 1
  ```

  Replace `2024` with the year and `1` with the day number you want to create. This will:

  - Create a new folder in `cmd/{YEAR}/day_{DAY}`
  - Copy the appropriate template (Go or Java) into the new directory
  - Fetch the input data for the day from the Advent of Code website

- **Run the solution**: To run the program for a specific day (automatically detects Go or Java), use:

  ```bash
  just run 2024 1
  ```

  Replace `2024` with the year and `1` with the day number you want to run. The command will automatically detect whether it's a Go or Java solution and run it accordingly.

- **Remove a day**: If you want to remove the day's folder and files, run:

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
    ├── main.go
    └── Main.java
```

- `cmd/{year}/day_{n}/`: Contains the solution for day `{n}` (Go or Java) and the input file for that day.
- `cmd/utils/`: Contains utility functions used in the solutions (`utils.go`, `directedGraph.go`).
- `go.mod`: Go module file, handling dependencies.
- `justfile`: File containing task automation commands with Just.
- `template/`: Contains template programs for new days (`main.go` for Go, `Main.java` for Java).
