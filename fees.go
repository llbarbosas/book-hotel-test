package main

import "time"

type MockFee struct {
	cond  FeeCond
	value float64
}

func (mf MockFee) Cond() FeeCond {
	return mf.cond
}

func (mf MockFee) Value() float64 {
	return mf.value
}

func NewMockFee(cond FeeCond, value float64) Fee {
	return MockFee{
		cond:  cond,
		value: value,
	}
}

func combineFees(fees ...FeeCond) FeeCond {
	return func(hd BookingInfo) bool {
		cond := true

		for _, fee := range fees {
			cond = cond && fee(hd)
		}

		return cond
	}
}

func weekdayFeeCond(hd BookingInfo) bool {
	day := hd.Day().Weekday()

	return day != time.Sunday && day != time.Saturday
}

func weekendFeeCond(hd BookingInfo) bool {
	return !weekdayFeeCond(hd)
}

func feeCondByClientType(ct ClientType) FeeCond {
	return func(hd BookingInfo) bool {
		return string(hd.ClientType()) == string(ct)
	}
}

var (
	defaultWeekdayFeeCond = combineFees(weekdayFeeCond, feeCondByClientType(DefaultClientType))
	defaultWeekendFeeCond = combineFees(weekendFeeCond, feeCondByClientType(DefaultClientType))
	rewardWeekdayFeeCond  = combineFees(weekdayFeeCond, feeCondByClientType(RewardClientType))
	rewardWeekendFeeCond  = combineFees(weekendFeeCond, feeCondByClientType(RewardClientType))
)
