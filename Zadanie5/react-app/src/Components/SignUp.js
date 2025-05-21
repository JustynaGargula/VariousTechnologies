import {useNavigate} from "react-router-dom";
import axios from "axios";

function SignUp(){
    const navigate = useNavigate();

    async function handleSignUp(e){
        e.preventDefault()
        const url = "http://localhost:1323/signup";
        const form = e.target;
        const name = form.elements.username.value;
        const password = form.elements.password.value;
        const email = form.elements.email.value;

        try {
            const res = await axios.post(url, {
                "Name": name,
                "Email": email,
                "Password": password
            })

            alert(res.data.message)
            navigate("/login")
        } catch (err) {
            const errMessage = err.response?.data?.error || "Sign up failed"
            alert(errMessage)
        }
    }

    return (
        <div>
            <h1>Sign in</h1>
            <form onSubmit={handleSignUp}>
                <div>
                    <label id="username">User name</label>
                    <input type="text" name="username"/>
                </div>
                <div>
                    <label id="email">Email</label>
                    <input type="email" name="email"/>
                </div>
                <div>
                    <label id="password">Password</label>
                    <input type="password" name="password"/>
                </div>
                <button type="submit">Sign in</button>
            </form>
        </div>
    )
}
export default SignUp