from locust import HttpUser, task, between
import random
import uuid
from gevent import spawn, sleep

def random_location():
    return {
        "lat": round(random.uniform(37.7, 37.8), 5),
        "lng": round(random.uniform(-122.5, -122.4), 5),
    }

# -----------------------------
# Rider Simulation
# -----------------------------
class RiderUser(HttpUser):

    def on_start(self):
        self.rider_id = str(uuid.uuid4())
        # Rider signup is not tracked on UI
        self.client.post(
            "/api/rides/signup",
            json={"id": self.rider_id, "name": f"Rider-{self.rider_id[:5]}"},
            name=None  # ignore in UI
        )

    @task(3)
    def request_ride(self):
        ride_data = {
            "rider_id": self.rider_id,
            "pickup": random_location(),
            "dropoff": random_location()
        }
        response = self.client.post(
            "/api/rides/request",
            json=ride_data,
            name="Rides Requested"
        )
        if response.status_code == 200:
            ride_id = response.json().get("ride_id")
            if ride_id:
                spawn(self.wait_for_completion, ride_id)

    def wait_for_completion(self, ride_id):
        max_wait = 180
        waited = 0
        interval = 5
        while waited < max_wait:
            # Polling status is not shown on UI
            status_resp = self.client.get(f"/api/rides/{ride_id}/status", name=None)
            status = status_resp.json().get("status")
            if status == "completed":
                break
            sleep(interval)
            waited += interval

# -----------------------------
# Driver Simulation
# -----------------------------
class DriverUser(HttpUser):

    @task
    def accept_available_ride(self):
        # Get available drivers (not shown in UI)
        drivers_resp = self.client.get("/api/drivers/available-drivers", name=None)
        available_drivers = drivers_resp.json()
        if not available_drivers:
            return
        driver_id = available_drivers[0]["id"]

        # Get available rides (shown in UI)
        rides_resp = self.client.get("/api/drivers/available-rides", name="Available Rides")
        available_rides = rides_resp.json()
        if not available_rides:
            return
        ride_id = available_rides[0]["id"]

        # Accept ride (not tracked)
        accept_resp = self.client.post(
            f"/api/rides/{ride_id}/accept",
            json={"driver_id": driver_id},
            name=None
        )

        if accept_resp.status_code == 200:
            spawn(self.complete_ride_later, ride_id, driver_id)

    def complete_ride_later(self, ride_id, driver_id):
        sleep(60)
        self.client.post(
            f"/api/rides/{ride_id}/complete",
            json={"driver_id": driver_id},
            name="Rides Completed"
        )
