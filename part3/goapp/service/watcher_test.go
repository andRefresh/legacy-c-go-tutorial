package service_test

import (
	"go-app/service"
	"testing"
	"time"
)

type MockSubscriber struct {
	Data MockData
}

type MockData struct {
	Field1 int
	Field2 int
}

func Test_StartWatch(t *testing.T) {
	subscriber := MockSubscriber{}

	w := service.NewWatcher()
	w.StartWatch("mock", subscriber.Data, 100)

	dataChanged := w.Subscribe(subscriber)
	timeout := time.NewTicker(200 * time.Millisecond)

	select {
	case data, open := <-dataChanged:
		if !open {
			t.Errorf("Error channel closed")
			return
		}

		if data.(service.WatchData).Data.(MockData).Field1 != 0 {
			t.Errorf("MockData values error")
		}

	case <-timeout.C:
		t.Errorf("data changed but no change event was raised")
		return
	}

}

func Test_Watch(t *testing.T) {
	subscriber := MockSubscriber{}

	w := service.NewWatcher()
	w.StartWatch("mock", &subscriber.Data, 50)

	dataChanged := w.Subscribe(subscriber)
	timeout := time.NewTicker(200 * time.Millisecond)

	<-dataChanged // First value discarded
	subscriber.Data.Field1 = 5

	select {
	case _, open := <-dataChanged:
		if !open {
			t.Errorf("Error channel closed")
			return
		}

		//if data.(service.WatchData).Data.(MockData).Field1 != 5 {
		//	t.Errorf("MockData values error. Field 1 was %d, should have been 5", data.(MockData).Field1)
		//}

	case <-timeout.C:
		t.Errorf("data changed but no change event was raised")
		return
	}

	subscriber.Data.Field1 = 6

	select {
	case _, open := <-dataChanged:
		if !open {
			t.Errorf("Error channel closed")
			return
		}

		//if data.(*MockData).Field1 != 6 {
		//	t.Errorf("MockData values error. Field 1 was %d, should have been 6", data.(MockData).Field1)
		//}

	case <-timeout.C:
		t.Errorf("data changed but no change event was raised")
		return
	}
}
