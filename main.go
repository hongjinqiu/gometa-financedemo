package main

import (
	"encoding/json"
	"flag"
	"github.com/hongjinqiu/gometa-financedemo/accountingperiod"
	"github.com/hongjinqiu/gometa-financedemo/accountinit"
	"github.com/hongjinqiu/gometa-financedemo/actiontest"
	"github.com/hongjinqiu/gometa-financedemo/article"
	"github.com/hongjinqiu/gometa-financedemo/articletype"
	"github.com/hongjinqiu/gometa-financedemo/balancetype"
	"github.com/hongjinqiu/gometa-financedemo/bank"
	"github.com/hongjinqiu/gometa-financedemo/bankaccount"
	"github.com/hongjinqiu/gometa-financedemo/bankaccountinit"
	"github.com/hongjinqiu/gometa-financedemo/bbspost"
	"github.com/hongjinqiu/gometa-financedemo/billpaymenttypeparameter"
	"github.com/hongjinqiu/gometa-financedemo/billreceivetypeparameter"
	"github.com/hongjinqiu/gometa-financedemo/billtype"
	"github.com/hongjinqiu/gometa-financedemo/cashaccount"
	"github.com/hongjinqiu/gometa-financedemo/cashaccountinit"
	"github.com/hongjinqiu/gometa-financedemo/currencytype"
	"github.com/hongjinqiu/gometa-financedemo/customer"
	"github.com/hongjinqiu/gometa-financedemo/customertype"
	"github.com/hongjinqiu/gometa-financedemo/demo"
	"github.com/hongjinqiu/gometa-financedemo/gatheringbill"
	"github.com/hongjinqiu/gometa-financedemo/incomeitem"
	"github.com/hongjinqiu/gometa-financedemo/incometype"
	"github.com/hongjinqiu/gometa-financedemo/measureunit"
	"github.com/hongjinqiu/gometa-financedemo/paybill"
	"github.com/hongjinqiu/gometa-financedemo/paypact"
	"github.com/hongjinqiu/gometa-financedemo/provider"
	"github.com/hongjinqiu/gometa-financedemo/providertype"
	"github.com/hongjinqiu/gometa-financedemo/pubreferencelog"
	"github.com/hongjinqiu/gometa-financedemo/systemparameter"
	"github.com/hongjinqiu/gometa-financedemo/sysuser"
	"github.com/hongjinqiu/gometa-financedemo/taxtype"
	"github.com/hongjinqiu/gometa/app"
	"github.com/hongjinqiu/gometa/console"
	. "github.com/hongjinqiu/gometa/error"
	"github.com/hongjinqiu/gometa/log"
	"github.com/hongjinqiu/gometa/route"
	"github.com/hongjinqiu/gometa/session"
	"net/http"
	"reflect"
	"runtime/debug"
	"strings"
)

var addr = flag.String("addr", ":8888", "address")

var constrollersDict []reflect.Type = []reflect.Type{}

func init() {
	constrollersDict = append(constrollersDict, reflect.TypeOf(console.Console{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(app.App{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(articletype.ArticleType{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(article.Article{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(sysuser.SysUser{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(accountingperiod.AccountingPeriod{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(accountinit.AccountInit{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(actiontest.ActionTest{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(balancetype.BalanceType{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(bank.Bank{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(bankaccount.BankAccount{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(bankaccountinit.BankAccountInit{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(bbspost.BbsPost{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(billpaymenttypeparameter.BillPaymentTypeParameter{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(billreceivetypeparameter.BillReceiveTypeParameter{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(billtype.BillType{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(cashaccount.CashAccount{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(cashaccountinit.CashAccountInit{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(currencytype.CurrencyType{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(customer.Customer{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(customertype.CustomerType{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(demo.Demo{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(gatheringbill.GatheringBill{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(incomeitem.IncomeItem{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(incometype.IncomeType{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(measureunit.MeasureUnit{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(paybill.PayBill{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(paypact.PayPact{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(provider.Provider{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(providertype.ProviderType{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(pubreferencelog.PubReferenceLog{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(systemparameter.SystemParameter{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(taxtype.TaxType{}))

	// 注:route.FilterLi里面不能添加>2的过滤器,否则go里面的url在并发访问时匹配会报错,例如:/app/comboview有可能会匹配到/app/combo,非常奇怪的问题,不知道为什么,你不信?你自己试一试
	route.FilterLi = append(route.FilterLi, SessionAdapter)
	route.FilterLi = append(route.FilterLi, UTF8AndBusinessPanicFilter)
}

func SessionAdapter(w http.ResponseWriter, r *http.Request, li []route.HttpHandleFilterFunc) {
	userAgent := strings.Join(r.Header["User-Agent"], ",")
	if strings.Index(userAgent, "Firefox") > -1 {
		session.SetToSession(w, r, "userId", "10")
	} else {
		session.SetToSession(w, r, "userId", "20")
	}
	session.SetToSession(w, r, "adminUserId", "1")
	li[0](w, r, li[1:])
}

func UTF8AndBusinessPanicFilter(w http.ResponseWriter, r *http.Request, li []route.HttpHandleFilterFunc) {
	defer func() {
		if x := recover(); x != nil {
			if reflect.TypeOf(x).Name() == "BusinessError" {
				err := x.(BusinessError)
				w.Header()["Content-Type"] = []string{"application/json; charset=utf-8"}
				data, marshalErr := json.Marshal(&map[string]interface{}{
					"success": false,
					"code":    err.Code,
					"message": err.Error(),
				})
				if marshalErr != nil {
					panic(marshalErr)
				}
				w.Write(data)
			} else {
				log.Error(x, "\n", string(debug.Stack()))
				panic(x)
			}
		}
	}()
	w.Header()["Content-Type"] = []string{"text/html; charset=utf-8"}
	li[0](w, r, li[1:])
}

func registerRoute() {
	route.RegisteStaticFilePath()
	route.RegisteReflectController(constrollersDict)

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		app.App{}.Index(w, r)
	})
}

//devsrvr -addr=:8001 -targetaddr=localhost:8888 D:\goworkspace\src\github.com\gometa-financedemo D:\goworkspace\src\github.com\gometa
//devsrvr -addr=:8001 -targetaddr=localhost:8888 /home/hongjinqiu/goworkspace/src/github.com/hongjinqiu/gometa-financedemo /home/hongjinqiu/goworkspace/src/github.com/hongjinqiu/gometa
func main() {
	flag.Parse()
	log.Warn("listening on:", *addr)

	registerRoute()

	log.Error(http.ListenAndServe(*addr, nil))
}
