# 📘 Business Requirements Document (BRD)  
**Project Name:** GoRide – A Ride Sharing App (Pet Project in Golang)  
**Author:** Md. Sayeed Rahman (https://github.com/sayeed1999)  
**Version:** 1.1  
**Date:** August 2025  

---

## 1. Purpose
The purpose of this document is to outline the high-level business requirements for **GoRide**, a simplified ride-sharing application built using the Go programming language. This project serves as a learning tool and simulates basic features of a real-world ride-hailing service like Pathao or Uber.

---

## 2. Background
Ride-sharing systems are a classic example of real-time, distributed systems. GoRide will help the developer (you) understand how to design and implement backend services with concurrency, APIs, caching, persistence, and modular architecture, while building something functional and fun.

---

## 3. Objectives
- Build a simple backend system using Go.
- Simulate basic ride-sharing logic (request, accept, complete).
- Use PostgreSQL for persistent storage.
- Use Redis for caching frequently accessed data like driver location and availability.
- Practice software engineering best practices including unit and integration testing.
- **Allow future replacement of auth layer (e.g., switch from basic login to SSO) without affecting ride-sharing core api.**
- Gain experience with documenting, testing, and extending APIs.

---

## 4. Scope

### In-Scope:
- User registration/login (as rider or driver)
- Ride request and ride acceptance
- Ride status management (requested → accepted → in-progress → completed)
- Driver availability toggling
- Mock location handling (with Redis caching)
- Redis integration for performance (e.g., storing active drivers)
- PostgreSQL as primary database
- Rate limiting and basic security practices
- Unit testing for individual components
- End-to-end (E2E) or integration testing for ride flow
- Basic logging and API testing

### Out-of-Scope (MVP):
- Real-time GPS tracking or websockets
- Payment gateway
- In-app chat or ratings
- Advanced pricing or surge logic
- Mobile or web frontend

---

## 5. Stakeholders

| Role       | Description              |
|------------|--------------------------|
| Developer  | You (solo project owner) |
| End User   | Simulated riders/drivers |
| System     | GoRide backend service   |

---

## 6. Assumptions
- The system will be backend-only.
- Location data will be static or mocked (no live GPS).
- Redis will be used for caching high-read data like driver availability.
- PostgreSQL will be used for persistence.
- The environment will be local/dev only; but must be ready for scaling or deployment.
- Simulated user actions via CLI/Postman/E2E testing.

---

## 7. Business Requirements

| ID     | Requirement                                                                 |
|--------|------------------------------------------------------------------------------|
| BR-01 | The system must allow users to register as a rider or driver.                |
| BR-02 | Riders must be able to request a ride by providing pickup and drop-off info. |
| BR-03 | Drivers must be able to view available ride requests.                        |
| BR-04 | Drivers must be able to accept ride requests.                                |
| BR-05 | The system must track the status of each ride.                               |
| BR-06 | Drivers can set themselves as available or unavailable.                      |
| BR-07 | Riders can cancel a ride before it's accepted.                               |
| BR-08 | All ride actions should be logged for debugging.                             |
| BR-09 | Redis must be used to cache and quickly access driver availability/location. |
| BR-10 | The system must include unit and integration/E2E tests for all features.     |

---

## 8. Architecture Principles

- **Modular Monolith**: The system must be built as a modular monolith, where core domains like Auth, Ride Management, Driver, and Rider are organized into well-separated packages/modules with clear interfaces.

- **Auth Decoupling**: Authentication and authorization must be fully decoupled from the ride-sharing domain logic, so that the auth mechanism (e.g., basic login) can be replaced with alternatives (e.g., SSO, OAuth2) without affecting core functionality.

- **Transport & Infrastructure Independence**: The domain logic (e.g., ride requests, matching) must not depend on REST handlers, database implementations, or external services directly. All interactions should occur via interfaces and abstractions, enabling easier testing and future changes.

- **Dependency Inversion**: High-level business logic must not depend on low-level implementation details. Use interfaces and dependency injection to manage interactions between modules and infrastructure.

---

## 9. Constraints
- Must use Go (Golang) for all backend logic.
- PostgreSQL must be used as the primary database.
- Redis must be integrated for caching.
- REST API only (no UI or websockets).
- CLI/Postman-based testing only.
- Must be modular and testable.
- Limited concurrency (basic simulation only).

---

## 10. Timeline (Estimated)

| Phase        | Duration     |
|--------------|--------------|
| BRD (Business Requirements Document)         | ✅ Complete  |
| PRD (Product Requirements Document)          | 2 days       |
| SDD (Software/System Design Document)        | 2–3 days     |
| TDD (Technical Design Documentation)         | 2 days       |
| Implementation                               | Ongoing      |
---

## 10. Success Criteria
- Full simulation of a ride lifecycle (request → accept → complete).
- Redis used for caching driver availability and location.
- PostgreSQL used for storing users, rides, and historical data.
- APIs are covered by unit tests and at least one E2E test flow.
- Modular Go code with room for future extensibility.

---
