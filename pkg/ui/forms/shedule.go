package forms

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	x "github.com/glsubri/gomponents-alpine"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/icons"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type (
	SheduleForm struct {
		Day       string `form:"day" validate:"required"`
		Transport string `form:"transport"`
		Guide     string `form:"guide"`
		Json      string `form:"json" validate:"required"`
		form.Submission
	}

	//SheduleForm struct {
	//	Day       string `form:"day" validate:"required"`
	//	Transport string `form:"transport" validate:"required"`
	//	Guide     string `form:"guide" validate:"required"`
	//	Json      string `form:"json" validate:"required"`
	//	form.Submission
	//}

	SheduleParam struct {
		//Trip       ent.Trip
		M0         int
		M1         int
		M2         int
		Y0         int
		Y1         int
		Y2         int
		Shedules   []Shedule
		Guides     []SheduleGuide
		Transports []SheduleTransport
	}
)

func (f *SheduleForm) Render(r *ui.Request, trip *SheduleParam) Node {

	header := func(text string) Node {
		return Div(
			Class("menu-title mt-3 uppercase bg-base-200 p-2"),
			Span(Text(text)),
		)
	}

	// type of resources
	optsType := make([]Choice, 2)
	optsType[0] = Choice{Label: "Транспорт", Value: "0"}
	optsType[1] = Choice{Label: "Гид", Value: "1"}
	//opts[2] = Choice{Label: "ORDER", Value: "333"}

	//guide
	optsGuide := make([]Choice, len(trip.Guides)+1)
	optsGuide[0] = Choice{Label: "Все", Value: "0"}
	for i := range len(trip.Guides) {
		optsGuide[i+1] = Choice{Label: trip.Guides[i].LastName + " " + trip.Guides[i].FirstName,
			Value: strconv.Itoa(trip.Guides[i].Id)}
	}

	//transport
	optsTransport := make([]Choice, len(trip.Transports)+1)
	optsTransport[0] = Choice{Label: "Все", Value: "0"}
	for i := range len(trip.Transports) {
		optsTransport[i+1] = Choice{Label: trip.Transports[i].Name,
			Value: strconv.Itoa(trip.Transports[i].Id)}
	}

	return Form(

		ID("shedule"),
		Method(http.MethodPost),
		Attr("hx-post", r.Path(routenames.SheduleSubmit)),

		H3(Text("Расписание")),
		P(),

		If(len(trip.Guides) == 0,
			Div(
				Span(
					Class("ml-8"),
					Text("Не задан список гидов!"),
				),
				Style("background-color: red;"),
			)),
		P(),
		//Transport ID == 0 => on foot!!!
		If(len(trip.Transports) == 0 || (len(trip.Transports) == 1 && trip.Transports[0].Id != 0),
			Div(
				Span(
					Class("ml-8"),
					Text("Не задан список транспорта!"),
				),
				Style("background-color: red;"),
			)),

		P(),
		Div(
			x.Cloak(),

			x.Data("{ shedule_json: 'json', selected_date: '', shedule_row: [], type_resource: '0', guide_id: '0', transport_id: '0', checked_month: '"+strconv.Itoa(trip.M0)+"'}"),
			x.Init("$watch('shedule_row', value => shedule_json=JSON.stringify(shedule_row))"),
			P(),
			//

			header("Дата"),

			P(),
			Div(
				MonthChooser(MonthChooserOptionsParams{
					Value: trip.M0,
					Options: []Choice{
						{Value: strconv.Itoa(trip.M0), Label: ui.MonthName(trip.M0) + " " + strconv.Itoa(trip.Y0)},
						{Value: strconv.Itoa(trip.M1), Label: ui.MonthName(trip.M1) + " " + strconv.Itoa(trip.Y1)},
						{Value: strconv.Itoa(trip.M2), Label: ui.MonthName(trip.M2) + " " + strconv.Itoa(trip.Y2)},
					},
				}),
			),
			//P(),
			//Raw(`<TimePicker value='2024-01-01 12:00' interval='60' />`),
			//Raw(`<div data-theme="light" className="rounded-2xl w-48" style={{ width: "390px" }}><DatePicker value="2025-01-01 12:00" /></div>`),
			//P(),

			Div(
				x.Show("checked_month == "+strconv.Itoa(trip.M0)),
				Label(),
				SheduleCalendar(OptionsParamsSheduleCalendar{
					Label: "",
					Month: trip.M0,
					Year:  trip.Y0,
					Options: initSheduleCalendarDays(trip.Shedules, time.Now().Day(), trip.M0, trip.Y0, trip.Guides,
						trip.Transports, false),
				}),
			),

			Div(
				x.Show("checked_month == "+strconv.Itoa(trip.M1)),
				Label(),
				SheduleCalendar(OptionsParamsSheduleCalendar{
					Label: "",
					Month: trip.M1,
					Year:  trip.Y1,
					Options: initSheduleCalendarDays(trip.Shedules, 0, trip.M1, trip.Y1, trip.Guides,
						trip.Transports, false),
				}),
			),

			Div(
				x.Show("checked_month == "+strconv.Itoa(trip.M2)),
				Label(),
				SheduleCalendar(OptionsParamsSheduleCalendar{
					Label: "",
					Month: trip.M2,
					Year:  trip.Y2,
					Options: initSheduleCalendarDays(trip.Shedules, 0, trip.M2, trip.Y2, trip.Guides,
						trip.Transports, true),
				}),
			),
			Div(
				Class("flex gap-2 mt-5"),
				Div(
					Text("Выбранная дата: "),
				),
				Div(
					x.Show("selected_date != ''"),
					Class("badge badge-success"),
					x.Text("selected_date"),
				),
			),

			P(Text("LIST EVENT")),
			//Template(
			//	x.For("s,index in shedule_row"),
			//	Div(
			//		Class("flex w-full gap-2 mt-5"),
			//		Input(
			//			x.Bind("value", "s.v"),
			//			Style("background-color: orange;"),
			//		),
			//		Div(
			//			icons.IconDelete(),
			//			x.On("click", "shedule_row.splice(index, 1);"), // TODO
			//x.On("click", ""+sendMessage(r.Context)),
			//),
			//),
			//),
			//P(),

			//Div(
			//	InputField(
			//
			//		InputFieldParams{
			//			Name: "WAY",
			//		}),
			//	Style("background-color: orange;"),
			//),
			//OptionsParamsSheduleCalendar struct {
			//Form    form.Form
			//Label   string
			//Year    int
			//Month   int
			//Value   int
			//Options []SheduleDate
			//}

			//calendar = calendar + "<template x-if='checked_month == " +
			//strconv.Itoa(trip.M0) + "'>" + convertNodeToString(SheduleCalendar(OptionsParamsSheduleCalendar{
			//Label:   "",
			//Month:   trip.M0,
			//Year:    trip.Y0,
			//Options: initSheduleCalendarDays(trip.Shedules, time.Now().Day(), trip.M0, trip.Y0, trip.Guides, trip.Transports, false),
			//})) + "</template>"
			//

			//P(),

			//Div(
			//	x.Show("tourist_count > 0 && order_transport != '' && order_begin != ''"),
			//header("Стоимость (Руб.)"),
			//),
			//Div(
			//	x.Show("tourist_count > 0 && order_transport != '' && order_begin != ''"),
			//	Class("ml-8 text-[#fcb700]"),
			//Class("ml-8 text-green-800"),
			//Strong(x.Text("order_cost")),
			//),
			//P(),
			//
			//If(trip.Trip.Type == 0, Div(
			//	header("Место и время подачи транспорта"),
			//)),
			//If(trip.Trip.Type == 1, Div(
			//	header("Место и время встречи гида"),
			//)),
			//Div(
			//	x.Show("tourist_count > 0 && order_transport != '' && order_begin != ''"),
			//	TextareaFieldPlace(TextareaFieldParamsPlace{
			//		Form:      f,
			//		FormField: "Place",
			//		Name:      "place",
			//		Label:     "",
			//		Value:     f.Place,
			//	}),
			//),
			//Div(
			//InputFieldTourist(
			//	InputFieldParamsTourist{
			//		Name: "Tourists",
			//	}),
			//
			InputFieldDay(
				InputFieldParamsDay{
					Name:  "Day",
					Model: "selected_date",
				}),

			//InputFieldBegin(
			//	InputFieldParamsBegin{
			//		Name: "Begin",
			//	}),
			//Div(
			InputFieldTransport(
				InputFieldParamsTransport{
					Name: "Transport",
				}),
			InputFieldTransport(
				InputFieldParamsTransport{
					Name: "Guide",
				}),

			//	Style("background-color: red;"),
			//),
			//InputFieldCost(
			//	InputFieldParamsCost{
			//		Name: "Cost",
			//	}),
			//Style("background-color: orange;"),

			//),

			Template(
				x.For("item, index in shedule_row"),
				Div(
					Class("flex flex-row flex-wrap ml-4 gap-4 bg-base-200 p-2 mt-2"),
					Div(
						Class("ml-4 mt-10"),
						Raw("&#x25A0;"),
						Style("color: orange;"),
						//x.Bind("xindex", "index"),
					),

					Div(
						Class("ml-4"),
						InputFieldTime(
							InputFieldParamsTime{
								Name:      "item.id",
								Label:     "Начало",
								InputType: "time",
								Model:     "item.begin",
							}),
					),

					Div(
						Class("ml-4"),
						InputFieldTime(
							InputFieldParamsTime{
								Name:      "item.id",
								Label:     "Окончание",
								InputType: "time",
								Model:     "item.end",
							}),
					),
					Div(
						Class("ml-4"),
						SelectListSheduleResourceType(
							OptionsParamsSheduleResourcesType{
								Label:   "Тип",
								Options: optsType,
								Model:   "item.type_resource",
								//Index: 9,
							},
						),
					),

					Template(
						x.If("item.type_resource == 1"),
						Class("ml-4"),
						SelectListSheduleResourceValue(
							OptionsParams{
								Name:    "Guide",
								Label:   "Гид",
								Options: optsGuide,
								Model:   "item.guide_id",
							},
						),
					),

					Template(
						x.If("item.type_resource == 0"),
						Class("ml-4"),
						SelectListSheduleResourceValue(
							OptionsParams{
								Name:    "Transport",
								Label:   "Транспорт",
								Options: optsTransport,
								Model:   "item.transport_id",
							},
						),
					),

					Div(
						Class("ml-4 w-96"),
						InputFieldSheduleComment(
							InputFieldParamsSheduleComment{
								Name:      "Comment",
								Label:     "Комментарий",
								InputType: "text",
								Model:     "item.comment",
							}),
					),
					Div(
						Class("mt-12 ml-8"),
						icons.IconDelete(),
						x.On("click", "shedule_row.splice(index, 1); "),
					),
				),
			),

			Div(
				InputFieldJson(
					InputFieldParamsJson{
						Name: "Json",
					}),
				Style("background-color: orange;"),
			),

			ControlGroup(
				FormButtonShedule(ColorInfo, "Сохранить"),
			),

			//Div(
			//	x.Text("guide_id"),
			//	Style("background-color: orange;"),
			//),
		),

		CSRF(r),
	)
}

func sendMessage(ctx echo.Context) string { // TODO check
	msg.Success(ctx, fmt.Sprintf("Successfully delete %s.", "DELETE"))
	return ""
}

func initSheduleCalendarDays(shedules []Shedule, today int,
	month int, year int, guides []SheduleGuide, transports []SheduleTransport, isLastMonth bool) []SheduleDate {

	firstDayOfWeek := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC) // first day of week
	daysOfMonth := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)  // days of month

	shift0 := int(firstDayOfWeek.Weekday())
	if shift0 == 0 {
		shift0 = 7 // Sunday
	}

	days := make([]SheduleDate, 42)

	for i := 0; i < 42; i++ {
		en := true
		if today > 0 && i < today+shift0-1 {
			en = false
		} // days in current month before today
		if i > (time.Now().Day()+2) && isLastMonth {
			en = false
		}
		days[i] = SheduleDate{Year: year, Month: month, Day: i - shift0 + 2, Label: i, IsEnabled: en,
			IsVisible: i >= shift0-1 && i < daysOfMonth.Day()+shift0-1, IsWeekend: (i+2)%7 == 0 || (i+1)%7 == 0}
	}

	//check shedule date -> is resource available
	for i := 0; i < 42; i++ {
		if days[i].IsVisible {

			resourcedShedules := ui.Filter(shedules, func(shedule Shedule) bool {
				return days[i].Year == shedule.Begin.Year() && days[i].Year == shedule.End.Year() &&
					days[i].Month == int(shedule.Begin.Month()) && days[i].Month == int(shedule.End.Month()) &&
					days[i].Day == shedule.Begin.Day() && days[i].Day == shedule.End.Day()
			})
			_ = resourcedShedules

			arr := make([]SheduleValue, len(resourcedShedules))
			for j := range resourcedShedules {
				arr[j] = SheduleValue(resourcedShedules[j])
			}
			days[i].Shedules = arr
		}
	}
	return days
}

//func hasFreeResourceForTransportNEW(shedules []Shedule, transports []OrderTransport) []OrderTransport {
//
//	shedulesTransport := ui.Filter(shedules, func(shedule Shedule) bool {
//		return shedule.Resource_type == ui.TRANSPORT_TYPE
//	})
//
//	arr := make([]OrderTransport, 0)
//	for k := range transports {
//		found := false
//		for j := range shedulesTransport {
//			if shedulesTransport[j].Resource_id == transports[k].Id {
//				found = true
//			}
//		}
//		if !found {
//			arr = append(arr, transports[k])
//		} else {
//			found = false
//		}
//	}
//	return arr
//}

//func hasFreeResourceForGuideNEW(shedules []Shedule, guides []OrderGuide) []OrderGuide {
//
//	shedulesGuide := ui.Filter(shedules, func(shedule Shedule) bool {
//		return shedule.Resource_type == ui.GUIDE_TYPE
//	})
//
//	arr := make([]OrderGuide, 0)
//	for k := range guides {
//		found := false
//		for j := range shedulesGuide {
//			if shedulesGuide[j].Resource_id == guides[k].Id {
//				found = true
//			}
//		}
//		if !found {
//			arr = append(arr, guides[k])
//		} else {
//			found = false
//		}
//	}
//	return arr
//}

//func generateTripDurationNEW(tr ent.Trip) []TripDuration {
//
//	var tripDurations []TripDuration
//
//	durationRest := ui.TIME_BETWEEN_TRIP * time.Hour // one hour rest between trips
//	_ = durationRest
//
//	tripDuration := calculateShedule(tr.Begin, tr.Duration)
//	condition := tripDuration.end.Before(tr.End)
//	if condition {
//		add to array
//tripDurations = append(tripDurations, tripDuration)
//}
//
//for condition {
//	b := tripDuration.end.Add(durationRest)
//	tripDuration = calculateShedule(b, tr.Duration)
//	durationToAdd := time.Hour*time.Duration(tr.Duration.Hour()) + time.Minute*time.Duration(tr.Duration.Minute()) +
//		time.Second*time.Duration(0)
//	checkEnd := tripDuration.end.Add(durationToAdd)
//	condition = checkEnd.Before(tr.End)
//	if condition {
//		tripDurations = append(tripDurations, tripDuration)
//	} else {
//		if tripDuration.end.Before(tr.End) {
//			tripDurations = append(tripDurations, tripDuration)
//		}
//	}
//}
//return tripDurations
//}
//
//func calculateSheduleNEW(begin time.Time, duration time.Time) TripDuration {
//
//	durationToAdd := time.Hour*time.Duration(duration.Hour()) + time.Minute*time.Duration(duration.Minute()) +
//		time.Second*time.Duration(0)
//	end := begin.Add(durationToAdd)
//	return TripDuration{begin: begin, end: end}
//}

/*
func convertNodeToStringNEW(myNode Node) string {

	// Use a strings.Builder to capture the rendered HTML
	var buf strings.Builder
	err := myNode.Render(&buf)
	if err != nil {
		fmt.Printf("Error rendering node: %v\n", err)
		return ""
	}

	// Get the resulting HTML string
	htmlString := buf.String()
	return htmlString
}*/

//func updateSheduleDayColor() string {
//msg.Success(ctx, fmt.Sprintf("Successfully delete %s.", "DELETE"))
//fmt.Println("#####s######## UPDATE")
//return ""
//}
