package buycycle_test

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/Goboolean/fetch-server/internal/infrastructure/ws"
	"github.com/Goboolean/fetch-server/internal/infrastructure/ws/buycycle"
	"github.com/Goboolean/fetch-server/internal/infrastructure/ws/mock"
	"github.com/Goboolean/fetch-server/internal/util/env"
	"github.com/Goboolean/fetch-server/internal/util/withintime"
	"github.com/Goboolean/shared/pkg/resolver"
	"github.com/joho/godotenv"
)



var instance ws.Fetcher

var (
	count int = 0
	receiver ws.Receiver = mock.NewMockReceiver(func() {
		count++
	})
)





func SetupBuycycle() {
	instance = buycycle.New(&resolver.ConfigMap{
		"HOST": os.Getenv("BUYCYCLE_HOST"),
		"PORT": os.Getenv("BUYCYCLE_PORT"),
	}, context.Background(), receiver)
}

func TeardownBuycycle() {
	instance.Close()
}


func TestMain(m *testing.M) {

	if err := os.Chdir(env.Root); err != nil {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	SetupBuycycle()
	code := m.Run()
	TeardownBuycycle()

	os.Exit(code)
}




func Test_Constructor(t *testing.T) {

	if flag := isMarketOn(); flag {
		t.Skip()
	}
}


var (
	once sync.Once
	withinDurationChecker *withintime.WithinDurationChecker
	isMarketOnCache bool
)

// Struct withinDurationChecker is initialized with information of the Korea stock market.
// Value isMarketOnCache is cached at the time of first call, therefore inconsistency beween tests may not occur.
// U can get the value by calling IsMarketOn() function.
func isMarketOn() bool {

	once.Do(func() {
		var err error

		withinDurationChecker, err = withintime.New(&withintime.Option{
			Location: "Asia/Seoul",
			Inclusion: withintime.ConditionList{
				time.Monday:    &withintime.Condition{StartTime: "09:00", EndTime: "15:30"},
				time.Tuesday:   &withintime.Condition{StartTime: "09:00", EndTime: "15:30"},
				time.Wednesday: &withintime.Condition{StartTime: "09:00", EndTime: "15:30"},
				time.Thursday:  &withintime.Condition{StartTime: "09:00", EndTime: "15:30"},
				time.Friday:    &withintime.Condition{StartTime: "09:00", EndTime: "15:30"},
			},
			Exclusion: withintime.ConditionList{
				// TODO: add holiday
				},
			}, nil)

		if err != nil {
			panic(err)
		}

		isMarketOnCache = withinDurationChecker.IsTimeNowWithinDuration()
	})
	
	return isMarketOnCache
}