package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	floodcontrol "task/internal/flood_control"

	"gopkg.in/yaml.v3"
)

type TestParams struct {
	UserID_from  int
	UserID_to    int
	Num_requests int
}

func main() {
	var test_params TestParams

	yamlFile, err := os.ReadFile("test_config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlFile, &test_params)
	if err != nil {
		log.Fatal(err)
	}

	fc := floodcontrol.NewFC(context.Background())

	for userID := test_params.UserID_from; userID < test_params.UserID_to; userID++ {
		for i := 0; i < test_params.Num_requests; i++ {
			// первые 10 запросов пройдут проверку, остальные 5 - нет
			ans, err := fc.Check(context.Background(), int64(userID))
			fmt.Println(i, ans, err)
			time.Sleep(time.Second) // 1 sec
		}
	}
}
