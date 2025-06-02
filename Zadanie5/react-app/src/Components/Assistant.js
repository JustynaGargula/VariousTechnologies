import axios from "axios";
import {useState} from "react";

function Assistant(){
    const [answer, setAnswer] = useState("...")
    async function handlePrompt(e) {
        e.preventDefault()
        setAnswer("Waiting for answer...")
        const url = "http://localhost:8000/ask"
        const form = e.target;
        const prompt = form.elements.prompt.value
        try {
            const res = await axios.post(url, {
                "prompt": prompt
            })
            setAnswer(res.data)
        } catch(err) {
            setAnswer("Got error: "+err)
        }
    }

    return(
        <div>
            <h1>Ask assistant about our store or products</h1>
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