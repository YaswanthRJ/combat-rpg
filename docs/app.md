Here is the **documentation for your Application Layer (`app`)**.

---

# Application Layer Documentation

## 1. Overview

The **Application Layer** manages the **game flow and orchestration of domain logic**.

While the **Domain Layer** defines rules (combat, creatures, actions), the **Application Layer**:

* Manages **campaigns**
* Controls **fight lifecycle**
* Coordinates **player actions**
* Maintains **game state storage**
* Ensures **valid game flow**

It acts as a **bridge between the Domain Layer and external layers** (such as APIs or controllers).

---

# 2. Campaign System

## File: `campaign.go`

### Purpose

Defines the **Campaign entity**, which represents a player's progress through multiple fights.

### `Campaign`

```go
type Campaign struct {
	ID            string
	Player        domain.Creature
	Fights        map[string]*domain.FightState
	ActiveFightID string
}
```

Fields:

| Field           | Description                      |
| --------------- | -------------------------------- |
| `ID`            | Unique campaign identifier       |
| `Player`        | Player's creature                |
| `Fights`        | All fights in this campaign      |
| `ActiveFightID` | ID of the currently active fight |

---

### Campaign Responsibilities

A campaign stores:

* Player state
* All fights fought in the campaign
* Current active fight
* Player HP persistence across fights

---

# 3. Campaign Store

## File: `campaign_store.go`

### Purpose

Provides **in-memory storage** for campaigns.

---

### `CampaignStore`

```go
type CampaignStore struct {
	mu        sync.Mutex
	campaigns map[string]*Campaign
}
```

Fields:

| Field       | Description                 |
| ----------- | --------------------------- |
| `mu`        | Mutex for thread safety     |
| `campaigns` | Map storing campaigns by ID |

---

### Why Mutex?

The mutex ensures **safe concurrent access** if multiple requests access the store.

Example:

```go
store.mu.Lock()
defer store.mu.Unlock()
```

---

### `NewCampaignStore()`

Creates a new store instance.

```go
func NewCampaignStore() *CampaignStore
```

Initializes:

```go
campaigns: make(map[string]*Campaign)
```

---

# 4. Campaign Service

## File: `campaign_service.go`

### Purpose

The **CampaignService** is the **main application service** that manages campaign gameplay.

It coordinates:

* Campaign creation
* Fight creation
* Player actions
* Game state updates

---

## `CampaignService`

```go
type CampaignService struct {
	store *CampaignStore
}
```

Dependencies:

| Dependency      | Purpose                       |
| --------------- | ----------------------------- |
| `CampaignStore` | Persistent state of campaigns |

---

### Constructor

```go
func NewCampaignService(store *CampaignStore) *CampaignService
```

Creates a new service instance.

---

# 5. Campaign Operations

---

# `StartCampaign()`

Creates a new campaign.

```go
func (s *CampaignService) StartCampaign(playerCreatureName string) (string, error)
```

### Steps

1. Lock store
2. Generate campaign ID
3. Generate player creature using domain
4. Create campaign object
5. Store campaign
6. Return campaign ID

### Example Flow

```
Client Request
      ↓
StartCampaign("Bandit")
      ↓
domain.GenerateCreature()
      ↓
Campaign created
      ↓
Campaign stored
      ↓
Return campaignID
```

---

# `StartFight()`

Creates a fight within a campaign.

```go
func (s *CampaignService) StartFight(
    campaignID string,
    enemyCreatureName string,
) (*domain.FightState, domain.CreatureTemplate, error)
```

---

### Validation

Checks:

1. Campaign exists
2. No active fight currently running
3. Enemy type is valid

---

### Steps

```
Fetch campaign
     ↓
Check active fight
     ↓
Fetch enemy template
     ↓
Generate enemy creature
     ↓
Create FightState
     ↓
Store fight in campaign
     ↓
Set ActiveFightID
```

---

### Return Values

| Value              | Purpose             |
| ------------------ | ------------------- |
| `FightState`       | Current fight state |
| `CreatureTemplate` | Enemy info for UI   |
| `error`            | Failure reason      |

---

# `PerformAction()`

Executes a player action during a fight.

```go
func (s *CampaignService) PerformAction(
    campaignID string,
    action domain.Action,
) (domain.ActionResult, *domain.FightState, error)
```

---

### Steps

```
Fetch campaign
     ↓
Get active fight
     ↓
Call domain.ResolveRound()
     ↓
Update campaign player state if fight ended
     ↓
Return action result + fight state
```

---

### Important Behavior

When a fight ends:

```go
currentCampaign.Player = fight.Player
```

This ensures:

✔ Player HP persists across fights.

---

# 6. Internal Helper Methods

These functions ensure **safe access to campaign data**.

---

## `getCampaignLocked()`

Fetches a campaign safely.

```go
func (s *CampaignService) getCampaignLocked(id string)
```

Error cases:

* Campaign does not exist.

---

## `getActiveFightLocked()`

Returns the current active fight.

```go
func (s *CampaignService) getActiveFightLocked(c *Campaign)
```

Error cases:

| Case                | Error                     |
| ------------------- | ------------------------- |
| No active fight     | `"no active fight found"` |
| Fight missing       | `"active fight missing"`  |
| Fight already ended | `"no ongoing fight"`      |

---

# 7. Application Layer Flow

```
External Layer (API / CLI)
           ↓
     CampaignService
           ↓
   Domain Combat Logic
           ↓
     Update Campaign
           ↓
        Return Result
```

---

# 8. Unit Tests

## File: `campaign_service_test.go`

Tests validate **correct campaign behavior**.

---

### `TestStartCampaignCreatesCampaign`

Checks:

* Campaign ID generated
* Campaign stored in store

---

### `TestStartFightWithoutCampaignFails`

Ensures:

* Starting fight on invalid campaign fails.

---

### `TestCannotStartFightIfActiveExists`

Ensures:

* Only **one active fight at a time**.

---

### `TestPerformActionWithoutFightFails`

Ensures:

* Player cannot perform action without fight.

---

### `TestPlayerHPPersistsAfterFight`

Ensures:

* Player HP is preserved after fight.

---

# 9. Application Layer Responsibilities

The Application Layer handles:

✔ Game progression
✔ Campaign lifecycle
✔ Fight lifecycle
✔ State persistence
✔ Coordination between domain logic and storage

---

It **does NOT handle**:

❌ Combat rules
❌ Damage calculation
❌ Creature definitions
❌ HTTP requests or UI

These belong to:

| Layer              | Responsibility |
| ------------------ | -------------- |
| Domain             | Game rules     |
| Application        | Game flow      |
| Infrastructure/API | Communication  |

---

# 10. Example Gameplay Flow

```
Start Campaign
      ↓
Start Fight
      ↓
Player Action
      ↓
Resolve Combat (Domain)
      ↓
Update Campaign State
      ↓
Return Action Result
```

---

If you want, send **Layer 3 (probably the API / transport layer)** and I’ll document it too. Then I can also give you a **clean architecture diagram of the whole backend**, which will make your project documentation look very professional.
