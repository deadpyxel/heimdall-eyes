
# Heimdall-Eyes

A simple uptime monitor written in Go

## Features

> Disclaimer: The goals of this tool are currently changing and being defined

- Simple command structure
- Full CLI Based
- Expose APIs
- Very small footprint (In memory, CPU and codebase)
- Cross platform
- Configurable using config files

## Installation

> WIP

## Configuration

> WIP

## Running Tests

To run tests, run the following command

```bash
go test -cover -v ./...
```
If you want to run the benchmarks:

```bash
go test -bench=. -v ./...
```

## Run Locally

Clone the project

```bash
git clone https://github.com/deadpyxel/heimdall-eyes.git
```

> SSH is extremely recommended

Go to the project directory

```bash
cd heimdall-eyes
```

Build the project locally

```bash
go build -o bin/ -v ./...
```

Run the app

```bash
./bin/cmd
```

## Acknowledgements

 - Gopher's Public Discord
