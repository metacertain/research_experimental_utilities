// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main 

import (
	"fmt"
	"strconv"
)

var (
	debug = "lat"
)

func main() {

	fmt.Println("Hi")

	var (
		E = 3
		S [][]int
		R []int
	)

	R = make([]int, E)
	S = make([][]int, E)
	for i:=0; i < E; i++ {
		R[i] = 1;					// -> No. d'responses
		S[i] = make([]int, 2)
		S[i][0] = 2					// No. d'parametres
		S[i][1] = R[i]				// No. d'responses
	}
	
//
//
//
//		1st division = type (0 route 1 param 2 response)
//		2nd division = items ( 0 1  1 i  2 n )
//		3rd division = fields ( 
//								0 : method, path, lore
//								1 : name, lore
//								2 : status code, status text, responseitembody
//								)
//		
//
//
//																				//
//																				//
//		[0][0][0], [0][0][1], [0][0][2]											//
//			method		path		lore										//	apiroute {}{}{}	<- 3 * 1
//																				//
//		[1][ params ][0], [1][ params ][1]										//	routeparameter{}{} <- i * 2
//					name			lore										//
//																				//
//																				//	routeresponse { < n * 3
//		[2][ responses ][0], [2][ responses ][1], [2][ responses ][2]			//		routeResponseItem { status code } { status text } { RouteresponseItemBody}
//					status code			status text			responseitembody	//	
//																				//	}
//																				//
//		
	fmt.Println("\nEndpoints in spec\n\n")

	for j:=0; j < E; j++ {
		for i:=0; i < S[j][0]; i++ {

		}
	}

	
	Endpoints := make([][][][]string, E)
	for j:=0; j < E ; j++ {
		Endpoints[j] = make ([][][]string, 3)

		for k:=0; k < 3; k++ {
			switch k {

				case 0:
					Endpoints[j][k] = make([][]string, 1)
					Endpoints[j][k][0] = make([]string, 3)
					Endpoints[j][k][0] = []string{"MET", "/path" + strconv.Itoa(j), "Endpoint " + strconv.Itoa(j) + " Description"}
					fmt.Printf("%v: , %v: , 0: \n", j, k)

				case 1:
					Endpoints[j][k] = make([][]string, S[j][0] )
					for l:=0; l < S[j][0]; l++ {
						Endpoints[j][k][l] = make([]string, 2 )
						Endpoints[j][k][l] = []string{"Variable " + strconv.Itoa(l) + " name", "Variable " + strconv.Itoa(l) + " Description"}
						fmt.Printf("%v: , %v: , %v: \n", j, k, l)
					}

				case 2:
					Endpoints[j][k] = make([][]string, S[j][1] )
					for l:=0; l < S[j][1]; l++ {
						Endpoints[j][k][l] = make([]string, 3 )
						Endpoints[j][k][l] = []string{"Answer " + strconv.Itoa(l) + " code", "Answer " + strconv.Itoa(l) + " name", "Answer " + strconv.Itoa(l) + " description"}
						fmt.Printf("%v: , %v: , %v: \n", j, k, l)	
					}
				

			}

		}

	}

	
	fmt.Println(Endpoints)
	fmt.Println(S)
	printStruct( Endpoints );
}

// End of ma  i n

func printStruct(data [][][][]string) {
	fmt.Printf("\nVisualizer \n\n")

										fmt.Printf("\n\\UseRawInputEncoding\n")

	for whoah := range data {

		fmt.Printf("\n")

		for whoahtype := range data[whoah] {

				switch whoahtype {
					case 0:
										fmt.Printf("\n\n\\begin{apiRoute}{" + 
										data[whoah][0][0][0] + "}{" + 
										data[whoah][0][0][1] + "}{" + 
										data[whoah][0][0][2] + "}\n{\n}\n{ }\n")

					case 1:
										fmt.Printf("\n\\begin{routeParameter} ")
						
						for whoahitem, _ := range data[whoah][whoahtype] {

										fmt.Printf("\n\\routeParamItem")

							for whoahfield := range data[whoah][whoahtype][whoahitem] {

										fmt.Printf("{%v}", data[whoah][whoahtype][whoahitem][whoahfield] )

							}
						
						}
										fmt.Printf("\n\\end{routeParameter}")
				
					case 2: 
										fmt.Printf("\n\\begin{routeResponse}{application/json}")
						for whoahitem, _ := range data[whoah][whoahtype] {

										fmt.Printf("\n\\begin{routeResponseItem}")

							for whoahfield := range data[whoah][whoahtype][whoahitem] {
								if whoahfield < 2 {
										fmt.Printf("{%v}", data[whoah][whoahtype][whoahitem][whoahfield] )
								}
								if whoahfield == 2 {
										fmt.Printf("\n\\begin{routeResponseItemBody}\n %v \n\\end{routeResponseItemBody} }", data[whoah][whoahtype][whoahitem][whoahfield] )	
								}

							}
										fmt.Printf("\n\\end{routeResponseItem}")
					}
										fmt.Printf("\n\\end{routeResponse}")
				

				}	
		
		}
				fmt.Printf("\n\\end{apiRoute}")


	}




//	for _, whoha := range s {
//		for index, _ := range whoha {	
//			if index < 10 {
//				fmt.Printf("  %v          &", index )				
//			} else {
//				fmt.Printf(" %v          &", index )				
//			}
//		}
//		break
//	}
//	for woah, whoha := range s {
//		fmt.Printf("\n#### Gain of +%v #### & ", woah+1)
//		for _, val := range whoha {	
//			fmt.Printf(" %f   & ", val )				
//		}
//
//	}
//	fmt.Printf("\n")
//	if debug {
//		for woah, whoha := range s {
//			fmt.Printf("\n#### Debug %v ####\n", woah)
//			for index, val := range whoha {	
//				if val < 0 || val > 1 {
//					fmt.Printf("(D: %v | T: %v | %.4f ) & ", woah, index, val )				
//				}
//			}
//	
//		}
//	}
}
