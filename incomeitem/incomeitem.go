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
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c IncomeItem) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}

	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c IncomeItem) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c IncomeItem) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c IncomeItem) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c IncomeItem) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c IncomeItem) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c IncomeItem) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeItemSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c IncomeItem) LogList(w http.ResponseWriter, r *http.Request) {
	result := c.RLogListCommon(w, r)

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
