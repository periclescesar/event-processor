# Event Processor
This is an event pipeline application
where several producers can post their events to be validated and persisted
so that another application can send the event to the end client.

Event contract validations will be described in a configuration file.

# Architecture
```mermaid
---
title: Integration Archteture
---
flowchart LR
    EP1(Event Producer) ---> B
    EP2(Event Producer) ---> B
    EP3(Event Producer) ---> B
    EP4(Event Producer) ---> B
    B[(MessageBroker)]   --> R

    subgraph EPs[Event Processor]
        direction LR
        R[Receiver] --> V
        V[Validator] --> P
        P[Persist]
    end

    P --> DB
    DB[(DataBase)] --> S
    S(Sender) --> C1
    S(Sender) --> C2
    S(Sender) --> C3
    C1([Client])
    C2([Client])
    C3([Client])
```

```mermaid
---
title: Event Processor activity diagram
---
flowchart TD
    EV[(events)] -.-> R
    R[Receiver]    --save--> US
    US[EventService] --validate--> Vt
    Vt[Validator] --get--> C[/Schema/]
    US --> V
    V{Event is Valid} --false--> RE[Reject Event]
    V --true--> P[Save on Repository]
    RE[Return Error] -.-> DLQ[(events.dlq)]
```

# Running the project
1. Cloning `.env.example`:
```shell
cd deployments
cp .env.example .env
```
You can adjust the `.env` file to suit your infrastructure needs. Set `LOG_LEVEL` to `trace` to increase verbosity and gain detailed insights into the application's internal processes.
2. Starting dependencies:
```shell
docker compose up -d rabbitmq mongodb
```

3. In deployments path, run terraform with:
```shell
export $(cat .env | xargs -I% echo TF_VAR_%)
cd terraform
terraform init
terraform apply
```

4. Up event-processor service:
```shell
cd ../
docker compose up event-processor
```
After build image you will see: 
```
starting event processor...
 [*] Waiting for messages
```

5. starting producers on another terminal:
```shell
cd deployments
docker compose up event-producer
```
Then you can see the events flowing from the producer to the processor
and being saved in mongoDb in the `events` collection.

# Creating a new event schema
You can use this website to create a JSON Schema: 

https://www.jsonschemavalidator.net/

using this file as a template: [event-base.schema.json](configs/events-schemas/event-base.schema.json)
and following specification [draft/2019-09](https://json-schema.org/draft/2019-09/json-schema-validation) 

After building and validating schema, save on path `configs/events-schemas/` with name `<eventType>.schema.json`,
replacing `<eventType>` with the eventType value expected to use this schema to validate them.

### example:
You can put this schema file on `configs/events-schemas/transaction-allowed.schema.json`:
```json
{
	"$id": "file://configs/events-schemas/transaction-allowed.schema.json",
	"$schema": "https://json-schema.org/draft/2019-09/schema",
	"title": "transaction allowed event",
	"description": "schema for all transaction allowed event",
	"required": [
		"eventType",
		"tenantId"
	],
	"type": "object",
	"properties": {
		"eventType": {
			"type": "string",
			"minLength": 3,
			"maxLength": 255
		},
		"tenantId": {
			"type": "string",
			"format": "uuid"
		},
		"transactionId": {
			"type": "string",
			"format": "uuid"
		},
		"userId": {
			"type": "string",
			"format": "uuid"
		},
		"value": {
			"type": "number"
		},
		"latitude": {
			"type": "number",
			"minimum": -90,
			"maximum": 90
		},
		"longitude": {
			"type": "number",
			"minimum": -180,
			"maximum": 180
		}
	},
	"additionalProperties": false
}
```
Then the `event-processor` will be ready to validate this new event. Like, for example:
```json
{
	"eventType": "transaction-created",
	"tenantId": "f25d45d5-213b-4de1-a35f-210f07fc55c4",
	"transactionId": "f25d45d5-213b-4de1-a35f-210f07fc55c4",
	"userId": "f25d45d5-213b-4de1-a35f-210f07fc55c4",
	"value": 2345.23
}
```
Or
```json
{
	"eventType": "transaction-created",
	"tenantId": "f25d45d5-213b-4de1-a35f-210f07fc55c4",
	"transactionId": "f25d45d5-213b-4de1-a35f-210f07fc55c4",
	"userId": "f25d45d5-213b-4de1-a35f-210f07fc55c4",
	"value": 2345.23,
	"latitude": 31,
	"longitude": 48
}
```
For testing, you can put this json events on `testdata` as a json file.
And run `event-producer` to publish these events on queue.

# Docs
* [asyncapi.yaml](docs/asyncapi.yaml) can be viewed on the website: https://studio.asyncapi.com/
* [JSON Schema](https://json-schema.org/)
* JSON Schema specification: [draft/2019-09](https://json-schema.org/draft/2019-09/json-schema-validation)

### ADRs
* [choose-schema-validator.md](docs/adr/choose-schema-validator.md)
* [choose-how-to-organize-events-schemas.md](docs/adr/choose-how-to-organize-events-schemas.md)

# For Devs:
Before any commit you need `golangci-lint` and `mockery` to generate mocks,
you can add manually or running:
```shell
sh ./scripts/install.sh
```

## Test
Before running unit tests, ensure that all mocks have been generated using:
```shell
mockery --all
```

# @TODO
* [X] Create Rabbitmq instance and provision it via terraform
* [X] Create a consumer for the events
* [X] ~~Create a configuration contract reader~~ 
* [X] Create a validator that uses the contract to reject invalid events 
* [X] Use a database to persist the events and be read in the future by the sender 
* [X] Create a docker container for the app
* [X] Increase coverage test
* [X] Improve logs
* [ ] Create acceptance tests 
* [ ] Create load tests 
* [ ] Create user on rabbitmq for application only
* [ ] Move to LocalStack 
* [X] Use terraform to declare resources
* [X] Create a parameterizable producer to put events in the pipeline
* [ ] Change Dockerfile to be multi-stage build
* [ ] Optimize flow
