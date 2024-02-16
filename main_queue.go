package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"strings"
)

func getLinksForDownload() string {
	contents, _ := ioutil.ReadFile("examples.txt")
	return string(contents)
}

func makeArrayWithLinks(myString string) []string {
	arr := strings.FieldsFunc(myString, func(r rune) bool {
		return r == ','
	})
	return arr
}
func prepareUrls(links []string) *list.List {
	queue := list.New()
	for _, link := range links {
		queue.PushBack(link)
	}
	return queue
}
func getDataForQueue() *list.List {
	arrayWithUrl := makeArrayWithLinks(getLinksForDownload())
	queue := prepareUrls(arrayWithUrl)
	return queue
}

func awaitTask(myValue string, myUrl string) <-chan string {
	fmt.Println(myValue)
	fmt.Println("Starting Task...")

	c := make(chan string)

	go func() {
		resp, err := http.Get(myUrl)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		c <- string(body)

		fmt.Println("...Done!")
	}()

	return c
}

func getMethods(myQueue *list.List) {
	fooType := reflect.TypeOf(myQueue)
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		fmt.Println(method.Name)

	}
}

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
