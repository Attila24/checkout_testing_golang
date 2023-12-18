package api

type AgodaBookingParams struct {
	Checkin  string
	Checkout string
	Adults   int64
	Price    int64
	Email    string
}

type bookingError struct {
	message string
}

func CreateAgodaBooking(params AgodaBookingParams) error {
	if params.Adults > 0 && params.Price > 0 {
		return nil
	}

	return bookingError{
		message: "booking failed",
	}
}

func (be bookingError) Error() string {
	return be.message
}
