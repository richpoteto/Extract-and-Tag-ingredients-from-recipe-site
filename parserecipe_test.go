package extractrecipe

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	log "github.com/schollz/logger"
	"github.com/stretchr/testify/assert"
)

func BenchmarkParse(b *testing.B) {
	log.SetLevel("error")
	for n := 0; n < b.N; n++ {
		r, _ := NewFromFile("testing/sites/lasagna.html")
		r.Parse()
	}
}

// TODO: test non recipes
// func TestParseNonRecipes(t *testing.T) {
// 	files := []string{
// 		"testing/sites/spain-europe-middle-class.html",
// 	}
// 	for _, f := range files {
// 		log.Infof("working on %s", f)
// 		r, err := NewFromFile(f)
// 		assert.Nil(t, err)
// 		err = r.Parse()
// 		assert.NotNil(t, err)
// 	}
// }

func ExampleRecipe1() {
	r, _ := NewFromURL("https://joyfoodsunshine.com/the-most-amazing-chocolate-chip-cookies/")
	r.Parse()
	fmt.Println(r.PrintIngredientList())
	fmt.Println(r)
	// Output:
	// 1 cup butter
	// 1 cup white sugar
	// 1 cup brown sugar
	// 2 tsp vanilla
	// 2 whole eggs
	// 3 cups flour
	// 1 tsp baking soda
	// 1/2 tsp baking powder
	// 1 tsp salt
	// 2 cups chocolate
}

func TestParseRecipes(t *testing.T) {
	urls := []string{}
	err := filepath.Walk("testing/sites", func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			urls = append(urls, strings.TrimLeft(path, "testing/sites/"))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, urlToTest := range urls {
		pathToFile := "testing/sites/" + urlToTest
		log.Infof("working on %s", pathToFile)
		r, err := NewFromFile(pathToFile)
		assert.Nil(t, err)
		err = r.Parse()
		assert.Nil(t, err)
		fmt.Println(urlToTest)
		fmt.Println(r.IngredientList())
	}
}
