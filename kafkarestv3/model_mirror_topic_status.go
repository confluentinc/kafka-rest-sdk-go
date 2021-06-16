/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3

// MirrorTopicStatus the model 'MirrorTopicStatus'
type MirrorTopicStatus string

// List of MirrorTopicStatus
const (
	MIRRORTOPICSTATUS_ACTIVE          MirrorTopicStatus = "ACTIVE"
	MIRRORTOPICSTATUS_FAILED          MirrorTopicStatus = "FAILED"
	MIRRORTOPICSTATUS_PAUSED          MirrorTopicStatus = "PAUSED"
	MIRRORTOPICSTATUS_STOPPED         MirrorTopicStatus = "STOPPED"
	MIRRORTOPICSTATUS_PENDING_STOPPED MirrorTopicStatus = "PENDING_STOPPED"
)
