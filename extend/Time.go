package extend

import "time"

type Time time.Time

const timeFormant = "2006-01-02 15:04:05"

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormant)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormant)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormant)
}
