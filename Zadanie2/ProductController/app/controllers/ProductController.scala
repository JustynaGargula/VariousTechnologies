package controllers
import scala.collection.mutable.ListBuffer
import models.Product

class ProductController {

  var products = ListBuffer(
    Product(1, "Spoon", 4.99),
    Product(2, "Fork", 5.99)
  )

  def addProduct(product: Product) = products.append(product)
  def getProductById(id: Int) = products.find(product => product.id == id)
  def getAllProducts() = products
  def setProductName(id:Int, newName: String) = getProductById(id).foreach(product => product.name = newName)
  def setProductPrice(id: Int, newPrice: Double) = getProductById(id).foreach(product => product.price = newPrice)
  def deleteProduct(id: Int) = getProductById(id).foreach(product => products -= product)


}
