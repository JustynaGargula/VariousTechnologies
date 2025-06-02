import requests
from fastapi import FastAPI, Body
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel

app = FastAPI()
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["GET", "POST"],
    allow_headers=["*"]
)


class PromptRequest(BaseModel):
    prompt: str


@app.post("/ask")
def ask_assistant(prompt_req: PromptRequest = Body(...)):
    llm_url = "http://localhost:11434/api/generate"

    payload = {
        "model": "llama3.2",
        "prompt": prompt_req.prompt,
        "stream": False
    }
    response = requests.post(url=llm_url, json=payload)

    if response.status_code == 200:
        response_data = response.json()
        return f"Assistant response: {response_data['response']}"
    else:
        return f"Error: {response.text}"
