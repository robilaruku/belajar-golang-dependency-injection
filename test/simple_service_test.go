package test

import (
	"fmt"
	"testing"

	"robi/belajar-golang-restful-api/simple"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService()
	fmt.Println(err)
	fmt.Println(simpleService)
}