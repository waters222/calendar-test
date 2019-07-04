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

// struct for return pair
type CalendarPair struct {
	FirstId  int
	SecondId int
}

// check everything brutally
// Complexity O(n^2)
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

// Complexity O(n^2)
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

// we sort the events for faster check
func FindOverlapPairsSort(evts CalendarEvents) (ret []CalendarPair) {
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

// Segment struct

type Segment struct {
	Start int64
	End   int64
	Ids   []int
}

func (c *Segment) ToString() string {
	return fmt.Sprintf("Segment %d -> %d with ids: %v", c.Start, c.End, c.Ids)
}

func (c *Segment) IsWithin(value int64) int {
	if value < c.Start {
		return -1
	} else if value > c.End {
		return 1
	} else {
		return 0
	}
}

func (c *Segment) SplitStart(evt CalendarEvent) []Segment {
	if c.Start == evt.Start {
		return []Segment{{c.Start, c.End, append(c.Ids, evt.Id)}}
	} else {
		return []Segment{
			{c.Start, evt.Start - 1, c.Ids},
			{evt.Start, c.End, append(c.Ids, evt.Id)},
		}
	}

}

func (c *Segment) SplitWithin(evt CalendarEvent) []Segment {
	if c.Start == evt.Start && c.End == evt.End {
		return []Segment{
			{c.Start, c.End, append(c.Ids, evt.Id)},
		}
	} else if c.Start == evt.Start {
		return []Segment{
			{c.Start, evt.End, append(c.Ids, evt.Id)},
			{evt.End + 1, c.End, c.Ids},
		}
	} else if c.End == evt.End {
		return []Segment{
			{c.Start, evt.Start - 1, c.Ids},
			{evt.Start, c.End, append(c.Ids, evt.Id)},
		}
	} else {
		return []Segment{
			{c.Start, evt.Start - 1, c.Ids},
			{evt.Start, evt.End, append(c.Ids, evt.Id)},
			{evt.End + 1, c.End, c.Ids},
		}
	}

}

func (c *Segment) SplitEnd(evt CalendarEvent) []Segment {
	if c.End == evt.End {
		return []Segment{
			{c.Start, c.End, append(c.Ids, evt.Id)},
		}
	} else {
		return []Segment{
			{c.Start, evt.End, append(c.Ids, evt.Id)},
			{evt.End + 1, c.End, c.Ids},
		}
	}

}

type Segments []Segment

func (c Segments) len() int {
	return len(c)
}
func (c Segments) FindSeg(value int64) (idx int) {
	idx = len(c) / 2
	for {
		result := c[idx].IsWithin(value)
		if result < 0 {
			// search left
			idx--
			if idx < 0 {
				// underflow
				return
			}
			idx = idx / 2
		} else if result > 0 {
			// search right
			idx++
			if idx >= len(c) {
				// overflow
				return
			}
			idx = (idx + +1 + len(c)) / 2
		} else {
			//found it
			return
		}
	}
}

func FindOverlapPairsSeg(evts CalendarEvents) (ret []CalendarPair) {
	length := len(evts)

	// if the input is less then 2 events then there must be no overlaps
	if length < 2 {
		return
	}

	segs := Segments{{Start: evts[0].Start, End: evts[0].End, Ids: []int{evts[0].Id}}}
	for i := 1; i < length; i++ {
		evt := evts[i]
		// find the segment which include start
		idx := segs.FindSeg(evt.Start)
		if idx < 0 {
			// underflow
			if evt.End < segs[0].Start {
				temp := Segments{Segment{Start: evt.Start, End: evt.End, Ids: []int{evt.Id}}}
				segs = append(temp, segs...)
			} else {
				for _, id := range segs[0].Ids {
					ret = append(ret, CalendarPair{id, evt.Id})
				}
				temp := Segments{Segment{Start: evt.Start, End: segs[0].Start - 1, Ids: []int{evt.Id}}}
				temp = append(temp, segs[0].SplitEnd(evt)...)
				segs = append(temp, segs[1:]...)
			}
		} else if idx == segs.len() {
			// overflow
			segs = append(segs, Segment{Start: evt.Start, End: evt.End, Ids: []int{evt.Id}})
		} else {
			// within
			if evt.End <= segs[idx].End {
				for _, id := range segs[idx].Ids {
					ret = append(ret, CalendarPair{id, evt.Id})
				}
				// total within
				temp := append(segs[:idx], segs[idx].SplitEnd(evt)...)
				segs = append(temp, segs[idx+1:]...)
			} else {
				// we need search right for more overlaps
				for _, id := range segs[idx].Ids {
					ret = append(ret, CalendarPair{id, evt.Id})
				}
				temp := append(segs[:idx], segs[idx].SplitStart(evt)...)
				end := false
				for p := idx + 1; p < segs.len(); p++ {
					for _, id := range segs[p].Ids {
						ret = append(ret, CalendarPair{id, evt.Id})
					}
					if evt.End <= segs[p].End {
						// last one
						temp = append(temp, segs[p].SplitEnd(evt)...)
						segs = append(temp, segs[p+1:]...)
						end = true
						break
					} else {
						segs[p].Ids = append(segs[p].Ids, evt.Id)
						temp = append(temp, segs[p])
					}
				}
				if !end {
					// our new seg is over the last one
					// so append a new seg here
					segs = append(segs, Segment{Start: segs[segs.len()-1].End + 1, End: evt.End, Ids: []int{evt.Id}})
				}
			}
		}

	}
	return
}
