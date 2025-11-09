# Feature Specification: Add Rect2D Structure

**Feature Branch**: `003-add-rect2d-struct`
**Created**: 2025-11-09
**Status**: Draft
**Input**: User description: "矩形を扱うRect2D構造体を追加する。Rect2Dは左上の点座標と、幅と高さの取得と設定が可能で、矩形の4点の座標を一度に取得できる関数も作成する。"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Create and inspect a rectangle (Priority: P1)

A developer needs to represent a rectangular area in a 2D space. They create a `Rect2D` instance by specifying its position and size.

**Why this priority**: This is the core functionality of the feature, enabling the basic representation of a rectangle.

**Independent Test**: A `Rect2D` can be created, and its properties (position, width, height) can be retrieved and verified to match the initial values.

**Acceptance Scenarios**:

1.  **Given** a position (e.g., x=10, y=20) and dimensions (e.g., width=100, height=50), **When** a `Rect2D` is created, **Then** its position property returns the correct `Vector2D` (10, 20), its width property returns 100, and its height property returns 50.

---

### User Story 2 - Retrieve rectangle corner points (Priority: P1)

A developer needs to get the coordinates of the four corners of a rectangle for tasks like collision detection, rendering, or geometric calculations.

**Why this priority**: This is a key utility function that makes the `Rect2D` struct immediately useful for common graphical and geometric applications.

**Independent Test**: A method can be called on any `Rect2D` instance, which returns the four corner points. The coordinates of these points can then be verified against expected values.

**Acceptance Scenarios**:

1.  **Given** a `Rect2D` at position (10, 20) with size (100, 50), **When** the corner points are requested, **Then** the returned points are Top-Left (10, 20), Top-Right (110, 20), Bottom-Left (10, 70), and Bottom-Right (110, 70).

---

### User Story 3 - Modify a rectangle (Priority: P2)

A developer needs to move or resize a rectangle after it has been created.

**Why this priority**: Modifying existing objects is a fundamental operation required for any dynamic use case, such as moving UI elements or animating sprites.

**Independent Test**: The position and size of an existing `Rect2D` instance can be updated through its setter methods, and the new values can be verified by calling the corresponding getter methods.

**Acceptance Scenarios**:

1.  **Given** an existing `Rect2D` at (10, 20), **When** its position is updated to (30, 40), **Then** its position property returns the new `Vector2D` (30, 40).
2.  **Given** an existing `Rect2D` with size (100, 50), **When** its size is updated to (120, 60), **Then** its width property returns 120 and its height property returns 60.

---

### Edge Cases

- What happens when width or height are set to zero? The rectangle would have no area.
- How does the system handle negative width or height? This should be considered invalid input. The behavior (e.g., error, clamping to zero) should be defined during implementation.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The system MUST provide a `Rect2D` struct.
- **FR-002**: The `Rect2D` struct MUST store its top-left corner's position using the existing `Vector2D` struct.
- **FR-003**: The `Rect2D` struct MUST store its width and height as floating-point numbers.
- **FR-004**: The `Rect2D` struct MUST provide methods to get and set the top-left corner's position.
- **FR-005**: The `Rect2D` struct MUST provide methods to get and set its width and height.
- **FR-006**: The `Rect2D` struct MUST provide a method to return the coordinates of its four corners (Top-Left, Top-Right, Bottom-Right, Bottom-Left) as an array or slice of `Vector2D` objects.

### Key Entities *(include if feature involves data)*

- **Rect2D**: Represents a 2D rectangle in a Cartesian coordinate system, defined by its top-left corner point (`Vector2D`), a width, and a height.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Developers can successfully instantiate a `Rect2D` object with a given position and size, and retrieve those same values without error.
- **SC-002**: The properties (position, size) of a `Rect2D` object can be successfully modified after instantiation, with the new values being retrieved correctly.
- **SC-003**: The corner calculation method correctly and consistently returns the four corner points for any valid `Rect2D` instance (positive width and height).