package currencytype

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type CurrencyTypeSupport struct {
	ActionSupport
}

type CurrencyType struct {
	BaseDataAction
}

func (c CurrencyType) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CurrencyType) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CurrencyTypeSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CurrencyType) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CurrencyType) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CurrencyType) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c CurrencyType) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c CurrencyType) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c CurrencyType) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = CurrencyTypeSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c CurrencyType) LogList(w http.ResponseWriter, r *http.Request) {
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
