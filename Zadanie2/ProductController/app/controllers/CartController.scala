package controllers

import models.{Cart, CartProduct}
import play.api.mvc.{Action, AnyContent, BaseController, ControllerComponents}

import javax.inject.Inject
import scala.collection.mutable.ListBuffer

class CartController @Inject()(val controllerComponents: ControllerComponents) extends BaseController{

  var products: ListBuffer[CartProduct] = ListBuffer(
    CartProduct(1, 1),
    CartProduct(2, 2)
  )
  var cart = Cart(products, 10.98, 1)

  def addProductToCart(id: Int): Action[AnyContent] = Action {
    products.find(cartProduct => cartProduct.productId == id) match {
      case Some(cartProduct) =>
        val index = products.indexWhere(cartProduct => cartProduct.productId == id)
        products.update(index, cartProduct.copy(quantity = cartProduct.quantity + 1))
      case None => products.append(CartProduct(id, 1))
    }
    Ok("Product added to cart")
  }

  def getCart(): Action[AnyContent] = Action {
    Ok(cart.toString)
  }

  def updateProductQuantity(productId: Int, newQuantity:Int): Action[AnyContent] = Action {
    products.indexWhere(product => product.productId == productId) match {
      case -1 => NotFound("Product not found")
      case index =>
        products.update(index, CartProduct(productId, newQuantity))
        Ok(s"Product with id $productId updated")
    }
  }

  def deleteProductFromCart(id: Int): Action[AnyContent] = Action {
    products.indexWhere(product => product.productId == id) match {
      case -1 => NotFound("Product not found")
      case index =>
        products.remove(index)
        Ok(s"Product $id deleted")
    }
  }

}
