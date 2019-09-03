/*
Package bsr is generated by gogll. Do not edit.

Copyright Marius Ackerman 2019

It implements a Binary Subtree Representation set as defined in

	Scott et al
	Derivation representation using binary subtree sets,
	Science of Computer Programming 175 (2019)

*/
package bsr

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/goccmack/gogll/examples/ambiguous1/parser/slot"
)

type bsr interface {
	LeftExtent() int
	RightExtent() int
	Pivot() int
}

var (
	set *bsrSet
	startSym string
)

type bsrSet struct {
	slotEntries   map[BSR]bool
	ignoredSlots   map[BSR]bool
	stringEntries map[stringBSR]bool
	rightExtent   int
	I			  []byte
}

// BSR is the binary subtree representation of a parsed nonterminal
type BSR struct {
	Label       slot.Label
	leftExtent  int
	pivot       int
	rightExtent int
}

type stringBSR struct {
	Label       slot.Label
	leftExtent  int
	pivot       int
	rightExtent int
}

func newSet(I []byte) *bsrSet {
	return &bsrSet{
		slotEntries:   make(map[BSR]bool),
		ignoredSlots:   make(map[BSR]bool),
		stringEntries: make(map[stringBSR]bool),
		rightExtent:	0,
		I: I,
	}
}

/*
Add a bsr to the set. (i,j) is the extent. k is the pivot.
*/
func Add(l slot.Label, i, k, j int) {
	// fmt.Printf("bsr.Add(%s,%d,%d,%d)\n", l,i,k,j)
	if l.EoR() {
		insert(BSR{l, i, k, j})
	} else {
		if l.Pos() > 1 {
			insert(stringBSR{l, i, k, j})
		}
	}
}

func AddEmpty(l slot.Label, i int) {
	insert(BSR{l, i, i, i})
}

func Contain(nt string, left, right int) bool {
	// fmt.Printf("bsr.Contain(%s,%d,%d)\n",nt,left,right)
	for e, _ := range set.slotEntries {
		// fmt.Printf("  (%s,%d,%d)\n",e.Label.Head(),e.leftExtent,e.rightExtent)
		if e.Label.Head() == nt && e.leftExtent == left && e.rightExtent == right {
			// fmt.Println("  true")
			return true
		}
	}
	// fmt.Println("  false")
	return false
}

// GetAll returns all BSR grammar slot entries
func GetAll() (bsrs []BSR) {
	for b := range set.slotEntries {
		bsrs = append(bsrs, b)
	}
	return
}

// GetRightExtent returns the right extent of the BSR set
func GetRightExtent() int {
	return set.rightExtent
}

// GetRoot returns the root of the parse tree of an unambiguous parse. 
// GetRoot fails if the parse was ambiguous. Use GetRoots() for ambiguous parses.
func GetRoot() BSR {
	rts := GetRoots()
	if len(rts) != 1 {
		failf("%d parse trees exist for start symbol %s", len(rts), startSym)
	}
	return rts[0]
}

// GetRoots returns all the roots of parse trees of the start symbol of the grammar.
func GetRoots() (roots []BSR) {
	for s, _ := range set.slotEntries {
		if s.Label.Head() == startSym && s.leftExtent == 0 && s.rightExtent == set.rightExtent {
			roots = append(roots, s)
		}
	}
	return
}

// GetNTSlot returns all slot entries of the NT, 'nt'
func GetNTSlot(nt string) (slots []BSR) {
	for bsr := range set.slotEntries {
		if bsr.Label.Head() == nt {
			slots = append(slots, bsr)
		}
	}
	return
}

func getString(l slot.Label, leftExtent, rightExtent int) stringBSR {
	for str, _ := range set.stringEntries {
		if str.Label == l && str.leftExtent == leftExtent && str.rightExtent == rightExtent {
			return str
		}
	}
	fmt.Printf("Error: no string %s left extent=%d right extent=%d pos=%d\n",
		strings.Join(l.Symbols()[:l.Pos()], " "), leftExtent, rightExtent, l.Pos())
	panic("must not happen")
}

func Init(startSymbol string, I []byte) {
	set = newSet(I)
	startSym = startSymbol
}

func insert(bsr bsr) {
	if bsr.RightExtent() > set.rightExtent {
		set.rightExtent = bsr.RightExtent()
	}
	switch s := bsr.(type) {
	case BSR:
		set.slotEntries[s] = true
	case stringBSR:
		set.stringEntries[s] = true
	default:
		panic(fmt.Sprintf("Invalid type %T", bsr))
	}
}

// Alternate returns the index of the grammar rule alternate.
func (b BSR) Alternate() int {
	return b.Label.Alternate()
}

// GetNTChild returns the BSR of occurrence i of nt in s.
// GetNTChild fails if s has ambiguous subtrees of occurrence i of nt.
func (b BSR) GetNTChild(nt string, i int) BSR {
	bsrs := b.GetNTChildren(nt, i)
	if len(bsrs) != 1 {
		ambiguousSlots := []string{}
		for _, c := range bsrs {
			ambiguousSlots = append(ambiguousSlots, c.String())
		}
		fail(b, "%s is ambiguous in %s\n  %s", nt, b, strings.Join(ambiguousSlots, "\n  "))
	}
	return bsrs[0]
}

// GetNTChildI returns the BSR of NT symbol[i] in s.
// GetNTChildI fails if s has ambiguous subtrees of NT i.
func (b BSR) GetNTChildI(i int) BSR {
	bsrs := b.GetNTChildrenI(i)
	if len(bsrs) != 1 {
		fail(b, "NT %d is ambiguous in %s", i, b)
	}
	return bsrs[0]
}

// GetNTChild returns all the BSRs of occurrence i of nt in s
func (b BSR) GetNTChildren(nt string, i int) []BSR {
	// fmt.Printf("GetNTChild(%s,%d) %s\n", nt, i, b)
	positions := []int{}
	for j, s := range b.Label.Symbols() {
		if s == nt {
			positions = append(positions, j)
		}
	}
	if len(positions) == 0 {
		fail(b, "Error: %s has no NT %s", b, nt)
	}
	return b.GetNTChildrenI(positions[i])
}

// GetNTChildI returns all the BSRs of NT symbol[i] in s
func (b BSR) GetNTChildrenI(i int) []BSR {
	// fmt.Printf("bsr.GetNTChildI(%d) %s\n", i, b)
	if i >= len(b.Label.Symbols()) {
		fail(b, "Error: cannot get NT child %d of %s", i, b)
	}
	if len(b.Label.Symbols()) == 1 {
		return getNTSlot(b.Label.Symbols()[i], b.pivot, b.rightExtent)
	}
	if len(b.Label.Symbols()) == 2 {
		if i == 0 {
			return getNTSlot(b.Label.Symbols()[i], b.leftExtent, b.pivot)
		}
		return getNTSlot(b.Label.Symbols()[i], b.pivot, b.rightExtent)
	}
	idx := b.Label.Index()
	str := stringBSR{b.Label, b.leftExtent, b.pivot, b.rightExtent}
	for idx.Pos > i+1 && idx.Pos > 2 {
		idx.Pos--
		str = getString(slot.GetLabel(idx.NT, idx.Alt, idx.Pos), str.leftExtent, str.pivot)
		// fmt.Printf("  %s\n", str)
	}
	if i == 0 {
		return getNTSlot(b.Label.Symbols()[i], str.leftExtent, str.pivot)
	}
	return getNTSlot(b.Label.Symbols()[i], str.pivot, str.rightExtent)
}

func (b BSR) GetString() string {
	return string(set.I[b.LeftExtent():b.RightExtent()])
}

func (b BSR) Ignore() {
	// fmt.Printf("bsr.Ignore %s\n", b)
	delete(set.slotEntries, b)
	set.ignoredSlots[b] = true
}

func (s BSR) LeftExtent() int {
	return s.leftExtent
}

func (s BSR) RightExtent() int {
	return s.rightExtent
}

func (s BSR) Pivot() int {
	return s.pivot
}

func (s BSR) String() string {
	return fmt.Sprintf("%s,%d,%d,%d - %s", s.Label, s.leftExtent, s.pivot, s.rightExtent, set.I[s.LeftExtent():s.RightExtent()])
}

func (s stringBSR) LeftExtent() int {
	return s.leftExtent
}

func (s stringBSR) RightExtent() int {
	return s.rightExtent
}

func (s stringBSR) Pivot() int {
	return s.pivot
}

func (s stringBSR) Empty() bool {
	return s.leftExtent == s.pivot && s.pivot == s.rightExtent
}

func (s stringBSR) String() string {
	// fmt.Printf("bsr.stringBSR.stringBSR(): %s, %d, %d, %d\n",
	// 	s.Label.Symbols(), s.leftExtent, s.pivot, s.rightExtent)
	ss := s.Label.Symbols()[:s.Label.Pos()]
	str := strings.Join(ss, " ")
	return fmt.Sprintf("%s,%d,%d,%d - %s", str, s.leftExtent, s.pivot,
		s.rightExtent, set.I[s.LeftExtent():s.RightExtent()])
}

func Dump() {
	DumpSlots()
	DumpStrings()
}

func DumpSlots() {
	fmt.Printf("Slots (%d)\n", len(getSlots()))
	for _, s := range getSlots() {
		DumpSlot(s)
	}
}

func DumpSlot(s BSR) {
	fmt.Println(s)
}

func DumpStrings() {
	fmt.Printf("Strings(%d)\n", len(getStrings()))
	for _, s := range getStrings() {
		dumpString(s)
	}
}

func dumpString(s stringBSR) {
	fmt.Println(s)
}

func getSlots() (slots []BSR) {
	for s := range set.slotEntries {
		slots = append(slots, s)
	}
	sort.Slice(slots,
		func(i, j int) bool {
			return slots[i].Label < slots[j].Label
		})
	return
}

func getStrings() (strings []stringBSR) {
	for s := range set.stringEntries {
		strings = append(strings, s)
	}
	sort.Slice(strings,
		func(i, j int) bool {
			return strings[i].Label < strings[j].Label
		})
	return
}

func getNTSlot(nt string, leftExtent, rightExtent int) (bsrs []BSR) {
	for sl, _ := range set.slotEntries {
		if sl.Label.Head() == nt && sl.leftExtent == leftExtent && sl.rightExtent == rightExtent {
			bsrs = append(bsrs, sl)
		}
	}
	return
}

func fail(b BSR, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	line, col := getLineColumn(b.LeftExtent(), set.I)
	fmt.Printf("Error in BSR: %s at line %d col %d\n", msg, line, col)
	os.Exit(1)
}

func failf(format string, args ...interface{}) {
	fmt.Printf("Error in BSR: %s\n", fmt.Sprintf(format, args...))
	os.Exit(1)
}

func decodeRune(str []byte) (string, rune, int) {
	if len(str) == 0 {
		return "$", '$', 0
	}
	r, sz := utf8.DecodeRune(str)
	if r == utf8.RuneError {
		panic(fmt.Sprintf("Rune error: %s", str))
	}
	switch r {
	case '\t', ' ':
		return "space", r, sz
	case '\n':
		return "\\n", r, sz
	}
	return string(str[:sz]), r, sz
}

func getLineColumn(cI int, I []byte) (line, col int) {
	line, col = 1, 1
	for j := 0; j < cI; {
		_, r, sz := decodeRune(I[j:])
		switch r {
		case '\n':
			line++
			col = 1
		case '\t':
			col += 4
		default:
			col++
		}
		j += sz
	}
	return
}

