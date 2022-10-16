package date

import "errors"

type DateTimeLayouts int32

const (

	//Date Layouts
	DATE_LAYOUT_1  DateTimeLayouts = iota
	DATE_LAYOUT_2  DateTimeLayouts = iota
	DATE_LAYOUT_3  DateTimeLayouts = iota
	DATE_LAYOUT_4  DateTimeLayouts = iota
	DATE_LAYOUT_5  DateTimeLayouts = iota
	DATE_LAYOUT_6  DateTimeLayouts = iota
	DATE_LAYOUT_7  DateTimeLayouts = iota
	DATE_LAYOUT_8  DateTimeLayouts = iota
	DATE_LAYOUT_9  DateTimeLayouts = iota
	DATE_LAYOUT_10 DateTimeLayouts = iota
	DATE_LAYOUT_11 DateTimeLayouts = iota
	DATE_LAYOUT_12 DateTimeLayouts = iota
	DATE_LAYOUT_13 DateTimeLayouts = iota
	DATE_LAYOUT_14 DateTimeLayouts = iota

	//Time Layouts
	TIME_LAYOUT_1 DateTimeLayouts = iota
	TIME_LAYOUT_2 DateTimeLayouts = iota
	TIME_LAYOUT_3 DateTimeLayouts = iota
	TIME_LAYOUT_4 DateTimeLayouts = iota

	//Date Time Layouts
	DATE_TIME_LAYOUT_1 DateTimeLayouts = iota
	DATE_TIME_LAYOUT_2 DateTimeLayouts = iota
	DATE_TIME_LAYOUT_3 DateTimeLayouts = iota
	DATE_TIME_LAYOUT_4 DateTimeLayouts = iota
	DATE_TIME_LAYOUT_5 DateTimeLayouts = iota
)

func (e DateTimeLayouts) Format() (format string, err error) {

	switch e {

	//Date Enums
	case DATE_LAYOUT_1:
		format = "2006-01-02"

	case DATE_LAYOUT_2:
		format = "20060102"

	case DATE_LAYOUT_3:
		format = "January 02, 2006"

	case DATE_LAYOUT_4:
		format = "02 January 2006"

	case DATE_LAYOUT_5:
		format = "02-Jan-2006"

	case DATE_LAYOUT_6:
		format = "01/02/06"

	case DATE_LAYOUT_7:
		format = "01/02/2006"

	case DATE_LAYOUT_8:
		format = "010206"

	case DATE_LAYOUT_9:
		format = "Jan-02-06"

	case DATE_LAYOUT_10:
		format = "Jan-02-2006"

	case DATE_LAYOUT_11:
		format = "06"

	case DATE_LAYOUT_12:
		format = "Mon"

	case DATE_LAYOUT_13:
		format = "Monday"

	case DATE_LAYOUT_14:
		format = "Jan-06"

	//Time Enums
	case TIME_LAYOUT_1:
		format = "15:04"

	case TIME_LAYOUT_2:
		format = "15:04:05"

	case TIME_LAYOUT_3:
		format = "3:04 PM"

	case TIME_LAYOUT_4:
		format = "03:04:05 PM"

	//Date Time Enums
	case DATE_TIME_LAYOUT_1:
		format = "2006-01-02T15:04:05"

	case DATE_TIME_LAYOUT_2:
		format = "2006-01-02T15:04:05-0700"

	case DATE_TIME_LAYOUT_3:
		format = "2 Jan 2006 15:04:05"

	case DATE_TIME_LAYOUT_4:
		format = "2 Jan 2006 15:04"

	case DATE_TIME_LAYOUT_5:
		format = "Mon, 2 Jan 2006 15:04:05 MST"

	default:
		err = errors.New("Date Layout Not Supported... Yet :( ")
	}

	return

}
