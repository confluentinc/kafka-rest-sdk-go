# ReplicationFactorChangeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**TopicName** | **string** |  | 
**DesiredReplicationFactor** | **int32** | The desired replication factor for the topic. | 
**Action** | [**ReplicationFactorChangeAction**](ReplicationFactorChangeAction.md) |  | 
**Status** | [**ReplicationFactorChangeStatus**](ReplicationFactorChangeStatus.md) |  | 
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time at which this Replication Factor Change was created. | [readonly] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time at which this Replication Factor Change was last updated. | [readonly] 
**ErrorCode** | Pointer to **int32** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Topic** | [**Relationship**](Relationship.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


