package utils

import (
	"testing"
	"time"

	"github.com/nzgogo/micro/codec"
)

type Product struct {
	PriceFixed *int64 `json:"priceFixed,omitempty" bson:"priceFixed,omitempty"`
}

func TestReplacePrice(t *testing.T) {
	price := int64(29999)
	product := Product{PriceFixed: &price}

	result := make(map[string]interface{})
	productBytes, err := codec.Marshal(product)
	if err = codec.Unmarshal(productBytes, &result); err != nil {
		return
	}
	//result["priceFixed"] = int64(299990000)
	//fmt.Println(reflect.TypeOf(result["priceFixed"]))
	//fmt.Println(fmt.Sprintf("%.0f", result["priceFixed"]))
	ReplacePrice(result, []string{
		"priceFixed",
	})
	if result["priceFixed"] != 299.99 {
		t.Error(
			"For", "priceFixed",
			"expected", 299.99,
			"got", result["priceFixed"],
		)
	}
}

func TestFirstDayOfISOWeek(t *testing.T) {
	//	year, week := time.Now().ISOWeek()
	loc, _ := time.LoadLocation("Pacific/Auckland")
	layout := "2006-01-02 15:04:05"

	// current week
	//begin := FirstDayOfISOWeek(year, week, time.Local)
	//expect, err := time.ParseInLocation(layout, "2018-07-30 00:00:00.000", loc)
	//if err != nil {
	//	t.Error(err.Error())
	//}
	//if !expect.Equal(begin) {
	//	t.Error("expected", expect,
	//		"got", begin)
	//}

	// get the starting time of the 1st week of 1985:
	begin := FirstDayOfISOWeek(1985, 1, time.Local)
	expect, err := time.ParseInLocation(layout, "1984-12-31 00:00:00.000", loc)
	if err != nil {
		t.Error(err.Error())
	}
	if !expect.Equal(begin) {
		t.Error("expected", expect,
			"got", begin)
	}
}
