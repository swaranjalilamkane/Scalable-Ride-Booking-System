# API Endpoint Report

This report documents all available endpoints in the application. All endpoints were verified using `src/test_full_suite.sh`.

## 1. Rider Endpoints

### 1.1 Rider Signup
- **Method:** `POST`
- **URL:** `/api/rides/signup`
- **Description:** Creates a new rider account.
- **Input:** `{"id": "string", "name": "string", "email": "string"}`
- **Test Status:** ✅ **PASSED**

### 1.2 Request Ride
- **Method:** `POST`
- **URL:** `/api/rides/request`
- **Description:** Requests a ride from pickup to dropoff location.
- **Input:** `{"rider_id": "string", "pickup": {"lat": float, "lng": float}, "dropoff": {"lat": float, "lng": float}}`
- **Test Status:** ✅ **PASSED**

### 1.3 Get Ride Status
- **Method:** `GET`
- **URL:** `/api/rides/:id/status`
- **Description:** Retrieves details of a specific ride.
- **Input:** URL Parameter `:id`
- **Test Status:** ✅ **PASSED**

### 1.4 Get Rider History
- **Method:** `GET`
- **URL:** `/api/rides/rider/:id/history`
- **Description:** Lists all rides associated with a rider.
- **Input:** URL Parameter `:id`
- **Test Status:** ✅ **PASSED**

### 1.5 Cancel Ride
- **Method:** `POST`
- **URL:** `/api/rides/:id/cancel`
- **Description:** Cancels a requested ride.
- **Input:** URL Parameter `:id`
- **Test Status:** ✅ **PASSED**

---

## 2. Driver Endpoints

### 2.1 Driver Signup
- **Method:** `POST`
- **URL:** `/api/drivers/signup`
- **Description:** Creates a new driver account.
- **Input:** `{"id": "string", "name": "string"}`
- **Test Status:** ✅ **PASSED**

### 2.2 Update Location
- **Method:** `POST`
- **URL:** `/api/drivers/:id/location`
- **Description:** Updates the driver's current coordinates.
- **Input:** `{"location": {"lat": float, "lng": float}}`
- **Test Status:** ✅ **PASSED**

### 2.3 Get Available Rides
- **Method:** `GET`
- **URL:** `/api/drivers/available-rides`
- **Description:** Lists all rides with status `requested`.
- **Input:** None
- **Test Status:** ✅ **PASSED**

### 2.4 Get Driver Rides
- **Method:** `GET`
- **URL:** `/api/drivers/:id/rides`
- **Description:** Lists historic rides for a specific driver.
- **Input:** URL Parameter `:id`
- **Test Status:** ✅ **PASSED**

---

## 3. Ride Actions (Driver Initiated)

### 3.1 Accept Ride
- **Method:** `POST`
- **URL:** `/api/rides/:id/accept`
- **Description:** Assigns a driver to a ride and changes status to `accepted`.
- **Input:** `{"driver_id": "string"}`
- **Test Status:** ✅ **PASSED**

### 3.2 Complete Ride
- **Method:** `POST`
- **URL:** `/api/rides/:id/complete`
- **Description:** Marks a ride as `completed` and frees the driver.
- **Input:** `{"driver_id": "string"}`
- **Test Status:** ✅ **PASSED**
