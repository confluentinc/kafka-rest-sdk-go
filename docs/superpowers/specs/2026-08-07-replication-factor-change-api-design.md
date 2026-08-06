# Replication Factor Change API — Design

## Goal

Add the Replication Factor Change API to `kafka-rest-sdk-go`: list tracked changes, submit a
change, cancel a change. This mirrors the three new endpoints added to `ce-kafka-rest`'s OpenAPI
spec in commits `4c9fdb63` (Confluent Cloud `openapi.yaml`) and `c33b2ee3` (on-prem
`consolidated-openapi.yaml`), tagged `ReplicationFactorChange (v3)`:

- `GET /clusters/{cluster_id}/replication-factor-changes` — list tracked changes
- `POST /clusters/{cluster_id}/replication-factor-changes:create` — submit a change (async)
- `POST /clusters/{cluster_id}/replication-factor-changes:cancel` — cancel an in-flight change

## Reference precedent

- PR #44 (`UpdatePartitionCountKafkaTopic`) — added a method to an *existing* API interface;
  established the "regen → scope down → remove unrelated files → restore go.mod" workflow.
- PR #46 (Share Group V3 API) — closer analog: added a *brand-new* multi-endpoint API surface.
  Its file list is the template for what this change touches:
  `client.go` (+3 lines: struct field + constructor line), `README.md` (+endpoint rows +model
  rows), `kafkarestv3/api/openapi.yaml` (embedded spec copy, scoped additions), a new
  `api_<name>_v3.go`, and one `model_*.go` + `docs/*.md` pair per schema. No mock files in either
  reference PR.

## Starting state

A previous (uncommitted, pre-context-clear) `make gen` run left correctly-generated code for this
feature already sitting in the working tree, alongside ~90 unrelated files from spec drift
(Share Group offsets, Streams Group, Broker Task V3, Replica Status V3, etc. — features added to
`consolidated-openapi.yaml` since this SDK was last regenerated, unrelated to this task) and a
`go.mod` toolchain-version bump. Plan: keep only the replication-factor-change files, discard the
rest, and hand-apply the wiring PR #46 shows `make gen` normally produces for a new API surface,
scoped to just this feature.

## Changes

**Keep as-is (already generated correctly):**
- `kafkarestv3/api_replication_factor_change_v3.go` — `ReplicationFactorChangeV3Api` interface +
  service (`ClustersClusterIdReplicationFactorChangesGet`,
  `ClustersClusterIdReplicationFactorChangescancelPost`,
  `ClustersClusterIdReplicationFactorChangescreatePost`)
- 8 model files: `model_change_replication_factor_request_data.go`,
  `model_change_replication_factor_batch_request_data.go`,
  `model_cancel_replication_factor_change_request_data.go`,
  `model_replication_factor_change_data.go` (+ `_all_of`, `_list`, `_list_all_of`),
  `model_replication_factor_change_cancel_data.go` (+ `_all_of`, `_list`, `_list_all_of`).
  `status` / `prior_status` are plain `string` fields — consistent with how other
  `x-extensible-enum` statuses are already generated elsewhere in this SDK (e.g.
  `RemoveBrokerTaskData`'s status fields); no separate enum type is expected.
- 12 matching `kafkarestv3/docs/*.md` files.

**Hand-edit (scoped to this feature only):**
1. `kafkarestv3/client.go` — add `ReplicationFactorChangeV3Api ReplicationFactorChangeV3Api`
   field to the `APIClient` struct (alphabetically after `ReplicaV3Api`) and
   `c.ReplicationFactorChangeV3Api = (*ReplicationFactorChangeV3ApiService)(&c.common)` in
   `NewAPIClient` (same alphabetical position).
2. `kafkarestv3/README.md` — add 3 endpoint rows to the "Documentation for API Endpoints" table
   and 12 model rows to the "Documentation for Models" list, both alphabetically positioned
   after the `ReplicaV3Api`/`Relationship`-area entries.
3. `kafkarestv3/api/openapi.yaml` (embedded spec copy) — graft in the same paths/schemas/responses
   from ce-kafka's `c33b2ee3` (on-prem `consolidated-openapi.yaml`, the Makefile's codegen
   source), inserted immediately after the existing
   `/clusters/{cluster_id}/remove-broker-tasks/{broker_id}` block, matching its relative position
   in `consolidated-openapi.yaml`.

**Discard:** all Share Group offsets / Streams Group / Broker Task V3 / Replica Status V3 /
Broker Replica Exclusion V3 / etc. untracked files unrelated to this feature, the `go.mod`
toolchain bump (restore via `git checkout`), and `.idea/`.

**Explicitly out of scope:** mock files (neither reference PR added them for a new API).

## Verification

`go build ./...` and `go vet ./...` inside `kafkarestv3/`, matching the testing section of both
reference PRs.

## Delivery

Work happens on branch `kl-4882-rf-change-rest-api` in `kafka-rest-sdk-go`. Commit locally; do
**not** push or open a PR — the user will review and push themselves.
