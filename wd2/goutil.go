// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main 

import (
	"fmt"
)

var (
	debug = "lat"
)

func main() {

	fmt.Println("Hi")

//
//
//
//		1st division = type (0 route 1 param 2 response)
//		2nd division = items ( 0 ~ 1   1 ~ i   2 ~ j   3 ~ k   4 ~ n )
//		3rd division = fields ( 
//								0 : method, path, lore
//								1 : name, lore
//								2 : name, lore
//								3 : name, lore								
//								4 : status code, status text, responseitembody
//								)
//		
//
//
//																				//
//																				//
//		[0][0][0], [0][0][1], [0][0][2]											//
//			method		path		lore										//	apiroute {0}{0}{}	<- 1 * 3
//																				//
//		[1][ params ][0], [1][ params ][1]										//	routeparameter[1]{i}{}  <- i * 2
//					name			lore										//
//																				//
//		[1][ params ][0], [1][ params ][1]										//	headerparameter[2]{j}{} <- j * 2
//					name			lore										//
//																				//
//		[1][ params ][0], [1][ params ][1]										//	queryparameter{3}{k} <- k * 2
//					name			lore										//
//																				//
//																				//	routeresponse{4}{n} <- n * 3
//		[2][ responses ][0], [2][ responses ][1], [2][ responses ][2]			//		routeResponseItem { status code } { status text } { RouteresponseItemBody}
//					status code			status text			responseitembody	//	
//																				//	}
//																				//
//		

	Lol := [][][][]string{
{ { {"PUT", "/bzz:/", 																			" Path description"} } , { {"file/collection", "as request body"} 															}	,		{ {"SWARM-TAG", "hex string" }, {"SWARM-STAMP", "hex string"}, {"SWARM-ENCRYPTION", "hex string"}, {"SWARM-PIN", "bool"}, {"SWARM-PARITIES", "integer"}  	},	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"POST", "/bzz:/", 																			" Path description"} } , { {"file/collection", "as request body"} 															}	,		{ {"SWARM-TAG", "hex string" }, {"SWARM-STAMP", "hex string"}, {"SWARM-ENCRYPTION", "hex string"}, {"SWARM-PIN", "bool"}, {"SWARM-PARITIES", "integer"}  	} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"POST", "/tags", 																			" Path description"} } , { {"other", "other parameter"} 																	}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"GET", "/tags", 																			" Path description"} } , { {"other", "other parameter"} 																	}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"GET", "/tags/\\{id\\}", 																	" Path description"} } , { {"id", "string"} 																				}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"DELETE", "/tags/\\{id\\}",																" Path description"} } , { {"id", "string"} 																				}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"GET", "/stamp", 																			" Path description"} } , { {"other", "other parameter"} 																	}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"GET", "/stamp/{id}",																		" Path description"} } , { {"other", "other parameter"} 																	}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"PUT", "/stamp", 																			" Path description"} } , { {"other", "other parameter"} 																	}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"DELETE", "/stamp", 																		" Path description"} } , { {"other", "other parameter"} 																	}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"POST", "/stamp", 																			" Path description"} } , { {"other", "other parameter"} 																	}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"GET", "/pin", 																			" Path description"} } , { {"other", "other parameter"} 																	}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"PUT", "/pin", 																			" Path description"} } , { {"other", "other parameter"} 																	}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"DELETE", "/pin", 																			" Path description"} } , { {"other", "other parameter"} 																	}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"GET", "/chunk/\\{reference\\}", 															" Path description"} } , { {"reference", "string"} 																			}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"POST", "/chunk/(?span=\\{span\\})", 														" Path description"} } , { 																									}	,		{ {"SWARM-TAG", "hex string" }, {"SWARM-STAMP", "hex string"}, {"SWARM-ENCRYPTION", "hex string"}, {"SWARM-PIN", "bool"}, {"SWARM-PARITIES", "integer"}  	} ,	 	{ {"span", "string"} 										} , 		{ { "204", "No Content", "" } } }, 
{ { {"GET", "/soc/\\{owner\\}/\\{id\\}(?key=\\{key\\})", 										" Path description"} } , { {"owner", "string"}, {"id", "string"} 															}	,		{  																																							} ,	 	{ {"key", "string"} 										} , 		{ { "204", "No Content", "" } } }, 
{ { {"POST", "/soc/\\{owner\\}/\\{id\\}?span=\\{span\\}\\&encrypt=\\{encrypt\\}", 				" Path description"} } , { {"owner", "string"}, {"id", "string"} 															}	,		{ {"SWARM-TAG", "hex string" }, {"SWARM-STAMP", "hex string"}, {"SWARM-ENCRYPTION", "hex string"}, {"SWARM-PIN", "bool"}, {"SWARM-PARITIES", "integer"}  	} ,	 	{ {"span", "string"}, {"encrypt", "hex string"} 			} , 		{ { "204", "No Content", "" } } }, 
{ { {"POST", "/file/",				 															" Path description"} } , { {"other", "other parameter"} 																	}	,		{ {"SWARM-TAG", "hex string" }, {"SWARM-STAMP", "hex string"}, {"SWARM-ENCRYPTION", "hex string"}, {"SWARM-PIN", "bool"}, {"SWARM-PARITIES", "integer"}  	} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"GET", "/file/\\{reference\\}", 															" Path description"} } , { {"reference", "string"} 																			}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"PUT", "/file/ ", 																			" Path description"} } , { {"file/collection", "as request body"} 															}	,		{ {"SWARM-TAG", "hex string" }, {"SWARM-STAMP", "hex string"}, {"SWARM-ENCRYPTION", "hex string"}, {"SWARM-PIN", "bool"}, {"SWARM-PARITIES", "integer"}  	} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"GET", "/manifest/\\{node:reference\\}/\\{path\\}", 										" Path description"} } , { {"node reference", "hex string"}, {"path", "string"} 											}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"DELETE", "/manifest/\\{node\\}/\\{path\\} ", 												" Path description"} } , { {"node", "string"}, {"path", "string"} 															}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"PUT", "/manifest/\\{old\\} ", 															" Path description"} } , { {"old", "string"} 																				}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"POST", "/access/\\{address\\} ", 															" Path description"} } , { {"address", "string"} 																			}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"GET", "/access/\\{address\\} ", 															" Path description"} } , { {"address", "string"} 																			}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"PUT", "/access/\\{root\\}/ ", 															" Path description"} } , { {"root", "string"} 																				}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"DELETE", "/access/\\{root\\} ", 															" Path description"} } , { {"root", "string"} 																				}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
{ { {"POST", "/pss/\\{topic\\} ", 																" Path description"} } , { {"topic", "string"} 																				}	,		{  																																							} ,	 	{ 															} , 		{ { "204", "No Content", "" } } }, 
	} 	


	
	fmt.Println(Lol)
	printStruct( Lol );
}

// End of ma  i n

func printStruct(data [][][][]string) {
	fmt.Printf("\nVisualizer \n\n")

											fmt.Printf("\n\\UseRawInputEncoding\n")

	for whoah := range data {

		fmt.Printf("\n")

		for whoahtype := range data[whoah] {
				if ( len(data[whoah][whoahtype]) > 0 ) {
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

											fmt.Printf("\n\\begin{headerParameter} ")
							
							for whoahitem, _ := range data[whoah][whoahtype] {

											fmt.Printf("\n\\headerParamItem")

								for whoahfield := range data[whoah][whoahtype][whoahitem] {

											fmt.Printf("{%v}", data[whoah][whoahtype][whoahitem][whoahfield] )

								}
							
							}
											fmt.Printf("\n\\end{headerParameter}")

						case 3:

											fmt.Printf("\n\\begin{queryParameter} ")
							
							for whoahitem, _ := range data[whoah][whoahtype] {

											fmt.Printf("\n\\queryParamItem")

								for whoahfield := range data[whoah][whoahtype][whoahitem] {

											fmt.Printf("{%v}", data[whoah][whoahtype][whoahitem][whoahfield] )

								}
							
							}
											fmt.Printf("\n\\end{queryParameter}")

						case 4: 
											fmt.Printf("\n\\begin{routeResponse}{application/json}")
							for whoahitem, _ := range data[whoah][whoahtype] {

											fmt.Printf("\n\\begin{routeResponseItem}")

								for whoahfield := range data[whoah][whoahtype][whoahitem] {
									if whoahfield < 2 {
											fmt.Printf("{%v}", data[whoah][whoahtype][whoahitem][whoahfield] )
									}
									if whoahfield == 2 {
											fmt.Printf("\n\\begin{routeResponseItemBody}\n %v \n\\end{routeResponseItemBody}", data[whoah][whoahtype][whoahitem][whoahfield] )	
									}

								}
											fmt.Printf("\n\\end{routeResponseItem}")
						}
											fmt.Printf("\n\\end{routeResponse}")
					

					}	
				}
		
		}
				fmt.Printf("\n\\end{apiRoute}\n\\\\\n\\\\\n")


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
