package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

type Page struct {
	title string
	body []byte
}

func (p *Page) save() error {
	filename := p.title + ".txt"
	return ioutil.WriteFile(filename,p.body,0600)
}

func loadPage(title string) (*Page,error){
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err!=nil {
		return nil,err
	}
	return (&Page{title: title, body: body}),nil
}


func hello( w http.ResponseWriter, r *http.Request ) {
	w.Write([]byte("Hello!!"))
}

func main() {

	p1 := &Page{ title: "Test", body: []byte("This is a test page")}
	p1.save()
	p2,err := loadPage("Test")
	if err != nil {
		fmt.Println("Error in reading file")
		return
	}
	fmt.Println(string(p2.body))

	http.HandleFunc("/",hello)
	http.ListenAndServe(":8080",nil)
	
}
