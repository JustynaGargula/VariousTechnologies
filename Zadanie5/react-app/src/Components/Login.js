import {useNavigate} from "react-router-dom";
import axios from "axios";

function Login(){
    const navigate = useNavigate()

    async function handleLogin(e) {
        e.preventDefault()
        const url = "http://localhost:1323/login";
        const form = e.target;
        const name = form.elements.username.value;
        const password = form.elements.password.value;

        try {
            const res = await axios.post(url, {
                "Name": name,
                "Password": password
            })
            console.log("Token: ", res.data.token)

            alert(res.data.message)
            navigate("/")
        } catch (err) {
            const errMessage = err.response?.data?.error || "Login failed"
            alert(errMessage)
        }
    }

    function handleGoogleLogin(){
        window.location.href = "http://localhost:1323/google-login";
    }

    return (
        <div>
            <h1>Log in</h1>
            <form onSubmit={handleLogin}>
                <div>
                    <label id="username">User name</label>
                    <input type="text" name="username"/>
                </div>
                <div>
                    <label id="password">Password</label>
                    <input type="password" name="password"/>
                </div>
                <button type="submit">Log in</button>
            </form>
            <p>or</p>
            <button onClick={handleGoogleLogin}>Login with Google</button>

        </div>
    )
}
export default Login