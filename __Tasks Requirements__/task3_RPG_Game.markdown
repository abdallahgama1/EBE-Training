# Task: Design and Implement a Character and Equipment System for a Text-Based RPG in Go

## Overview
You are tasked with designing and implementing a simple character and equipment system for a text-based role-playing game (RPG) using Object-Oriented Programming (OOP) principles in Go. The system should support different character Structs with unique attributes and allow characters to equip weapons and armor that modify their stats. The focus is on creating a modular, extensible OOP design with clear requirements.

## Requirements

### 1. Character Structs
- Create a base `Character` struct with attributes:
  - `name` (string)
  - `health` (int)
  - `strength` (int)
  - `defense` (int)
  - `magic` (int)
- Implement at least three types of characters:
  - `Warrior`: High `strength` and `defense`, low `magic`
  - `Mage`: High `magic`, low `strength`
  - `Archer`: Moderate `strength` and `health`, low `magic`
- Use composition or interfaces to model character types (Go does not support classical inheritance).

### 2. Item Hierarchy
- Create a base `Item` struct with basic properties (e.g., `name`).
- Implement at least two types of items:
  - `Weapon`: Includes a `damageBonus` property (int) that boosts `strength`.
  - `Armor`: Includes a `defenseBonus` property (int) that boosts `defense`.
- Use composition or interfaces to manage different item types.

### 3. Equipping Items
- Characters must be able to equip:
  - One `Weapon`
  - One `Armor`
- Equipping an item modifies the character's stats (e.g., `strength` increases by `damageBonus` from a weapon).

### 4. Stat Calculation
- Add a method to the `Character` struct (e.g., `GetTotalStats`) that calculates and returns the character's effective stats, combining base stats with bonuses from equipped items.

### 5. Equipment Restrictions
- Enforce class-specific equipment rules:
  - `Warrior`: Can equip heavy swords but not magic staves.
  - `Mage`: Can equip staves but not heavy armor.
  - `Archer`: Can equip bows but not heavy shields.
- Implement logic to prevent invalid equipment assignments.

### 6. Text-Based Interface
- Provide a simple command-line interface that allows the user to:
  - Create a character and select their class.
  - Equip a weapon and armor from a predefined list of items.
  - Display the character's current stats and equipped items.



## Example Usage
To illustrate how the system should function, the following example walks through the process of creating a character, equipping items, and displaying stats, using a Warrior named "Aragorn."

### Step 1: Character Creation
- The user selects "Warrior" and names the character "Aragorn."
- The system creates Aragorn with the following base stats:
  - **Health**: 100
  - **Strength**: 15
  - **Defense**: 10
  - **Magic**: 5

### Step 2: Item Selection and Equipping
- The user is presented with a list of predefined items:
  - **Weapons**:
    - "Heavy Sword" (Damage Bonus: 5)
    - "Magic Staff" (Damage Bonus: 3)
    - "Bow" (Damage Bonus: 4)
  - **Armor**:
    - "Plate Armor" (Defense Bonus: 5)
    - "Leather Armor" (Defense Bonus: 3)
    - "Magic Robe" (Defense Bonus: 2)
- The user chooses "Heavy Sword" and "Plate Armor."
- The system checks the class restrictions:
  - Warriors can equip "Heavy Sword" and "Plate Armor" (valid choices).
- The items are successfully equipped.

### Step 3: Stat Modification
- Equipping "Heavy Sword" adds its `damageBonus` (5) to Aragorn's `strength`:
  - Base Strength: 15 + 5 = **Total Strength: 20**
- Equipping "Plate Armor" adds its `defenseBonus` (5) to Aragorn's `defense`:
  - Base Defense: 10 + 5 = **Total Defense: 15**
- Aragorn's updated stats are now:
  - **Health**: 100
  - **Strength**: 20
  - **Defense**: 15
  - **Magic**: 5

### Step 4: Displaying Stats
- The user selects the option to display Aragorn's stats.
- The system shows the following:
  ```
  Character: Aragorn (Warrior)
  Health: 100
  Strength: 20
  Defense: 15
  Magic: 5
  Equipped Weapon: Heavy Sword (+5 damage)
  Equipped Armor: Plate Armor (+5 defense)
  ```

### Step 5: Enforcing Equipment Restrictions
- The user attempts to equip "Magic Staff" to Aragorn.
- The system checks the restriction:
  - Warriors cannot equip magic staves.
- It displays an error message: "Warriors cannot equip magic staves."
- Aragorn's weapon remains "Heavy Sword."

This example demonstrates how the system handles character creation, item equipping, stat modification, and restriction enforcement, ensuring the requirements are met in a clear and interactive way.


## Guidelines
- **Composition**: Use composition to build character and item structures (no inheritance in Go).
- **Interfaces**: Define interfaces for behaviors that vary between character types or items (e.g., equipping items, calculating stats).
- **Encapsulation**: Protect character stats and equipment within the `Character` struct, using methods to access or modify them.
- **Modularity**: Design the system so new character types or item types can be added with minimal changes.

## Deliverables
- **Design Document**: A brief outline or diagram showing the struct relationships and interfaces.
- **Implementation**: Code in Go, following OOP principles as much as possible.
- **Working Interface**: A functional text-based program demonstrating the requirements.
