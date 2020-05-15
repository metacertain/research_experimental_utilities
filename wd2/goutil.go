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
{ { {"PUT", "/bzz:/",	 																"Append to upload, returns new manifest"} 											},	{ {"file/collection", "as request body"} 															},	{ {"SWARM-TAG", "hex string" }, {"SWARM-STAMP", "hex string"}, {"SWARM-ENCRYPTION", "hex string"}, {"SWARM-PIN", "bool"}, {"SWARM-PARITIES", "integer"}  	} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"POST", "/bzz:/", 																	"Upload file or collection, returns tag used for tracking global sync state"} 		},	{ {"file/collection", "as request body"} 															},	{ {"SWARM-TAG", "hex string" }, {"SWARM-STAMP", "hex string"}, {"SWARM-ENCRYPTION", "hex string"}, {"SWARM-PIN", "bool"}, {"SWARM-PARITIES", "integer"}  	} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"POST", "/tags", 																	"Create new tag"} 																	},	{ 							 																		},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"GET", "/tags", 																	"Get all tags"} 																	},	{ 							 																		},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"GET", "/tags/\\{id\\}", 															"Path description"} 																},	{ {"id", "string"} 																					},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"DELETE", "/tags/\\{id\\}",														"Path description"} 																},	{ {"id", "string"}													 								},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"GET", "/stamp", 																	"View all postage stamps"} 															},	{ 							 																		},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"GET", "/stamp/{id}",																"View postage stamp with id"} 														},	{ {"id", ""} 																						},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"PUT", "/stamp/{id}", 																"Top up postage stamp with id"} 													},	{ {"id", ""} 																						},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"DELETE", "/stamp/{id}", 															"Drain and expire stamp with id"} 													},	{ {"id", ""} 																						},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"POST", "/stamp/{id}", 															"Create a new postage stamp"} 														},	{ {"id", ""} 																						},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"GET", "/pin/{id}?offset=\\{offset}&length=\\{length}", 							"View Pinned Content and metadata"} 												},	{ 							 																		},	{  																																							} 	,	{ {"offset", "integer"}, {"length", "integer"}				}  , 	{ { "200", "Ok", "" } } }, 
{ { {"POST", "/pin/{id}", 																"Pin an already uploaded content"} 													},	{ {"id", ""} 																						},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"PUT", "/pin/{id}", 																"Path description?"} 																},	{ {"id", ""} 																						},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"DELETE", "/pin/{id}", 															"Remove pinning from content"} 														},	{ {"id", ""} 																						},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"GET", "/chunk/\\{reference\\}", 													"Retrieve chunk as in \\ref\\{def:retrieve\\}"} 									},	{ {"reference", "string"} 																			},	{ {"SWARM-ENCRYPTION", "hex string"}																														} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"POST", "/chunk/(?span=\\{span\\})", 												"Create chunk as in \\ref\\{def:store\\}"} 											},	{ 							 																		},	{ {"SWARM-TAG", "hex string" }, {"SWARM-STAMP", "hex string"}, {"SWARM-ENCRYPTION", "hex string"}, {"SWARM-PIN", "bool"}, {"SWARM-PARITIES", "integer"}  	} 	,	{ {"span", "string"} 										}  , 	{ { "200", "Ok", "" } } }, 
{ { {"GET", "/soc/\\{owner\\}/\\{id\\}(?key=\\{key\\})", 								"Retrieve single owner chunk as in \\ref\\{def:soc-retrieve\\} "} 					},	{ {"owner", "eth address of single owner"}, {"id", "identifier within owner namespace"} 			},	{  																																							} 	,	{ {"key", "string"} 										}  , 	{ { "200", "Ok", "" } } }, 
{ { {"POST", "/soc/\\{owner\\}/\\{id\\}?span=\\{span\\}\\&encrypt=\\{encrypt\\}", 		"Create new single owner chunk as in \\ref\\{def:soc-store\\} "} 					},	{ {"owner", "eth address of single owner"}, {"id", "identifier within owner namespace"} 			},	{ {"SWARM-TAG", "hex string" }, {"SWARM-STAMP", "hex string"}, {"SWARM-ENCRYPTION", "hex string"}, {"SWARM-PIN", "bool"}, {"SWARM-PARITIES", "integer"}  	} 	,	{ {"span", "string"}, {"encrypt", "hex string"} 			}  , 	{ { "200", "Ok", "" } } }, 
{ { {"POST", "/file/",				 													"Path description as in \\ref\\{def:swarm-hash\\}"} 								},	{ 							 																		},	{ {"SWARM-TAG", "hex string" }, {"SWARM-STAMP", "hex string"}, {"SWARM-ENCRYPTION", "hex string"}, {"SWARM-PIN", "bool"}, {"SWARM-PARITIES", "integer"}  	} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"GET", "/file/\\{reference\\}", 													"Retrieve file by reference as in \\ref\\{def:file-retrieval\\}"} 					},	{ {"reference", "string"} 																			},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"PUT", "/file/ ", 																	"Update file, returns new reference"} 												},	{ {"file/collection", "as request body"} 															},	{ {"SWARM-TAG", "hex string" }, {"SWARM-STAMP", "hex string"}, {"SWARM-ENCRYPTION", "hex string"}, {"SWARM-PIN", "bool"}, {"SWARM-PARITIES", "integer"}  	} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"GET", "/manifest/\\{node:reference\\}/\\{path\\}", 								"Lookup manifest by path as in \\ref\\{def:manifests-lookup\\}"} 					},	{ {"node reference", "hex string"}, {"path", "string"} 												},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"DELETE", "/manifest/\\{node\\}/\\{path\\} ", 										"Delete manifest by path as in \\ref\\{def:manifests-remove\\}"} 					},	{ {"node", "string"}, {"path", "string"} 															},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"PUT", "/manifest/\\{old\\} ", 													"Update manifest as in \\ref\\{def:manifest-update\\}"} 							},	{ {"old", "string"} 																				},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"POST", "/manifest/\\{old\\}/\\{new\\}", 											"Merge manifests as in \\ref\\{def:manifest-merge\\}"} 								},	{ {"old", "string"}, {"new", "string"} 																},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"POST", "/access/\\{address\\} ", 													"Lock ACT for address as in \\ref\\{def:ac-api\\}"} 								},	{ {"address", "string"} 																			},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"GET", "/access/\\{address\\} ", 													"Unlock ACT for address as in \\ref\\{def:ac-api\\}"} 								},	{ {"address", "string"}																				},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"PUT", "/access/\\{root\\}/ ", 													"Add ACT for root hash as in \\ref\\{def:act-api\\}"} 								},	{ {"root", "string"} 																				},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"DELETE", "/access/\\{root\\} ", 													"Remove ACT for root hash as in \\ref\\{def:act-api\\}"} 							},	{ {"root", "string"}  																				},	{  																																							} 	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"POST", "/pss/\\{recipient\\}/\\{topic\\}(?targets=\\{targets\\})", 				"Send private message to recipient(s) as in \\ref\\{def:send\\}"}					},	{ {"recipient", "string"}, {"topic", "string"}														},	{																																							}	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
{ { {"GET", "/pss/\\{topic\\}/\\{channel\\}",											"Recieve on channel as in \\ref\\{def:recieve\\}"}									},	{ {"topic", "string"}, {"channel", "string"} 														},	{																																							}	,	{ 															}  , 	{ { "200", "Ok", "" } } }, 
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
