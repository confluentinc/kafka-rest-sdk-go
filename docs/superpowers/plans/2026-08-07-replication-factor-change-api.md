# Replication Factor Change API Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add the Replication Factor Change API (list / create / cancel) to `kafka-rest-sdk-go`, on branch `kl-4882-rf-change-rest-api`, committed locally only (no push, no PR).

**Architecture:** A stale, uncommitted `make gen` run already produced correct Go code for this feature (interface/service + 8 models + 12 docs) alongside ~90 unrelated files from unrelated spec drift. This plan discards the unrelated files, restores `go.mod`, and hand-writes the three "wiring" files a full codegen run would otherwise touch (`client.go`, `README.md`, the embedded `kafkarestv3/api/openapi.yaml`), scoped to just this feature. All YAML added to `api/openapi.yaml` follows conventions verified against existing structurally-identical entries already in that file (e.g. `AlterBrokerReplicaExclusionData`/`RemoveBrokerTaskData`): `paths`/`components.schemas`/`components.requestBodies`/`components.responses` are insertion-ordered maps (new entries go where the source spec places them), `allOf`-composed schemas split into a `<Name>` wrapper (just an `allOf` list) plus a `<Name>_allOf` fragment (`properties`+`required`, no `type`/`example`), every object's own keys are sorted alphabetically except `properties`/`example` (which preserve declaration order) and `required` lists (which are alphabetized), and shared 400/401/429/5XX response bodies are inlined verbatim from `ce-kafka-rest`'s current `BadRequestErrorResponse`/`UnauthorizedErrorResponse`/`TooManyRequestsErrorResponse`/`ServerErrorResponse` components.

**Tech Stack:** Go (kafkarestv3 SDK package), OpenAPI 3 YAML.

## Global Constraints

- Do not push or open a PR — all work stays local on branch `kl-4882-rf-change-rest-api`.
- Do not add mock files (neither reference PR #44 nor #46 added them for a new API).
- Do not touch any files unrelated to replication-factor-change (Share Group offsets, Streams Group, Broker Task V3, Replica Status V3, Broker Replica Exclusion V3, etc.) — these came from unrelated spec drift and must be discarded, not committed.
- `kafkarestv3/go.mod` must end up identical to `git show HEAD:kafkarestv3/go.mod` (no toolchain/indirect-dependency bump).

---

### Task 1: Discard unrelated generated files and restore go.mod

**Files:**
- Delete: `.idea/` (directory)
- Delete: `kafkarestv3/api_broker_replica_exclusion_v3.go`, `kafkarestv3/api_broker_task_v3.go`, `kafkarestv3/api_remove_broker_task_v3.go`, `kafkarestv3/api_replica_status_v3.go`, `kafkarestv3/api_streams_group_v3.go`
- Delete: `kafkarestv3/docs/AlterShareGroupOffsetsRequestData.md`, `kafkarestv3/docs/AlterShareGroupOffsetsResponseData.md`, `kafkarestv3/docs/AlterShareGroupOffsetsResponseDataPartitions.md`, `kafkarestv3/docs/BrokerReplicaExclusionV3Api.md`, `kafkarestv3/docs/BrokerTaskV3Api.md`, `kafkarestv3/docs/CreateAclRequestDataList.md`, `kafkarestv3/docs/CreateAclRequestDataListAllOf.md`, `kafkarestv3/docs/GroupConfigData.md`, `kafkarestv3/docs/GroupConfigDataAllOf.md`, `kafkarestv3/docs/GroupConfigDataList.md`, `kafkarestv3/docs/GroupConfigDataListAllOf.md`, `kafkarestv3/docs/InlineResponse409.md`, `kafkarestv3/docs/LinkCategory.md`, `kafkarestv3/docs/PartitionLevelTruncationData.md`, `kafkarestv3/docs/RemoveBrokerTaskV3Api.md`, `kafkarestv3/docs/ReplicaStatusV3Api.md`, `kafkarestv3/docs/ResetType.md`, `kafkarestv3/docs/StreamsGroupData.md`, `kafkarestv3/docs/StreamsGroupDataAllOf.md`, `kafkarestv3/docs/StreamsGroupDataList.md`, `kafkarestv3/docs/StreamsGroupDataListAllOf.md`, `kafkarestv3/docs/StreamsGroupMemberAssignmentData.md`, `kafkarestv3/docs/StreamsGroupMemberAssignmentDataAllOf.md`, `kafkarestv3/docs/StreamsGroupMemberData.md`, `kafkarestv3/docs/StreamsGroupMemberDataAllOf.md`, `kafkarestv3/docs/StreamsGroupMemberDataList.md`, `kafkarestv3/docs/StreamsGroupMemberDataListAllOf.md`, `kafkarestv3/docs/StreamsGroupSubtopologyData.md`, `kafkarestv3/docs/StreamsGroupSubtopologyDataAllOf.md`, `kafkarestv3/docs/StreamsGroupSubtopologyDataList.md`, `kafkarestv3/docs/StreamsGroupSubtopologyDataListAllOf.md`, `kafkarestv3/docs/StreamsGroupV3Api.md`, `kafkarestv3/docs/StreamsTaskData.md`, `kafkarestv3/docs/StreamsTaskDataAllOf.md`, `kafkarestv3/docs/StreamsTaskDataList.md`, `kafkarestv3/docs/StreamsTaskDataListAllOf.md`, `kafkarestv3/docs/UpdateGroupConfigRequestData.md`
- Delete: `kafkarestv3/model_alter_share_group_offsets_request_data.go`, `kafkarestv3/model_alter_share_group_offsets_response_data.go`, `kafkarestv3/model_alter_share_group_offsets_response_data_partitions.go`, `kafkarestv3/model_create_acl_request_data_list.go`, `kafkarestv3/model_create_acl_request_data_list_all_of.go`, `kafkarestv3/model_group_config_data.go`, `kafkarestv3/model_group_config_data_all_of.go`, `kafkarestv3/model_group_config_data_list.go`, `kafkarestv3/model_group_config_data_list_all_of.go`, `kafkarestv3/model_inline_response_409.go`, `kafkarestv3/model_link_category.go`, `kafkarestv3/model_partition_level_truncation_data.go`, `kafkarestv3/model_reset_type.go`, `kafkarestv3/model_streams_group_data.go`, `kafkarestv3/model_streams_group_data_all_of.go`, `kafkarestv3/model_streams_group_data_list.go`, `kafkarestv3/model_streams_group_data_list_all_of.go`, `kafkarestv3/model_streams_group_member_assignment_data.go`, `kafkarestv3/model_streams_group_member_assignment_data_all_of.go`, `kafkarestv3/model_streams_group_member_data.go`, `kafkarestv3/model_streams_group_member_data_all_of.go`, `kafkarestv3/model_streams_group_member_data_list.go`, `kafkarestv3/model_streams_group_member_data_list_all_of.go`, `kafkarestv3/model_streams_group_subtopology_data.go`, `kafkarestv3/model_streams_group_subtopology_data_all_of.go`, `kafkarestv3/model_streams_group_subtopology_data_list.go`, `kafkarestv3/model_streams_group_subtopology_data_list_all_of.go`, `kafkarestv3/model_streams_task_data.go`, `kafkarestv3/model_streams_task_data_all_of.go`, `kafkarestv3/model_streams_task_data_list.go`, `kafkarestv3/model_streams_task_data_list_all_of.go`, `kafkarestv3/model_update_group_config_request_data.go`
- Modify (restore): `kafkarestv3/go.mod`

**Interfaces:** None — pure cleanup task, no code contracts produced or consumed.

- [ ] **Step 1: Confirm current untracked/modified state matches expectations**

Run: `git -C /Users/princepurohit/src/kafka-rest-sdk-go status --porcelain=v1`
Expected: `M kafkarestv3/go.mod` plus the untracked files listed above (both the ones to delete and the 12 replication-factor-change files to keep: `api_replication_factor_change_v3.go`, the 8 `model_*replication_factor*`/`model_cancel_replication_factor*`/`model_canceled_partition_data.go`/`model_change_replication_factor*` files, and the 12 matching `docs/*.md` files).

- [ ] **Step 2: Delete the unrelated directory and files**

```bash
rm -rf /Users/princepurohit/src/kafka-rest-sdk-go/.idea
rm /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/api_broker_replica_exclusion_v3.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/api_broker_task_v3.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/api_remove_broker_task_v3.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/api_replica_status_v3.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/api_streams_group_v3.go
rm /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/AlterShareGroupOffsetsRequestData.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/AlterShareGroupOffsetsResponseData.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/AlterShareGroupOffsetsResponseDataPartitions.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/BrokerReplicaExclusionV3Api.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/BrokerTaskV3Api.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/CreateAclRequestDataList.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/CreateAclRequestDataListAllOf.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/GroupConfigData.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/GroupConfigDataAllOf.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/GroupConfigDataList.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/GroupConfigDataListAllOf.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/InlineResponse409.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/LinkCategory.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/PartitionLevelTruncationData.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/RemoveBrokerTaskV3Api.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/ReplicaStatusV3Api.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/ResetType.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupData.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupDataAllOf.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupDataList.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupDataListAllOf.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupMemberAssignmentData.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupMemberAssignmentDataAllOf.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupMemberData.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupMemberDataAllOf.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupMemberDataList.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupMemberDataListAllOf.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupSubtopologyData.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupSubtopologyDataAllOf.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupSubtopologyDataList.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupSubtopologyDataListAllOf.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsGroupV3Api.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsTaskData.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsTaskDataAllOf.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsTaskDataList.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/StreamsTaskDataListAllOf.md \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/docs/UpdateGroupConfigRequestData.md
rm /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_alter_share_group_offsets_request_data.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_alter_share_group_offsets_response_data.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_alter_share_group_offsets_response_data_partitions.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_create_acl_request_data_list.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_create_acl_request_data_list_all_of.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_group_config_data.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_group_config_data_all_of.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_group_config_data_list.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_group_config_data_list_all_of.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_inline_response_409.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_link_category.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_partition_level_truncation_data.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_reset_type.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_group_data.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_group_data_all_of.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_group_data_list.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_group_data_list_all_of.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_group_member_assignment_data.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_group_member_assignment_data_all_of.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_group_member_data.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_group_member_data_all_of.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_group_member_data_list.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_group_member_data_list_all_of.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_group_subtopology_data.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_group_subtopology_data_all_of.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_group_subtopology_data_list.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_group_subtopology_data_list_all_of.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_task_data.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_task_data_all_of.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_task_data_list.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_streams_task_data_list_all_of.go \
   /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/model_update_group_config_request_data.go
```

- [ ] **Step 3: Restore go.mod**

Run: `git -C /Users/princepurohit/src/kafka-rest-sdk-go checkout -- kafkarestv3/go.mod`

- [ ] **Step 4: Verify only the replication-factor-change files remain untracked**

Run: `git -C /Users/princepurohit/src/kafka-rest-sdk-go status --porcelain=v1`
Expected: exactly these 21 untracked entries (no `M` lines):
```
?? kafkarestv3/api_replication_factor_change_v3.go
?? kafkarestv3/docs/CancelReplicationFactorChangeRequestData.md
?? kafkarestv3/docs/CanceledPartitionData.md
?? kafkarestv3/docs/ChangeReplicationFactorBatchRequestData.md
?? kafkarestv3/docs/ChangeReplicationFactorRequestData.md
?? kafkarestv3/docs/ReplicationFactorChangeCancelData.md
?? kafkarestv3/docs/ReplicationFactorChangeCancelDataAllOf.md
?? kafkarestv3/docs/ReplicationFactorChangeCancelDataList.md
?? kafkarestv3/docs/ReplicationFactorChangeCancelDataListAllOf.md
?? kafkarestv3/docs/ReplicationFactorChangeData.md
?? kafkarestv3/docs/ReplicationFactorChangeDataAllOf.md
?? kafkarestv3/docs/ReplicationFactorChangeDataList.md
?? kafkarestv3/docs/ReplicationFactorChangeDataListAllOf.md
?? kafkarestv3/docs/ReplicationFactorChangeV3Api.md
?? kafkarestv3/model_cancel_replication_factor_change_request_data.go
?? kafkarestv3/model_change_replication_factor_batch_request_data.go
?? kafkarestv3/model_change_replication_factor_request_data.go
?? kafkarestv3/model_canceled_partition_data.go
?? kafkarestv3/model_replication_factor_change_cancel_data.go
?? kafkarestv3/model_replication_factor_change_cancel_data_all_of.go
?? kafkarestv3/model_replication_factor_change_cancel_data_list.go
?? kafkarestv3/model_replication_factor_change_cancel_data_list_all_of.go
?? kafkarestv3/model_replication_factor_change_data.go
?? kafkarestv3/model_replication_factor_change_data_all_of.go
?? kafkarestv3/model_replication_factor_change_data_list.go
?? kafkarestv3/model_replication_factor_change_data_list_all_of.go
```

- [ ] **Step 5: Commit**

```bash
git -C /Users/princepurohit/src/kafka-rest-sdk-go add kafkarestv3/api_replication_factor_change_v3.go kafkarestv3/model_cancel_replication_factor_change_request_data.go kafkarestv3/model_change_replication_factor_batch_request_data.go kafkarestv3/model_change_replication_factor_request_data.go kafkarestv3/model_canceled_partition_data.go kafkarestv3/model_replication_factor_change_cancel_data.go kafkarestv3/model_replication_factor_change_cancel_data_all_of.go kafkarestv3/model_replication_factor_change_cancel_data_list.go kafkarestv3/model_replication_factor_change_cancel_data_list_all_of.go kafkarestv3/model_replication_factor_change_data.go kafkarestv3/model_replication_factor_change_data_all_of.go kafkarestv3/model_replication_factor_change_data_list.go kafkarestv3/model_replication_factor_change_data_list_all_of.go kafkarestv3/docs/CancelReplicationFactorChangeRequestData.md kafkarestv3/docs/CanceledPartitionData.md kafkarestv3/docs/ChangeReplicationFactorBatchRequestData.md kafkarestv3/docs/ChangeReplicationFactorRequestData.md kafkarestv3/docs/ReplicationFactorChangeCancelData.md kafkarestv3/docs/ReplicationFactorChangeCancelDataAllOf.md kafkarestv3/docs/ReplicationFactorChangeCancelDataList.md kafkarestv3/docs/ReplicationFactorChangeCancelDataListAllOf.md kafkarestv3/docs/ReplicationFactorChangeData.md kafkarestv3/docs/ReplicationFactorChangeDataAllOf.md kafkarestv3/docs/ReplicationFactorChangeDataList.md kafkarestv3/docs/ReplicationFactorChangeDataListAllOf.md kafkarestv3/docs/ReplicationFactorChangeV3Api.md
git -C /Users/princepurohit/src/kafka-rest-sdk-go commit -m "Add generated interface, models, and docs for Replication Factor Change API"
```

---

### Task 2: Register ReplicationFactorChangeV3Api in client.go

**Files:**
- Modify: `kafkarestv3/client.go`

**Interfaces:**
- Consumes: `ReplicationFactorChangeV3Api` interface and `ReplicationFactorChangeV3ApiService` type from `kafkarestv3/api_replication_factor_change_v3.go` (Task 1).
- Produces: `APIClient.ReplicationFactorChangeV3Api` field, usable by SDK consumers as `client.ReplicationFactorChangeV3Api.ClustersClusterIdReplicationFactorChangesGet(...)` etc.

- [ ] **Step 1: Add the struct field**

In `kafkarestv3/client.go`, find:
```go
	ReplicaV3Api ReplicaV3Api

	ShareGroupV3Api ShareGroupV3Api
```
Replace with:
```go
	ReplicaV3Api ReplicaV3Api

	ReplicationFactorChangeV3Api ReplicationFactorChangeV3Api

	ShareGroupV3Api ShareGroupV3Api
```

- [ ] **Step 2: Add the constructor line**

In `kafkarestv3/client.go`, find:
```go
	c.ReplicaV3Api = (*ReplicaV3ApiService)(&c.common)
	c.ShareGroupV3Api = (*ShareGroupV3ApiService)(&c.common)
```
Replace with:
```go
	c.ReplicaV3Api = (*ReplicaV3ApiService)(&c.common)
	c.ReplicationFactorChangeV3Api = (*ReplicationFactorChangeV3ApiService)(&c.common)
	c.ShareGroupV3Api = (*ShareGroupV3ApiService)(&c.common)
```

- [ ] **Step 3: Build**

Run: `cd /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3 && go build ./...`
Expected: no errors.

- [ ] **Step 4: Commit**

```bash
git -C /Users/princepurohit/src/kafka-rest-sdk-go add kafkarestv3/client.go
git -C /Users/princepurohit/src/kafka-rest-sdk-go commit -m "Register ReplicationFactorChangeV3Api in APIClient"
```

---

### Task 3: Update kafkarestv3/README.md doc index

**Files:**
- Modify: `kafkarestv3/README.md`

**Interfaces:** None — documentation only.

- [ ] **Step 1: Add the 3 endpoint rows**

In `kafkarestv3/README.md`, find:
```
*ReplicaV3Api* | [**ClustersClusterIdBrokersBrokerIdPartitionReplicasGet**](docs/ReplicaV3Api.md#clustersclusteridbrokersbrokeridpartitionreplicasget) | **Get** /clusters/{cluster_id}/brokers/{broker_id}/partition-replicas | Search Replicas by Broker
*ShareGroupV3Api* | [**GetKafkaShareGroup**](docs/ShareGroupV3Api.md#getkafkasharegroup) | **Get** /clusters/{cluster_id}/share-groups/{group_id} | Get Share Group
```
Replace with:
```
*ReplicaV3Api* | [**ClustersClusterIdBrokersBrokerIdPartitionReplicasGet**](docs/ReplicaV3Api.md#clustersclusteridbrokersbrokeridpartitionreplicasget) | **Get** /clusters/{cluster_id}/brokers/{broker_id}/partition-replicas | Search Replicas by Broker
*ReplicationFactorChangeV3Api* | [**ClustersClusterIdReplicationFactorChangesGet**](docs/ReplicationFactorChangeV3Api.md#clustersclusteridreplicationfactorchangesget) | **Get** /clusters/{cluster_id}/replication-factor-changes | List Replication Factor Changes
*ReplicationFactorChangeV3Api* | [**ClustersClusterIdReplicationFactorChangescancelPost**](docs/ReplicationFactorChangeV3Api.md#clustersclusteridreplicationfactorchangescancelpost) | **Post** /clusters/{cluster_id}/replication-factor-changes:cancel | Cancel Replication Factor Change
*ReplicationFactorChangeV3Api* | [**ClustersClusterIdReplicationFactorChangescreatePost**](docs/ReplicationFactorChangeV3Api.md#clustersclusteridreplicationfactorchangescreatepost) | **Post** /clusters/{cluster_id}/replication-factor-changes:create | Change Replication Factor
*ShareGroupV3Api* | [**GetKafkaShareGroup**](docs/ShareGroupV3Api.md#getkafkasharegroup) | **Get** /clusters/{cluster_id}/share-groups/{group_id} | Get Share Group
```

- [ ] **Step 2: Add the "Cancel.../Canceled..." model rows (C-section)**

In `kafkarestv3/README.md`, find:
```
 - [BrokerTaskType](docs/BrokerTaskType.md)
 - [CellConfigData](docs/CellConfigData.md)
```
Replace with:
```
 - [BrokerTaskType](docs/BrokerTaskType.md)
 - [CancelReplicationFactorChangeRequestData](docs/CancelReplicationFactorChangeRequestData.md)
 - [CanceledPartitionData](docs/CanceledPartitionData.md)
 - [CellConfigData](docs/CellConfigData.md)
```

- [ ] **Step 3: Add the "Change..." model rows (C-section)**

In `kafkarestv3/README.md`, find:
```
 - [CellDataAllOf](docs/CellDataAllOf.md)
 - [ClusterConfigData](docs/ClusterConfigData.md)
```
Replace with:
```
 - [CellDataAllOf](docs/CellDataAllOf.md)
 - [ChangeReplicationFactorBatchRequestData](docs/ChangeReplicationFactorBatchRequestData.md)
 - [ChangeReplicationFactorRequestData](docs/ChangeReplicationFactorRequestData.md)
 - [ClusterConfigData](docs/ClusterConfigData.md)
```

- [ ] **Step 4: Add the "ReplicationFactorChange..." model rows (R-section)**

In `kafkarestv3/README.md`, find:
```
 - [ReplicaStatusDataListAllOf](docs/ReplicaStatusDataListAllOf.md)
 - [Resource](docs/Resource.md)
```
Replace with:
```
 - [ReplicaStatusDataListAllOf](docs/ReplicaStatusDataListAllOf.md)
 - [ReplicationFactorChangeCancelData](docs/ReplicationFactorChangeCancelData.md)
 - [ReplicationFactorChangeCancelDataAllOf](docs/ReplicationFactorChangeCancelDataAllOf.md)
 - [ReplicationFactorChangeCancelDataList](docs/ReplicationFactorChangeCancelDataList.md)
 - [ReplicationFactorChangeCancelDataListAllOf](docs/ReplicationFactorChangeCancelDataListAllOf.md)
 - [ReplicationFactorChangeData](docs/ReplicationFactorChangeData.md)
 - [ReplicationFactorChangeDataAllOf](docs/ReplicationFactorChangeDataAllOf.md)
 - [ReplicationFactorChangeDataList](docs/ReplicationFactorChangeDataList.md)
 - [ReplicationFactorChangeDataListAllOf](docs/ReplicationFactorChangeDataListAllOf.md)
 - [Resource](docs/Resource.md)
```

- [ ] **Step 5: Commit**

```bash
git -C /Users/princepurohit/src/kafka-rest-sdk-go add kafkarestv3/README.md
git -C /Users/princepurohit/src/kafka-rest-sdk-go commit -m "Add Replication Factor Change API to README doc index"
```

---

### Task 4: Add paths to embedded kafkarestv3/api/openapi.yaml

**Files:**
- Modify: `kafkarestv3/api/openapi.yaml`

**Interfaces:** None — this file is a documentation copy, not consumed by Go code.

- [ ] **Step 1: Insert the 3 new paths**

In `kafkarestv3/api/openapi.yaml`, find (the end of the `/clusters/{cluster_id}/remove-broker-tasks/{broker_id}` block, immediately before the next path):
```
      summary: Get Remove Broker Task
      tags:
      - RemoveBrokerTask
  /clusters/{cluster_id}/brokers/{broker_id}:unregister:
```
Replace with:
```
      summary: Get Remove Broker Task
      tags:
      - RemoveBrokerTask
  /clusters/{cluster_id}/replication-factor-changes:
    get:
      description: |-
        [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

        Return all tracked Replication Factor Changes in the cluster specified with ``cluster_id``.
      parameters:
      - description: The Kafka cluster ID.
        example: cluster-1
        explode: false
        in: path
        name: cluster_id
        required: true
        schema:
          type: string
        style: simple
      responses:
        "200":
          content:
            application/json:
              example:
                kind: KafkaReplicationFactorChangeList
                metadata:
                  self: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/replication-factor-changes
                  next: null
                data:
                - kind: KafkaReplicationFactorChange
                  metadata:
                    self: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/replication-factor-changes/topic-1
                    resource_name: crn:///kafka=cluster-1/replication-factor-changes=topic-1
                  cluster_id: cluster-1
                  topic_name: topic-1
                  replication_factor: 6
                  status: IN_PROGRESS
                  created_at: 2026-08-07T07:20:50Z
                  updated_at: 2026-08-07T07:22:10Z
                  error_code: null
                  error_message: null
                  topic:
                    related: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/topics/topic-1
                - kind: KafkaReplicationFactorChange
                  metadata:
                    self: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/replication-factor-changes/topic-2
                    resource_name: crn:///kafka=cluster-1/replication-factor-changes=topic-2
                  cluster_id: cluster-1
                  topic_name: topic-2
                  replication_factor: 2
                  status: FAILED
                  created_at: 2026-08-07T07:19:05Z
                  updated_at: 2026-08-07T07:19:40Z
                  error_code: 10038
                  error_message: The requested replication factor is unattainable due to placement constraints.
                  topic:
                    related: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/topics/topic-2
              schema:
                $ref: '#/components/schemas/ReplicationFactorChangeDataList'
          description: The list of tracked Replication Factor Changes.
        "400":
          content:
            application/json:
              examples:
                bad_request_cannot_deserialize:
                  description: Thrown when trying to deserialize an integer from non-integer
                    data.
                  value:
                    error_code: 400
                    message: 'Cannot deserialize value of type `java.lang.Integer`
                      from String "A": not a valid `java.lang.Integer` value'
                unsupported_version_exception:
                  description: Thrown when the version of this API is not supported
                    in the underlying Kafka cluster.
                  value:
                    error_code: 40035
                    message: The version of this API is not supported in the underlying
                      Kafka cluster.
              schema:
                $ref: '#/components/schemas/Error'
          description: Indicates a bad request error. It could be caused by an unexpected
            request body format or other forms of request validation failure.
        "401":
          content:
            application/json:
              examples:
                kafka_authentication_failed:
                  description: Thrown when using Basic authentication with wrong Kafka
                    credentials.
                  value:
                    error_code: 40101
                    message: Authentication failed
              schema:
                $ref: '#/components/schemas/Error'
          description: Indicates a client authentication error. Kafka authentication
            failures will contain error code 40101 in the response body.
        "429":
          content:
            text/html:
              example:
                description: A sample response from Jetty's DoSFilter.
                value: <html> <head> <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
                  <title>Error 429 Too Many Requests</title> </head> <body> <h2>HTTP
                  ERROR 429 Too Many Requests</h2> <table> <tr> <th>URI:</th> <td>/v3/clusters/my-cluster</td>
                  </tr> <tr> <th>STATUS:</th> <td>429</td> </tr> <tr> <th>MESSAGE:</th>
                  <td>Too Many Requests</td> </tr> <tr> <th>SERVLET:</th> <td>default</td>
                  </tr> </table> </body> </html>
              schema:
                type: string
          description: Indicates that a rate limit threshold has been reached, and
            the client should retry again later.
        "5XX":
          content:
            application/json:
              examples:
                generic_internal_server_error:
                  description: Thrown for generic HTTP 500 errors.
                  value:
                    error_code: 500
                    message: Internal Server Error
                produce_v3_missing_schema:
                  description: Thrown when the specified schema cannot be fetched
                    from Schema Registry.
                  value:
                    error_code: 50002
                    message: Error when fetching latest schema version. subject =
                      my-topic
              schema:
                $ref: '#/components/schemas/Error'
          description: A server-side problem that might not be addressable from the
            client side. Retriable Kafka errors will contain error code 50003 in the
            response body.
      summary: List Replication Factor Changes
      tags:
      - ReplicationFactorChange (v3)
  /clusters/{cluster_id}/replication-factor-changes:create:
    post:
      description: |-
        [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

        Submit a Replication Factor change for one or more topics in the cluster specified with ``cluster_id``. The change is applied asynchronously; poll the List Replication Factor Changes API to follow progress.
      parameters:
      - description: The Kafka cluster ID.
        example: cluster-1
        explode: false
        in: path
        name: cluster_id
        required: true
        schema:
          type: string
        style: simple
      requestBody:
        $ref: '#/components/requestBodies/ChangeReplicationFactorRequest'
      responses:
        "204":
          description: No Content
        "400":
          content:
            application/json:
              examples:
                bad_request_cannot_deserialize:
                  description: Thrown when trying to deserialize an integer from non-integer
                    data.
                  value:
                    error_code: 400
                    message: 'Cannot deserialize value of type `java.lang.Integer`
                      from String "A": not a valid `java.lang.Integer` value'
                unsupported_version_exception:
                  description: Thrown when the version of this API is not supported
                    in the underlying Kafka cluster.
                  value:
                    error_code: 40035
                    message: The version of this API is not supported in the underlying
                      Kafka cluster.
              schema:
                $ref: '#/components/schemas/Error'
          description: Indicates a bad request error. It could be caused by an unexpected
            request body format or other forms of request validation failure.
        "401":
          content:
            application/json:
              examples:
                kafka_authentication_failed:
                  description: Thrown when using Basic authentication with wrong Kafka
                    credentials.
                  value:
                    error_code: 40101
                    message: Authentication failed
              schema:
                $ref: '#/components/schemas/Error'
          description: Indicates a client authentication error. Kafka authentication
            failures will contain error code 40101 in the response body.
        "429":
          content:
            text/html:
              example:
                description: A sample response from Jetty's DoSFilter.
                value: <html> <head> <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
                  <title>Error 429 Too Many Requests</title> </head> <body> <h2>HTTP
                  ERROR 429 Too Many Requests</h2> <table> <tr> <th>URI:</th> <td>/v3/clusters/my-cluster</td>
                  </tr> <tr> <th>STATUS:</th> <td>429</td> </tr> <tr> <th>MESSAGE:</th>
                  <td>Too Many Requests</td> </tr> <tr> <th>SERVLET:</th> <td>default</td>
                  </tr> </table> </body> </html>
              schema:
                type: string
          description: Indicates that a rate limit threshold has been reached, and
            the client should retry again later.
        "5XX":
          content:
            application/json:
              examples:
                generic_internal_server_error:
                  description: Thrown for generic HTTP 500 errors.
                  value:
                    error_code: 500
                    message: Internal Server Error
                produce_v3_missing_schema:
                  description: Thrown when the specified schema cannot be fetched
                    from Schema Registry.
                  value:
                    error_code: 50002
                    message: Error when fetching latest schema version. subject =
                      my-topic
              schema:
                $ref: '#/components/schemas/Error'
          description: A server-side problem that might not be addressable from the
            client side. Retriable Kafka errors will contain error code 50003 in the
            response body.
      summary: Change Replication Factor
      tags:
      - ReplicationFactorChange (v3)
  /clusters/{cluster_id}/replication-factor-changes:cancel:
    post:
      description: |-
        [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

        Cancel the in-flight Replication Factor Change for the topics specified in the request body, in the cluster specified with ``cluster_id``. Cancellation does not roll back replicas already applied.
      parameters:
      - description: The Kafka cluster ID.
        example: cluster-1
        explode: false
        in: path
        name: cluster_id
        required: true
        schema:
          type: string
        style: simple
      requestBody:
        $ref: '#/components/requestBodies/CancelReplicationFactorChangeRequest'
      responses:
        "200":
          content:
            application/json:
              example:
                kind: KafkaReplicationFactorChangeCancelList
                metadata:
                  self: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/replication-factor-changes
                  next: null
                data:
                - kind: KafkaReplicationFactorChangeCancel
                  metadata:
                    self: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/replication-factor-changes/topic-1
                    resource_name: crn:///kafka=cluster-1/replication-factor-changes=topic-1
                  cluster_id: cluster-1
                  topic_name: topic-1
                  prior_status: IN_PROGRESS
                  error_code: null
                  error_message: null
                  canceled_partitions:
                  - partition_id: 0
                  - partition_id: 1
                  topic:
                    related: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/topics/topic-1
                - kind: KafkaReplicationFactorChangeCancel
                  metadata:
                    self: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/replication-factor-changes/topic-2
                    resource_name: crn:///kafka=cluster-1/replication-factor-changes=topic-2
                  cluster_id: cluster-1
                  topic_name: topic-2
                  prior_status: null
                  error_code: 10037
                  error_message: No Replication Factor Change is in progress for topic topic-2.
                  canceled_partitions: []
                  topic:
                    related: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/topics/topic-2
              schema:
                $ref: '#/components/schemas/ReplicationFactorChangeCancelDataList'
          description: The list of canceled Replication Factor Changes. Topics for
            which no change was tracked are included with a non-null error_code and
            error_message.
        "400":
          content:
            application/json:
              examples:
                bad_request_cannot_deserialize:
                  description: Thrown when trying to deserialize an integer from non-integer
                    data.
                  value:
                    error_code: 400
                    message: 'Cannot deserialize value of type `java.lang.Integer`
                      from String "A": not a valid `java.lang.Integer` value'
                unsupported_version_exception:
                  description: Thrown when the version of this API is not supported
                    in the underlying Kafka cluster.
                  value:
                    error_code: 40035
                    message: The version of this API is not supported in the underlying
                      Kafka cluster.
              schema:
                $ref: '#/components/schemas/Error'
          description: Indicates a bad request error. It could be caused by an unexpected
            request body format or other forms of request validation failure.
        "401":
          content:
            application/json:
              examples:
                kafka_authentication_failed:
                  description: Thrown when using Basic authentication with wrong Kafka
                    credentials.
                  value:
                    error_code: 40101
                    message: Authentication failed
              schema:
                $ref: '#/components/schemas/Error'
          description: Indicates a client authentication error. Kafka authentication
            failures will contain error code 40101 in the response body.
        "429":
          content:
            text/html:
              example:
                description: A sample response from Jetty's DoSFilter.
                value: <html> <head> <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
                  <title>Error 429 Too Many Requests</title> </head> <body> <h2>HTTP
                  ERROR 429 Too Many Requests</h2> <table> <tr> <th>URI:</th> <td>/v3/clusters/my-cluster</td>
                  </tr> <tr> <th>STATUS:</th> <td>429</td> </tr> <tr> <th>MESSAGE:</th>
                  <td>Too Many Requests</td> </tr> <tr> <th>SERVLET:</th> <td>default</td>
                  </tr> </table> </body> </html>
              schema:
                type: string
          description: Indicates that a rate limit threshold has been reached, and
            the client should retry again later.
        "5XX":
          content:
            application/json:
              examples:
                generic_internal_server_error:
                  description: Thrown for generic HTTP 500 errors.
                  value:
                    error_code: 500
                    message: Internal Server Error
                produce_v3_missing_schema:
                  description: Thrown when the specified schema cannot be fetched
                    from Schema Registry.
                  value:
                    error_code: 50002
                    message: Error when fetching latest schema version. subject =
                      my-topic
              schema:
                $ref: '#/components/schemas/Error'
          description: A server-side problem that might not be addressable from the
            client side. Retriable Kafka errors will contain error code 50003 in the
            response body.
      summary: Cancel Replication Factor Change
      tags:
      - ReplicationFactorChange (v3)
  /clusters/{cluster_id}/brokers/{broker_id}:unregister:
```

- [ ] **Step 2: Validate YAML syntax**

Run: `ruby -ryaml -e "YAML.load_file('/Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/api/openapi.yaml')" && echo OK`
Expected: `OK` (no parse errors).

---

### Task 5: Add requestBodies, responses, and schemas to embedded kafkarestv3/api/openapi.yaml

**Files:**
- Modify: `kafkarestv3/api/openapi.yaml`

**Interfaces:** None — documentation only. Depends on Task 4 (paths reference these components by name).

- [ ] **Step 1: Insert the 2 new requestBodies**

In `kafkarestv3/api/openapi.yaml`, find:
```
    BrokerReplicaExclusionBatchRequest:
      content:
        application/json:
          example:
            data:
            - broker_id: 1
              reason: The broker is to be removed.
            - broker_id: 2
              reason: The broker is to be removed.
          schema:
            $ref: '#/components/schemas/BrokerReplicaExclusionBatchRequestData'
      description: Alter Broker Replica Exclusions.
  responses:
```
Replace with:
```
    BrokerReplicaExclusionBatchRequest:
      content:
        application/json:
          example:
            data:
            - broker_id: 1
              reason: The broker is to be removed.
            - broker_id: 2
              reason: The broker is to be removed.
          schema:
            $ref: '#/components/schemas/BrokerReplicaExclusionBatchRequestData'
      description: Alter Broker Replica Exclusions.
    ChangeReplicationFactorRequest:
      content:
        application/json:
          example:
            data:
            - topic_name: topic-1
              replication_factor: 6
            - topic_name: topic-2
              replication_factor: 2
            timeout_ms: 60000
          schema:
            $ref: '#/components/schemas/ChangeReplicationFactorBatchRequestData'
      description: Replication Factor changes to submit, keyed by topic.
    CancelReplicationFactorChangeRequest:
      content:
        application/json:
          example:
            topic_names:
            - topic-1
            - topic-2
            timeout_ms: 60000
          schema:
            $ref: '#/components/schemas/CancelReplicationFactorChangeRequestData'
      description: Topic names whose Replication Factor Change should be canceled.
  responses:
```

- [ ] **Step 2: Insert the 2 new named responses**

In `kafkarestv3/api/openapi.yaml`, find the end of `ListRemoveBrokerTaskResponse` (immediately before `NoContentResponse`):
```
              broker:
                related: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/brokers/2
          schema:
            $ref: '#/components/schemas/RemoveBrokerTaskDataList'
      description: The list of remove broker tasks.
    NoContentResponse:
```
Replace with:
```
              broker:
                related: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/brokers/2
          schema:
            $ref: '#/components/schemas/RemoveBrokerTaskDataList'
      description: The list of remove broker tasks.
    ListReplicationFactorChangeResponse:
      content:
        application/json:
          example:
            kind: KafkaReplicationFactorChangeList
            metadata:
              self: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/replication-factor-changes
              next: null
            data:
            - kind: KafkaReplicationFactorChange
              metadata:
                self: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/replication-factor-changes/topic-1
                resource_name: crn:///kafka=cluster-1/replication-factor-changes=topic-1
              cluster_id: cluster-1
              topic_name: topic-1
              replication_factor: 6
              status: IN_PROGRESS
              created_at: 2026-08-07T07:20:50Z
              updated_at: 2026-08-07T07:22:10Z
              error_code: null
              error_message: null
              topic:
                related: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/topics/topic-1
            - kind: KafkaReplicationFactorChange
              metadata:
                self: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/replication-factor-changes/topic-2
                resource_name: crn:///kafka=cluster-1/replication-factor-changes=topic-2
              cluster_id: cluster-1
              topic_name: topic-2
              replication_factor: 2
              status: FAILED
              created_at: 2026-08-07T07:19:05Z
              updated_at: 2026-08-07T07:19:40Z
              error_code: 10038
              error_message: The requested replication factor is unattainable due to placement constraints.
              topic:
                related: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/topics/topic-2
          schema:
            $ref: '#/components/schemas/ReplicationFactorChangeDataList'
      description: The list of tracked Replication Factor Changes.
    ListCancelReplicationFactorChangeResponse:
      content:
        application/json:
          example:
            kind: KafkaReplicationFactorChangeCancelList
            metadata:
              self: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/replication-factor-changes
              next: null
            data:
            - kind: KafkaReplicationFactorChangeCancel
              metadata:
                self: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/replication-factor-changes/topic-1
                resource_name: crn:///kafka=cluster-1/replication-factor-changes=topic-1
              cluster_id: cluster-1
              topic_name: topic-1
              prior_status: IN_PROGRESS
              error_code: null
              error_message: null
              canceled_partitions:
              - partition_id: 0
              - partition_id: 1
              topic:
                related: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/topics/topic-1
            - kind: KafkaReplicationFactorChangeCancel
              metadata:
                self: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/replication-factor-changes/topic-2
                resource_name: crn:///kafka=cluster-1/replication-factor-changes=topic-2
              cluster_id: cluster-1
              topic_name: topic-2
              prior_status: null
              error_code: 10037
              error_message: No Replication Factor Change is in progress for topic topic-2.
              canceled_partitions: []
              topic:
                related: https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/topics/topic-2
          schema:
            $ref: '#/components/schemas/ReplicationFactorChangeCancelDataList'
      description: The list of canceled Replication Factor Changes. Topics for which
        no change was tracked are included with a non-null error_code and error_message.
    NoContentResponse:
```

- [ ] **Step 3: Insert the 8 new top-level schemas**

In `kafkarestv3/api/openapi.yaml`, find:
```
    AlterBrokerReplicaExclusionDataList:
      allOf:
      - $ref: '#/components/schemas/ResourceCollection'
      - $ref: '#/components/schemas/AlterBrokerReplicaExclusionDataList_allOf'
    UnregisterBrokerData:
```
Replace with:
```
    AlterBrokerReplicaExclusionDataList:
      allOf:
      - $ref: '#/components/schemas/ResourceCollection'
      - $ref: '#/components/schemas/AlterBrokerReplicaExclusionDataList_allOf'
    ChangeReplicationFactorRequestData:
      example:
        topic_name: topic_name
        replication_factor: 0
      properties:
        topic_name:
          type: string
        replication_factor:
          type: integer
      required:
      - replication_factor
      - topic_name
      type: object
    ChangeReplicationFactorBatchRequestData:
      example:
        data:
        - topic_name: topic_name
          replication_factor: 0
        - topic_name: topic_name
          replication_factor: 0
        timeout_ms: 0
      properties:
        data:
          items:
            $ref: '#/components/schemas/ChangeReplicationFactorRequestData'
          type: array
        timeout_ms:
          description: 'The time in milliseconds to wait for the request to complete. Default: 60000'
          nullable: true
          type: integer
      required:
      - data
      type: object
    CancelReplicationFactorChangeRequestData:
      example:
        topic_names:
        - topic_names
        - topic_names
        timeout_ms: 0
      properties:
        topic_names:
          items:
            type: string
          type: array
        timeout_ms:
          description: 'The time in milliseconds to wait for the request to complete. Default: 60000'
          nullable: true
          type: integer
      required:
      - topic_names
      type: object
    ReplicationFactorChangeData:
      allOf:
      - $ref: '#/components/schemas/Resource'
      - $ref: '#/components/schemas/ReplicationFactorChangeData_allOf'
    ReplicationFactorChangeDataList:
      allOf:
      - $ref: '#/components/schemas/ResourceCollection'
      - $ref: '#/components/schemas/ReplicationFactorChangeDataList_allOf'
    CanceledPartitionData:
      example:
        partition_id: 0
      properties:
        partition_id:
          type: integer
      required:
      - partition_id
      type: object
    ReplicationFactorChangeCancelData:
      allOf:
      - $ref: '#/components/schemas/Resource'
      - $ref: '#/components/schemas/ReplicationFactorChangeCancelData_allOf'
    ReplicationFactorChangeCancelDataList:
      allOf:
      - $ref: '#/components/schemas/ResourceCollection'
      - $ref: '#/components/schemas/ReplicationFactorChangeCancelDataList_allOf'
    UnregisterBrokerData:
```

- [ ] **Step 4: Insert the 4 new `_allOf` fragment schemas**

In `kafkarestv3/api/openapi.yaml`, find:
```
    AlterBrokerReplicaExclusionDataList_allOf:
      properties:
        data:
          items:
            $ref: '#/components/schemas/AlterBrokerReplicaExclusionData'
          type: array
      required:
      - data
    RemoveBrokerTaskData_allOf:
```
Replace with:
```
    AlterBrokerReplicaExclusionDataList_allOf:
      properties:
        data:
          items:
            $ref: '#/components/schemas/AlterBrokerReplicaExclusionData'
          type: array
      required:
      - data
    ReplicationFactorChangeData_allOf:
      properties:
        cluster_id:
          type: string
        topic_name:
          type: string
        replication_factor:
          description: The desired replication factor for the topic.
          type: integer
        status:
          type: string
          x-extensible-enum:
          - PENDING
          - IN_PROGRESS
          - COMPLETED
          - FAILED
          - UNKNOWN
        created_at:
          description: The date and time at which this Replication Factor Change
            was created.
          example: 2019-10-12T07:20:50Z
          format: date-time
          readOnly: true
          type: string
        updated_at:
          description: The date and time at which this Replication Factor Change
            was last updated.
          example: 2019-10-12T07:20:50Z
          format: date-time
          readOnly: true
          type: string
        error_code:
          nullable: true
          type: integer
        error_message:
          nullable: true
          type: string
        topic:
          $ref: '#/components/schemas/Relationship'
      required:
      - cluster_id
      - created_at
      - replication_factor
      - status
      - topic
      - topic_name
      - updated_at
    ReplicationFactorChangeDataList_allOf:
      properties:
        data:
          items:
            $ref: '#/components/schemas/ReplicationFactorChangeData'
          type: array
      required:
      - data
    ReplicationFactorChangeCancelData_allOf:
      properties:
        cluster_id:
          type: string
        topic_name:
          type: string
        prior_status:
          description: Status just before cancellation (``PENDING`` or ``IN_PROGRESS``);
            null if no tracked change existed for the topic.
          nullable: true
          type: string
          x-extensible-enum:
          - PENDING
          - IN_PROGRESS
          - COMPLETED
          - FAILED
          - UNKNOWN
        error_code:
          nullable: true
          type: integer
        error_message:
          nullable: true
          type: string
        canceled_partitions:
          items:
            $ref: '#/components/schemas/CanceledPartitionData'
          type: array
        topic:
          $ref: '#/components/schemas/Relationship'
      required:
      - canceled_partitions
      - cluster_id
      - error_code
      - error_message
      - prior_status
      - topic
      - topic_name
    ReplicationFactorChangeCancelDataList_allOf:
      properties:
        data:
          items:
            $ref: '#/components/schemas/ReplicationFactorChangeCancelData'
          type: array
      required:
      - data
    RemoveBrokerTaskData_allOf:
```

- [ ] **Step 5: Validate YAML syntax**

Run: `ruby -ryaml -e "YAML.load_file('/Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3/api/openapi.yaml')" && echo OK`
Expected: `OK` (no parse errors).

- [ ] **Step 6: Commit**

```bash
git -C /Users/princepurohit/src/kafka-rest-sdk-go add kafkarestv3/api/openapi.yaml
git -C /Users/princepurohit/src/kafka-rest-sdk-go commit -m "Add Replication Factor Change API to embedded openapi.yaml"
```

---

### Task 6: Final verification

**Files:** None (verification only).

**Interfaces:** None.

- [ ] **Step 1: Build the whole SDK module**

Run: `cd /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3 && go build ./...`
Expected: no errors, no output.

- [ ] **Step 2: Vet the whole SDK module**

Run: `cd /Users/princepurohit/src/kafka-rest-sdk-go/kafkarestv3 && go vet ./...`
Expected: no errors, no output.

- [ ] **Step 3: Confirm go.mod is unchanged from HEAD before this branch**

Run: `git -C /Users/princepurohit/src/kafka-rest-sdk-go diff master -- kafkarestv3/go.mod`
Expected: no output (empty diff).

- [ ] **Step 4: Confirm working tree is clean and branch has exactly the expected commits**

Run: `git -C /Users/princepurohit/src/kafka-rest-sdk-go status --porcelain=v1`
Expected: no output (clean).

Run: `git -C /Users/princepurohit/src/kafka-rest-sdk-go log --oneline master..HEAD`
Expected: 4 commits (Task 1, 2, 3, 5 commits), most-recent first.

**Do not push and do not open a PR** — this branch stays local for the user to review and push themselves.
