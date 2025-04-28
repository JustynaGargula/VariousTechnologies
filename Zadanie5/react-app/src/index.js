import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import Products from './Components/Products'
import Payment from "./Components/Payment";
import Cart from "./Components/Cart";
import {BrowserRouter as Router, Link, Route, Routes} from "react-router-dom";

function Index() {
    return (
        <Router>
            <nav>
                <Link to="/">Products</Link> | <Link to="/cart">Cart</Link> | <Link to="/payment">Payment</Link>
            </nav>

            <Routes>
                <Route path="/" element={<Products />} />
                <Route path="/cart" element={<Cart />} />
                <Route path="/payment" element={<Payment></Payment>}></Route>
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
