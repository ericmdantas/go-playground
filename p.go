package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
)

func main() {
	bf, _ := ioutil.ReadFile("f.txt")

	var w bytes.Buffer

	gw := gzip.NewWriter(&w)
	gzip.NewWriterLevel(gw, gzip.BestCompression)
	gw.Write(bf)
	defer gw.Close()

	ioutil.WriteFile("f.gz", w.Bytes(), 0644)
	f1, _ := os.Open("f.gz")
	defer f1.Close()

	gr, _ := gzip.NewReader(f1)
	defer gr.Close()

	b, _ := ioutil.ReadAll(f1)

	ioutil.WriteFile("f2.txt", b, 0644)
}
