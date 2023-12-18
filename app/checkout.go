package app

import "checkout_testing/app/api"

type checkoutInterface interface {
	Call(BookingParams) (Booking, error)
}

type checkout struct{}

type checkoutError struct {
	message string
}

func NewCheckout() checkoutInterface {
	return &checkout{}
}

func (c *checkout) Call(params BookingParams) (booking Booking, err error) {
	err = validate_params(params)
	if err != nil {
		return booking, err
	}

	err = createExternalBooking(params)
	if err != nil {
		return booking, checkoutError{message: "provider failure"}
	}

	return createInternalBooking(params)
}

func validate_params(params BookingParams) error {
	if params.Checkout != "" && params.Checkin != "" && params.Price != 0 && params.Adults != 0 {
		return nil
	}

	return checkoutError{message: "missing param"}
}

func createExternalBooking(params BookingParams) error {
	agodaParams := api.AgodaBookingParams{
		Checkin:  params.Checkin,
		Checkout: params.Checkout,
		Adults:   params.Adults,
		Price:    params.Price,
		Email:    params.Email,
	}

	return api.CreateAgodaBooking(agodaParams)
}

func createInternalBooking(params BookingParams) (Booking, error) {
	return Booking{
		Checkout: params.Checkout,
		Checkin:  params.Checkin,
		Price:    params.Price,
		Adults:   params.Adults,
		Status:   "active",
	}, nil
}

func (ce checkoutError) Error() string {
	return ce.message
}
