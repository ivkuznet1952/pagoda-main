package handlers

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/guide"
	"github.com/mikestefanello/pagoda/ent/shedule"
	"github.com/mikestefanello/pagoda/ent/transport"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/pages"
)

type (
	SheduleORM struct {
		orm *ent.Client
	}
)

func init() {
	Register(new(SheduleORM))
}

func (h *SheduleORM) Init(c *services.Container) error {
	h.orm = c.ORM
	return nil
}

func (h *SheduleORM) Routes(g *echo.Group) {

	g.POST("/shedule", h.Submit).Name = routenames.SheduleSubmit
	g.GET("/shedule", h.Page).Name = routenames.Shedule
}

func (h *SheduleORM) Page(ctx echo.Context) error {

	var f = form.Get[forms.SheduleForm](ctx)

	now := time.Now()
	//now := time.Date(2026, 3, 31, 0, 0, 0, 0, time.Local) // days of month

	var v []struct {
		Id            int
		Resource_type int
		Resource_id   int
		Begin         time.Time
		End           time.Time
		Comment       string
	}

	errShedule := h.orm.Shedule.Query().
		Where(
			shedule.BeginGT(now),
		).
		Unique(true).
		Select(shedule.FieldID, shedule.FieldResourceType, shedule.FieldResourceID, shedule.FieldBegin,
			shedule.FieldEnd, shedule.FieldComment).Scan(ctx.Request().Context(), &v)

	shedules := make([]forms.Shedule, len(v))
	if errShedule == nil {
		for i := range v {
			shedules[i] = v[i]
		}
	}

	transports := make([]components.SheduleTransport, 0)

	var trans []struct {
		Id   int
		Name string
	}

	//
	errTransport := h.orm.Transport.Query().
		Where(
			transport.Active(true),
		).
		Unique(true).
		Select(transport.FieldID, transport.FieldName).Scan(ctx.Request().Context(), &trans)

	//_ = transports
	if errTransport == nil {
		for i := range trans {
			transports = append(transports, components.SheduleTransport{Id: trans[i].Id, Name: trans[i].Name})
		}
	}

	if errTransport != nil {
		fmt.Println("ERR_TRANSPORT: " + errTransport.Error())
		//tr = nil
	}

	var gdes []struct {
		Id        int
		FirstName string
		LastName  string
	}
	_ = gdes

	errGuides := h.orm.Guide.Query().
		Where(
			guide.Active(true),
		).
		Unique(true).
		Select(guide.FieldID, guide.FieldFirstname, guide.FieldLastname).Scan(ctx.Request().Context(), &gdes)
	guides := make([]components.SheduleGuide, len(gdes))
	_ = guides
	if errGuides == nil {
		for i := range gdes {
			guides[i] = components.SheduleGuide{Id: gdes[i].Id, FirstName: gdes[i].FirstName, LastName: gdes[i].LastName}
		}
	}
	if errGuides != nil {
		fmt.Println("ERR_GUIDES: " + errGuides.Error())
		//tr = nil
	}

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

	t := forms.SheduleParam{M0: m0, M1: m1, M2: m2, Y0: y0, Y1: y1, Y2: y2, Shedules: shedules,
		Guides: guides, Transports: transports}
	//fmt.Println("t:", t)

	return pages.SheduleUs(ctx, f, &t)

}

//func GetOrderNumberByID(ctx echo.Context, client *ent.Client, orderNumberID int) (*ent.OrderNumber, error) {
//The Get method is available on the specific entity client
//orderNumber, err := client.OrderNumber.Get(ctx.Request().Context(), orderNumberID)
//if err != nil {
//	// Handle error (e.g., ent.IsNotFound(err) can check if the user wasn't found)
//	return nil, fmt.Errorf("failed to get orderNumber: %w", err)
//}
//return orderNumber, nil
//}

func (h *SheduleORM) Submit(ctx echo.Context) error {

	fmt.Println("########################### SUBMIT")

	var input forms.SheduleForm
	//fmt.Println("########### Action:" + input.SAction)
	//fmt.Println("########### Json:" + input.Json)

	err := form.Submit(ctx, &input)
	_ = err

	//fmt.Println(err)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.Page(ctx)
	default:
		return err
	}

	fmt.Println("########### Day:" + input.Day)
	fmt.Println("########### Json:" + input.Json)
	fmt.Println("########### SAction:" + input.SAction)

	//input.SAction = "update"
	//x.Init("")
	//msg.Success(ctx, fmt.Sprintf("Данные сохранены успешно!"))

	/*
		orderNumber, err := GetOrderNumberByID(ctx, h.orm, 1)
		if err != nil {
			fmt.Println("########; " + err.Error())
			return h.Page(ctx)
		}
		_ = orderNumber

		day, err := time.Parse("2006-01-02", input.Day)
		_ = day
		fmt.Println(day)

		before, _, _ := strings.Cut(input.Begin, " - ")

		begin, err := time.Parse("15:04", before)
		_ = begin
		// parse begin
		tripId, errTripId := strconv.Atoi(input.Tripid)
		if errTripId != nil {
			return h.Page(ctx)
		}
		_ = tripId
		touristCount, errTouristCount := strconv.Atoi(input.Tourists)
		if errTouristCount != nil {
			return h.Page(ctx)
		}
		_ = touristCount
		transportId, errTransportId := strconv.Atoi(input.Transport)
		if errTransportId != nil {
			return h.Page(ctx)
		}
		_ = transportId
		guideId, errGuideId := strconv.Atoi(input.Guide)
		if errGuideId != nil {
			return h.Page(ctx)
		}
		_ = guideId
		cost, errCost := strconv.Atoi(input.Cost)
		if errCost != nil {
			return h.Page(ctx)
		}
		_ = cost

		order, errOrder := h.orm.GOrder.Create().
			SetNum(orderNumber.Num).
			SetTripID(tripId).
			SetTouristCount(touristCount).
			SetDay(day).
			SetBegin(begin).
			SetTransportID(transportId).
			SetGuideID(guideId).
			SetCost(cost).
			SetStatus(0).
			SetPayStatus(0).
			SetPaidSum(0).
			SetCustomerID(1). // TODO USER.ID
			SetPlace(input.Place).
			SetComment("").
			SetCreated(time.Now()).
			SetUpdated(time.Now()).
			SetCreatedBy(1). // TODO USER.ID
			SetArchived(false).
			Save(ctx.Request().Context())
		_ = order
		if errOrder != nil {
			fmt.Println("######## ERROR ORDER; " + errOrder.Error())
			return h.Page(ctx)
		}
		errUpdateOrderNumber := h.orm.OrderNumber.UpdateOneID(1).SetNum(orderNumber.Num + 1).Exec(ctx.Request().Context())
		if errUpdateOrderNumber != nil {
			fmt.Println("######## ERROR UPDATE ORDER NUMBER " + errUpdateOrderNumber.Error())
		}
	*/
	//form.Clear(ctx)
	//msg.Success(ctx, fmt.Sprintf("The task has been created. Check the logs in %d seconds.", "input.Delay"))
	//form.Clear(ctx)
	return h.Page(ctx)
}
