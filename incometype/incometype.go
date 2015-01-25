package incometype

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type IncomeTypeSupport struct {
	ActionSupport
}

type IncomeType struct {
	BaseDataAction
}

func (c IncomeType) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeTypeSupport{}
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c IncomeType) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeTypeSupport{}

	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c IncomeType) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeTypeSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c IncomeType) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeTypeSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c IncomeType) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeTypeSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c IncomeType) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeTypeSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c IncomeType) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeTypeSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c IncomeType) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = IncomeTypeSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c IncomeType) LogList(w http.ResponseWriter, r *http.Request) {
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
