package functional_test

import (
	"errors"
	"testing"
	"time"

	"github.com/andrewAlizaga/golang_simple_type_convertions/internal/date"
)

func TestCaseGetDateLayout(t *testing.T) {
	t.Log("INVOKE TEST - TestCaseGetDateLayout")

	result, err := date.GetDateLayout("Mon")

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	} else {
		t.Log("CONVERTION RESULTS: ")
		t.Log(result)

		if len(result) == 0 {
			err = errors.New("EMPTY ARRAY, NO ABLE TO FIND DATE STRING MATCHES, CHECK THE STRING IS COMPATIBLE O NOT")
			t.Logf(err.Error())
			t.Fail()
		} else {
			for idx, lay := range result {
				t.Log(idx)
				t.Log("layout: ", lay)
				t.Log("format: ", lay.Format())
			}
		}
	}

}

func TestDateToSpeech(t *testing.T) {
	t.Log("INVOKE TEST - TestDateToSpeech")

	dateTime := time.Now()

	t.Log("TESTING DATE: ")
	t.Log(dateTime)

	speech, err := date.GetSpeech(dateTime, date.LANGUAGE_SPANISH)

	if err == nil {
		t.Log("GOT SPEECH: ", speech)
	}else {
		t.Fatalf(err.Error())
		return
	}

	//call controller
}

func TestGetDateStringToSpeech(t *testing.T) {
	t.Log("INVOKE TEST - TestCaseGetDateLayout")

	result, err := date.GetDateLayout("Mon")

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	} else {
		t.Log("CONVERTION RESULTS: ")
		t.Log(result)

		if len(result) == 0 {
			err = errors.New("EMPTY ARRAY, NO ABLE TO FIND DATE STRING MATCHES, CHECK THE STRING IS COMPATIBLE O NOT")
			t.Logf(err.Error())
			t.Fail()
		} else {
			for idx, lay := range result {
				t.Log(idx)
				t.Log("layout: ", lay)
				t.Log("format: ", lay.Format())
			}
		}
	}
}
