package main

import "time"

type MockBookingInfo struct {
	day        time.Time
	clientType ClientType
}

func (mdi MockBookingInfo) Day() time.Time {
	return mdi.day
}

func (mdi MockBookingInfo) ClientType() ClientType {
	return mdi.clientType
}

func NewMockBookingInfo(d time.Time, ct ClientType) BookingInfo {
	return MockBookingInfo{
		day:        d,
		clientType: ct,
	}
}

type MockHotel struct {
	name   string
	fees   []Fee
	rating uint8
}

func (mh MockHotel) Name() string {
	return mh.name
}

func (mh MockHotel) Rating() uint8 {
	return mh.rating
}

func (mh MockHotel) ApplicableFee(daily BookingInfo) *Fee {
	for _, fee := range mh.fees {
		if fee.Cond()(daily) {
			return &fee
		}
	}

	return nil
}

func (mh MockHotel) ReservationPrice(checkin, checkout time.Time, ct ClientType) float64 {
	days := daysBetween(checkin, checkout)
	amount := 0.0

	for _, day := range days {
		daily := NewMockBookingInfo(day, ct)

		fee := mh.ApplicableFee(daily)

		if fee != nil {
			amount += (*fee).Value()
		}
	}

	return amount
}

type MockHotelChain struct {
	hotels []Hotel
}

func (hc MockHotelChain) LowestReservationPriceHotel(checkin, checkout time.Time, clientType ClientType) string {
	if len(hc.hotels) < 0 {
		return ""
	}

	var (
		lowestHotel            Hotel
		lowestReservationTotal float64
	)

	for i, hotel := range hc.hotels {
		reservationTotal := hotel.ReservationPrice(checkin, checkout, clientType)

		isLow := lowestReservationTotal > reservationTotal
		isHighRating := lowestReservationTotal == reservationTotal && lowestHotel.Rating() < hotel.Rating()

		if i == 0 || isLow || isHighRating {
			lowestReservationTotal = reservationTotal
			lowestHotel = hotel
		}
	}

	return lowestHotel.Name()
}
