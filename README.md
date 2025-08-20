# Nebulo - Distributed P2P Storage System

[![CI](https://github.com/manab-pr/nebulo/actions/workflows/ci.yml/badge.svg)](https://github.com/manab-pr/nebulo/actions/workflows/ci.yml)
[![Release](https://github.com/manab-pr/nebulo/actions/workflows/release.yml/badge.svg)](https://github.com/manab-pr/nebulo/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/manab-pr/nebulo)](https://goreportcard.com/report/github.com/manab-pr/nebulo)
[![Docker Pulls](https://img.shields.io/docker/pulls/manab-pr/nebulo-server)](https://hub.docker.com/r/manab-pr/nebulo-server)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Nebulo is a distributed storage system that allows you to use multiple devices as a unified storage network. Instead of renting cloud storage, you can utilize your own devices to create a personal distributed storage solution.

## Features

- **Device Management**: Register and manage multiple storage devices
- **File Storage**: Store files across your device network
- **Queued Transfers**: Handle offline devices with queued file transfers
- **Storage Analytics**: Monitor storage usage across all devices
- **File Search**: Search and locate files across the network
- **P2P Communication**: Direct device-to-device file transfers

## Architecture

This project follows Clean Architecture principles with SOLID design patterns:

```
├── cmd/                    # Application entry points
│   ├── server/            # Main API server
│   └── device-server/     # Device-side server
├── config/                # Configuration management
├── container/             # Dependency injection
├── internal/              # Internal packages
├── modules/               # Feature modules (Clean Architecture)
│   ├── devices/          # Device management
│   ├── files/            # File operations
│   ├── transfers/        # Queued transfers
│   ├── storage/          # Storage analytics
│   └── search/           # File search
```

Each module follows the Clean Architecture pattern:
- `domain/` - Business logic and entities
- `data/` - Data access layer
- `presentation/` - HTTP handlers and routes

## API Endpoints

### Device Management
- `POST /api/v1/devices/register` - Register a new device
- `POST /api/v1/devices/heartbeat` - Device heartbeat
- `GET /api/v1/devices` - List all devices
- `DELETE /api/v1/devices/:id` - Remove device

### File Storage
- `POST /api/v1/files/store` - Store a file
- `GET /api/v1/files/:fileId` - Get file metadata
- `GET /api/v1/files` - List all files
- `DELETE /api/v1/files/:fileId` - Delete file

### Queued Transfers
- `GET /api/v1/transfers/pending/:deviceId` - Get pending transfers
- `POST /api/v1/transfers/complete` - Mark transfer complete
- `DELETE /api/v1/transfers/:id` - Cancel transfer

### Storage Overview
- `GET /api/v1/storage/summary` - Storage summary
- `GET /api/v1/storage/device/:deviceId` - Device storage info

### Search & Query
- `GET /api/v1/files/search?name=xyz` - Search files
- `GET /api/v1/files/location/:fileId` - Get file location

### Internal Device Server
- `POST /internal/store` - Store file on device
- `GET /internal/files/:id` - Retrieve file from device
- `GET /internal/storage` - Get device storage info
- `POST /internal/confirm/:fileId` - Confirm file stored

## Getting Started

### Prerequisites
- Go 1.21+
- MongoDB
- Redis (optional)

### Configuration

Copy `.env.example` to `.env` and configure:

```bash
# Server Configuration
PORT=8080
APP_ENV=development
APP_NAME=Nebulo

# Database Configuration
MONGO_URI=mongodb://localhost:27017
MONGO_DATABASE=nebulo

# Redis Configuration (optional)
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# File Storage Configuration
STORAGE_PATH=./storage
MAX_FILE_SIZE=100MB

# Device Network Configuration
DEVICE_SERVER_PORT=8081
HEARTBEAT_INTERVAL=30s
TRANSFER_TIMEOUT=300s
```

### Running the Application

#### Using Docker Compose (Recommended)
```bash
# Clone the repository
git clone https://github.com/manab-pr/nebulo.git
cd nebulo

# Start all services (MongoDB, Redis, Server, Device)
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

#### Using Make (Development)
```bash
# Set up development environment
make dev

# Build both binaries
make build

# Run main server
make run

# Run device server (in another terminal)
make run-device

# Run tests
make test

# Run with coverage
make test-coverage
```

#### Manual Setup
1. Start the main API server:
```bash
go run cmd/server/main.go
```

2. Start the device server (on each device):
```bash
go run cmd/device-server/main.go
```

### Building

```bash
# Using Make
make build                    # Build both binaries
make build-server            # Build main server only
make build-device            # Build device server only
make release-build           # Build for multiple platforms

# Manual build
go build -o nebulo-server cmd/server/main.go
go build -o nebulo-device cmd/device-server/main.go
```

## How It Works

1. **Device Registration**: Each device registers with the main server, providing its storage capacity and network information.

2. **File Storage**: When you upload a file, the system selects an online device with sufficient space and transfers the file.

3. **Offline Handling**: If the target device is offline, the transfer is queued and executed when the device comes back online.

4. **File Retrieval**: Files can be retrieved by their metadata, and the system will locate and serve them from the appropriate device.

5. **Health Monitoring**: Devices send regular heartbeats to maintain their online status and report storage usage.

## Technology Stack

- **Backend**: Go (Gin framework)
- **Database**: MongoDB
- **Cache**: Redis (optional)
- **Architecture**: Clean Architecture
- **Logging**: Uber Zap
- **Validation**: Go Validator
- **Containerization**: Docker & Docker Compose
- **CI/CD**: GitHub Actions

## Deployment

### Docker Hub Images

Pre-built Docker images are available on Docker Hub:

- **Main Server**: `manab-pr/nebulo-server:latest`
- **Device Server**: `manab-pr/nebulo-device:latest`

### Production Deployment

```bash
# Using docker-compose with production settings
curl -LO https://raw.githubusercontent.com/manab-pr/nebulo/main/docker-compose.yml
docker-compose up -d
```

### Environment Variables

Required environment variables for production:

```bash
# Security
JWT_SECRET=your-secure-jwt-secret-key
MONGO_URI=mongodb://username:password@host:port/database

# Optional
REDIS_ADDR=redis:6379
STORAGE_PATH=/app/storage
MAX_FILE_SIZE=1GB
```

### Kubernetes Deployment

Kubernetes manifests are available in the `k8s/` directory (coming soon).

## CI/CD Pipeline

The project uses GitHub Actions for continuous integration and deployment:

### Continuous Integration
- **Code Quality**: golangci-lint, go vet, staticcheck
- **Testing**: Unit tests with race detection
- **Security**: gosec security scanner
- **Build**: Multi-platform binary builds
- **Coverage**: Codecov integration

### Continuous Deployment
- **Docker Images**: Automatic builds for AMD64 and ARM64
- **GitHub Releases**: Automated releases with cross-platform binaries
- **Versioning**: Semantic versioning with git tags

### Secrets Required
For GitHub Actions to work properly, configure these repository secrets:
- `DOCKER_USERNAME`: Docker Hub username
- `DOCKER_PASSWORD`: Docker Hub password/token

## Development

### Prerequisites
- Go 1.21+
- Docker & Docker Compose
- Make (optional, for convenience)

### Quick Start
```bash
# Clone and setup
git clone https://github.com/manab-pr/nebulo.git
cd nebulo
make dev

# Run tests
make test

# Start development servers
make run          # Main server
make run-device   # Device server (in another terminal)
```

### Available Make Targets
```bash
make help         # Show all available targets
make build        # Build both binaries
make test         # Run tests
make lint         # Run linter
make docker       # Build and run with Docker
make clean        # Clean build artifacts
```

### API Testing

#### Quick API Test
```bash
# Run automated API tests
chmod +x scripts/test-api.sh
./scripts/test-api.sh
```

#### Postman Collection
Complete Postman collection and environments are available in the `postman/` directory:

- **Collection**: `postman/Nebulo_API_Collection.json`
- **Local Environment**: `postman/Local_Development_Environment.json`
- **Docker Environment**: `postman/Docker_Environment.json`

**Import Instructions:**
1. Open Postman
2. Import the collection and environment files
3. Select the appropriate environment
4. Start testing the API endpoints

**Test Flow:**
1. Health Check → Verify server is running
2. Register Device → Get device ID
3. Device Heartbeat → Mark device as online  
4. Upload File → Store test file
5. Query Storage → Check storage stats
6. Search Files → Test search functionality

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests for your changes
5. Ensure tests pass (`make test`)
6. Ensure code quality (`make lint`)
7. Commit your changes (`git commit -m 'Add amazing feature'`)
8. Push to the branch (`git push origin feature/amazing-feature`)
9. Create a Pull Request

### Coding Standards
- Follow Go best practices and idioms
- Write tests for new functionality
- Update documentation as needed
- Ensure all CI checks pass

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.