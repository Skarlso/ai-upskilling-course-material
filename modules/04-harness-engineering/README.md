# Module 4: Harness engineering

[Back to the course](../../README.md) | Previous: [Context engineering](../03-context-engineering/README.md) | Next: [Tools and protocols](../05-tools-and-protocols/README.md)

Time to give your agent a bit of a rest. We'll take a look at a commercial one with real tasks now on a real repository.

You already know that a model is a tiny part of the whole story. Everything around a model, the thing we are building, is the _harness_: instructions, tools, checks, everything. Birgitta Böckeler splits it into two kinds of things. Guides steer the agent _before_ it acts and sensors check what it did _after_.

This is where our hoarded knowledge gets to be in play! The way you like errors wrapped, functions written, preferences on if vs switch statements, solutions to hard problems, snippets of code, interface design, etc. Written down in a repo where the agent reads it _every time_. That is how it ends up writing the code you would have written. Mostly. Largely.

One rule for this module, and for the rest of your career with these tools: **when the agent repeats a mistake, fix the harness, not the prompt!!!!!!!!!** THIS IS INCREDIBLY IMPORTANT!
Make a rule in `AGENTS.md`. Make a mod/plugin. Write down the code, the context, whatever! Fix. The. Harness!

## Read

[Harness engineering for coding agent users](https://martinfowler.com/articles/harness-engineering.html) by Birgitta Böckeler.

## Build

Pick a real Go repo you work on.

1. **Write the instructions file.** `AGENTS.md`: how to build, test, conventions, and restrictions. Keep it short! Every line costs _context_.
2. **Make the checks one command.** `go vet ./...`, your linter and `go test ./...`. Tell the agent that these tools are the gates for it declaring success on a task.
3. **Do two real tasks** with a commercial coding agent. Just small ones and notice any errors it makes. Any small error is fine.
4. **Fix the harness.** Each time it does something you don't want, don't skip over it! _Ask_ what would have prevented that action! It, or you, will turn it into either a line in the instructions, or a lint rule, or a test, and you'll add that to the harness.

As said, be a good citizen to others. Review the diff yourself before anybody else does, and write the pull request description by hand. It's common courtesy.

## Done when

You have fixed at least one repeated agent mistake by changing the harness, not the prompt.

## Discuss

Which checks were cheapest to add and resulted in the most things fixed? A simple `DO NOT` or a small test or maybe a mod that prevents further errors?

## Extras

- Run a second coding agent on the same harness. See what stayed the same. Most of it should, and that's why it won't matter which agent you are running.
- Create a skill: a small how-to the agent loads only when it is required.
- Start. Hoarding. Knowledge. I cannot express this hard enough.

## Optional reading

- "Hoard things you know how to do" in [Agentic Engineering Patterns](https://simonwillison.net/guides/agentic-engineering-patterns/) by Simon Willison
- [Harness engineering](https://openai.com/index/harness-engineering/) by OpenAI
- [Skill Issue](https://humanlayer.dev/blog/skill-issue-harness-engineering-for-coding-agents) by HumanLayer
- The vendor's own course for whichever coding agent you use
