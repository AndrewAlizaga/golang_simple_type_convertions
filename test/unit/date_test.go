package unit_test

import (
	"log"
	"testing"

	datev1 "github.com/andrewAlizaga/golang_simple_type_convertions/internal/date"
)

func TestLayoutFormat(t *testing.T) {

	log.Println("INVOKE TEST - TestLayoutFormat")

	t.Log("starting test")

	layouts, err := datev1.GetDateLayout("04/02/22")

	if err != nil {
		t.Logf("TEST HAS FAILED")
		t.Fatalf(err.Error())
		t.Fail()
	} else {
		t.Log("Test successful, layouts: ")
		log.Println(layouts)

		for idx, lay := range(layouts) {
			log.Println(lay.Format())
			log.Println(idx)
		}
		
	}

}
