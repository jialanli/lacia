package utils

import (
	"testing"
)

func TestCopyStructByCustomByTagGraceful(t *testing.T) {

	type example1 struct {
		Name    string `json:"name"`
		Age1    int    `json:"age"`
		GenderX string
	}

	type example2 struct {
		Name string `json:"name"`
		Age2 int    `json:"age"`
	}

	type example3 struct {
		Name3  string `json:"name"`
		Age3   string `json:"age"`
		Gender string `json:"gender"`
	}

	e1 := example1{Name: "name1", Age1: 99}

	e2 := example2{}
	CopyStructByCustomTagGraceful(&e1, &e2, "json")
	t.Logf("e2: %+v", e2)

	e3 := example3{}
	CopyStructByCustomTagGraceful(&e1, &e3, "json")
	t.Logf("e3: %+v", e3)
}

func TestCopyStructByTagGraceful(t *testing.T) {

	type example1 struct {
		Name    string `json:"name"`
		Age1    int    `json:"age"`
		GenderX string
	}

	type example2 struct {
		Name string `json:"name"`
		Age2 int    `json:"age"`
	}

	type example3 struct {
		Name3 string `json:"name"`
		Age3  int    `json:"age"`
		//Gender string `json:"gender"`
	}

	e1 := example1{Name: "name1", Age1: 99}

	e2 := example2{}
	CopyStructByTagGraceful(&e1, &e2)
	t.Logf("e2: %+v", e2)

	e3 := example3{}
	CopyStructByTagGraceful(&e1, &e3)
	t.Logf("e3: %+v", e3)
}

func TestCopyStructByFieldGraceful(t *testing.T) {

	type example1 struct {
		Name string
		Age  int
		//GenderX string
	}

	type example2 struct {
		Name string
		Age  int
	}

	type example3 struct {
		Name   string
		Age    int
		Gender string
	}

	e1 := example1{Name: "name1", Age: 99}

	e2 := example2{}
	CopyStructByFieldGraceful(&e1, &e2)
	t.Logf("e2: %+v", e2)

	e3 := example3{}
	CopyStructByFieldGraceful(&e1, &e3)
	t.Logf("e3: %+v", e3)
}
