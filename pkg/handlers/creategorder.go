package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/gcost"
	"github.com/mikestefanello/pagoda/ent/guide"
	"github.com/mikestefanello/pagoda/ent/shedule"
	"github.com/mikestefanello/pagoda/ent/transport"
	"github.com/mikestefanello/pagoda/ent/trip"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/pages"
)

type (
	GOrderORM struct {
		orm *ent.Client
	}
)

func init() {
	Register(new(GOrderORM))
}

func (h *GOrderORM) Init(c *services.Container) error {
	h.orm = c.ORM
	return nil
}

func (h *GOrderORM) Routes(g *echo.Group) {

	g.POST("/about", h.Submit).Name = routenames.GOrderSubmit
	g.GET("/creategorder/:trip_id", h.Page).Name = routenames.CreateGOrder
}

func (h *GOrderORM) Page(ctx echo.Context) error {

	//fmt.Println("############f : 00000")
	var f = form.Get[forms.GOrderForm](ctx)

	tripId := ctx.Param("trip_id")
	//fmt.Println("############f tripId: " + tripId)
	tid, _ := strconv.Atoi(tripId)

	//fmt.Println("############f : " + strconv.Itoa(tid))

	//begin1 := time.Date(2026, 2, 27, 10, 15, 0, 0, time.UTC)
	//end1 := time.Date(2026, 2, 27, 10, 40, 0, 0, time.UTC)
	//h.orm.Shedule.UpdateOneID(7).SetBegin(begin1).SetEnd(end1).Exec(ctx.Request().Context())
	//
	//begin2 := time.Date(2026, 2, 19, 9, 35, 0, 0, time.UTC)
	//end2 := time.Date(2026, 2, 19, 9, 45, 0, 0, time.UTC)
	//h.orm.Shedule.UpdateOneID(5).SetBegin(begin2).SetEnd(end2).Exec(ctx.Request().Context())

	if tid > 0 {

		now := time.Now()
		tr, errTrip := h.orm.Trip.Query().Where(trip.ID(tid)).Only(ctx.Request().Context())

		gcosts, errGCost := h.orm.GCost.Query().Where(gcost.TripID(tid)).All(ctx.Request().Context())
		_ = gcosts

		costs := make([]forms.Cost, len(gcosts))
		if errGCost == nil {
			for i := range costs {
				costs[i] = forms.Cost{TransportId: gcosts[i].TransportID, Cost: gcosts[i].Cost}
			}
		}

		var v []struct {
			ResourceType int
			ResourceId   int
			Begin        time.Time
			End          time.Time
		}

		errShedule := h.orm.Shedule.Query().
			Where(
				shedule.BeginGT(now),
			).
			Unique(true).
			Select(shedule.FieldResourceType, shedule.FieldResourceID, shedule.FieldBegin, shedule.FieldEnd).Scan(ctx.Request().Context(), &v)

		shedules := make([]forms.Shedule, len(v))
		if errShedule == nil {
			for i := range v {
				shedules[i] = v[i]
			}
		}

		transports := make([]components.OrderTransport, 0)
		if tr != nil && tr.Type == 0 {

			var trans []struct {
				Id        int
				Name      string
				Max_count int
				Min_count int
			}
			//_ = trans
			//transports := make([]components.OrderTransport, len(trans))

			//
			errTransport := h.orm.Transport.Query().
				Where(
					transport.Active(true),
				).
				Unique(true).
				Select(transport.FieldID, transport.FieldName, transport.FieldMinCount, transport.FieldMaxCount).Scan(ctx.Request().Context(), &trans)

			//transports := make([]components.OrderTransport, len(trans))
			//_ = transports
			if errTransport == nil {
				for i := range trans {
					//transports[i] = components.OrderTransport{Id: trans[i].Id, Name: trans[i].Name, Min_count: trans[i].Min_count,
					//	Max_count: trans[i].Max_count, Cost: 0}
					transports = append(transports, components.OrderTransport{Id: trans[i].Id, Name: trans[i].Name, Min_count: trans[i].Min_count,
						Max_count: trans[i].Max_count, Cost: 0})
					filteredCost := ui.Filter(costs, func(cost forms.Cost) bool {
						return cost.TransportId == trans[i].Id
					})
					if (len(filteredCost)) == 1 {
						transports[i].Cost = filteredCost[0].Cost
					}
				}
			}

			if errTransport != nil {
				fmt.Println("ERR_TRANSPORT: " + errTransport.Error())
				tr = nil
			}
		}

		if tr != nil && tr.Type == 1 {
			transports = append(transports, components.OrderTransport{Id: 0, Name: "Пешая экскурсия", Min_count: 1,
				Max_count: 10, Cost: 0})

			filteredCost := ui.Filter(gcosts, func(cost *ent.GCost) bool {
				return cost.TransportID == 0
			})
			if (len(filteredCost)) == 1 {
				transports[0].Cost = filteredCost[0].Cost
			}
		}

		//fmt.Println(transports)
		var gdes []struct {
			Id int
		}
		_ = gdes

		errGuides := h.orm.Guide.Query().
			Where(
				guide.Active(true),
			).
			Unique(true).
			Select(guide.FieldID).Scan(ctx.Request().Context(), &gdes)
		guides := make([]components.OrderGuide, len(gdes))
		_ = guides
		if errGuides == nil {
			for i := range gdes {
				guides[i] = components.OrderGuide{Id: gdes[i].Id}
			}
		}
		if errGuides != nil {
			fmt.Println("ERR_GUIDES: " + errGuides.Error())
			tr = nil
		}

		// create trip shedule for current date
		if errTrip != nil {
			fmt.Println("ERR: " + errTrip.Error())
			tr = nil
		}
		if errGCost != nil {
			fmt.Println("ERR: " + errGCost.Error())
			tr = nil
		}

		//

		if tr != nil {
			m0 := int(now.Month())
			m1 := m0 + 1
			if m1 > 12 {
				m1 = 1
			}
			m2 := m1 + 1
			if m2 > 12 {
				m2 = 1
			}

			// year
			y0 := now.Year()
			y1 := y0
			if m1 < m0 {
				y1 = y1 + 1
			}
			y2 := y1
			if m2 < m1 {
				y2 = y2 + 1
			}

			t := forms.GOrderParam{Trip: *tr, M0: m0, M1: m1, M2: m2, Y0: y0, Y1: y1, Y2: y2, Shedules: shedules,
				Guides: guides, Transports: transports}
			return pages.GOrderUs(ctx, f, &t)
		}
	}
	return pages.GOrderUs(ctx, f, nil)
}

func (h *GOrderORM) Submit(ctx echo.Context) error {

	var input forms.GOrderForm

	err := form.Submit(ctx, &input)
	//_ = err
	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.Page(ctx)
	default:
		return err
	}

	//fmt.Println("#####//////////////////////111111 handler input DATE: " + input.Day)
	//fmt.Println("#####//////////////////////111111 handler input HOUR: " + input.Begin)
	//fmt.Println("#####//////////////////////111111 handler input TRANSPORT: " + input.Transport)
	//fmt.Println("#####//////////////////////111111 handler input TRANSPORT: " + input.Cost)
	//fmt.Println("#####//////////////////////111111 handler input PLACE: " + input.Place)
	fmt.Println("#####//////////////////////111111 handler input GUIDEID: " + input.Guide)
	//fmt.Println("##### SUBMIT") oklch(0.2326 0.014 253.100006
	return h.Page(ctx)
}
