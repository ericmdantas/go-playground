package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	r, _ := http.Get("http://jsonplaceholder.typicode.com/photos")
	bBody, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	ioutil.WriteFile("r.json", bBody, 0644)

	var bts bytes.Buffer

	gw := gzip.NewWriter(&bts)
	gw.Write(bBody)
	gw.Close()

	ioutil.WriteFile("r.gz", bts.Bytes(), 0644)

	jf, _ := os.Open("r.json")
	gf, _ := os.Open("r.gz")
	stat, _ := jf.Stat()
	defer jf.Close()

	rb := make([]byte, stat.Size())

	rG, _ := gzip.NewReader(gf)
	rG.Read(rb)

	ioutil.WriteFile("r2.json", rb, 0644)
}
