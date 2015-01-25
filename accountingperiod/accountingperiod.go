package accountingperiod

import (
	"encoding/json"
	"fmt"
	. "github.com/hongjinqiu/gometa/common"
	. "github.com/hongjinqiu/gometa/component"
	. "github.com/hongjinqiu/gometa/controllers"
	. "github.com/hongjinqiu/gometa/error"
	"github.com/hongjinqiu/gometa/global"
	. "github.com/hongjinqiu/gometa/model"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func init() {
}

type AccountingPeriodSupport struct {
	ActionSupport
}

func (o AccountingPeriodSupport) AfterNewData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	masterData := (*bo)["A"].(map[string]interface{})

	year := time.Now().Year()
	masterData["accountingYear"] = year

	(*bo)["A"] = masterData

	numAccountingPeriod, err := strconv.Atoi(fmt.Sprint(masterData["numAccountingPeriod"]))
	if err != nil {
		panic(err)
	}
	detailDataLi := []interface{}{}

	modelTemplateFactory := ModelTemplateFactory{}
	dataSetId := "B"
	for i := 0; i < numAccountingPeriod; i++ {
		data := modelTemplateFactory.GetDataSetNewData(dataSource, dataSetId, *bo)
		data["id"] = "AfterNewData" + fmt.Sprint(i)
		data["sequenceNo"] = i + 1
		numStr := fmt.Sprint(i + 1)
		if i+1 < 10 {
			numStr = "0" + numStr
		}
		startDateStr := fmt.Sprint(year) + numStr + "01"
		startDate, err := strconv.Atoi(startDateStr)
		if err != nil {
			panic(err)
		}
		data["startDate"] = startDate
		startTime, err := time.Parse("20060102", startDateStr)
		if err != nil {
			panic(err)
		}
		nextMonthTime := startTime.AddDate(0, 1, -1)
		data["endDate"], err = strconv.Atoi(nextMonthTime.Format("20060102"))
		if err != nil {
			panic(err)
		}
		detailDataLi = append(detailDataLi, data)
	}

	(*bo)["B"] = detailDataLi
}

/**
删除前判断被用，会计期内有单据则视为被用
*/
func (o AccountingPeriodSupport) BeforeDeleteData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	bDataSetLi := (*bo)["B"].([]interface{})
	firstLineData := bDataSetLi[0].(map[string]interface{})
	lastLineData := bDataSetLi[len(bDataSetLi)-1].(map[string]interface{})

	commonUtil := CommonUtil{}
	firstStartDate := commonUtil.GetIntFromMap(firstLineData, "startDate")
	lastEndDate := commonUtil.GetIntFromMap(lastLineData, "endDate")

	qb := QuerySupport{}
	session, _ := global.GetConnection(sessionId)
	queryMap := map[string]interface{}{
		"A.billDate": map[string]interface{}{
			"$gte": firstStartDate,
			"$lt":  lastEndDate,
		},
	}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		queryMap[k] = v
	}

	modelTemplateFactory := ModelTemplateFactory{}
	dataSourceIdLi := []string{"GatheringBill", "PayBill"}
	for _, dataSourceId := range dataSourceIdLi {
		tmpDataSource := modelTemplateFactory.GetDataSource(dataSourceId)
		collectionName := modelTemplateFactory.GetCollectionName(tmpDataSource)
		_, found := qb.FindByMapWithSession(session, collectionName, queryMap)
		if found {
			panic(BusinessError{Message: "会计期范围内存在单据，不能删除"})
		}
	}
}

type AccountingPeriod struct {
	BaseDataAction
}

func (c AccountingPeriod) renderCommon(w http.ResponseWriter, r *http.Request, modelRenderVO ModelRenderVO) {
	bo := modelRenderVO.Bo
	relationBo := modelRenderVO.RelationBo
	dataSource := modelRenderVO.DataSource
	usedCheckBo := modelRenderVO.UsedCheckBo
	// 重新修改usedCheckBo,改为查询单据,只要在会计期内存在单据,则视为被用
	modelTemplateFactory := ModelTemplateFactory{}
	strId := modelTemplateFactory.GetStrId(bo)
	if strId != "" && strId != "0" {
		bDataSetLi := bo["B"].([]interface{})
		firstLineData := bDataSetLi[0].(map[string]interface{})
		lastLineData := bDataSetLi[len(bDataSetLi)-1].(map[string]interface{})

		commonUtil := CommonUtil{}
		firstStartDate := commonUtil.GetIntFromMap(firstLineData, "startDate")
		lastEndDate := commonUtil.GetIntFromMap(lastLineData, "endDate")

		qb := QuerySupport{}
		//GatheringBill,PayBill
		sessionId := global.GetSessionId()
		global.SetGlobalAttr(sessionId, "userId", fmt.Sprint(modelRenderVO.UserId))
		defer global.CloseSession(sessionId)
		session, _ := global.GetConnection(sessionId)

		queryMap := map[string]interface{}{
			"A.billDate": map[string]interface{}{
				"$gte": firstStartDate,
				"$lt":  lastEndDate,
			},
		}
		permissionSupport := PermissionSupport{}
		permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, modelRenderVO.FormTemplate.Security)
		for k, v := range permissionQueryDict {
			queryMap[k] = v
		}

		dataSourceIdLi := []string{"GatheringBill", "PayBill"}
		for _, dataSourceId := range dataSourceIdLi {
			tmpDataSource := modelTemplateFactory.GetDataSource(dataSourceId)
			collectionName := modelTemplateFactory.GetCollectionName(tmpDataSource)
			_, found := qb.FindByMapWithSession(session, collectionName, queryMap)
			if found {
				// 主数据集设置被用标记
				if usedCheckBo["A"] == nil {
					usedCheckBo["A"] = map[string]interface{}{}
				}
				masterUsedCheck := usedCheckBo["A"].(map[string]interface{})
				usedCheckBo["A"] = masterUsedCheck
				masterUsedCheck[strId] = true

				// 分录数据集设置被用标记
				if usedCheckBo["B"] == nil {
					usedCheckBo["B"] = map[string]interface{}{}
				}
				detailUsedCheck := usedCheckBo["B"].(map[string]interface{})
				usedCheckBo["B"] = usedCheckBo["B"]
				for _, detailData := range bDataSetLi {
					detailDataDict := detailData.(map[string]interface{})
					detailUsedCheck[fmt.Sprint(detailDataDict["id"])] = true
				}
				break
			}
		}
	}

	modelIterator := ModelIterator{}
	var result interface{} = ""
	modelIterator.IterateAllFieldBo(dataSource, &bo, &result, func(fieldGroup FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
		if (*data)[fieldGroup.Id] != nil {
			(*data)[fieldGroup.Id] = fmt.Sprint((*data)[fieldGroup.Id])
		}
	})
	format := r.FormValue("format")
	if strings.ToLower(format) == "json" {
		//		c.Response.ContentType = "application/json; charset=utf-8"
		//		return c.RenderJson(map[string]interface{}{
		//			"bo":          bo,
		//			"relationBo":  relationBo,
		//			"usedCheckBo": usedCheckBo,
		//		})
		w.Header()["Content-Type"] = []string{"application/json; charset=utf-8"}
		data, err := json.Marshal(&map[string]interface{}{
			"bo":          bo,
			"relationBo":  relationBo,
			"usedCheckBo": usedCheckBo,
		})
		if err != nil {
			panic(err)
		}
		w.Write(data)
		return
	}
	//c.Response.ContentType = "text/html; charset=utf-8"
	//	return c.Render()
	return
}

func (c AccountingPeriod) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.SaveCommon(w, r)
	c.renderCommon(w, r, modelRenderVO)
}

func (c AccountingPeriod) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountingPeriodSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.renderCommon(w, r, modelRenderVO)
}

func (c AccountingPeriod) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.renderCommon(w, r, modelRenderVO)
}

func (c AccountingPeriod) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.renderCommon(w, r, modelRenderVO)
}

func (c AccountingPeriod) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.renderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c AccountingPeriod) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.renderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountingPeriod) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.renderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c AccountingPeriod) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountingPeriodSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.renderCommon(w, r, modelRenderVO)
}

func (c AccountingPeriod) LogList(w http.ResponseWriter, r *http.Request) {
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
