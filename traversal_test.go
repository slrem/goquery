package goquery

import (
	"testing"
)

func TestFind(t *testing.T) {
	sel := Doc().Root.Find("div.row-fluid")
	AssertLength(t, sel.Nodes, 9)
}

func TestFindNotSelf(t *testing.T) {
	sel := Doc().Root.Find("h1").Find("h1")
	AssertLength(t, sel.Nodes, 0)
}

func TestFindInvalidSelector(t *testing.T) {
	defer AssertPanic(t)
	Doc().Root.Find(":+ ^")
}

func TestChainedFind(t *testing.T) {
	sel := Doc().Root.Find("div.hero-unit").Find(".row-fluid")
	AssertLength(t, sel.Nodes, 4)
}

func TestChildren(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content").Children()
	AssertLength(t, sel.Nodes, 5)
}

func TestContents(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content").Contents()
	AssertLength(t, sel.Nodes, 13)
}

func TestChildrenFiltered(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content").ChildrenFiltered(".hero-unit")
	AssertLength(t, sel.Nodes, 1)
}

func TestContentsFiltered(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content").ContentsFiltered(".hero-unit")
	AssertLength(t, sel.Nodes, 1)
}

func TestChildrenFilteredNone(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content").ChildrenFiltered("a.btn")
	AssertLength(t, sel.Nodes, 0)
}

func TestParent(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid").Parent()
	AssertLength(t, sel.Nodes, 3)
}

func TestParentBody(t *testing.T) {
	sel := Doc().Root.Find("body").Parent()
	AssertLength(t, sel.Nodes, 1)
}

func TestParentFiltered(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid").ParentFiltered(".hero-unit")
	AssertLength(t, sel.Nodes, 1)
	AssertClass(t, sel, "hero-unit")
}

func TestParents(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid").Parents()
	AssertLength(t, sel.Nodes, 8)
}

func TestParentsFiltered(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid").ParentsFiltered("body")
	AssertLength(t, sel.Nodes, 1)
}
