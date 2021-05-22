package main

import (
	"os"
	"testing"

	"github.com/kchevelev/stretchbinarytree/data"
	"github.com/kchevelev/stretchbinarytree/util"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestStretchK2(t *testing.T) {
	// set test parameters
	root := data.GetSampleTree()
	k := 2

	// run function
	stretch(&root, k)

	//validate results
	var after string
	expected := "L:6 L:6 L:40 L:40 R:28 R:28 R:17 R:17 L:9 L:9 R:3 R:3 "
	util.FlattenTree(&root, "L", &after)
	if after != expected {
		t.Fatalf("expected %q,\n got %q", expected, after)
	}
}

func TestStretchK3(t *testing.T) {
	// set test parameters
	root := data.GetSampleTree()
	k := 3

	// run function
	stretch(&root, k)

	//validate results
	var after string
	expected := "L:4 L:4 L:4 L:27 L:27 L:27 R:18 R:18 R:18 R:11 R:11 R:11 L:6 L:6 L:6 R:2 R:2 R:2 "
	util.FlattenTree(&root, "L", &after)
	if after != expected {
		t.Fatalf("expected %q,\n got %q", expected, after)
	}
}
