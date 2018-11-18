package main

import (
	"fmt"
	"testing"
)

func TestPrintAdd(t *testing.T) {
	//t.SkipNow()
	for _, item := range [][]int{
		{1, 2, 3},
		{4, 5, 9},
		//{5, 6, 9},
	} {

		if PrintAdd(item[0], item[1]) != item[2] {
			t.Errorf("error")
		}
	}
}

func TestAll(t *testing.T)  {
	t.Run("TestPrintAdd", TestPrintAdd)
}

func TestMain(m *testing.M)  {
	fmt.Println("test main...")
	m.Run()
}


// 性能
func BenchmarkPrintAdd(b *testing.B) {
	for n:= 0 ; n < b.N; n++ {
		PrintAdd(1,2)

	}
}

/*
每个都要import testing

 */
