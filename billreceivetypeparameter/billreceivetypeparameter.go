package billreceivetypeparameter


import (
	. "github.com/hongjinqiu/gometa/controllers"
	"strings"
	"net/http"
	"encoding/json"
)

func init() {
}

type BillReceiveTypeParameterSupport struct {
	ActionSupport
}

type BillReceiveTypeParameter struct {

	BaseDataAction
}

func (c BillReceiveTypeParameter) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BillReceiveTypeParameter) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	
	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BillReceiveTypeParameter) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BillReceiveTypeParameter) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BillReceiveTypeParameter) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c BillReceiveTypeParameter) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BillReceiveTypeParameter) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c BillReceiveTypeParameter) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillReceiveTypeParameterSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BillReceiveTypeParameter) LogList(w http.ResponseWriter, r *http.Request) {
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
