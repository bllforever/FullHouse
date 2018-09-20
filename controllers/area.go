
package controllers
import(
"fmt"    
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/cache"
   _"github.com/astaxie/beego/cache/redis"
  "encoding/json"
   "time"
  "FullHouse/models"
)
type AreaControllers struct{
    beego.Controller
}
func (this *AreaControllers)SendJson(resquest interface{}){
        this.Data["json"]=resquest
        this.ServeJSON()
}
func (this *AreaControllers)GetArea(){
    beego.Info("AreaControllers GetArea SUCC...")
    resp:=make(map[string]interface{})
   resp["errno"]=models.RECODE_OK
   resp["errmsg"]=models.RecodeText(models.RECODE_OK)
    defer this.SendJson(resp)
     //1.从缓存中读取数据
     //1.1连接redis
     cache_conn,err0:=cache.NewCache("redis",`{"key":"FullHouse","conn":"127.0.0.1:6379","dbnum":"0"}`)
     if err0!=nil{
         beego.Info("cache.NewCache err...")
        resp["errno"]=models.RECODE_SERVERERR
        resp["errmsg"]=models.RecodeText(models.RECODE_SERVERERR) 
         return
     }
   //2.从缓存中查询数据
   //3.缓存中无数据，将数据库中的数据返回给前端
   area_info_value:=cache_conn.Get("area_info")
   if area_info_value!=nil{
       beego.Info("cache.Get SUCC.....")
     //将数据Json格式转换为map格式
       var area_info interface{}
     json.Unmarshal(area_info_value.([]byte),&area_info)
       resp["data"]=area_info
     return
   }
   //3.1数据库中指定表
   var areas []*models.Area
   o:=orm.NewOrm()
   qs:=o.QueryTable("area")
   //3.2从表中读取数据
   num,err1:=qs.All(&areas)
   if err1!=nil{
       resp["errno"]=models.RECODE_DATAERR
       resp["errmsg"]=models.RecodeText(models.RECODE_DATAERR) 
       return
   }
   beego.Info("num=",num)
   resp["data"]=areas
   //3.3将数据库中获得的数据转换成json
   area_info_str,err2:=json.Marshal(areas)
   if err2!=nil{
       beego.Info("json.Marshal err....")
        resp["errno"]=models.RECODE_SERVERERR
        resp["errmsg"]=models.RecodeText(models.RECODE_SERVERERR) 
         return
   }
   fmt.Printf("%s\n",area_info_str)
   //3.4将json格式的数据存到redis中
   cache_conn.Put("xxx","value",time.Second*3600)
   str:=cache_conn.Get("xxx")
   fmt.Printf("%s\n",str)
   if err3:=cache_conn.Put("area_info",area_info_str,time.Second*3600);err3!=nil{
       beego.Info("cache_conn Put err....")
        resp["errno"]=models.RECODE_SERVERERR
        resp["errmsg"]=models.RecodeText(models.RECODE_SERVERERR) 
         return
   }
   return

}

