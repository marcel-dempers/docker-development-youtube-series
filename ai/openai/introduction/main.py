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