Here is the **documentation for your Transport Layer (`transport`)**.

---

# Transport Layer Documentation

## 1. Overview

The **Transport Layer** exposes the application functionality through **HTTP APIs**.

Its responsibilities are to:

* Accept **HTTP requests**
* Parse **JSON request bodies**
* Call **Application Layer services**
* Convert **domain objects into API-friendly views**
* Return **JSON responses**

This layer **does not contain business logic**.
It simply acts as the **communication interface between clients and the backend**.

---

# 2. Handler

## File: `handler.go`

### Purpose

The `Handler` struct defines all HTTP endpoints and connects them to the **CampaignService**.

---

### `Handler`

```go
type Handler struct {
	service *app.CampaignService
}
```

Dependency:

| Dependency        | Purpose             |
| ----------------- | ------------------- |
| `CampaignService` | Executes game logic |

---

### Constructor

```go
func NewHandler(service *app.CampaignService) *Handler
```

Creates a handler instance with access to the application service.

---

# 3. API Endpoints

The transport layer exposes **three endpoints**.

| Endpoint          | Method | Purpose                      |
| ----------------- | ------ | ---------------------------- |
| `/campaign/start` | POST   | Start a new campaign         |
| `/fight/start`    | POST   | Start a fight                |
| `/fight/action`   | POST   | Perform an action in a fight |

---

# 4. Start Campaign Endpoint

## `POST /campaign/start`

### Handler Function

```go
func (h *Handler) StartCampaign(w http.ResponseWriter, r *http.Request)
```

---

### Request Body

```json
{
  "Creature": "Bandit"
}
```

Field:

| Field    | Description          |
| -------- | -------------------- |
| Creature | Player creature type |

---

### Processing Steps

```
Decode JSON request
       ↓
Call CampaignService.StartCampaign()
       ↓
Generate campaign ID
       ↓
Return JSON response
```

---

### Response

```json
{
  "campaignID": "uuid-value"
}
```

---

### Error Cases

| Error                 | HTTP Code |
| --------------------- | --------- |
| Invalid JSON          | 400       |
| Invalid creature type | 400       |

---

# 5. Start Fight Endpoint

## `POST /fight/start`

### Handler Function

```go
func (h *Handler) StartFight(w http.ResponseWriter, r *http.Request)
```

---

### Request Body

```json
{
  "campaignId": "uuid",
  "enemy": "Bandit"
}
```

Fields:

| Field      | Description         |
| ---------- | ------------------- |
| campaignId | Campaign identifier |
| enemy      | Enemy creature type |

---

### Processing Steps

```
Decode request
      ↓
Call CampaignService.StartFight()
      ↓
Create FightState
      ↓
Convert to FightView
      ↓
Return JSON response
```

---

### Response

Example:

```json
{
  "player": {
    "hp": 30,
    "maxHP": 30,
    "actions": [
      { "id": 0, "name": "Heavy Attack" },
      { "id": 1, "name": "Fast Attack" }
    ]
  },
  "enemy": {
    "name": "Bandit",
    "hp": 20,
    "maxHP": 20,
    "description": "A common bandit"
  },
  "status": "ongoing"
}
```

---

# 6. Perform Action Endpoint

## `POST /fight/action`

### Handler Function

```go
func (h *Handler) PerformAction(w http.ResponseWriter, r *http.Request)
```

---

### Request Body

```json
{
  "campaignId": "uuid",
  "action": 0
}
```

Fields:

| Field      | Description         |
| ---------- | ------------------- |
| campaignId | Campaign identifier |
| action     | Action ID           |

---

### Action Validation

Before calling the service, the handler verifies that the action exists:

```go
if _, ok := domain.ActionPool[req.Action]; !ok
```

---

### Processing Steps

```
Decode request
      ↓
Validate action
      ↓
Call CampaignService.PerformAction()
      ↓
Resolve combat round
      ↓
Build fight view
      ↓
Return result + updated view
```

---

### Response

```json
{
  "result": {
    "playerDamageDealt": 10,
    "enemyDamageDealt": 5,
    "actionNumber": 1,
    "fightEnded": false
  },
  "view": {
    "player": { ... },
    "enemy": { ... },
    "status": "ongoing"
  }
}
```

---

# 7. View Models

## File: `fight_view.go`

### Purpose

View models transform **domain data structures into API responses**.

Domain objects often contain extra data or internal structure that should not be exposed directly.

---

# `FightView`

Represents the **fight state returned to the client**.

```go
type FightView struct {
	Player PlayerView
	Enemy  EnemyView
	Status string
}
```

Fields:

| Field  | Description   |
| ------ | ------------- |
| Player | Player status |
| Enemy  | Enemy status  |
| Status | Fight state   |

---

# `PlayerView`

```go
type PlayerView struct {
	HP      int
	MaxHP   int
	Actions []ActionView
}
```

Fields:

| Field   | Description       |
| ------- | ----------------- |
| HP      | Current health    |
| MaxHP   | Maximum health    |
| Actions | Available actions |

---

# `EnemyView`

```go
type EnemyView struct {
	Name        string
	HP          int
	MaxHP       int
	Description string
}
```

Fields:

| Field       | Description       |
| ----------- | ----------------- |
| Name        | Enemy name        |
| HP          | Current HP        |
| MaxHP       | Maximum HP        |
| Description | Enemy description |

---

# `ActionView`

```go
type ActionView struct {
	ID   domain.Action
	Name string
}
```

Used to expose **player actions to the client**.

Example:

```json
{
  "id": 0,
  "name": "Heavy Attack"
}
```

---

# 8. View Conversion

## `ToFightView()`

Converts a **FightState** into a **FightView**.

```go
func ToFightView(state *domain.FightState, template domain.CreatureTemplate) FightView
```

---

### Conversion Steps

```
Read fight state
       ↓
Extract player HP
       ↓
Convert player actions
       ↓
Extract enemy data
       ↓
Map fight status
       ↓
Return FightView
```

---

# 9. Status Mapping

## `mapStatus()`

Converts internal fight status to API-friendly strings.

```go
func mapStatus(s domain.FightStatus) string
```

Mapping:

| Domain Status | API Response |
| ------------- | ------------ |
| Ongoing       | `"ongoing"`  |
| PlayerWon     | `"won"`      |
| PlayerLost    | `"lost"`     |

---

# 10. Transport Layer Architecture

```
Client
   ↓
HTTP Request
   ↓
Transport Layer (Handlers)
   ↓
Application Layer (CampaignService)
   ↓
Domain Layer (Combat Logic)
   ↓
Result Returned
```

---

# 11. Layer Responsibilities Summary

| Layer           | Responsibility            |
| --------------- | ------------------------- |
| **Transport**   | HTTP API and JSON         |
| **Application** | Campaign and game flow    |
| **Domain**      | Combat rules and entities |

---

# 12. Full Backend Flow

```
Client
  ↓
POST /campaign/start
  ↓
CampaignService.StartCampaign()
  ↓
Domain.GenerateCreature()
  ↓
Campaign created
```

```
Client
  ↓
POST /fight/start
  ↓
CampaignService.StartFight()
  ↓
Create FightState
  ↓
Return FightView
```

```
Client
  ↓
POST /fight/action
  ↓
CampaignService.PerformAction()
  ↓
Domain.ResolveRound()
  ↓
Return ActionResult + Updated View
```

---

If you want, I can also give you a **very clean "Project Architecture Documentation" page** (with diagrams + explanations) that ties **Domain + Application + Transport together**, which is usually expected in backend project documentation.
