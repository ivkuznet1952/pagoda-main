package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	. "maragu.dev/gomponents"
)

func GOrderUs(ctx echo.Context, form *forms.GOrderForm, trip *forms.GOrderParam) error {
	r := ui.NewRequest(ctx)
	r.Title = "Contact us"
	r.Metatags.Description = "Get in touch with us."

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

			//h := &TripORM
			//append( Text("DDDDDDDDDD"))
			//form.append(test())

			return form.Render(r, trip)
		}),
		Iff(form.IsDone(), func() Node {
			//fmt.Println("/////////// 00000000000000000000")
			//return form.Render(r)
			//return Text("Дата:" + form.Day)
			return Badge(ColorSuccess, form.Day+" "+form.Begin) //+" "+strconv.Itoa(form.TouristCount))
		}),

		//Iff(!form.IsDone(), func() Node {
		//var input forms.TripTest
		//form.Submit(ctx, &input)
		//return form.Render(r)
		//}),
	}
	//g = append(g, Text("DDDDDDDDDD"))
	//err := r.Render(layouts.Primary, Group{})
	//if err != nil {
	//	return err
	//}
	r = ui.NewRequest(ctx)
	//fmt.Println("///////////////////////// pages/triptest.go TripTestUs 20")
	return r.Render(layouts.Primary, g)
}
