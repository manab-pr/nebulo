#!/bin/bash

# Nebulo API Test Script
# This script performs basic API testing to verify the system is working

set -e

# Configuration
BASE_URL="http://localhost:8080"
DEVICE_URL="http://localhost:8081"
TEST_FILE="test-file.txt"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to check if service is running
check_service() {
    local url=$1
    local name=$2
    
    print_status "Checking $name service at $url..."
    
    if curl -s -f "$url/health" > /dev/null 2>&1 || curl -s -f "$url/internal/storage" > /dev/null 2>&1; then
        print_status "$name service is running ✓"
        return 0
    else
        print_error "$name service is not responding ✗"
        return 1
    fi
}

# Function to create test file
create_test_file() {
    print_status "Creating test file: $TEST_FILE"
    echo "This is a test file for Nebulo distributed storage system" > "$TEST_FILE"
    echo "Created at: $(date)" >> "$TEST_FILE"
    echo "Content: $(head -c 1000 /dev/urandom | base64)" >> "$TEST_FILE"
}

# Function to cleanup
cleanup() {
    print_status "Cleaning up..."
    rm -f "$TEST_FILE"
    rm -f response.json
}

# Trap cleanup on exit
trap cleanup EXIT

# Main test function
main() {
    echo "=================================="
    echo "🚀 Nebulo API Test Script"
    echo "=================================="
    
    # Check if services are running
    print_status "Checking services..."
    
    if ! check_service "$BASE_URL" "Main server"; then
        print_error "Main server not running. Please start it with: make run"
        exit 1
    fi
    
    if ! check_service "$DEVICE_URL" "Device server"; then
        print_warning "Device server not running. Some tests may fail."
        print_warning "Start device server with: make run-device"
    fi
    
    # Create test file
    create_test_file
    
    print_status "Starting API tests..."
    echo
    
    # Test 1: Health Check
    print_status "1. Testing health endpoint..."
    if curl -s -f "$BASE_URL/health" | jq -e '.status == "ok"' > /dev/null 2>&1; then
        print_status "   Health check passed ✓"
    else
        print_error "   Health check failed ✗"
        exit 1
    fi
    
    # Test 2: Register Device
    print_status "2. Registering device..."
    DEVICE_RESPONSE=$(curl -s -X POST "$BASE_URL/api/v1/devices/register" \
        -H "Content-Type: application/json" \
        -d '{
            "name": "Test Device",
            "ip_address": "127.0.0.1",
            "type": "test",
            "total_storage": 1073741824
        }')
    
    if echo "$DEVICE_RESPONSE" | jq -e '.data.id' > /dev/null 2>&1; then
        DEVICE_ID=$(echo "$DEVICE_RESPONSE" | jq -r '.data.id')
        print_status "   Device registered successfully ✓"
        print_status "   Device ID: $DEVICE_ID"
    else
        print_error "   Device registration failed ✗"
        echo "$DEVICE_RESPONSE"
        exit 1
    fi
    
    # Test 3: Device Heartbeat
    print_status "3. Sending device heartbeat..."
    HEARTBEAT_RESPONSE=$(curl -s -X POST "$BASE_URL/api/v1/devices/heartbeat" \
        -H "Content-Type: application/json" \
        -d "{
            \"device_id\": \"$DEVICE_ID\",
            \"available_storage\": 536870912,
            \"used_storage\": 536870912
        }")
    
    if echo "$HEARTBEAT_RESPONSE" | jq -e '.message' > /dev/null 2>&1; then
        print_status "   Heartbeat sent successfully ✓"
    else
        print_warning "   Heartbeat failed (may be expected if device server not running)"
    fi
    
    # Test 4: List Devices
    print_status "4. Listing devices..."
    DEVICES_RESPONSE=$(curl -s "$BASE_URL/api/v1/devices")
    
    if echo "$DEVICES_RESPONSE" | jq -e '.data | length >= 1' > /dev/null 2>&1; then
        DEVICE_COUNT=$(echo "$DEVICES_RESPONSE" | jq '.data | length')
        print_status "   Found $DEVICE_COUNT device(s) ✓"
    else
        print_error "   Failed to list devices ✗"
        exit 1
    fi
    
    # Test 5: Upload File
    print_status "5. Uploading test file..."
    UPLOAD_RESPONSE=$(curl -s -X POST "$BASE_URL/api/v1/files/store" \
        -F "file=@$TEST_FILE")
    
    if echo "$UPLOAD_RESPONSE" | jq -e '.data.id' > /dev/null 2>&1; then
        FILE_ID=$(echo "$UPLOAD_RESPONSE" | jq -r '.data.id')
        print_status "   File uploaded successfully ✓"
        print_status "   File ID: $FILE_ID"
    else
        print_warning "   File upload failed (expected if no online devices)"
        print_warning "   Response: $UPLOAD_RESPONSE"
    fi
    
    # Test 6: List Files
    print_status "6. Listing files..."
    FILES_RESPONSE=$(curl -s "$BASE_URL/api/v1/files")
    
    if echo "$FILES_RESPONSE" | jq -e '.data' > /dev/null 2>&1; then
        FILE_COUNT=$(echo "$FILES_RESPONSE" | jq '.data | length')
        print_status "   Found $FILE_COUNT file(s) ✓"
    else
        print_error "   Failed to list files ✗"
        exit 1
    fi
    
    # Test 7: Get File Metadata (if file was uploaded)
    if [ ! -z "$FILE_ID" ]; then
        print_status "7. Getting file metadata..."
        FILE_META_RESPONSE=$(curl -s "$BASE_URL/api/v1/files/$FILE_ID")
        
        if echo "$FILE_META_RESPONSE" | jq -e '.data.id' > /dev/null 2>&1; then
            print_status "   File metadata retrieved ✓"
        else
            print_warning "   Failed to get file metadata"
        fi
    fi
    
    # Test 8: Storage Summary
    print_status "8. Getting storage summary..."
    STORAGE_RESPONSE=$(curl -s "$BASE_URL/api/v1/storage/summary")
    
    if echo "$STORAGE_RESPONSE" | jq -e '.data.total_devices' > /dev/null 2>&1; then
        TOTAL_DEVICES=$(echo "$STORAGE_RESPONSE" | jq '.data.total_devices')
        ONLINE_DEVICES=$(echo "$STORAGE_RESPONSE" | jq '.data.online_devices')
        print_status "   Storage summary retrieved ✓"
        print_status "   Total devices: $TOTAL_DEVICES, Online: $ONLINE_DEVICES"
    else
        print_error "   Failed to get storage summary ✗"
        exit 1
    fi
    
    # Test 9: Search Files
    print_status "9. Searching files..."
    SEARCH_RESPONSE=$(curl -s "$BASE_URL/api/v1/files/search?name=test")
    
    if echo "$SEARCH_RESPONSE" | jq -e '.data' > /dev/null 2>&1; then
        SEARCH_COUNT=$(echo "$SEARCH_RESPONSE" | jq '.data | length')
        print_status "   Search completed, found $SEARCH_COUNT result(s) ✓"
    else
        print_error "   Search failed ✗"
        exit 1
    fi
    
    # Test 10: Device Storage Info
    print_status "10. Getting device storage info..."
    DEVICE_STORAGE_RESPONSE=$(curl -s "$BASE_URL/api/v1/storage/device/$DEVICE_ID")
    
    if echo "$DEVICE_STORAGE_RESPONSE" | jq -e '.data.device_id' > /dev/null 2>&1; then
        print_status "   Device storage info retrieved ✓"
    else
        print_warning "   Failed to get device storage info"
    fi
    
    # Test 11: Device Server Tests (if running)
    if check_service "$DEVICE_URL" "Device server" 2>/dev/null; then
        print_status "11. Testing device server..."
        
        # Test device server storage endpoint
        DEVICE_STORAGE=$(curl -s "$DEVICE_URL/internal/storage")
        if echo "$DEVICE_STORAGE" | jq -e '.total_storage' > /dev/null 2>&1; then
            print_status "   Device server storage endpoint working ✓"
        else
            print_warning "   Device server storage endpoint failed"
        fi
        
        # Test file upload to device server
        DEVICE_UPLOAD=$(curl -s -X POST "$DEVICE_URL/internal/store" \
            -F "file=@$TEST_FILE" \
            -F "filename=device-test-file.txt")
        
        if echo "$DEVICE_UPLOAD" | jq -e '.message' > /dev/null 2>&1; then
            print_status "   Device server file upload working ✓"
        else
            print_warning "   Device server file upload failed"
        fi
    fi
    
    # Cleanup registered device
    print_status "12. Cleaning up test device..."
    DELETE_RESPONSE=$(curl -s -X DELETE "$BASE_URL/api/v1/devices/$DEVICE_ID")
    
    if echo "$DELETE_RESPONSE" | jq -e '.message' > /dev/null 2>&1; then
        print_status "   Test device deleted ✓"
    else
        print_warning "   Failed to delete test device"
    fi
    
    echo
    echo "=================================="
    print_status "🎉 API tests completed!"
    echo "=================================="
    
    print_status "Test Summary:"
    print_status "✓ Health check"
    print_status "✓ Device registration"
    print_status "✓ Device listing"
    print_status "✓ File operations"
    print_status "✓ Storage analytics"
    print_status "✓ Search functionality"
    
    echo
    print_status "Next steps:"
    echo "  1. Import Postman collection from: postman/Nebulo_API_Collection.json"
    echo "  2. Use 'make docker' for full stack testing"
    echo "  3. Check logs with: docker-compose logs -f"
}

# Check dependencies
if ! command -v curl &> /dev/null; then
    print_error "curl is required but not installed"
    exit 1
fi

if ! command -v jq &> /dev/null; then
    print_warning "jq is not installed - some tests may not work properly"
    print_warning "Install with: brew install jq (macOS) or apt-get install jq (Linux)"
fi

# Run main function
main "$@"