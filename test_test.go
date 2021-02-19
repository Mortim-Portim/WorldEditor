package main

import (
	"fmt"
	"testing"
)

func BenchmarkF(b *testing.B) {
	/*wrld := GE.GetWorldStructure(0, 0, 0, 0, 100, 100, 0, 0)

	for i := 0; i < 100; i++ {
		x := float64(rand.Intn(99))
		y := float64(rand.Intn(99))
		obj := &GE.StructureObj{Hitbox: GE.GetRectangle(x, y, x+1, y+1)}
		wrld.AddStructObj(obj)
	}*/

	for i := 0; i < 1000000; i++ {
		fmt.Print("D")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Print("F")
		//GE.FindPathMat(wrld, [2]int{0, 0}, [2]int{99, 99})
	}

	fmt.Println()
}

//703548
