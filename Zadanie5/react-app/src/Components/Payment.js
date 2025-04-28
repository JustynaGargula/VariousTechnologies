import {useState} from "react";
import axios from "axios";
import {useLocation} from "react-router-dom";


function Payment() {
    const url = "http://localhost:1323/payment";
    const [formData, setFormData] = useState({
        amount: "",
        cardNumber: "",
        expiryDate: "",
        cvv: ""
    });

    const location = useLocation();
    const { totalAmount } = location.state || { totalAmount: 0 };

    const sendPayment = async(err) => {
        err.preventDefault();
        // console.log("Got data from form: ", formData)
        try {
            const res = await axios.post(url, formData);
            console.log("Payment successful:", res.data);
            alert("Payment successful!");
        }
        catch (error) {
            console.error("Payment error:", error);
            alert("Payment error");
        }
    }

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData(prev => ({ ...prev, [name]: value }));
    };

    return (
        <div>
            <h1>Payment page</h1>
            <p>Amount to pay: {totalAmount}</p>
            <form onSubmit={sendPayment}>
                <div>
                    <label>Card number</label>
                    <input type="text" name="cardNumber" value={formData.cardNumber} required  onChange={handleChange}/>
                </div>
                <div>
                    <label>Expiry date</label>
                    <input type="date" name="expiryDate" value={formData.expiryDate} required  onChange={handleChange}/>
                </div>
                <div>
                    <label>CVV</label>
                    <input type="text" name="cvv" value={formData.cvv} required  onChange={handleChange}/>
                </div>
                <button type="submit">Pay</button>
            </form>
        </div>
    )

}

export default Payment;