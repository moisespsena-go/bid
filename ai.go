package keyb

import "database/sql/driver"

type AI struct {
	Key
}

func (this *AI) Value() (driver.Value, error) {
	if this.Key == "" {
		this.Key = New()
	}
	return this.Key.Value()
}
