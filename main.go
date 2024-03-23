package main

import (
	"context"
	"fmt"
	"time"

	"task/internal/flood_control"
)

func main() {
	fc := floodcontrol.NewFC(context.Background())

	for userID := 215; userID < 220; userID++ {
		for i := 0; i < 15; i++ {
			// первые 10 запросов пройдут проверку, остальные 5 - нет
			ans, err := fc.Check(context.Background(),int64(userID))
			fmt.Println(i, ans, err)
			time.Sleep(time.Second) // 1 sec
		}
	}
}