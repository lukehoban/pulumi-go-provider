package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/blang/semver"
	provider "github.com/iwahbe/pulumi-go-provider"
	"github.com/iwahbe/pulumi-go-provider/resource"
)

type Enum int

const (
	A Enum = iota
	C
	T
	G
)

type Strct struct {
	Enum  Enum     `pulumi:"enum"`
	Names []string `pulumi:"names"`
}

func (s *Strct) Annotate(a resource.Annotator) {
	a.Describe(&s, "This is a holder for enums")
	a.Describe(&s.Names, "Names for the default value")

	a.SetDefault(&s.Enum, A)
}

func main() {
	println(reflect.TypeOf((*Enum)(nil)).Elem().String())

	err := provider.Run("schema-test", semver.Version{Minor: 1},
		provider.Types(
			provider.Enum[Enum](
				provider.EnumVal("A", A),
				provider.EnumVal("C", C),
				provider.EnumVal("T", T),
				provider.EnumVal("G", G)),
			&Strct{}),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
