package controllers

import models.Category
import play.api.mvc.{Action, AnyContent, BaseController, ControllerComponents}

import javax.inject.Inject
import scala.collection.mutable.ListBuffer

class CategoryController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  private val categories: ListBuffer[Category] = ListBuffer(
    Category(1, "Kitchenware", "Items for the kitchen"),
    Category(2, "Electronics", "Electronic devices and accessories")
  )

  def getAllCategories: Action[AnyContent] = Action {
    Ok(categories.toString)
  }

  def getCategoryById(id: Int): Action[AnyContent] = Action {
    categories.find(_.id == id) match {
      case Some(category) => Ok(category.toString)
      case None => NotFound("Category not found")
    }
  }

  def addCategory(id: Int, name: String, description: String): Action[AnyContent] = Action {
    categories.append(Category(id, name, description))
    Created(s"Category '$name' added")
  }

  def updateCategory(id: Int, newName: String, newDescription: String): Action[AnyContent] = Action {
    categories.indexWhere(_.id == id) match {
      case -1 => NotFound("Category not found")
      case index =>
        categories.update(index, Category(id, newName, newDescription))
        Ok(s"Category $id updated")
    }
  }

  def deleteCategory(id: Int): Action[AnyContent] = Action {
    categories.indexWhere(_.id == id) match {
      case -1 => NotFound("Category not found")
      case index =>
        categories.remove(index)
        Ok(s"Category $id deleted")
    }
  }
}
