package main

import "testing"

func TestEntrada1(t *testing.T) {
	checkin := Date(16, 3, 2020)
	checkout := Date(18, 3, 2020)
	clientType := ClientType("regular")

	wanted := "Parque das flores"
	got := RedeHoteis.LowestReservationPriceHotel(checkin, checkout, clientType)

	if got != wanted {
		t.Errorf("LowestReservationPriceHotel = %s; want %s", got, wanted)
	}
}

func TestEntrada2(t *testing.T) {
	checkin := Date(20, 3, 2020)
	checkout := Date(22, 3, 2020)
	clientType := ClientType("regular")

	wanted := "Jardim Botânico"
	got := RedeHoteis.LowestReservationPriceHotel(checkin, checkout, clientType)

	if got != wanted {
		t.Errorf("LowestReservationPriceHotel = %s; want %s", got, wanted)
	}
}

func TestEntrada3(t *testing.T) {
	checkin := Date(26, 3, 2020)
	checkout := Date(28, 3, 2020)
	clientType := ClientType("fidelidade")

	wanted := "Mar Atlântico"
	got := RedeHoteis.LowestReservationPriceHotel(checkin, checkout, clientType)

	if got != wanted {
		t.Errorf("LowestReservationPriceHotel = %s; want %s", got, wanted)
	}
}
