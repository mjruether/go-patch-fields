# Contributing to gopatchfields

## Welcome!

We appreciate your interest in contributing to gopatchfields. This document provides guidelines for contributing to the project.

## Getting Started

1. Fork the repository
2. Clone your fork
3. Create a new branch for your feature or bugfix
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Development Setup

1. Ensure you have Go 1.21 or later installed
2. Install development dependencies:
   ```bash
   go mod download
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   ```

## Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -race -coverprofile=coverage.txt -covermode=atomic ./...

# Run linters
golangci-lint run
```

## Submitting Changes

1. Ensure all tests pass
2. Run linters and fix any issues
3. Commit your changes with a clear, descriptive commit message
4. Push to your fork and submit a pull request

## Code Style

- Follow standard Go formatting (use `gofmt`)
- Write clear, concise comments
- Keep functions and methods focused and relatively short
- Add tests for new functionality
- Ensure good test coverage

## Reporting Issues

- Use GitHub Issues to report bugs or request features
- Provide a clear description of the issue
- Include steps to reproduce for bugs
- Provide example code or use cases for feature requests

## Code of Conduct

- Be respectful and inclusive
- Provide constructive feedback
- Collaborate and help each other

## Questions?

If you have any questions, feel free to open an issue or reach out to the maintainers.

Thank you for contributing to gopatchfields!
