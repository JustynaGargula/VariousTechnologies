package controllers
import scala.collection.mutable.ListBuffer
import models.Product
import play.api.mvc.{Action, AnyContent, BaseController, ControllerComponents}
import play.api.mvc.Results.{Ok, Created, NotFound}

import javax.inject.Inject

class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController{

  var products = ListBuffer(
    Product(1, "Spoon", 4.99),
    Product(2, "Fork", 5.99)
  )

  def addProduct(id: Int, name: String, price: Double): Action[AnyContent] = Action{
    products.append(Product(id, name, price))
    Created(s"Product $name added")
  }
  def getProductById(id: Int): Action[AnyContent] = Action {
    products.find(product => product.id == id) match {
      case Some(product) => Ok(product.toString)
      case None => NotFound("Product not found")
    }
  }
  def getAllProducts(): Action[AnyContent] = Action {
    Ok(products.toString)
  }
  def updateProduct(id: Int, newName: String, newPrice: Double): Action[AnyContent] = Action {
    products.indexWhere(product => product.id == id) match {
      case -1 => NotFound("Product not found")
      case index =>
        products.update(index, Product(id, newName, newPrice))
        Ok(s"Product $id updated")
    }
  }

   def deleteProduct(id: Int): Action[AnyContent] = Action {
    products.indexWhere(product => product.id == id) match {
      case -1 => NotFound("Product not found")
      case index =>
        products.remove(index)
        Ok(s"Product $id deleted")
    }
  }


}
