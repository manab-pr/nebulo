# Nebulo API Reference

Quick reference for all Nebulo API endpoints.

## 🏥 Health & Status

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Server health check |

```bash
curl http://localhost:8080/health
```

## 📱 Device Management

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/v1/devices/register` | Register a new device |
| `POST` | `/api/v1/devices/heartbeat` | Send device heartbeat |
| `GET` | `/api/v1/devices` | List all devices |
| `DELETE` | `/api/v1/devices/{id}` | Remove device |

### Register Device
```bash
curl -X POST http://localhost:8080/api/v1/devices/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My Laptop",
    "ip_address": "192.168.1.100",
    "type": "laptop",
    "total_storage": 107374182400
  }'
```

### Device Heartbeat
```bash
curl -X POST http://localhost:8080/api/v1/devices/heartbeat \
  -H "Content-Type: application/json" \
  -d '{
    "device_id": "DEVICE_ID_HERE",
    "available_storage": 85899345920,
    "used_storage": 21474836480
  }'
```

## 📁 File Management

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/v1/files/store` | Store file (multipart) |
| `GET` | `/api/v1/files/{fileId}` | Get file metadata |
| `GET` | `/api/v1/files` | List all files |
| `DELETE` | `/api/v1/files/{fileId}` | Delete file |

### Store File
```bash
curl -X POST http://localhost:8080/api/v1/files/store \
  -F "file=@/path/to/your/file.txt" \
  -F "target_device=DEVICE_ID_OPTIONAL"
```

## 🔄 Transfer Management

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/v1/transfers/pending/{deviceId}` | Get pending transfers |
| `POST` | `/api/v1/transfers/complete` | Mark transfer complete |
| `DELETE` | `/api/v1/transfers/{id}` | Cancel transfer |

### Complete Transfer
```bash
curl -X POST http://localhost:8080/api/v1/transfers/complete \
  -H "Content-Type: application/json" \
  -d '{
    "transfer_id": "TRANSFER_ID_HERE",
    "success": true,
    "error_msg": ""
  }'
```

## 📊 Storage Analytics

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/v1/storage/summary` | Aggregated storage stats |
| `GET` | `/api/v1/storage/device/{deviceId}` | Device storage info |

```bash
# Get storage summary
curl http://localhost:8080/api/v1/storage/summary

# Get device storage
curl http://localhost:8080/api/v1/storage/device/DEVICE_ID_HERE
```

## 🔍 Search & Query

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/v1/files/search?name={query}` | Search files by name |
| `GET` | `/api/v1/files/location/{fileId}` | Get file location info |

```bash
# Search files
curl "http://localhost:8080/api/v1/files/search?name=test"

# Get file location
curl http://localhost:8080/api/v1/files/location/FILE_ID_HERE
```

## 🖥️ Internal Device Server (Port 8081)

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/internal/store` | Store file on device |
| `GET` | `/internal/files/{id}` | Retrieve file from device |
| `GET` | `/internal/storage` | Get device storage info |
| `POST` | `/internal/confirm/{fileId}` | Confirm file storage |

### Store File on Device
```bash
curl -X POST http://localhost:8081/internal/store \
  -F "file=@/path/to/file.txt" \
  -F "filename=unique-filename.txt"
```

### Get Device Storage
```bash
curl http://localhost:8081/internal/storage
```

## 📋 Response Format

All API responses follow this format:

### Success Response
```json
{
  "message": "Operation completed successfully",
  "data": {
    // Response data here
  }
}
```

### Error Response
```json
{
  "error": "Error message describing what went wrong"
}
```

## 📝 Common Data Types

### Device Object
```json
{
  "id": "64f8b8c8e4b0123456789abc",
  "name": "My Laptop",
  "ip_address": "192.168.1.100",
  "type": "laptop",
  "total_storage": 107374182400,
  "available_storage": 85899345920,
  "used_storage": 21474836480,
  "status": "online",
  "last_heartbeat": "2024-01-15T10:30:00Z",
  "created_at": "2024-01-15T09:00:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

### File Object
```json
{
  "id": "64f8b8c8e4b0123456789def",
  "name": "uuid_original-filename.txt",
  "original_name": "original-filename.txt",
  "size": 1048576,
  "mime_type": "text/plain",
  "stored_on": "64f8b8c8e4b0123456789abc",
  "status": "stored",
  "created_at": "2024-01-15T10:00:00Z",
  "updated_at": "2024-01-15T10:00:00Z"
}
```

### Storage Summary
```json
{
  "total_devices": 3,
  "online_devices": 2,
  "offline_devices": 1,
  "total_storage": 322122547200,
  "used_storage": 64424509440,
  "available_storage": 257698037760,
  "total_files": 42
}
```

## 🚨 HTTP Status Codes

| Code | Meaning |
|------|---------|
| `200` | OK - Request successful |
| `201` | Created - Resource created successfully |
| `400` | Bad Request - Invalid request data |
| `404` | Not Found - Resource not found |
| `500` | Internal Server Error - Server error |

## 🛠️ Testing Tools

### Using cURL
```bash
# Set base URL
export BASE_URL="http://localhost:8080"

# Health check
curl $BASE_URL/health

# Register device and extract ID
DEVICE_ID=$(curl -X POST $BASE_URL/api/v1/devices/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","ip_address":"127.0.0.1","type":"test","total_storage":1000000000}' | \
  jq -r '.data.id')

echo "Device ID: $DEVICE_ID"
```

### Using Postman
1. Import collection: `postman/Nebulo_API_Collection.json`
2. Import environment: `postman/Local_Development_Environment.json`
3. Set environment variables as needed
4. Start testing!

### Using the Test Script
```bash
# Run comprehensive API tests
make test-api
```

## 🔧 Development & Debugging

### Start Services
```bash
# Local development
make run          # Main server (terminal 1)
make run-device   # Device server (terminal 2)

# Docker
make docker       # All services
```

### Check Logs
```bash
# Docker logs
docker-compose logs -f nebulo-server
docker-compose logs -f nebulo-device

# Health checks
make health         # Main server
make health-device  # Device server
```

### Database Inspection
```bash
# Connect to MongoDB
docker exec -it nebulo-mongodb mongosh nebulo

# List collections
show collections

# View devices
db.devices.find().pretty()

# View files
db.files.find().pretty()
```

---

For more detailed testing scenarios and examples, see the [Postman Documentation](postman/README.md).