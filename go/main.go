/* 
 Copyright (c) 2017, 2018, 2019, 2020, 2021 KhaiPhong

 Licensed under the Apache License, Version 2.0 (the "License");

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  
Home is an injected microservice for (1) Facebook clone / Home pages,
(2) Twitter clone / Mu / Miss You, (3) LinkedIn clone / Confidential HR

Each active user has its Home / Mu / CHR from the Home services.
*/

package main

import "fmt"	
	
func main() {
  fmt.Println("Hello, 世界")
}

/*
  (1) Home page templates,

  (2) Miss You Mu with one paragraph,
  
  (3) Confidential HR
  
func Home() { //to enable Home <=> its aggregate
  fmt.Printf("Get connection to the user private/ public Home")
}
func Mu() { //to enable Mu <=> its aggregate
  fmt.Printf("Get connection to the user private/ public Mu")
}
func Chr() { //to enable Chr <=> public/private HR
  fmt.Printf("Get connection to the user data store")
}

*/
