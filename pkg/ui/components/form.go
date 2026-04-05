package components

import (
	"fmt"
	"strconv"
	"time"

	x "github.com/glsubri/gomponents-alpine"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type (
	InputFieldParams struct {
		Form        form.Form
		FormField   string
		Name        string
		InputType   string
		Label       string
		Value       string
		Placeholder string
		Help        string
	}

	InputFieldParamsSheduleComment struct {
		//Form        form.Form
		//FormField   string
		Name      string
		InputType string
		Label     string
		Value     string
		Model     string
		//Placeholder string
		//Help        string
	}

	InputFieldParamsJson struct {
		//	Form      form.Form
		//	FormField string
		Name string
		//	InputType string
		Label string
		Model string
		//	Value     string
	}

	InputFieldParamsTime struct {
		//Form        form.Form
		//FormField   string
		Name      string
		InputType string
		Label     string
		Value     string
		Model     string
		//Index int
		//Placeholder string
		//Help        string
	}

	InputFieldParamsBegin struct {
		Name  string
		Label string
	}

	InputFieldParamsDay struct {
		Name  string
		Label string
		Model string
	}

	InputFieldParamsSheduleAction struct {
		Name  string
		Label string
		Model string
	}

	InputFieldParamsTourist struct {
		Name  string
		Label string
	}

	InputFieldParamsTransport struct {
		Name  string
		Label string
	}

	InputFieldParamsTest struct {
		Name  string
		Label string
	}

	InputFieldParamsCost struct {
		Name  string
		Label string
	}

	InputFieldParamsGuideId struct {
		Name  string
		Label string
	}

	InputFieldParamsTripId struct {
		Name string
		//Label string
		Value string
	}

	FileFieldParams struct {
		Name  string
		Label string
		Help  string
	}

	OptionsParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Value     string
		Model     string
		Options   []Choice
		Help      string
	}

	OptionsParamsSheduleResourcesType struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Value     string
		Model     string
		Options   []Choice
		Index     int
		//Help      string
	}

	//OptionsParamsResourceType struct {
	//	Name      string
	//	Label     string
	//	Value     string
	//	Model     string
	//	Options   []Choice
	//}

	MonthChooserOptionsParams struct {
		//Form      form.Form
		//FormField string
		Name    string
		Label   string
		Value   int
		Options []Choice
		//Help      string
	}

	OptionsParamsCalendar struct {
		//Form    form.Form
		Label   string
		Year    int
		Month   int
		Value   int
		Options []ChoiceDate
	}

	OptionsParamsSheduleCalendar struct {
		//Form          form.Form
		//FormField     string
		Label       string
		Year        int
		Month       int
		SelectedDay string
		//SelectedMonth int
		//Value   int
		Options []SheduleDate
	}

	OptionsParamsShedule struct {
		Label   string
		Value   int
		Options []ChoiceDateShedule
	}

	Choice struct {
		Value string
		Label string
	}

	HourDuration struct {
		H0              int
		M0              int
		H1              int
		M1              int
		OrderTransports []OrderTransport
		OrderGuides     []OrderGuide
	}

	OptionsParamsOrderPeriod struct {
		Label   string
		Name    string
		Value   string
		Options []HourDuration
	}

	OrderTransport struct {
		Id        int
		Name      string
		Cost      int
		Max_count int
		Min_count int
	}

	SheduleTransport struct {
		Id   int
		Name string
	}

	OrderGuide struct {
		Id int
	}
	SheduleGuide struct {
		Id        int
		FirstName string
		LastName  string
	}

	ChoiceDate struct {
		Year            int
		Month           int
		Day             int
		Label           int
		IsEnabled       bool
		IsVisible       bool
		IsWeekend       bool
		HourDurations   []HourDuration
		OrderTransports []OrderTransport
	}

	SheduleDate struct {
		Year      int
		Month     int
		Day       int
		Label     int
		IsEnabled bool
		IsVisible bool
		IsWeekend bool
		Shedules  []SheduleValue
	}

	SheduleValue struct {
		Id            int
		Resource_type int
		Resource_id   int
		Begin         time.Time
		End           time.Time
		Comment       string
	}

	ChoiceDateShedule struct {
		Begin time.Time
		End   time.Time
		Value string
	}

	TextareaFieldParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Value     string
		//Help      string
	}

	TextareaFieldParamsPlace struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Value     string
		//Help      string
	}

	CheckboxParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Checked   bool
	}
)

func ControlGroup(controls ...Node) Node {
	return Div(
		Class("mt-2 flex gap-2"),
		Group(controls),
	)
}

func TextareaField(el TextareaFieldParams) Node {
	return Fieldset(
		el.Label,
		Textarea(
			//Class("textarea h-24 w-2/3 "+formFieldStatusClass(el.Form, el.FormField)),
			Class("textarea h-24 w-full"),
			ID(el.Name),
			Name(el.Name),
			Text(el.Value),
		),
		//Help(el.Help),
		formFieldErrors(el.Form, el.FormField),
	)
}

func TextareaFieldPlace(el TextareaFieldParamsPlace) Node {
	return Fieldset(
		el.Label,
		Textarea(
			//Class("textarea h-24 w-2/3 "+formFieldStatusClass(el.Form, el.FormField)),
			Class("textarea h-24 w-full"),
			ID(el.Name),
			Name(el.Name),
			Text(el.Value),
			x.Model("order_place"),
		),
		//Help(el.Help),
		formFieldErrors(el.Form, el.FormField),
	)
}

/*
	func Radios(el OptionsParams) Node {
		buttons := make(Group, len(el.Options))
		for i, opt := range el.Options {
			id := "radio-" + el.Name + "-" + opt.Value

			buttons[i] = Div(
				Class("mb-2 ml-2"),
				//
				Input(
					ID(id),
					Type("radio"),
					Name(el.Name),
					Value(opt.Value),
					Class("radio mr-1 "+formFieldStatusClass(el.Form, el.FormField)),
					If(el.Value == opt.Value, Checked()),
					ds.Bind("checked_month"),
				),
				Label(
					Text(opt.Label),
					For(id),
				),
			)

		}

		return Fieldset(
			el.Label,
			Div(
				Class("flex flex-row"),
				buttons,
			),
			formFieldErrors(el.Form, el.FormField),
		)
	}
*/
func MonthChooser(el MonthChooserOptionsParams) Node {
	buttons := make(Group, len(el.Options))

	for i, opt := range el.Options {
		id := "radio-" + el.Name + "-" + opt.Value

		buttons[i] = Div(
			Class("mb-2 ml-2"),
			//
			Input(
				ID(id),
				Type("radio"),
				Name(el.Name),
				Value(opt.Value),
				Class("radio mr-1 "), //+formFieldStatusClass(el.Form, el.FormField)),
				If(strconv.Itoa(el.Value) == opt.Value, Checked()),
				x.On("click", "if (checked_month != "+opt.Value+") {order_day=''; order_begin = ''; "+
					"begin_list = []; transport_list = [];}"),
				x.Model("checked_month"),
			),
			Label(
				Text(opt.Label),
				For(id),
			),
		)
	}

	return Fieldset(
		el.Label,
		Div(
			Class("flex flex-row"),
			buttons,
		),
	)
}

func setOrderPeriod(h []HourDuration) string {

	s := "["
	for i := range h {
		value := addZeroToStr(strconv.Itoa(h[i].H0)) + ":" + addZeroToStr(strconv.Itoa(h[i].M0)) + " - " +
			addZeroToStr(strconv.Itoa(h[i].H1)) + ":" + addZeroToStr(strconv.Itoa(h[i].M1))
		if len(h[i].OrderGuides) > 0 {
			s = s + "{id: '" + strconv.Itoa(i) + "', value: '" + value + "', transports: " +
				setOrderTransport(h[i].OrderTransports) + ", guideid: " + strconv.Itoa(h[i].OrderGuides[0].Id) +
				", active: false},"
		} else {
			s = s + "{id: '" + strconv.Itoa(i) + "', value: '" + value + "', transports: " +
				setOrderTransport(h[i].OrderTransports) + ", guideid: 0" +
				", active: false},"
		}
	}
	s = s[0:len(s)-1] + "]"
	return s
}

func setOrderTransport(h []OrderTransport) string {

	//fmt.Println(len(h))
	s := "["
	for i := range h {
		value := "{ Id: " + strconv.Itoa(h[i].Id) + ", Name:'" + h[i].Name + "', Cost:" + strconv.Itoa(h[i].Cost) +
			", active: false},"
		s = s + value
	}
	s = s[0:len(s)-1] + "]"
	//fmt.Println(s)
	return s
}

//func onchange(i int) string {
//	return "console.log('########### VALUE CHANGED: order[" + strconv.Itoa(i) + "]')"

//return "order[" + strconv.Itoa(i) + "] ="
//}

//func TestSheduleRow(el OptionsParamsShedule) Node {
//
//	fields := make(Group, len(el.Options))
//
//	for i, opt := range el.Options {
//		_ = opt
//		fields[i] = Div(
//			Input(
//				Class("w-44 ml-1"),
//				x.Model("shedule_row["+strconv.Itoa(i)+"].v"),
//				Style("background-color: green;"),
//			),
//			P(),
//			P(),
//		)
//	}
//
//	return FieldsetCalendar(
//		el.Label,
//
//		Div(
//			fields,
//		),
//	)
//
//}

func Calendar(el OptionsParamsCalendar) Node {

	buttons := make(Group, len(el.Options))

	for i, opt := range el.Options {

		buttons[i] = Div(
			x.Ref("dayDiv_"+strconv.Itoa(el.Month)+"_"+strconv.Itoa(opt.Label)),
			x.On("click", updateDayColor(el, i, opt)),
			Class("flex flex-row ml-4"),
			Div(
				Raw("&#x25A0;"),
			),
			Span(
				Class("w-4 ml-1"),
				Strong(Text(strconv.Itoa(opt.Day))),
			),
			If(!opt.IsVisible, Style("color: transparent")),
			If(opt.IsVisible && opt.IsEnabled && opt.IsWeekend, Style("color: red")),
			If(opt.IsVisible && !opt.IsEnabled, Style("color: gray")),
		)

	}

	return FieldsetCalendar(
		el.Label,
		Div(
			Class("flex flex-row ml-4"),
			Span(
				Class("w-[47px]"),
				Strong(Text("пн")),
			),
			Span(
				Class("w-[47px]"),
				Strong(Text("вт")),
			),
			Span(
				Class("w-[47px]"),
				Strong(Text("ср")),
			),
			Span(
				Class("w-[47px]"),
				Strong(Text("чт")),
			),
			Span(
				Class("w-[47px]"),
				Strong(Text("пт")),
			),
			Span(
				Class("w-[47px] text-red-600"),
				Strong(Text("сб")),
			),
			Span(
				Class("w-[47px] text-red-600"),
				Strong(Text("вс")),
			),
		),
		Div(
			calendarLine(buttons),
		),
	)

}

// add 0 if string length == 1
func addZeroToStr(s string) string {
	if len(s) == 1 {
		return "0" + s
	}
	return s
}

// set day as string
func getDay(opt ChoiceDate) string {
	if opt.IsEnabled && opt.IsVisible {
		return addZeroToStr(strconv.Itoa(opt.Year)) + "-" + addZeroToStr(strconv.Itoa(opt.Month)) + "-" + strconv.Itoa(opt.Day)
	}
	return ""
}

func getDayShedule(opt SheduleDate) string {
	if opt.IsVisible {
		return addZeroToStr(strconv.Itoa(opt.Year)) + "-" + addZeroToStr(strconv.Itoa(opt.Month)) + "-" + strconv.Itoa(opt.Day)
	}
	return ""
}

func getDaySheduleFormatted(opt SheduleDate) string {
	if opt.IsVisible {
		return strconv.Itoa(opt.Day) + "." + addZeroToStr(strconv.Itoa(opt.Month)) + "." + addZeroToStr(strconv.Itoa(opt.Year))
	}
	return ""
}

// create line of calendar
func calendarLine(b Group) Node {

	var g Group
	for j := range 6 {
		g = append(g,
			Div(
				Class("flex flex-row"),
				b[j*7:(j+1)*7],
			))
	}
	return g
}

func updateDayColor(el OptionsParamsCalendar, current int, opt ChoiceDate) string {

	if !opt.IsVisible || !opt.IsEnabled {
		return ""
	}

	el.Value = current

	var s = "transport_list = []; order_cost = ''; order_place = '';  order_day='" + getDay(el.Options[current]) + "'; "

	if len(opt.HourDurations) > 0 {
		a := setOrderPeriod(opt.HourDurations)
		s = s + "begin_list=" + a + "; order_cost = ''; order_place = ''; if (begin_list.length == 1) " +
			"begin_list[0].active = true;"
	}

	l := len(el.Options)
	ref := "$refs.dayDiv_" + strconv.Itoa(el.Month)

	if !el.Options[current].IsVisible {
		return ref + "_" + strconv.Itoa(current) + ".style.color = 'transparent';"
	}

	for i := 0; i < l; i++ {

		if el.Options[i].IsVisible {
			if i == current {
				if el.Options[i].IsEnabled {
					s = s + ref + "_" + strconv.Itoa(i) + ".style.color = 'green';"
				}
				if !el.Options[i].IsEnabled {
					s = s + ref + "_" + strconv.Itoa(i) + ".style.color = 'gray';"
				}
			}
			if i != current {
				if el.Options[i].IsEnabled && !el.Options[i].IsWeekend {
					s = s + ref + "_" + strconv.Itoa(i) + ".style.color = 'white';"
				}
				if el.Options[i].IsEnabled && el.Options[i].IsWeekend {
					s = s + ref + "_" + strconv.Itoa(i) + ".style.color = 'red';"
				}
				if !el.Options[i].IsEnabled {
					s = s + ref + "_" + strconv.Itoa(i) + ".style.color = 'gray';"
				}
			}
		}
		if !el.Options[i].IsVisible {
			s = s + ref + "_" + strconv.Itoa(i) + ".style.color = 'transparent';"
		}
	}
	return s
}

// FieldsetCalendar set fields of calendar
func FieldsetCalendar(label string, els ...Node) Node {
	return FieldSet(
		Class("fieldset"),
		If(len(label) > 0, Legend(
			Class("fieldset-legend"),
			Text(label),
		)),
		Group(els),
	)
}

func SheduleCalendar(el OptionsParamsSheduleCalendar) Node {

	buttons := make(Group, len(el.Options))

	selectedDay := 0
	selectedMonth := 0
	if el.SelectedDay != "" {
		selectedDay, selectedMonth = parseDate(el.SelectedDay)
	}

	for i, opt := range el.Options {

		buttons[i] = Div(
			x.Ref("dayDiv_"+strconv.Itoa(el.Month)+"_"+strconv.Itoa(opt.Label)),
			x.On("click", updateSheduleDayColor(el, i, opt)),
			Class("flex flex-row ml-4"),
			Div(
				Raw("&#x25A0;"),
			),
			Span(
				Class("w-4 ml-1"),
				Strong(Text(strconv.Itoa(opt.Day))),
			),
			If(!opt.IsVisible, Style("color: transparent")),
			If(opt.IsVisible && opt.IsEnabled && opt.IsWeekend, Style("color: red")),
			If(opt.IsVisible && !opt.IsEnabled, Style("color: gray")),
			//If(i == 28, x.Init("$refs.dayDiv_4_28.click(); "+updateSheduleDayColor(el, i, opt))),
			If(el.Month == selectedMonth && i == selectedDay+1, x.Init("$el.click(); "+updateSheduleDayColor(el, i, opt))),
		)

	}

	return FieldsetCalendar(
		el.Label,
		Div(
			Class("flex flex-row ml-4"),
			Span(
				Class("w-[47px]"),
				Strong(Text("пн")),
			),
			Span(
				Class("w-[47px]"),
				Strong(Text("вт")),
			),
			Span(
				Class("w-[47px]"),
				Strong(Text("ср")),
			),
			Span(
				Class("w-[47px]"),
				Strong(Text("чт")),
			),
			Span(
				Class("w-[47px]"),
				Strong(Text("пт")),
			),
			Span(
				Class("w-[47px] text-red-600"),
				Strong(Text("сб")),
			),
			Span(
				Class("w-[47px] text-red-600"),
				Strong(Text("вс")),
			),
		),
		Div(
			calendarLine(buttons),
		),
	)

}

func parseDate(d string) (int, int) {
	//fmt.Println("d: ", d)
	if d == "" {
		return 0, 0
	}
	//layout := "2006-01-02"  //15:04:05
	layout := "02.01.2006"
	//value := "2024-05-15"

	//fmt.Println("PARSE")
	t, err := time.Parse(layout, d)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, 0
	}
	//fmt.Println(t)
	return t.Day(), int(t.Month())
}

//func testShedule(el OptionsParamsSheduleCalendar, current int, opt SheduleDate) Node {
//	updateSheduleDayColor(el, current, opt)
//	x.Init("$refs.dayDiv_4_27.click()")
//	return Style("color: transparent")
//}

func updateSheduleDayColor(el OptionsParamsSheduleCalendar, current int, opt SheduleDate) string {

	if !opt.IsVisible {
		return ""
	}
	if !opt.IsEnabled {
		return ""
	}
	//selected_date
	//el.Value = current
	//el.SelectedDay = current
	s := ""
	arr := ""
	for j := range opt.Shedules {
		arr = arr + "{id:'" + strconv.Itoa(opt.Shedules[j].Id) + "'," +
			"begin:'" + addZeroToStr(strconv.Itoa(opt.Shedules[j].Begin.Hour())) + ":" +
			addZeroToStr(strconv.Itoa(opt.Shedules[j].Begin.Minute())) + "'," +
			"end:'" + addZeroToStr(strconv.Itoa(opt.Shedules[j].End.Hour())) + ":" +
			addZeroToStr(strconv.Itoa(opt.Shedules[j].End.Minute())) + "'," +
			"type_resource: " + strconv.Itoa(opt.Shedules[j].Resource_type) + ","
		if opt.Shedules[j].Resource_type == 1 {
			arr = arr + "transport_id:0,"
			arr = arr + "guide_id:" + strconv.Itoa(opt.Shedules[j].Resource_id) + ","
		}
		if opt.Shedules[j].Resource_type == 0 {
			arr = arr + "transport_id:" + strconv.Itoa(opt.Shedules[j].Resource_id) + ","
			arr = arr + "guide_id:0,"
		}
		arr = arr + "comment:'" + opt.Shedules[j].Comment + "'," +
			"},"
	}
	s = " shedule_row = [" + arr + "]; selected_date = '" + getDaySheduleFormatted(opt) + "';"

	l := len(el.Options)
	ref := "$refs.dayDiv_" + strconv.Itoa(el.Month)

	if !el.Options[current].IsVisible {
		return ref + "_" + strconv.Itoa(current) + ".style.color = 'transparent';"
	}

	for i := 0; i < l; i++ {

		if el.Options[i].IsVisible {
			if i == current {
				if el.Options[i].IsEnabled {
					s = s + ref + "_" + strconv.Itoa(i) + ".style.color = 'green';"
				}
				if !el.Options[i].IsEnabled {
					s = s + ref + "_" + strconv.Itoa(i) + ".style.color = 'gray';"
				}
			}
			if i != current {
				if !el.Options[i].IsWeekend {
					s = s + ref + "_" + strconv.Itoa(i) + ".style.color = 'white';"
				}
				if el.Options[i].IsWeekend {
					s = s + ref + "_" + strconv.Itoa(i) + ".style.color = 'red';"
				}
				if !el.Options[i].IsEnabled {
					s = s + ref + "_" + strconv.Itoa(i) + ".style.color = 'gray';"
				}
			}
		}
		if !el.Options[i].IsVisible {
			s = s + ref + "_" + strconv.Itoa(i) + ".style.color = 'transparent';"
		}
	}
	return s
}

func SelectList(el OptionsParams) Node {
	buttons := make(Group, len(el.Options))

	for i, opt := range el.Options {

		buttons[i] = Option(
			Text(opt.Label),
			Value(opt.Value),
			If(opt.Value == el.Value, Attr("selected")),
		)
	}

	return Fieldset(
		el.Label,
		Select(
			Class("select "+formFieldStatusClass(el.Form, el.FormField)),
			Name(el.Name),
			buttons,
			x.Model(el.Model),
		),
		//formFieldErrors(el.Form, el.FormField),
	)
}

func SelectListSheduleResourceValue(el OptionsParams) Node {
	buttons := make(Group, len(el.Options))

	for i, opt := range el.Options {

		buttons[i] = Option(
			Text(opt.Label),
			Value(opt.Value),
			If(opt.Value == el.Value, Attr("selected")),
		)
	}

	return Fieldset(
		el.Label,
		Select(
			Class("w-96 select "+formFieldStatusClass(el.Form, el.FormField)),
			Name(el.Name),
			buttons,
			x.Model(el.Model),
		),
		//formFieldErrors(el.Form, el.FormField),
	)
}

func SelectListSheduleResourceType(el OptionsParamsSheduleResourcesType) Node {
	buttons := make(Group, len(el.Options))

	for i, opt := range el.Options {
		buttons[i] = Option(
			Text(opt.Label),
			Value(opt.Value),
			If(opt.Value == el.Value, Attr("selected")),
		)
	}

	return Fieldset(
		el.Label,
		Select(
			Class("select "+formFieldStatusClass(el.Form, el.FormField)),
			Name(el.Name),
			buttons,
			x.Model(el.Model),
			//x.On("click", "if (type_resource == 0) {shedule_row[xindex].transport_id = 0}; if (type_resource == 1) {shedule_row[xindex].guide_id = 0};"),
			//x.On("click", "console.log($el.hasAttribute('selected'))"),
		),
		//formFieldErrors(el.Form, el.FormField),
	)
}

func Checkbox(el CheckboxParams) Node {
	return Div(
		Label(
			Class("label"),
			Input(
				Class("checkbox"),
				Type("checkbox"),
				Name(el.Name),
				If(el.Checked, Checked()),
				Value("true"),
			),
			Text(" "+el.Label),
		),
		formFieldErrors(el.Form, el.FormField),
	)
}

func InputField(el InputFieldParams) Node {
	return Fieldset(
		el.Label,
		Input(
			ID(el.Name),
			Name(el.Name),
			Type(el.InputType),
			//Class("input "+formFieldStatusClass(el.Form, el.FormField)),
			Class("input"),
			Value(el.Value),
			If(el.Placeholder != "", Placeholder(el.Placeholder)),
		),
		//Help(el.Help),
		formFieldErrors(el.Form, el.FormField),
	)
}

func InputFieldSheduleComment(el InputFieldParamsSheduleComment) Node {
	return Fieldset(
		el.Label,
		Input(
			ID(el.Name),
			Name(el.Name),
			Type(el.InputType),
			//Class("input "+formFieldStatusClass(el.Form, el.FormField)),
			Class("input"),
			Value(el.Value),
			x.Model(el.Model),
			//If(el.Placeholder != "", Placeholder(el.Placeholder)),
		),
		//Help(el.Help),
		//formFieldErrors(el.Form, el.FormField),
	)
}

func InputFieldTime(el InputFieldParamsTime) Node {
	//value = el.Value
	return Fieldset(
		el.Label,
		Input(
			Name(el.Name),
			Type(el.InputType),
			Class("input w-24"),
			x.Model(el.Model),
		),
	)
}

func InputFieldTourist(el InputFieldParamsTourist) Node {
	return Fieldset(
		el.Label,
		Input(
			x.Model("tourist_count"),
			Name(el.Name),
			Class("hidden"),
		),
	)
}

func InputFieldDay(el InputFieldParamsDay) Node {
	return Fieldset(
		el.Label,
		Input(
			x.Model(el.Model),
			Name(el.Name),
			Class("hidden"),
		),
	)
}

func InputFieldSheduleAction(el InputFieldParamsSheduleAction) Node {
	return Fieldset(
		el.Label,
		Input(
			x.Model(el.Model),
			Name(el.Name),
			//Class("hidden"),
		),
	)
}

func InputFieldJson(el InputFieldParamsJson) Node {
	return Fieldset(
		el.Label,
		Input(
			//x.Model("shedule_json"),
			x.Model(el.Model),
			Name(el.Name),
			//Type("hidden"),
		),
	)
}

//func test(s string) string {
//	fmt.Println("######:" + s)
//	return ""
//}

func InputFieldBegin(el InputFieldParamsBegin) Node {
	return Fieldset(
		el.Label,
		Input(
			x.Model("order_begin"),
			Name(el.Name),
			Class("hidden"),
		),
	)
}

func InputFieldTransport(el InputFieldParamsTransport) Node {
	return Fieldset(
		el.Label,
		Input(
			x.Model("order_transport"),
			Name(el.Name),
			Class("hidden"),
		),
	)
}

func InputFieldTest(el InputFieldParamsTest) Node {
	return Fieldset(
		el.Label,
		Input(
			x.Bind("value", "JSON.stringify(shedule_row)"),
			Name(el.Name),
			//Class("hidden"),
		),
	)
}

func InputFieldCost(el InputFieldParamsCost) Node {
	return Fieldset(
		el.Label,
		Input(
			x.Model("order_cost"),
			Name(el.Name),
			Class("hidden"),
		),
	)
}

func InputFieldGuideId(el InputFieldParamsGuideId) Node {
	return Fieldset(
		el.Label,
		Input(
			x.Model("order_guideid"),
			Name(el.Name),
			Class("hidden"),
		),
	)
}

func InputFieldTripId(el InputFieldParamsTripId) Node {
	return Input(
		Name(el.Name),
		Value(el.Value),
		Class("hidden"),
	)
}

func Fieldset(label string, els ...Node) Node {
	return FieldSet(
		Class("fieldset"),
		If(len(label) > 0, Legend(
			Class("fieldset-legend"),
			Text(label),
		)),
		Group(els),
	)
}

func FileField(el FileFieldParams) Node {
	return Fieldset(
		el.Label,
		Input(
			Type("file"),
			Class("file-input"),
			Name(el.Name),
		),
		//Help(el.Help),
	)
}

func formFieldStatusClass(fm form.Form, formField string) string {
	switch {
	case fm == nil:
		return ""
	case !fm.IsSubmitted():
		return ""
	case fm.FieldHasErrors(formField):
		return "input-error"
	default:
		return "input-success"
	}

}

func formFieldErrors(fm form.Form, field string) Node {
	if fm == nil {
		return nil
	}

	errs := fm.GetFieldErrors(field)
	if len(errs) == 0 {
		return nil
	}

	g := make(Group, len(errs))
	for i, err := range errs {
		g[i] = Div(
			Class("text-error"),
			Text(err),
		)
	}

	return g
}

func CSRF(r *ui.Request) Node {
	return Input(
		Type("hidden"),
		Name("csrf"),
		Value(r.CSRF),
	)
}

func FormButton(color Color, label string) Node {
	return Button(
		Class("btn "+buttonColor(color)),
		//x.Bind("disabled", "tourist_count == 0"),
		Text(label),
	)
}

func FormButtonOrder(color Color, label string) Node {
	return Button(
		Class("btn w-full md:w-44 lg:w-44 "+buttonColor(color)),
		x.Bind("disabled", "tourist_count == 0 || order_day == '' || order_begin == '' || "+
			"order_transport == '' || order_guideid == '' || order_place == ''"),
		Text(label),
	)
}

//func FormButtonShedule(color Color, label string) Node {
//	return Button(
//		Class("btn w-full md:w-44 lg:w-44 "+buttonColor(color)),
//		Text(label),
//	)
//}

func ButtonLink(color Color, href, label string) Node {
	return A(
		Href(href),
		Class("btn "+buttonColor(color)),
		Text(label),
	)
}

func buttonColor(color Color) string {
	// Only colors being used are included so unused styles are not compiled.
	switch color {
	case ColorPrimary:
		return "btn-primary"
	case ColorInfo:
		return "btn-info"
	case ColorAccent:
		return "btn-accent"
	case ColorError:
		return "btn-error"
	case ColorLink:
		return "btn-link"
	default:
		return ""
	}
}
