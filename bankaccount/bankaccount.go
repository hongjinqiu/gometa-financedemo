package bankaccount

import (
	"encoding/json"
	"fmt"
	. "github.com/hongjinqiu/gometa/component"
	. "github.com/hongjinqiu/gometa/controllers"
	. "github.com/hongjinqiu/gometa/error"
	"github.com/hongjinqiu/gometa/global"
	. "github.com/hongjinqiu/gometa/model"
	. "github.com/hongjinqiu/gometa/model/handler"
	. "github.com/hongjinqiu/gometa/mongo"
	"net/http"
	"strconv"
	"strings"
)

func init() {
}

type BankAccountSupport struct {
	ActionSupport
}

/**
* 为避免并发问题,重设amtOriginalCurrencyBalance为数据库中值
 */
func (o BankAccountSupport) BeforeSaveData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
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

		collectionName := "BankAccount"
		boInDb, found := qb.FindByMapWithSession(session, collectionName, queryMap)
		if !found {
			panic(BusinessError{Message: "银行账户保存前，银行账户未找到"})
		}

		bDataSetInDbLi := boInDb["B"].([]interface{})
		boInDb["B"] = bDataSetInDbLi

		bDataSetLi := (*bo)["B"].([]interface{})
		(*bo)["B"] = bDataSetLi

		for _, itemInDb := range bDataSetInDbLi {
			dataSetLineInDb := itemInDb.(map[string]interface{})
			currencyTypeIdInDb := fmt.Sprint(dataSetLineInDb["currencyTypeId"])

			for i, item := range bDataSetLi {
				dataSetLine := item.(map[string]interface{})
				bDataSetLi[i] = dataSetLine
				currencyTypeId := fmt.Sprint(dataSetLine["currencyTypeId"])

				if currencyTypeIdInDb == currencyTypeId {
					dataSetLine["amtOriginalCurrencyBalance"] = dataSetLineInDb["amtOriginalCurrencyBalance"]
					break
				}
			}
		}
	}
}

func (c BankAccountSupport) AfterNewData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	modelTemplateFactory := ModelTemplateFactory{}
	dataSetId := "B"
	data := modelTemplateFactory.GetDataSetNewData(dataSource, dataSetId, *bo)

	// 设置默认的币别
	qb := QuerySupport{}
	session, _ := global.GetConnection(sessionId)
	collection := "CurrencyType"
	query := map[string]interface{}{
		"A.code": "RMB",
	}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		query[k] = v
	}

	currencyType, found := qb.FindByMapWithSession(session, collection, query)
	if !found {
		panic(BusinessError{Message: "没有找到币别人民币，请先配置默认币别"})
	}
	data["currencyTypeId"] = currencyType["id"]

	(*bo)["B"] = []interface{}{
		data,
	}
}

func (c BankAccountSupport) AfterSaveData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}, diffDataRowLi *[]DiffDataRow) {
	for _, item := range *diffDataRowLi {
		if item.FieldGroupLi[0].GetDataSetId() == "B" { // 币别分录
			if item.SrcData != nil && item.DestData != nil { // 修改
				// 旧数据反过账,新数据正过账
				c.logBankAccountCurrencyType(sessionId, dataSource, *bo, item, BEFORE_UPDATE)
				c.logBankAccountCurrencyType(sessionId, dataSource, *bo, item, AFTER_UPDATE)
			} else if item.SrcData == nil && item.DestData != nil { // 新增
				// 新数据正过账
				c.logBankAccountCurrencyType(sessionId, dataSource, *bo, item, ADD)
			} else if item.SrcData != nil && item.DestData == nil { // 删除
				c.logBankAccountCurrencyType(sessionId, dataSource, *bo, item, DELETE)
			}
		}
	}
}

func (c BankAccountSupport) logBankAccountCurrencyType(sessionId int, bankAccountDataSource DataSource, bankAccountBo map[string]interface{}, diffDataRow DiffDataRow, diffDataType int) {
	if diffDataType == BEFORE_UPDATE { // 不管
		return
	}

	var addData map[string]interface{}
	var deleteData map[string]interface{}
	var afterUpdateData map[string]interface{}
	if diffDataType == ADD {
		addData = *(diffDataRow.DestData)
	} else if diffDataType == AFTER_UPDATE {
		afterUpdateData = *(diffDataRow.DestData)
	} else if diffDataType == DELETE {
		deleteData = diffDataRow.SrcData
	}

	bankAccountMasterData := bankAccountBo["A"].(map[string]interface{})
	bo := map[string]interface{}{}
	collectionName := "BankAccountCurrencyType"
	session, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)

	if diffDataType == AFTER_UPDATE { // 重新获取一遍bo
		beforeUpdateData := diffDataRow.SrcData
		qb := QuerySupport{}
		query := map[string]interface{}{
			"A.bankAccountId":  bankAccountMasterData["id"],
			"A.currencyTypeId": beforeUpdateData["currencyTypeId"],
		}
		bankAccountCurrencyType, found := qb.FindByMapWithSession(session, collectionName, query)
		if found {
			bo = bankAccountCurrencyType
		} else {
			bo["A"] = map[string]interface{}{}
		}
	}
	if diffDataType == ADD {
		bo["A"] = map[string]interface{}{
			"code":             bankAccountMasterData["code"],
			"name":             bankAccountMasterData["name"],
			"bankAccountId":    bankAccountMasterData["id"],
			"bankId":           bankAccountMasterData["bankId"],
			"accountProperty":  bankAccountMasterData["accountProperty"],
			"currencyTypeId":   addData["currencyTypeId"],
			"bankAccountBId":   addData["id"],
			"amtSecureBalance": addData["amtSecureBalance"],
			"amtLimits":        addData["amtLimits"],
			// TODO 安全余额,限额
		}
	} else if diffDataType == AFTER_UPDATE {
		boMaster := bo["A"].(map[string]interface{})
		boMaster["code"] = bankAccountMasterData["code"]
		boMaster["name"] = bankAccountMasterData["name"]
		boMaster["bankAccountId"] = bankAccountMasterData["id"]
		boMaster["bankId"] = bankAccountMasterData["bankId"]
		boMaster["accountProperty"] = bankAccountMasterData["accountProperty"]
		boMaster["currencyTypeId"] = afterUpdateData["currencyTypeId"]
		boMaster["bankAccountBId"] = afterUpdateData["id"]
		boMaster["amtSecureBalance"] = afterUpdateData["amtSecureBalance"]
		boMaster["amtLimits"] = afterUpdateData["amtLimits"]
		// TODO 安全余额,限额
		bo["A"] = boMaster
	} else if diffDataType == BEFORE_UPDATE { // 不管

	} else if diffDataType == DELETE {
		// 直接delete,return
		query := map[string]interface{}{
			"A.bankAccountId":  bankAccountMasterData["id"],
			"A.currencyTypeId": deleteData["currencyTypeId"],
		}
		txnManager.RemoveAll(txnId, collectionName, query)
		return
	}

	modelTemplateFactory := ModelTemplateFactory{}
	dataSourceModelId := "BankAccountCurrencyType"
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	modelTemplateFactory.ConvertDataType(dataSource, &bo)
	strId := modelTemplateFactory.GetStrId(bo)
	bankAccountAction := BankAccount{}
	if strId == "" || strId == "0" {
		masterSeqName := GetCollectionSequenceName(collectionName)
		masterSeqId := GetSequenceNo(db, masterSeqName)
		bo["_id"] = masterSeqId
		bo["id"] = masterSeqId
		boMaster := bo["A"].(map[string]interface{})
		boMaster["id"] = masterSeqId
		bo["A"] = boMaster
		bankAccountAction.SetCreateFixFieldValue(sessionId, dataSource, &bo)
	} else {
		bankAccountAction.SetModifyFixFieldValue(sessionId, dataSource, &bo)
	}

	if diffDataType == ADD {
		txnManager.Insert(txnId, collectionName, bo)
	} else if diffDataType == AFTER_UPDATE {
		if strId == "" || strId == "0" {
			txnManager.Insert(txnId, collectionName, bo)
		} else {
			txnManager.Update(txnId, collectionName, bo)
		}
	} else if diffDataType == BEFORE_UPDATE { // 不管

	} else if diffDataType == DELETE { // 前一步骤已经return,不管

	}
}

func (c BankAccountSupport) AfterDeleteData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	// 直接删除,整个删除 账户币别中的数据
	bankAccountMasterData := (*bo)["A"].(map[string]interface{})
	_, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	query := map[string]interface{}{
		"A.bankAccountId": bankAccountMasterData["id"],
	}
	collectionName := "BankAccountCurrencyType"
	txnManager.RemoveAll(txnId, collectionName, query)
}

type BankAccount struct {
	BaseDataAction
}

func (c BankAccount) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankAccountSupport{}
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BankAccount) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankAccountSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BankAccount) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankAccountSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BankAccount) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankAccountSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BankAccount) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankAccountSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c BankAccount) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankAccountSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BankAccount) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankAccountSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c BankAccount) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankAccountSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BankAccount) LogList(w http.ResponseWriter, r *http.Request) {
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
