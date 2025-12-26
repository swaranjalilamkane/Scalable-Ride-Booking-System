#!/bin/bash

# Base URL
URL="http://localhost:8080/api"

echo "🚗 Starting Ride Booking Demo..."
echo "--------------------------------"

# 1. Signup Rider
echo "1. Registering Rider (Rider One)..."
curl -s -X POST "$URL/rides/signup" \
  -H "Content-Type: application/json" \
  -d '{"id":"r1", "name":"Rider One", "email":"rider1@example.com"}' | python3 -m json.tool

echo -e "\n"

# 2. Signup Driver
echo "2. Registering Driver (Driver One)..."
curl -s -X POST "$URL/drivers/signup" \
  -H "Content-Type: application/json" \
  -d '{"id":"d1", "name":"Driver One"}' | python3 -m json.tool

echo -e "\n"

# 3. Request Ride
echo "3. Requesting Ride (Rider One)..."
RESPONSE=$(curl -s -X POST "$URL/rides/request" \
  -H "Content-Type: application/json" \
  -d '{
    "rider_id":"r1", 
    "pickup":{"lat":37.7749, "lng":-122.4194}, 
    "dropoff":{"lat":37.7849, "lng":-122.4094}
  }')
echo $RESPONSE | python3 -m json.tool

# Extract Ride ID (simple grep/sed as dependency-free fallback)
RIDE_ID=$(echo $RESPONSE | grep -o '"id":"[^"]*' | cut -d'"' -f4)
echo "   > Created Ride ID: $RIDE_ID"

echo -e "\n"

# 4. Accept Ride
echo "4. Driver Accepting Ride ($RIDE_ID)..."
curl -s -X POST "$URL/rides/$RIDE_ID/accept" \
  -H "Content-Type: application/json" \
  -d '{"driver_id":"d1"}' | python3 -m json.tool

echo -e "\n"

# 5. Check Status
echo "5. Checking Ride Status..."
curl -s -X GET "$URL/rides/$RIDE_ID/status" | python3 -m json.tool

echo -e "\n"

# 6. Complete Ride
echo "6. Completing Ride..."
curl -s -X POST "$URL/rides/$RIDE_ID/complete" \
  -H "Content-Type: application/json" \
  -d '{"driver_id":"d1"}' | python3 -m json.tool

echo -e "\n--------------------------------"
echo "✅ Demo Complete!"
