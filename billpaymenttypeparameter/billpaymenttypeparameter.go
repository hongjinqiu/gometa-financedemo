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
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c BillPaymentTypeParameter) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}

	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c BillPaymentTypeParameter) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c BillPaymentTypeParameter) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c BillPaymentTypeParameter) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c BillPaymentTypeParameter) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BillPaymentTypeParameter) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c BillPaymentTypeParameter) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillPaymentTypeParameterSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c BillPaymentTypeParameter) LogList(w http.ResponseWriter, r *http.Request) {
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
