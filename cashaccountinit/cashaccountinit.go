package cashaccountinit

import (
	"encoding/json"
	"fmt"
	. "github.com/hongjinqiu/gometa-financedemo/accountinout"
	. "github.com/hongjinqiu/gometa/component"
	. "github.com/hongjinqiu/gometa/controllers"
	. "github.com/hongjinqiu/gometa/error"
	"github.com/hongjinqiu/gometa/global"
	. "github.com/hongjinqiu/gometa/model"
	. "github.com/hongjinqiu/gometa/model/handler"
	. "github.com/hongjinqiu/gometa/mongo"
	"github.com/hongjinqiu/gometa/session"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func init() {
}

type CashAccountInitSupport struct {
	ActionSupport
}

func (c CashAccountInitSupport) AfterSaveData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}, diffDateRowLi *[]DiffDataRow) {
	for _, item := range *diffDateRowLi {
		if item.SrcData != nil && item.DestData != nil { // 修改
			// 旧数据反过账,新数据正过账
			c.logCashAccount(sessionId, dataSource, item.SrcData, BEFORE_UPDATE)
			c.logCashAccount(sessionId, dataSource, *(item.DestData), AFTER_UPDATE)
		} else if item.SrcData == nil && item.DestData != nil { // 新增
			// 新数据正过账
			c.logCashAccount(sessionId, dataSource, *(item.DestData), ADD)
		}
	}
}

func (c CashAccountInitSupport) AfterDeleteData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	// 反过账
	data := (*bo)["A"].(map[string]interface{})
	c.logCashAccount(sessionId, dataSource, data, DELETE)
}

/**
* 检查赤字
 */
func (c CashAccountInitSupport) checkLimitsControl(sessionId int, diffDateRowAllLi []DiffDataRow, continueAnyAll string) {
	accountInOutService := AccountInOutService{}
	forbidLi, warnLi := accountInOutService.CheckCashAccountDiffDataRowLimitControl(sessionId, diffDateRowAllLi)
	if len(forbidLi) > 0 {
		panic(BusinessError{
			Code:    LIMIT_CONTROL_FORBID,
			Message: strings.Join(forbidLi, "<br />"),
		})
	}
	if len(warnLi) > 0 && continueAnyAll == "false" {
		panic(BusinessError{
			Code:    LIMIT_CONTROL_WARN,
			Message: strings.Join(warnLi, "<br />"),
		})
	}
}

func (c CashAccountInitSupport) logCashAccount(sessionId int, dataSource DataSource, data map[string]interface{}, diffDataType int) {
	accountId, err := strconv.Atoi(fmt.Sprint(data["accountId"]))
	if err != nil {
		panic(err)
	}
	currencyTypeId, err := strconv.Atoi(fmt.Sprint(data["currencyTypeId"]))
	if err != nil {
		panic(err)
	}
	accountInOutParam := AccountInOutParam{
		AccountId:      accountId,
		CurrencyTypeId: currencyTypeId,
		AmtIncrease:    fmt.Sprint(data["amtEarly"]),
		DiffDataType:   diffDataType,
	}

	accountInOutService := AccountInOutService{}
	accountInOutService.LogCashAccountInOut(sessionId, accountInOutParam)
}

type CashAccountInit struct {
	BaseDataAction
}

func (c CashAccountInit) SaveCommon(w http.ResponseWriter, r *http.Request) ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", session.GetFromSession(w, r, "userId"))
	global.SetGlobalAttr(sessionId, "adminUserId", session.GetFromSession(w, r, "adminUserId"))
	defer global.CloseSession(sessionId)
	defer c.RollbackTxn(sessionId)

	userId, err := strconv.Atoi(session.GetFromSession(w, r, "userId"))
	if err != nil {
		panic(err)
	}

	// 自己拆,再循环保存
	//	dataSourceModelId := c.Params.Form.Get("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := r.FormValue("formTemplateId")
	jsonBo := r.FormValue("jsonData")

	bo := map[string]interface{}{}
	err = json.Unmarshal([]byte(jsonBo), &bo)
	if err != nil {
		panic(err)
	}
	continueAnyAll := "false"
	if bo["continueAnyAll"] != nil && fmt.Sprint(bo["continueAnyAll"]) != "" {
		continueAnyAll = bo["continueAnyAll"].(string)
	}

	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)
	//	modelTemplateFactory.ConvertDataType(dataSource, &bo)

	cashAccountInitLi := []map[string]interface{}{}
	bDataSetLi := bo["B"].([]interface{})
	nowIdLi := []int{}
	for _, item := range bDataSetLi {
		mapItem := item.(map[string]interface{})
		line := map[string]interface{}{
			"A": mapItem,
		}
		if mapItem["id"] != nil {
			line["id"] = mapItem["id"]
			line["_id"] = mapItem["id"]
		} else {
			line["id"] = ""
			line["_id"] = ""
		}
		modelTemplateFactory.ConvertDataType(dataSource, &line)
		strId := fmt.Sprint(line["id"])
		if strId != "" {
			intId, _ := strconv.Atoi(strId)
			nowIdLi = append(nowIdLi, intId)
		}
		cashAccountInitLi = append(cashAccountInitLi, line)
	}
	queryData := bo["A"].(map[string]interface{})
	strAccountId := fmt.Sprint(queryData["accountId"])
	// 先处理删除的数据,并弄到差异数据中
	diffDataRowAllLi := []DiffDataRow{}
	toDeleteLi := c.dealDelete(sessionId, dataSource, formTemplate, queryData, nowIdLi)
	modelIterator := ModelIterator{}
	for _, item := range toDeleteLi {
		var result interface{} = ""
		modelIterator.IterateDataBo(dataSource, &item, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
			diffDataRow := DiffDataRow{
				FieldGroupLi: fieldGroupLi,
				SrcData:      *data,
				SrcBo:        item,
			}
			diffDataRowAllLi = append(diffDataRowAllLi, diffDataRow)
		})
	}

	bo = map[string]interface{}{
		"_id": 0,
		"id":  0,
		"A": map[string]interface{}{
			"id":        0,
			"accountId": strAccountId,
		},
	}
	bDataSetLi = []interface{}{}
	for i, _ := range cashAccountInitLi {
		cashAccountInit := &cashAccountInitLi[i]
		strId := modelTemplateFactory.GetStrId(*cashAccountInit)
		if strId == "" || strId == "0" {
			c.SetCreateFixFieldValue(sessionId, dataSource, cashAccountInit)
		} else {
			c.SetModifyFixFieldValue(sessionId, dataSource, cashAccountInit)
			editMessage, isValid := c.RActionSupport.EditValidate(sessionId, dataSource, formTemplate, *cashAccountInit)
			if !isValid {
				panic(editMessage)
			}
		}
		// 这样只会是新增和修改的数据
		c.RActionSupport.BeforeSaveData(sessionId, dataSource, formTemplate, cashAccountInit)
		financeService := FinanceService{}
		diffDataRowLi := financeService.SaveData(sessionId, dataSource, cashAccountInit)
		c.RActionSupport.AfterSaveData(sessionId, dataSource, formTemplate, cashAccountInit, diffDataRowLi)

		bDataSetLi = append(bDataSetLi, (*cashAccountInit)["A"])

		for _, diffDataRowItem := range *diffDataRowLi {
			diffDataRowAllLi = append(diffDataRowAllLi, diffDataRowItem)
		}
	}
	cashAccountInitSupport := c.RActionSupport.(CashAccountInitSupport)
	cashAccountInitSupport.checkLimitsControl(sessionId, diffDataRowAllLi, continueAnyAll)
	bo["B"] = bDataSetLi

	usedCheckBo := map[string]interface{}{}

	columnModelData := templateManager.GetColumnModelDataForFormTemplate(sessionId, formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.CommitTxn(sessionId)
	return ModelRenderVO{
		UserId:      userId,
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
	}
}

func (c CashAccountInit) dealDelete(sessionId int, dataSource DataSource, formTemplate FormTemplate, queryData map[string]interface{}, nowIdLi []int) []map[string]interface{} {
	modelTemplateFactory := ModelTemplateFactory{}
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	_, db := global.GetConnection(sessionId)
	queryMap := map[string]interface{}{
		"_id": map[string]interface{}{
			"$nin": nowIdLi,
		},
		"A.accountType": 1,
	}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		queryMap[k] = v
	}

	strAccountId := fmt.Sprint(queryData["accountId"])
	if strAccountId != "" && strAccountId != "0" {
		accountIdLi := []int{}
		for _, item := range strings.Split(strAccountId, ",") {
			accountId, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			accountIdLi = append(accountIdLi, accountId)
		}
		queryMap["A.accountId"] = map[string]interface{}{
			"$in": accountIdLi,
		}
	}
	queryMapByte, err := json.Marshal(&queryMap)
	if err != nil {
		panic(err)
	}
	log.Println("dealDelete,collectionName:" + collectionName + ", query:" + string(queryMapByte))
	toDeleteLi := []map[string]interface{}{}
	coll := db.C(collectionName)
	err = coll.Find(queryMap).All(&toDeleteLi)
	if err != nil {
		panic(err)
	}
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	modelIterator := ModelIterator{}
	usedCheck := UsedCheck{}
	for _, item := range toDeleteLi {
		if usedCheck.CheckUsed(sessionId, dataSource, item) {
			panic(BusinessError{Message: "已被用，不能删除"})
		}

		// 删除被用,帐户初始化不存在被用,不判断被用
		var result interface{} = ""
		modelIterator.IterateDataBo(dataSource, &item, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
			usedCheck.DeleteAll(sessionId, fieldGroupLi, *data)
		})

		c.RActionSupport.BeforeDeleteData(sessionId, dataSource, formTemplate, &item)
		_, removeResult := txnManager.Remove(txnId, collectionName, item)
		if !removeResult {
			panic("删除失败")
		}
		c.RActionSupport.AfterDeleteData(sessionId, dataSource, formTemplate, &item)
	}
	return toDeleteLi
}

func (c CashAccountInit) GetDataCommon(w http.ResponseWriter, r *http.Request) ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", session.GetFromSession(w, r, "userId"))
	global.SetGlobalAttr(sessionId, "adminUserId", session.GetFromSession(w, r, "adminUserId"))
	defer global.CloseSession(sessionId)
	defer c.RollbackTxn(sessionId)

	userId, err := strconv.Atoi(session.GetFromSession(w, r, "userId"))
	if err != nil {
		panic(err)
	}

	//	dataSourceModelId := r.FormValue("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := r.FormValue("formTemplateId")

	session, _ := global.GetConnection(sessionId)
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"A.accountType": 1,
	}
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		queryMap[k] = v
	}

	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	c.RActionSupport.BeforeGetData(sessionId, dataSource, formTemplate)
	// 需要进行循环处理,再转回来,
	pageNo := 1
	pageSize := 1000
	orderBy := ""

	queryMapByte, err := json.MarshalIndent(&queryMap, "", "\t")
	if err != nil {
		panic(err)
	}
	log.Println("dealDelete,collectionName:" + collectionName + ", query:" + string(queryMapByte))
	result := querySupport.IndexWithSession(session, collectionName, queryMap, pageNo, pageSize, orderBy)
	items := result["items"].([]interface{})
	dataSetLi := []interface{}{}
	for _, item := range items {
		line := item.(map[string]interface{})
		bo := map[string]interface{}{
			"_id": line["id"],
			"id":  line["id"],
			"A":   line["A"],
		}
		modelTemplateFactory.ConvertDataType(dataSource, &bo)
		c.RActionSupport.AfterGetData(sessionId, dataSource, formTemplate, &bo)
		dataSetLi = append(dataSetLi, bo["A"])
	}
	bo := map[string]interface{}{
		"_id": 0,
		"id":  0,
		"A": map[string]interface{}{
			"id":        0,
			"accountId": 0,
		},
		"B": dataSetLi,
	}

	//	usedCheck := UsedCheck{}
	//	usedCheckBo := usedCheck.GetFormUsedCheck(sessionId, dataSource, bo)
	usedCheckBo := map[string]interface{}{}

	columnModelData := templateManager.GetColumnModelDataForFormTemplate(sessionId, formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})

	//	modelTemplateFactory.ConvertDataType(dataSource, &bo)

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	return ModelRenderVO{
		UserId:      userId,
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
	}
}

func (c CashAccountInit) EditDataCommon(w http.ResponseWriter, r *http.Request) ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", session.GetFromSession(w, r, "userId"))
	global.SetGlobalAttr(sessionId, "adminUserId", session.GetFromSession(w, r, "adminUserId"))
	defer global.CloseSession(sessionId)
	defer c.RollbackTxn(sessionId)

	userId, err := strconv.Atoi(session.GetFromSession(w, r, "userId"))
	if err != nil {
		panic(err)
	}

	//	dataSourceModelId := r.FormValue("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := r.FormValue("formTemplateId")
	queryDataStr := r.FormValue("queryData")
	queryData := map[string]interface{}{}
	err = json.Unmarshal([]byte(queryDataStr), &queryData)
	if err != nil {
		panic(err)
	}
	strAccountId := fmt.Sprint(queryData["accountId"])
	//	strId := r.FormValue("id")
	//	id, err := strconv.Atoi(strId)
	//	if err != nil {
	//		panic(err)
	//	}

	session, _ := global.GetConnection(sessionId)
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"A.accountType": 1,
	}
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		queryMap[k] = v
	}

	if strAccountId != "" && strAccountId != "0" {
		accountIdLi := []int{}
		for _, item := range strings.Split(strAccountId, ",") {
			accountId, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			accountIdLi = append(accountIdLi, accountId)
		}
		queryMap["A.accountId"] = map[string]interface{}{
			"$in": accountIdLi,
		}
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	// 需要进行循环处理,再转回来,
	pageNo := 1
	pageSize := 1000
	orderBy := ""
	result := querySupport.IndexWithSession(session, collectionName, queryMap, pageNo, pageSize, orderBy)
	items := result["items"].([]interface{})
	dataSetLi := []interface{}{}
	for _, item := range items {
		line := item.(map[string]interface{})
		bo := map[string]interface{}{
			"_id": line["id"],
			"id":  line["id"],
			"A":   line["A"],
		}

		modelTemplateFactory.ConvertDataType(dataSource, &bo)
		editMessage, isValid := c.RActionSupport.EditValidate(sessionId, dataSource, formTemplate, bo)
		if !isValid {
			panic(editMessage)
		}

		c.RActionSupport.BeforeEditData(sessionId, dataSource, formTemplate, &bo)
		c.RActionSupport.AfterEditData(sessionId, dataSource, formTemplate, &bo)
		dataSetLi = append(dataSetLi, bo["A"])
	}

	bo := map[string]interface{}{
		"_id": 0,
		"id":  0,
		"A": map[string]interface{}{
			"id":        0,
			"accountId": strAccountId,
		},
		"B": dataSetLi,
	}

	//	usedCheck := UsedCheck{}
	//	usedCheckBo := usedCheck.GetFormUsedCheck(sessionId, dataSource, bo)
	usedCheckBo := map[string]interface{}{}

	columnModelData := templateManager.GetColumnModelDataForFormTemplate(sessionId, formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.CommitTxn(sessionId)
	return ModelRenderVO{
		UserId:       userId,
		Bo:           bo,
		RelationBo:   relationBo,
		UsedCheckBo:  usedCheckBo,
		DataSource:   dataSource,
		FormTemplate: formTemplate,
	}
}

func (c CashAccountInit) GiveUpDataCommon(w http.ResponseWriter, r *http.Request) ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", session.GetFromSession(w, r, "userId"))
	global.SetGlobalAttr(sessionId, "adminUserId", session.GetFromSession(w, r, "adminUserId"))
	defer global.CloseSession(sessionId)
	defer c.RollbackTxn(sessionId)

	userId, err := strconv.Atoi(session.GetFromSession(w, r, "userId"))
	if err != nil {
		panic(err)
	}

	//	dataSourceModelId := r.FormValue("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := r.FormValue("formTemplateId")
	queryDataStr := r.FormValue("queryData")
	queryData := map[string]interface{}{}
	err = json.Unmarshal([]byte(queryDataStr), &queryData)
	if err != nil {
		panic(err)
	}
	strAccountId := fmt.Sprint(queryData["accountId"])
	//	strId := r.FormValue("id")
	//	id, err := strconv.Atoi(strId)
	//	if err != nil {
	//		panic(err)
	//	}

	session, _ := global.GetConnection(sessionId)
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"A.accountType": 1,
	}
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		queryMap[k] = v
	}

	if strAccountId != "" && strAccountId != "0" {
		accountIdLi := []int{}
		for _, item := range strings.Split(strAccountId, ",") {
			accountId, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			accountIdLi = append(accountIdLi, accountId)
		}
		queryMap["A.accountId"] = map[string]interface{}{
			"$in": accountIdLi,
		}
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	pageNo := 1
	pageSize := 1000
	orderBy := ""
	result := querySupport.IndexWithSession(session, collectionName, queryMap, pageNo, pageSize, orderBy)
	items := result["items"].([]interface{})
	dataSetLi := []interface{}{}
	for _, item := range items {
		line := item.(map[string]interface{})
		bo := map[string]interface{}{
			"_id": line["id"],
			"id":  line["id"],
			"A":   line["A"],
		}
		modelTemplateFactory.ConvertDataType(dataSource, &bo)
		c.RActionSupport.BeforeGiveUpData(sessionId, dataSource, formTemplate, &bo)
		c.RActionSupport.AfterGiveUpData(sessionId, dataSource, formTemplate, &bo)
		dataSetLi = append(dataSetLi, bo["A"])
	}
	bo := map[string]interface{}{
		"_id": 0,
		"id":  0,
		"A": map[string]interface{}{
			"id":        0,
			"accountId": strAccountId,
		},
		"B": dataSetLi,
	}

	//	usedCheck := UsedCheck{}
	//	usedCheckBo := usedCheck.GetFormUsedCheck(sessionId, dataSource, bo)
	usedCheckBo := map[string]interface{}{}

	columnModelData := templateManager.GetColumnModelDataForFormTemplate(sessionId, formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.CommitTxn(sessionId)
	return ModelRenderVO{
		UserId:      userId,
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
	}
}

func (c CashAccountInit) RefreshDataCommon(w http.ResponseWriter, r *http.Request) ModelRenderVO {
	sessionId := global.GetSessionId()
	global.SetGlobalAttr(sessionId, "userId", session.GetFromSession(w, r, "userId"))
	global.SetGlobalAttr(sessionId, "adminUserId", session.GetFromSession(w, r, "adminUserId"))
	defer global.CloseSession(sessionId)
	defer c.RollbackTxn(sessionId)

	userId, err := strconv.Atoi(session.GetFromSession(w, r, "userId"))
	if err != nil {
		panic(err)
	}

	//	dataSourceModelId := r.FormValue("dataSourceModelId")
	dataSourceModelId := "AccountInit"
	formTemplateId := r.FormValue("formTemplateId")
	queryDataStr := r.FormValue("queryData")
	queryData := map[string]interface{}{}
	err = json.Unmarshal([]byte(queryDataStr), &queryData)
	if err != nil {
		panic(err)
	}
	strAccountId := fmt.Sprint(queryData["accountId"])
	//	strId := r.FormValue("id")
	//	id, err := strconv.Atoi(strId)
	//	if err != nil {
	//		panic(err)
	//	}

	session, _ := global.GetConnection(sessionId)
	querySupport := QuerySupport{}
	queryMap := map[string]interface{}{
		"A.accountType": 1,
	}
	templateManager := TemplateManager{}
	formTemplate := templateManager.GetFormTemplate(formTemplateId)
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		queryMap[k] = v
	}

	if strAccountId != "" && strAccountId != "0" {
		accountIdLi := []int{}
		for _, item := range strings.Split(strAccountId, ",") {
			accountId, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			accountIdLi = append(accountIdLi, accountId)
		}
		queryMap["A.accountId"] = map[string]interface{}{
			"$in": accountIdLi,
		}
	}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSource := modelTemplateFactory.GetDataSource(dataSourceModelId)
	collectionName := modelTemplateFactory.GetCollectionName(dataSource)
	// 需要进行循环处理,再转回来,
	pageNo := 1
	pageSize := 1000
	orderBy := ""
	result := querySupport.IndexWithSession(session, collectionName, queryMap, pageNo, pageSize, orderBy)
	items := result["items"].([]interface{})
	dataSetLi := []interface{}{}
	for _, item := range items {
		line := item.(map[string]interface{})
		bo := map[string]interface{}{
			"_id": line["id"],
			"id":  line["id"],
			"A":   line["A"],
		}
		modelTemplateFactory.ConvertDataType(dataSource, &bo)
		c.RActionSupport.BeforeRefreshData(sessionId, dataSource, formTemplate, &bo)
		c.RActionSupport.AfterRefreshData(sessionId, dataSource, formTemplate, &bo)
		dataSetLi = append(dataSetLi, bo["A"])
	}
	bo := map[string]interface{}{
		"_id": 0,
		"id":  0,
		"A": map[string]interface{}{
			"id":        0,
			"accountId": strAccountId,
		},
		"B": dataSetLi,
	}

	//	usedCheck := UsedCheck{}
	//	usedCheckBo := usedCheck.GetFormUsedCheck(sessionId, dataSource, bo)
	usedCheckBo := map[string]interface{}{}

	columnModelData := templateManager.GetColumnModelDataForFormTemplate(sessionId, formTemplate, bo)
	bo = columnModelData["bo"].(map[string]interface{})
	relationBo := columnModelData["relationBo"].(map[string]interface{})

	modelTemplateFactory.ClearReverseRelation(&dataSource)
	c.CommitTxn(sessionId)
	return ModelRenderVO{
		UserId:      userId,
		Bo:          bo,
		RelationBo:  relationBo,
		UsedCheckBo: usedCheckBo,
		DataSource:  dataSource,
	}
}

func (c CashAccountInit) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountInitSupport{}
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CashAccountInit) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountInitSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CashAccountInit) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountInitSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CashAccountInit) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountInitSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CashAccountInit) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountInitSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c CashAccountInit) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountInitSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CashAccountInit) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountInitSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c CashAccountInit) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CashAccountInitSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CashAccountInit) LogList(w http.ResponseWriter, r *http.Request) {
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
