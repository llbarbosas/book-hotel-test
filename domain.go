package main

import "time"

type BookingInfo interface {
	Day() time.Time
	ClientType() ClientType
}

type FeeCond func(BookingInfo) bool

type Fee interface {
	Cond() FeeCond
	Value() float64
}

type Hotel interface {
	Name() string
	Rating() uint8
	ApplicableFee(BookingInfo) *Fee
	ReservationPrice(checkin, checkout time.Time, clientType ClientType) float64
}

type ClientType string

const (
	DefaultClientType = "regular"
	RewardClientType  = "fidelidade"
)

type HotelChain interface {
	LowestReservationPriceHotel(checkin, checkout time.Time, clientType ClientType) string
}
