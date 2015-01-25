package billtype

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type BillTypeSupport struct {
	ActionSupport
}

type BillType struct {
	BaseDataAction
}

func (c BillType) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillTypeSupport{}
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c BillType) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillTypeSupport{}

	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c BillType) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillTypeSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c BillType) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillTypeSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c BillType) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillTypeSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c BillType) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillTypeSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BillType) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillTypeSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c BillType) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BillTypeSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c BillType) LogList(w http.ResponseWriter, r *http.Request) {
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
