package ntp

import (
	"github.com/beevik/ntp"
	"time"
)

// Now : Программа возвращает текущее время
func Now() (time.Time, error) {
	return ntp.Time("pool.ntp.org")
}
