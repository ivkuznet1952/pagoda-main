package forms

import (
	"net/http"
	"time"

	x "github.com/glsubri/gomponents-alpine"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/icons"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type Contact struct {
	//Email      string `form:"email" validate:"required,email"`
	Email string `form:"email" validate:"required"`
	//Department string `form:"department" validate:"required,oneof=sales marketing hr"`
	Message string `form:"message" validate:"required"`
	Test    string `form:"test"`
	form.Submission
}

func (f *Contact) Render(r *ui.Request) Node {
	f.Email = "contact@mail.com"
	arr := make([]ChoiceDateShedule, 2)
	arr[0] = ChoiceDateShedule{Begin: time.Now(), End: time.Now(), Value: "value1"}
	arr[1] = ChoiceDateShedule{Begin: time.Now(), End: time.Now(), Value: "value2"}
	_ = arr
	//arr : = ChoiceDateShedule{}

	return Form(
		ID("contact"),
		Method(http.MethodPost),
		Attr("hx-post", r.Path(routenames.ContactSubmit)),
		Div(
			x.Data("{ shedule_row: [{v:'1111'}, {v:'2222'}]} "),

			InputField(InputFieldParams{
				Form:      f,
				FormField: "Email",
				Name:      "email",
				InputType: "email",
				Label:     "Email address",
				Value:     f.Email,
			}),
			/*
				Radios(OptionsParams{
					Form:      f,
					FormField: "Department",
					Name:      "department",
					Label:     "Department",
					Value:     f.Department,
					Options: []Choice{
						{Value: "sales", Label: "Sales"},
						{Value: "marketing", Label: "Marketing"},
						{Value: "hr", Label: "HR"},
					},
				}),
			*/
			TextareaField(TextareaFieldParams{
				Form:      f,
				FormField: "Message",
				Name:      "message",
				Label:     "Message",
				Value:     f.Message,
			}),

			P(),
			//Div(
			//	x.Ref("testshedule"),
			//	TestSheduleRow(OptionsParamsShedule{
			//		Label: "TestSheduleRow",
			//		Options: arr,
			//	}),
			//),
			//Raw(`<template x-for="s in shedule_row"> <input x-bind:value="s.v"> </template>`),
			/*		Raw("<template x-for='s,index in shedule_row '> {<div>"+
					" <input x-bind:value='s.v'/> "+
					//"<svg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 32 32' stroke-width='3.5' stroke='currentColor' class='w-5 h-5'><path stroke-linecap='round' stroke-linejoin='round' d='m14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21q.512.078 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48 48 0 0 0-3.478-.397m-12 .562q.51-.088 1.022-.165m0 0a48 48 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a52 52 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a49 49 0 0 0-7.5 0'></path></svg>"+
					//"<button class='btn' x-on:click='shedule_row.splice(index, 1)'> удалить </button>"+
					//"</div>"+
					convertNodeToString(
						Div(
						//Class("flex flex-row ml-4"),
						icons.IconDelete(),
						Span(
							Class("w-4 ml-1"),
						),

						x.On("click", "shedule_row.splice(index, 1)"),
					),
						)+
					" </template>"),
			*/

			P(Text("LIST EVENT")),
			Template(
				x.For("s,index in shedule_row"),
				Div(
					Class("flex w-full gap-2 mt-5"),
					Input(
						x.Bind("value", "s.v"),
						Style("background-color: orange;"),
					),
					Div(
						icons.IconDelete(),
						x.On("click", "shedule_row.splice(index, 1)"),
					),
				),
			),
			P(),
			Div(
				InputFieldTest(
					InputFieldParamsTest{
						Name: "Test",
					}),
				Style("background-color: red;"),
			),

			P(),
			Div(
				Class("flex flex-row ml-4"),
				icons.IconRowPlus(),
				P(),

				Span(
					Class("w-4 ml-1"),
					Strong(Text("Добавить")),
				),
				x.On("click", "shedule_row.push({v:'3333'});"),
			),
			P(),
			Div(
				Class("flex flex-row ml-4"),
				//icons.IconRowMinus(),
				icons.IconDelete(),
				//Text("Удалить"),
				Span(
					Class("w-4 ml-1"),
					Strong(Text("Удалить")),
				),
				x.On("click", ""),
			),

			//Div(
			//	Span(
			//Class("i-lucide-menu w-6 h-6"),
			//Style("background-color: red;"),
			//Text("aa"),
			//),
			//Raw('<span class="i-lucide-menu w-6 h-6"></span>'),
			//),

			//<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
			//<path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
			//</svg>
			//<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
			//<path stroke-linecap="round" stroke-linejoin="round" d="M12 9v6m3-3H9m12 0a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
			//</svg>

		),
		ControlGroup(
			FormButton(ColorPrimary, "Submit"),
		),
		CSRF(r),
	)

}

//func add() string{
//	return ""
//}
