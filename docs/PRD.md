# 📗 Product Requirements Document (PRD)  
**Project Name:** GoRide – A Ride Sharing App (Pet Project in Golang)  
**Author:** Md. Sayeed Rahman (https://github.com/sayeed1999)  
**Version:** 1.0  
**Date:** August 2025  

---

## 1. Purpose  
To define the product requirements for **GoRide**, a simplified backend-only ride-sharing application built in Go. The project will simulate real-world ride-hailing behavior and provide APIs for customers and drivers to interact through a ride lifecycle.

---

## 2. Goals & Objectives  
- Provide a modular backend system using Go (modular monolith architecture).  
- Allow users to act as **customers** (ride requesters) or **drivers** (ride providers).  
- Implement complete ride lifecycle: request → accept → complete.  
- Use PostgreSQL for persistence.  
- Offer REST API as default interface, with **GraphQL** as an optional query-only interface.  
- Enable clear separation of concerns between modules (auth, ride, customer, driver).  
- Ensure extensibility (e.g., auth system swap without affecting ride logic).

---

## 3. Scope  

### In-Scope (MVP)
- Customer & Driver registration/login (via Auth module).
- Ride request by customer.
- Ride acceptance by driver.
- Ride status transitions: `requested → accepted → in-progress → completed → archived`.
- Driver availability toggling.
- Ride cancellation by customer (before or after acceptance).
- REST API for all operations.
- **GraphQL query interface** for flexible data fetching:
  - Nested queries for ride history, current ride status, user profile data.
- PostgreSQL for all persistence.
- Unit and integration (E2E) testing.

### Out-of-Scope (MVP)
- Real-time location tracking (e.g., GPS).
- Payment gateway or billing.
- In-app messaging or driver rating system.
- Kafka, Redis, or any external services.
- Admin dashboard or back-office tools.
- **GraphQL mutations** (e.g., create/cancel/accept rides).

---

## 4. Assumptions  
- The system will run locally or in a dev environment.
- Users can have both customer and driver roles, selected per session.
- Auth module manages users and roles but is **decoupled** from ride logic.
- Modules interact via well-defined interfaces.
- No real-time updates; status changes are polled.

---

## 5. User Roles  
- **Customer**: A user requesting rides.
- **Driver**: A user accepting rides and fulfilling them.
- A single user can hold both roles but operate in only one role at a time.

---

## 6. Functional Requirements (User Stories)

### 🚗 Customer Stories

#### US-C01: Customer Registration & Login  
**As a customer**, I want to register and log in to the system.  
**Acceptance Criteria**:
- Registration stores user data in Auth module.
- JWT or session token returned after login.

#### US-C02: Request a Ride  
**As a customer**, I want to request a ride by specifying pickup and drop-off locations.  
**Acceptance Criteria**:
- Request is stored with `requested` status.
- Request is visible to available drivers.

#### US-C03: Cancel Ride (Before Acceptance)  
**As a customer**, I want to cancel my ride request before a driver accepts it.  
**Acceptance Criteria**:
- Ride is removed or marked `cancelled`.
- Available drivers no longer see it.

#### US-C04: Cancel Ride (After Acceptance)  
**As a customer**, I want to cancel even after a driver accepts, in case of a poor experience.  
**Acceptance Criteria**:
- Ride is marked `cancelled_by_customer`.
- Driver is notified or flagged.

#### US-C05: View My Ride History (GraphQL-supported)  
**As a customer**, I want to view all my previous rides with driver details.  
**Acceptance Criteria**:
- REST: Returns list of rides.
- GraphQL: Returns ride list with nested driver info.

---

### 🧑‍✈️ Driver Stories

#### US-D01: Driver Registration & Login  
**As a driver**, I want to register/login and access the system.  
**Acceptance Criteria**:
- Handled via Auth module.
- Driver role is assigned upon registration.

#### US-D02: Toggle Availability  
**As a driver**, I want to set myself as available or unavailable for rides.  
**Acceptance Criteria**:
- Available drivers are matched to ride requests.
- GraphQL optional: return availability + current status.

#### US-D03: View Open Ride Requests  
**As a driver**, I want to see unassigned ride requests.  
**Acceptance Criteria**:
- REST endpoint lists nearby or latest ride requests.

#### US-D04: Accept Ride Request  
**As a driver**, I want to accept a ride and begin service.  
**Acceptance Criteria**:
- Ride status changes to `accepted`, then `in-progress`.

#### US-D05: Complete Ride  
**As a driver**, I want to mark a ride complete once it ends.  
**Acceptance Criteria**:
- Status transitions to `completed`.

#### US-D06: View Ride History (GraphQL-supported)  
**As a driver**, I want to view past rides I’ve completed.  
**Acceptance Criteria**:
- REST: Basic history view.
- GraphQL: Nested view with customer and location info.

---

## 7. Non-Functional Requirements  

| Category       | Requirement |
|----------------|-------------|
| Architecture   | Modular Monolith. Each domain is independently testable and loosely coupled. |
| Modularity     | Ride module must not depend on Auth logic; use interface-based dependency. |
| Security       | All APIs require authentication and role-based authorization. |
| Testing        | Unit tests for modules, integration tests for workflows. |
| API Design     | Follow RESTful standards and clean GraphQL schema design (no mutations). |
| Performance    | Lightweight enough for local execution; no real-time constraints. |
| GraphQL        | Query-only support, with nested schemas and role-protected access. |

---

## 8. Constraints  

- Written in **Golang**  
- Backend only – no mobile or web UI  
- Kafka, or real-time systems will be in future improvements (not in MVP).
- GraphQL limited to queries (no mutations in MVP)
- Auth must be replaceable (e.g., SSO later) without affecting business logic  
- Only PostgreSQL used for persistent data  

---

## 9. Success Criteria  

- Customer can request, cancel, and view rides.  
- Driver can view, accept, complete rides.  
- All endpoints covered with tests.  
- GraphQL query layer works for history and profile info.  
- Modular design validated by ability to swap out auth without breakage.  
