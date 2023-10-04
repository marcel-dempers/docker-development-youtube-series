# Introduction to Open AI

## Overview 

What is [Open AI](https://openai.com/) ? 

* Research company on AI development
* Builds and provides models
* Builds and provides a standard protocol for using AI

What is a model ?

I see a model as a language super database. </br>
Instead of writing a query, that is slow to query a traditional database like SQL, you can throw a question at a model and it gives you an answer really fast </br>

Model examples:
* GPT 3.5 
* GPT 4

## Getting started 

The best way to get started and to understand OpenAI, is to learn hands on

* Create an OpenAI account [here](https://openai.com/) 

## Chat GPT 

Here you can find the link to [ChatGPT](https://chat.openai.com/)

## Open AI Playground 

Here you can find the link to the [OpenAI Playground](https://platform.openai.com/playground)

## Build an AI powered app

We can start with a `main.py` that reads a message

```
import sys

message = sys.argv[0]

```
Then we will need the code from the Open AI playground and add it to our `main.py`. </br>
Move the `import` statements to the top </br>

Once you have tidied up everything, you can get the response message from the AI:

```
responseMessage = response.choices[0].message.content
```

Let's build our app

```
cd ai\openai\introduction
docker build . -t ai-app
```

Set my OpenAI API key

```
$ENV:OPENAI_API_KEY=""
```

Run our AI App:

```
docker run -it -e OPENAI_API_KEY=$ENV:OPENAI_API_KEY ai-app
```

When we run the app, notice it has no concept of memory. </br>
The playground works because it keeps track of all the user and AI messages and keeps appending new messages to it </br>
So it can track the conversation.

Let's keep track of messages, by writing it to a local file </br>
We will also take the system message out and keep it as a constant in our code </br>

Full example:

```
import sys
import os
import json
import openai

openai.api_key = os.getenv("OPENAI_API_KEY")

#read the incoming message
message = sys.argv[1]
user_message = {
  "role" : "user",
  "content" : message
}

systemMessage = {
  "role": "system",
  "content": "You are a kubernetes exper that can assist developers with troubleshooting deployments\n\nTo help the developer you will need to know the namespaces as well as the pod name. Ask for missing information\n\nGenerate a command to help the developer surface logs or information\n"
}

# read the cached user messages if there are any
userMessages = []
if os.path.isfile("messages.json"):
  with open('messages.json', newline='') as messagesFile:
    data = messagesFile.read()
    userMessages = json.loads(data)

# add the new message to it and update the cached messages
userMessages.append(user_message)
with open('messages.json', 'w', newline='') as messagesFile:
  msgJSON = json.dumps(userMessages)
  messagesFile.write(msgJSON)
  print(msgJSON)

messages = []
messages.append(systemMessage)
messages.extend(userMessages)

response = openai.ChatCompletion.create(
  model="gpt-3.5-turbo",
  messages=messages,
  temperature=1,
  max_tokens=256,
  top_p=1,
  frequency_penalty=0,
  presence_penalty=0
)

responseMessage = response.choices[0].message.content
print(responseMessage)

```

Now we can mount our volume so we persist the cache of messages 

```
docker run -it -e OPENAI_API_KEY=$ENV:OPENAI_API_KEY -v ${PWD}:/app ai-app "can you help me with my deployment?"
Of course! I'd be happy to help with your deployment. Could you please provide me with the namespace and the name of the pod you're encountering issues with?

docker run -it -e OPENAI_API_KEY=$ENV:OPENAI_API_KEY -v ${PWD}:/app ai-app "my pod is pod-123" 
Sure, I can help you with your deployment. Can you please provide me with the namespace in which the pod is running?

docker run -it -e OPENAI_API_KEY=$ENV:OPENAI_API_KEY -v ${PWD}:/app ai-app "its in the products namespace"
Great! To surface the logs for the pod "pod-123" in the "products" namespace, you can use the following command:

```shell
kubectl logs -n products pod-123
```

This command will retrieve the logs for the specified pod in the given namespace. Make sure you have the necessary permissions to access the namespace.
```