package model_test

import (
	"testing"
	"time"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewDate(t *testing.T) {
	inputTime := time.Date(2024, 9, 22, 15, 0, 0, 0, time.UTC)
	expectedTime := time.Date(2024, 9, 22, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))

	date := model.NewDate(inputTime)

	assert.Equal(t, expectedTime, time.Time(date))
}

func TestString_Date(t *testing.T) {
	inputTime := time.Date(2024, 9, 22, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
	date := model.NewDate(inputTime)

	assert.Equal(t, "2024-09-22", date.String())
}
