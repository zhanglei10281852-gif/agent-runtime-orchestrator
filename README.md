# Agent Runtime Orchestrator

Agent Runtime Orchestrator is a production-oriented Go backend for controlling how enterprise AI agents invoke MCP servers and other privileged tools. It keeps tool releases, trust boundaries, short-lived approvals, execution capacity, policy evidence, and audit history in one durable workflow without calling an external online service.

## Architecture

- `cmd/server` wires configuration, SQLite, services, HTTP middleware, workers, and graceful shutdown.
- `cmd/seed-user` creates or updates a local account with a bcrypt password hash.
- `internal/domain` owns tool and request state machines, risk ranges, authorization roles, and business errors.
- `internal/service` coordinates catalog, execution authorization, receipts, incident review, bulk intake, queries, and authentication.
- `internal/repository` defines persistence contracts; HTTP handlers never issue SQL.
- `internal/storage/sqlite` provides versioned migrations, transactional writes, optimistic updates, pagination, restart recovery, and durable job claiming.
- `internal/httpapi` exposes authenticated JSON APIs with stable error codes and request IDs.
- `internal/worker` retries durable deliveries, expires pending approvals, records permanent failures, and stops on context cancellation.
- `internal/audit` records actor, request, action, object, outcome, and metadata in the owning transaction.

## Business Model

The versioned SQLite schema contains 14 related tables. The most important relationships are:

- `users` own revocable, expiring `sessions`. Roles are `agent_developer`, `tool_operator`, `security_reviewer`, and `compliance_auditor`.
- `workspaces` define allowed policy-risk ranges, execution duration limits, review deadlines, and a business time zone.
- `trust_zones` represent requester and execution boundaries with local-day capacity and cutoff rules.
- `tool_revisions` register a versioned MCP/tool contract in one workspace and requester zone.
- `execution_pools` provide attested capacity and may be reserved by only one request.
- `execution_requests` bind a workspace, two trust zones, one execution pool, and one or more tool revisions through `execution_request_tools`.
- `approval_tasks` implement expiring two-person review. The requester cannot approve their own task.
- `execution_receipts` carry ordered policy-risk signals. Out-of-range receipts open or extend `policy_incidents` and block affected tool revisions.
- `review_decisions`, `audit_events`, `idempotency_records`, and `outbox_jobs` preserve review evidence, replay protection, and asynchronous delivery.

Submitting an execution request is atomic: it checks the active workspace and trust zones, applies business-day capacity, validates tool expiry and protocol metadata, reserves every tool revision, reserves an attested execution pool, persists the request, stores the idempotent response, enqueues delivery, and writes an audit event. Any intermediate failure rolls the transaction back. Version predicates and unique constraints protect reservations from concurrent writers.

## State Machines

- Workspace: `draft -> active -> archived`.
- Tool revision: `registered -> verified -> reserved -> executing -> executed`, with blocked, approved, and rejected policy outcomes.
- Execution request: `submitted -> authorized -> executing -> completed -> archived`; submitted or authorized requests may be cancelled.
- Execution pool: `available -> reserved -> allocated`, plus reconciliation and retirement.
- Approval task: `pending -> accepted|rejected|expired`.
- Policy incident: `open -> reviewing -> cleared|rejected`.

An execution cannot start until its revisions and pool are reserved and its readiness constraints pass. Cancellation releases both tool revisions and capacity. Archival is rejected while approvals or policy incidents remain unresolved.

## Configuration

Copy `.env.example` values into the process environment. No password or token is committed. Create a local developer account explicitly:

```bash
DATABASE_PATH=./data/agent-runtime.db \
BOOTSTRAP_EMAIL=developer@example.test \
BOOTSTRAP_PASSWORD='change-this-local-password' \
BOOTSTRAP_DISPLAY_NAME='Agent Developer' \
BOOTSTRAP_ROLE=agent_developer \
go run ./cmd/seed-user
```

Start the API with `go run ./cmd/server`. Startup applies unapplied migrations. `GET /healthz` reports process liveness and `GET /readyz` checks the database. SIGINT/SIGTERM cancels workers before the HTTP shutdown deadline.

## HTTP API

Authentication endpoints are `POST /api/v1/auth/login` and authenticated `POST /api/v1/auth/logout`. A successful login returns a bearer token backed by a server-side session; logout revokes it and expired sessions are rejected.

The main workflow uses:

- `POST /api/v1/workspaces` and `POST /api/v1/workspaces/{id}/activate`
- `POST /api/v1/trust-zones` and `POST /api/v1/execution-pools`
- `POST /api/v1/tool-revisions`, bulk registration, and `POST /api/v1/tool-revisions/{id}/verify`
- `POST /api/v1/execution-requests` with an `Idempotency-Key` header
- request `authorize`, `begin`, `complete`, `archive`, and `cancel` operations
- `POST /api/v1/execution-requests/{id}/approval-tasks` and approval resolution
- `POST /api/v1/execution-requests/{id}/receipts` and policy-incident review
- paginated execution, incident, and audit queries plus a platform summary

Errors use one stable response shape:

```json
{"error":{"code":"business_conflict","message":"...","request_id":"req_..."}}
```

## Build And Test

```bash
go test ./... -count=1
go test -race ./... -count=1
go vet ./...
go build ./...
```

Tests use temporary on-disk SQLite databases and cover migrations, rollback, optimistic conflicts, concurrent reservations, restart recovery, pagination, authentication and revocation, HTTP error contracts, lifecycle transitions, idempotency, worker retry/cancellation, and time boundaries.

Build and run the Linux image:

```bash
docker build --platform linux/amd64 -t agent-runtime-orchestrator .
docker run --rm -p 8080:8080 -v agent-runtime-data:/data agent-runtime-orchestrator
```

The image uses the repository's real `cmd/server` entry point and stores its database at `/data/agent-runtime.db` by default.
