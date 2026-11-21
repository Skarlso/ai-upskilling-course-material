# Agentic Engineering: The Course

## Intro

Hello and welcome to my course on becoming an Agentic Engineer. These are interesting times!
Our jobs are changing. Our world is changing. I'm not going to pretend that everything is good and dandy, but it can be better!
In order to stay sane, we need to change our thinking. We need to change how we approach this new world. And it can be so much better! It can be _interesting_ and challenging! Here's how.

### Shift away focus from technology and coding

For most of us, coding is the part where we came in and had the most fun. And I hear you. Sometimes, I miss coding too. Then I write a couple of functions, loops, whatever, by hand, and I'm okay. Because there were also the boring parts. Creating the 100000th little function, loop, whatever. I'm now glad I don't have to do that. Because there are times when I know how something needs to be solved, I sketch it out on paper, I know the decision, but writing the actual code would take weeks.

Now, it takes 20 minutes instead. And because I seeded my agent with hoarded knowledge of my own (more about this later in the course) it will generally create the same code I would have written. So I didn't just save myself the trouble of typing out stuff, but gave myself the ability and TIME to really think a problem through from beginning to end, because I knew that actually writing it would take mere minutes instead of weeks or months.

### We are builders

Another reason why I do this is because I _love_ and _live_ for `creating` something that people will ENJOY using! People will find it useful. People will be happy when they can do something they really wanted because of a tool that I wrote.

I have had a list of things I would like to do since 2001. It has over 40 items. Guess how many I managed to actually create or test in 20-something years? Three. My ideas either required me to learn JavaScript / Frontend development / Obscure app frameworks or just would have taken an insane amount of time to actually _type_ out.

I'm no longer bound by this. I'm no longer bound by technology. I'm no longer bound by frameworks, SDKs, and coveted knowledge of obscure OS behaviors. I'm free. There are no strings attached to me. I can try out an idea in _minutes_. But! And here is the most important part of this all. I have a general understanding and knowledge of software design that helps me grasp, even if I don't know a certain technology, whether it will scale, work, or crash immediately.

And this is crucial. This is why I think it makes sense, no, it's _required_, to still learn software engineering. Because you need the skills to utilize this technology _efficiently_.

I might not know the Bootstrap framework, but I know enough to understand that a CSS or a JavaScript segment is inefficient enough that the browser will crash. I don't know C, but I know if an eBPF script does something silly, and I can run a hundred tests as well with agents to make sure it works. I know enough about software to design, use and test it even if I don't understand _yet_ the particular technology being utilized.

I will not compromise on quality. And this is _not_ vibe coding. Vibe coding would be to say, meh. I don't care. I. CARE.

And I know. I hear you. A lot of corporations just want to use it to replace people. Evil corporations are evil. They will always be. We can't do anything about that. What we CAN do is build with care. Build with love. And continue to build with quality.

There are only two rules. Never compromise on quality and don't be rude to your fellow reviewers. Think about the problem. Review the solution. And test your code. And write the description by hand. Give the reviewer the courtesy that you reviewed your own code first.

We will refuse to be a [meat proxy](https://meatproxy.me/). When you interact with others, don't make your agent speak for you. That's just rude.

So come, join me in being responsible _Agentic Engineers_ who will NOT compromise on their values. Who will steer the future in the right direction and not let corporations and greed take away what we built for decades. Be proud. And I wish the best to you on this journey.

## The Course

Now, this course is NOT about building LLMs and learning Python. This course is about learning to use agents and agentic coding better. AND! To build up your general software skills and design skills that atrophied over time because you were so focused on coding. Systems thinking, distributed software building, concurrency, multi-tenancy... These are your new mantras. And not, what algorithm to choose or which loop to run.

Those are almost always irrelevant. We will still solve puzzles. Just on the macro-scale and not the micro-scale. Don't get me wrong, we will still read, understand and validate most of the code. We have to. We don't compromise on quality. But with a different mindset and with a different purpose.

But, I digress. Simon Willison has better words than I do here, and we will meet him along the way.

This is a self-paced course meant to be done together with friends. You can also do it all alone, that's fine too. But it's nice to share with someone. Across nine modules you build one coding agent in Go, and learn the details that make agents what they are today: context, harnesses, tools, verification, security and system design. There is a small final project at the end that will put all of this together.

There is no schedule. Go at your own pace, skip the extras if you have to, and pick it up again whenever you have time.

The course is tool-neutral. Use any coding agent and any model provider. It doesn't matter.

### What you leave with

These are optional goals. Feel free to ignore any of these.

- A coding agent you wrote yourself and understand line by line (yes, we'll do this by hand)
- A harness on a real repo, plus checks that the agent runs itself
- An MCP server your agent uses
- An eval suite with a measured pass rate, plus holdout scenarios
- Sandboxing and threat modelling
- A small project software factory where an issue goes in and a pull request comes out

### What you need

You should be comfortable in Go. You do not need a machine learning background.

- Go
- An API key for one model provider
- Access to one commercial coding agent
- A container runtime such as Docker, from module 8

### How it works

There are nine modules and the final project. Do the modules in order, because each one extends the agent from the one before it.

- **Read.** The lore for the module.
- **Build.** The task for the module. ~2 hours or so of diligent work.
- **Done when.** An output that will tell you it's done.
- **Discuss.** Anything worth talking about with the group (or yourself! seriously. talk to yourself.)
- **Extras and optional reading.** Only if you have the time and the interest really.

The done-when line is the minimum. Everything past that is optional.

### Modules

| # | Module | You leave with |
| --- | --- | --- |
| 1 | [Orientation and the loop, part 1](modules/01-loop-part-1/README.md) | An agent, written by you, that can chat and read files |
| 2 | [The loop, part 2](modules/02-loop-part-2/README.md) | An agent that edits files, runs commands and fixes a failing test |
| 3 | [Context engineering](modules/03-context-engineering/README.md) | Full request logs, token counts and an instructions file |
| 4 | [Harness engineering](modules/04-harness-engineering/README.md) | A harness on a real repo, with checks the agent runs itself |
| 5 | [Tools and protocols](modules/05-tools-and-protocols/README.md) | An MCP server that the agent uses |
| 6 | [Verification, part 1](modules/06-verification-part-1/README.md) | A task suite with a measured pass rate |
| 7 | [Verification, part 2](modules/07-verification-part-2/README.md) | Holdout scenarios and a list of your agent's failure types |
| 8 | [Security and sandboxing](modules/08-security-and-sandboxing/README.md) | A sandboxed shell tool and a one-page threat model |
| 9 | [System design](modules/09-system-design/README.md) | Budget caps and a human approval step |
| | [Final project](project/README.md) | Issue in, pull request out |
| | [The fundamentals](fundamentals/README.md) | The skills underneath, as a reading list |

### Doing it together

People will progress at different paces, but there are things that make being in a group worthwhile:

- **Sharing is caring.** A few lines about what you did, what failed, or an interesting thing you discovered. Here is a template if you are introverted: [`templates/module-log.md`](templates/module-log.md).
- **Do show if something is NOT working.** Equally important to figure out why something fails not just when it works.
- **Read each other's agent output.** Review each other's agents' output! This is why we are here after all. Understanding agentic output. Finding patterns in the responses so reviewing becomes a LOT easier.
- **Meet when you have time.** Whenever a couple of people have finished a set of work, discuss it.

### Final project

**Issue in, pull request out.** Build a small pipeline that takes one issue on a toy repo and produces a pull request that solves it without human intervention. This is everything that we learned so far pieced together. And do be mindful here, we aren't building a fully autonomous factory to replace people. Far from it. It requires careful iterations, a very _VERY_ well written detailed issue. As much detail as you can get, to reduce iterative drift and, ultimately, to understand if the result is even what you wanted in the first place.

The pull request contains the plan, the result, a holdout scenario score and how much it cost. Then, a human reviews it and decides if it can be merged. If not, the process is continued.

There are two alternative tracks, a harness study and a tool study, if you are busy. The full brief is in [`project/README.md`](project/README.md).

## Safety

- Run your home-built agent on throwaway repos with throwaway credentials until you have finished module 8.
- Do NOT point a home-built agent at your employer's code without permission.
- Keep secrets out of any logs and transcripts you share.
- Set yourself a spending limit for model usage if you have one. A local model is usually fine for all the tasks we are about to do.

## Repo layout

```
README.md              this file
modules/
  01-loop-part-1/README.md
  02-loop-part-2/README.md
  03-context-engineering/README.md
  04-harness-engineering/README.md
  05-tools-and-protocols/README.md
  06-verification-part-1/README.md
  07-verification-part-2/README.md
  08-security-and-sandboxing/README.md
  09-system-design/README.md
project/README.md      brief, tracks, what to show at the end
templates/
  module-log.md
  threat-model.md
  project-spec.md
toy-repo/              shared Go repo for the verification modules
_config.yml            website config (GitHub Pages)
_layouts/default.html  website page layout
assets/style.css       website styling
```

Your agent lives in your own repo.

## Fundamentals

I also added a list of skills you should do a refresher on! These are the ones that are most important now... Distributed Systems, System Thinking, Observability practices, and a lot more. You can find these under [Fundamentals README](./fundamentals/README.md).

## Credits

This course is a reading order and a set of exercises built on other people's freely available work. The main sources:

- [Agentic Engineering Patterns](https://simonwillison.net/guides/agentic-engineering-patterns/), [The Five Levels](https://simonwillison.net/2026/Jan/28/the-five-levels/) and [The lethal trifecta](https://simonwillison.net/2025/Jun/16/the-lethal-trifecta/) by Simon Willison
- [How to Build an Agent](https://ampcode.com/how-to-build-an-agent) by Thorsten Ball
- [how-to-build-a-coding-agent](https://github.com/ghuntley/how-to-build-a-coding-agent) by Geoffrey Huntley
- [Building effective agents](https://www.anthropic.com/engineering/building-effective-agents) by Anthropic
- [12-Factor Agents](https://github.com/humanlayer/12-factor-agents) and [Skill Issue](https://humanlayer.dev/blog/skill-issue-harness-engineering-for-coding-agents) by HumanLayer
- [Harness engineering for coding agent users](https://martinfowler.com/articles/harness-engineering.html) by Birgitta Böckeler
- [Harness engineering](https://openai.com/index/harness-engineering/) by OpenAI
- [Model Context Protocol](https://modelcontextprotocol.io/) and its [Go SDK](https://github.com/modelcontextprotocol/go-sdk)
- [The StrongDM Software Factory](https://factory.strongdm.ai/)
- The [evals email series](https://ai.hamel.dev/eval-course) and the book [Evals for AI Engineers](https://www.oreilly.com/library/view/evals-for-ai/9798341660717/) by Hamel Husain and Shreya Shankar
- [Open Policy Agent](https://www.openpolicyagent.org/)

If a link is broken, please open an issue or a pull request.
