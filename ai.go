package bid

import "database/sql/driver"

type AI struct {
	BID
}

func (this *AI) Value() (driver.Value, error) {
	if this.IsZero() {
		this.BID = New()
	}
	return this.BID.Value()
}
