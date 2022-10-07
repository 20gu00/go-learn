package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi, this is home page")
}

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	//全部读取,流一般都是读一次,全部读
	body, err := io.ReadAll(r.Body) //Body
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		// 记住要返回，不然就还会执行后面的代码

		return
	}
	// 类型转换，将 []byte 转换为 string
	fmt.Fprintf(w, "read the data: %s \n", string(body))

	//这种地方容易错,比如框架没设计好,或者使用了中间件已经读取了body,比如打日志
	//所以它是无法判断是读取过了body还是没读取到,还是body就是这样
	// 尝试再次读取，啥也读不到，但是也不会报错
	body, err = io.ReadAll(r.Body)
	if err != nil {
		// 不会进来这里
		fmt.Fprintf(w, "read the data one more time got error: %v", err)
		return
	}
	fmt.Fprintf(w, "read the data one more time: [%s] and read data length %d \n", string(body), len(body))
}

func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	//GetBody原则上可以重复读,但原生的http.Request默认值是nil
	//一般来说不会读取多次
	if r.GetBody == nil {
		fmt.Fprint(w, "GetBody is nil \n")
	} else {
		fmt.Fprintf(w, "GetBody not nil \n")
	}
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	//查询参数
	//http://ip:port/path?a=b&aa=bb
	values := r.URL.Query() //map[sting][]string
	fmt.Fprintf(w, "query is %v\n", values)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	//实际上不一定拿得到值,很多事空的,唯一一个百分百拿到的是path
	//request.URL可能很多字段都是空的
	//随便加个header测试a-Ba-aa:20
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}

func header(w http.ResponseWriter, r *http.Request) {
	//会自动帮你格式化好
	//一般X开头表示自定义的header,会自动格式化好大小写
	fmt.Fprintf(w, "header is %v\n", r.Header)
}

func form(w http.ResponseWriter, r *http.Request) {
	//form表单,测试时也是使用参数?a=b&aa=bb
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
	//使用Form建议加上个hearder Content-Type: application/x-www-form-urlencoded
	//表单,实际还是用json通信比较多
	err := r.ParseForm() //必须先调用ParseForm才有效
	if err != nil {
		fmt.Fprintf(w, "parse form error %v\n", r.Form)
	}
	fmt.Fprintf(w, "after parse form %v\n", r.Form)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/body/once", readBodyOnce)
	http.HandleFunc("/body/multi", getBodyIsNil)
	//http://127.0.0.1:8080/url/query?name=cjq,y&n=y
	http.HandleFunc("/url/query", queryParams)
	http.HandleFunc("/header", header)
	http.HandleFunc("/wholeUrl", wholeUrl)
	http.HandleFunc("/form", form)

	//可以监听多个端口,多个服务暴露api,实际上还是同一个进程
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

//Get方法发不了body
