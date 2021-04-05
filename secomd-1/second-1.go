package main

import "fmt"
type Author struct {
	Name string             //名字
	VIP bool                //是否是高贵的带会员
	Icon string             //头像
	Signature string        //签名
	Focus int               //关注人数
}
type Video struct {
	VideoName string          //视频名字
	PlaybackTimes int        //播放次数
	Comments int             //评论数
	AuthorInformation Author  //作者信息
	VS []Video             //推荐视频
}


func main(){
var a Video
	b:=Video{VideoName: "一号推荐"}
	c:=Video{VideoName: "二号推荐"}
a=Video{
	AuthorInformation:
		Author{Name: "黑马",VIP: false,Signature: "go语言大法好",Focus: 900000},
	VideoName:"go语言入门",
	PlaybackTimes: 45000,
	Comments: 700,
VS:
	[]Video{b,c}}



fmt.Print(a)

}
