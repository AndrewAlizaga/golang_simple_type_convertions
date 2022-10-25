package pkg

import (
	"time"

	datev1 "github.com/andrewAlizaga/golang_simple_type_convertions/internal/date"
)

//Access point for lib consumption
//case 1, get date string layout and format
func GetDateLayoutDateAndFormat(dateString string) (layouts []datev1.DateTimeLayouts, err error){
	layouts, err = datev1.GetDateLayout(dateString)
	return
}


//Access point for lib consumption
//case 2, date to redable format
func DateParseSpeech(date time.Time, language datev1.Language) (speech string, err error){
	speech, err = datev1.GetSpeech(date, language)
	return
}

//Access point for lib consumption
//Get 
//func FindDatesOnText(speech string) (dates []time.Time){
//	speech, err = datev1.GetSpeech(date, language)
//	return
//}

