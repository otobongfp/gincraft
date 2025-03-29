# Gincraft

A CLI tool for quickly scaffolding Gin web applications with a clean architecture structure.

## Features

- 🚀 Quick project scaffolding
- 📁 Clean architecture structure
- 🔧 Pre-configured with best practices
- 📝 Comprehensive documentation
- 🧪 Built-in testing setup

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/otobongfp/gincraft.git
cd gincraft

# Install dependencies
go mod tidy

# Build the binary
go build -o gincraft

# Install globally (optional)
sudo mv gincraft /usr/local/bin/
```

### Using Go Install

```bash
go install github.com/otobongfp/gincraft@latest
```

## Usage

Create a new Gin project:

```bash
gincraft new myapp
```

This will create a new project with the following structure:

```
myapp/
├── controllers/     # HTTP request handlers
├── routes/         # Route definitions
├── services/       # Business logic
├── models/         # Data models
├── main.go         # Application entry point
├── go.mod          # Go module file
└── .gitignore      # Git ignore file
```

## Development

### Prerequisites

- Go 1.21 or later
- Git

### Local Development

1. Clone the repository:

   ```bash
   git clone https://github.com/otobongfp/gincraft.git
   cd gincraft
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build the binary:

   ```bash
   go build -o gincraft
   ```

4. Test locally:
   ```bash
   ./gincraft new testapp
   ```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT
