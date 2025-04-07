function Product() {
    let Product = {
        id: 1,
        name: "Book",
        description: "Self development book"
    }
    return (
        <div>
            <h1>Product page</h1>
            <h2>Available products</h2>
            <ul>
                <li key={Product.id}>
                    <ul>
                        <li>Name: {Product.name}</li>
                        <li>Description: {Product.description}</li>
                    </ul>
                </li>
            </ul>
        </div>
    )
}