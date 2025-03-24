package main

import (
	"fmt"
	"log"
	"time"

	"github.com/casbin/casbin/v2"
)

func main() {
	init_start := time.Now()
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		log.Fatalf("unable to create Casbin enforcer: %v", err)
	}
	init_elapsed := time.Since(init_start)
	fmt.Printf("init took %s\n", init_elapsed)

	start := time.Now()
	if res, _ := e.Enforce("amy", "company", "alarm_management"); res {
		// permit alice to read data1
		fmt.Printf("200: %s\n", res)
	} else {
		// deny the request, show an error
		fmt.Printf("403: %s\n", res)
	}

	elapsed := time.Since(start)
	fmt.Printf("1st took %s\n", elapsed)

	start2 := time.Now()
	if res, _ := e.Enforce("amy", "eu_site/wfjakf", "view_detail6"); res {
		// permit alice to read data1
		fmt.Printf("200: %s\n", res)
	} else {
		// deny the request, show an error
		fmt.Printf("403: %s\n", res)
	}

	elapsed2 := time.Since(start2)
	fmt.Printf("last took %s\n", elapsed2)
}
