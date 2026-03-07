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
		H0 int
		M0 int
		H1 int
		M1 int
	}

	OptionsParamsOrderPeriod struct {
		Label   string
		Name    string
		Value   string
		Options []HourDuration
	}

	ChoiceDate struct {
		Year      int
		Month     int
		Day       int
		Label     int
		IsEnabled bool
		IsVisible bool
		IsWeekend bool
		//IsSelected    bool
		HourDurations []HourDuration
	}

	TextareaFieldParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Value     string
		Help      string
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
			Class("textarea h-24 w-2/3 "+formFieldStatusClass(el.Form, el.FormField)),
			ID(el.Name),
			Name(el.Name),
			Text(el.Value),
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
				x.On("click", "if (checked_month != "+opt.Value+") {order_day=''};"),
				x.Model("checked_month"),
			),
			Label(
				Text(opt.Label),
				For(id),
			),
		)

	}
	//fmt.Println("####### __________________")
	return Fieldset(
		el.Label,
		Div(
			Class("flex flex-row"),
			buttons,
		),
		//formFieldErrors(el.Form, el.FormField),
	)
}

func OrderPeriod(h []HourDuration) string {

	s := "["
	for i := range h {
		value := addZeroToStr(strconv.Itoa(h[i].H0)) + ":" + addZeroToStr(strconv.Itoa(h[i].M0)) + " - " +
			addZeroToStr(strconv.Itoa(h[i].H1)) + ":" + addZeroToStr(strconv.Itoa(h[i].M1))
		s = s + "{id: '" + strconv.Itoa(i) + "', value: '" + value + "', active: false},"
	}
	s = s[0:len(s)-1] + "]"
	return s
}

func Calendar(el OptionsParamsCalendar) Node {

	buttons := make(Group, len(el.Options))

	for i, opt := range el.Options {

		buttons[i] = Div(

			Attr("x-ref", "dayDiv_"+strconv.Itoa(el.Month)+"_"+strconv.Itoa(opt.Label)),
			//Attr("x-model", "order_day"),
			Attr("@click", updateDayColor(el, i, opt)),
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
			Attr("x-show", "begin_list.length > 0"),
			Text("Период проведение экскурсии (начало/окончание)"),
		),
		Div(
			Class("flex flex-wrap gap-8"),
			Attr("x-id", "['list-item']"),
			Template(
				Attr("x-for", "(item, index) in begin_list"),
				Div(
					//Attr(":id", "$id('list-item', item.id)"),
					Strong(
						Attr(":style", "item.active && { color: 'green'}"),
						Attr("x-text", "item.value"),
						Attr("@click", "order_begin = item.value; "+
							" for (let i = 0; i < begin_list.length; i++) {"+
							" begin_list[i].active = false;  "+
							"}; "+
							"; setTimeout(() => {item.active=true;}, 20); "),
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

	var s = "order_day='" + getDay(el.Options[current]) + "'; "
	if opt.IsVisible || opt.IsEnabled {
		if len(opt.HourDurations) == 0 {
			s = s + "begin_list = []; begin_list[0]='Нет доступного времени экскурсии!'; "
		}
		if len(opt.HourDurations) > 0 {
			a := OrderPeriod(opt.HourDurations)
			s = s + "begin_list=" + a + "; "
		}
	}
	if !opt.IsVisible || !opt.IsEnabled {
		s = s + "begin_list=[]; "
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
			Class("input "+formFieldStatusClass(el.Form, el.FormField)),
			Value(el.Value),
			If(el.Placeholder != "", Placeholder(el.Placeholder)),
		),
		//Help(el.Help),
		formFieldErrors(el.Form, el.FormField),
	)
}

/*
func InputFieldTouristCount(el InputFieldParams) Node {
	return Fieldset(
		el.Label,
		Input(
			ID(el.Name),
			Name(el.Name),
			Type(el.InputType),
			Class("input "+formFieldStatusClass(el.Form, el.FormField)),
			Value(el.Value),
			If(el.Placeholder != "", Placeholder(el.Placeholder)),
			Attr("x-model", "tourist_count"),
			//Attr("x-ref", "tourist_count_field"),
			Attr("@click", "if (tourist_count < 0) tourist_count = 0;"),
		),
		//Help(el.Help),
		formFieldErrors(el.Form, el.FormField),
	)
} */

func InputFieldDay(el InputFieldParamsDay) Node {
	return Fieldset(
		el.Label,
		Input(
			//Attr("x-model", "order_day"),
			x.Model("order_day"),
			Name(el.Name),
			//Class("hidden"),
		),
	)
}

func InputFieldBegin(el InputFieldParamsBegin) Node {
	return Fieldset(
		el.Label,
		Input(
			//Attr("x-model", "order_begin"),
			x.Model("order_begin"),
			Name(el.Name),
			//Class("hidden"),
		),
	)
}

//func Help(text string) Node {
//	return If(len(text) > 0, Div(
//		Class("label"),
//		Text(text),
//	))
//}

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
