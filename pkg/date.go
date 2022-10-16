package pkg

import (
	datev1 "github.com/andrewAlizaga/golang_simple_type_convertions/internal/date"
)

//Access point for lib consumption
//case 1, get date string layout and format
func GetDateLayoutDateAndFormat(dateString string) (layouts []datev1.DateTimeLayouts, err error){
	layouts, err = datev1.GetDateLayout(dateString)
	return
}
