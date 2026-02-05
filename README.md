# PIXN Network

PIXN Network is an early-stage execution layer for AI agents, designed for Web3 and infrastructure developers who need a simple, verifiable way to submit and manage AI-related tasks via an API.

This repository represents a working MVP skeleton, not a finished product.

---

## What problem does PIXN solve?

Developers experimenting with AI agents often face the same issues:
- No clear execution boundary between agent logic and infrastructure
- No simple way to submit tasks, track execution, and integrate later with decentralized systems
- Overengineered frameworks before basic execution is proven

PIXN focuses on execution first, not intelligence.

---

## What PIXN is (current stage)

- A minimal HTTP API for task submission
- A clear execution boundary (input → accepted → future processing)
- A foundation for integrating agent logic later
- A demonstrable, runnable MVP

---

## What PIXN is NOT

- Not an AI model
- Not an autonomous agent
- Not a decentralized network
- Not a production-ready system

---

## Current MVP Status

✔ API server runs locally  
✔ /task endpoint available (POST-only)  
✔ Demo execution confirmed on Windows  
✔ CI pipeline active

See docs/demo.md for details.

---

## Repository structure


---

## Roadmap (high-level)

- Define agent execution interface
- Add in-memory task lifecycle
- Introduce pluggable execution backends
- Explore decentralized execution models

No timelines are promised.

---

## Philosophy

Execution before intelligence.  
Working code before narratives.

---

## Disclaimer

This is an early-stage MVP.  
Interfaces and structure may change.
