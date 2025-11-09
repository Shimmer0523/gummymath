# Tasks: Add Rect2D Structure

**Input**: Design documents from `/specs/003-add-rect2d-struct/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md

**Tests**: Unit tests are required by the project constitution and will be created for all new logic.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- This is a single project Go library. New files will be at the repository root.

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Create the necessary files for the new feature.

- [X] T001 [P] Create the source file for the `Rect2D` struct in `rect2d.go`
- [X] T002 [P] Create the test file for the `Rect2D` struct in `rect2d_test.go`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Define the core data structure that all other logic will depend on.

- [X] T003 Define the `Rect2D` struct in `rect2d.go` with `Position` (`Vector2D`), `Width` (`float64`), and `Height` (`float64`) fields, based on `data-model.md`.

---

## Phase 3: User Story 1 - Create and inspect a rectangle (Priority: P1) 🎯 MVP

**Goal**: Enable developers to create a `Rect2D` instance and verify its properties.

**Independent Test**: Create a `Rect2D` with specific values and assert that its public fields match the values provided on creation.

### Implementation for User Story 1

- [X] T004 [US1] Implement the `NewRect2D` constructor function in `rect2d.go`.
- [X] T005 [US1] Write a unit test `TestNewRect2D` in `rect2d_test.go` to verify that the constructor correctly assigns the position, width, and height.

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently.

---

## Phase 4: User Story 2 - Retrieve rectangle corner points (Priority: P1)

**Goal**: Allow developers to get the four corner coordinates of any `Rect2D` instance.

**Independent Test**: For a `Rect2D` with a known position and size, call the `Corners()` method and assert that the returned `Vector2D` points match the calculated expected coordinates.

### Implementation for User Story 2

- [X] T006 [US2] Implement the `Corners() [4]Vector2D` method on the `Rect2D` struct in `rect2d.go`.
- [X] T007 [US2] Write a unit test `TestCorners` in `rect2d_test.go` to verify the corner calculation is correct for a sample rectangle.

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently.

---

## Phase 5: User Story 3 - Modify a rectangle (Priority: P2)

**Goal**: Allow developers to modify the properties of an existing `Rect2D` instance.

**Independent Test**: Create a `Rect2D`, change its public `Position`, `Width`, and `Height` fields, and then assert that the fields hold the new values.

### Implementation for User Story 3

- [X] T008 [US3] Write a unit test `TestRect2DModification` in `rect2d_test.go`. This test will create a `Rect2D`, modify its public fields directly, and assert that the changes were successful. (No implementation is needed in `rect2d.go` as the fields are public).

**Checkpoint**: All user stories should now be independently functional.

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Finalize the feature with documentation and code formatting.

- [X] T009 [P] Add GoDoc comments in Japanese to the `Rect2D` struct and its methods in `rect2d.go` as required by the constitution.
- [X] T010 Run `go fmt ./...` to ensure all code in the project is correctly formatted.
- [X] T011 Run all project tests with `go test ./...` to confirm that the new feature works and hasn't introduced any regressions.

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: Can start immediately.
- **Foundational (Phase 2)**: Depends on `T001` from Setup.
- **User Stories (Phase 3-5)**: Depend on Foundational phase completion.
- **Polish (Phase 6)**: Depends on all user stories being complete.

### User Story Dependencies

- **User Story 1 (P1)**: Depends on `T003`.
- **User Story 2 (P1)**: Depends on `T003`.
- **User Story 3 (P2)**: Depends on `T003`.

All user stories are independent of each other and can be developed in parallel after the Foundational phase is complete.

### Within Each User Story

- Implementation task first, then the corresponding test task.

### Parallel Opportunities

- `T001` and `T002` can run in parallel.
- Once `T003` is complete, work on User Story 1, 2, and 3 can happen in parallel.
- `T009` (documentation) can be done in parallel with other polish tasks.

---

## Implementation Strategy

### MVP First (User Stories 1 & 2)

1. Complete Phase 1 & 2 (Setup & Foundational).
2. Complete Phase 3 (User Story 1).
3. Complete Phase 4 (User Story 2).
4. **STOP and VALIDATE**: The core functionality (creation and inspection) is now complete.
5. Proceed to other stories and polish.

### Incremental Delivery

1. Complete Setup + Foundational → `Rect2D` struct exists.
2. Add User Story 1 → `Rect2D` can be created. Testable.
3. Add User Story 2 → `Rect2D` corners can be calculated. Testable.
4. Add User Story 3 → `Rect2D` modification is verified. Testable.
5. Complete Polish phase.
