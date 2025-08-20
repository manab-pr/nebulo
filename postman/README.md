# Nebulo API Testing with Postman

This directory contains Postman collection and environments for testing the Nebulo distributed storage API.

## 📁 Files

- `Nebulo_API_Collection.json` - Complete Postman collection with all API endpoints
- `Local_Development_Environment.json` - Environment for local development testing
- `Docker_Environment.json` - Environment for Docker-based testing

## 🚀 Quick Start

### 1. Import Collection and Environment

1. Open Postman
2. Click **Import** button
3. Import `Nebulo_API_Collection.json`
4. Import one of the environment files:
   - `Local_Development_Environment.json` for local testing
   - `Docker_Environment.json` for Docker testing

### 2. Start Nebulo Services

#### Local Development
```bash
# Terminal 1 - Start main server
make run

# Terminal 2 - Start device server
make run-device
```

#### Docker
```bash
# Start all services
docker-compose up -d
```

### 3. Test the API

Select the appropriate environment in Postman and start testing!

## 📝 Testing Flow

### Basic Testing Sequence

1. **Health Check**
   - Test: `GET /health`
   - Verify server is running

2. **Register a Device**
   - Test: `POST /api/v1/devices/register`
   - Copy the returned `device_id` to environment variables

3. **Send Device Heartbeat**
   - Test: `POST /api/v1/devices/heartbeat`
   - Use the `device_id` from step 2

4. **Upload a File**
   - Test: `POST /api/v1/files/store`
   - Copy the returned `file_id` to environment variables

5. **Check Storage Summary**
   - Test: `GET /api/v1/storage/summary`
   - Verify storage usage

## 🧪 Test Scenarios

### Scenario 1: Basic File Operations
```
1. Register Device → Copy device_id
2. Device Heartbeat → Confirm device online
3. Store File → Copy file_id
4. Get File Metadata → Verify file info
5. List All Files → Confirm file in list
6. Delete File → Cleanup
```

### Scenario 2: Device Management
```
1. Register Device → Copy device_id
2. List All Devices → Verify device in list
3. Device Heartbeat → Update storage info
4. Get Device Storage → Check specific device storage
5. Delete Device → Cleanup
```

### Scenario 3: Search & Analytics
```
1. Register Device → Copy device_id
2. Store Multiple Files → Upload various files
3. Search Files → Test search functionality
4. Get Storage Summary → Check aggregated stats
5. Get File Locations → Verify file placement
```

### Scenario 4: Transfer Management
```
1. Register Device → Device goes online
2. Create Transfer → (This happens automatically when storing files to offline devices)
3. Get Pending Transfers → Check queued transfers
4. Complete Transfer → Mark transfer as done
```

## 🔧 Environment Variables

The collection uses these environment variables:

| Variable | Description | Example |
|----------|-------------|---------|
| `base_url` | Main server URL | `http://localhost:8080` |
| `device_url` | Device server URL | `http://localhost:8081` |
| `device_id` | Device identifier | Auto-populated after registration |
| `file_id` | File identifier | Auto-populated after upload |
| `transfer_id` | Transfer identifier | Auto-populated when needed |
| `file_name` | File name for device operations | `test-file.txt` |

## 📋 API Endpoints Reference

### Main Server (Port 8080)

#### Health & Status
- `GET /health` - Server health check

#### Device Management
- `POST /api/v1/devices/register` - Register new device
- `POST /api/v1/devices/heartbeat` - Send device heartbeat
- `GET /api/v1/devices` - List all devices
- `DELETE /api/v1/devices/{id}` - Remove device

#### File Management
- `POST /api/v1/files/store` - Store file (multipart/form-data)
- `GET /api/v1/files/{fileId}` - Get file metadata
- `GET /api/v1/files` - List all files
- `DELETE /api/v1/files/{fileId}` - Delete file

#### Transfer Management
- `GET /api/v1/transfers/pending/{deviceId}` - Get pending transfers
- `POST /api/v1/transfers/complete` - Mark transfer complete
- `DELETE /api/v1/transfers/{id}` - Cancel transfer

#### Storage Analytics
- `GET /api/v1/storage/summary` - Aggregated storage stats
- `GET /api/v1/storage/device/{deviceId}` - Device-specific storage

#### Search & Query
- `GET /api/v1/files/search?name={query}` - Search files
- `GET /api/v1/files/location/{fileId}` - Get file location

### Device Server (Port 8081)

#### Internal Operations
- `POST /internal/store` - Store file on device (multipart/form-data)
- `GET /internal/files/{id}` - Retrieve file from device
- `GET /internal/storage` - Get device storage info
- `POST /internal/confirm/{fileId}` - Confirm file storage

## 🐛 Troubleshooting

### Common Issues

1. **Connection Refused**
   - Verify servers are running: `curl http://localhost:8080/health`
   - Check Docker containers: `docker-compose ps`

2. **File Upload Fails**
   - Ensure you're using `multipart/form-data`
   - Check file size limits in configuration
   - Verify device has sufficient storage

3. **Device Not Found**
   - Register device first
   - Check device_id in environment variables
   - Verify device heartbeat was sent

4. **Empty Responses**
   - Check database connection (MongoDB)
   - Verify Redis connection (optional)
   - Check server logs: `docker-compose logs`

### Debug Tips

1. **Check Server Logs**
   ```bash
   # Local
   # Check terminal output
   
   # Docker
   docker-compose logs -f nebulo-server
   docker-compose logs -f nebulo-device
   ```

2. **Verify Database State**
   ```bash
   # Connect to MongoDB
   docker exec -it nebulo-mongodb mongosh nebulo
   
   # List collections
   show collections
   
   # Check devices
   db.devices.find().pretty()
   
   # Check files
   db.files.find().pretty()
   ```

3. **Test with cURL**
   ```bash
   # Health check
   curl http://localhost:8080/health
   
   # Register device
   curl -X POST http://localhost:8080/api/v1/devices/register \
     -H "Content-Type: application/json" \
     -d '{"name":"Test Device","ip_address":"127.0.0.1","type":"test","total_storage":1000000000}'
   ```

## 📊 Example Test Data

### Device Registration
```json
{
    "name": "My Laptop",
    "ip_address": "192.168.1.100",
    "type": "laptop",
    "total_storage": 107374182400
}
```

### Device Heartbeat
```json
{
    "device_id": "{{device_id}}",
    "available_storage": 85899345920,
    "used_storage": 21474836480
}
```

### Transfer Completion
```json
{
    "transfer_id": "{{transfer_id}}",
    "success": true,
    "error_msg": ""
}
```

## 📈 Performance Testing

For performance testing, consider:

1. **Load Testing** - Use Postman's Collection Runner for multiple iterations
2. **Concurrent Uploads** - Test multiple file uploads simultaneously
3. **Large Files** - Test with files of various sizes
4. **Network Simulation** - Test with network delays/failures

## 🔒 Security Testing

Test security aspects:

1. **Input Validation** - Send malformed requests
2. **File Types** - Upload various file types
3. **Size Limits** - Test file size restrictions
4. **Path Traversal** - Test for directory traversal vulnerabilities

---

Happy Testing! 🚀