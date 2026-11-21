# The fundamentals

[Back to the course](../README.md)

In the intro I told you that there are skills that atrophied over time when you were focused too deep in the guts and bowels of services, controllers, servers, clients, and code in general.

Systems thinking. Distributed systems. Design. This page is here so you might make an effort now to refocus and relearn some of the lost skills you had or never had in the first place.

It is _important_ to learn these things. Why? Because writing code got cheap. Deciding what to build, and proving that what you got is right, did not! Everything on this page is about those two things.

This is not a module. I didn't include any specific projects here. It's completely at your discretion. It's an opinionated reading list basically. Pick a topic that you would like to pursue and strengthen or that pairs well with the module you are on currently.

**If you only want a couple, focus on these:** testing theory, distributed systems and software design. Best bang for the buck.

## Software design

**Pairs with:** [module 4](../modules/04-harness-engineering/README.md)

Agents do their best work in code with small, clear boundaries. The context window is limited, therefore it's important to define it as well as possible. When you designed something well for your colleagues because you thought that they will understand it better, you need to do the same thing now, just for your harness.

- _A Philosophy of Software Design_ by John Ousterhout. Short and it's the best single read on deep modules.
- [_Building Evolutionary Architectures 2nd Edition_](https://www.oreilly.com/library/view/building-evolutionary-architectures/9781492097532/) by Neal Ford, Rebecca Parsons, Patrick Kua and Pramod Sadalage.

## Technical writing

**Pairs with:** [module 4](../modules/04-harness-engineering/README.md), [module 5](../modules/05-tools-and-protocols/README.md) and the [final project](../project/README.md)

Now, more than ever, it's important to learn how to write technical lingo correctly! I know, writing documentation is usually the _boring_ part of software engineering. I know. But it's a must. And since LLMs are just not that good, and arguably will never be good, at subjective writing, it's imperative that WE DO learn how to write well...

Besides. You need to be able to formulate _correctly_ what you want. And that requires precise formulation.

- Google's free [Technical Writing courses](https://developers.google.com/tech-writing). Start with Technical Writing One.

## Testing theory

**Pairs with:** [module 6](../modules/06-verification-part-1/README.md)

The more tests you have the better. It's that simple. If you tell the agent that it cannot touch the tests and that tests are what prove the functionality, it's going to be a bit harder for the agent to cheat.

Property-based testing and fuzzing help you to tell the agent "this must always be true".

- [Fuzzing in Go](https://go.dev/doc/security/fuzz/).
- A property-testing library such as `pgregory.net/rapid`.

## Statistics for evals

**Pairs with:** [module 6](../modules/06-verification-part-1/README.md) and [module 7](../modules/07-verification-part-2/README.md)

The agent doesn't give the same answer twice, so a result is not a boolean but a proportion with an error.
But, "11 out of 12" or "12 out of 15" are numbers that are _definitive_. You don't need much math, but you do _want_ this much at least.

- _Statistics Done Wrong_ by Alex Reinhart. Short and funny and interesting.
- [Adding Error Bars to Evals](https://arxiv.org/pdf/2411.00640) from Anthropic.
- [Evals for AI Engineers](https://www.oreilly.com/library/view/evals-for-ai/9798341660717/) by Shreya Shankar and Hamel Husain covers this for agents specifically.

## Threat modelling and isolation

**Pairs with:** [module 8](../modules/08-security-and-sandboxing/README.md)

This one is really self-explanatory why you need to understand boundaries, security and threat modelling.

- [Threat Modeling: Designing for Security](https://www.amazon.co.uk/Threat-Modeling-Designing-Adam-Shostack/dp/1118809998?th=1) by Adam Shostack is a very old one from 2014.
- [Threat Modeling: Designing for Security in an AI World](https://shostack.org/books/threat-modeling-book-2nd-edition) by Adam Shostack as well is coming out in 2027 and is the second edition of this. Highly recommend it when it comes out.

## Distributed systems

**Pairs with:** [module 9](../modules/09-system-design/README.md) and the [final project](../project/README.md)

Understanding how, when, if to scale something is one of the most crucial elements of designing good software. I can't stress enough how important it is to understand how to design something that won't collapse when more than five people are trying to use it at the same time.

- [Designing Data-Intensive Applications, 2nd edition](https://www.oreilly.com/library/view/designing-data-intensive-applications/9781098119058/) by Martin Kleppmann and Chris Riccomini. "The Book".
- _Release It!_, 2nd edition, by Michael Nygard. Timeouts, circuit breakers and bulkheads, and more.
- [Designing Distributed Systems 2nd Edition](https://learning.oreilly.com/library/view/designing-distributed-systems/9781098156343/) by Brendan Burns is a fantastic primer for everything about distribution and it's _only_ 220 pages compared to the Data-Intensive Applications, which is almost 700!

## Concurrency

**Pairs with:** [module 6](../modules/06-verification-part-1/README.md) and [module 9](../modules/09-system-design/README.md)

Concurrency is equally important. Most of the time though this will be very low level engineering. So I'm not sure what to recommend here. If anyone has good suggestions I can add here, please feel free to open an issue.

- _Concurrency in Go_ by Katherine Cox-Buday.

## Observability

**Pairs with:** [module 3](../modules/03-context-engineering/README.md) and the [final project](../project/README.md)

Another one of the cornerstones of good software is that you have the ability to observe what's happening with it internally and not just externally on the output. Metrics are crucial! Understanding good metrics is god tier.

- OpenTelemetry's tracing basics.
- [Observability Engineering, 2nd edition](https://www.oreilly.com/library/view/-/9781098179915/). It has new material on LLMs.

## Systems thinking

**Pairs with:** the [final project](../project/README.md)

When writing code stops being the slow part, something else becomes the slow part. Usually review. Feedback loops, queues and bottlenecks tell you where your little factory is going to stall before it does.

- _Thinking in Systems_ by Donella Meadows. A slim book with no code in it. Perfect for travelling.
- _The Goal_ by Eliyahu Goldratt. A novel about a factory. Yes, really.
- [Systems Thinking for Agentic AI](https://www.amazon.de/-/en/Systems-Thinking-Agentic-guardrails-observability/dp/1808828496), a _very_ recent book on applying systems thinking for agents.

## Reading code

**Pairs with:** [module 4](../modules/04-harness-engineering/README.md)

You are going to read far more code than you write from now on. It is a skill, and it gets better with practice like any other.

- The chapters on understanding code in [Agentic Engineering Patterns](https://simonwillison.net/guides/agentic-engineering-patterns/) by Simon Willison.

## For later: specifications

**Pairs with:** nothing yet.

This one is an interesting one. I'm debating still on the usefulness as this is very much on the fringe of things. But if you DO have the time, I would recommend delving into TLA+ just for the sake of understanding something completely new and different.

- [Learn TLA+](https://learntla.com/) by Hillel Wayne. Start with PlusCal, which reads a lot more like code.
- A good exercise once you have done the project: try modelling your pipeline's states.
