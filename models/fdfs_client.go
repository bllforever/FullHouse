package models
import (
"fmt"
    "github.com/weilaihui/fdfs_client"
)
func FDFSUploadByFilename(filename string)(groupname string,remotefileid string,err error){
    fdfsClient, err0:=fdfs_client.NewFdfsClient("/home/itheima/workspace/go/src/github.com/weilaihui/fdfs_client/client.conf")
    if err != nil {
          fmt.Println("NewFdfsClient err=",err0)
          return "","",err0
    }
    uploadResponse, err1:= fdfsClient.UploadByFilename(filename)
    if err != nil {
        fmt.Println("fdfs_client.UploadByfilename error %s", err1)
          return "","",err1

    }
    groupname= uploadResponse.GroupName;
    remotefileid=uploadResponse.RemoteFileId;
    err=nil
    return
}
func FDFSUploadByBuffer(fileBuffer []byte, suffix string)(groupname string,remotefileid string,err error){
    fdfsClient, err0:=fdfs_client.NewFdfsClient("/home/itheima/workspace/go/src/github.com/weilaihui/fdfs_client/client.conf") 
    if err0 != nil {
          fmt.Println("NewFdfsClient err=",err0)
          return "","",err0

    }
  //  fileBuffer := make([]byte, fileSize)

  uploadResponse, err1:= fdfsClient.UploadByBuffer(fileBuffer, suffix)
    if err1 != nil {
          fmt.Println("fdfsClient.UploadByBuffer err=",err1)
          return "","",err1
    }

    groupname= uploadResponse.GroupName;
    remotefileid=uploadResponse.RemoteFileId;
    err=nil
    return
}
