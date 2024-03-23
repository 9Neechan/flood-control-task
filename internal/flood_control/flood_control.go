package floodcontrol

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"task/internal/configs"

	"github.com/redis/go-redis/v9"
)

var DB *redis.Client = configs.ConnectRedisDB()
var K int = configs.EnvMaxReq()
var N int64 = configs.EnvInterval()

type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}

func NewFC(ctx context.Context) *FC {
	return &FC{
		Ctx:    ctx,
	}
}

type FC struct {
	// Тип NumCheckPerSecond реализует интерфейс, реализуя его методы
	Ctx    context.Context
}

func (fc *FC) Check(ctx context.Context, userID int64) (bool, error) {
	userID_Str := strconv.FormatInt(userID, 10)
	curTime := time.Now().Unix()

	// Вставляем в БД текущий запрос
	err := DB.LPush(ctx, userID_Str, curTime).Err()
	if err != nil {
		return false, err
	}

	// Достаем из БД все предыдущие запросы по UserID
	vals, err := DB.LRange(ctx, userID_Str, 0, -1).Result()
	if err != nil {
		return false, err
	}

	// Кол-во запросов, входящих в проверяемый промежуток времени
	num_req := len(vals)

	// обрезаем данные в БД, которые не входят в проверяемый промежуток времени
	for ind, el := range vals {
		num, _ := strconv.ParseInt(el, 10, 64)
		if num == curTime - N {
			err := DB.LTrim(ctx, userID_Str, int64(ind), int64(len(vals)-1)).Err()
			if err != nil {
				fmt.Println(err)
			}
			num_req = len(vals) - ind
			break
		}
	}

	// Проверяем условие флуд-контроля
	if num_req > K {
		return false, &RequestsLimitError{}
	}

	return true, nil
}

type RequestsLimitError struct{}

func (e *RequestsLimitError) Error() string {
	return "The limit of the maximum allowed number of requests has been reached according to the flood control rules"
}
