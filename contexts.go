package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/guregu/kami"
	"golang.org/x/net/context"
)

func indexHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	k := ctx.Value("key").(*Something)
	whatis := reflect.TypeOf(k)
	fmt.Println(k)
	fmt.Println(whatis)

	test := k.a
	fmt.Println(test)
}

type Something struct {
	a string
}

func main() {
	m := kami.New()

	ctx := context.Background()
	// ctx = context.WithValue(ctx, "key", "val")
	// val := struct {
	// 	a string
	// }{
	// 	"a string",
	// }

	val := Something{
		a: "string",
	}

	ctx = context.WithValue(ctx, "key", &val)
	m.Context = ctx

	m.Get("/", indexHandler)

	log.Fatalln(http.ListenAndServe(":8080", m))
}
