package calendar

import (
	"fmt"
	"sort"
)

// the calendar event struct
type CalendarEvent struct {
	Id    int
	Start int64 // start date in unix timestamp
	End   int64 // end date in unix timestamp
}

func (c *CalendarEvent) ToString() string {
	return fmt.Sprintf("Calendar event: %d -> %d", c.Start, c.End)
}
func (c *CalendarEvent) IsValid() bool {
	return c.Start >= 0 && c.Start <= c.End
}

func (c *CalendarEvent) isOverlap(evt CalendarEvent) bool {
	return c.Start >= evt.Start && c.Start <= evt.End ||
		c.End >= evt.Start && c.End <= evt.End ||
		evt.Start >= c.Start && evt.Start <= c.End ||
		evt.End >= c.Start && evt.End <= c.End
}

// using Golang's sort interface
// we sort the calendar events by using Start date
type CalendarEvents []CalendarEvent

func (c CalendarEvents) Len() int {
	return len(c)
}

func (c CalendarEvents) Swap(i, j int) {
	temp := c[i]
	c[i] = c[j]
	c[j] = temp
}

func (c CalendarEvents) Less(i, j int) bool {
	return c[i].Start < c[j].Start
}

// struct for return pair
type CalendarPair struct {
	FirstId  int
	SecondId int
}

// check everything brutally
func FindOverlapPairsBrutal(evts CalendarEvents) (ret []CalendarPair) {
	length := len(evts)

	// if the input is less then 2 events then there must be no overlaps
	if length < 2 {
		return
	}

	for i := 0; i < length; i++ {
		for p := i + 1; p < length; p++ {
			if evts[i].isOverlap(evts[p]) {
				ret = append(ret, CalendarPair{evts[i].Id, evts[p].Id})
			}
		}
	}

	return
}

// we sort the events for faster check
func FindOverlapPairs(evts CalendarEvents) (ret []CalendarPair) {
	length := len(evts)

	// if the input is less then 2 events then there must be no overlaps
	if length < 2 {
		return
	}

	// for better performance we can sort the event by start time
	sort.Sort(evts)

	for i := 0; i < length; i++ {
		for p := i + 1; p < length; p++ {
			// well the current event end date is before next one's start data so no need to rest events
			if evts[i].End < evts[p].Start {
				break
			}
			if evts[i].isOverlap(evts[p]) {
				// we make sure firstId always smaller
				if evts[i].Id <= evts[p].Id {
					ret = append(ret, CalendarPair{evts[i].Id, evts[p].Id})
				} else {
					ret = append(ret, CalendarPair{evts[p].Id, evts[i].Id})
				}

			}
		}
	}
	return

}
