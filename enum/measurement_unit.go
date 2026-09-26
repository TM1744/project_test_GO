package enum

type MeasurementUnit string

const (
	UnitUnit   MeasurementUnit = "UNIT"
	Meter      MeasurementUnit = "METER"
	Centimeter MeasurementUnit = "CENTIMETER"
	Millimeter MeasurementUnit = "MILLIMETER"
	Decimeter  MeasurementUnit = "DECIMETER"
	Liter      MeasurementUnit = "LITER"
	Millimiter MeasurementUnit = "MILLIMITER"
	Kilo       MeasurementUnit = "KILO"
	Gram       MeasurementUnit = "GRAM"
	Milligram  MeasurementUnit = "MILLIGRAM"
)

func (m MeasurementUnit) IsValid() bool {
	switch m {
	case UnitUnit, Meter, Centimeter, Millimeter, Decimeter, Liter, Millimiter, Kilo, Gram, Milligram:
		return true
	}
	return false
}
