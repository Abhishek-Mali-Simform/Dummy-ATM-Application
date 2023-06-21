package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["Application/controllers:ATMController"] = append(beego.GlobalControllerRouter["Application/controllers:ATMController"],
        beego.ControllerComments{
            Method: "CheckBalanceFromATM",
            Router: "/check-balance",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:ATMController"] = append(beego.GlobalControllerRouter["Application/controllers:ATMController"],
        beego.ControllerComments{
            Method: "DepositMoneyFromATM",
            Router: "/deposit",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:ATMController"] = append(beego.GlobalControllerRouter["Application/controllers:ATMController"],
        beego.ControllerComments{
            Method: "WithdrawMoneyFromATM",
            Router: "/withdraw",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:AccountController"] = append(beego.GlobalControllerRouter["Application/controllers:AccountController"],
        beego.ControllerComments{
            Method: "RegisterAccount",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:BankController"] = append(beego.GlobalControllerRouter["Application/controllers:BankController"],
        beego.ControllerComments{
            Method: "RetrieveAccountInfo",
            Router: "/accounts-info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:BankController"] = append(beego.GlobalControllerRouter["Application/controllers:BankController"],
        beego.ControllerComments{
            Method: "RetrieveCardInfo",
            Router: "/cards-info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:BankController"] = append(beego.GlobalControllerRouter["Application/controllers:BankController"],
        beego.ControllerComments{
            Method: "CheckBalanceFromBank",
            Router: "/check-balance",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:BankController"] = append(beego.GlobalControllerRouter["Application/controllers:BankController"],
        beego.ControllerComments{
            Method: "DepositMoneyFromBank",
            Router: "/deposit",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:BankController"] = append(beego.GlobalControllerRouter["Application/controllers:BankController"],
        beego.ControllerComments{
            Method: "WithdrawMoneyFromBank",
            Router: "/withdraw",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:GoogleLoginController"] = append(beego.GlobalControllerRouter["Application/controllers:GoogleLoginController"],
        beego.ControllerComments{
            Method: "GoogleCallback",
            Router: "/callback",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:GoogleLoginController"] = append(beego.GlobalControllerRouter["Application/controllers:GoogleLoginController"],
        beego.ControllerComments{
            Method: "GoogleLogin",
            Router: "/login",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:MainController"] = append(beego.GlobalControllerRouter["Application/controllers:MainController"],
        beego.ControllerComments{
            Method: "FrontEndLoginPage",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:MainController"] = append(beego.GlobalControllerRouter["Application/controllers:MainController"],
        beego.ControllerComments{
            Method: "FrontEndAddPayeePage",
            Router: "/add-payee",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:MainController"] = append(beego.GlobalControllerRouter["Application/controllers:MainController"],
        beego.ControllerComments{
            Method: "FrontEndDashboardPage",
            Router: "/dashboard",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:MainController"] = append(beego.GlobalControllerRouter["Application/controllers:MainController"],
        beego.ControllerComments{
            Method: "Default",
            Router: "/default",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:MainController"] = append(beego.GlobalControllerRouter["Application/controllers:MainController"],
        beego.ControllerComments{
            Method: "FrontEndUserBalanceInfoPage",
            Router: "/get-balance",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:MainController"] = append(beego.GlobalControllerRouter["Application/controllers:MainController"],
        beego.ControllerComments{
            Method: "FrontEndLogout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:MainController"] = append(beego.GlobalControllerRouter["Application/controllers:MainController"],
        beego.ControllerComments{
            Method: "FrontEndRegisterPage",
            Router: "/register-user",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:MainController"] = append(beego.GlobalControllerRouter["Application/controllers:MainController"],
        beego.ControllerComments{
            Method: "FrontEndTransactionHistoryPage",
            Router: "/transaction-history",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:MainController"] = append(beego.GlobalControllerRouter["Application/controllers:MainController"],
        beego.ControllerComments{
            Method: "FrontEndTransferFundsPage",
            Router: "/transfer-funds",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:MainController"] = append(beego.GlobalControllerRouter["Application/controllers:MainController"],
        beego.ControllerComments{
            Method: "FrontEndGetAllUserPage",
            Router: "/users",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:MicrosoftLoginController"] = append(beego.GlobalControllerRouter["Application/controllers:MicrosoftLoginController"],
        beego.ControllerComments{
            Method: "MicrosoftCallback",
            Router: "/callback",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:MicrosoftLoginController"] = append(beego.GlobalControllerRouter["Application/controllers:MicrosoftLoginController"],
        beego.ControllerComments{
            Method: "MicrosoftLogin",
            Router: "/login",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:PayeeController"] = append(beego.GlobalControllerRouter["Application/controllers:PayeeController"],
        beego.ControllerComments{
            Method: "RegisterPayee",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:PayeeController"] = append(beego.GlobalControllerRouter["Application/controllers:PayeeController"],
        beego.ControllerComments{
            Method: "RetrieveAllPayee",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:TransactionController"] = append(beego.GlobalControllerRouter["Application/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "RetrieveTransactionHistory",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:TransactionController"] = append(beego.GlobalControllerRouter["Application/controllers:TransactionController"],
        beego.ControllerComments{
            Method: "RetrieveTransactionSummary",
            Router: "/report",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:UserController"] = append(beego.GlobalControllerRouter["Application/controllers:UserController"],
        beego.ControllerComments{
            Method: "RetrieveUserInfo",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:UserController"] = append(beego.GlobalControllerRouter["Application/controllers:UserController"],
        beego.ControllerComments{
            Method: "SignInUser",
            Router: "/sign-in",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Application/controllers:UserController"] = append(beego.GlobalControllerRouter["Application/controllers:UserController"],
        beego.ControllerComments{
            Method: "SignUpUser",
            Router: "/sign-up",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
