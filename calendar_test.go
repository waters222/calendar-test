package calendar

import (
	"testing"
)

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
	ret := FindOverlapPairs(events)
	if len(ret) != 1 {
		t.Logf("it should return 1 pair instead of %d", len(ret))
		t.FailNow()
	}
	if ret[0].FirstId != 0 || ret[0].SecondId != 1 {
		t.Logf("firstId %d, secondId %d is wrong", ret[0].FirstId, ret[0].SecondId)
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 99, 100}}
	ret = FindOverlapPairs(events)
	if len(ret) != 1 {
		t.Logf("it should return 1 pair instead of %d", len(ret))
		t.FailNow()
	}
	if ret[0].FirstId != 0 || ret[0].SecondId != 1 {
		t.Logf("firstId %d, secondId %d is wrong", ret[0].FirstId, ret[0].SecondId)
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 101, 200}}
	ret = FindOverlapPairs(events)
	if len(ret) != 0 {
		t.Logf("it should has no overlap")
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 101, 200}, {2, 0, 2000}}
	ret = FindOverlapPairs(events)
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
	ret = FindOverlapPairs(events)
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
