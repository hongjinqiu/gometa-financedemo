package customer

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type CustomerSupport struct {
	ActionSupport
}

type Customer struct {
	BaseDataAction
}

func (c Customer) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerSupport{}
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c Customer) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c Customer) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c Customer) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c Customer) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c Customer) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c Customer) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c Customer) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c Customer) LogList(w http.ResponseWriter, r *http.Request) {
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
