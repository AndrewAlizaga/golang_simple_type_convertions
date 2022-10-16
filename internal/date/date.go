package date

import (
	"log"
	"time"
)

/*
func dashCheck(stringDate string) (matched bool) {
	matched = strings.Contains(stringDate, "-")
	return
}

func slashCheck(stringDate string) (matched bool) {
	matched = strings.Contains(stringDate, "/")
	return
}

func timeCheck(stringDate string) (matched bool) {
	matched = strings.Contains(stringDate, ":")
	return
}

func tZoneCheck(stringDate string) (matched bool) {
	matched = strings.Contains(stringDate, "T")
	return
}
*/

func GetDateLayot(stringDate string) (matchingLayouts []DateTimeLayouts, err error) {
	log.Println("INVOKE GetDateLayot")

	format := DATE_LAYOUT_1.Format()
	log.Println("DATE LAYOUT FORMATT`ING TEST: ", format)

	//Real Date detecting here
	//check := dashCheck(stringDate)
	var convertion time.Time = time.Now()

	convertion, convertionErr := time.Parse(DATE_LAYOUT_1.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_LAYOUT_1)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_LAYOUT_2.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_LAYOUT_2)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_LAYOUT_3.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_LAYOUT_3)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_LAYOUT_4.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_LAYOUT_4)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_LAYOUT_5.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_LAYOUT_5)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_LAYOUT_6.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_LAYOUT_6)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_LAYOUT_7.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_LAYOUT_7)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_LAYOUT_8.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_LAYOUT_8)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_LAYOUT_9.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_LAYOUT_9)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_LAYOUT_10.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_LAYOUT_10)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_LAYOUT_11.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_LAYOUT_11)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_LAYOUT_12.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_LAYOUT_12)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_LAYOUT_13.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_LAYOUT_13)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_LAYOUT_14.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_LAYOUT_14)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(TIME_LAYOUT_1.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, TIME_LAYOUT_1)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(TIME_LAYOUT_2.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, TIME_LAYOUT_2)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(TIME_LAYOUT_3.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, TIME_LAYOUT_3)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(TIME_LAYOUT_4.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, TIME_LAYOUT_4)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUT_1.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUT_1)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUT_2.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUT_2)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUT_3.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUT_3)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUT_4.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUT_4)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUT_5.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUT_5)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUT_Layout.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUT_Layout)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUTS_ANSIC.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUTS_ANSIC)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUTS_UnixDate.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUTS_UnixDate)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUTS_RubyDate.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUTS_RubyDate)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUTS_RFC822.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUTS_RFC822)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUTS_RFC822Z.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUTS_RFC822Z)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUTS_RFC850.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUTS_RFC850)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUTS_RFC1123.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUTS_RFC1123)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUTS_RFC1123Z.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUTS_RFC1123Z)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUTS_RFC3339.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUTS_RFC3339)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUTS_RFC3339Nano.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUTS_RFC3339Nano)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}
	convertion, convertionErr = time.Parse(DATE_TIME_LAYOUTS_Kitchen.Format(), stringDate)
	if convertionErr == nil {
		matchingLayouts = append(matchingLayouts, DATE_TIME_LAYOUTS_Kitchen)
		log.Println("CONVERTION SUCCESS: ", convertion.String())
	}else{
		log.Println("error parsing to objective layout format: ", convertionErr.Error())
		
	}

	log.Println("MATCHING LAYOUTS")
	log.Println(matchingLayouts)
	return
}
