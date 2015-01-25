package customertype

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type CustomerTypeSupport struct {
	ActionSupport
}

type CustomerType struct {
	BaseDataAction
}

func (c CustomerType) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerTypeSupport{}
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CustomerType) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerTypeSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CustomerType) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerTypeSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CustomerType) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerTypeSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CustomerType) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerTypeSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c CustomerType) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerTypeSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CustomerType) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerTypeSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c CustomerType) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CustomerTypeSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CustomerType) LogList(w http.ResponseWriter, r *http.Request) {
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
