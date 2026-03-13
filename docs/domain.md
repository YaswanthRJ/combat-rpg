Here is **clean documentation for your Domain Layer**. This is typically the **core business logic layer** in a layered backend architecture.

---

# Domain Layer Documentation

## 1. Overview

The **Domain Layer** contains the **core game logic and data models** for the combat system.
It defines:

* Game **entities** (Creature, FightState)
* **Actions** a creature can perform
* **Combat mechanics** (damage calculation, turn resolution)
* **Game state transitions** (win/lose conditions)

This layer is **independent of frameworks, databases, or APIs** and purely represents the **rules of the game**.

---

# 2. Action System

## File: `action.go`

### Purpose

Defines all possible **actions** in combat and their properties.

### Types

#### `Actionkind`

Represents the **category of an action**.

```go
type Actionkind int
```

Possible values:

| Value           | Meaning                    |
| --------------- | -------------------------- |
| `ActionAttack`  | Offensive action           |
| `ActionDefend`  | Defensive action           |
| `ActionUtility` | Utility/non-damage actions |

---

#### `Action`

Represents a **specific combat move**.

```go
type Action int
```

Available actions:

| Action        | Description             |
| ------------- | ----------------------- |
| `HeavyAttack` | High damage attack      |
| `FastAttack`  | Low damage quick attack |
| `Block`       | Defensive move          |

---

### `ActionData`

Stores metadata about each action.

```go
type ActionData struct {
    Name       string
    Kind       Actionkind
    Multiplier float64
}
```

Fields:

| Field        | Description         |
| ------------ | ------------------- |
| `Name`       | Human-readable name |
| `Kind`       | Type of action      |
| `Multiplier` | Damage multiplier   |

---

### `ActionPool`

A lookup table mapping actions to their data.

```go
var ActionPool = map[Action]ActionData
```

Example:

| Action      | Multiplier | Kind   |
| ----------- | ---------- | ------ |
| HeavyAttack | 1.3        | Attack |
| FastAttack  | 0.8        | Attack |
| Block       | 0.25       | Defend |

---

# 3. Creature System

## File: `creature.go`

### Purpose

Defines **combat entities** and templates used to generate creatures.

---

### `Creature`

Represents an **instance of a fighter in battle**.

```go
type Creature struct {
    HP      int
    MaxHP   int
    Attack  int
    Defense int
    Actions []Action
}
```

Fields:

| Field   | Description       |
| ------- | ----------------- |
| HP      | Current health    |
| MaxHP   | Maximum health    |
| Attack  | Attack power      |
| Defense | Damage reduction  |
| Actions | Available actions |

---

### `CreatureTemplate`

Represents a **template used to generate creatures**.

```go
type CreatureTemplate struct {
    Name        string
    MaxHP       int
    Attack      int
    Defense     int
    Actions     []Action
    Description string
}
```

---

### `CreaturePool`

A registry of predefined creature templates.

```go
var CreaturePool = map[string]CreatureTemplate
```

Examples:

| Creature | HP | Attack | Defense |
| -------- | -- | ------ | ------- |
| Bandit   | 20 | 15     | 10      |
| Soldier  | 30 | 20     | 25      |

---

### `GenerateCreature()`

Creates a **new creature instance from a template**.

```go
func GenerateCreature(Name string) (Creature, error)
```

Steps:

1. Looks up the creature template from `CreaturePool`
2. Initializes a new `Creature`
3. Copies the action list
4. Returns the generated creature

Example:

```go
enemy, err := GenerateCreature("Bandit")
```

---

# 4. Fight State

## File: `fight_state.go`

### Purpose

Tracks the **current state of a combat encounter**.

---

### `FightStatus`

Represents the **status of a fight**.

```go
type FightStatus int
```

Possible states:

| Status       | Meaning               |
| ------------ | --------------------- |
| `Ongoing`    | Fight is active       |
| `PlayerWon`  | Player defeated enemy |
| `PlayerLost` | Player was defeated   |

---

### `FightState`

Stores the **entire combat state**.

```go
type FightState struct {
    Player       Creature
    Enemy        Creature
    FightStatus  FightStatus
    ActionNumber int
}
```

Fields:

| Field        | Description          |
| ------------ | -------------------- |
| Player       | Player creature      |
| Enemy        | Enemy creature       |
| FightStatus  | Current fight result |
| ActionNumber | Turn counter         |

---

# 5. Action Result

## File: `action_result.go`

### Purpose

Represents the **result of a combat turn**.

```go
type ActionResult struct {
    PlayerDamageDealt int
    EnemyDamageDealt  int
    ActionNumber      int
    FightEnded        bool
}
```

Fields:

| Field             | Description            |
| ----------------- | ---------------------- |
| PlayerDamageDealt | Damage dealt to enemy  |
| EnemyDamageDealt  | Damage dealt to player |
| ActionNumber      | Current turn number    |
| FightEnded        | Whether combat ended   |

---

# 6. Combat Logic

## File: `combat.go`

### Purpose

Contains the **core combat engine**.

---

## `ResolveRound()`

Main function that processes **one combat turn**.

```go
func ResolveRound(state *FightState, action Action) (ActionResult, error)
```

Execution flow:

1. Increase turn counter
2. Resolve **player action**
3. Apply damage to enemy
4. Check if enemy died
5. Resolve **enemy attack**
6. Apply damage to player
7. Check if player died
8. Return result

---

### Combat Flow

```
Player Action
     ↓
Damage to Enemy
     ↓
Enemy Dead? → Player Won
     ↓
Enemy Attack
     ↓
Damage to Player
     ↓
Player Dead? → Player Lost
```

---

## `ResolvePlayerAction()`

Processes the **player's selected action**.

```go
func ResolvePlayerAction(state *FightState, action Action)
```

Logic:

* Fetch action from `ActionPool`
* If action is **defense → activate block**
* If action is **attack → calculate damage**

---

## `ResolveEnemyAction()`

Handles **enemy attack phase**.

```go
func ResolveEnemyAction(state *FightState, playerBlocked bool)
```

Features:

* Enemy always uses **base multiplier (1.0)**
* If player blocked → damage reduced to **25%**

---

## `calculateDamage()`

Core damage formula.

```go
damage = attack * multiplier - defense
```

Rules:

* Damage cannot be negative
* Returns integer value

---

## `applyDamage()`

Applies damage to a creature.

```go
func applyDamage(creature *Creature, damage int) bool
```

Returns:

```
true → creature died
false → still alive
```

---

## `applyMultiplier()`

Applies percentage-based reduction.

Example:

```
Block → 0.25 multiplier
```

---

# 7. Unit Tests

## File: `combat_test.go`

### Purpose

Tests the correctness of the **combat system**.

---

### `TestHeavyAttackDealsDamage`

Verifies:

* Heavy attack damage calculation
* Enemy HP reduction

---

### `TestPlayerWin`

Ensures:

* Player victory detected
* Fight ends correctly

---

### `TestPlayerLosesFight`

Ensures:

* Enemy can defeat player
* Fight status updates properly

---

### `TestInvalidAction`

Ensures:

* Invalid actions produce an error

---

# 8. Domain Layer Responsibilities

The domain layer handles:

✔ Core game rules
✔ Combat logic
✔ Entities and data models
✔ State transitions
✔ Game balance parameters

It **does NOT handle**:

❌ HTTP requests
❌ Databases
❌ External services
❌ UI

---

# 9. Example Fight Flow

```
Start Fight
   ↓
Create FightState
   ↓
Player selects action
   ↓
ResolveRound()
   ↓
Apply Player Damage
   ↓
Apply Enemy Damage
   ↓
Update FightState
   ↓
Return ActionResult
```

---

If you want, when you send **Layer 2**, I can also help you produce a **full architecture diagram for the whole backend** (which will make your project documentation much stronger).
