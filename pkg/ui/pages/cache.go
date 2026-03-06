package pages

import (
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
)

func UpdateCache(ctx echo.Context, form *forms.Cache) error {
	r := ui.NewRequest(ctx)
	r.Title = "Set a cache entry"
	//node := Div(Textf("Hello %s", value))
	//node = node
	return r.Render(layouts.Primary, form.Render(r))
}
