# LicenseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**Category** | **string** |  | 
**CategoryShortName** | **string** |  | 
**LicenseType** | **string** |  | 
**ExpiresAt** | [**time.Time**](time.Time.md) | Absent when the license does not expire (e.g., the built-in Developer / Free-Forever license). | [optional] 
**Audience** | **string** |  | [optional] 
**LicenseJwt** | **string** |  | 
**Status** | **string** |  | [readonly] 
**TopicName** | **string** | Internal Kafka topic the license is stored in. Absent when not applicable. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


