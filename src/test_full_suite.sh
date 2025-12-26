#!/bin/bash

# Base URL
URL="http://localhost:8080/api"

echo "🧪 Starting Comprehensive Endpoint Test..."
echo "-----------------------------------------"

# Helper function to print headers
print_step() {
  echo -e "\n🔹 $1"
}

# 1. RIDER SIGNUP
print_step "1. POST /api/rides/signup (Rider Signup)"
curl -s -X POST "$URL/rides/signup" \
  -H "Content-Type: application/json" \
  -d '{"id":"r_test", "name":"Test Rider", "email":"test@example.com"}' | python3 -m json.tool

# 2. DRIVER SIGNUP
print_step "2. POST /api/drivers/signup (Driver Signup)"
curl -s -X POST "$URL/drivers/signup" \
  -H "Content-Type: application/json" \
  -d '{"id":"d_test", "name":"Test Driver"}' | python3 -m json.tool

# 3. UPDATE DRIVER LOCATION
print_step "3. POST /api/drivers/:id/location (Update Driver Location)"
curl -s -X POST "$URL/drivers/d_test/location" \
  -H "Content-Type: application/json" \
  -d '{"location":{"lat":40.7128, "lng":-74.0060}}' | python3 -m json.tool

# 4. REQUEST RIDE
print_step "4. POST /api/rides/request (Request Ride)"
RESPONSE=$(curl -s -X POST "$URL/rides/request" \
  -H "Content-Type: application/json" \
  -d '{
    "rider_id":"r_test", 
    "pickup":{"lat":40.7128, "lng":-74.0060}, 
    "dropoff":{"lat":40.7580, "lng":-73.9855}
  }')
echo $RESPONSE | python3 -m json.tool
RIDE_ID=$(echo $RESPONSE | grep -o '"id":"[^"]*' | cut -d'"' -f4)
echo "   > Ride ID extracted: $RIDE_ID"

# 5. GET AVAILABLE RIDES (Driver)
print_step "5. GET /api/drivers/available-rides (Get Available Rides)"
curl -s -X GET "$URL/drivers/available-rides" | python3 -m json.tool

# 6. GET RIDE STATUS
print_step "6. GET /api/rides/:id/status (Get Ride Status)"
curl -s -X GET "$URL/rides/$RIDE_ID/status" | python3 -m json.tool

# 7. ACCEPT RIDE
print_step "7. POST /api/rides/:id/accept (Accept Ride)"
curl -s -X POST "$URL/rides/$RIDE_ID/accept" \
  -H "Content-Type: application/json" \
  -d '{"driver_id":"d_test"}' | python3 -m json.tool

# 8. GET DRIVER RIDES
print_step "8. GET /api/drivers/:id/rides (Get Driver Rides)"
curl -s -X GET "$URL/drivers/d_test/rides" | python3 -m json.tool

# 9. COMPLETE RIDE
print_step "9. POST /api/rides/:id/complete (Complete Ride)"
curl -s -X POST "$URL/rides/$RIDE_ID/complete" \
  -H "Content-Type: application/json" \
  -d '{"driver_id":"d_test"}' | python3 -m json.tool

# 10. GET RIDER HISTORY
print_step "10. GET /api/rides/rider/:id/history (Get Rider History)"
curl -s -X GET "$URL/rides/rider/r_test/history" | python3 -m json.tool

# 11. CANCEL RIDE (Test with new ride)
print_step "11. POST /api/rides/:id/cancel (Cancel Ride)"
# Create fresh ride
RIDE_2_RESP=$(curl -s -X POST "$URL/rides/request" -H "Content-Type: application/json" -d '{"rider_id":"r_test", "pickup":{"lat":10,"lng":10}, "dropoff":{"lat":20,"lng":20}}')
RIDE_ID_2=$(echo $RIDE_2_RESP | grep -o '"id":"[^"]*' | cut -d'"' -f4)
echo "   > Created temporary ride: $RIDE_ID_2"
# Cancel it
curl -s -X POST "$URL/rides/$RIDE_ID_2/cancel" \
  -H "Content-Type: application/json" | python3 -m json.tool

echo -e "\n ✅ Test Suite Complete"
