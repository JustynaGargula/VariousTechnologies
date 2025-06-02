import React, {useEffect, useState} from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import Products from './Components/Products'
import Payment from "./Components/Payment";
import Cart from "./Components/Cart";
import {BrowserRouter as Router, Link, Route, Routes} from "react-router-dom";
import Login from "./Components/Login";
import SignUp from "./Components/SignUp";
import Assistant from "./Components/Assistant";

function Index() {
    const [cartItems, setCartItems] = useState([]);
    //            { id: 1, name: 'Produkt A', quantity: 2, price: 50 },
    //            { id: 2, name: 'Produkt B', quantity: 1, price: 100 },

    const addToCart = (product) => {
        setCartItems(prevItems => [...prevItems, product]);
    };

    useEffect(() => {
        const params = new URLSearchParams(window.location.search);
        const token = params.get("token");
        if(token) {
            localStorage.setItem("jwt", token);
            window.history.replaceState({}, document.title, "/");
        }
    }, []);

    return (
        <Router>
            <nav>
                <Link to="/">Products</Link> | <Link to="/cart">Cart</Link> | <Link to="/payment">Payment</Link>
                | <Link to="/login">Log In</Link> | <Link to="/signup">Sign Up</Link>
                | <Link to="/assistant">Assistant</Link>
            </nav>

            <Routes>
                <Route path="/" element={<Products addToCart={addToCart} />} />
                <Route path="/cart" element={<Cart cartItems={cartItems} />} />
                <Route path="/payment" element={<Payment></Payment>}></Route>
                <Route path="/login" element={<Login></Login>}></Route>
                <Route path="/signup" element={<SignUp></SignUp>}></Route>
                <Route path="/assistant" element={<Assistant></Assistant>}></Route>
            </Routes>
        </Router>
    );
}

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Index />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
