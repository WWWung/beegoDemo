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
	beego.Router("/", &controllers.HomeController{})

	//	关于我们
	beego.Router("/summary", &controllers.SummaryController{}, "get:GetSummaryPage")

	//	新闻中心
	beego.Router("/news", &controllers.NewsController{}, "get:GetNewsPage")
	beego.Router("/newsDetail", &controllers.NewsController{}, "get:GetDetailPage")

	//	案例展示
	beego.Router("/customer", &controllers.CustomerController{}, "get:GetCustomerPage")
	beego.Router("/customerDetail", &controllers.CustomerController{}, "get:GetDetailPage")

	//	产品中心
	beego.Router("/products", &controllers.ProductsController{})
	beego.Router("/productDetail", &controllers.ProductsController{}, "get:GetProductDetailPage")

	//	账号管理
	beego.Router("/register", &controllers.UserController{}, "get:GetRegisterPage")
	beego.Router("/login", &controllers.UserController{}, "get:GetLoginPage")
	beego.Router("/user.api", &controllers.UserController{}, "*:APIhandler")
	beego.Router("/admin/users", &controllers.UserController{}, "get:GetAdminUsersPage")

	//	下载中心
	beego.Router("/downloadCenter", &controllers.DownloadCenterController{})
	beego.Router("/downloadCenter/preview", &controllers.DownloadCenterController{}, "*:GetPriviewPage")

	//留言
	beego.Router("/contact", &controllers.ContactController{})
	beego.Router("/contact.api", &controllers.ContactController{}, "*:APIhandler")
	beego.Router("/admin/contact", &controllers.ContactController{}, "get:GetAdminContactPage")

	//	后台管理页面
	beego.Router("/admin", &controllers.AdminController{})

	//	后台-产品中心
	beego.Router("/admin/products", &controllers.ProductsController{}, "get:GetAdminPage")
	beego.Router("/products.api", &controllers.ProductsController{}, "*:APIhandler")

	//	后台-产品详细页面
	beego.Router("/admin/productDetail", &controllers.ProductsDetailController{})

	//	后台-文件管理
	beego.Router("/admin/filesManage", &controllers.FilesManageController{})
	beego.Router("/admin/fileDetail", &controllers.FilesManageController{}, "get:GetFileDetailPage")
	beego.Router("/filesManage.api", &controllers.FilesManageController{}, "*:APIhandler")

	//	后台-简介（企业简介、系统简介）
	beego.Router("/admin/summary", &controllers.SummaryController{})
	beego.Router("/admin/summaryDetail", &controllers.SummaryController{}, "get:GetSummaryDetailPage")
	beego.Router("/summary.api", &controllers.SummaryController{}, "*:APIhandler")

	//	后台-新闻中心
	beego.Router("/admin/news", &controllers.NewsController{})
	beego.Router("/admin/newsDetail", &controllers.NewsController{}, "get:GetNewsDetailPage")
	beego.Router("/news.api", &controllers.NewsController{}, "*:APIhandler")

	//	后台-案例展示
	beego.Router("/admin/customer", &controllers.CustomerController{})
	beego.Router("/admin/customerDetail", &controllers.CustomerController{}, "get:GetCustomerDetailPage")
	beego.Router("/customer.api", &controllers.CustomerController{}, "*:APIhandler")

	//	后台-友情链接
	beego.Router("/admin/friendUrl", &controllers.FriendURLController{})
	beego.Router("/admin/friendUrlDetail", &controllers.FriendURLController{}, "get:GetFriendURLDetailPage")
	beego.Router("/friendUrl.api", &controllers.FriendURLController{}, "*:APIhandler")

	//	后台管理页面登录
	beego.Router("/admin/login", &controllers.AdminLoginController{})
	beego.Router("/admin/loginin", &controllers.AdminLoginController{}, "post:Loginin")

	//	后台管理页面注册
	beego.Router("/admin/register", &controllers.AdminRegisController{})
	beego.Router("/admin/regis", &controllers.AdminRegisController{}, "post:Regis")
}
