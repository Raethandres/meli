package domain

import (
	"testing"
	"github.com/matryer/is"
)

func TestGetLocation(t *testing.T) {

	t.Run("When get location with 3 distance params", func(t *testing.T) {
		is := is.New(t)
		x,y := GetLocation(485.41, 265.75, 600.52)
		
		is.Equal(float32(-99.99947), x)
		is.Equal(float32(74.9959), y)
	})


	t.Run("When try to get location with below 3 distance params", func(t *testing.T) {
		is := is.New(t)
		x,y := GetLocation(485.41, 265.75)
		
		is.Equal(float32(0), x)
		is.Equal(float32(0), y)
	})
}