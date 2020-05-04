// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main 

import (
	"fmt"
	"math"
//	"io/ioutil"
)

var (
	debug = false
)

func main() {

	fmt.Println("Hi")

	var (
		B int64 = 16
		ND int64 = 31
		G = int64( math.Log2(float64(B)) + 1 )
//		H = 0
		R float64 = 0
		final float64 = 0
		a int64 = 0
		i int64 = 0
		j int64 = 0
		n int64 = 0
	)

	fmt.Println("\nDiversity based gain probabilities for loss of T")

	var BDBG = make([][]float64, G)
	var Suite = make([]int64, G)

	fmt.Println("\nSuitable candidates by Proximity Gain G\n")
	for j=0; j < G ; j++ {
		Suite[j] = int64(math.Pow(2, float64( G-1-j )))
		fmt.Printf("%v \\\\\\|||/// %v \n", j+1, Suite[j])
	}
	
	for j=0; j < G; j++ {
		if debug {
			fmt.Printf("\n\n This is row %v calc debug \n\n", j)
		}
		BDBG[j] = make([]float64, B + 1)
		for i=0; i <= B; i++ {
				
				// This didnt work sane above some values
				//BDBG[j][i] = 1 - ( float64( Fctn( B - Suite[j] ) * Fctn( i ) ) / float64( Fctn( i - Suite[j] ) * Fctn( B ) ) )

				if debug {
					fmt.Printf("\n\nBEGINNING calculation for cell j: %v i: %v \n", j, i)
					fmt.Printf("\n S:%v/T:%v &  T!/(T-S)! :  %v  &  B!/(B-S)! (1/~ needed): %v  \n", Suite[j], i, intvFctnShw(i, Suite[j]), intvFctnShw( B, Suite[j] ) )
				}
				BDBG[j][i] = 1 - ( intvFctnShw( i, Suite[j] ) / intvFctnShw( B, Suite[j] ) )
				// This is only safe because Suite[j] is <= B, otherwise would deterministically become division by zero			
		}
	}

	printStruct(BDBG)

	var D float64

	fmt.Println("\nAverage hop count for ND with G\n")

	for n=0; n < ND; n++ {

		D = float64(n) / float64(G)
		R := int64(math.Floor( D ))
//		H := int64(math.Ceil( D ))
		b1 := n - ( R * G )
		fmt.Printf("########### N = %v, G = %v, A= %v, B= %v       Average hop count for Neigborhood depth: %v based on gain: %v :", n, G, R, b1, n, G)
		final = 0;
		if debug {
			fmt.Printf("\nConsisting of: \n")
		}
		for a = 0; a <= R ; a++ {
			tmp := a*G + b1 ;
			final += Sn1p2(tmp)
		}
		final += Sn1p2(n+1)
		if debug {
			fmt.Printf("(b) Result: ")
		}
		fmt.Printf("(%v)   %v \n", b1, final)
		//fmt.Printf("%v\n\n", H)

	}

	fmt.Println(R)

}

// End of main

// Utilities for avg hop count

// this is only safe/sane for non-negative n
func Sn1p2(n int64)(result float64) {
	if 0 <= n {
		powah := -1*float64(n)
		result = 1 - math.Pow(2, powah )
		if debug {
			fmt.Printf("Sn(%v) /// %v \n", n, result)
		}
		return
	}
	return 0
}


// Utilities for P(G, Lossof(T))

func Fctn(n int64)(result int64) {
	if (n > 0) {
		result = n * Fctn(n-1)
		return result
	} 
	if (n < 0) {
		fmt.Println("Done questionable math")
		return 0
	}
	return 1
}

// This function represents up! / (up-down)! or „last "down" terms of the up! factorial product”
func intvFctnShw(up, down int64)(result float64) {
	if debug {
		fmt.Printf("intvFctn called for %v  \\\\\\|||///  %v \n", up, down)
	}
	return intvFctn(up, down)
}

func intvFctn(up, down int64)(result float64) {
	var beat int64
	result = 1
	
	if down <= up {
		for beat = 0; beat < down; beat++ {
			if debug {
				fmt.Printf(" %v * ", up - beat)
			}
			result = result * float64(up - beat)
		}
		if debug { 
			fmt.Printf("\n")
		}
		return 
	}
	if up < down {
		if debug {
			fmt.Printf(" 0 * \n")
		}
		return 0
	}
	fmt.Printf("This just happened")
	return

}

func printStruct(s [][]float64) {
	fmt.Printf("\nVisualizer \n\n#### Loss of(T) #### &")

	for _, whoha := range s {
		for index, _ := range whoha {	
			if index < 10 {
				fmt.Printf("  %v          &", index )				
			} else {
				fmt.Printf(" %v          &", index )				
			}
		}
		break
	}
	for woah, whoha := range s {
		fmt.Printf("\n#### Gain of +%v #### & ", woah+1)
		for _, val := range whoha {	
			fmt.Printf(" %f   & ", val )				
		}

	}
	fmt.Printf("\n")
	if debug {
		for woah, whoha := range s {
			fmt.Printf("\n#### Debug %v ####\n", woah)
			for index, val := range whoha {	
				if val < 0 || val > 1 {
					fmt.Printf("(D: %v | T: %v | %.4f ) & ", woah, index, val )				
				}
			}
	
		}
	}
}
