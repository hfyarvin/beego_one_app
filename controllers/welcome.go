package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

// WelcomeController operations for Welcome
type WelcomeController struct {
	beego.Controller
}

// URLMapping ...
func (c *WelcomeController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Welcome
// @Param	body		body 	models.Welcome	true		"body for Welcome content"
// @Success 201 {object} models.Welcome
// @Failure 403 body is empty
// @router / [post]
func (c *WelcomeController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Welcome by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Welcome
// @Failure 403 :id is empty
// @router /:id [get]
func (c *WelcomeController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Welcome
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Welcome
// @Failure 403
// @router / [get]
func (c *WelcomeController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Welcome
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Welcome	true		"body for Welcome content"
// @Success 200 {object} models.Welcome
// @Failure 403 :id is not int
// @router /:id [put]
func (c *WelcomeController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Welcome
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *WelcomeController) Delete() {

}

//GetPage

func (c *WelcomeController) WelcomeIndex() {
	fmt.Println("Welcome index page.")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Name"] = beego.AppConfig.String("name")
	c.TplName = "welcome/welcome.tpl" //default file_folder is views
}

func (c *WelcomeController) WelcomeMyName() {
	c.Ctx.Output.Body([]byte(fmt.Sprintf("welcome %s", beego.AppConfig.String("name"))))
}
