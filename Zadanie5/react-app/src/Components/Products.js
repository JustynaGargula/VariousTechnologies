import {useEffect, useState} from "react";

function Products({addToCart}) {

    const [products, setProducts] = useState([])
    const url = "http://localhost:1323/products";
    useEffect(() => {
        fetch(url).then((response) => response.json())
            .then((data) => setProducts(data))
            .catch( (error) => {
                console.log("Couldn't fetch products. Got error: ", error);
            })
    }, []);

    function addProductToCart(product){
        addToCart(product)
    }

    return (
        <div>
            <h1>Product page</h1>
            <h2>Available products</h2>
            <ul>
                { products.map((product) => (
                    <li key={product.id}>
                        {product.name}
                        <ul>
                            <li>Price: {product.price}</li>
                            <li>Description: {product.description}</li>
                        </ul>
                        <button onClick={addProductToCart(product)}>Add to cart</button>
                    </li>
                ))}
            </ul>
        </div>
    )
}

export default Products;