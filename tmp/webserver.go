package main

import (
	"encoding/json"
	"fmt"
	"github.com/tylertreat/BoomFilters"
	"io"
	//"math/rand"

	"bytes"
	"github.com/caio/go-tdigest"
	"net/http"
	"strconv"
)

type DSResponse struct {
	Status int
	Body   []string
}

//var buffer bytes.Buffer
var sbf = boom.NewDefaultStableBloomFilter(10000, 0.01)
var cms = boom.NewCountMinSketch(0.001, 0.99)
var topk = boom.NewTopK(0.001, 0.99, 5)
var hll, err = boom.NewDefaultHyperLogLog(0.1)
var t = tdigest.New(100)

func AddToBloom(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("key")
	fmt.Println(key)
	sbf.Add([]byte(key))
	s := make([]string, 0, 5)
	s = append(s, "success")
	res := &DSResponse{
		Status: 200,
		Body:   s}
	resB, _ := json.Marshal(res)
	io.WriteString(w, string(resB))
	//io.WriteString(w, key)
}
func CheckBloom(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	s := make([]string, 0, 1)
	if sbf.Test([]byte(key)) {

		s = append(s, "success")
		res := &DSResponse{
			Status: 200,
			Body:   s}
		resB, _ := json.Marshal(res)
		io.WriteString(w, string(resB))

		//io.WriteString(w,"Entry Found")

	} else {
		s = append(s, "No entry For Key")
		res := &DSResponse{
			Status: 200,
			Body:   s}
		resB, _ := json.Marshal(res)
		io.WriteString(w, string(resB))

		//io.WriteString(w,"Entry Not Found")
	}

}
func AddToCMS(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	cms.Add([]byte(key))
	s := make([]string, 0, 1)
	s = append(s, "success")
	res := &DSResponse{
		Status: 200,
		Body:   s}
	resB, _ := json.Marshal(res)
	io.WriteString(w, string(resB))

	//io.WriteString(w, "Added")
}
func GetCountForKey(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	i := cms.Count([]byte(key))
	str := strconv.FormatUint(i, 10)
	s := make([]string, 0, 1)
	s = append(s, str)
	res := &DSResponse{
		Status: 200,
		Body:   s}
	resB, _ := json.Marshal(res)

	//fmt.Println("Frequency =",i)
	io.WriteString(w, string(resB))
}

func AddToTopK(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	topk.Add([]byte(key))
}
func getTopK(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Entered")

	//buffer := bytes.NewBufferString("")
	s := make([]string, 0, 5)
	//var buffer = bytes.Buffer
	for _, element := range topk.Elements() {
		buffer := bytes.NewBufferString("")
		buffer.WriteString(string((element).Data))
		buffer.WriteString("-")
		str := strconv.FormatUint((element).Freq, 10)
		buffer.WriteString(str)
		//fmt.Println( (element).Freq)
		s = append(s, string(buffer.String()))
		//   buffer.WriteString(string(element))
	}
	res := &DSResponse{
		Status: 200,
		Body:   s}
	resB, _ := json.Marshal(res)
	io.WriteString(w, string(resB))
}

func AddToHLL(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	hll.Add([]byte(key))
}

func GetHLL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("count", hll.Count())
	//io.WriteString(w,hll.Count())
}
func AddToTDigest(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	i, err := strconv.ParseFloat(key, 64)
	if err != nil {
		t.Add(i, 1)
	}
}
func GetTDigest(w http.ResponseWriter, r *http.Request) {
	n := t.Quantile(0.5)
	str := strconv.FormatFloat(n, 'f', 6, 64)
	s := make([]string, 0, 1)
	s = append(s, str)
	res := &DSResponse{
		Status: 200,
		Body:   s}
	resB, _ := json.Marshal(res)
	io.WriteString(w, string(resB))
}
func main() {
	http.HandleFunc("/check", CheckBloom)
	http.HandleFunc("/add", AddToBloom)
	http.HandleFunc("/addKey", AddToCMS)

	http.HandleFunc("/checkKey", GetCountForKey)
	http.HandleFunc("/addK", AddToTopK)
	http.HandleFunc("/getK", getTopK)
	http.HandleFunc("/addToH", AddToHLL)
	http.HandleFunc("/getH", GetHLL)
	http.HandleFunc("/addtdigest", AddToTDigest)
	http.HandleFunc("/gettdigest", GetTDigest)
	fmt.Println("About to start")
	http.ListenAndServe(":8000", nil)
}
