package app

type BookingParams struct {
	Checkin  string
	Checkout string
	Adults   int64
	Price    int64
	Email    string
}

type Booking struct {
	Checkin  string
	Checkout string
	Adults   int64
	Price    int64
	Email    string
	Status   string
}

type application struct {
	Bookings []Booking
	checkout checkoutInterface
}

func NewApplication(checkout checkoutInterface) application {
	if checkout == nil {
		checkout = NewCheckout()
	}

	return application{
		checkout: checkout,
	}
}

func (app *application) CreateBooking(params BookingParams) error {
	booking, err := app.checkout.Call(params)
	if err != nil {
		return err
	}

	app.Bookings = append(app.Bookings, booking)
	return nil
}
