// Package routers
// @APIVersion 1.0.0
// @Title BANKING APPLICATION
// @Description bank application to perform simple withdraw, deposit, check balance operation
// @Contact abhishek.m@simformsolutions.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Application/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	backendNameSpace := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/atm",
			beego.NSInclude(&controllers.ATMController{}),
		),
		beego.NSNamespace("/bank",
			beego.NSInclude(&controllers.BankController{}),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(&controllers.UserController{}),
		),
		beego.NSNamespace("/account",
			beego.NSInclude(&controllers.AccountController{}),
		),
		beego.NSNamespace("/google",
			beego.NSInclude(&controllers.GoogleLoginController{}),
		),
		beego.NSNamespace("/transaction",
			beego.NSInclude(&controllers.TransactionController{}),
		),
		beego.NSNamespace("/microsoft",
			beego.NSInclude(&controllers.MicrosoftLoginController{}),
		),
		beego.NSNamespace("/payee",
			beego.NSInclude(&controllers.PayeeController{}),
		),
	)

	frontendNamespace := beego.NewNamespace("/world-bank",
		beego.NSInclude(&controllers.MainController{}),
	)
	beego.AddNamespace(backendNameSpace, frontendNamespace)
}
