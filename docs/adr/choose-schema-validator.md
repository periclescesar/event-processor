# Choose Schema Validator

## Status
accepted

Date: 2024-08-30

## Context

To validate the structure of events passed through the application,
I can create a complete piece of code to read a model structure
and implement a validator that uses this structure for validation.
However, if I use a ready-made library for this, 
I will simplify my work and gain development time, quality, and better testability of the function.

## Decision

In this way, a standard already used in the market for this purpose is [JSON Schema](https://json-schema.org/),
which references the following libraries for manipulating and validating schemas:

| id | Lib                        | first commit | last commit | imports | imported by |
|----|----------------------------|--------------|-------------|---------|-------------|
| 1  | santhosh-tekuri/jsonschema | 2024-04-17   | 2024-07-30  | 24      | 19          |
| 2  | jsonschema                 | 2024-05-07   | 2024-07-30  | 23      | 4           |
| 3  | qri-io/jsonschema          | 2018-01-12   | 2021-08-30  | 18      | 123         |
| 4  | gojsonschema               | 2013-02-26   | 2020-10-27  | 25      | 1780        |

To ensure a library has fewer side effects from sub-dependencies,
we need to choose one with a low number of sub-dependencies. 
In this regard, all options are quite close, so none will be excluded based on this criterion alone.

For greater assurance of quality,
we need a library that has been used for a longer time and by a larger number of applications.
In this case, the `gojsonschema` library is the most widely used,
with a significant lead over the second most used library in this aspect.

It is also important that the library is up-to-date to avoid issues with future updates and to comply with new standards.
The latest version of the JSON Schema specification is `2020-12`, which is supported only by libraries `1` and `2`.
The penultimate version is `2019-09`, supported by library `3`,
while library `4` only supports the third-to-last version of the specification, `draft-07`.

Based on this information, library 3, `qri-io/jsonschema`, is the one that best meets the current needs of the tool.

## Consequences
1. This will allow me to focus on other parts of development.
2. There will be no need to document how to use the schema, just references to the specification.
3. I will have greater assurance of the quality of the validation tool.
4. It makes it easier for new users who are already familiar with the specification.
5. We won’t be able to use the most current specification.