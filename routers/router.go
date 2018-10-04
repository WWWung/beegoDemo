package routers

import (
	"test/controllers"

	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

func init() {
	//	验证码
	beego.Handler("/captcha/*.png", captcha.Server(240, 70))
	beego.Router("/captcha", &controllers.CaptchaController{}, "post:ReloadCaptcha")

	//	上传文件
	beego.Router("upload", &controllers.FileController{}, "post:UploadFile")

	//	首页
	beego.Router("/", &controllers.MainController{})

	//	后台管理页面
	beego.Router("/admin", &controllers.AdminController{})

	//	后台-产品中心
	beego.Router("/admin/products", &controllers.ProductsController{})
	beego.Router("/admin/productsList", &controllers.ProductsController{}, "post:ProductsList")
	beego.Router("/admin/products.api", &controllers.ProductsController{}, "post:APIhandler")

	//	后台-产品详细页面
	beego.Router("/admin/productDetail", &controllers.ProductsDetailController{})

	//	后台管理页面登录
	beego.Router("/admin/login", &controllers.AdminLoginController{})
	beego.Router("/admin/loginin", &controllers.AdminLoginController{}, "post:Loginin")

	//	后台管理页面注册
	beego.Router("/admin/register", &controllers.AdminRegisController{})
	beego.Router("/admin/regis", &controllers.AdminRegisController{}, "post:Regis")
}
