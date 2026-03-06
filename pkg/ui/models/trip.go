package models

import (
	"strconv"
	"time"

	"github.com/mikestefanello/pagoda/pkg/pager"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

//type (
//	Posts struct {
//		Posts []Post
//		Pager pager.Pager
//	}
//
//	Post struct {
//		ID          int
//		Title, Body string
//	}
//)

type (
	Trips struct {
		Trips []Trip
		Pager pager.Pager
	}

	Trip struct {
		ID          int
		Name        string
		Description string
		Duration    int
		Comment     string
		Begin       time.Time
		End         time.Time
		Active      bool
		Photo       string
	}
)

//func (p *Posts) Render(path string) Node {
//	g := make(Group, len(p.Posts))
//
//	for i, post := range p.Posts {
//		g[i] = post.Render()
//	}
//
//	return Div(
//		ID("posts"),
//		Ul(
//			Class("list bg-base-100 rounded-box shadow-md not-prose"),
//			g,
//		),
//		Div(Class("mb-4")),
//		Pager(p.Pager.Page, path, !p.Pager.IsEnd(), "#posts"),
//	)
//}

//func (p *Post) Render() Node {
//	return Li(
//		Class("list-row"),
//		Div(
//			Class("text-4xl font-thin opacity-30 tabular-nums"),
//			Text(fmt.Sprintf("%02d", p.ID)),
//		),
//		Div(
//			Img(
//				Class("size-10 rounded-box"),
//				Src(ui.StaticFile("gopher.png")),
//				Alt("Gopher"),
//			),
//		),
//		Div(
//			Class("list-col-grow"),
//			Div(
//				Text(p.Title),
//			),
//			Div(
//				Class("text-xs font-semibold opacity-60"),
//				Text(p.Body),
//			),
//		),
//	)
//}

func (p *Trips) Render(path string) Node {

	//fmt.Println("///////// 92:", path)
	//r.Path(routenames.Home)
	g := make(Group, len(p.Trips))

	for i, trip := range p.Trips {
		g[i] = trip.Render()
	}

	return Div(
		ID("trips"),
		Ul(
			Class("list bg-base-100 rounded-box shadow-md not-prose"),
			g,
		),
		Div(Class("mb-4")),
		Pager(p.Pager.Page, path, !p.Pager.IsEnd(), "#trips"),
	)
}

func (t *Trip) Render() Node {
	//func (t *Trip) Render(r *ui.Request) Node {
	return Li(
		Class("list-row"),
		//Div(
		//	Class("text-4xl font-thin opacity-30 tabular-nums"),
		//	Text(fmt.Sprintf("%02d", t.ID)),
		//),
		Div(
			Img(
				Class("size-10 rounded-box"),
				Src(ui.StaticFile("gopher.png")),
				Alt("Gopher"),
			),
		),
		Div(
			Class("list-col-grow"),
			Div(
				Text(t.Name),
			),
			Div(
				Class("text-xs font-semibold opacity-60"),
				Text(t.Description),
			),
		),
		ControlGroup(
			//FormButton(ColorPrimary, "Login"),
			ButtonLink(ColorWarning, "/"+routenames.CreateGOrder+"/"+strconv.Itoa(t.ID), "Заказать"),
		),
	)
}
