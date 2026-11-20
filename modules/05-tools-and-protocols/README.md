# Module 5: Tools and protocols

[Back to the course](../../README.md) | Previous: [Harness engineering](../04-harness-engineering/README.md) | Next: [Verification, part 1](../06-verification-part-1/README.md)

Okay, now, we are going to expose/export/excavate your tools from your agent.

That is the job of the Model Context Protocol. You write a server that offers tools, and any client that understands the language of the MCP will be able to use them. Your agent, commercial agents, whatever comes next year.

And this is the lesson: A tool is described by a prompt. Yep. The model decides when and how to call it from that prompt and nothing else. You'll create a description that your agent will most definitely misread and then you fix it and try again. This is the skill here. Writing a description that will make your agent understand what the tool is used for.

## Read

The introduction at [modelcontextprotocol.io](https://modelcontextprotocol.io/).

## Build

1. **Pick a system you know well.** Your pet project or notes folder or something that you really understand well so you'll know immediately if something is not correct.
2. **Write a read-only MCP server** for it with the [Go SDK](https://github.com/modelcontextprotocol/go-sdk). Two or three tools. The stdio transport is the simplest place to start.
3. **Make your agent a client.** On startup, ask the server for its tools and give them to the agent.
4. **Ask it things.** Especially things it can only know through your server. Models, JSON schemas, or YAML objects, or something that it couldn't know from somewhere else. Like... what's the speed of an unladen swallow?

Read-only is on purpose. Keep it that way until you have done module 8.

## Done when

Your agent answers a question it could only know the answer to through your server! And you have rewritten at least one tool description because the model misused it.

## Discuss

What made a tool description work? And what made the model ignore one?

## Extras

- Connect the same server to a commercial coding agent. Everything should work the same way.
- Make the error messages useful _to the model_. And I can't stress this enough. "Not found" tells nothing to the model. "No item with that id, try the list tool first" makes more sense.
- Add a tool that actually writes something with your permission.

## Optional reading

- The two MCP courses in [Anthropic Academy](https://anthropic.skilljar.com/). This is Anthropic's course, but the server protocol is the same, and it's free.
