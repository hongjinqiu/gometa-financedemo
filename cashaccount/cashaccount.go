package cashaccount

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/component"
	. "github.com/hongjinqiu/gometa/controllers"
	. "github.com/hongjinqiu/gometa/error"
	"github.com/hongjinqiu/gometa/global"
	. "github.com/hongjinqiu/gometa/model"
	"net/http"
	"strconv"
	"strings"
)

func init() {
}

type CashAccountSupport struct {
	ActionSupport
}

func (o CashAccountSupport) AfterNewData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	masterData := (*bo)["A"].(map[string]interface{})
	(*bo)["A"] = masterData

	session, _ := global.GetConnection(sessionId)
	qb := QuerySupport{}
	query := map[string]interface{}{
		"A.code": "RMB",
	}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		query[k] = v
	}

	collectionName := "CurrencyType"
	result, found := qb.FindByMapWithSession(session, collectionName, query)
	if found {
		masterData["currencyTypeId"] = result["id"]
	}
}

/**
* 为避免并发问题,重设amtOriginalCurrencyBalance为数据库中值
 */
func (o CashAccountSupport) BeforeSaveData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	session, _ := global.GetConnection(sessionId)
	modelTemplateFactory := ModelTemplateFactory{}
	strId := modelTemplateFactory.GetStrId(*bo)
	if strId != "" && strId != "0" {
		id, err := strconv.Atoi(strId)
		if err != nil {
			panic(err)
		}
		qb := QuerySupport{}
		queryMap := map[string]interface{}{
			"_id": id,
		}
		permissionSupport := PermissionSupport{}
		permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
		for k, v := range permissionQueryDict {
			queryMap[k] = v
		}

		collectionName := "CashAccount"
		boInDb, found := qb.FindByMapWithSession(session, collectionName, queryMap)
		if !found {
			panic(BusinessError{Message: "现金账户保存前，现金账户未找到"})
		}

		masterData := (*bo)["A"].(map[string]interface{})
		(*bo)["A"] = masterData

		masterDataInDb := boInDb["A"].(map[string]interface{})
		boInDb["A"] = masterDataInDb

		masterData["amtOriginalCurrencyBalance"] = masterDataInDb["amtOriginalCurrencyBalance"]
	}
}

type CashAccount struct {
	BaseDataAction
}

func (c CashAccount) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountSupport{}
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CashAccount) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CashAccount) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CashAccount) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CashAccount) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c CashAccount) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CashAccount) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c CashAccount) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CashAccount) LogList(w http.ResponseWriter, r *http.Request) {
	result := c.LogListCommon(w, r)

	format := r.FormValue("format")
	if strings.ToLower(format) == "json" {
		w.Header()["Content-Type"] = []string{"application/json; charset=utf-8"}
		data, err := json.Marshal(&result)
		if err != nil {
			panic(err)
		}
		w.Write(data)
		return
	}
	return
}
