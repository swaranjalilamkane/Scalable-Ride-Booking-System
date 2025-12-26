# Ride Booking App - Single Instance Manual

This guide explains how to build, run, and test the Ride Booking application in a single-instance mode.

> **Note:** This version uses **in-memory storage**. If you stop the server, all user and ride data will be lost.

## 1. Prerequisites
- **Go 1.25+** installed (`go version` to check)
- **Terminal** (Mac/Linux)

## 2. Compile and Run

1.  Open your terminal and navigate to the `src` directory:
    ```bash
    cd src
    ```

2.  Build the application:
    ```bash
    go build -o ride-app .
    ```

3.  Run the server:
    ```bash
    ./ride-app
    ```
    *You should see: "Starting Unified Ride Booking Service on port 8080..."*

## 3. Using the App

### Option A: Automated Demo (Recommended)
We have included a script that runs through the entire flow (Signup -> Request -> Accept -> Complete).

1.  Keep the server running in one terminal.
2.  Open a **new terminal window**.
3.  Navigate to `src` and run:
    ```bash
    sh demo.sh
    ```

### Option B: Manual API Testing
You can use `curl` to interact with the API manually.

#### 1. Rider Signup
```bash
curl -X POST http://localhost:8080/api/rides/signup \
  -H "Content-Type: application/json" \
  -d '{"id":"r1", "name":"Alice", "email":"alice@test.com"}'
```

#### 2. Driver Signup
```bash
curl -X POST http://localhost:8080/api/drivers/signup \
  -H "Content-Type: application/json" \
  -d '{"id":"d1", "name":"Bob"}'
```

#### 3. Request a Ride
```bash
curl -X POST http://localhost:8080/api/rides/request \
  -H "Content-Type: application/json" \
  -d '{
    "rider_id":"r1", 
    "pickup":{"lat":37.77, "lng":-122.41}, 
    "dropoff":{"lat":37.80, "lng":-122.45}
  }'
```
*Tip: Copy the `id` from the response (e.g., `"ride_1"`).*

#### 4. Accept Ride (as Driver)
```bash
curl -X POST http://localhost:8080/api/rides/ride_1/accept \
  -H "Content-Type: application/json" \
  -d '{"driver_id":"d1"}'
```

#### 5. Check Status
```bash
curl http://localhost:8080/api/rides/ride_1/status
```

## 4. Stopping the Server
Press `Ctrl + C` in the terminal where the server is running.
