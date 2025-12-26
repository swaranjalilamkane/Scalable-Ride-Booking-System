# CS6650GroupProject

Authors: Jainum Sanghavi and Swaranjali Lamkane

#  Ride Booking System


## Codebase Description
We are building a scalable ride booking system that connects riders with drivers in real time. The system will allow users to request rides and drivers to accept or decline them, with dynamic ride state tracking. The application will be containerized using Docker and deployed on AWS ECS to ensure scalability and fault isolation.

---

##  Core Architecture

### Rider API Endpoints
- **POST /api/rides/signup:** Create a new rider account  
- **POST /api/rides/request:** Submit a new ride request (pickup, dropoff)  
- **GET /api/rides/:id/status:** Retrieve the current ride status  
- **POST /api/rides/:id/cancel:** Cancel an existing ride request  

### Driver API Endpoints
- **POST /api/drivers/signup:** Create a new driver account  
- **POST /api/drivers/:id/location:** Update driver’s current location  
- **POST /api/rides/:id/accept:** Driver accepts a ride  
- **POST /api/rides/:id/complete:** Driver marks the ride as completed  

---

## System Components

- **Ride Matching Service:** Matches riders with available drivers using proximity-based logic.  
- **Driver Location Service:** Tracks real-time driver locations for matching and availability.  
- **Ride State Service:** Maintains the lifecycle of a ride (requested → accepted → in progress → completed).  
- **Load Balancer + API Gateway:** Routes incoming traffic to multiple backend containers.  
- **Dockerized Microservices:** Each service is containerized and deployable independently.  
- **AWS ECS (Fargate):** Provides horizontal scalability with container orchestration.  
