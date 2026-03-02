# AGENTS — Development Rules and Collaboration Model for `pdd`

This document defines the roles, responsibilities, and development rules for contributors.  
The project follows principles of clarity, determinism, and minimalism.

---

## Roles

### Maintainer
- Owns the architecture and long‑term direction.
- Ensures simplicity and consistency across the codebase.
- Reviews contributions and enforces project rules.
- Maintains the integrity of `.puzzles.xml` and puzzle formats.

### Contributor
- Implements small, isolated changes.
- Respects the project structure and design principles.
- Avoids unnecessary complexity.
- Writes clean, readable, deterministic code.

---

## Development Principles

### Small, incremental changes
Every change must be small, focused, and easy to review.

### Transparency
All communication happens through:
- code,
- commit messages,
- puzzles.

No hidden context or undocumented behavior.

### Minimal state
The project stores only:
- source files,
- `.puzzles.xml`.

No databases, caches, or external services.

### Predictable behavior
The tool must behave deterministically:
- no guessing,
- no fuzzy matching,
- no hidden side effects.

### Resilience to editing
Developers may freely:
- move files,
- reorder code,
- delete timestamps.

The tool must handle these cases gracefully:
- missing timestamp → new puzzle,
- removed comment → CLOSED puzzle.

---

## Puzzle Rules

### Comment format

```
TODO: [1772488261057134868] description
FIXME: [1772488261057145147] description
```

After scanning:

```
TODO: [1772488261057185904] [timestamp] description
```

### Editing descriptions
Changing the text does not create a new puzzle.

### Removing timestamps
This is treated as creating a new puzzle.  
The old one becomes `CLOSED`.

### Removing comments
The puzzle is marked as `CLOSED`.

---

## Commit Rules

- One logical change per commit.
- Commit messages describe *what* changed, not *why*.
- `.puzzles.xml` must be committed if it changed as a result of code edits.
- No large, multi‑purpose commits.

---

## Code Style

- Short, focused functions.
- No global state.
- Minimal dependencies.
- Explicit data structures.
- No hidden behavior or clever tricks.

---

## Project Evolution

The project evolves incrementally.  
New features are added only if they:
- do not complicate the architecture,
- preserve determinism,
- require no external services,
- keep cognitive load low.

If a feature introduces doubt, it is not added.

---

## Summary

`pdd` must remain:
- simple,
- robust,
- predictable,
- minimalistic.

All decisions favor clarity and long‑term maintainability.
