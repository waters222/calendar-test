package calendar

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

var testEvents []CalendarEvent

func TestMain(m *testing.M) {

	rand.Seed(111)
	count := 500
	maxValue := 1000
	// lets create benchmark needed array
	testEvents = make([]CalendarEvent, count)
	for i := 0; i < count; i++ {
		start := rand.Int63n(int64(maxValue))
		left := int64(maxValue) - start
		end := rand.Int63n(left) + start
		temp := CalendarEvent{i, start, end}
		if !temp.IsValid() {
			fmt.Println(fmt.Sprintf("[ERROR] calenar event is invalid: %s", temp.ToString()))
			os.Exit(1)
		}
		testEvents[i] = temp
	}
	m.Run()
}

func compareTwoPairArrays(left, right []CalendarPair) (bool, []CalendarPair, []CalendarPair) {
	ret := true
	if len(left) != len(right) {
		ret = false
	}

	length := len(right)

	var leftNotFound []CalendarPair

	for _, evt := range left {
		found := false
		for i := 0; i < length; i++ {
			temp := right[i]
			if evt.FirstId == temp.FirstId &&
				evt.SecondId == temp.SecondId {
				// found
				found = true
				if i+1 < len(right) {
					right[i] = right[length-1]
					length--
				}
				break
			}
		}
		if !found {
			ret = false
			leftNotFound = append(leftNotFound, evt)
		}
	}
	return ret, leftNotFound, right[0:length]
}

func printPairs(evts []CalendarEvent, pairs []CalendarPair) {
	for _, pair := range pairs {
		fmt.Println(fmt.Sprintf("pair: %d, %d -> %d, %d & %d, %d", pair.FirstId, pair.SecondId,
			evts[pair.FirstId].Start, evts[pair.FirstId].End,
			evts[pair.SecondId].Start, evts[pair.SecondId].End))
	}

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

func TestSegments_FindSeg(t *testing.T) {
	segs := Segments{
		{10, 100, nil},
		{200, 300, nil},
		{400, 500, nil},
		{600, 700, nil},
		{800, 900, nil},
	}

	if _, err := segs.FindSeg(9); err != errorTooSmall {
		t.Log("should be too small")
		t.FailNow()
	}

	if _, err := segs.FindSeg(901); err != errorTooLarge {
		t.Log("should be too large")
		t.FailNow()
	}

	if idx, err := segs.FindSeg(101); err != errorInBetween {
		t.Log("should be in between")
		t.FailNow()
	} else if idx != 0 {
		t.Logf("in between should be 0 instead of %d", idx)
		t.FailNow()
	}

	if idx, err := segs.FindSeg(350); err != errorInBetween {
		t.Log("should be in between")
		t.FailNow()
	} else if idx != 1 {
		t.Logf("in between should be 1 instead of %d", idx)
		t.FailNow()
	}

	if _, err := segs.FindSeg(501); err != errorInBetween {
		t.Log("should be in between")
		t.FailNow()
	}
	if _, err := segs.FindSeg(701); err != errorInBetween {
		t.Log("should be in between")
		t.FailNow()
	}

	if idx, err := segs.FindSeg(201); err != nil || idx != 1 {

		t.FailNow()
	}

	if idx, err := segs.FindSeg(20); err != nil || idx != 0 {

		t.FailNow()
	}
	if idx, err := segs.FindSeg(450); err != nil || idx != 2 {

		t.FailNow()
	}
	if idx, err := segs.FindSeg(800); err != nil || idx != 4 {

		t.FailNow()
	}

	segs = Segments{
		{10, 15, nil},
		{16, 71, nil},
	}
	if idx, err := segs.FindSeg(16); err != nil || idx != 1 {
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
		t.Logf("it should has no overlap: %v", ret)
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 101, 200}, {2, 0, 2000}}
	ret = FindOverlapPairsSeg(events)
	if len(ret) != 2 {
		t.Logf("it should return 2 pair instead of %d", len(ret))
		t.FailNow()
	}
	events = []CalendarEvent{{0, 0, 100}, {1, 101, 200}, {2, 0, 2000}, {3, 101, 200}}
	ret = FindOverlapPairsSeg(events)
	if len(ret) != 4 {
		t.Logf("it should return 4 pair instead of %d", len(ret))
		t.FailNow()
	}
	baseline := FindOverlapPairsBrutal(events)
	if isSame, leftNotFound, rightNotFound := compareTwoPairArrays(ret, baseline); !isSame {
		printPairs(events, leftNotFound)
		fmt.Println("=========================")
		printPairs(events, rightNotFound)
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 130, 200}, {3, 150, 160}, {4, 110, 120}}
	ret = FindOverlapPairsSeg(events)
	if len(ret) != 1 {
		t.Logf("it should return  pair instead of %d", len(ret))
		t.FailNow()
	}
	baseline = FindOverlapPairsBrutal(events)
	if isSame, leftNotFound, rightNotFound := compareTwoPairArrays(ret, baseline); !isSame {
		printPairs(events, leftNotFound)
		fmt.Println("=========================")
		printPairs(events, rightNotFound)
		t.FailNow()
	}
}

func TestFindOverlapPairBucket(t *testing.T) {
	events := []CalendarEvent{{0, 0, 100}, {1, 0, 110}}
	ret := FindOverlapPairBucket(events, 50)
	if len(ret) != 1 {
		t.Logf("it should return 1 pair instead of %d", len(ret))
		t.FailNow()
	}
	if ret[0].FirstId != 0 || ret[0].SecondId != 1 {
		t.Logf("firstId %d, secondId %d is wrong", ret[0].FirstId, ret[0].SecondId)
		t.FailNow()
	}

	events = []CalendarEvent{{0, 0, 100}, {1, 101, 200}, {2, 0, 2000}}
	ret = FindOverlapPairBucket(events, 50)
	if len(ret) != 2 {
		t.Logf("it should return 2 pair instead of %d", len(ret))
		t.FailNow()
	}

}

func TestFindOverlapPairsSort(t *testing.T) {
	events := make([]CalendarEvent, len(testEvents))
	copy(events, testEvents)
	ret := FindOverlapPairsSort(events)
	baseline := FindOverlapPairsBrutal(events)
	if isSame, leftNotFound, rightNotFound := compareTwoPairArrays(ret, baseline); !isSame {
		printPairs(events, leftNotFound)
		fmt.Println("=========================")
		printPairs(events, rightNotFound)
		t.FailNow()
	}
}

func TestFindOverlapPairsSeg2(t *testing.T) {
	events := make([]CalendarEvent, len(testEvents))
	copy(events, testEvents)
	ret := FindOverlapPairsSeg(events)
	baseline := FindOverlapPairsBrutal(events)
	if isSame, leftNotFound, rightNotFound := compareTwoPairArrays(ret, baseline); !isSame {
		fmt.Println("==== seg ====")
		printPairs(events, leftNotFound)
		fmt.Println("########")
		fmt.Println("==== baseline ====")
		printPairs(events, rightNotFound)
		fmt.Println("########")
		t.FailNow()
	}
}

//
func Benchmark_brutal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindOverlapPairsBrutal(testEvents)
	}

}
func Benchmark_bucket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindOverlapPairBucket(testEvents, 200)
	}
}

func Benchmark_seq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindOverlapPairsBrutal(testEvents)
	}
}

func Benchmark_sort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindOverlapPairsBrutal(testEvents)
	}
}
