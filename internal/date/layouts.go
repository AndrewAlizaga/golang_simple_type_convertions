package date

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
	DATE_TIME_LAYOUT_1            DateTimeLayouts = iota
	DATE_TIME_LAYOUT_2            DateTimeLayouts = iota
	DATE_TIME_LAYOUT_3            DateTimeLayouts = iota
	DATE_TIME_LAYOUT_4            DateTimeLayouts = iota
	DATE_TIME_LAYOUT_5            DateTimeLayouts = iota
	DATE_TIME_LAYOUT_Layout       DateTimeLayouts = iota
	DATE_TIME_LAYOUTS_ANSIC       DateTimeLayouts = iota
	DATE_TIME_LAYOUTS_UnixDate    DateTimeLayouts = iota
	DATE_TIME_LAYOUTS_RubyDate    DateTimeLayouts = iota
	DATE_TIME_LAYOUTS_RFC822      DateTimeLayouts = iota
	DATE_TIME_LAYOUTS_RFC822Z     DateTimeLayouts = iota
	DATE_TIME_LAYOUTS_RFC850      DateTimeLayouts = iota
	DATE_TIME_LAYOUTS_RFC1123     DateTimeLayouts = iota
	DATE_TIME_LAYOUTS_RFC1123Z    DateTimeLayouts = iota
	DATE_TIME_LAYOUTS_RFC3339     DateTimeLayouts = iota
	DATE_TIME_LAYOUTS_RFC3339Nano DateTimeLayouts = iota
	DATE_TIME_LAYOUTS_Kitchen     DateTimeLayouts = iota

	// Handy time stamps.
	DATE_TIME_LAYOUTS_Stamp      DateTimeLayouts = iota
	DATE_TIME_LAYOUTS_StampMilli DateTimeLayouts = iota
	DATE_TIME_LAYOUTS_StampMicro DateTimeLayouts = iota
	DATE_TIME_LAYOUTS_StampNano  DateTimeLayouts = iota
)

func (e DateTimeLayouts) Format() (format string) {

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

	case DATE_TIME_LAYOUT_Layout:
		format = "01/02 03:04:05PM '06 -0700"

	case DATE_TIME_LAYOUTS_ANSIC:
		format = "Mon Jan _2 15:04:05 2006"

	case DATE_TIME_LAYOUTS_UnixDate:
		format = "Mon Jan _2 15:04:05 MST 2006"

	case DATE_TIME_LAYOUTS_RubyDate:
		format = "Mon Jan 02 15:04:05 -0700 2006"

	case DATE_TIME_LAYOUTS_RFC822:
		format = "02 Jan 06 15:04 MST"

	case DATE_TIME_LAYOUTS_RFC822Z:
		format = "02 Jan 06 15:04 -0700"

	case DATE_TIME_LAYOUTS_RFC850:
		format = "Monday, 02-Jan-06 15:04:05 MST"

	case DATE_TIME_LAYOUTS_RFC1123:
		format = "Mon, 02 Jan 2006 15:04:05 MST"

	case DATE_TIME_LAYOUTS_RFC1123Z:
		format = "Mon, 02 Jan 2006 15:04:05 -0700"

	case DATE_TIME_LAYOUTS_RFC3339:
		format = "2006-01-02T15:04:05Z07:00"

	case DATE_TIME_LAYOUTS_RFC3339Nano:
		format = "2006-01-02T15:04:05.999999999Z07:00"

	case DATE_TIME_LAYOUTS_Kitchen:
		format = "3:04PM"

	case DATE_TIME_LAYOUTS_Stamp:
		format = "Jan _2 15:04:05"

	case DATE_TIME_LAYOUTS_StampMilli:
		format = "Jan _2 15:04:05.000"

	case DATE_TIME_LAYOUTS_StampMicro:
		format = "Jan _2 15:04:05.000000"

	case DATE_TIME_LAYOUTS_StampNano:
		format = "Jan _2 15:04:05.000000000"

	default:
		format = ""
	}

	return

}
