package db

import (
	"reflect"
	"testing"

	"github.com/arijitnayak92/taskAfford/TODONEW/schema"
	"github.com/arijitnayak92/taskAfford/TODONEW/testdb"
	_ "github.com/lib/pq"
)

//...
func TestPostgres_Insert(t *testing.T) {
	postgres := &Postgres{testdb.Setup()}
	defer postgres.Close()

	todo := &schema.Todo{
		Title:  "title1",
		Note:   "note1",
		Status: false,
	}

	got, err := postgres.Insert(todo)
	if err != nil {
		t.Fatal(err)
	}

	want := 1

	if got != want {
		t.Fatal(err)
	}
}

//...
func TestPostgres_GetAll(t *testing.T) {
	postgres := &Postgres{testdb.Setup()}
	defer postgres.Close()

	todo := &schema.Todo{
		Title:  "title1",
		Note:   "note1",
		Status: false,
	}

	_, err := postgres.Insert(todo)
	if err != nil {
		t.Fatal(err)
	}

	got, err := postgres.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	want := []schema.Todo{
		{
			ID:     1,
			Title:  "title1",
			Note:   "note1",
			Status: false,
		},
	}

	if equal(got, want) {
		t.Fatalf("Want: %v, Got: %v", want, got)
	}
}

//...
func TestPostgres_Delete(t *testing.T) {
	postgres := &Postgres{testdb.Setup()}
	defer postgres.Close()

	todo := &schema.Todo{
		Title:  "title1",
		Note:   "note1",
		Status: false,
	}

	id, err := postgres.Insert(todo)
	if err != nil {
		t.Fatal(err)
	}

	err = postgres.Delete(id)
	if err != nil {
		t.Fatal(err)
	}

	got, err := postgres.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(got) > 0 {
		t.Fatal("The record is not deleted.")
	}
}

func equal(got interface{}, want interface{}) bool {
	return reflect.DeepEqual(got, want)
}