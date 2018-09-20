package controllers
import(
    
	"github.com/astaxie/beego"
    "FullHouse/models"
)
type IndexControllers struct{
    beego.Controller
}
func (this *IndexControllers)SendJson(resp interface{}){
   this.Data["json"]=resp
   this.ServeJSON()
}
func (this *IndexControllers)GetIndexInfo(){
    beego.Info("IndexController SUCC......")
    resp:=make(map[string]interface{})
    resp["errno"]=models.RECODE_OK
    resp["errmsg"]=models.RecodeText(models.RECODE_OK)
    defer this.SendJson(resp)
}
