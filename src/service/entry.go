package main

import (
	"fmt"
	"time"
	"net/http"
	"./tools"
	"./sql"
	"encoding/json"
	"math/rand"
	"sync"
	"runtime"
	"path/filepath"
	"os"
	"strings"
)

func main() {
	str := tools.UniqueValue()
	fmt.Println(str)

	res := tools.GeneratRandomNumber(1,6, 4)
	fmt.Println(res)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(1000000)
	fmt.Println(num)

	year := time.Now().Year()

	fmt.Println(year)

	fmt.Println(time.Now().Unix())

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))


	testStruct()
	// localhost:8080/login/
	//http.HandleFunc("/login/", LoginHandle)
	//http.ListenAndServe(":8080", nil)
	//fmt.Println("Hello world")

	//time.Sleep(10 * time.Second)



}

// 路由设置
func LoginHandle(w http.ResponseWriter, r *http.Request)  {
	tools.LogInfo("info log test")
	tools.LogInfo("hello world")
	fmt.Println("Login is running...")
	fmt.Fprint(w, "Hello")

	mysql.InitInfo()
	sql := "select * from app_user where id = ?"
	row, _:= mysql.FetchRow(sql, 1)
	fmt.Println(*row)
	tools.Debug(*row)
	v, err:= json.Marshal(*row)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, string(v))
	result := make(map[string]interface{})
	result["code"] = 100
	result["list"] = *row

	res, _ := json.Marshal(result)
	fmt.Fprint(w, string(res))

	insertSql := "insert app_user (name,password,salt) values (?,?,?)"
	id, err := mysql.InsertRow(insertSql, 1,3,4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

	res01 := make(map[string]string)
	res01["c"] = "c1"
	res01["f"] = "f1"
	res01["sfasdfs"] = "dsfs"
	res01["dddd"] = "12235722"
	res01["ad"] = ""
	res01["sign"] = "12235722"
	str := tools.DoMd5(res01)
	tools.LogInfo(str)
	fmt.Fprint(w,  str)

}


// 为类型添加方法  GO 语言面向对象与面向过程
type Integer int

func (a Integer) less (b Integer) bool  {
	return a < b
}

func Integer_Less(a Integer, b Integer) bool  {
	return a < b
}

// 结构体  1.对象， 2.对象的指针
type person struct {
	name string
	age int
	sex bool
}

// 对象初始化
// 1


func testStruct()  {
	// 对象类型初始化
	var per person
	per.name = "name"
	per.age = 18
	per.sex = true

	fmt.Println(per)
	//
	//per1 := person{"name", 18, true}
	//fmt.Println(per1)
	//
	//per2 := person{ name:"name", age:18}
	//fmt.Println(per2)

	// 对象指针初始化

	p1 := new(person)
	p1.name = "name"
	p1.age = 18
	p1.sex = true
	fmt.Println(p1)

	//p2 := &person{"name", 18, true}
	//fmt.Println(p2)

	//p3 := &person{name:"name", age:18}
	//fmt.Println(p3)

	doAny()
	path := currentPath()
	fmt.Println(path)
}

// 并发
func doAny()  {
	//runtime.GOMAXPROCS(1) // 设置逻辑处理器个数
	runtime.GOMAXPROCS(runtime.NumCPU()) // 根据电脑的实际物理核数 设置逻辑处理器个数
	var wg sync.WaitGroup  // 技术信号量
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i< 500 ;i++  {
			fmt.Println("A",i)
		}
	}()

	go func() {
		defer wg.Done()  // 执行完毕减法1
		for i:= 1; i < 100; i ++ {
			fmt.Println("B",i)
		}
	}()
	wg.Wait()  // 如果计数器大于0 就阻塞
}

func currentPath()  string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(os.Args)
	execpath, err := os.Executable()
	fmt.Println(execpath)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\","/", -1)
}
