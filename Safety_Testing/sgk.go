package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"crypto/tls"
)


//配置host 或者在国外主机上执行
//104.223.12.164 qq.findmima.com qun.findmima.com
//208.64.31.157 usa.findmima.com

//https client
var tr = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
var client = &http.Client{Transport: tr}



func main(){
	key:="12888888" //搜索的关键字，可以是QQ、邮箱、用户名、手机号
	qq:="12888888" //搜索的QQ号
	souqq(qq)
	souqun(qq)
	souusa(key)
	//sou163() //http://163.donothackme.club/search  keyword:12888888@qq.com _token:KyDZW0UnntinG6ftpbXCz1veajecoaNkIIN27bW1
}


func souqq(key string)  {
	//tabs from https://qq.findmima.com/ajax.php?act=database
	tabs := []string{"qq_mima_1","qq_mima_2","qq_mima_3","qq_mima_4","qq_mima_5","qq_mima_6"}
	url := "https://qq.findmima.com/ajax.php?act=select"//qq密码
	sel := "1" //
	mac := "1"//左侧匹配
	contents:= ""
	content:= ""
	for _,tab := range tabs {
		//fmt.Println("\nqq tab:", tab)
		content = HttpPostForm(sel, mac, key, tab, url)
		//fmt.Println("a"+content+"b")
		if (content) != ""{
        contents = contents+content+"\n"
    	}
    }
    fmt.Println("qq "+key+" contents:"+ contents)
}


func souqun(key string)  {
	tabs := []string{"QQGroupData1","QQGroupData11","QQGroupData2","QQGroupData3","QQGroupData4","QQGroupData5","QQGroupData6","QQGroupData7","QQGroupData8","QQGroupData9"}
	url := "https://qun.findmima.com/ajax.php?act=select"//qun信息
	mac := "2"//精确匹配
	sel := "1" //
	contents:= ""
	content:= ""
	for _,tab := range tabs {
		content =HttpPostForm(sel, mac, key, tab, url)
        if len(content) >0{
        contents = contents+content+"\n"
    	}
    }
    fmt.Println("qun "+key+" contents:", contents)
}

func souusa(key string)  {
	//泄露网站列表 https://usa.findmima.com/info/list.html
	tabs := []string{"email1","email2","email3","email4","email5","mydb1","mydb10","mydb11","mydb12","mydb13","mydb2","mydb3","mydb4","mydb5","mydb6","mydb7","mydb8","mydb9","sgk2017"}
	url := "https://usa.findmima.com/ajax.php?act=select"//相关密码
	sel := "3" //
	mac := "1"//左侧匹配
	contents:= ""
	content:= ""
	for _,tab := range tabs {
		content = HttpPostForm(sel, mac, key, tab, url)
		if (content)  != ""{
        contents = contents+content+"\n"
    	}
    }
    fmt.Println("usa "+key+" contents:", contents)

}


func HttpPostForm(sel, mac, key, tab, urls string) string {
	v := url.Values{}
	v.Set("select_act", sel)
	v.Set("match_act", mac)
	v.Set("key", key)
	v.Set("table", tab)
	
	param := ioutil.NopCloser(strings.NewReader(v.Encode()))
	//fmt.Println(v) //打印post正文
	req, _ := http.NewRequest("POST", urls, param)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	//req.Header.Set("Referer", "https://qq.findmima.com/")
	req.Header.Set("Cookie", "__cfduid=d903a303beef26234448e4d4a154bfb491487920839; PHPSESSID=g1pa5cb15thsmf4jq74nj0so32; _ga=GA1.2.490953270.1487920721; _gid=GA1.2.1577856359.1504957732'")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(data)) //打印返回
	return string(data)

}