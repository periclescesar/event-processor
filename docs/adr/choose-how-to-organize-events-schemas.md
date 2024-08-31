# Choose how to organize events schemas

## Status
accepted

Date: 2024-08-31

## Context
Given that the application will process events coming from different sources,
having different structures and validation rules between them,
we need everyone to have some way of distinguishing 
them from the others in order to obtain the correct validation scheme.
It is also necessary to have a parameter to identify which client the event belongs to.

## Decision
Therefore, using JSON Schema we have 2 ways to handle this requirement:
### 1. ~~Using a single root schema~~
With a single root schema, we need to create a file with the following structure:
```json
{
  "$id": "file://configs/events-schemas/event-base.schema.json",
  "$schema": "https://json-schema.org/draft/2019-09/schema",
  "title": "base for all events",
  "description": "this schema should be used as a base for all events",
  "required": ["eventType", "tenant"],
  "type": "object",
  "properties": {
    "eventType": {
      "type": "string",
      "minLength": 3,
      "maxLength": 255
    },
    "tenant": {
      "type": "string",
      "minLength": 3,
      "maxLength": 255
    }
  },
  "anyOf": [
    {
      "type": "object",
      "properties": {
        "street_address": { "type": "string" },
        "city": { "type": "string" },
        "state": { "type": "string" }
      },
      "required": ["street_address", "city", "state"]
    }
  ]
}
```
Inside the `anyOf` parameter we place the definitions of the other events.

The problem with this approach is that we cannot use the `"additionalProperties": false` parameter,
which would serve to restrict additional attributes in the event,
and because of this nothing prevents unwanted data from being passed to the next step.
We also do not guarantee validation of the schema according to the `eventType`,
if you have two events with the same structure, but with different validations,
they may generate a false positive in the validation. Furthermore,
this structure does not favor the maintainability of the schemas.

### :heavy_check_mark: 2. Using multiple schema files
We can use the following scheme as a template for creating new events:
```json
{
  "$id": "file://configs/events-schemas/event-base.schema.json",
  "$schema": "https://json-schema.org/draft/2019-09/schema",
  "title": "base for all events",
  "description": "this schema should be used as a base for all events",
  "required": ["eventType", "tenant"],
  "type": "object",
  "properties": {
    "eventType": {
      "type": "string",
      "minLength": 3,
      "maxLength": 255
    },
    "tenant": {
      "type": "string",
      "minLength": 3,
      "maxLength": 255
    }
  }
}
```
Based on this JSON Schema we can declare new events, where the file name must be the `eventType` value:
`file://configs/events-schemas/<eventType>.schema.json`

This way we can easily load the schema that validates the event according to the `eventType`,
without side effects and allowing full use of the JSON Schema specification.
You just need to ensure via code that all specifications are an extension of `event-base.schema.json`.

## Consequences

We will have to ensure that the specifications follow `event-base.schema.json`
as a base while maintaining all properties.