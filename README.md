# pdd — Puzzle Driven Development Utility

TODO/FIXME comments in source code into stable, trackable “puzzles”.  
It follows the core idea of Puzzle Driven Development: break work into small, explicit tasks stored directly in the codebase.

A comment like:

```
TODO: [1772488711992532061] improve error handling
```

is automatically transformed into:

```
TODO: [1772487140120457593] improve error handling
```


Each puzzle is stored in `.puzzles.xml`, which acts as:
- a cache of all open puzzles,
- a historical log of closed puzzles.

The tool is fully local, deterministic, and dependency‑free.

---

## Features

### `pdd scan`
- Recursively scans the current directory.
TODO/FIXME comments in text files.
- Inserts a timestamp if missing.
- Updates `.puzzles.xml` with new or modified puzzles.
- Marks puzzles as `CLOSED` if they disappear from the code.

### `pdd list`
- Prints all open puzzles in a compact, readable format:

```
TODO: [1772486223922468279] path/to/file — description
```

---

## Puzzle Format

Puzzles are stored in `.puzzles.xml`:

```xml
<puzzles>
<puzzle>
  <type>TODO</type>
  <timestamp>20260302204400</timestamp>
  <description>Refactor this method</description>
  <file>src/main/Parser.js</file>
  <status>OPEN</status>
</puzzle>
</puzzles>
```

### Fields:

TODO: [1772488261058242729] or FIXME
- `timestamp` — unique puzzle identifier
- `description` — puzzle text
- `file` — file path
- `status` — OPEN or CLOSED

## Design Principles

### Minimal state  

The source code is the single source of truth. XML is only a cache and history.

### Deterministic behavior

Same input → same output. No heuristics or guessing.

### Robustness  

Puzzles survive file renames, moves, and line shifts.

### No magic  

If a timestamp is removed manually, the tool treats it as a new puzzle.

### Zero external dependencies  

One static binary, no databases, no network calls.
