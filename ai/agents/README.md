# Introduction to Agents

<!-- #TODO: YouTube link Agents -->

In this video guide we discuss what an AI agent is from a DevOps, SRE & Platform Engineering perspective. </br>

Engineers will play a key role in:

* Hosting Agents in cloud, Clusters, Local, etc
* Understand traffic routing to agents
* Securing Agentic systems

In this guide we learn about:

* What an agent is
* Different forms of AI Agents
* Context engineering
* Agents.md & System Instructions
* Subagents
* Agent Skills

## Agents.md

AI Agents usually follow their own System Instructions depending on the Agent type. For example Gemini looks for `GEMINI.md` at the root of the repo. Claude looks for `CLAUDE.md`

There is a open standard for [AGENTS.md](https://agents.md/). </br>

Our example [AGENTS.md](./AGENTS.md) can be copied to the root of this repo to demonstrate a system instruction for an agent using a tool like OpenCode. </br>

For a deeper dive in setting up agents for each AI (Gemini,Claude,Copilot,etc), checkout the playlist [here](https://www.youtube.com/playlist?list=PLHq1uqvAteVtCsuixHwiNPQGZ1x5UVLo9)

## Subagents

Agent usually have one contenxt window. </br>
Subagents provide capability to have dedicated agents for each specilized task with their own context windows. </br> For example, you can have a dedicated agent for troubleshooting or debugging, vs Code Review. This means you can keep the context window of the main agent free. </br> 

Our example [subagent](./k8s_expert.md) can be copied into the subagent directory for OpenCode which is `.opencode/agents/k8s_expert/k8s_expert.md` and our OpenCode terminal will pick it up on restart. </br>

For a deeper dive in setting up subagents for each AI (Gemini,Claude,Copilot,etc), checkout the playlist [here](https://www.youtube.com/playlist?list=PLHq1uqvAteVtCsuixHwiNPQGZ1x5UVLo9)


## Agent Skills

If you look at Agents, you will notice the markdown file may become quite large. </br>
We have a lot of commands and scripts inside that file. 

Agent Skills allow us to take this to the next level. Instead of having all the "brains" in the Agents.md file, we can break up functionality and place them in skills. </br>

In this video, we'll discuss why skills are so important and how to use them </br>

There is an open standard written for [Agent Skills](https://agentskills.io/home) as well. </br>

Our example skills folder can be copied to the dedicated `.agents/skills/<name-of-skills>` folder where our OpenCode terminal will pick it up. 

The main thing you will notice once we use skills, is that the Agents instruction file will become a lot leaner and we can essentially move all the "how to"'s out and into scripts. </br>

The instructions then just clearly point to skills that the agent can execute. </br>

We can summarise it as: 

```markdown
---
name: k8s_expert
description: A specialized kubernetes agent specialised in local cluster development and troubleshooting.
---

You help engineers provision, inspect, and debug Kubernetes clusters running on their local machine.

## Persona & Behaviour

- Speak like a senior platform engineer: direct, precise, no hand-holding
- Always confirm tool availability before running commands
- Prefer `kind` for local clusters unless the user specifies otherwise
- When creating clusters, always confirm the cluster name and Kubernetes version first
- When troubleshooting, always check events before logs

## Available Skills

* Provision test k8s infrastructure for testing technical guides using provided available skill

```

We can proceed to use the skill by either referring to it in a prompt or depending on the AI app we are using, for example we can use `/skills` in OpenCode to execute a skill. </br>