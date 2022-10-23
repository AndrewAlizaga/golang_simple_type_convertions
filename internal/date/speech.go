package date

import (
	"errors"
	"log"
	"strconv"
	"time"
)


func GetSpeech(date time.Time, language Language) (speech string, err error) {
	log.Println("INVOKE GetSpeech")

	transformedDate := date.Format("January 2, 2006")


	log.Println("transformed date: ", transformedDate)

	if err == nil {


		switch language {
			case LANGUAGE_SPANISH:
				speech = strconv.FormatInt(int64(date.Day()), 10) + " de " + SpanishMonths[date.Month()-1] + " del " + strconv.FormatInt(int64(date.Year()), 10)

			case LANGUAGE_ENGLISH:
				speech = EnglishMonths[date.Month()-1] + " " + strconv.FormatInt(int64(date.Day()), 10) + " of " + strconv.FormatInt(int64(date.Year()), 10)

			default: 
				err = errors.New("unsuported language passed")
		}
	}
	
	log.Println("SPEECH: ", speech)

	//Real Date detecting here
	//check := dashCheck(stringDate)
	return
}
