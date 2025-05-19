import requests

llamaUrl = "http://localhost:11434/api/generate"

examplePayload = {
    "model": "llama3",
    "prompt": "Tell me something about this shop",
    "stream": False
}

response = requests.post(url=llamaUrl, json=examplePayload)

responseData = response.json()
print("Assistant response: ", responseData["response"])

# if __name__ == '__main__':
#     print_hi('PyCharm')
