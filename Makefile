INIT_DIR=../kafka-rest-sdk-go/
GENERATOR_DIR=../openapi-generator
OPENAPI_SPEC_DIR=../ce-kafka-rest/api/v3/consolidated-openapi.yaml
SDK_OUTPUT_DIR=../kafka-rest-sdk-go/kafkarestv3


.PHONY: gen-deps
gen-deps:
	# You only have to run gen-deps once.
	# https://github.com/travisjeffery/mocker
	go get github.com/travisjeffery/mocker/cmd/mocker
	go install github.com/travisjeffery/mocker/cmd/mocker

	# https://github.com/confluentinc/openapi-generator
	cd .. && git clone https://github.com/confluentinc/openapi-generator.git
	cd ../openapi-generator && git checkout ath-add-interfaces && mvn clean install


.PHONY: gen
gen:
	cd $(GENERATOR_DIR) && java -jar modules/openapi-generator-cli/target/openapi-generator-cli.jar generate \
		-i $(OPENAPI_SPEC_DIR) \
		-g go \
		-o $(SDK_OUTPUT_DIR) \
		--generate-alias-as-model \
		--enable-post-process-file \
		--git-user-id confluentinc \
		--git-repo-id kafka-rest-sdk-go/kafkarestv3 \
		-p enumClassPrefix=true,packageName=kafkarestv3,generateInterfaces=true
	cd $(INIT_DIR) && cd "kafkarestv3"
	find . -type f -name "*.go" -print0 | xargs -0 sed -i '' -e 's/== 5XX/>= 500/g'
	gofmt -w .
	echo "Remember to git restore kafkarestv3/go.mod"

.PHONY: gen-mock
gen-mock:
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_aclv3.go --pkg mock $(SDK_OUTPUT_DIR)/api_aclv3.go ACLV3Api
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_balancer_status_v3.go --pkg mock $(SDK_OUTPUT_DIR)/api_balancer_status_v3.go BalancerStatusV3Api
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_broker.go --pkg mock $(SDK_OUTPUT_DIR)/api_broker.go BrokerApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_broker_replica_exclusion.go --pkg mock $(SDK_OUTPUT_DIR)/api_broker_replica_exclusion.go BrokerReplicaExclusionApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_broker_task.go --pkg mock $(SDK_OUTPUT_DIR)/api_broker_task.go BrokerTaskApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_broker_v3.go --pkg mock $(SDK_OUTPUT_DIR)/api_broker_v3.go BrokerV3Api
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_cluster_linking_v3.go --pkg mock $(SDK_OUTPUT_DIR)/api_cluster_linking_v3.go ClusterLinkingV3Api
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_cluster_v3.go --pkg mock $(SDK_OUTPUT_DIR)/api_cluster_v3.go ClusterV3Api
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_configs_v3.go --pkg mock $(SDK_OUTPUT_DIR)/api_configs_v3.go ConfigsV3Api
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_consumer_group_v3.go --pkg mock $(SDK_OUTPUT_DIR)/api_consumer_group_v3.go ConsumerGroupV3Api
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_partition.go --pkg mock $(SDK_OUTPUT_DIR)/api_partition.go PartitionApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_partition_v3.go --pkg mock $(SDK_OUTPUT_DIR)/api_partition_v3.go PartitionV3Api
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_records.go --pkg mock $(SDK_OUTPUT_DIR)/api_records.go RecordsApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_remove_broker_task.go --pkg mock $(SDK_OUTPUT_DIR)/api_remove_broker_task.go RemoveBrokerTaskApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_replica.go --pkg mock $(SDK_OUTPUT_DIR)/api_replica.go ReplicaApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_replica_status.go --pkg mock $(SDK_OUTPUT_DIR)/api_replica_status.go ReplicaStatusApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_replica_v3.go --pkg mock $(SDK_OUTPUT_DIR)/api_replica_v3.go ReplicaV3Api
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_topic_v3.go --pkg mock $(SDK_OUTPUT_DIR)/api_topic_v3.go TopicV3Api