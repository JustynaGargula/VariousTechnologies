object Data {
    var categories = mutableListOf<Category>(
        Category(1,"electronics", "Electronic things."),
        Category(2,"cosmetics", "Products for hygiene and skincare"),
        Category(3,"kitchenware", "Cutlery and other things for kitchen"));

    var products = mutableListOf<Product>(
        Product(1, "Fork", 7.99, categories[2]),
        Product(2, "Spoon", 8.99, categories[2]),
        Product(3, "Soap", 12.99, categories[1])
    )

    fun getAllCategories(): String{
        return categories.joinToString("\n") { category ->
            "**${category.name}** (ID: ${category.id})\n\t${category.description ?: "_brak opisu_"}"
        }
    }

    fun getProductsForCategory(categoryId: Long): String{
        var categoryProducts = mutableListOf<Product>()
        for (product in products){
            if(product.category.id == categoryId){
                categoryProducts.add(product)
            }
        }
        if (categoryProducts.size>0){
            return categoryProducts.joinToString("\n") { product: Product ->
                "**${product.name}** (ID: ${product.id}), price: ${product.price}"
            }
        }
        else {
            return "No products in this category"
        }

    }
}