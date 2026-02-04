# PIXN Demo

This document demonstrates the current MVP concept of PIXN.

## What this demo shows

- How a task would be submitted to the system
- How the system accepts and processes the task
- What kind of response is expected

This demo represents the **current state**, not future promises.

---

## Demo Scenario

A client sends a task request to the PIXN API.

### Example request

```bash
curl -X POST http://localhost:8080/task \
  -H "Content-Type: application/json" \
  -d '{
    "id": "task-001",
    "prompt": "Generate a short explanation of PIXN"
  }'
{
  "status": "accepted",
  "task_id": "task-001"
}
