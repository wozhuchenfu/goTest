package elasticSearchTest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
	//"regexp"
)

type employee struct {
	First_name string
	Last_name string
	Age int
	About string
	Interests []string
}


func Add()  {

	//更新索引与添加索引相同
	client := &http.Client{}
	var emp1 employee
	emp1.Age = 23
	emp1.About = "I like to play computer"
	emp1.First_name = "张"
	emp1.Last_name = "三"
	emp1.Interests = []string{"basketball","football"}
	b,_ := json.Marshal(emp1)
	fmt.Println(string(b))
	fmt.Println("+++++++++++++++++")

	r,err := http.NewRequest(http.MethodPut,"http://localhost:9200/megacorp/employee/4?pretty",strings.NewReader(string(b)))
	resp,_ :=client.Do(r)
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	fmt.Println("===========================",resp.StatusCode)

}
/*
path 例：(megacorp/employee/1)
 */
func Search(path string)  {
	//client := &http.Client{}
	resp,err := http.Get("http://localhost:9200/"+path+"?pretty")
	if err!=nil {
		fmt.Println(err)
	}
	b,err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	fmt.Println("=============",resp.StatusCode,"=============")

}

func Delete(path string)  {
	client := &http.Client{}
	r,err := http.NewRequest(http.MethodDelete,"http://localhost:9200/"+path+"?pretty",nil)
	if err!=nil {
		fmt.Println(err)
	}
	resp,err := client.Do(r)
	if err!=nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	b,err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

}
/*
查询所有
localhost:9200/megacorp/employee/_search?pretty
 */

func FindAll()  {
	resp,err := http.Get("http://localhost:9200/megacorp/employee/_search?pretty")
	if err!=nil {
		fmt.Println(err)
	}
	b,err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

/*
条件查找
localhost:9200/megacorp/employee/_search?q=last_name:Smith&pretty
 */
func FindByName(name string)  {
	resp,err := http.Get("http://localhost:9200/megacorp/employee/_search?q=last_name:"+name+"&pretty")
	if err!=nil {
		fmt.Println(err)
	}
	b,err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

/*
GET 'localhost:9200/megacorp/employee/_search?pretty' -H 'Content-Type: application/json' -d'
{
    "query" : {
        "match" : {
            "last_name" : "Smith"
        }
    }
}'


 */

type query struct {
	Query *match
}
type match struct {
	Match *People
}

type People struct {
	Last_name string
}


func FindByJson(name string)  {
	p := &People{"123456"}
	m := &match{p}
	q := query{m}
	//jsons := "{query : {match : {last_name :"+ name+"}}"
	jsons,_ := json.Marshal(q)
	s := strings.ToLower(string(jsons))
	s2:=strings.Replace(s,"123456",name,-1)
	fmt.Println("s",s)
	fmt.Println("s2",s2)
	fmt.Println(string(jsons))
	client := &http.Client{}
	r,err := http.NewRequest(http.MethodGet,"http://localhost:9200/megacorp/employee/_search?pretty",strings.NewReader(s2))
	if err!=nil {
		fmt.Println(err)
	}
	r.Header.Set("Content-Type","application/json")
	defer r.Body.Close()
	res,err := client.Do(r)
	if err!=nil {
		fmt.Println(err)
	}
	b,err := ioutil.ReadAll(res.Body)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

}

/*
curl -XGET 'localhost:9200/megacorp/employee/_search?pretty' -H 'Content-Type: application/json' -d'
{
    "query" : {
        "bool": {
            "must": {
                "match" : {
                    "last_name" : "smith"
                }
            },
            "filter": {
                "range" : {
                    "age" : { "gt" : 30 }
                }
            }
        }
    }
}
'

 */

func FindByFilter(age int,name string)  {
	a := make(map[string]*map[string]*map[string]*map[string]*map[string]string)
	b := make(map[string]string)
	b["last_name"] = name
	c := make(map[string]*map[string]string)
	c["match"] = &b
	d := make(map[string]*map[string]*map[string]string)
	d["must"] = &c
	e := make(map[string]*map[string]*map[string]*map[string]string)
	e["bool"] = &d
	a["query"] = &e
	s,err := json.Marshal(a)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(string(s))


	f := make(map[string]int)
	g := make(map[string]map[string]int)
	h := make(map[string]map[string]map[string]int)
	i := make(map[string]map[string]map[string]map[string]int)
	f["gt"] = age
	g["age"] = f
	h["range"] = g
	i["filter"] = h
	js,err := json.Marshal(i)
	if err!=nil {
		fmt.Println(err)
	}
	jss := string(js)
	fmt.Println(jss)

    k := make([]interface{},2)
    k[0] = a
    k[1] = i
    //k = append(k,a,i)
    j,err:=json.Marshal(k)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(string(j))
	x := string(j)
	//rex, _ := regexp.Compile("p([a-z]+)ch")
	//xl := len(x)
	x = strings.Replace(x,"[","",-1)
	x = strings.Replace(x,"]","",-1)
	fmt.Println("=========strings.Replace==========")
	fmt.Println( x)
	r,err := http.NewRequest(http.MethodGet,"http://localhost:9200/megacorp/employee/_search?pretty",strings.NewReader(x))
	if err!=nil {
		fmt.Println(err)
	}
	r.Header.Set("Content-Type","application/json")
	client := &http.Client{}
	resp,err := client.Do(r)
	if err!=nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	bs,err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	fmt.Println("r.URL.RequestURI()",r.URL.RequestURI())
	fmt.Println("r.URL.RawQuery",r.URL.RawQuery)
	fmt.Println("请求参数：")
	for k,v := range r.URL.Query(){
		fmt.Println("k",k)
		fmt.Println("v",v)
		for a,b := range v{
			fmt.Println("a",a)
			fmt.Println("b",b)
		}
	}

}

//使用 HEAD 指令来检查文档是否存在






