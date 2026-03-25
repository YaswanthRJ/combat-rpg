# Combat Simulator Backend

A layered backend system that simulates turn-based combat between creatures.
Players can start campaigns, engage in fights, and perform actions during combat rounds.

---

## 🚀 Features

* Turn-based combat engine
* Campaign-based progression system
* REST API for gameplay interaction
* Clean layered architecture
* Fully unit-testable core logic

---

## 🧠 Architecture

This project follows a **layered architecture**:

```
Client (Frontend / CLI / Postman)
           │
           ▼
   Transport Layer (HTTP API)
           │
           ▼
   Application Layer (Game Flow)
           │
           ▼
     Domain Layer (Game Logic)
```

### Layers

| Layer       | Responsibility                 |
| ----------- | ------------------------------ |
| Domain      | Core combat logic and rules    |
| Application | Campaign and fight lifecycle   |
| Transport   | HTTP APIs and request handling |

---

## ⚙️ Tech Stack

* Go (Golang)
* net/http
* JSON-based REST APIs

---

## 🧩 Core Concepts

### 🧬 Domain Layer (Game Engine)

* Defines **Creature, Actions, FightState**
* Implements **combat logic**
* Handles **damage calculation and win conditions**

### 🎮 Application Layer (Game Flow)

* Manages **campaigns and fights**
* Controls **player progression**
* Ensures valid gameplay rules

### 🌐 Transport Layer (API)

* Exposes REST endpoints
* Handles JSON requests/responses
* Converts domain objects into API views

---

## 📡 API Endpoints

### 1. Start Campaign

```
POST /campaign/start
```

**Request**

```json
{
  "Creature": "Soldier"
}
```

**Response**

```json
{
  "campaignID": "uuid"
}
```

---

### 2. Start Fight

```
POST /fight/start
```

**Request**

```json
{
  "campaignId": "uuid",
  "enemy": "Bandit"
}
```

---

### 3. Perform Action

```
POST /fight/action
```

**Request**

```json
{
  "campaignId": "uuid",
  "action": 0
}
```

---

## ⚔️ Combat System

Each combat round:

```
Player Action
     ↓
Damage to Enemy
     ↓
Enemy Attack
     ↓
Damage to Player
     ↓
Check Win/Loss
```

### Example Actions

| Action      | Type   | Effect        |
| ----------- | ------ | ------------- |
| HeavyAttack | Attack | High damage   |
| FastAttack  | Attack | Low damage    |
| Block       | Defend | Reduce damage |

---

## 🧪 Testing

The project includes unit tests for:

* Campaign creation
* Fight lifecycle validation
* Combat logic correctness
* Player state persistence

Run tests:

```bash
go test ./...
```

---

## ▶️ Running the Project

### Start Server

```bash
go run .
```

Server runs on:

```
http://localhost:8080
```

---

## 🧪 Example Request (cURL)

```bash
curl -X POST http://localhost:8080/campaign/start \
  -H "Content-Type: application/json" \
  -d '{"Creature":"Soldier"}'
```

---

## 🔑 Design Principles

* **Separation of Concerns**
* **Testability**
* **Scalability**
* **Clean Architecture**

---

## 📈 Future Improvements

- Persistent storage (PostgreSQL)
- Smarter enemy AI with decision-making logic
- Additional combat mechanics (status effects, abilities)
- Creature Personalities
- Ai powered taunt system integration

---

## 📌 Project Status

Version: **v1**

This version focuses on:

* Core architecture
* Combat engine
* API design

---

## 👨‍💻 Author

Yaswanth Raj

---
