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
	lis.TypeReg.Set(new(user.Pro))
	lis.TypeReg.Set(new(cat.Cat))
}

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

	fmt.Printf("Length of Repo is %d\n", repo.Len())
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

	if err := repo.Push("some key 2", push_tom); err != nil {
		log.Println(err)
		return
	}

	if err := repo.Push("some key 3", push_bob); err != nil {
		log.Println(err)
		return
	}

	if err := repo.Push("some key 4", push_kitty); err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < 4; i++ {
		key, msg, err := repo.Pop()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Key: '%s'\n", key)
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
			log.Printf("Type error, expected *user.Pro or *cat.Cat, go %T", t)
		}
		fmt.Println("-------------------------------------------------------")
	}

	log.Println("Done")

}
