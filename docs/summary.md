Below is a **clean Overall Architecture Documentation** that ties your **Domain, Application, and Transport layers together**.

This is typically placed at the **top of project documentation or README**.

---

# Combat Simulator Backend

## Overall Architecture Documentation

---

# 1. Project Overview

The **Combat Simulator Backend** is a layered backend system that simulates turn-based combat between creatures.
Players start a **campaign**, engage in **fights**, and perform **actions** during combat rounds.

The system follows a **layered architecture** that separates responsibilities into three main layers:

1. **Domain Layer** – Core game rules and combat logic
2. **Application Layer** – Game flow and campaign management
3. **Transport Layer** – HTTP API interface for external clients

This separation improves:

* Maintainability
* Testability
* Scalability
* Code clarity

---

# 2. High Level Architecture

```
Client (Frontend / CLI / Postman)
           │
           ▼
   Transport Layer
   (HTTP API Handlers)
           │
           ▼
   Application Layer
   (Campaign Service)
           │
           ▼
     Domain Layer
   (Game Logic Engine)
```

---

# 3. Layer Responsibilities

| Layer           | Responsibility                        |
| --------------- | ------------------------------------- |
| **Transport**   | Handles HTTP requests and responses   |
| **Application** | Manages campaigns and fight lifecycle |
| **Domain**      | Implements combat rules and entities  |

---

# 4. Domain Layer

### Purpose

The **Domain Layer** contains the **core game logic** and represents the rules of the combat simulator.

It is completely **independent of external systems** such as APIs or databases.

### Main Responsibilities

* Define game entities
* Define combat actions
* Implement damage calculations
* Resolve combat rounds
* Maintain fight state transitions

### Key Components

| Component          | Description                           |
| ------------------ | ------------------------------------- |
| `Creature`         | Represents a combat entity            |
| `CreatureTemplate` | Template used to generate creatures   |
| `FightState`       | Current combat state                  |
| `Action`           | Combat actions available to creatures |
| `ResolveRound()`   | Core combat engine                    |

### Example Combat Logic

```
Player chooses action
        ↓
Calculate damage
        ↓
Apply damage to enemy
        ↓
Enemy attacks
        ↓
Apply damage to player
        ↓
Check victory conditions
```

---

# 5. Application Layer

### Purpose

The **Application Layer** orchestrates gameplay by coordinating domain logic and managing game progression.

It ensures that the game follows valid rules such as:

* A campaign must exist before starting a fight
* Only one fight can be active at a time
* Player state persists between fights

### Main Responsibilities

* Campaign lifecycle management
* Fight lifecycle management
* Player action processing
* State persistence during gameplay

### Key Components

| Component         | Description                |
| ----------------- | -------------------------- |
| `CampaignService` | Core gameplay service      |
| `Campaign`        | Player campaign state      |
| `CampaignStore`   | In-memory campaign storage |

### Example Game Flow

```
Start Campaign
      ↓
Start Fight
      ↓
Perform Player Action
      ↓
Resolve Combat (Domain)
      ↓
Update Campaign State
```

---

# 6. Transport Layer

### Purpose

The **Transport Layer** exposes the application functionality through **HTTP APIs**.

It converts incoming requests into service calls and returns structured responses.

### Responsibilities

* Handle HTTP requests
* Parse JSON request bodies
* Call application services
* Convert domain objects into API views
* Return JSON responses

### Key Components

| Component       | Description                           |
| --------------- | ------------------------------------- |
| `Handler`       | HTTP request handler                  |
| `FightView`     | API representation of a fight         |
| `ToFightView()` | Converts domain state to API response |

### Available API Endpoints

| Endpoint          | Method | Description          |
| ----------------- | ------ | -------------------- |
| `/campaign/start` | POST   | Start a new campaign |
| `/fight/start`    | POST   | Start a fight        |
| `/fight/action`   | POST   | Perform an action    |

---

# 7. Complete Request Flow

Example: **Player performs an attack**

```
Client
  │
  ▼
POST /fight/action
  │
  ▼
Transport Layer
(Handler.PerformAction)
  │
  ▼
Application Layer
(CampaignService.PerformAction)
  │
  ▼
Domain Layer
(ResolveRound)
  │
  ▼
Combat logic executed
  │
  ▼
Result returned
  │
  ▼
Transport formats JSON response
  │
  ▼
Client receives result
```

---

# 8. Example Gameplay Sequence

### 1. Start Campaign

```
POST /campaign/start
```

Response:

```
campaignID = "uuid"
```

---

### 2. Start Fight

```
POST /fight/start
```

Creates a fight between player and enemy.

---

### 3. Perform Combat Actions

```
POST /fight/action
```

Each request resolves **one combat round**.

---

# 9. Data Flow Between Layers

```
Transport Layer
      │
      │  HTTP Request
      ▼
Application Layer
      │
      │  Service Calls
      ▼
Domain Layer
      │
      │  Combat Logic
      ▼
Application Layer
      │
      │  Updated State
      ▼
Transport Layer
      │
      │  JSON Response
      ▼
Client
```

---

# 10. Key Design Principles

### Separation of Concerns

Each layer handles a specific responsibility:

| Layer       | Concern       |
| ----------- | ------------- |
| Domain      | Game rules    |
| Application | Game workflow |
| Transport   | Communication |

---

### Testability

* Domain logic is **fully unit-testable**
* Application services are **tested independently**
* Transport layer remains **thin**

---

### Scalability

The architecture allows easy future improvements such as:

* Database persistence
* Multiplayer support
* WebSocket real-time fights
* Additional combat mechanics

---


---

# 11. Benefits of This Architecture

✔ Clear separation of responsibilities
✔ Easy to test business logic
✔ Flexible API layer
✔ Simple state management
✔ Scalable structure for future features

