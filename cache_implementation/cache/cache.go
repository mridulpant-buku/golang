package cache

//what are the 2 things cache shall have.
//for now let us just assume we have a doubly linked list for LRU caching
import (
	algo "cache_implementation/algorithms"
	domain "cache_implementation/domain"
	exception "cache_implementation/exception"
	"fmt"
)

type Cache struct {
	Capacity int
	Storage  map[int]domain.Pokemon
	Dll      *algo.DoubleLinkedList
}

// creating a constructor in golang
func NewCache(capacity int) *Cache {
	return &Cache{
		capacity, make(map[int]domain.Pokemon), algo.NewDoublyLinkedList(),
	}
}

/*
- function cache to :
	- set, add an element to the key
	- remove
	- get

*/
//moveNodeToEnd is a combination of removeAtBeginning and insertAtEnd
//how the algorithm works for the LRU cache
//removeAtBeginning ->
//insertAtEnd ->
//leastused are at the front and most used are at the end.
func (self *Cache) Get(key int) (domain.Pokemon, error) {
	//returns a cache value based on the key
	//if key not present in self.Storage then throw an Error
	value, exists := self.Storage[key]
	if !exists {
		return domain.Pokemon{}, exception.KeyNotFoundError{Key: fmt.Sprintf("%v", key)}
	}
	self.Dll.RemoveFromBeginning()
	self.Dll.InsertAtEnd(key)
	return value, nil
}

func (self *Cache) Set(key int, value domain.Pokemon) {
	//set a value to the in-memory cache based on the storage and eviction policy.
	// based on the storage and eviction policy , how?

	if self.Dll.Len >= self.Capacity {
		self.Evict()
	}
	//case of updation
	_, exists := self.Storage[key]
	if exists {
		self.Dll.RemoveFromBeginning()
	}
	self.Storage[key] = value
	self.Dll.InsertAtEnd(key)

}

func (self *Cache) Delete(key int) error {
	//remove a value from the in-memory cache.
	_, exists := self.Storage[key]
	if !exists {
		return exception.KeyNotFoundError{Key: fmt.Sprintf("%v", key)}
	}
	delete(self.Storage, key)
	return nil
}

func (self *Cache) Evict() {
	// what is the strategy to evict?
	keyToDelete := self.Dll.RemoveFromBeginning()
	self.Delete(keyToDelete)
}

// when they are using pointers and & , and when they are not.
//only when you are changin the values in the parent struct type from a child struct , we require to pass by reference and not through pass by value.?
//In Go, when defining methods on a struct, it's common to use a pointer to the struct as the receiver type, especially when you want to modify the struct's fields within the method. This allows you to update the state of the cache. So, it's recommended to change the receiver type to a pointer:
