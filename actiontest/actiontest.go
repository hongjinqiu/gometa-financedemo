package actiontest

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/component"
	. "github.com/hongjinqiu/gometa/controllers"
	. "github.com/hongjinqiu/gometa/model"
	. "github.com/hongjinqiu/gometa/model/handler"
	"net/http"
	"strings"
)

func init() {
}

type ActionTestSupport struct {
	ActionSupport
}

func (c ActionTestSupport) BeforeSaveData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}) {
	println("ActionTestSupport BeforeSaveData")
}

func (c ActionTestSupport) AfterSaveData(sessionId int, dataSource DataSource, formTemplate FormTemplate, bo *map[string]interface{}, diffDateRowLi *[]DiffDataRow) {
	println("ActionTestSupport AfterSaveData")
}

type ActionTest struct {
	BillAction
}

func (c ActionTest) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c ActionTest) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ActionTestSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c ActionTest) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c ActionTest) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c ActionTest) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c ActionTest) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c ActionTest) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c ActionTest) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c ActionTest) LogList(w http.ResponseWriter, r *http.Request) {
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

func (c ActionTest) CancelData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ActionTestSupport{}
	modelRenderVO := c.CancelDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c ActionTest) UnCancelData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ActionSupport{}
	modelRenderVO := c.UnCancelDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}
