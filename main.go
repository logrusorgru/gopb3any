//
// Copyright (c) 2015 Konstanin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See LICENSE file for more details or see below.
//

//
//        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004
//
// Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>
//
// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.
//
//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
//  0. You just DO WHAT THE FUCK YOU WANT TO.
//

package main

import (
	// our types (messages)
	"github.com/logrusorgru/gopb3any/cat"
	"github.com/logrusorgru/gopb3any/user"

	// lifo storage
	"github.com/logrusorgru/gopb3any/lis"

	"fmt"
	"log"
)

func init() {
	// register our types
	lis.TypeReg.Set(new(user.Pro))
	lis.TypeReg.Set(new(cat.Cat))
}

// example function
func example_no1() error {
	usr := new(user.Pro)
	usr.Name = "Alice"
	usr.Age = 21
	usr.Languages = []string{"c", "go", "javascript"}
	// push
	err := repo.Push("alice", usr)
	if err != nil {
		return err
	}

	fmt.Printf("Length of repo is %d\n", repo.Len())
	k, alice, err := repo.Pop()
	if err != nil {
		return err
	}

	fmt.Println("Key:", k)
	fmt.Printf("Type is %T\n", alice)
	fmt.Printf("Name is %s\n", alice.(*user.Pro).Name)
	fmt.Printf("Age is %d\n", alice.(*user.Pro).Age)
	fmt.Printf("Languages are %v\n", alice.(*user.Pro).Languages)

	fmt.Println("===============================================")

	return nil
}

// create new lifo storage
var repo = new(lis.Repo)

func main() {

	// see function
	if err := example_no1(); err != nil {
		log.Println(err)
		return
	}

	// create units
	push_alice := &user.Pro{
		Name:      "Alice",
		Age:       21,
		Languages: []string{"c", "go", "javascript"},
	}

	push_tom := &cat.Cat{
		Name:         "Tom",
		TailLength:   40,
		FavoriteFood: "pork on coals",
	}

	push_bob := &user.Pro{
		Name:      "Bob",
		Age:       98,
		Languages: []string{"fortran", "ada", "assembler"},
	}

	push_kitty := &cat.Cat{
		Name:         "Kitty",
		TailLength:   25,
		FavoriteFood: "milk",
	}

	// push to repo
	if err := repo.Push("some key 1", push_alice); err != nil {
		log.Println(err)
		return
	}

	// push
	if err := repo.Push("some key 2", push_tom); err != nil {
		log.Println(err)
		return
	}

	// push
	if err := repo.Push("some key 3", push_bob); err != nil {
		log.Println(err)
		return
	}

	// push
	if err := repo.Push("some key 4", push_kitty); err != nil {
		log.Println(err)
		return
	}

	// show length of repo
	fmt.Printf("Length of repo is %d\n", repo.Len())
	fmt.Println("-----------------------------------------------")

	// pop pop pop pop
	for i := 0; i < 4; i++ {
		key, msg, err := repo.Pop()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Key: '%s'\n", key)
		// detect type of message
		switch t := msg.(type) {
		case *user.Pro:
			fmt.Printf(
				"Name: %s\nAge: %d\nLanguages: %v\n",
				t.Name, t.Age, t.Languages,
			)
		case *cat.Cat:
			fmt.Printf(
				"Name: %s\nTailLength: %d\nFavoriteFood: %v\n",
				t.Name, t.TailLength, t.FavoriteFood,
			)
		default:
			log.Printf("Type error, expected *user.Pro or *cat.Cat, go %T\n", t)
		}
		fmt.Println("-----------------------------------------------")
	}

	log.Println("Done")

}
