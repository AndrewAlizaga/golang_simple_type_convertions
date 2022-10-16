package date

import (
	"log"
)

func GetDateLayot(stringDate string) (dateTimeLayout DateTimeLayouts, err error){
	log.Println("INVOKE GetDateLayot")

	format, err := DATE_LAYOUT_1.Format()

	if err == nil {
		log.Println("DATE LAYOUT FORMATTING TEST: ", format)
	}else {
		log.Println(err.Error())
	}
	
	return
}
