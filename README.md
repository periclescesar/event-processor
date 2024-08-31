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
    R[Receiver]    --> UC
    UC[SaveEventUC] <--> V
    V{isValid}
    C[Contract]    -.usedBy.-> V
    
    V --invalid--> RV[Reject Event]
    V --valid--> P[Persist]
    RV[Reject Event] --> DLQ(events.dlq)
```

# running project
1. Cloning `.env.example`
```shell
cd deployments
cp .env.example .env
```
2. starting dependencies
```shell
docker compose up -d
```
3. in cd deployments path, run terraform with
```shell
export $(cat .env | xargs -I% echo TF_VAR_%)
cd terraform
terraform init
terraform plan
terraform apply
```

# Creating a new event schema
You can use this website to create a JSON Schema: 

https://www.jsonschemavalidator.net/

using this file as a template: [event-base.schema.json](configs/events-schemas/event-base.schema.json)

After building and validating schema, save on path `configs/events-schema/` with name `<eventType>.schema.json`,
replacing `<eventType>` with the eventType value expected to use this schema to validate them.

# ADRs
* [choose-schema-validator.md](docs/adr/choose-schema-validator.md)
* [choose-how-to-organize-events-schemas.md](docs/adr/choose-how-to-organize-events-schemas.md)

# @TODO
 * [ ] Create Rabbitmq instance and provision it via terraform
 * [ ] Create a consumer for the events
 * [ ] Create a configuration contract reader 
 * [ ] Create a validator that uses the contract to reject invalid events 
 * [ ] Use a database to persist the events and be read in the future by the sender 
 * [ ] Create a parameterizable producer to put events in the pipeline 
 * [ ] Create acceptance tests 
 * [ ] Create load tests 
 * [ ] Create user on rabbitmq for application only
 * [ ] Optimize flow 
 * [ ] Create a docker container for the app 
 * [ ] Move to LocalStack 
 * [ ] Use terraform to declare resources