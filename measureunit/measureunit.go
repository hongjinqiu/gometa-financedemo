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
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c MeasureUnit) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c MeasureUnit) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c MeasureUnit) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c MeasureUnit) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c MeasureUnit) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c MeasureUnit) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c MeasureUnit) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = MeasureUnitSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c MeasureUnit) LogList(w http.ResponseWriter, r *http.Request) {
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
