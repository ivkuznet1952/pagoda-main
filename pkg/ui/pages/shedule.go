package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"

	. "maragu.dev/gomponents"
	//"github.com/mikestefanello/pagoda/pkg/ui/forms"
	//"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	//. "maragu.dev/gomponents"
)

func SheduleUs(ctx echo.Context, form *forms.SheduleForm, trip *forms.SheduleParam) error {

	r := ui.NewRequest(ctx)
	r.Title = "Расписание"
	r.Metatags.Description = "Расписание"

	g := Group{

		//Iff(r.Htmx.Target != "triptest", func() Node {
		//return Card(CardParams{
		//	Title: "Card component",
		//	Body: Group{
		//		Span(Text("This is an example of a form with inline, server-side validation and HTMX-powered AJAX submissions without writing a single line of JavaScript.")),
		//		Span(Text("Only the form below will update async upon submission.")),
		//	},
		//	Color: ColorWarning,
		//	Size:  SizeMedium,
		//})
		//	return Text("SUBMIT")
		//}),
		//Iff(form.IsDone(), func() Node {
		//
		//	return Card(CardParams{
		//		Title: "Thank you!",
		//		Body: Group{
		//			Span(Text(form.Message)),
		//			Span(Text("No email was actually sent but this entire operation was handled server-side and degrades without JavaScript enabled.")),
		//		},
		//		Color: ColorSuccess,
		//		Size:  SizeLarge,
		//	})
		//}),
		//Iff(form.IsDone(), func() Node {
		//	return Badge(ColorSuccess, form.Message)
		//}),

		Iff(!form.IsDone(), func() Node {
			//fmt.Println(trip)
			//fmt.Println("/////////////////////////  00000000000000")
			return form.Render(r, trip)
		}),
		//Iff(form.IsDone(), func() Node {
		//return Badge(ColorSuccess, "Все отлично!!!!!!")
		//fmt.Println(trip)
		//form.IsSubmitted()
		//return form.Render(r, trip)
		//}),
		Iff(form.IsDone(), func() Node {
			//	msg.Success(ctx, fmt.Sprintf("Данные сохранены успешно!"))
			return form.Render(r, trip)
			//return Badge(ColorSuccess, "Все отлично!!!!!!")
		}),

		//msg.Success(ctx, fmt.Sprintf("Данные сохранены успешно!")),
		//Iff(!form.IsDone(), func() Node {
		//var input forms.TripTest
		//form.Submit(ctx, &input)
		//return form.Render(r)
		//}),
		//x.Init("alert('CLIKED'); $refs.dayDiv_4_27.click();"),

	}

	//g = append(g, Alert(ColorSuccess, "УСПЕШНО"))
	//g = append(g, msg.Success(ctx, fmt.Sprintf("Данные сохранены успешно!")))

	//)
	//Template(

	//x.Init("setTimeout(() => show_notif = false, 3000)")
	//x.Init("setTimeout(() => checked_month = 5, 3000)"),
	//x.Init("alert('CLIKED'); $refs.dayDiv_4_27.click();"),
	//Text("DDDDDDDDDD"),
	//x.On("click", "alert('CLIKED'); $refs.dayDiv_4_27.click"),
	//	),
	//x.Bind("checked_month", "4"),
	//x.Bind("selected_date", "27.04.2026"),

	//err := r.Render(layouts.Primary, Group{})
	//if err != nil {
	//	return err
	//}

	r = ui.NewRequest(ctx)
	return r.Render(layouts.Primary, g)
}
