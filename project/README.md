# Final project: issue in, pull request out

[Back to the course](../README.md) | Previous: [System design](../modules/09-system-design/README.md)

You made it! You came so far, look at what you accomplished! Well done. You deserve a pat on the back.

You went through nine modules and built an agent that loops, has a harness, uses tools, has an eval suite, is sandboxed and has budget tracking. Now, we will put everything together.

We will build a small pipeline that takes one issue on a toy repo and creates a pull request with a solution.

Let me be clear about what this is _not_. We are not building a factory to replace people. What you will find out, very quickly, is how much care and thought goes in at both ends. The issue has to be written _really_ well, with as much detail as you can give it, or the result will be vastly different from what you expect. And at the other end a human still has to look at it and decide if it is even what they wanted. You can play around with this as much as you like and see for yourself how writing the issue will influence the result.

And even then, no matter how well you will write the issue, you will quickly realize that there are a _great many things_ that are never really written down, yet still matter. These are the non-functional requirements that just come with experience and thought. And then you will understand the shortcomings of such a factory, but you will also see the benefits. It can get you started quickly and then you can _iterate_ from there. Scalability, maintainability, observability, correctness, compliance, metrics... there are at least a dozen more of these that you will just never write down in an issue unless you would like to create a 20 something pages long research paper for a green button on a webpage.

## The Pipeline

This is the flow we are aiming for:

1. **Spec goes in.** An issue with detailed acceptance criteria. Use [`templates/project-spec.md`](../templates/project-spec.md).
2. **Implementation.** Your agent does the work, inside the sandbox from module 8.
3. **Gates.** Build, tests and linters HAVE to pass. If they don't, the failure goes back to the agent, _once_.
4. **Holdout scenarios.** Scenarios the agent does not see are run on the result and scored, like in module 7.
5. **Budget.** The run stops when it hits the limit from module 9.
6. **Evidence out.** A pull request that has the plan in it, the gate results, the scenario score and the cost of it all.

At the end, you review the pull request. If it's not good enough, improve the issue. If it keeps making some mistakes that you notice are the same, improve the _harness_. Don't forget to understand that repeated mistakes are the harness's fault!

## Keep it simple

Just keep it simple. No need for many retries, Kubernetes, CI, second agent, etc. The goal here is to keep it small.

## Measure it

Run the pipeline on at least three issues. One issue is not a metric. Write down:

- the holdout scenario score
- the cost per run
- the minutes of _your_ attention each run needed

The last metric is the interesting one.

## Show it off!

When you share it with the group, show a run! Talk about and compare numbers from other pipelines. Talk about what kind of failures you encountered or what hilarious situations the agent got into. It's all part of the process.

## Two other tracks

If you are too busy, just try these two things:

- **Harness study.** Take a real repo and five real issues. Measure a coding agent's pass rate before and after you build a harness for it.
- **Tool study.** Build an MCP server for a real system, plus an eval suite that measures whether agents use it correctly. Pizza Hut has an MCP server... Just saying. The ideas are limitless.

## And after

The pull request your pipeline opens is a draft from a machine. The usual rules still apply! Review it yourself first. Understand it. Test it. And if it goes to another human, write the description by hand! This is _important_. Remember the Agentic Engineer motto. Don't be rude to other people.
