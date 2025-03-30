package metrics

import "time"

const (
	TrafficOffPeak     = "off-peak"
	TrafficPeak        = "peak"
	TrafficScaleUp     = "scale-up"
	TrafficScaleDown   = "scale-down"
	TrafficProfileName = "resiliency:type/traffic/metrics/traffic"
)

type TrafficProfile struct {
	Week [7][24]string
}

func NewTrafficProfile() *TrafficProfile {
	c := new(TrafficProfile)
	return c
}

func (t *TrafficProfile) Now() string {
	ts := time.Now().UTC()
	day := ts.Weekday()
	hour := ts.Hour()
	return t.Week[day][hour]
}

/*
func (t *TrafficProfile) IsMedium(tr string) bool {
	return tr == trafficMedium
}

func (t *TrafficProfile) IsHigh(tr string) bool {
	return tr == trafficHigh
}

func (t *TrafficProfile) IsLow(tr string) bool {
	return tr == trafficLow
}

//func dayHour(t *TrafficPofile)


*/
/*
func GetCalendar(h messaging.Notifier, agentId string, msg *messaging.Message) *ProcessingCalendar {
	if !msg.IsContentType(ContentTypeCalendar) {
		return nil
	}
	if p, ok := msg.Body.(*ProcessingCalendar); ok {
		return p
	}
	h.Notify(CalendarTypeErrorStatus(agentId, msg.Body))
	return nil
}


*/
