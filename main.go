package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"
	"syscall/js"
)

var (
	host string
	hash string
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("url", js.FuncOf(getCurrentUrl))
	js.Global().Set("copy", js.FuncOf(copyURL))
	fmt.Println("Initialized!")
	<-c
}
func getCurrentUrl(this js.Value, args []js.Value) interface{} {
	//fmt.Println("Getting url...")
	host = args[0].String()
	hash = strings.ReplaceAll(args[1].String(), "#", "")
	//encoded := strings.ReplaceAll(args[0].String(), "#", "")

	fmt.Println("Host:" + host)
	fmt.Println("Hash:" + hash)

	if hash != "" {
		text := decode(hash)
		document := js.Global().Get("document")
		editor := document.Call("getElementById", "editor")
		editor.Set("value", text)
	}

	return js.ValueOf(nil)
}

func copyURL(this js.Value, args []js.Value) interface{} {
	//host = args[0].String()
	fmt.Println("Copy")

	document := js.Global().Get("document")
	editor := document.Call("getElementById", "editor")
	text := editor.Get("value").String()
	//fmt.Println("Text\n" + text)
	encoded := encode(text)

	url := host + "#" + encoded
	fmt.Println("URL:" + url)

	js.Global().Get("navigator").Get("clipboard").Call("writeText", js.ValueOf(url))

	return nil
}
func encode(text string) string {
	textbytes := []byte(text)

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	_, err := gz.Write(textbytes)
	if err != nil {
		fmt.Println(err)
	}
	gz.Close()

	return base64.URLEncoding.EncodeToString(b.Bytes())
}
func decode(encoded string) string {

	decodedbytes, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Print("error: ")
		fmt.Println(err)
	}
	gz, err1 := gzip.NewReader(bytes.NewReader(decodedbytes))
	if err1 != nil {
		fmt.Println(err1)
	}
	//fix for Go >1.13
	decoded, err2 := ioutil.ReadAll(gz)
	if err2 != nil {
		fmt.Println(err2)
	}

	return string(decoded)
}
