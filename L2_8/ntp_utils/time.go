package ntp_utils

import (
	"github.com/beevik/ntp"
	"time"
)

func Now() (time.Time, error) {
	return ntp.Time("pool.ntp.org") // Программа возвращает текущее время
}
