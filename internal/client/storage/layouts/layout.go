package layouts

type Layout string

const (
	LayoutDate        Layout = "01/02/2006"
	LayoutDateAndTime Layout = "01/02/2006 15:04:05"
)

func (l Layout) ToString() string {
	return string(l)
}
