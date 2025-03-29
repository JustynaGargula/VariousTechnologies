import controllers.ProductController
import models.Product

object ProductControllerTest extends App {
  val controller = new ProductController

  println("📌 Wszystkie produkty:")
  println(controller.getAllProducts())

  println("\n📌 Dodawanie produktu:")
  controller.addProduct(Product(3, "Knife", 6.99))
  println(controller.getAllProducts())

  println("\n📌 Pobieranie produktu o ID 2:")
  println(controller.getProductById(2))

  println("\n📌 Zmiana nazwy produktu o ID 1:")
  controller.setProductName(1, "Golden Spoon")
  println(controller.getAllProducts())

  println("\n📌 Usunięcie produktu o ID 2:")
  controller.deleteProduct(2)
  println(controller.getAllProducts())
}
