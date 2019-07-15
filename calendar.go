package calendar

import (
	"fmt"
	"github.com/pkg/errors"
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
				if evts[i].Id < evts[p].Id {
					ret = append(ret, CalendarPair{evts[i].Id, evts[p].Id})
				} else {
					ret = append(ret, CalendarPair{evts[p].Id, evts[i].Id})
				}

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
		newIds := make([]int, len(c.Ids))
		copy(newIds, c.Ids)
		return []Segment{
			{c.Start, evt.Start - 1, c.Ids},
			{evt.Start, c.End, append(newIds, evt.Id)},
		}
	}

}

func (c *Segment) SplitWithin(evt CalendarEvent) []Segment {
	if c.Start == evt.Start && c.End == evt.End {
		return []Segment{
			{c.Start, c.End, append(c.Ids, evt.Id)},
		}
	} else if c.Start == evt.Start {
		newIds := make([]int, len(c.Ids))
		copy(newIds, c.Ids)
		return []Segment{
			{c.Start, evt.End, append(newIds, evt.Id)},
			{evt.End + 1, c.End, c.Ids},
		}
	} else if c.End == evt.End {
		newIds := make([]int, len(c.Ids))
		copy(newIds, c.Ids)
		return []Segment{
			{c.Start, evt.Start - 1, c.Ids},
			{evt.Start, c.End, append(newIds, evt.Id)},
		}
	} else {
		newIds0 := make([]int, len(c.Ids))
		copy(newIds0, c.Ids)
		newIds1 := make([]int, len(c.Ids))
		copy(newIds1, c.Ids)
		return []Segment{
			{c.Start, evt.Start - 1, newIds0},
			{evt.Start, evt.End, append(newIds1, evt.Id)},
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
		newIds := make([]int, len(c.Ids))
		copy(newIds, c.Ids)
		return []Segment{
			{c.Start, evt.End, append(newIds, evt.Id)},
			{evt.End + 1, c.End, c.Ids},
		}
	}

}

type Segments []Segment

func (c Segments) len() int {
	return len(c)
}

var errorTooSmall = errors.New("too small")
var errorTooLarge = errors.New("too large")
var errorInBetween = errors.New("in between")

func (c Segments) FindSeg(value int64) (idx int, err error) {
	left := 0
	right := len(c)
	for left < right {
		idx = (left + right) / 2
		result := c[idx].IsWithin(value)
		if result < 0 {
			// search left
			right = idx
		} else if result > 0 {
			// search right
			left = idx + 1
		} else {
			return
		}
	}
	if right == 0 {
		err = errorTooSmall
	} else if left == len(c) {
		err = errorTooLarge
	} else {
		err = errorInBetween

	}
	return
}

func (c Segments) InsertSegments(idx int, segs []Segment) Segments {
	segs = append(segs, c[idx:]...)
	return append(c[:idx], segs...)

}

func (c Segments) AddSeg(idx int, id int, start, end int64, pairs map[CalendarPair]bool) Segments {
	var insertions []Segment

	finished := false
	leftSegs := c[:idx]
	for idx < len(c) {
		seg := c[idx]
		if end < seg.Start {
			insertions = append(insertions, Segment{start, end, []int{id}})
			finished = true
			break
		} else {
			// adding the pair into hash map
			for _, v := range seg.Ids {
				pair := CalendarPair{FirstId: v, SecondId: id}
				if _, ok := pairs[pair]; !ok {
					pairs[pair] = true
				}
			}
			if start < seg.Start {
				insertions = append(insertions, Segment{start, seg.Start - 1, []int{id}})
			}

			if end <= seg.End {
				insertions = append(insertions, seg.SplitEnd(CalendarEvent{id, start, end})...)
				finished = true
				idx++
				break
			} else {
				seg.Ids = append(seg.Ids, id)
				insertions = append(insertions, seg)
				start = seg.End + 1
				idx++
			}

		}
	}

	if finished {
		if idx < len(c) {
			insertions = append(insertions, c[idx:]...)
		}
	} else {
		insertions = append(insertions, Segment{start, end, []int{id}})
	}

	return append(leftSegs, insertions...)
}

func FindOverlapPairsSeg(evts CalendarEvents) (ret []CalendarPair) {
	length := len(evts)

	// if the input is less then 2 events then there must be no overlaps
	if length < 2 {
		return
	}
	pairs := make(map[CalendarPair]bool)

	segs := Segments{{Start: evts[0].Start, End: evts[0].End, Ids: []int{evts[0].Id}}}
	for i := 1; i < length; i++ {
		evt := evts[i]
		// find the segment which include start
		idx, err := segs.FindSeg(evt.Start)
		if err == errorTooSmall {
			// underflow
			segs = segs.AddSeg(0, evt.Id, evt.Start, evt.End, pairs)
		} else if err == errorTooLarge {
			// overflow
			segs = append(segs, Segment{Start: evt.Start, End: evt.End, Ids: []int{evt.Id}})
		} else if err == errorInBetween {
			if evt.Start >= segs[idx].Start {
				idx++
			}
			segs = segs.AddSeg(idx, evt.Id, evt.Start, evt.End, pairs)
		} else {
			// within
			for _, id := range segs[idx].Ids {
				pair := CalendarPair{FirstId: id, SecondId: evt.Id}
				if _, ok := pairs[pair]; !ok {
					pairs[pair] = true
				}
			}
			if evt.End <= segs[idx].End {
				// totally within
				splits := append(segs[idx].SplitWithin(evt), segs[idx+1:]...)
				segs = append(segs[:idx], splits...)
			} else {
				newStart := segs[idx].End + 1
				splits := segs[idx].SplitStart(evt)
				offset := len(splits)
				splits = append(splits, segs[idx+1:]...)
				segs = append(segs[:idx], splits...)
				segs = segs.AddSeg(idx+offset, evt.Id, newStart, evt.End, pairs)
			}
		}
	}
	for pair := range pairs {
		ret = append(ret, pair)
	}
	return
}

type bucketSeg struct {
	evts []*CalendarEvent
}

func (c *bucketSeg) Add(event *CalendarEvent) {
	c.evts = append(c.evts, event)
}

func (c *bucketSeg) GetPair(event *CalendarEvent) (ret []CalendarPair) {
	for _, e := range c.evts {
		if e.isOverlap(*event) {
			ret = append(ret, CalendarPair{e.Id, event.Id})
		}
	}
	return
}

func (c *CalendarEvent) getHash(size int64) (start, end int64) {
	start = c.Start / size
	end = c.End / size
	return
}

func FindOverlapPairBucket(evts CalendarEvents, size int64) (ret []CalendarPair) {
	buckets := make(map[int64]*bucketSeg)
	pairs := make(map[CalendarPair]bool)
	for idx := range evts {
		start, end := evts[idx].getHash(size)
		if bucket, ok := buckets[start]; !ok {
			buckets[start] = &bucketSeg{[]*CalendarEvent{&evts[idx]}}
		} else {
			newPairs := bucket.GetPair(&evts[idx])
			for i := 0; i < len(newPairs); i++ {
				pairs[newPairs[i]] = true
			}
			bucket.Add(&evts[idx])
		}
		if end > start {
			if bucket, ok := buckets[end]; !ok {
				buckets[end] = &bucketSeg{[]*CalendarEvent{&evts[idx]}}
			} else {
				newPairs := bucket.GetPair(&evts[idx])
				for i := 0; i < len(newPairs); i++ {
					pairs[newPairs[i]] = true
				}
				bucket.Add(&evts[idx])
			}
			for k := start + 1; k < end; k++ {
				if bucket, ok := buckets[end]; !ok {
					buckets[end] = &bucketSeg{[]*CalendarEvent{&evts[idx]}}
				} else {
					for i := 0; i < len(bucket.evts); i++ {
						pairs[CalendarPair{bucket.evts[i].Id, evts[idx].Id}] = true
					}
					bucket.Add(&evts[idx])
				}
			}
		}

	}
	for pair := range pairs {
		ret = append(ret, pair)
	}
	return
}
