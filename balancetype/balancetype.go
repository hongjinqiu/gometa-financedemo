package balancetype

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type BalanceTypeSupport struct {
	ActionSupport
}

type BalanceType struct {
	BaseDataAction
}

func (c BalanceType) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BalanceTypeSupport{}
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BalanceType) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BalanceTypeSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BalanceType) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BalanceTypeSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BalanceType) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BalanceTypeSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BalanceType) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BalanceTypeSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c BalanceType) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BalanceTypeSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c BalanceType) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BalanceTypeSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c BalanceType) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BalanceTypeSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c BalanceType) LogList(w http.ResponseWriter, r *http.Request) {
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
