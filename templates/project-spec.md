# Project spec

This is the spec for your [final project](../project/README.md).

## Track

Pipeline, harness study or tool study. Whatever you would like to get out of this study.

## The repo

Which repo the pipeline works on, and how to build and test it.

## The issue

The text exactly as the agent will see it. Write it with care, because this is most of the work. What is wrong or missing, what the result should do, and anything it must not change.

## Acceptance criteria

A short list. Each one should be something a person or a command can check.

- 
- 
- 

## Out of scope

What the agent should leave alone.

## Gates

The commands that must pass, in order.

```
go build ./...
go vet ./...
go test ./...
```

## Holdout scenarios

Titles only. The scenarios themselves live somewhere the agent can't read.

1. 
2. 
3. 

## Budget

| Limit | Value |
| --- | --- |
| Tokens per run | |
| Time per run | |
| Spend per run | |

## Design sketch

The steps of the pipeline, what drives them (a script, a CI workflow), and where the agent runs.

## Evidence in the pull request

What the pull request will contain: the plan, the gate results, the scenario score, the cost.

## Open questions

What you are unsure about. Ask the group.
