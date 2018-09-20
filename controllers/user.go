package controllers
import(
    
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    "FullHouse/models"
    "encoding/json"
    "path"
    "fmt"
)
type UserControllers struct{
    beego.Controller
}
func (this *UserControllers)SendJson(resp interface{}){
   this.Data["json"]=resp
   this.ServeJSON()
}
func (this *UserControllers)Register(){
    beego.Info("UserControlleIRegister SUCC......")
    resp:=make(map[string]interface{})
    resp["errno"]=models.RECODE_OK
    resp["errmsg"]=models.RecodeText(models.RECODE_OK)
    defer this.SendJson(resp)
    var requestdata=make(map[string]interface{})
    json.Unmarshal(this.Ctx.Input.RequestBody,&requestdata)
    beego.Info("用户名:",requestdata["mobile"])
    beego.Info("密码:",requestdata["password"])
    beego.Info("验证码:",requestdata["sms_code"])
    if requestdata["mobile"]==""||requestdata["password"]==""||requestdata["sms_code"]==""{
        resp["errno"]=models.RECODE_NODATA
        resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)
        return
    }
    user:=models.User{};
   o:=orm.NewOrm()
    user.Mobile=requestdata["mobile"].(string)
    user.Password_hash=requestdata["password"].(string)
    user.Name=requestdata["mobile"].(string)
   id,err0:=o.Insert(&user)
   if err0!=nil{
       resp["errno"]=models.RECODE_DATAERR
       resp["errmsg"]=models.RecodeText(models.RECODE_DATAERR)
       return
   }
   beego.Info("User id=",id)
   this.SetSession("name",user.Mobile)
   this.SetSession("user_id",id)
   this.SetSession("mobile",user.Mobile)
   return
}
func (this *UserControllers)Login(){
    beego.Info("UserControlleILogin SUCC......")
    resp:=make(map[string]interface{})
    resp["errno"]=models.RECODE_OK
    resp["errmsg"]=models.RecodeText(models.RECODE_OK)
    defer this.SendJson(resp)
    var requestdata=make(map[string]interface{})
    json.Unmarshal(this.Ctx.Input.RequestBody,&requestdata)
    beego.Info("用户名:",requestdata["mobile"])
    beego.Info("密码:",requestdata["password"])
    user:=models.User{}
    o:=orm.NewOrm()
    qs:=o.QueryTable("user")
    err0:=qs.Filter("name",requestdata["mobile"].(string)).One(&user)
    if err0!=nil{
        beego.Info("用户已存在")
    resp["errno"]=models.RECODE_NODATA
    resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)
    return
    }
    if user.Password_hash!=requestdata["password"].(string){
        beego.Info("用户密码错误")
    resp["errno"]=models.RECODE_NODATA
    resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)
    return
    }
    this.SetSession("user_id",user.Id)
    this.SetSession("name",user.Name)

    return
}
func (this *UserControllers)UploadAvatar(){
    resp:=make(map[string]interface{})
    resp["errno"]=models.RECODE_OK
    resp["errmsg"]=models.RecodeText(models.RECODE_OK)
    defer this.SendJson(resp)
    file,filehead,err0:=this.GetFile("avatar")
    if err0!=nil{
        beego.Info("GetFile  Err.....")
        resp["errno"]=models.RECODE_NODATA
    resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)
    return
    }
    var fileBuffer =make([]byte,filehead.Size)
    if  _, err1:=file.Read(fileBuffer);err1!=nil{
        beego.Info("file.Read  Err.....")
        resp["errno"]=models.RECODE_NODATA
    resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)
    return
    }
    suffix:=path.Ext(filehead.Filename)
    beego.Info(suffix[1:])
    groupname,fileid,err2:=models.FDFSUploadByBuffer(fileBuffer,suffix[1:])
    if err2!=nil{
        beego.Info("FDFSUPloadByBuffer Err.....")
        resp["errno"]=models.RECODE_NODATA
    resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)
    return

    }
    beego.Info("groupname:",groupname,"fileid",fileid)
    user_id:=this.GetSession("user_id")
    fmt.Println("user_id:",user_id)
    user:=models.User{}
    o:=orm.NewOrm()
    user.Id=user_id.(int)
    user.Avatar_url=fileid
    if _,err3:=o.Update(&user,"avatar_url");err3!=nil{
        beego.Info("Update Err.....")
        resp["errno"]=models.RECODE_NODATA
    resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)
    return

    }
   url:="http://192.168.199.146:10086/"+fileid
   beego.Info(url)
   var url_map=make(map[string]interface{})
   url_map["avatar_url"]=url
   resp["data"]=url_map
   return
}
func (this *UserControllers)GetAuthor(){
    beego.Info("UserControllers GetAuthor SUCC......")
    resp:=make(map[string]interface{})
    resp["errno"]=models.RECODE_OK
    resp["errmsg"]=models.RecodeText(models.RECODE_OK)
    defer this.SendJson(resp)
    user_id:=this.GetSession("user_id")
    user:=models.User{}
    o:=orm.NewOrm()
    qs:=o.QueryTable("user")
    err0:=qs.Filter("id",user_id).One(&user)
    if err0!=nil{
    resp["errno"]=models.RECODE_NODATA
    resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)
     return 
    }
    var user_info=make(map[string]interface{})
    user_info["user_id"]=user.Id
    user_info["name"]=user.Name
    user_info["password"]=user.Password_hash
    user_info["mobile"]=user.Mobile
    user_info["real_name"]=user.Real_name
    user_info["id_card"]=user.Id_card
    user_info["avatar"]=user.Avatar_url
    resp["data"]=user_info
    return
}
func (this *UserControllers)PostAuthor(){
    beego.Info("UserControllers PostAuthor SUCC......")
    resp:=make(map[string]interface{})
    resp["errno"]=models.RECODE_OK
    resp["errmsg"]=models.RecodeText(models.RECODE_OK)
    defer this.SendJson(resp)
    user_id:=this.GetSession("user_id")
    beego.Info("user_id",user_id)
    user:=models.User{}
    o:=orm.NewOrm()
    var real_info=make(map[string]interface{})
    json.Unmarshal(this.Ctx.Input.RequestBody,&real_info)
    if real_info["real_name"]==""||real_info["id_card"]==""{
    resp["errno"]=models.RECODE_DATAERR
    resp["errmsg"]=models.RecodeText(models.RECODE_DATAERR)
     return 
    }
    user.Real_name=real_info["real_name"].(string)
    user.Id_card=real_info["id_card"].(string)
    _,err1:=o.Update(&user,"real_name","id_card")
    if err1!=nil{
        beego.Info("Update ERR.......")
    resp["errno"]=models.RECODE_DATAERR
    resp["errmsg"]=models.RecodeText(models.RECODE_DATAERR)
     return 
    }
    beego.Info("user_real_name:",user.Real_name,"user_id_card:",user.Id_card)
    this.SetSession("user_id",user.Id)
    this.SetSession("user_name",user.Real_name)
    return
}
