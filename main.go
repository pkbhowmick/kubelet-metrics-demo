package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	URL = "http://prometheus-kube-prometheus-kubelet.kube-system.svc:10250/metrics/cadvisor"
)

func main()  {
	res, err := http.Get(URL)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
