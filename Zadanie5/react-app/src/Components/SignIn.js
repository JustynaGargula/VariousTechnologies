import {useNavigate} from "react-router-dom";

function SignIn(){
    var navigate = useNavigate()
    function handleSignIn(){
        alert("Successfully signed in!")
        navigate("/login")
    }

    return (
        <div>
            <h1>Sign in</h1>
            <form action={handleSignIn}>
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
export default SignIn