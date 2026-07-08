# Distributed Ride Booking Platform (Capstone Project)

## 📌 Project Overview
The platform simulates the core workflows of an Uber-like application, orchestrating multiple microservices to handle trip previews, ride requesting, real-time driver matching, and secure payments. 

### Core Tech Stack:
- **Language:** Go (1.26+)
- **Database:** PostgreSQL (with BSON mapping)
- **Message Broker:** RabbitMQ (for event-driven communication)
- **APIs & Protocols:** HTTP/REST, gRPC, WebSockets (for real-time updates)
- **Telemetry:** OpenTelemetry & Jaeger (Distributed Tracing)
- **Deployment & Orchestration:** Docker, Kubernetes (K8s), and Tilt for local development

---

## 🔄 Trip & Payment Lifecycle (Target Flow)

Here is the event-driven sequence diagram detailing the end-to-end trip creation, driver dispatch, and payment flow:

```mermaid
sequenceDiagram
  participant User as User
  participant APIGateway as APIGateway
  participant TripService as TripService
  participant DriverService as DriverService
  participant Driver as Driver
  participant PaymentService as PaymentService
  participant Stripe as Stripe API
  participant OSRM as OSRM API

  User ->> APIGateway: Preview Trip Route
  APIGateway ->> TripService: gRPC: PreviewTrip
  TripService ->> OSRM: HTTP: Get & Calculate Route in real world map
  OSRM ->> TripService: HTTP: Route coordinates
  TripService ->> APIGateway: HTTP: Trip route info
  APIGateway ->> User: Display Trip Preview UI

  User ->> APIGateway: Create Trip Request
  APIGateway ->> TripService: gRPC: CreateTrip
  Note over TripService, DriverService: Trip Exchange
  TripService -->>+ DriverService: trip.event.created
  Note right of TripService: Event: New trip needs a driver
  DriverService ->> Driver: driver.cmd.trip_request
  Note right of DriverService: Command: New trip available
  Driver ->> APIGateway: WebSocket: driver.cmd.trip_accept
  Note left of Driver: Command: Accept trip request
  APIGateway -->> TripService: driver.cmd.trip_accept
  Note right of APIGateway: Command forwarded to RabbitMQ (DriverTripResponseQueue)
  Note over TripService: Process driver acceptance, update trip status...
  TripService -->> PaymentService: trip.event.driver_assigned
  Note right of TripService: Event: Create payment session
  PaymentService ->> Stripe: Create Checkout Session
  Stripe -->> PaymentService: Session Created
  PaymentService -->> APIGateway: payment.event.session_created
  Note right of PaymentService: Event: Payment UI ready
  APIGateway ->> User: Show Payment UI
  Note right of APIGateway: WebSocket: Payment form
  User ->> Stripe: Complete Payment
  Stripe ->> APIGateway: Webhook: Payment Success
  APIGateway -->>+ TripService: payment.event.success
  Note right of PaymentService: Event: Payment completed
  
  Note right of TripService: Trip considered done
```

---
