package google_get_calendar_events

import (
	"google.golang.org/api/calendar/v3"
	"time"
)

type GoogleGetCalendarEvents struct {
	Name        string
	Time        time.Duration
	Description string
}

func Get(service *calendar.Service) []GoogleGetCalendarEvents {

	var personalGoogleCalendarEvents []GoogleGetCalendarEvents

	t := time.Now()

	// start at the day => it indicates "From 00:00:00 at current day current month current year
	startOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	// end of the day => it indicates until "23:59:59 at current day current month current year
	endOfDay := startOfDay.Add(24 * time.Hour)

	events, err := service.Events.List("primary").SingleEvents(true).TimeMin(startOfDay.Format(time.RFC3339)).TimeMax(endOfDay.Format(time.RFC3339)).OrderBy("startTime").Do()

	if err != nil {
		println("🙍🏻‍♂️ No se han encontrado eventos para el día actual")
		return personalGoogleCalendarEvents
	}

	for _, event := range events.Items {
		personalGoogleCalendarEvents = append(personalGoogleCalendarEvents, GoogleGetCalendarEvents{
			Name:        event.Summary,
			Description: event.Description,
			Time: func() time.Duration {
				if event.Start.DateTime != "" && event.End.DateTime != "" {
					start, _ := time.Parse(time.RFC3339, event.Start.DateTime)
					end, _ := time.Parse(time.RFC3339, event.End.DateTime)
					return end.Sub(start)
				}
				d, _ := time.ParseDuration("8h")
				return d
			}(),
		})
	}
	return personalGoogleCalendarEvents
}
