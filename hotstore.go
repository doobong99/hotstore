package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 'http://apis.data.go.kr/B553077/api/open/sdsc/storeListInUpjong?ServiceKey=HBL%2FY0u39MG3pXAvaoRngGVc96fOUxmYD2aLu6L0MCCZjDure7gOD2e1ikw0kjRkPg%2F2pu9J0znei3JR1Geq3g%3D%3D&type=json&divId=indsLclsCd&key=Q&numOfRows=10&pageNo=1'

	arg_numOfRows := os.Getenv("NUMROWS")
	arg_pageNo := os.Getenv("PAGENO")

	ServiceKey := "HBL%2FY0u39MG3pXAvaoRngGVc96fOUxmYD2aLu6L0MCCZjDure7gOD2e1ikw0kjRkPg%2F2pu9J0znei3JR1Geq3g%3D%3D"
	reqUrl := "http://apis.data.go.kr/B553077/api/open/sdsc/storeListInUpjong?ServiceKey="
	var buf bytes.Buffer
	buf.WriteString(reqUrl)
	buf.WriteString(ServiceKey)

	//req, err := http.NewRequest("GET", "http://openapi.airkorea.or.kr/openapi/services/rest/ArpltnInforInqireSvc/getCtprvnRltmMesureDnsty?ServiceKey=BAQxmz5Vk%2FdlnOIOZtAHqmWSCv1PN7nvFsInlUu8xUTAKZGvabiPxPwv6cxIUhfJqS1p08DFyI5HoBy41FssWQ%3D%3D", nil)
	req, err := http.NewRequest("GET", buf.String(), nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	param := req.URL.Query()

	param.Add("type", "json")
	param.Add("divId", "indsLclsCd")
	param.Add("key", "Q")
	param.Add("numOfRows", arg_numOfRows)
	param.Add("pageNo", arg_pageNo)

	req.URL.RawQuery = param.Encode()

	fmt.Println(req.URL.String())
	fmt.Println("--------------------------------------------------------------------------------------------------")

	// Client객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//-------------------------------------------------------------------------------------------------------------------------------------
	//데이터 가공 및 출력
	//-------------------------------------------------------------------------------------------------------------------------------------

	// 결과 출력(string)
	resBytes, _ := ioutil.ReadAll(resp.Body)
	str := string(resBytes) //바이트를 문자열로
	fmt.Println(str)
	fmt.Println("1--------------------------------------------------------------------------------------------------")

	// 결과 출력(Json)
	var jsonResponse JsonResponse
	err = json.Unmarshal(resBytes, &jsonResponse)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonResponse)

	jSlice := jsonResponse.Body.Items

	for i := range jSlice {
		fmt.Println(jSlice[i].BizesNm, jSlice[i].LnoAdr, jSlice[i].Lon, jSlice[i].Lat)
	}

	fmt.Println("2--------------------------------------------------------------------------------------------------")

	//-------------------------------------------------------------------------------------------------------------------------------------
	//mysql connection
	//-------------------------------------------------------------------------------------------------------------------------------------

	// db, err := sql.Open("mysql", "jonas:ahffk18!@tcp(127.0.0.1:3306)/test")
	// db, err := sql.Open("mysql", "admin:EvzdKZX1XR@tcp(10.178.230.98:3306)/hotspot_store")
	db, err := sql.Open("mysql", "admin:EvzdKZX1XR@tcp(169.56.96.234:3306)/hotspot_store")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// insert parsed data to mysql
	for i := range jSlice {

		ss, err := strconv.Atoi(jSlice[i].BizesID)
		if err != nil {
			panic(err)
		}

		result, err := db.Exec("INSERT INTO store VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", i+70000, nil, "", jSlice[i].Lat, jSlice[i].Lon, jSlice[i].IndsSclsNm, ss, jSlice[i].LnoAdr, jSlice[i].BizesNm, jSlice[i].RdnmAdr)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(result)
	}

}
