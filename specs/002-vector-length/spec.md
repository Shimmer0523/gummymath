# Feature Specification: Get Vector Length

**Feature Branch**: `002-vector-length`
**Created**: 2025-10-31
**Status**: Draft
**Input**: User description: "Vector2Dにベクトルの長さを取得する関数を追加する"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Calculate Vector Magnitude (Priority: P1)

As a developer using the 2D vector library, I want to calculate the length (magnitude) of a vector so that I can use it in physics calculations, distance checks, or other mathematical operations.

**Why this priority**: This is a fundamental and essential feature for any vector library.

**Independent Test**: Can be fully tested by creating a `Vector2D` object and calling the `Length()` function to verify the returned value.

**Acceptance Scenarios**:

1.  **Given** a `Vector2D` object with components (3, 4), **When** the `Length()` function is called, **Then** the return value is 5.
2.  **Given** a `Vector2D` object with components (0, 0), **When** the `Length()` function is called, **Then** the return value is 0.
3.  **Given** a `Vector2D` object with components (-3, -4), **When** the `Length()` function is called, **Then** the return value is 5.

### Edge Cases

- What happens when the vector components are very large or very small? The calculation should maintain floating-point precision.
- What happens with a zero vector (0,0)? The length should be 0.

## Requirements *(mandatory)*

### Functional Requirements

-   **FR-001**: A `Length()` function MUST be added to the `Vector2D` type.
-   **FR-002**: The `Length()` function MUST calculate the vector's magnitude using the formula `sqrt(x*x + y*y)`.
-   **FR-003**: The `Length()` function MUST return a `float64` value.

### Key Entities *(include if feature involves data)*

-   **Vector2D**: Represents a 2D vector with `X` and `Y` components.

## Success Criteria *(mandatory)*

### Measurable Outcomes

-   **SC-001**: The `Length()` function returns the correct magnitude for any given vector, verifiable through manual calculation.
-   **SC-002**: The function's calculation is accurate to the standard precision of `float64`.