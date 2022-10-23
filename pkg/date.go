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

/*

// LATAM DAYS ARRAY
var latamDays = [...]string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado", "Domingo"}

// USA DAYS ARRAY
var usaDays = [...]string{"Monday", "Thuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

// LATAM DAYS ARRAY
var latamMonths = [...]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio",
	"Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

// USA Months ARRAY
var usaMonths = [...]string{"January", "February", "March", "April", "May", "June", "July",
	"Agust", "September", "October", "November", "December"}


*/