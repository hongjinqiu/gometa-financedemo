package measureunit

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type MeasureUnitSupport struct {
	ActionSupport
}

type MeasureUnit struct {
	BaseDataAction
}

func (c MeasureUnit) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c MeasureUnit) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}

	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c MeasureUnit) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c MeasureUnit) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c MeasureUnit) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c MeasureUnit) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c MeasureUnit) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c MeasureUnit) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c MeasureUnit) LogList(w http.ResponseWriter, r *http.Request) {
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
