package enum

type PersonType string

const (
	Natural PersonType = "NATURAL"
	Legal   PersonType = "LEGAL"
)

func (p PersonType) IsValid() bool {
	switch p {
	case Natural, Legal:
		return true
	}
	return false
}
