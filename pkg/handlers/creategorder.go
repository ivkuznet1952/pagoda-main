package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/guide"
	"github.com/mikestefanello/pagoda/ent/shedule"
	"github.com/mikestefanello/pagoda/ent/transport"
	"github.com/mikestefanello/pagoda/ent/trip"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/pages"
)

type (
	GOrderORM struct {
		orm *ent.Client
	}

	//TripDuration struct {
	//	begin time.Time
	//	end   time.Time
	//}

	//GOrderParam struct {
	//	Trip     ent.Trip
	//	M0       int
	//	M1       int
	//	M2       int
	//	Y0       int
	//	Y1       int
	//	Y2       int
	//	Shedules []ent.Shedule
	//Days0 []components.ChoiceDate
	//Days1 []components.ChoiceDate
	//Days2 []components.ChoiceDate
	//}
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
		tr, err := h.orm.Trip.Query().Where(trip.ID(tid)).Only(ctx.Request().Context())

		var v []struct {
			Resource_type int
			Begin         time.Time
			End           time.Time
		}

		errShedule := h.orm.Shedule.Query().
			Where(
				shedule.BeginGT(now),
			).
			Unique(true).
			Select(shedule.FieldResourceType, shedule.FieldBegin, shedule.FieldEnd).Scan(ctx.Request().Context(), &v)

		shedules := make([]forms.Shedule, len(v))
		if errShedule == nil {
			for i := range v {
				shedules[i] = v[i]
			}
		}

		var t []struct {
			Id        int
			Name      string
			Max_count int
			Min_count int
		}
		_ = t

		errTransport := h.orm.Transport.Query().
			Where(
				transport.Active(true),
			).
			Unique(true).
			Select(transport.FieldID, transport.FieldName, transport.FieldMinCount, transport.FieldMaxCount).Scan(ctx.Request().Context(), &t)
		transports := make([]forms.Transport, len(t))
		_ = transports
		if errTransport == nil {
			for i := range t {
				_ = i
				transports[i] = t[i]
			}
		}
		if errTransport != nil {
			fmt.Println("ERR_TRANSPORT: " + errTransport.Error())
			tr = nil
		}

		guideCount, errGuide := h.orm.Guide.Query().Where(guide.Active(true)).Count(ctx.Request().Context())

		// create trip shedule for current date
		if err != nil {
			fmt.Println("ERR: " + err.Error())
			tr = nil
		}
		if errGuide != nil {
			fmt.Println("ERR_GUIDE: " + errGuide.Error())
			tr = nil
		}

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

			t := forms.GOrderParam{Trip: *tr, M0: m0, M1: m1, M2: m2, Y0: y0, Y1: y1, Y2: y2, Shedules: shedules, GuideCount: guideCount}
			return pages.GOrderUs(ctx, f, &t)
		}
	}

	//if tid == 0 {
	//	return pages.GOrderUs(ctx, f, nil)
	//}

	return pages.GOrderUs(ctx, f, nil)
}

func (h *GOrderORM) Submit(ctx echo.Context) error {

	var input forms.GOrderForm
	//fmt.Println(input)

	err := form.Submit(ctx, &input)
	//fmt.Println("#####////////////////////// ERROR: " + err)
	//_ = err
	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.Page(ctx)
	default:
		return err
	}

	fmt.Println("#####//////////////////////111111 handler input DATE: " + input.Day)
	fmt.Println("#####//////////////////////111111 handler input HOUR: " + input.Begin)
	//fmt.Println(input)

	//fmt.Println("##### SUBMIT")
	return h.Page(ctx)
}
