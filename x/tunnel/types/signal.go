package types

import "fmt"

// NewSignalDeviation creates a new SignalDeviation instance.
func NewSignalDeviation(
	signalID string,
	softDeviationBPS uint64,
	hardDeviationBPS uint64,
) SignalDeviation {
	return SignalDeviation{
		SignalID:         signalID,
		SoftDeviationBPS: softDeviationBPS,
		HardDeviationBPS: hardDeviationBPS,
	}
}

// NewLatestSignalPrices creates a new LatestSignalPrices instance.
func NewLatestSignalPrices(
	tunnelID uint64,
	signalPrices []SignalPrice,
	timestamp int64,
) LatestSignalPrices {
	return LatestSignalPrices{
		TunnelID:     tunnelID,
		SignalPrices: signalPrices,
		Timestamp:    timestamp,
	}
}

// Validate validates the latest signal prices.
func (latestSignalPrices LatestSignalPrices) Validate() error {
	if latestSignalPrices.TunnelID == 0 {
		return fmt.Errorf("tunnel ID cannot be 0")
	}
	if len(latestSignalPrices.SignalPrices) == 0 {
		return fmt.Errorf("signal prices cannot be empty")
	}
	if latestSignalPrices.Timestamp < 0 {
		return fmt.Errorf("timestamp cannot be negative")
	}
	return nil
}

// UpdateSignalPrices updates the signal prices in the latest signal prices.
func (latestSignalPrices *LatestSignalPrices) UpdateSignalPrices(newSignalPrices []SignalPrice) {
	// create a map of new signal prices
	newSpMap := make(map[string]SignalPrice)
	for _, sp := range newSignalPrices {
		newSpMap[sp.SignalID] = sp
	}

	// update signal prices
	for i, sp := range latestSignalPrices.SignalPrices {
		if newSp, ok := newSpMap[sp.SignalID]; ok {
			latestSignalPrices.SignalPrices[i] = newSp
		}
	}
}

// NewSignalPrice creates a new SignalPrice instance.
func NewSignalPrice(
	signalID string,
	price uint64,
) SignalPrice {
	return SignalPrice{
		SignalID: signalID,
		Price:    price,
	}
}
