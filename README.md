# Infra-as-LLM

## Background

This project started as an idea to automate my own job. As a cloud architect today I am responsible for taking enterprise problems and translating them into engineering solutions with the help of engineering teams.

I am usually researching the latest cloud trends, best practices, and cost saving opportunities. 

What if an LLM agent (or agents) could do my job? The research is all online available, and an LLM can probably research faster than me. All major cloud providers have SDKs and APIs that can be integrated into the agent’s workflow. In theory, it should be quite easy.

## Challenges

1. The biggest challenge is validation, and this requires pre-training

`How do I know that the solution the AI has given me, is the best, most cost-effective, and resilient architecture for the problem at hand?`


2. The next challenge is working with multiple agents

`How can I efficiently create a set of instructions from a user prompt, and delegate the instructions to different agents, and then have them come back together at the end to put together the entire solution?`


3. The last challenge, and perhaps the easiest to solve is integration

`How do I integrate my agents to effectively take actions using the cloud APIs? For example, if my agent decides to build a lambda function for my solution, how can I make sure that it knows which api call(s) to make?`

