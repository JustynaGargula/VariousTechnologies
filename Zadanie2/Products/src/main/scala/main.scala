import scala.collection.mutable.ListBuffer

@main
def main(): Unit = {
  println("Hello world!")
}

case class Product(id: Int, name: String, price: Double)

var products = ListBuffer(
  Product(1, "Spoon", 4.99),
  Product(2, "Fork", 5.99)
)