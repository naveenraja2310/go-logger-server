# Go Logger Server

A simple and secure HTTP server written in Go that provides structured logging with automatic file rotation and a REST API to read logs.

## 🚀 Features

- **Structured Logging**: Uses [Zap](https://github.com/uber-go/zap) for high-performance, structured logging
- **Automatic Log Rotation**: Implements [Lumberjack](https://github.com/natefinch/lumberjack) for automatic log file rotation
- **HTTP API**: RESTful endpoint to read log files remotely
- **Security**: Built-in path traversal protection to prevent unauthorized file access
- **Reverse Order**: Logs are displayed with the most recent entries first
- **Lightweight**: Minimal dependencies and resource usage

## 📋 Prerequisites

- Go 1.20 or later
- Basic understanding of REST APIs (optional)

## 🔧 Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/go-logger-server.git
   cd go-logger-server
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Run the server**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080` by default.

## 📖 Usage

### Starting the Server

```bash
go run main.go
```

You should see output similar to:
```
Server starting on :8080
Logger initialized successfully
```

### Reading Logs via API

**Endpoint**: `GET /readfile`

**Parameters**:
- `path` (required): The log file name (e.g., `demo.log`)

**Example Request**:
```bash
curl "http://localhost:8080/readfile?path=demo.log"
```

**Example Response**:
```
2024-01-15 14:30:25 INFO  Application started successfully
2024-01-15 14:30:20 INFO  Database connection established
2024-01-15 14:30:15 INFO  Server initialization complete
```

### Log File Structure

- **Location**: All logs are stored in the `root` directory
- **Format**: Structured JSON or plain text (configurable)
- **Rotation**: Files automatically rotate when they reach 1MB
- **Backup**: Old log files are kept with numbered suffixes (e.g., `demo.log.1`, `demo.log.2`)

## 📁 Project Structure

```
go-logger-server/
├── main.go              # Main application file
├── demo.log         # Default log file
├── go.mod              # Go module file
├── go.sum              # Go dependencies
└── README.md           # This file
```

## ⚙️ Configuration

### Log Rotation Settings
By default, logs rotate with these settings:
- **Max Size**: 1MB per file
- **Max Backups**: 3 old files kept
- **Max Age**: 28 days
- **Compression**: Disabled

### Server Settings
- **Port**: 8080 (default)
- **Host**: localhost
- **Log Level**: Info

## 🔍 API Reference

### GET /readfile

Retrieves the contents of a log file in reverse chronological order.

**Example**:
```bash
# Success
curl "http://localhost:8080/readfile"

## 🛠️ Development

### Building for Production
```bash
go build -o go-logger-server main.go
./go-logger-server
```

### Testing
```bash
go test ./...
```

### Adding Custom Logs
To add your own log entries programmatically:

```go
logger.Info("Your message here")
logger.Error("Error message", zap.String("key", "value"))
```
