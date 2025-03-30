package models

import scala.collection.mutable.ListBuffer

case class Cart(products: ListBuffer[CartProduct], totalPrice: Double, id: Int)