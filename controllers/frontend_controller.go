package controllers

// MainController operations for front end
type MainController struct {
	CommonFunction
}

// URLMapping ...
func (controller *MainController) URLMapping() {
	controller.Mapping("Default", controller.Default)
	controller.Mapping("FrontEndLoginPage", controller.FrontEndLoginPage)
	controller.Mapping("FrontEndRegisterPage", controller.FrontEndRegisterPage)
	controller.Mapping("FrontEndGetAllUserPage", controller.FrontEndGetAllUserPage)
	controller.Mapping("FrontEndDashboardPage", controller.FrontEndDashboardPage)
	controller.Mapping("FrontEndUserBalanceInfoPage", controller.FrontEndUserBalanceInfoPage)
	controller.Mapping("FrontEndLogout", controller.FrontEndLogout)
	controller.Mapping("FrontEndTransactionHistoryPage", controller.FrontEndTransactionHistoryPage)
	controller.Mapping("FrontEndAddPayeePage", controller.FrontEndAddPayeePage)
	controller.Mapping("FrontEndTransferFundsPage", controller.FrontEndTransferFundsPage)
}

// Default ...
// @Title Retrieve User(s) Info
// @Description get User Info based on token, id or all users
// @Success 200 Success Message
// @Failure 403
// @router /default [get]
func (controller *MainController) Default() {
	controller.Data["Website"] = "beego.me"
	controller.Data["Email"] = "astaxie@gmail.com"
	controller.TplName = "index.tpl"
}

// FrontEndLoginPage ...
// @Title FrontEndLoginPage
// @Description Front End Login Page
// @Success 200 Success Message
// @Failure 403
// @router / [get]
func (controller *MainController) FrontEndLoginPage() {
	controller.TplName = "login.html"
}

// FrontEndRegisterPage ...
// @Title FrontEndRegisterPage
// @Description Front End Register Page
// @Success 200 Success Message
// @Failure 403
// @router /register-user [get]
func (controller *MainController) FrontEndRegisterPage() {
	controller.TplName = "register_user.html"
}

// FrontEndGetAllUserPage ...
// @Title FrontEndGetAllUserPage
// @Description Front End Get All UserPage
// @Success 200 Success Message
// @Failure 403
// @router /users [get]
func (controller *MainController) FrontEndGetAllUserPage() {
	controller.TplName = "listUsers.html"
}

// FrontEndDashboardPage ...
// @Title FrontEndDashboardPage
// @Description Front End Dashboard Page
// @Success 200 Success Message
// @Failure 403
// @router /dashboard [get]
func (controller *MainController) FrontEndDashboardPage() {
	token, claim, check := controller.GetCredentials()
	if check {
		controller.TplName = "dashboard.html"
		controller.Data["token"] = token
		controller.Data["claim"] = claim
	} else {
		controller.Redirect("/world-bank", 302)
	}
}

// FrontEndLogout ...
// @Title FrontEndLogout
// @Description Front End Logout
// @Success 200 Success Message
// @Failure 403
// @router /logout [get]
func (controller *MainController) FrontEndLogout() {
	_, _, check := controller.GetCredentials()
	if check {
		if controller.RemoveCredentials() {
			controller.Redirect("/world-bank", 302)
		}
	} else {
		controller.InvalidMethod()
	}
}

// FrontEndUserBalanceInfoPage ...
// @Title FrontEndUserBalanceInfoPage
// @Description Front End Balance Info Page
// @Success 200 Success Message
// @Failure 403
// @router /get-balance [get]
func (controller *MainController) FrontEndUserBalanceInfoPage() {
	token, claim, check := controller.GetCredentials()
	if check {
		controller.TplName = "user_balance_info.html"
		controller.Data["token"] = token
		controller.Data["claim"] = claim
	} else {
		controller.Redirect("/world-bank", 302)
	}
}

// FrontEndTransactionHistoryPage ...
// @Title FrontEndTransactionHistoryPage
// @Description Front End Transaction history Page
// @Success 200 Success Message
// @Failure 403
// @router /transaction-history [get]
func (controller *MainController) FrontEndTransactionHistoryPage() {
	token, claim, check := controller.GetCredentials()
	if check {
		controller.TplName = "transaction_history.html"
		controller.Data["token"] = token
		controller.Data["claim"] = claim
	} else {
		controller.Redirect("/world-bank", 302)
	}
}

// FrontEndAddPayeePage ...
// @Title FrontEndAddPayeePage
// @Description Front End Add Payee Page
// @Success 200 Success Message
// @Failure 403
// @router /add-payee [get]
func (controller *MainController) FrontEndAddPayeePage() {
	token, claim, check := controller.GetCredentials()
	if check {
		controller.TplName = "add_payee.html"
		controller.Data["token"] = token
		controller.Data["claim"] = claim
	} else {
		controller.Redirect("/world-bank", 302)
	}
}

// FrontEndTransferFundsPage ...
// @Title FrontEndTransferFundsPage
// @Description Front End Transfer Funds Page
// @Success 200 Success Message
// @Failure 403
// @router /transfer-funds [get]
func (controller *MainController) FrontEndTransferFundsPage() {
	token, claim, check := controller.GetCredentials()
	if check {
		controller.TplName = "transfer_funds.html"
		controller.Data["token"] = token
		controller.Data["claim"] = claim
	} else {
		controller.Redirect("/world-bank", 302)
	}
}
