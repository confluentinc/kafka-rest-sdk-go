# AlterBrokerHealthRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrokerIds** | **[]int32** | Broker IDs whose health is being altered. | 
**Component** | **string** | Which broker component health is being altered. | [optional] [default to COMPONENT_UNSPECIFIED]
**HealthStatus** | **string** | The health status to set for the given component. | 
**Reason** | **string** | An explanation for altering broker health, e.g. an incident ticket. | 
**Force** | **bool** | Apply the change even if the cluster has reached the maximum number of demoted brokers. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


