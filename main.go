package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	jardimBotanico = MockHotel{
		name:   "Jardim Botânico",
		rating: 4,
		fees: []Fee{
			NewMockFee(defaultWeekdayFeeCond, 160),
			NewMockFee(rewardWeekdayFeeCond, 110),
			NewMockFee(defaultWeekendFeeCond, 60),
			NewMockFee(rewardWeekendFeeCond, 50),
		},
	}
	parqueDasFlores = MockHotel{
		name:   "Parque das flores",
		rating: 3,
		fees: []Fee{
			NewMockFee(defaultWeekdayFeeCond, 110),
			NewMockFee(rewardWeekdayFeeCond, 80),
			NewMockFee(defaultWeekendFeeCond, 90),
			NewMockFee(rewardWeekendFeeCond, 80),
		},
	}
	marAtlantico = MockHotel{
		name:   "Mar Atlântico",
		rating: 5,
		fees: []Fee{
			NewMockFee(defaultWeekdayFeeCond, 220),
			NewMockFee(rewardWeekdayFeeCond, 100),
			NewMockFee(defaultWeekendFeeCond, 150),
			NewMockFee(rewardWeekendFeeCond, 40),
		},
	}
	RedeHoteis = MockHotelChain{
		hotels: []Hotel{jardimBotanico, parqueDasFlores, marAtlantico},
	}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	lineReader := makeLineReader(reader)

	checkInDate := parseDate(lineReader("Check in: "))
	checkOutDate := parseDate(lineReader("Check out: "))
	clientType := parseReward(lineReader("Reward client? "))

	fmt.Println(RedeHoteis.LowestReservationPriceHotel(checkInDate, checkOutDate, clientType))
}
