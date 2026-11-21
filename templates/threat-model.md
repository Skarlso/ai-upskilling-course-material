# Threat model

One page. Fill it in for your agent as it stands today. Used in [module 8](../modules/08-security-and-sandboxing/README.md).

## The agent

What it is, what it is for, and where it runs.

## What it can reach

| | What exactly |
| --- | --- |
| Files it can read | |
| Files it can write | |
| Commands it can run | |
| Network it can reach | |
| Credentials it holds | |
| Tools from MCP servers | |

## The three legs

For each one: is it present, and through what?

| Leg | Present? | Through what |
| --- | --- | --- |
| Access to private data | | |
| Exposure to untrusted content | | |
| A way to send data out | | |

## My setups

One row per way you run the agent. Which leg does each setup take away?

| Setup | Leg removed | How |
| --- | --- | --- |
| Bare, on my machine | | |
| Shell tool in a container, no network | | |
| | | |

## What happened when I attacked it

What you planted, where, and what the agent did. Before the sandbox and after.

## Worst case

If this agent were talked into anything at all, what is the worst it could do with what it can reach today?

## What I would still not let it touch

And why.
