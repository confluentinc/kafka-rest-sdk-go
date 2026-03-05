# ShareGroupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**ShareGroupId** | **string** |  | 
**State** | **string** |  | 
**Coordinator** | [**Relationship**](Relationship.md) |  | 
**Consumers** | [**Relationship**](Relationship.md) |  | 
**ConsumerCount** | **int32** | Number of consumers in this share group | 
**PartitionCount** | **int32** | Total number of partitions assigned to this share group across all consumers | 
**AssignedTopicPartitions** | [**[]ShareGroupTopicPartitionData**](ShareGroupTopicPartitionData.md) | List of topic-partitions assigned to this share group, including those from empty groups | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


