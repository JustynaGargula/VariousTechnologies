import {useState} from "react";
import axios from "axios";

function Cart() {
    const url = "http://localhost:1323/cart";

    const [cartItems, setCartItems] = useState([
        { id: 1, name: 'Produkt A', quantity: 2, price: 50 },
        { id: 2, name: 'Produkt B', quantity: 1, price: 100 },
    ]);

    const totalAmount = cartItems.reduce((sum, item) => sum + item.price * item.quantity, 0);


    const sendCart = async (err) => {
        try {
            const res = await axios.post(url, cartItems)
            console.log("Successfully sent cart data: ", res.data);
            alert("Successfully sent cart data");
        }
        catch(error) {
            console.error("Sending cart data error:", error);
            alert("Failed to send cart data");
        }

    };

    return (
        <div>
            <h1>Cart page</h1>
            <h2>Items in cart:</h2>
            <ul>
                {cartItems.map(item => (
                    <li id={item.id}>{item.name}, quantity: {item.quantity}, price: {item.price}</li>
                ))}
            </ul>
            <p>Total: {totalAmount}</p>
            <button onClick={sendCart}>Go to payment</button>
        </div>
    )
}

export default Cart;