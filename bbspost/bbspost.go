package bbspost

import (
	"encoding/json"
	"fmt"
	. "github.com/hongjinqiu/gometa/common"
	. "github.com/hongjinqiu/gometa/component"
	. "github.com/hongjinqiu/gometa/controllers"
	. "github.com/hongjinqiu/gometa/error"
	"github.com/hongjinqiu/gometa/global"
	. "github.com/hongjinqiu/gometa/model"
	. "github.com/hongjinqiu/gometa/model/handler"
	"github.com/hongjinqiu/gometa/mongo"
	. "github.com/hongjinqiu/gometa/mongo"
	"net/http"
	"strconv"
	"strings"
)

func init() {
}

type BbsPostSupport struct {
	ActionSupport
}

func (c BbsPostSupport) AfterSaveData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}, diffDateRowLi *[]DiffDataRow) {
	master := (*bo)["A"].(map[string]interface{})
	if fmt.Sprint(master["type"]) == "1" { // 主题帖
		c.bbsPostAfterSaveData(sessionId, dataSource, bo, diffDateRowLi)
	} else if fmt.Sprint(master["type"]) == "2" { // 主题帖回复
		c.bbsPostReplyAfterSaveData(sessionId, dataSource, bo, diffDateRowLi)
	}
}

func (c BbsPostSupport) bbsPostReplyAfterSaveData(sessionId int, dataSource DataSource, bo *map[string]interface{}, diffDateRowLi *[]DiffDataRow) {
	session, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)

	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	// 更新bbsPostId对应的主题帖的 lastReplyBy = self,lastReplyTime = currentTime,
	commonUtil := CommonUtil{}
	dateUtil := DateUtil{}
	master := (*bo)["A"].(map[string]interface{})
	qb := QuerySupport{}
	mainBbsPostQuery := map[string]interface{}{
		"_id": commonUtil.GetIntFromMap(master, "bbsPostId"),
	}
	bbsPostCollectionName := "BbsPost"
	mainBbsPost, found := qb.FindByMapWithSession(session, bbsPostCollectionName, mainBbsPostQuery)
	if !found {
		panic(BusinessError{Message: "主题帖未找到"})
	}
	mainBbsPostMaster := mainBbsPost["A"].(map[string]interface{})
	mainBbsPost["A"] = mainBbsPostMaster
	mainBbsPostMaster["lastReplyBy"] = userId
	mainBbsPostMaster["lastReplyTime"] = dateUtil.GetCurrentYyyyMMddHHmmss()
	if _, updateResult := txnManager.Update(txnId, bbsPostCollectionName, mainBbsPost); !updateResult {
		panic(BusinessError{Message: "主题帖更新失败"})
	}

	for _, item := range *diffDateRowLi {
		isNewOrUpdate := (item.SrcData != nil && item.DestData != nil) || (item.SrcData == nil && item.DestData != nil)
		if isNewOrUpdate {
			// 旧数据反过账,新数据正过账,修改后,更新 bbsPostId,readBy,bbsPostId
			bbsPostId := commonUtil.GetIntFromMap(*item.DestData, "bbsPostId")
			c.addOrUpdateBbsPostRead(sessionId, bbsPostId)
		}
	}
}

func (c BbsPostSupport) bbsPostAfterSaveData(sessionId int, dataSource DataSource, bo *map[string]interface{}, diffDateRowLi *[]DiffDataRow) {
	commonUtil := CommonUtil{}
	for _, item := range *diffDateRowLi {
		if item.SrcData != nil && item.DestData != nil { // 修改
			// 旧数据反过账,新数据正过账,修改后,更新 bbsPostId,readBy,bbsPostId
			bbsPostId := commonUtil.GetIntFromMap(*item.DestData, "id")
			c.addOrUpdateBbsPostRead(sessionId, bbsPostId)
		} else if item.SrcData == nil && item.DestData != nil { // 新增
			// 新数据正过账,新增记阅读记录,
			bbsPostId := commonUtil.GetIntFromMap(*item.DestData, "id")
			c.addBbsPostRead(sessionId, bbsPostId)
		}
	}
}

func (c BbsPostSupport) addOrUpdateBbsPostRead(sessionId int, bbsPostId int) {
	session, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	bbsPost := BbsPost{}
	modelTemplateFactory := ModelTemplateFactory{}
	bbsPostReadDS := modelTemplateFactory.GetDataSource("BbsPostRead")

	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	dateUtil := DateUtil{}
	qb := QuerySupport{}

	bbsPostRead, found := qb.FindByMapWithSession(session, "BbsPostRead", map[string]interface{}{
		"A.bbsPostId": bbsPostId,
		"A.readBy":    userId,
	})
	if found {
		bbsPost.SetModifyFixFieldValue(sessionId, bbsPostReadDS, &bbsPostRead)
		bbsPostReadA := bbsPostRead["A"].(map[string]interface{})
		bbsPostRead["A"] = bbsPostReadA

		bbsPostReadA["lastReadTime"] = dateUtil.GetCurrentYyyyMMddHHmmss()
		_, updateResult := txnManager.Update(txnId, "BbsPostRead", bbsPostRead)
		if !updateResult {
			panic(BusinessError{Message: "更新意见反馈阅读记录失败"})
		}
	} else {
		c.addBbsPostRead(sessionId, bbsPostId)
	}
}

func (c BbsPostSupport) addBbsPostRead(sessionId int, bbsPostId int) {
	_, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	bbsPost := BbsPost{}
	modelTemplateFactory := ModelTemplateFactory{}
	bbsPostReadDS := modelTemplateFactory.GetDataSource("BbsPostRead")

	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}
	dateUtil := DateUtil{}
	sequenceNo := mongo.GetSequenceNo(db, "bbsPostReadId")
	bbsPostRead := map[string]interface{}{
		"_id": sequenceNo,
		"id":  sequenceNo,
		"A": map[string]interface{}{
			"id":           sequenceNo,
			"bbsPostId":    bbsPostId,
			"readBy":       userId,
			"lastReadTime": dateUtil.GetCurrentYyyyMMddHHmmss(),
		},
	}
	bbsPost.SetCreateFixFieldValue(sessionId, bbsPostReadDS, &bbsPostRead)
	txnManager.Insert(txnId, "BbsPostRead", bbsPostRead)
}

func (o BbsPostSupport) BeforeDeleteData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	session, _ := global.GetConnection(sessionId)
	// 已回复过的帖子不可删除
	qb := QuerySupport{}
	bbsPostCollectionName := "BbsPost"
	bbsPostReplyQuery := map[string]interface{}{
		"A.bbsPostId": (*bo)["id"],
	}
	permissionSupport := PermissionSupport{}
	permissionQueryDict := permissionSupport.GetPermissionQueryDict(sessionId, formTemplate.Security)
	for k, v := range permissionQueryDict {
		bbsPostReplyQuery[k] = v
	}

	_, found := qb.FindByMapWithSession(session, bbsPostCollectionName, bbsPostReplyQuery)
	if found {
		panic(BusinessError{Message: "存在回复的主题帖不可删除"})
	}
}

func (c BbsPostSupport) AfterDeleteData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	// 反过账
	_, db := global.GetConnection(sessionId)
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	userId, err := strconv.Atoi(fmt.Sprint(global.GetGlobalAttr(sessionId, "userId")))
	if err != nil {
		panic(err)
	}

	query := map[string]interface{}{
		"A.bbsPostId": (*bo)["id"],
		"A.readBy":    userId,
	}
	txnManager.RemoveAll(txnId, "BbsPostRead", query)
}

type BbsPost struct {
	BaseDataAction
}

func (c BbsPost) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BbsPostSupport{}
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BbsPost) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BbsPostSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BbsPost) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BbsPostSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BbsPost) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BbsPostSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BbsPost) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BbsPostSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c BbsPost) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BbsPostSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BbsPost) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BbsPostSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c BbsPost) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BbsPostSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BbsPost) LogList(w http.ResponseWriter, r *http.Request) {
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
