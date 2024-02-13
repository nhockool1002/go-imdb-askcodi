package animals

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
)

type Animal struct {
	Name              string
	Progress          int
	Color             string
	ProgressFromRedis string
}

var (
	adjectives   = []string{"Loving", "Timid", "Furious", "Shiny", "Mechanical", "Pissed", "Cuddly"}
	animalNames  = []string{"Treefloof", "Murder Mittens", "Patience Monkey", "Forest Gorgi", "Wizard Cow", "Formal Chikcen"}
	colors       = []string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m"}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (a *Animal) Race(rdb *redis.Client) {
	status := rdb.Set(a.Name, a.Progress, 600*time.Second)
	if err := status.Err(); err != nil {
		log.Fatal("Redis error:", err)
	}
	val, err := rdb.Get(a.Name).Result()
	if err != nil {
		log.Fatal(err)
	}
	advance := rand.Intn(4)
	a.Progress += advance
	a.ProgressFromRedis = val
}

func CreateAnimal() Animal {
	adjective := adjectives[rand.Intn(len(adjectives))]
	name := animalNames[rand.Intn(len(animalNames))]
	color := colors[rand.Intn(len(colors))]
	return Animal{
		Name:  fmt.Sprintf("%s %s", adjective, name),
		Color: color,
	}
}

func IsDuplicate(animal Animal, animalList []Animal) bool {
	for _, a := range animalList {
		if animal.Name == a.Name {
			return true
		}
	}
	return false
}
