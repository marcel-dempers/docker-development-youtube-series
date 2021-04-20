package main 

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
)

func test(){

	pods, err := clientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}