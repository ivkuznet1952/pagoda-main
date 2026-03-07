package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/trip"
	"github.com/mikestefanello/pagoda/pkg/pager"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/ui/models"
	"github.com/mikestefanello/pagoda/pkg/ui/pages"
	"golang.org/x/net/context"
)

type Pages struct {
	orm *ent.Client
}

func init() {
	Register(new(Pages))
}

func (h *Pages) Init(c *services.Container) error {
	h.orm = c.ORM
	return nil
}

func (h *Pages) Routes(g *echo.Group) {
	g.GET("/", h.Home).Name = routenames.Home
	g.GET("/about", h.About).Name = routenames.About
}

func (h *Pages) Home(ctx echo.Context) error {
	pgr := pager.NewPager(ctx, 4)

	return pages.Home(ctx, &models.Trips{
		Trips: h.fetchTrips(&pgr),
		Pager: pgr,
	})
}

// fetchPosts is a mock example of fetching posts to illustrate how paging works.
func (h *Pages) fetchTrips(pager *pager.Pager) []models.Trip {

	//res, err := h.orm.User.
	//	Query().
	////	All(ctx.Request().Context())
	//	All(context.Background())

	r, err := h.orm.Trip.Query().Where(trip.Active(true)).All(context.Background())

	//fmt.Println(len(r))
	if err != nil {
		fmt.Printf("error: #%s\n", err)
	}

	//pager.SetItems(20)
	//trips := make([]models.Trip, 20)

	pager.SetItems(len(r))
	trips := make([]models.Trip, len(r))

	//for k := range posts {
	//	posts[k] = models.Post{
	//		ID:    k + 1,
	//		Title: fmt.Sprintf("Post example #%d", k+1),
	//		Body:  fmt.Sprintf("Lorem ipsum example #%d ddolor sit amet, consectetur adipiscing elit. Nam elementum vulputate tristique.", k+1),
	//	}
	//}

	for k := range trips {
		trips[k] = models.Trip{
			ID: r[k].ID,
			//Title: fmt.Sprintf("Post example #%d", k+1),
			//Body:  fmt.Sprintf("Lorem ipsum example #%d ddolor sit amet, consectetur adipiscing elit. Nam elementum vulputate tristique.", k+1),
			//Name: fmt.Sprintf("Post example #%d", k+1),
			Name: r[k].Name,
		}
	}
	//for l := range r {
	//	trips[l].Name = r[l].Name
	//}
	var end = min(len(r), pager.GetOffset()+pager.ItemsPerPage)
	//return trips[pager.GetOffset() : pager.GetOffset()+pager.ItemsPerPage]
	return trips[pager.GetOffset():end]
	//for l := range r {
	//	posts[l].Body = r[l].Name
	//}
	//return posts[pager.GetOffset() : pager.GetOffset()+pager.ItemsPerPage]
}

func (h *Pages) About(ctx echo.Context) error {
	return pages.About(ctx)
}
