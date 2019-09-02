package utils

import (
	"fmt"
	"math"
	"testing"
	"time"
)

type Chat struct {
	name string
	legs int
}

func (c *Chat) String() string {
	return fmt.Sprintf("Chat(%v,%v)", c.name, c.legs)
}

func radioactive(t *testing.T, c *Chat) {
	t.Log(c)
	c.legs++
	t.Log(c)
}

func TestStruct(t *testing.T) {
	c := &Chat{
		name: "Amy",
		legs: 4,
	}
	t.Log(c)
	radioactive(t, c)
	t.Log(c)
}

func TestTimer(t *testing.T) {
	a := time.AfterFunc(1*time.Second, func() { t.Log("Hello") })

	c := <-a.C
	t.Logf("%v", c)

}

func TestExp(t *testing.T) {
	t0 := time.Now()
	DT := 10 * time.Millisecond
	f := func() float64 {
		dt := time.Now().Sub(t0)
		return 100 * math.Exp(-float64((dt / DT).Nanoseconds()))
	}

	for x := f(); x > 90; {
		t.Log(x)
		time.Sleep(DT)
		if time.Now().Sub(t0) > time.Second {
			x = 0
		}
	}
}
