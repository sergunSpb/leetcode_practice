package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var service_1_url = "http://service1/fillItems/"
var service_2_url = "http://service2/scoreItems/"
var service_3_url = "http://service3/logItems/"

func main() {
	urls := []string{"ya.ru", "mail.ru"}
	res := make(chan SomeResponse)
	wg := sync.WaitGroup{}
	sem := make(chan interface{}, 1000)
	for _, v := range urls {
		wg.Add(1)
		go ProcessUrl(v, &wg, sem, res)
	}
	wg.Wait()
	for i := range res {
		fmt.Println(i)
	}
}

func ProcessUrl(url string, wg *sync.WaitGroup, sem chan interface{}, result chan SomeResponse) {
	defer func() {
		wg.Done()
		<-sem
	}()
	sem <- nil

	res, err := http.Get("url")
	if err != nil {
		fmt.Println("Got some error", err)
		return
	}
	resp_struct := struct {
		item_ids []int
	}{}

	json.NewDecoder(res.Body).Decode(&resp_struct)

	res_chan := make(chan SomeResponse)
	go func() {
		wg := sync.WaitGroup{}
		for _, item := range resp_struct.item_ids {
			wg.Add(1)
			go func() {
				defer wg.Done()
				res1, _ := http.Get(fmt.Sprintf("%s%s", service_1_url, item))
				res2, _ := http.Get(fmt.Sprintf("%s%s", service_2_url, item))
				res3, _ := http.Get(fmt.Sprintf("%s%s", service_3_url, item))
				res_chan <- BussinesLogic(res1, res2, res3)
			}()
		}
		wg.Wait()
		close(res_chan)
	}()
	for i := range res_chan {
		result <- i
	}
}

func BussinesLogic(s1, s2, s3 *http.Response) SomeResponse {
	return SomeResponse{}
}

type SomeResponse struct{}
