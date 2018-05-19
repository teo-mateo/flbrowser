package rtorrent

import (
	"fmt"
	"encoding/xml"
	"log"
	"net/http"
	"errors"
	"io/ioutil"
	"bytes"
)

var Ru string
var Rp string
var RActive string
var RSession string
var RDownloads string

func Url() string {
	return fmt.Sprintf("http://%s:%s@h.bardici.ro:8008/RPC2", Ru, Rp)
}

func CallAndUnmarshal2(method string, id string) (*CmethodResponse, error){
	mc := NewMethodCall(method, &id)
	return CallAndUnmarshal(mc)
}

func CallAndUnmarshal(input MethodCall) (*CmethodResponse, error){
	bytes, err := Call(input)
	if err != nil {
		return nil, err
	}

	response := CmethodResponse{}
	err = xml.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func Call(input MethodCall) ([]byte, error){

	inputBytes, err := xml.Marshal(input)
	if err != nil{
		return nil, err
	}

	log.Println(string(inputBytes))

	buf := bytes.NewBuffer(inputBytes)

	client := http.Client{}
	log.Println("POST " + Url())
	req, err := http.NewRequest("POST", Url(), buf)
	if err != nil{
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil{
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil{
		return nil, err
	}

	log.Println(string(bytes))

	return bytes, nil
}

func RPC_system_listMethods() ([]string, error) {
	mc := MethodCall{
		MethodName:"system.listMethods",
	}

	response, err := CallAndUnmarshal(mc)
	if err != nil{
		return nil, err
	}

	return response.GetStrings()
}

func RPC_download_list() ([]string, error){

	mc := NewMethodCall("download_list", nil)
	response, err := CallAndUnmarshal(mc)
	if err != nil{
		return nil, err
	}

	return response.GetStrings()
}



func RPC_d_name(id string) (string, error){
	mc := NewMethodCall("d.name", &id)
	response, err := CallAndUnmarshal(mc)
	if err != nil{
		return "", err
	}
	return response.GetString()
}

func RPC_d_is_open(id string) (bool, error){
	mc := NewMethodCall("d.is_open", &id)
	response, err := CallAndUnmarshal(mc)
	if err!= nil{
		return false, err
	}
	return response.GetBool()
}

func RPC_d_is_active(id string) (bool, error){
	response, err := CallAndUnmarshal(NewMethodCall("d.is_active", &id))
	if err != nil{
		return false, err
	}
	return response.GetBool()
}

func RPC_d_is_hash_checked(id string) (bool, error){
	response, err := CallAndUnmarshal(NewMethodCall("d.is_hash_checked", &id))
	if err != nil{
		return false, err
	}
	return response.GetBool()
}

func RPC_d_up_total(id string) (int, error){
	mc := NewMethodCall("d.up.total", &id)
	response, err := CallAndUnmarshal(mc)
	if err != nil{
		return 0, err
	}
	return response.GetInt()
}

func RPC_d_down_total(id string) (int, error){
	mc := NewMethodCall("d.down.total", &id)
	response, err := CallAndUnmarshal(mc)
	if err != nil{
		return 0, err
	}
	return response.GetInt()
}

func RPC_id__bool(method string, id string) (bool, error){
	response, err := CallAndUnmarshal(NewMethodCall(method, &id))
	if err != nil{
		return false, err
	}
	return response.GetBool()
}

func RPC_id__int(method string, id string) (int, error){
	response, err := CallAndUnmarshal(NewMethodCall(method, &id))
	if err != nil{
		return 0, err
	}
	return response.GetInt()
}

func RPC_id__string(method string, id string) (string, error){
	response, err := CallAndUnmarshal(NewMethodCall(method, &id))
	if err != nil {
		return "", err
	}
	return response.GetString()
}

