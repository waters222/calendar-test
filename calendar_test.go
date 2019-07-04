package calendar

import (
	"testing"
)

func compareTwoPairArrays(left, right []CalendarPair) bool {
	if len(left) != len(right) {
		return false
	}
	for _, evt := range left {
		found := false
		for i, temp := range right {
			if evt.FirstId == temp.FirstId &&
				evt.SecondId == temp.SecondId {
				// found
				found = true
				if i+1 < len(right) {
					if i > 0 {
						right = append(right[:i-1], right[i+1:]...)
					} else {
						right = right[i+1:]
					}
				}
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestIsEventValid(t *testing.T) {
	cal := CalendarEvent{0, 100, 200}
	if !cal.IsValid() {
		t.Logf("%s should be valid", cal.ToString())
	}

	cal = CalendarEvent{0, 201, 200}
	if cal.IsValid() {
		t.Logf("%s should not be valid", cal.ToString())
	}
	cal = CalendarEvent{0, -200, -100}
	if cal.IsValid() {
		t.Logf("%s should not be valid", cal.ToString())
	}
}

func TestOverlap(t *testing.T) {
	cal := CalendarEvent{0, 100, 200}

	testCal := CalendarEvent{1, 100, 200}
	if !cal.isOverlap(testCal) {
		t.Logf("%s should be overlaps with %s", cal.ToString(), testCal.ToString())
		t.Fail()
	}

	testCal = CalendarEvent{1, 0, 300}
	if !cal.isOverlap(testCal) {
		t.Logf("%s should be overlaps with %s", cal.ToString(), testCal.ToString())
		t.Fail()
	}

	testCal = CalendarEvent{1, 101, 150}
	if !cal.isOverlap(testCal) {
		t.Logf("%s should be overlaps with %s", cal.ToString(), testCal.ToString())
		t.Fail()
	}

	testCal = CalendarEvent{1, 0, 100}
	if !cal.isOverlap(testCal) {
		t.Logf("%s should be overlaps with %s", cal.ToString(), testCal.ToString())
		t.Fail()
	}

	testCal = CalendarEvent{1, 200, 250}
	if !cal.isOverlap(testCal) {
		t.Logf("%s should be overlaps with %s", cal.ToString(), testCal.ToString())
		t.Fail()
	}

	testCal = CalendarEvent{1, 0, 99}
	if cal.isOverlap(testCal) {
		t.Logf("%s should not be overlaps with %s", cal.ToString(), testCal.ToString())
		t.Fail()
	}

	testCal = CalendarEvent{1, 201, 202}
	if cal.isOverlap(testCal) {
		t.Logf("%s should not be overlaps with %s", cal.ToString(), testCal.ToString())
		t.Fail()
	}

}

func TestCalendarBrutal(t *testing.T) {
	events := []CalendarEvent{{0, 0, 100}, {1, 0, 100}}
	ret := FindOverlapPairsBrutal(events)
	if len(ret) != 1 {
		t.Logf("it should return 1 pair instead of %d", len(ret))
		t.FailNow()
	}
	if ret[0].FirstId != 0 || ret[0].SecondId != 1 {
		t.Logf("firstId %d, secondId %d is wrong", ret[0].FirstId, ret[0].SecondId)
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 99, 100}}
	ret = FindOverlapPairsBrutal(events)
	if len(ret) != 1 {
		t.Logf("it should return 1 pair instead of %d", len(ret))
		t.FailNow()
	}
	if ret[0].FirstId != 0 || ret[0].SecondId != 1 {
		t.Logf("firstId %d, secondId %d is wrong", ret[0].FirstId, ret[0].SecondId)
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 101, 200}}
	ret = FindOverlapPairsBrutal(events)
	if len(ret) != 0 {
		t.Logf("it should has no overlap")
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 101, 200}, {2, 0, 2000}}
	ret = FindOverlapPairsBrutal(events)
	if len(ret) != 2 {
		t.Logf("it should return 2 pair instead of %d", len(ret))
		t.FailNow()
	}
	if ret[0].FirstId != 0 || ret[0].SecondId != 2 || ret[1].FirstId != 1 || ret[1].SecondId != 2 {
		t.Logf("first pair %v", ret[0])
		t.Logf("second pair %v", ret[1])
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 101, 200}, {2, 0, 2000}, {3, 101, 200}}
	ret = FindOverlapPairsBrutal(events)
	if len(ret) != 4 {
		t.Logf("it should return 4 pair instead of %d", len(ret))
		t.FailNow()
	}
	if ret[0].FirstId != 0 || ret[0].SecondId != 2 ||
		ret[1].FirstId != 1 || ret[1].SecondId != 2 ||
		ret[2].FirstId != 1 || ret[2].SecondId != 3 ||
		ret[3].FirstId != 2 || ret[3].SecondId != 3 {

		t.Logf("first pair %v", ret[0])
		t.Logf("second pair %v", ret[1])
		t.Logf("third pair %v", ret[2])
		t.Logf("fourth pair %v", ret[3])

		t.FailNow()
	}
}

func TestCalendar(t *testing.T) {
	events := []CalendarEvent{{0, 0, 100}, {1, 0, 100}}
	ret := FindOverlapPairsSort(events)
	if len(ret) != 1 {
		t.Logf("it should return 1 pair instead of %d", len(ret))
		t.FailNow()
	}
	if ret[0].FirstId != 0 || ret[0].SecondId != 1 {
		t.Logf("firstId %d, secondId %d is wrong", ret[0].FirstId, ret[0].SecondId)
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 99, 100}}
	ret = FindOverlapPairsSort(events)
	if len(ret) != 1 {
		t.Logf("it should return 1 pair instead of %d", len(ret))
		t.FailNow()
	}
	if ret[0].FirstId != 0 || ret[0].SecondId != 1 {
		t.Logf("firstId %d, secondId %d is wrong", ret[0].FirstId, ret[0].SecondId)
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 101, 200}}
	ret = FindOverlapPairsSort(events)
	if len(ret) != 0 {
		t.Logf("it should has no overlap")
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 101, 200}, {2, 0, 2000}}
	ret = FindOverlapPairsSort(events)
	if len(ret) != 2 {
		t.Logf("it should return 2 pair instead of %d", len(ret))
		t.FailNow()
	}
	if ret[0].FirstId != 0 || ret[0].SecondId != 2 || ret[1].FirstId != 1 || ret[1].SecondId != 2 {
		t.Logf("first pair %v", ret[0])
		t.Logf("second pair %v", ret[1])
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 101, 200}, {2, 0, 2000}, {3, 101, 200}}
	ret = FindOverlapPairsSort(events)
	if len(ret) != 4 {
		t.Logf("it should return 4 pair instead of %d", len(ret))
		t.FailNow()
	}
	if ret[0].FirstId != 0 || ret[0].SecondId != 2 ||
		ret[1].FirstId != 1 || ret[1].SecondId != 2 ||
		ret[2].FirstId != 2 || ret[2].SecondId != 3 ||
		ret[3].FirstId != 1 || ret[3].SecondId != 3 {

		t.Logf("first pair %v", ret[0])
		t.Logf("second pair %v", ret[1])
		t.Logf("third pair %v", ret[2])
		t.Logf("fourth pair %v", ret[3])

		t.FailNow()
	}
}

func TestSegment_IsWithin(t *testing.T) {
	seg := Segment{100, 200, nil}
	if seg.IsWithin(0) != -1 {
		t.FailNow()
	}
	if seg.IsWithin(201) != 1 {
		t.FailNow()
	}
	if seg.IsWithin(105) != 0 {
		t.FailNow()
	}
}

func TestSegment_SplitStart(t *testing.T) {
	seg := Segment{100, 200, []int{0}}

	ret := seg.SplitStart(CalendarEvent{1, 105, 203})
	if len(ret) != 2 {
		t.FailNow()
	}
	if ret[0].Start != seg.Start {
		t.FailNow()
	}
	if ret[0].End != 105-1 {
		t.FailNow()
	}
	if ret[0].Ids[0] != 0 {
		t.FailNow()
	}

	if ret[1].Start != 105 {
		t.FailNow()
	}
	if ret[1].End != seg.End {
		t.FailNow()
	}
	if ret[1].Ids[0] != 0 || ret[1].Ids[1] != 1 {
		t.FailNow()
	}
}

func TestSegment_SplitEnd(t *testing.T) {
	seg := Segment{100, 200, []int{0}}

	ret := seg.SplitEnd(CalendarEvent{1, 6, 105})
	if len(ret) != 2 {
		t.FailNow()
	}
	if ret[0].Start != seg.Start {
		t.FailNow()
	}
	if ret[0].End != 105 {
		t.FailNow()
	}
	if ret[0].Ids[0] != 0 && ret[0].Ids[1] != 1 {
		t.FailNow()
	}

	if ret[1].Start != 105+1 {
		t.FailNow()
	}
	if ret[1].End != seg.End {
		t.FailNow()
	}
	if ret[1].Ids[0] != 0 {
		t.FailNow()
	}
}

func TestSegment_SplitWithin(t *testing.T) {
	seg := Segment{100, 200, []int{0}}

	// total overlaps
	ret := seg.SplitWithin(CalendarEvent{1, 100, 200})
	if len(ret) != 1 {
		t.FailNow()
	}
	if ret[0].Start != seg.Start ||
		ret[0].End != seg.End ||
		ret[0].Ids[0] != 0 ||
		ret[0].Ids[1] != 1 {
		t.Log("split within total overlaps failed")
		t.FailNow()
	}

	ret = seg.SplitWithin(CalendarEvent{1, 100, 150})
	if len(ret) != 2 {
		t.FailNow()
	}
	if ret[0].Start != seg.Start ||
		ret[0].End != 150 ||
		ret[0].Ids[0] != 0 ||
		ret[0].Ids[1] != 1 ||
		ret[1].Start != 150+1 ||
		ret[1].End != seg.End ||
		ret[1].Ids[0] != 0 {
		t.Log("split within same start overlaps failed")
		t.FailNow()
	}

	ret = seg.SplitWithin(CalendarEvent{1, 150, 200})
	if len(ret) != 2 {
		t.FailNow()
	}
	if ret[0].Start != seg.Start ||
		ret[0].End != 150-1 ||
		ret[0].Ids[0] != 0 ||

		ret[1].Start != 150 ||
		ret[1].End != seg.End ||
		ret[1].Ids[0] != 0 ||
		ret[1].Ids[1] != 1 {
		t.Log("split within same end overlaps failed")
		t.FailNow()
	}

	ret = seg.SplitWithin(CalendarEvent{1, 150, 180})
	if len(ret) != 3 {
		t.FailNow()
	}
	if ret[0].Start != seg.Start ||
		ret[0].End != 150-1 ||
		ret[0].Ids[0] != 0 ||

		ret[1].Start != 150 ||
		ret[1].End != 180 ||
		ret[1].Ids[0] != 0 ||
		ret[1].Ids[1] != 1 ||

		ret[2].Start != 180+1 ||
		ret[2].End != seg.End ||
		ret[2].Ids[0] != 0 {
		t.Log("split within within overlaps failed")
		t.FailNow()
	}
}

func TestFindOverlapPairsSeg(t *testing.T) {
	events := []CalendarEvent{{0, 0, 100}, {1, 0, 100}}
	ret := FindOverlapPairsSeg(events)
	if len(ret) != 1 {
		t.Logf("it should return 1 pair instead of %d", len(ret))
		t.FailNow()
	}
	if ret[0].FirstId != 0 || ret[0].SecondId != 1 {
		t.Logf("firstId %d, secondId %d is wrong", ret[0].FirstId, ret[0].SecondId)
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 99, 100}}
	ret = FindOverlapPairsSeg(events)
	if len(ret) != 1 {
		t.Logf("it should return 1 pair instead of %d", len(ret))
		t.FailNow()
	}
	if ret[0].FirstId != 0 || ret[0].SecondId != 1 {
		t.Logf("firstId %d, secondId %d is wrong", ret[0].FirstId, ret[0].SecondId)
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 101, 200}}
	ret = FindOverlapPairsSeg(events)
	if len(ret) != 0 {
		t.Logf("it should has no overlap")
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 101, 200}, {2, 0, 2000}}
	ret = FindOverlapPairsSeg(events)
	if len(ret) != 2 {
		t.Logf("it should return 2 pair instead of %d", len(ret))
		t.FailNow()
	}
	if ret[0].FirstId != 0 || ret[0].SecondId != 2 || ret[1].FirstId != 1 || ret[1].SecondId != 2 {
		t.Logf("first pair %v", ret[0])
		t.Logf("second pair %v", ret[1])
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 101, 200}, {2, 0, 2000}, {3, 101, 200}}
	ret = FindOverlapPairsSeg(events)
	if len(ret) != 4 {
		t.Logf("it should return 4 pair instead of %d", len(ret))
		t.FailNow()
	}
	baseline := FindOverlapPairsBrutal(events)
	if !compareTwoPairArrays(ret, baseline) {
		t.FailNow()
	}
}
