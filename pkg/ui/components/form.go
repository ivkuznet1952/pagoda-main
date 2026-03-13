package components

import (
	"strconv"

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

	InputFieldParamsBegin struct {
		Name  string
		Label string
	}

	InputFieldParamsDay struct {
		Name  string
		Label string
	}

	InputFieldParamsTourist struct {
		Name  string
		Label string
	}

	InputFieldParamsTransport struct {
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
		Options   []Choice
		Help      string
	}

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

	OrderGuide struct {
		Id int
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
		//formFieldErrors(el.Form, el.FormField),
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
		//for j := range h[i].OrderTransports {
		//s = s + " transport_list: " + setOrderTransport(h[i].OrderTransports)
		//}
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

		Div(
			//x.Show("order_day != '' && begin_list.length > 0"),
			Class("menu-title mt-3 uppercase bg-base-200 p-2"),
			Span(Text("Период проведение экскурсии (начало/окончание)")),
		),
		Div(
			Class("flex flex-wrap gap-8 ml-8"),
			x.Show("order_day != ''"),
			Template(
				x.For("(item, index) in begin_list"),
				Div(
					Strong(
						x.Bind("style", "item.active && { color: 'green'}"),
						x.Text("item.value"),
						x.On("click", "order_transport=''; order_cost = ''; order_place = ''; "+
							"order_begin = item.value; transport_list = []; "+
							" for (let i = 0; i < begin_list.length; i++) {"+
							" begin_list[i].active = false;  "+
							"}; "+
							"; setTimeout(() => {"+
							"item.active=true;"+
							"transport_list = item.transports; if (transport_list.length == 1) {transport_list[0].active = true; order_transport =  transport_list[0].Id; order_cost = transport_list[0].Cost}"+
							" order_cost = transport_list[0].Cost; order_guideid = item.guideid; "+
							"}, 20);"),
					),
				),
			),
		),
		//for (let i = 0; i < item.transports[index].length; i++) {transport_list.push(item.transports[i]};
		Div(
			//x.Show("order_day != '' order_cost != '' && order_transport != '' && transport_list.length > 0"),
			Class("menu-title mt-3 uppercase bg-base-200 p-2"),
			Span(Text("Транспорт")),
		),
		Div(
			Class("flex flex-wrap gap-8 ml-8"),
			Template(
				x.For("item in transport_list"),
				Div(
					Strong(
						x.Bind("style", "item.active && { color: 'green'}"),
						x.Text("item.Name"),
						//x.Text("item.Id"),
						x.On("click", "if (item.Id == 0) {order_transport = '0'} else {order_transport = item.Id; } ; order_cost = item.Cost; "+
							" for (let i = 0; i < transport_list.length; i++) {"+
							" transport_list[i].active = false;  "+
							"}; "+
							"; setTimeout(() => {item.active=true; }, 20); "),
					),
				),
			),
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
		return strconv.Itoa(opt.Day) + "." + addZeroToStr(strconv.Itoa(opt.Month)) + "." +
			addZeroToStr(strconv.Itoa(opt.Year))
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

	el.Value = current

	var s = "transport_list = []; order_cost = ''; order_place = '';  order_day='" + getDay(el.Options[current]) + "'; "
	if opt.IsVisible || opt.IsEnabled {
		if len(opt.HourDurations) == 0 {
			s = s + "begin_list = []; begin_list[0]='Нет доступного времени экскурсии!'; "
		}
		//if len(opt.OrderTransports) == 0 {
		//	s = s + "transport_list = []; transport_list[0]='Нет доступного транспорта!'; "
		//}
		if len(opt.HourDurations) > 0 {
			a := setOrderPeriod(opt.HourDurations)
			s = s + "begin_list=" + a + "; order_cost = ''; order_place = ''; if (begin_list.length == 1) " +
				"begin_list[0].active = true;"
		}
		//if len(opt.OrderTransports) > 0 {
		//	a := setOrderTransport(opt.OrderTransports)
		//	s = s + "transport_list=" + a + "; "
		//}
	}
	if !opt.IsVisible || !opt.IsEnabled {
		//s = s + "begin_list=[]; transport_list=[];"
	}

	l := len(el.Options)
	ref := "$refs.dayDiv_" + strconv.Itoa(el.Month)
	// why ?
	//for i := 0; i < l; i++ {
	//	if i == current && !el.Options[i].IsVisible {
	//		return ref + "_" + strconv.Itoa(i) + ".style.color = 'transparent';"
	//	}
	//}

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

/*
func formFieldErrorsCalendar(fm form.Form, field string) Node {
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
*/
/*
func formFieldStatusClassCalendar(fm form.Form, formField string) string {

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
} */

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
		),
		//Help(el.Help),
		formFieldErrors(el.Form, el.FormField),
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
			x.Model("order_day"),
			Name(el.Name),
			Class("hidden"),
		),
	)
}

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
		Class("btn "+buttonColor(color)),
		x.Bind("disabled", "tourist_count == 0 || order_day == '' || order_begin == '' || "+
			"order_transport == '' || order_guideid == '' || order_place == ''"),
		Text(label),
	)
}

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
