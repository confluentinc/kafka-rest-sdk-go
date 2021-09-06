# KAFKA-REST-SDK-GO

This repo holds the generated go sdks for [kafka-rest](https://github.com/confluentinc/kafka-rest) API. Each version of the MDS API is held in a versioned directory and are updated independently.

See Makefile to see how to update a version of the kafka-rest sdk.

## How to generate/update the sdk
See Makefile:

    make gen-deps
    # go install github.com/travisjeffery/mocker/cmd/mocker
    # cd .. && git clone https://github.com/confluentinc/openapi-generator.git
	# cd ../openapi-generator && git checkout ath-add-interfaces && mvn clean install

    make gen
    # cd $(GENERATOR_DIR) && java -jar modules/openapi-generator-cli/target/openapi-generator-cli.jar generate \
	# 	-i $(OPENAPI_SPEC_DIR) \
	# 	-g go \
	# 	-o $(SDK_OUTPUT_DIR) \
  	# 	-p enumClassPrefix=true,packageName=kafkarestv3,generateInterfaces=true
    # Remember to git restore kafkarestv3/go.mod

TODO: Fix to fit prod style

## kafkarestv1
See [README.md](https://github.com/confluentinc/kafka-rest-sdk-go/blob/master/kafkarestv3/README.md) for description of generated go sdk.

## Example Usage
TODO: How to use this repo, how will we set up the Jenkins and Artifactory processes.

## License

Copyright 2021 Confluent Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
