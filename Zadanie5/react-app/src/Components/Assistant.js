import axios from "axios";
import {useEffect, useState} from "react";

const greetings = ["Welcome!", "What can I help you with?",
    "What can I do for you?", "Hello!", "Hi! What's up?"];

const endings = ["Thank you for your time.", "Hope it helped. Bye",
    "Don't hesitate to ask more", "Goodbye.", "Can I help with anything else?"]

const prompt_context = "You are running online shop. Come up with answer " +
    "matching imaginary data. Don't provide any additional information. " +
    "I prefer short answers."

function Assistant(){
    const [answer, setAnswer] = useState("...");
    const [greeting, setGreeting] = useState();
    useEffect(() => {
        setGreeting(getRandomArrayItem(greetings))
    }, []);
    async function handlePrompt(e) {
        e.preventDefault();
        setAnswer("Waiting for answer...");
        const url = "http://localhost:8000/ask";
        const form = e.target;
        const prompt = form.elements.prompt.value;
        try {
            const res = await axios.post(url, {
                "prompt": prompt_context+" "+prompt
            })
            setAnswer(res.data+" "+getRandomArrayItem(endings));
        } catch(err) {
            setAnswer("Got error: "+err);
        }
    }

    function getRandomArrayItem(a){
        return a[Math.floor(Math.random()*a.length)];
    }

    return(
        <div>
            <h1>Ask assistant about our store or products</h1>
            <p>{greeting}</p>
            <form onSubmit={handlePrompt}>
                <div>
                    <label id="prompt">Put your prompt here:</label>
                    <input type="text" name="prompt"></input>
                </div>

                <button type="submit">Send</button>
            </form>
            <h2>Answer from assistant:</h2>
            <p>{answer}</p>
        </div>
    )
}

export default Assistant