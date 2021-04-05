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
	VideoName string         //视频名字
	PlaybackTimes int        //播放次数
	Comments int             //评论数
	AuthorInformation Author  //作者信息
	VS []Video              //推荐视频
	ThumbsUp  int          //点赞数
	Collection int         //收藏数
	Coin     int            //投币数
}

func (r *Video) ThumbsUps()  {      //点赞
	r.ThumbsUp +=1
}
func (r *Video) Collections()  {        //收藏
	r.Collection +=1
}
func (r *Video) Coins()  {          //投币
	r.Coin +=1
}
func (r *Video) SanLian(){         //一键三连
	r.ThumbsUps()
	r.Collections()
	r.Coins()
}

func VideoRelease(author string,video string)Video{
return Video{AuthorInformation:Author{Name: author},VideoName: video}
}

func main(){
	a:=VideoRelease("黑马","go语言入门")
    fmt.Println(a)
	a.SanLian()
	fmt.Println(a)
}
