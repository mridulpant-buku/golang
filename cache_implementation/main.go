package main

import (
	algo "cache_implementation/algorithms"
	cache "cache_implementation/cache"
	domain "cache_implementation/domain"
	"fmt"
)

func main() {

	d := algo.DoubleLinkedList{}
	d.InsertAtEnd(1)
	d.InsertAtEnd(2)
	d.InsertAtEnd(3)
	d.RemoveFromBeginning()
	d.Display()
	d.InsertAtEnd(5)
	d.Display()
	value := domain.Pokemon{Id: 1, Name: "pikachu", Height: 3.4, Weight: 40.3, Type: "htype", Abilities: "can swim"}
	hashMap := make(map[int]domain.Pokemon)
	hashMap[value.Id] = value
	c := cache.Cache{Capacity: 4, Storage: hashMap}
	fmt.Println(hashMap)
	fmt.Println(c.Storage)
	//fmt.Println(c.Dll)

	cache := cache.Cache{5, make(map[int]domain.Pokemon), algo.DoubleLinkedList{}}
	//cache.put(value.Id, value)

	pokemonDataset := []domain.Pokemon{
		{Id: 1, Name: "pikachu", Height: 3.4, Weight: 40.3, Type: "htype", Abilities: "can swim"},
		{Id: 2, Name: "bulbasaur", Height: 5.4, Weight: 60.2, Type: "ftype", Abilities: "can fire"},
		{Id: 3, Name: "raichu", Height: 4.3, Weight: 32, Type: "gtype", Abilities: "can bullet"},
	}

	for k, v := range pokemonDataset {
		cache.Set(k, v)
	}

	val, _ := cache.Get(1)
	fmt.Println(val)
}
