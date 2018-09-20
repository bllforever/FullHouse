package controllers
import(
    
	"github.com/astaxie/beego"
    "FullHouse/models"
)
type SessionControllers struct{
    beego.Controller
}
func (this *SessionControllers)SendJson(resp interface{}){
   this.Data["json"]=resp
   this.ServeJSON()
}
func (this *SessionControllers)DelSessionName(){
beego.Info("SessionController DelSessionName SUCC.....")
    resp:=make(map[string]interface{})
    resp["errno"]=models.RECODE_OK
    resp["errmsg"]=models.RecodeText(models.RECODE_OK)
    defer this.SendJson(resp)
    //if this.GetSession("name")!=nil{
        this.DelSession("name")
  //  }
   // if this.GetSession("mobile")!=nil{
        this.DelSession("mobile")
    //}
   // if this.GetSession("user_id")!=nil{
        this.DelSession("user_id")
   // }
    return

}
func (this *SessionControllers)GetSessionName(){
    beego.Info("SessionController GetSessionName SUCC.....")
    resp:=make(map[string]interface{})
    resp["errno"]=models.RECODE_SESSIONERR
    resp["errmsg"]=models.RecodeText(models.RECODE_SESSIONERR)
    defer this.SendJson(resp)
    var namemap=make(map[string]interface{})
    name:=this.GetSession("name")
    if name!=nil{
    namemap["name"]=name
    resp["errno"]=models.RECODE_OK
    resp["errmsg"]=models.RecodeText(models.RECODE_OK)
    resp["data"]=namemap

}
    return
}
