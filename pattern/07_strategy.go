package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
	An In-Memory-Cache it is of limited size.
	Whenever it reaches its maximum size that some old entries from the cache need to be evicted.
	This eviction can happen via several algorithms.

	Strategy pattern is used to decouple our Cache class with the algorithm such that
	we should be able to change the algorithm at run time.
	Also Cache class should not change when a new algorithm is being added.

	Our main Cache class will embed evictionAlgo interface.
	Instead of implementing all types of eviction algorithms in itself, our Cache class will
	delegate all it to the evictionAlgo interface.

	We need this when an object needs to support different behavior and you want to change
	the behavior at run time.
*/

type Cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) Get(key string) {
	delete(c.storage, key)
}

func InitCache(e evictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

type evictionAlgo interface {
	evict(c *Cache)
}

type fifo struct {
}

func (l *fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type lru struct {
}

func (l *lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}

type lfu struct {
}

func (l *lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
}

func StrategyExample() {
	lfu := &lfu{}
	cache := InitCache(lfu)
	cache.Add("a", "1")
	cache.Add("b", "2")
	cache.Add("c", "3")
	lru := &lru{}
	// changing the behavior at runtime
	cache.setEvictionAlgo(lru)
	cache.Add("d", "4")
	fifo := &fifo{}
	// changing the behavior at runtime
	cache.setEvictionAlgo(fifo)
	cache.Add("e", "5")
}
