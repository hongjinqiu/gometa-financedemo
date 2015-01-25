package incomeitem

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type IncomeItemSupport struct {
	ActionSupport
}

type IncomeItem struct {
	BaseDataAction
}

func (c IncomeItem) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c IncomeItem) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c IncomeItem) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c IncomeItem) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c IncomeItem) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c IncomeItem) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c IncomeItem) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c IncomeItem) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c IncomeItem) LogList(w http.ResponseWriter, r *http.Request) {
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
