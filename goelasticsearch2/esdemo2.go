package main

import (
	"fmt"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"os"
	"reflect"
)
var client *elastic.Client
var host = "http://127.0.0.1:9200/"

type Employee struct {
	FirstName string`json:"first_name"`
	LastName string`json:"last_name"`
	Age int`json:"age"`
	About string`json:"about"`
	Interests []string`json:"interests"`
}
func init(){ //这个会自动调用
	errlogg := log.New(os.Stdout, "App", log.LstdFlags)
	var err error
	client,err=elastic.NewClient(elastic.SetErrorLog(errlogg),elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	version, _ := client.ElasticsearchVersion(host)
	fmt.Printf("Elasticsearch version %s\n", version)

}
//新增的核心功能
func IndexAdd(company string,tp string,e Employee,id string,ctx context.Context){
	put, err := client.Index().Index(company).Type(tp).Id(id).BodyJson(e).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put.Id, put.Index, put.Type)
}
//新增，需要调用上面的功能
func create()  {
	e1 := Employee{"Jane", "Smith", 32, "I like to collect rock albums", []string{"music","foods","handbags"}}
	e2:= Employee{"Jack", "Smith", 25, "I like to go rock climbing", []string{"music","sports"}}
	e3 := Employee{"Jess", "Smith", 18, "I like to build cabinets\"", []string{"music","sex","money"}}

	company:="kennys"
    tp:="employee"
    ctx:=context.Background()
    IndexAdd(company,tp,e1,"1",ctx)
    IndexAdd(company,tp,e2,"2",ctx)
    IndexAdd(company,tp,e3,"3",ctx)

}
//删除
func delete() {
	res, err := client.Delete().Index("kennys").
		Type("employee").
		Id("1").
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}
//修改
func update() {
	res, err := client.Update().
		Index("kennys").
		Type("employee").
		Id("2").
		Doc(map[string]interface{}{"age": 88}).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)
}

//查找
func gets() {
	//通过id查找
	get1, err := client.Get().Index("kennys").Type("employee").Id("1").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
}
//搜索
func query() {
	var res *elastic.SearchResult
	var err error
	//取所有
	res, err = client.Search("kennys").Type("employee").Do(context.Background())
	printEmployee(res, err)
	//字段相等
	q := elastic.NewQueryStringQuery("last_name:Smith")
	res, err = client.Search("kennys").Type("employee").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	printEmployee(res, err)
	//条件查询
	//年龄大于30岁的
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	res, err = client.Search("kennys").Type("employee").Query(q).Do(context.Background())
	printEmployee(res, err)
	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
	res, err = client.Search("kennys").Type("employee").Query(matchPhraseQuery).Do(context.Background())
	printEmployee(res, err)
	//分析 interests
	aggs := elastic.NewTermsAggregation().Field("interests")
	res, err = client.Search("kennys").Type("employee").Aggregation("all_interests", aggs).Do(context.Background())
	printEmployee(res, err)
}
//简单分页
func list(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := client.Search("kennys").
		Type("employee").
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	printEmployee(res, err)
}
//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}

func main(){
   //create()
	//query()
	gets()
}