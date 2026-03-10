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
		Name      string
		Tourists  string `form:"tourists" validate:"required"`
		Day       string `form:"day" validate:"required"`
		Begin     string `form:"begin" validate:"required"`
		Transport string `form:"transport" validate:"required"`
		form.Submission
	}

	Shedule struct {
		Resource_type int
		Begin         time.Time
		End           time.Time
	}

	Transport struct {
		Id        int
		Name      string
		Max_count int
		Min_count int
	}

	Cost struct {
		//Id        int
		Cost         int
		Transport_id int
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
		GuideCount int
		Transports []Transport
		GCosts     []Cost
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
			x.On("click", " tourist_array.forEach(c => c.active=false); item.active=true; "+
				" $nextTick(() => { tourist_count = item.value });"),
		),
	)

	calendar := ""

	for i := range 10 {
		_ = i
		calendar = calendar + "<template x-cloak x-if='tourist_count ==" + strconv.Itoa(i+1) + "&& checked_month == " +
			strconv.Itoa(trip.M0) + "'>" + convertNodeToString(Calendar(OptionsParamsCalendar{
			//Label:   strconv.Itoa(i + 1),
			Label:   "",
			Month:   trip.M0,
			Year:    trip.Y0,
			Options: initCalendarDays(trip.Shedules, i+1, trip.Trip, time.Now().Day(), trip.M0, trip.Y0, trip.GuideCount, trip.Transports),
		})) + "</template>"
		calendar = calendar + "<template x-cloak x-if='tourist_count ==" + strconv.Itoa(i+1) + "&& checked_month == " +
			strconv.Itoa(trip.M1) + "'>" + convertNodeToString(Calendar(OptionsParamsCalendar{
			Label:   "",
			Month:   trip.M1,
			Year:    trip.Y1,
			Options: initCalendarDays(trip.Shedules, i+1, trip.Trip, 0, trip.M1, trip.Y1, trip.GuideCount, trip.Transports),
		})) + "</template>"
		calendar = calendar + "<template x-cloak x-if='tourist_count ==" + strconv.Itoa(i+1) + "&& checked_month == " +
			strconv.Itoa(trip.M2) + "'>" + convertNodeToString(Calendar(OptionsParamsCalendar{
			Label:   "",
			Month:   trip.M2,
			Year:    trip.Y2,
			Options: initCalendarDays(trip.Shedules, i+1, trip.Trip, 0, trip.M2, trip.Y2, trip.GuideCount, trip.Transports),
		})) + "</template>"
	}

	return Form(

		ID("gorder"),
		Method(http.MethodPost),
		Attr("hx-post", r.Path(routenames.GOrderSubmit)),

		//g,

		//	Badge(ColorSuccess, f.Day),

		Div(
			x.Data("{ order_transport: '', order_day: '', order_begin: '', begin_list: [], transport_list: [], tourist: 0, "+
				"tourist_array: [{value:1, active: false }, {value:2, active: false }, {value:3, active: false },"+
				"{value:4, active: false }, {value:5, active: false }, {value:6, active: false },"+
				"{value:7, active: false }, {value:8, active: false }, {value:9, active: false }, "+
				"{value:10, active: false }],"+
				" tourist_count: 0, checked_month: '"+strconv.Itoa(trip.M0)+"'}"),
			P(),
			header("Наименование экскурсии"),
			P(),
			Div(
				//Class("bg-base-200 p-2"),
				Text(trip.Trip.Name),
			),

			P(),
			header("Количество туристов (включая детей 12 лет и старше)"),
			P(),
			Div(
				Class("flex flex-wrap gap-8"),
				touristCountList,
			),

			P(),

			Div(
				x.Show("tourist_count > 0"),
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
				InputFieldTourist(
					InputFieldParamsTourist{
						Name: "Tourists",
					}),
				Style("background-color: gray;"),
			),
			P(),
			Div(
				InputFieldDay(
					InputFieldParamsDay{
						Name: "Day",
					}),
				Style("background-color: green;"),
			),
			P(),
			Div(
				InputFieldBegin(
					InputFieldParamsBegin{
						Name: "Begin",
					}),
				Style("background-color: blue;"),
			),
			P(),
			Div(
				InputFieldTransport(
					InputFieldParamsTransport{
						Name: "Transport",
					}),
				Style("background-color: red; "),
			),
			P(),
		),

		ControlGroup(
			FormButton(ColorPrimary, "Оформить заказ"),
		),
		CSRF(r),
	)
}

// func initCalendarDays(node Node, touristCount int, trip ent.Trip, today int, days []ChoiceDate,
func initCalendarDays(shedules []Shedule, touristCount int, trip ent.Trip, today int,
	month int, year int, guideCount int, transports []Transport) []ChoiceDate {

	firstDayOfWeek := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC) // first day of week
	daysOfMonth := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)  // days of month

	//for i := range shedules {
	//	fmt.Println(shedules[i])
	//}

	shift0 := int(firstDayOfWeek.Weekday())
	if shift0 == 0 {
		shift0 = 7 // Sunday
	}

	//fmt.Println("#######  INIT guideCount:" + strconv.Itoa(guideCount))

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

		days[i] = ChoiceDate{Year: year, Month: month, Day: i - shift0 + 2, Label: i, IsEnabled: en,
			IsVisible: i >= shift0-1 && i < daysOfMonth.Day()+shift0-1, IsWeekend: (i+2)%7 == 0 || (i+1)%7 == 0}
	}

	trip.Begin = time.Date(year, time.Month(month), 1, trip.Begin.Hour(), trip.Begin.Minute(), 0, 0, time.UTC)
	trip.End = time.Date(year, time.Month(month), 1, trip.End.Hour(), trip.End.Minute(), 0, 0, time.UTC)
	trip.Duration = time.Date(year, time.Month(month), 1, trip.Duration.Hour(), trip.Duration.Minute(), 0, 0, time.UTC)

	tripDurations := generateTripDuration(trip)
	//fmt.Println(len(tripDurations))
	//check trip date -> is resource available
	for i := 0; i < 42; i++ {
		if days[i].IsEnabled && days[i].IsVisible {
			// TODO
			for j := range tripDurations {
				//if month == 3 && (days[i].Day == 10) {
				b1 := time.Date(year, time.Month(month), days[i].Day, tripDurations[j].begin.Hour(), tripDurations[j].begin.Minute(), 0, 0, time.UTC)
				e1 := time.Date(year, time.Month(month), days[i].Day, tripDurations[j].end.Hour(), tripDurations[j].end.Minute(), 0, 0, time.UTC)

				hasFreeGuide := hasFreeResourceForDate(shedules, ui.GUIDE_TYPE, guideCount, b1, e1)

				allowTransports := ui.Filter(transports, func(transport Transport) bool {
					return touristCount >= transport.Min_count && touristCount <= transport.Max_count
				})
				//_ = allowTransports

				arr := make([]OrderTransport, len(allowTransports))

				for k := range allowTransports {
					//_ = k
					arr[k] = OrderTransport{Id: allowTransports[k].Id, Name: allowTransports[k].Name}
				}
				//_ = arr
				days[i].OrderTransports = arr
				hasFreeTransport := hasFreeResourceForDate(shedules, ui.TRANSPORT_TYPE, len(arr), b1, e1)
				//_ = hasFreeTransport
				if hasFreeGuide && hasFreeTransport {
					hourDuration := HourDuration{H0: b1.Hour(), M0: b1.Minute(), H1: e1.Hour(), M1: e1.Minute()}
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

// check resources
func hasFreeResourceForDate(shedules []Shedule, resourceType int, resourceCount int, b time.Time, e time.Time) bool {
	// Use the generic Filter function
	resourcedShedules := ui.Filter(shedules, func(shedule Shedule) bool {
		return shedule.Resource_type == resourceType && (((shedule.Begin.Before(b) || shedule.Begin.Equal(b)) &&
			(shedule.End.After(b) || shedule.End.Equal(e)) && (shedule.End.Before(e) || shedule.End.Equal(e))) ||
			((shedule.Begin.After(b) || shedule.Begin.Equal(b)) && (shedule.End.Before(e) || shedule.End.Equal(e))) ||
			((shedule.Begin.After(b) || shedule.Begin.Equal(b)) && (shedule.Begin.Before(e) || shedule.Begin.Equal(e)) &&
				(shedule.End.After(e) || shedule.End.Equal(e))))
	})
	_ = resourcedShedules

	return len(resourcedShedules) < resourceCount
}

func generateTripDuration(tr ent.Trip) []TripDuration {

	var tripDurations []TripDuration

	durationRest := time.Hour // one hour rest between trips
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

//func getCalendar(trip *GOrderParam, touristCount int) Node {
//
//	return Calendar(OptionsParamsCalendar{
//		Label:   "111111111",
//		Month:   trip.M0,
//		Year:    trip.Y0,
//		Options: initCalendarDays(trip.Shedules, touristCount, trip.Trip, time.Now().Day(), 2, 2026, 2),
//	})
//}

//func baka(group Group) string {
//	test(myNode)
//fmt.Println("##################### baka")
//group = append(group, Text("VALUE 11111111111111111111"))
//return ""
//}
