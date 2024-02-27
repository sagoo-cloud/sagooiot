package product

import (
	"context"
	"testing"
)

func TestInitProductForTd(t *testing.T) {
	s := devInitNew()
	if err := s.InitProductForTd(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func TestInitDeviceForTd(t *testing.T) {
	s := devInitNew()
	if err := s.InitDeviceForTd(context.TODO()); err != nil {
		t.Fatal(err)
	}
}
