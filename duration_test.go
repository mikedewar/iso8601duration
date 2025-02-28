package duration

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	t.Parallel()

	// test with bad format
	_, err := FromString("asdf")
	assert.Equal(t, err, ErrBadFormat)

	// test with month
	_, err = FromString("P1M")
	assert.Equal(t, err, ErrNoMonth)

	// test with good full string
	dur, err := FromString("P1Y2DT3H4M5S")
	assert.Nil(t, err)
	assert.Equal(t, 1, dur.Years)
	assert.Equal(t, 2, dur.Days)
	assert.Equal(t, 3, dur.Hours)
	assert.Equal(t, 4, dur.Minutes)
	assert.Equal(t, 5, dur.Seconds)

	// test with good week string
	dur, err = FromString("P1W")
	assert.Nil(t, err)
	assert.Equal(t, 1, dur.Weeks)
}

func TestString(t *testing.T) {
	t.Parallel()

	// test empty
	d := Duration{}
	assert.Equal(t, d.String(), "P")

	// test only larger-than-day
	d = Duration{Years: 1, Days: 2}
	assert.Equal(t, d.String(), "P1Y2D")

	// test only smaller-than-day
	d = Duration{Hours: 1, Minutes: 2, Seconds: 3}
	assert.Equal(t, d.String(), "PT1H2M3S")

	// test full format
	d = Duration{Years: 1, Days: 2, Hours: 3, Minutes: 4, Seconds: 5}
	assert.Equal(t, d.String(), "P1Y2DT3H4M5S")

	// test week format
	d = Duration{Weeks: 1}
	assert.Equal(t, d.String(), "P1W")
}

func TestToDuration(t *testing.T) {
	t.Parallel()

	d := Duration{Years: 1}
	assert.Equal(t, d.ToDuration(), time.Hour*24*365)

	d = Duration{Weeks: 1}
	assert.Equal(t, d.ToDuration(), time.Hour*24*7)

	d = Duration{Days: 1}
	assert.Equal(t, d.ToDuration(), time.Hour*24)

	d = Duration{Hours: 1}
	assert.Equal(t, d.ToDuration(), time.Hour)

	d = Duration{Minutes: 1}
	assert.Equal(t, d.ToDuration(), time.Minute)

	d = Duration{Seconds: 1}
	assert.Equal(t, d.ToDuration(), time.Second)
}

func TestParseISO8601Duration(t *testing.T) {
	var d time.Duration
	dd, err := ParseISO8601Duration("P")
	assert.Nil(t, err)
	assert.Equal(t, d, dd)

	// test only larger-than-day
	d, err = time.ParseDuration("8808h") // 367 days
	if err != nil {
		log.Fatal(err)
	}
	dd, err = ParseISO8601Duration("P1Y2D ")
	assert.Equal(t, d, dd)

	// test only smaller-than-day
	d, err = time.ParseDuration("1h2m3s")
	if err != nil {
		log.Fatal(err)
	}
	dd, err = ParseISO8601Duration("PT1H2M3S")
	assert.Equal(t, d, dd)

	// test full format
	d, err = time.ParseDuration("8811h4m5s")
	if err != nil {
		log.Fatal(err)
	}
	dd, err = ParseISO8601Duration("P1Y2DT3H4M5S")
	assert.Equal(t, d, dd)

	// test week format
	d, err = time.ParseDuration("168h")
	if err != nil {
		log.Fatal(err)
	}
	dd, err = ParseISO8601Duration("P1W")
	assert.Equal(t, d, dd)

}
