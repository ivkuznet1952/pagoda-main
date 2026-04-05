package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/pages"
)

type Contact struct {
	mail *services.MailClient
}

func init() {
	//fmt.Println("///////// handler init 20")
	Register(new(Contact))
}

func (h *Contact) Init(c *services.Container) error {
	h.mail = c.Mail
	return nil
}

func (h *Contact) Routes(g *echo.Group) {
	//fmt.Println("///////// handler Routes 31")
	g.GET("/contact", h.Page).Name = routenames.Contact
	g.POST("/contact", h.Submit).Name = routenames.ContactSubmit
}

func (h *Contact) Page(ctx echo.Context) error {
	//fmt.Println("///////// handler Page 37")
	//var input forms.Contact
	//input.Message = "1111111"
	//var f = form.Get[forms.Contact](ctx)
	//f.Email = "11@mail.com"
	//s.Message = s.Message + "AAAAAAA"
	return pages.ContactUs(ctx, form.Get[forms.Contact](ctx))

	//return pages.ContactUs(ctx, s)
}

func (h *Contact) Submit(ctx echo.Context) error {
	var input forms.Contact
	//input.Message = input.Message + "22222"
	//fmt.Println("/////////000000 handler Submit 51:" + input.Message)
	err := form.Submit(ctx, &input)
	_ = err
	//fmt.Println("///////// CONTACT MESSAGE handler Submit 52: " + input.Message)
	//fmt.Println("///////// CONTACT EMAIL Submit 53: " + input.Email)
	//fmt.Println("///////// CONTACT TEST Submit 53: " + input.Test)
	//switch err.(type) {
	//case nil:
	//case validator.ValidationErrors:
	//	return h.Page(ctx)
	//default:
	//	return err
	//}
	//fmt.Println("/////////22222 handler Submit 51:" + input.Message)
	//err = h.mail.
	//	Compose().
	//	To(input.Email).
	//	Subject("Contact form submitted").
	//	Body(fmt.Sprintf("The message is: %s", input.Message)).
	//	Send(ctx)
	//
	//if err != nil {
	//	return fail(err, "unable to send email")
	//}

	return h.Page(ctx)
}
