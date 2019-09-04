package main

import (
	"fmt"
	"html/template"
	"strings"
	"time"
	"net/http"
	"bufio"
	"os"
	"strconv"
//	"dbop"
)

var fhandle http.Handler

type GoodsInfo struct{
	Name string
	Price int
	//Input string
}

type DataInfo struct{
	NUM string
	MODLOG string
	LOG string
	HELPXD string
	//HELPSHD string
	Goods []* GoodsInfo
}

func LoadPrices(sel string)([]*GoodsInfo, int){
	ret:=make ([]* GoodsInfo,0,100)
	selprice:=0
	if file,err:=os.Open("prices.dat");err==nil{
		defer file.Close()
		rd:=bufio.NewReader(file)
		for {
			line,_,err:=rd.ReadLine()
			if err!=nil{
				break
			}
			data:=strings.TrimSpace(string(line))
			if data==""{
				continue
			}
			piece:=strings.Split(data,"----")
			if len(piece)!=2 {
				continue
			}
			if price,err:=strconv.Atoi(piece[0]);err!=nil{
				continue
			}else{
				if sel == piece[1] {
					selprice=price
				}
	//			text:=fmt.Sprintf("%s\t\t<input type=\"button\" value=\"%d\t\" onclick=\"ondhval('%s', value)\">", piece[1],price,piece[1])
				curgoods:= &GoodsInfo{piece[1], price}
				ret=append(ret,curgoods)
			}
		}
	}
	return ret,selprice
}

func gametime(w http.ResponseWriter, r *http.Request) {
	usr,_:=r.Cookie("logname")
	psw,_:=r.Cookie("passwd")
	if usr==nil{
		http.Redirect(w,r,"/",http.StatusFound)
		return
	}else if usr.Value!="ipad"{
			fmt.Fprintf(w,"请使用正确的用户登录后访问本页面\n")
			return
	}
	fmt.Fprintf(w,"User=%s, passwd=%s",usr.Value,psw.Value)
}

func xdinfo(w http.ResponseWriter, r *http.Request) {

	ck,_:=r.Cookie("logname")
	if ck==nil {
		http.Redirect(w, r, "/",http.StatusFound)
		return
	}else if ck.Value!="xd"{
			fmt.Fprintf(w,"请使用正确的用户登录后访问本页面\n")
			return
	}


	curxd := 0
	xdlog:="" //make([]string,0,50)
	modlog:=""
	alllog:=make([]string,0,50)
	allmodlog:=make([]string,0,50)

	filexd, err := os.Open("xd.dat")
	if err != nil {
		filexd, _ = os.Create("xd.dat")
		fmt.Fprintf(filexd,"0")
	} else {
		fmt.Fscanf(filexd, "%d%d", &curxd)
	}
	filexd.Close()

	filelog,err:=os.Open("log.dat")
	if err!=nil{
		filelog,_=os.Create("log.dat")
	}else{
		bf:=bufio.NewReader(filelog)
		for{
			line,err:=bf.ReadString('\n')
			if err!=nil{
				break
			}
			alllog=append(alllog,line)
		}
	}
	filelog.Close()
	cnt:=len(alllog)
	for i:=cnt-1;i>=0;i--{
		xdlog+=alllog[i]
	}

	if modfile,err:=os.Open("modlog.dat");err==nil{
/*	if err1!=nil{
		modfile,_=os.Create("modlog.dat")
	}else{*/
		bf:=bufio.NewReader(modfile)
		for{
			line,err:=bf.ReadString('\n')
			if err!=nil{
				break
			}
			allmodlog=append(allmodlog,line)
		}
		modfile.Close()
	}
	modcnt:=len(allmodlog)
	for i:=modcnt-1;i>=0;i--{
		modlog+=allmodlog[i]
	}

	var helpxd string
	filehlp,err:=os.Open("helpxd.dat")
	if err==nil{
		bf:=bufio.NewReader(filehlp)
		for{
			line,err:=bf.ReadString('\n')
			if err!=nil{
				break
			}
			helpxd+=line
		}
	}
	filehlp.Close()

/*	var helpshd string
	filehlp,err=os.Open("helpshd.dat")
	if err==nil{
		bf:=bufio.NewReader(filehlp)
		for{
			line,err:=bf.ReadString('\n')
			if err!=nil{
				break
			}
			helpshd+=line
		}
	}
	filehlp.Close()
*/

	if r.Method == "GET" {
		t, _ := template.ParseFiles("xdinfo.tpl")
		prices,_:=LoadPrices("")
		num:=DataInfo{
	//	num := make(map[string]string)
		NUM: fmt.Sprintf("%d", curxd),
		//SHD:	fmt.Sprintf("%d",curshd),
		MODLOG: modlog,
		LOG: xdlog,
		HELPXD: helpxd,
	//	HELPSHD: helpshd,
		Goods: prices,
		}
		t.Execute(w, num)
	}else{
		r.ParseForm()
		obj:=r.Form["obj"][0]
	//	spend,err:=strconv.Atoi(strings.TrimSpace(r.Form["spendxd"][0]))
		_,spend:=LoadPrices(obj)
//		spend:=prices[obj].Price
		if spend==0{
			fmt.Fprintf(w, "<script> alert(\"无此商品!\")</script>")
			return
		}
		if spend> curxd {
			fmt.Fprintf(w, "<script> alert(\"学豆不够!\")</script>")
			return
		}
		curxd-=spend
		tmpfile,_:=os.Create("xd.dat")
		fmt.Fprintf(tmpfile,"%d",curxd)
		tmpfile.Close()

		tm:=time.Now().Local()
		alllog=append(alllog,fmt.Sprintf("%d年%d月%d日 %d:%02d 用%d学豆 兑换了 %s\n", tm.Year(),tm.Month(),tm.Day(),tm.Hour(),tm.Minute(),spend,obj))
		tmpfile,_=os.Create("log.dat")
		for _,v:=range alllog {
			fmt.Fprintf(tmpfile,"%s",v)
		}
		tmpfile.Close()
//		fmt.Fprintf(w,"<script> alert(\"兑换成功\")</script>")
		http.Redirect(w, r, "xdinfo",http.StatusFound)
	}
}

func modify(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		curxd := 0
		file, err := os.Open("xd.dat")
		if err != nil {
			file, _ = os.Create("xd.dat")
		} else {
			fmt.Fscanf(file, "%d", &curxd)
		}
		file.Close()

		t, _ := template.ParseFiles("modify.tpl")
		num := make(map[string]string)
		strxd:=fmt.Sprintf("%d",curxd)
		num["Num"] = strxd
		num["LASTXD"]=strxd
		t.Execute(w, num)
	} else {
		r.ParseForm()
		xd := r.Form["xd"][0]
		reason:=r.Form["reason"][0]
		psw := r.Form["root"][0]
		if psw == "rootabc123" {
			nxd, err1 := strconv.Atoi(xd)
			if err1 != nil {
				fmt.Fprintf(w, "<script> alert(\"数字格式不正确!\")</script>")
			} else {
				file, _ := os.Create("xd.dat")
				fmt.Fprintf(file, "%d\n", nxd)
				file.Close()

				modlog,err:=os.OpenFile("modlog.dat",os.O_RDWR,0666)
				if err!=nil{
					modlog,_=os.Create("modlog.dat")
				}
				modlog.Seek(0,2)
				lastxd:=r.Form["lastxd"][0]
				tm:=time.Now().Local()
				//alllog=append(alllog,fmt.Sprintf("%d年%d月%d日 %d:%02d 用%d学豆 兑换了 %s\n", tm.Year(),tm.Month(),tm.Day(),tm.Hour(),tm.Minute(),spend,obj))
				fmt.Fprintf(modlog,"%d年%d月%d日 %d:%02d 由于:%s， 从 %s 变为 %d\n",tm.Year(),tm.Month(),tm.Day(),tm.Hour(),tm.Minute(),reason,lastxd,nxd)
				modlog.Close()
				fmt.Fprintf(w, "<script> alert(\"更新成功!\")</script>")
				fmt.Fprintf(w, "<font size= 20>当前学豆: %d </font>\n", nxd)
			}
		} else {
			fmt.Fprintf(w, "<script> alert(\"密码不正确!\")</script>")
		}
	}
}

func codes(w http.ResponseWriter, r *http.Request){
	ck,_:=r.Cookie("logname")
	if ck==nil {
		http.Redirect(w, r, "/",http.StatusFound)
	}else{
		if ck.Value!="mdd"{
			fmt.Fprintf(w,"请使用正确的用户登录后访问本页面\n")
		}else{
		h:=w.Header()
		h.Add("Pragma","no-cache")
		h.Add("Cache-Control","no-cache")
		fhandle.ServeHTTP(w,r)
		}
	}
}


func logon(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("a.test")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		psw := r.Form["password"][0]
		user := r.Form["username"][0]
		if psw == "123321" && user == "mdd" {
			//http.Redirect(w, r, "KZV2RJ7WZZ83ugUHReEDMNS0FE6J6wLQ9vydtMZP1m9O1zBX8woUOZAkBgns9mz0+9kGVO9AlH7PDGhoX5bn2WnXA6fDiWxfb2RecpMBb4mR13vkaL1ltvvLWNP7xSkrVdR2rhoy9beXQw/", http.StatusFound)
			//			fmt.Fprintf(w,"OK!\n")
			ck:=http.Cookie{Name:"logname",Value:"mdd"}
            http.SetCookie(w,&ck)
            http.Redirect(w, r, "codes/", http.StatusFound)
		} else if psw == "8192" && user == "xd" {
			ck:=http.Cookie{Name:"logname",Value:"xd"}
			http.SetCookie(w,&ck)
			http.Redirect(w, r, "xdinfo",http.StatusFound)
		} else if user=="ipad" && (psw=="gametime" || psw=="Gametime."){
			ckv:=http.Cookie{Name:"logname",Value:user}
			ckw:=http.Cookie{Name:"passwd",Value:psw}
			http.SetCookie(w,&ckv)
			http.SetCookie(w,&ckw)
			http.Redirect(w,r,"gametime",http.StatusFound)
		} else {
			fmt.Fprintln(w, "<font size= 20> 哼哼，用户名/密码不对哦~~~</font>")
			fmt.Fprintln(w, "<script> alert(\"小样，去死吧!\")</script>")
		}
	}
}

func main() {
	http.HandleFunc("/", logon)
	//	h:=http.FileServer(http.Dir("."))
	//http.HandleFunc("/",mainpage)
	fhandle=http.StripPrefix("/codes/",http.FileServer(http.Dir("/root/filesvr/data")))
	http.HandleFunc("/codes/",codes)
	http.HandleFunc("/login", logon)
	http.HandleFunc("/modify", modify)
	http.HandleFunc("/xdinfo", xdinfo)
	http.HandleFunc("/gametime", gametime)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Printf("Error:", err)
	}
}
