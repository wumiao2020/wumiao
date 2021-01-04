package extend

import "time"

type Time time.Time

const TimeFormant = "2006-01-02 15:04:05"

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormant)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormant)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(TimeFormant)
}
