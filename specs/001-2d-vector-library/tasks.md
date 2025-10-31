---
description: "Task list for 2D Vector Calculation Library"
---

# Tasks: 2D Vector Calculation Library

**Input**: Design documents from `/specs/001-2d-vector-library/`
**Prerequisites**: plan.md (required), spec.md (required for user stories)

**Tests**: Unit tests are included for each user story.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- **Single project**: `src/`, `tests/` at repository root

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [X] T001 Create project directory `src/vector2`
- [X] T002 Create test directory `tests/vector2`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [X] T003 Create `Vector2D` struct and `New` constructor in `src/vector2/vector2.go`
- [X] T004 Create initial test file `tests/vector2/vector2_test.go`

**Checkpoint**: Foundation ready - user story implementation can now begin.

---

## Phase 3: User Story 1 - Basic Vector Arithmetic (Priority: P1) 🎯 MVP

**Goal**: As a developer, I want to perform basic arithmetic operations on 2D vectors, such as addition, subtraction, and scaling.

**Independent Test**: The arithmetic operations can be tested by creating simple vectors, performing the operations, and asserting that the resulting vector has the expected components.

### Implementation for User Story 1

- [X] T005 [US1] Implement `Add(v *Vector2D) *Vector2D` method in `src/vector2/vector2.go`
- [X] T006 [US1] Implement `Sub(v *Vector2D) *Vector2D` method in `src/vector2/vector2.go`
- [X] T007 [US1] Implement `Scale(s float64) *Vector2D` method in `src/vector2/vector2.go`
- [X] T008 [US1] Write unit tests for `Add`, `Sub`, and `Scale` in `tests/vector2/vector2_test.go`

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently.

---

## Phase 4: User Story 2 - Vector Products (Priority: P2)

**Goal**: As a developer, I want to calculate the dot product and cross product of 2D vectors.

**Independent Test**: The product operations can be tested by creating vectors, performing the calculations, and asserting that the result is the expected scalar (dot product) or float (cross product).

### Implementation for User Story 2

- [X] T009 [US2] Implement `Dot(v *Vector2D) float64` method in `src/vector2/vector2.go`
- [X] T010 [US2] Implement `Cross(v *Vector2D) float64` method in `src/vector2/vector2.go`
- [X] T011 [US2] Write unit tests for `Dot` and `Cross` in `tests/vector2/vector2_test.go`

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently.

---

## Phase 5: User Story 3 - Vector Transformations (Priority: P3)

**Goal**: As a developer, I want to rotate a 2D vector by a given angle.

**Independent Test**: The rotation can be tested by rotating a vector by a known angle (e.g., 90 degrees) and asserting that the new vector components are correct.

### Implementation for User Story 3

- [X] T012 [US3] Implement `Rotate(angle float64) *Vector2D` method in `src/vector2/vector2.go`
- [X] T013 [US3] Write unit test for `Rotate` in `tests/vector2/vector2_test.go`

**Checkpoint**: All user stories should now be independently functional.

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [X] T014 Add GoDoc documentation to all public functions in `src/vector2/vector2.go`
- [X] T015 Validate code against `quickstart.md` usage examples.

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately.
- **Foundational (Phase 2)**: Depends on Setup completion.
- **User Stories (Phase 3+)**: All depend on Foundational phase completion.
- **Polish (Final Phase)**: Depends on all desired user stories being complete.

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (Phase 2).
- **User Story 2 (P2)**: Can start after Foundational (Phase 2).
- **User Story 3 (P3)**: Can start after Foundational (Phase 2).

### Parallel Opportunities

- User stories can be implemented in parallel after the Foundational phase is complete.

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational
3. Complete Phase 3: User Story 1
4. **STOP and VALIDATE**: Test User Story 1 independently.

### Incremental Delivery

1. Complete Setup + Foundational.
2. Add User Story 1 → Test independently.
3. Add User Story 2 → Test independently.
4. Add User Story 3 → Test independently.
