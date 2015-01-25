package billpaymenttypeparameter

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type BillPaymentTypeParameterSupport struct {
	ActionSupport
}

type BillPaymentTypeParameter struct {
	BaseDataAction
}

func (c BillPaymentTypeParameter) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BillPaymentTypeParameter) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BillPaymentTypeParameter) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BillPaymentTypeParameter) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BillPaymentTypeParameter) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c BillPaymentTypeParameter) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BillPaymentTypeParameter) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c BillPaymentTypeParameter) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BillPaymentTypeParameter) LogList(w http.ResponseWriter, r *http.Request) {
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
