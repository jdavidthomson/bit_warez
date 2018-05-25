package main

import (
    "net/http"
    "log"
    "encoding/json"
    "io/ioutil"
)
type ProcessJson struct{
	Process		map[string]interface{}		`json:"process"`
}

func SubmitProcess(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var process_json ProcessJson
	err = json.Unmarshal(body, &process_json)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
/*	log.Println(process_json.Process["process_name"].(string))
	arguments := process_json.Process["process_arguments"].([]interface{})
	for _, value := range arguments{
		log.Println(value.(map[string]interface{}))
	}
*/
	btc_address := GetNewAddress()
	log.Println(btc_address)
	json.NewEncoder(w).Encode(btc_address)
}
func PaidEndpoint(w http.ResponseWriter, r *http.Request){

}
