package main

import (
	//"context"
	"time"
		"fmt"
	//"net/http"
	//"strconv"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//log "github.com/sirupsen/logrus"
)

func TestSchedule(){
	for {
			fmt.Println("sleeping...")
			time.Sleep(20 * time.Second)
	}
}