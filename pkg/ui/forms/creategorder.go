package forms

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	x "github.com/glsubri/gomponents-alpine"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type (
	GOrderForm struct {
		Tripid    string `form:"tripid"`
		Tourists  string `form:"tourists" validate:"required"`
		Day       string `form:"day" validate:"required"`
		Begin     string `form:"begin" validate:"required"`
		Transport string `form:"transport" validate:"required"`
		Cost      string `form:"cost" validate:"required"`
		Place     string `form:"place" validate:"required"`
		Guide     string `form:"guide" validate:"required"`
		form.Submission
	}

	Shedule struct {
		Id            int
		Resource_type int
		Resource_id   int
		Begin         time.Time
		End           time.Time
		Comment       string
	}

	Cost struct {
		Cost        int
		TransportId int
	}

	GOrderParam struct {
		Trip       ent.Trip
		M0         int
		M1         int
		M2         int
		Y0         int
		Y1         int
		Y2         int
		Shedules   []Shedule
		Guides     []OrderGuide
		Transports []OrderTransport
	}

	TripDuration struct {
		begin time.Time
		end   time.Time
	}
)

func (f *GOrderForm) Render(r *ui.Request, trip *GOrderParam) Node {

	header := func(text string) Node {
		return Div(
			Class("menu-title mt-3 uppercase bg-base-200 p-2"),
			Span(Text(text)),
		)
	}

	touristCountList := Template(
		x.For("item in tourist_array"),
		Strong(
			x.Bind("style", "item.active && { color: 'green'}"),
			x.Text("item.value"),
			x.On("click", " order_day = ''; order_begin = [], order_transport = ''; transport_list = []; "+
				"order_cost = ''; order_place = ''; tourist_array.forEach(c => c.active=false); item.active=true; "+
				" $nextTick(() => { tourist_count = item.value });"),
		),
	)

	calendar := ""

	for i := range 10 {
		_ = i
		calendar = calendar + "<template x-if='tourist_count ==" + strconv.Itoa(i+1) + "&& checked_month == " +
			strconv.Itoa(trip.M0) + "'>" + convertNodeToString(Calendar(OptionsParamsCalendar{
			Label:   "",
			Month:   trip.M0,
			Year:    trip.Y0,
			Options: initCalendarDays(trip.Shedules, i+1, trip.Trip, time.Now().Day(), trip.M0, trip.Y0, trip.Guides, trip.Transports, false),
		})) + "</template>"
		calendar = calendar + "<template x-if='tourist_count ==" + strconv.Itoa(i+1) + "&& checked_month == " +
			strconv.Itoa(trip.M1) + "'>" + convertNodeToString(Calendar(OptionsParamsCalendar{
			Label:   "",
			Month:   trip.M1,
			Year:    trip.Y1,
			Options: initCalendarDays(trip.Shedules, i+1, trip.Trip, 0, trip.M1, trip.Y1, trip.Guides, trip.Transports, false),
		})) + "</template>"
		calendar = calendar + "<template x-cloak x-if='tourist_count ==" + strconv.Itoa(i+1) + "&& checked_month == " +
			strconv.Itoa(trip.M2) + "'>" + convertNodeToString(Calendar(OptionsParamsCalendar{
			Label:   "",
			Month:   trip.M2,
			Year:    trip.Y2,
			Options: initCalendarDays(trip.Shedules, i+1, trip.Trip, 0, trip.M2, trip.Y2, trip.Guides, trip.Transports, true),
		})) + "</template>"
	}

	warningArray := make([]string, 6)
	warningArray[0] = "Задайте количество туристов!"
	warningArray[1] = "Задайте день проведения экскурсии!"
	warningArray[2] = "Задайте период проведения экскурсии!"
	warningArray[3] = "Задайте транспорт!"
	warningArray[4] = "Задайте время и место подачи транспорта!"
	warningArray[5] = "Задайте удобное время и место встречи гида!"

	f.Tripid = strconv.Itoa(trip.Trip.ID)

	return Form(

		ID("gorder"),
		Method(http.MethodPost),
		Attr("hx-post", r.Path(routenames.GOrderSubmit)),

		H3(Text("Новый заказ")),
		P(),

		If(len(trip.Guides) == 0,
			Div(
				Span(
					Class("ml-8"),
					Text("Не задан список гидов!"),
				),
				Style("background-color: red;"),
			)),
		// Transport ID == 0 => on foot!!!
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
			x.Data("{ order_guideid: '', order_place: '', order_cost: '', order_transport: '', order_day: '',"+
				" order_begin: '', begin_list: [], transport_list: [], "+
				"tourist_array: [{value:1, active: false }, {value:2, active: false }, {value:3, active: false },"+
				"{value:4, active: false }, {value:5, active: false }, {value:6, active: false },"+
				"{value:7, active: false }, {value:8, active: false }, {value:9, active: false }, "+
				"{value:10, active: false }], warn: [{e: false}, {e: false},{e: false}, {e: false}, {e: false}, {e: false}],"+
				" tourist_count: 0, checked_month: '"+strconv.Itoa(trip.M0)+"'}"),

			Template(
				x.If("tourist_count == 0"),
				Badge(ColorWarning, warningArray[0]),
			),
			Template(
				x.If("tourist_count > 0 && order_day == ''"),
				Badge(ColorWarning, warningArray[1]),
			),
			Template(
				x.If("tourist_count > 0 && order_day != '' && order_begin == ''"),
				Badge(ColorWarning, warningArray[2]),
			),
			Template(
				x.If("tourist_count > 0 && order_day != '' && order_begin != '' && order_transport == ''"),
				Badge(ColorWarning, warningArray[3]),
			),

			If(trip.Trip.Type == 0, Template(
				x.If("tourist_count > 0 && order_day != '' && order_begin != '' && order_transport != '' && "+
					" order_place == ''"),
				Badge(ColorWarning, warningArray[4]),
			)),

			If(trip.Trip.Type == 1, Template(
				x.If("tourist_count > 0 && order_day != '' && order_begin != '' && order_transport != '' && "+
					" order_place == ''"),
				Badge(ColorWarning, warningArray[5]),
			)),

			header("Наименование экскурсии"),
			P(),
			Div(
				Class("ml-8"),
				Text(trip.Trip.Name),
			),

			P(),
			header("Количество туристов (включая детей 12 лет и старше)"),
			P(),
			Div(
				Class("flex flex-wrap gap-8 ml-8"),
				touristCountList,
			),

			P(),

			Div(
				header("Дата экскурсии"),
			),
			P(),
			Div(
				x.Show("tourist_count > 0"),
				MonthChooser(MonthChooserOptionsParams{
					Value: trip.M0,
					Options: []Choice{
						{Value: strconv.Itoa(trip.M0), Label: ui.MonthName(trip.M0) + " " + strconv.Itoa(trip.Y0)},
						{Value: strconv.Itoa(trip.M1), Label: ui.MonthName(trip.M1) + " " + strconv.Itoa(trip.Y1)},
						{Value: strconv.Itoa(trip.M2), Label: ui.MonthName(trip.M2) + " " + strconv.Itoa(trip.Y2)},
					},
				}),
			),

			Raw(calendar),

			Div(
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
							Class("fieldset"),
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
			//
			Div(
				Class("menu-title mt-3 uppercase bg-base-200 p-2"),
				Span(Text("Транспорт")),
			),
			Div(
				Class("flex flex-wrap gap-8 ml-8"),
				Template(
					x.For("item in transport_list"),
					Div(
						Strong(
							Class("fieldset"),
							x.Bind("style", "item.active && { color: 'green'}"),
							x.Text("item.Name"),
							x.On("click", "if (item.Id == 0) {order_transport = '0'} else {order_transport = item.Id; } ; order_cost = item.Cost; "+
								" for (let i = 0; i < transport_list.length; i++) {"+
								" transport_list[i].active = false;  "+
								"}; "+
								"; setTimeout(() => {item.active=true; }, 20); "),
						),
					),
				),
			),

			P(),

			Div(
				header("Стоимость (Руб.)"),
			),
			Div(
				x.Show("tourist_count > 0 && order_transport != '' && order_begin != ''"),
				Class("ml-8 text-green-800"),
				Strong(x.Text("order_cost")),
			),
			P(),

			If(trip.Trip.Type == 0, Div(
				header("Место и время подачи транспорта"),
			)),
			If(trip.Trip.Type == 1, Div(
				header("Место и время встречи гида"),
			)),
			Div(
				x.Show("tourist_count > 0 && order_transport != '' && order_begin != ''"),
				TextareaFieldPlace(TextareaFieldParamsPlace{
					Form:      f,
					FormField: "Place",
					Name:      "place",
					Label:     "",
					Value:     f.Place,
				}),
			),

			InputFieldTourist(
				InputFieldParamsTourist{
					Name: "Tourists",
				}),

			InputFieldDay(
				InputFieldParamsDay{
					Name:  "Day",
					Model: "order_day",
				}),

			InputFieldBegin(
				InputFieldParamsBegin{
					Name: "Begin",
				}),

			InputFieldTransport(
				InputFieldParamsTransport{
					Name: "Transport",
				}),

			InputFieldCost(
				InputFieldParamsCost{
					Name: "Cost",
				}),

			InputFieldGuideId(
				InputFieldParamsGuideId{
					Name: "Guide",
				}),

			InputFieldTripId(
				InputFieldParamsTripId{
					Name:  "Tripid",
					Value: f.Tripid,
				}),

			ControlGroup(
				FormButtonOrder(ColorInfo, "Оформить заказ"),
			),
		),

		CSRF(r),
	)
}

func initCalendarDays(shedules []Shedule, touristCount int, trip ent.Trip, today int,
	month int, year int, guides []OrderGuide, transports []OrderTransport, isLastMonth bool) []ChoiceDate {

	firstDayOfWeek := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC) // first day of week
	daysOfMonth := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)  // days of month

	shift0 := int(firstDayOfWeek.Weekday())
	if shift0 == 0 {
		shift0 = 7 // Sunday
	}

	days := make([]ChoiceDate, 42)

	if touristCount == 0 {
		return days
	}

	for i := 0; i < 42; i++ {
		//
		en := true
		if today > 0 && i < today+shift0-1 {
			en = false
		} // days in current month before today
		if i > (time.Now().Day()+2) && isLastMonth {
			en = false
		}
		days[i] = ChoiceDate{Year: year, Month: month, Day: i - shift0 + 2, Label: i, IsEnabled: en,
			IsVisible: i >= shift0-1 && i < daysOfMonth.Day()+shift0-1, IsWeekend: (i+2)%7 == 0 || (i+1)%7 == 0}
	}

	trip.Begin = time.Date(year, time.Month(month), 1, trip.Begin.Hour(), trip.Begin.Minute(), 0, 0, time.UTC)
	trip.End = time.Date(year, time.Month(month), 1, trip.End.Hour(), trip.End.Minute(), 0, 0, time.UTC)
	trip.Duration = time.Date(year, time.Month(month), 1, trip.Duration.Hour(), trip.Duration.Minute(), 0, 0, time.UTC)

	tripDurations := generateTripDuration(trip)
	//check trip date -> is resource available
	for i := 0; i < 42; i++ {
		if days[i].IsEnabled && days[i].IsVisible {
			//
			for j := range tripDurations {

				b := time.Date(year, time.Month(month), days[i].Day, tripDurations[j].begin.Hour(), tripDurations[j].begin.Minute(), 0, 0, time.UTC)
				e := time.Date(year, time.Month(month), days[i].Day, tripDurations[j].end.Hour(), tripDurations[j].end.Minute(), 0, 0, time.UTC)

				//if days[i].Day == 27 && days[i].Month == 3 {
				//	fmt.Println(shedules)
				//	fmt.Println("shedules ###################")
				//}

				resourcedShedules := ui.Filter(shedules, func(shedule Shedule) bool {
					return ((shedule.Begin.Before(b) || shedule.Begin.Equal(b)) &&
						(shedule.End.After(b) || shedule.End.Equal(e)) && (shedule.End.Before(e) || shedule.End.Equal(e))) ||
						((shedule.Begin.After(b) || shedule.Begin.Equal(b)) && (shedule.End.Before(e) || shedule.End.Equal(e))) ||
						((shedule.Begin.After(b) || shedule.Begin.Equal(b)) && (shedule.Begin.Before(e) || shedule.Begin.Equal(e)) &&
							(shedule.End.After(e) || shedule.End.Equal(e)))
				})
				_ = resourcedShedules
				//if days[i].Day == 27 && days[i].Month == 3 {
				//	fmt.Println(b)
				//	fmt.Println(e)
				//fmt.Println(resourcedShedules)
				//fmt.Println(shedules)
				//}

				allowTransports := ui.Filter(transports, func(transport OrderTransport) bool {
					return touristCount >= transport.Min_count && touristCount <= transport.Max_count
				})
				_ = allowTransports

				orderTransports := hasFreeResourceForTransport(resourcedShedules, allowTransports)
				orderGuides := hasFreeResourceForGuide(resourcedShedules, guides)

				if len(orderTransports) > 0 && len(orderGuides) > 0 {
					hourDuration := HourDuration{H0: b.Hour(), M0: b.Minute(), H1: e.Hour(), M1: e.Minute(), OrderTransports: orderTransports, OrderGuides: orderGuides}
					days[i].HourDurations = append(days[i].HourDurations, hourDuration)
				}
			}
		}
		if len(days[i].HourDurations) == 0 {
			days[i].IsEnabled = false
		}
		//}
	}
	return days
}

func hasFreeResourceForTransport(shedules []Shedule, transports []OrderTransport) []OrderTransport {

	shedulesTransport := ui.Filter(shedules, func(shedule Shedule) bool {
		return shedule.Resource_type == ui.TRANSPORT_TYPE
	})

	arr := make([]OrderTransport, 0)
	for k := range transports {
		found := false
		for j := range shedulesTransport {
			if shedulesTransport[j].Resource_id == transports[k].Id {
				found = true
			}
		}
		if !found {
			arr = append(arr, transports[k])
		} else {
			found = false
		}
	}
	return arr
}

func hasFreeResourceForGuide(shedules []Shedule, guides []OrderGuide) []OrderGuide {

	shedulesGuide := ui.Filter(shedules, func(shedule Shedule) bool {
		return shedule.Resource_type == ui.GUIDE_TYPE
	})

	arr := make([]OrderGuide, 0)
	for k := range guides {
		found := false
		for j := range shedulesGuide {
			if shedulesGuide[j].Resource_id == guides[k].Id {
				found = true
			}
		}
		if !found {
			arr = append(arr, guides[k])
		} else {
			found = false
		}
	}
	return arr
}

func generateTripDuration(tr ent.Trip) []TripDuration {

	var tripDurations []TripDuration

	durationRest := ui.TIME_BETWEEN_TRIP * time.Hour // one hour rest between trips
	_ = durationRest

	tripDuration := calculateShedule(tr.Begin, tr.Duration)
	condition := tripDuration.end.Before(tr.End)
	if condition {
		// add to array
		tripDurations = append(tripDurations, tripDuration)
	}

	for condition {
		b := tripDuration.end.Add(durationRest)
		tripDuration = calculateShedule(b, tr.Duration)
		durationToAdd := time.Hour*time.Duration(tr.Duration.Hour()) + time.Minute*time.Duration(tr.Duration.Minute()) +
			time.Second*time.Duration(0)
		checkEnd := tripDuration.end.Add(durationToAdd)
		condition = checkEnd.Before(tr.End)
		if condition {
			tripDurations = append(tripDurations, tripDuration)
		} else {
			if tripDuration.end.Before(tr.End) {
				tripDurations = append(tripDurations, tripDuration)
			}
		}
	}
	//for i := range tripDurations {
	//tripDurations[i].begin = time.Date(d.Year(), d.Month(), d.Day(), tripDurations[i].begin.Hour(), tripDurations[i].
	//	begin.Hour(), 0, 0, time.UTC)
	//tripDurations[i].end = time.Date(d.Year(), d.Month(), d.Day(), tripDurations[i].end.Hour(), tripDurations[i].end.
	//	Hour(), 0, 0, time.UTC)
	//customFormatBegin := tripDurations[i].begin.Format("3:04")
	//customFormatEnd := tripDurations[i].end.Format("3:04")
	//fmt.Println("Begin:", customFormatBegin, "End:", customFormatEnd)
	//}
	return tripDurations
}

func calculateShedule(begin time.Time, duration time.Time) TripDuration {

	durationToAdd := time.Hour*time.Duration(duration.Hour()) + time.Minute*time.Duration(duration.Minute()) +
		time.Second*time.Duration(0)
	end := begin.Add(durationToAdd)
	return TripDuration{begin: begin, end: end}
}

func convertNodeToString(myNode Node) string {

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
}
