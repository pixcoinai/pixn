# Technical Next Steps

This document outlines the next concrete technical steps for PIXN,
based strictly on the current Pre-MVP state.

No speculative features or timelines are included.

---

## Current Baseline

What is already implemented:
- Runnable HTTP API
- Task submission endpoint
- Clear execution boundary
- CI pipeline
- Verified local demo

---

## Next Step 1 — Define Agent Interface

Goal:
- Formalize how accepted tasks will be executed in the future

Actions:
- Define a minimal agent execution interface
- Specify input/output contracts
- Keep implementation decoupled from API layer

No execution logic is added at this step.

---

## Next Step 2 — In-Memory Task Lifecycle

Goal:
- Prove end-to-end task handling without persistence

Actions:
- Store tasks in memory
- Track basic task states (accepted → processed)
- Expose minimal task status endpoint

This step validates execution flow only.

---

## Next Step 3 — Execution Backend Abstraction

Goal:
- Avoid tight coupling between API and execution logic

Actions:
- Introduce an execution handler abstraction
- Allow swapping execution backends later
- Keep default implementation minimal

---

## Explicit Non-Goals

The following are intentionally excluded:
- AI model integration
- Autonomous agent behavior
- Persistent storage
- Distributed or decentralized execution

---

## Guiding Principle

Each step must:
- Be runnable
- Be testable
- Reduce uncertainty

Execution precedes intelligence.
