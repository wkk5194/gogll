
// Package slot is generated by gogll. Do not edit. 
//
//  Copyright 2019 Marius Ackerman
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
package slot

import(
	"bytes"
    "fmt"
    
    "github.com/goccmack/gogll/parser/symbols"
)

type Label int

const(
	GoGLL0R0 Label = iota
	GoGLL0R1
	GoGLL0R2
	LexAlternates0R0
	LexAlternates0R1
	LexAlternates1R0
	LexAlternates1R1
	LexAlternates1R2
	LexAlternates1R3
	LexBracket0R0
	LexBracket0R1
	LexBracket1R0
	LexBracket1R1
	LexBracket2R0
	LexBracket2R1
	LexBracket3R0
	LexBracket3R1
	LexGroup0R0
	LexGroup0R1
	LexGroup0R2
	LexGroup0R3
	LexOneOrMore0R0
	LexOneOrMore0R1
	LexOneOrMore0R2
	LexOneOrMore0R3
	LexOptional0R0
	LexOptional0R1
	LexOptional0R2
	LexOptional0R3
	LexRule0R0
	LexRule0R1
	LexRule0R2
	LexRule0R3
	LexRule0R4
	LexSymbol0R0
	LexSymbol0R1
	LexSymbol1R0
	LexSymbol1R1
	LexSymbol1R2
	LexSymbol2R0
	LexSymbol2R1
	LexSymbol3R0
	LexSymbol3R1
	LexSymbol4R0
	LexSymbol4R1
	LexSymbol4R2
	LexSymbol5R0
	LexSymbol5R1
	LexZeroOrMore0R0
	LexZeroOrMore0R1
	LexZeroOrMore0R2
	LexZeroOrMore0R3
	NT0R0
	NT0R1
	Package0R0
	Package0R1
	Package0R2
	RegExp0R0
	RegExp0R1
	RegExp1R0
	RegExp1R1
	RegExp1R2
	Rule0R0
	Rule0R1
	Rule1R0
	Rule1R1
	Rules0R0
	Rules0R1
	Rules1R0
	Rules1R1
	Rules1R2
	SyntaxAlternate0R0
	SyntaxAlternate0R1
	SyntaxAlternate1R0
	SyntaxAlternate1R1
	SyntaxAlternates0R0
	SyntaxAlternates0R1
	SyntaxAlternates1R0
	SyntaxAlternates1R1
	SyntaxAlternates1R2
	SyntaxAlternates1R3
	SyntaxRule0R0
	SyntaxRule0R1
	SyntaxRule0R2
	SyntaxRule0R3
	SyntaxRule0R4
	SyntaxSymbol0R0
	SyntaxSymbol0R1
	SyntaxSymbol1R0
	SyntaxSymbol1R1
	SyntaxSymbol2R0
	SyntaxSymbol2R1
	SyntaxSymbols0R0
	SyntaxSymbols0R1
	SyntaxSymbols1R0
	SyntaxSymbols1R1
	SyntaxSymbols1R2
	TokID0R0
	TokID0R1
	UnicodeClass0R0
	UnicodeClass0R1
	UnicodeClass1R0
	UnicodeClass1R1
	UnicodeClass2R0
	UnicodeClass2R1
	UnicodeClass3R0
	UnicodeClass3R1
	UnicodeClass4R0
	UnicodeClass4R1
)

type Slot struct {
	NT      symbols.NT
	Alt     int
	Pos     int
	Symbols symbols.Symbols
    Label 	Label
}

type Index struct {
	NT      symbols.NT
	Alt     int
	Pos     int
}

func GetAlternates(nt symbols.NT) []Label {
	alts, exist := alternates[nt]
	if !exist {
		panic(fmt.Sprintf("Invalid NT %s", nt))
	}
	return alts
}

func GetLabel(nt symbols.NT, alt, pos int) Label {
	l, exist := slotIndex[Index{nt,alt,pos}]
	if exist {
		return l
	}
	panic(fmt.Sprintf("Error: no slot label for NT=%s, alt=%d, pos=%d", nt, alt, pos))
}

func (l Label) EoR() bool {
	return l.Slot().EoR()
}

func (l Label) Head() symbols.NT {
	return l.Slot().NT
}

func (l Label) Index() Index {
	s := l.Slot()
	return Index{s.NT, s.Alt, s.Pos}
}

func (l Label) Alternate() int {
	return l.Slot().Alt
}

func (l Label) Pos() int {
	return l.Slot().Pos
}

func (l Label) Slot() *Slot {
	s, exist := slots[l]
	if !exist {
		panic(fmt.Sprintf("Invalid slot label %d", l))
	}
	return s
}

func (l Label) String() string {
	return l.Slot().String()
}

func (l Label) Symbols() symbols.Symbols {
	return l.Slot().Symbols
}

func (s *Slot) EoR() bool {
	return s.Pos >= len(s.Symbols)
}

func (s *Slot) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : ", s.NT)
	for i, sym := range s.Symbols {
		if i == s.Pos {
			fmt.Fprintf(buf, "∙")
		}
		fmt.Fprintf(buf, "%s ", sym)
	}
	if s.Pos >= len(s.Symbols) {
		fmt.Fprintf(buf, "∙")
	}
	return buf.String()
}

var slots = map[Label]*Slot{ 
    GoGLL0R0: {
        symbols.NT_GoGLL, 0, 0, 
        symbols.Symbols{  
            symbols.NT_Package, 
            symbols.NT_Rules,
        }, 
        GoGLL0R0, 
    },
    GoGLL0R1: {
        symbols.NT_GoGLL, 0, 1, 
        symbols.Symbols{  
            symbols.NT_Package, 
            symbols.NT_Rules,
        }, 
        GoGLL0R1, 
    },
    GoGLL0R2: {
        symbols.NT_GoGLL, 0, 2, 
        symbols.Symbols{  
            symbols.NT_Package, 
            symbols.NT_Rules,
        }, 
        GoGLL0R2, 
    },
    LexAlternates0R0: {
        symbols.NT_LexAlternates, 0, 0, 
        symbols.Symbols{  
            symbols.NT_RegExp,
        }, 
        LexAlternates0R0, 
    },
    LexAlternates0R1: {
        symbols.NT_LexAlternates, 0, 1, 
        symbols.Symbols{  
            symbols.NT_RegExp,
        }, 
        LexAlternates0R1, 
    },
    LexAlternates1R0: {
        symbols.NT_LexAlternates, 1, 0, 
        symbols.Symbols{  
            symbols.NT_RegExp, 
            symbols.T_23, 
            symbols.NT_LexAlternates,
        }, 
        LexAlternates1R0, 
    },
    LexAlternates1R1: {
        symbols.NT_LexAlternates, 1, 1, 
        symbols.Symbols{  
            symbols.NT_RegExp, 
            symbols.T_23, 
            symbols.NT_LexAlternates,
        }, 
        LexAlternates1R1, 
    },
    LexAlternates1R2: {
        symbols.NT_LexAlternates, 1, 2, 
        symbols.Symbols{  
            symbols.NT_RegExp, 
            symbols.T_23, 
            symbols.NT_LexAlternates,
        }, 
        LexAlternates1R2, 
    },
    LexAlternates1R3: {
        symbols.NT_LexAlternates, 1, 3, 
        symbols.Symbols{  
            symbols.NT_RegExp, 
            symbols.T_23, 
            symbols.NT_LexAlternates,
        }, 
        LexAlternates1R3, 
    },
    LexBracket0R0: {
        symbols.NT_LexBracket, 0, 0, 
        symbols.Symbols{  
            symbols.NT_LexGroup,
        }, 
        LexBracket0R0, 
    },
    LexBracket0R1: {
        symbols.NT_LexBracket, 0, 1, 
        symbols.Symbols{  
            symbols.NT_LexGroup,
        }, 
        LexBracket0R1, 
    },
    LexBracket1R0: {
        symbols.NT_LexBracket, 1, 0, 
        symbols.Symbols{  
            symbols.NT_LexOptional,
        }, 
        LexBracket1R0, 
    },
    LexBracket1R1: {
        symbols.NT_LexBracket, 1, 1, 
        symbols.Symbols{  
            symbols.NT_LexOptional,
        }, 
        LexBracket1R1, 
    },
    LexBracket2R0: {
        symbols.NT_LexBracket, 2, 0, 
        symbols.Symbols{  
            symbols.NT_LexZeroOrMore,
        }, 
        LexBracket2R0, 
    },
    LexBracket2R1: {
        symbols.NT_LexBracket, 2, 1, 
        symbols.Symbols{  
            symbols.NT_LexZeroOrMore,
        }, 
        LexBracket2R1, 
    },
    LexBracket3R0: {
        symbols.NT_LexBracket, 3, 0, 
        symbols.Symbols{  
            symbols.NT_LexOneOrMore,
        }, 
        LexBracket3R0, 
    },
    LexBracket3R1: {
        symbols.NT_LexBracket, 3, 1, 
        symbols.Symbols{  
            symbols.NT_LexOneOrMore,
        }, 
        LexBracket3R1, 
    },
    LexGroup0R0: {
        symbols.NT_LexGroup, 0, 0, 
        symbols.Symbols{  
            symbols.T_0, 
            symbols.NT_LexAlternates, 
            symbols.T_1,
        }, 
        LexGroup0R0, 
    },
    LexGroup0R1: {
        symbols.NT_LexGroup, 0, 1, 
        symbols.Symbols{  
            symbols.T_0, 
            symbols.NT_LexAlternates, 
            symbols.T_1,
        }, 
        LexGroup0R1, 
    },
    LexGroup0R2: {
        symbols.NT_LexGroup, 0, 2, 
        symbols.Symbols{  
            symbols.T_0, 
            symbols.NT_LexAlternates, 
            symbols.T_1,
        }, 
        LexGroup0R2, 
    },
    LexGroup0R3: {
        symbols.NT_LexGroup, 0, 3, 
        symbols.Symbols{  
            symbols.T_0, 
            symbols.NT_LexAlternates, 
            symbols.T_1,
        }, 
        LexGroup0R3, 
    },
    LexOneOrMore0R0: {
        symbols.NT_LexOneOrMore, 0, 0, 
        symbols.Symbols{  
            symbols.T_5, 
            symbols.NT_LexAlternates, 
            symbols.T_6,
        }, 
        LexOneOrMore0R0, 
    },
    LexOneOrMore0R1: {
        symbols.NT_LexOneOrMore, 0, 1, 
        symbols.Symbols{  
            symbols.T_5, 
            symbols.NT_LexAlternates, 
            symbols.T_6,
        }, 
        LexOneOrMore0R1, 
    },
    LexOneOrMore0R2: {
        symbols.NT_LexOneOrMore, 0, 2, 
        symbols.Symbols{  
            symbols.T_5, 
            symbols.NT_LexAlternates, 
            symbols.T_6,
        }, 
        LexOneOrMore0R2, 
    },
    LexOneOrMore0R3: {
        symbols.NT_LexOneOrMore, 0, 3, 
        symbols.Symbols{  
            symbols.T_5, 
            symbols.NT_LexAlternates, 
            symbols.T_6,
        }, 
        LexOneOrMore0R3, 
    },
    LexOptional0R0: {
        symbols.NT_LexOptional, 0, 0, 
        symbols.Symbols{  
            symbols.T_7, 
            symbols.NT_LexAlternates, 
            symbols.T_8,
        }, 
        LexOptional0R0, 
    },
    LexOptional0R1: {
        symbols.NT_LexOptional, 0, 1, 
        symbols.Symbols{  
            symbols.T_7, 
            symbols.NT_LexAlternates, 
            symbols.T_8,
        }, 
        LexOptional0R1, 
    },
    LexOptional0R2: {
        symbols.NT_LexOptional, 0, 2, 
        symbols.Symbols{  
            symbols.T_7, 
            symbols.NT_LexAlternates, 
            symbols.T_8,
        }, 
        LexOptional0R2, 
    },
    LexOptional0R3: {
        symbols.NT_LexOptional, 0, 3, 
        symbols.Symbols{  
            symbols.T_7, 
            symbols.NT_LexAlternates, 
            symbols.T_8,
        }, 
        LexOptional0R3, 
    },
    LexRule0R0: {
        symbols.NT_LexRule, 0, 0, 
        symbols.Symbols{  
            symbols.NT_TokID, 
            symbols.T_3, 
            symbols.NT_RegExp, 
            symbols.T_4,
        }, 
        LexRule0R0, 
    },
    LexRule0R1: {
        symbols.NT_LexRule, 0, 1, 
        symbols.Symbols{  
            symbols.NT_TokID, 
            symbols.T_3, 
            symbols.NT_RegExp, 
            symbols.T_4,
        }, 
        LexRule0R1, 
    },
    LexRule0R2: {
        symbols.NT_LexRule, 0, 2, 
        symbols.Symbols{  
            symbols.NT_TokID, 
            symbols.T_3, 
            symbols.NT_RegExp, 
            symbols.T_4,
        }, 
        LexRule0R2, 
    },
    LexRule0R3: {
        symbols.NT_LexRule, 0, 3, 
        symbols.Symbols{  
            symbols.NT_TokID, 
            symbols.T_3, 
            symbols.NT_RegExp, 
            symbols.T_4,
        }, 
        LexRule0R3, 
    },
    LexRule0R4: {
        symbols.NT_LexRule, 0, 4, 
        symbols.Symbols{  
            symbols.NT_TokID, 
            symbols.T_3, 
            symbols.NT_RegExp, 
            symbols.T_4,
        }, 
        LexRule0R4, 
    },
    LexSymbol0R0: {
        symbols.NT_LexSymbol, 0, 0, 
        symbols.Symbols{  
            symbols.T_2,
        }, 
        LexSymbol0R0, 
    },
    LexSymbol0R1: {
        symbols.NT_LexSymbol, 0, 1, 
        symbols.Symbols{  
            symbols.T_2,
        }, 
        LexSymbol0R1, 
    },
    LexSymbol1R0: {
        symbols.NT_LexSymbol, 1, 0, 
        symbols.Symbols{  
            symbols.T_9, 
            symbols.T_19,
        }, 
        LexSymbol1R0, 
    },
    LexSymbol1R1: {
        symbols.NT_LexSymbol, 1, 1, 
        symbols.Symbols{  
            symbols.T_9, 
            symbols.T_19,
        }, 
        LexSymbol1R1, 
    },
    LexSymbol1R2: {
        symbols.NT_LexSymbol, 1, 2, 
        symbols.Symbols{  
            symbols.T_9, 
            symbols.T_19,
        }, 
        LexSymbol1R2, 
    },
    LexSymbol2R0: {
        symbols.NT_LexSymbol, 2, 0, 
        symbols.Symbols{  
            symbols.T_10,
        }, 
        LexSymbol2R0, 
    },
    LexSymbol2R1: {
        symbols.NT_LexSymbol, 2, 1, 
        symbols.Symbols{  
            symbols.T_10,
        }, 
        LexSymbol2R1, 
    },
    LexSymbol3R0: {
        symbols.NT_LexSymbol, 3, 0, 
        symbols.Symbols{  
            symbols.NT_LexBracket,
        }, 
        LexSymbol3R0, 
    },
    LexSymbol3R1: {
        symbols.NT_LexSymbol, 3, 1, 
        symbols.Symbols{  
            symbols.NT_LexBracket,
        }, 
        LexSymbol3R1, 
    },
    LexSymbol4R0: {
        symbols.NT_LexSymbol, 4, 0, 
        symbols.Symbols{  
            symbols.T_14, 
            symbols.T_19,
        }, 
        LexSymbol4R0, 
    },
    LexSymbol4R1: {
        symbols.NT_LexSymbol, 4, 1, 
        symbols.Symbols{  
            symbols.T_14, 
            symbols.T_19,
        }, 
        LexSymbol4R1, 
    },
    LexSymbol4R2: {
        symbols.NT_LexSymbol, 4, 2, 
        symbols.Symbols{  
            symbols.T_14, 
            symbols.T_19,
        }, 
        LexSymbol4R2, 
    },
    LexSymbol5R0: {
        symbols.NT_LexSymbol, 5, 0, 
        symbols.Symbols{  
            symbols.NT_UnicodeClass,
        }, 
        LexSymbol5R0, 
    },
    LexSymbol5R1: {
        symbols.NT_LexSymbol, 5, 1, 
        symbols.Symbols{  
            symbols.NT_UnicodeClass,
        }, 
        LexSymbol5R1, 
    },
    LexZeroOrMore0R0: {
        symbols.NT_LexZeroOrMore, 0, 0, 
        symbols.Symbols{  
            symbols.T_22, 
            symbols.NT_LexAlternates, 
            symbols.T_24,
        }, 
        LexZeroOrMore0R0, 
    },
    LexZeroOrMore0R1: {
        symbols.NT_LexZeroOrMore, 0, 1, 
        symbols.Symbols{  
            symbols.T_22, 
            symbols.NT_LexAlternates, 
            symbols.T_24,
        }, 
        LexZeroOrMore0R1, 
    },
    LexZeroOrMore0R2: {
        symbols.NT_LexZeroOrMore, 0, 2, 
        symbols.Symbols{  
            symbols.T_22, 
            symbols.NT_LexAlternates, 
            symbols.T_24,
        }, 
        LexZeroOrMore0R2, 
    },
    LexZeroOrMore0R3: {
        symbols.NT_LexZeroOrMore, 0, 3, 
        symbols.Symbols{  
            symbols.T_22, 
            symbols.NT_LexAlternates, 
            symbols.T_24,
        }, 
        LexZeroOrMore0R3, 
    },
    NT0R0: {
        symbols.NT_NT, 0, 0, 
        symbols.Symbols{  
            symbols.T_15,
        }, 
        NT0R0, 
    },
    NT0R1: {
        symbols.NT_NT, 0, 1, 
        symbols.Symbols{  
            symbols.T_15,
        }, 
        NT0R1, 
    },
    Package0R0: {
        symbols.NT_Package, 0, 0, 
        symbols.Symbols{  
            symbols.T_17, 
            symbols.T_19,
        }, 
        Package0R0, 
    },
    Package0R1: {
        symbols.NT_Package, 0, 1, 
        symbols.Symbols{  
            symbols.T_17, 
            symbols.T_19,
        }, 
        Package0R1, 
    },
    Package0R2: {
        symbols.NT_Package, 0, 2, 
        symbols.Symbols{  
            symbols.T_17, 
            symbols.T_19,
        }, 
        Package0R2, 
    },
    RegExp0R0: {
        symbols.NT_RegExp, 0, 0, 
        symbols.Symbols{  
            symbols.NT_LexSymbol,
        }, 
        RegExp0R0, 
    },
    RegExp0R1: {
        symbols.NT_RegExp, 0, 1, 
        symbols.Symbols{  
            symbols.NT_LexSymbol,
        }, 
        RegExp0R1, 
    },
    RegExp1R0: {
        symbols.NT_RegExp, 1, 0, 
        symbols.Symbols{  
            symbols.NT_LexSymbol, 
            symbols.NT_RegExp,
        }, 
        RegExp1R0, 
    },
    RegExp1R1: {
        symbols.NT_RegExp, 1, 1, 
        symbols.Symbols{  
            symbols.NT_LexSymbol, 
            symbols.NT_RegExp,
        }, 
        RegExp1R1, 
    },
    RegExp1R2: {
        symbols.NT_RegExp, 1, 2, 
        symbols.Symbols{  
            symbols.NT_LexSymbol, 
            symbols.NT_RegExp,
        }, 
        RegExp1R2, 
    },
    Rule0R0: {
        symbols.NT_Rule, 0, 0, 
        symbols.Symbols{  
            symbols.NT_LexRule,
        }, 
        Rule0R0, 
    },
    Rule0R1: {
        symbols.NT_Rule, 0, 1, 
        symbols.Symbols{  
            symbols.NT_LexRule,
        }, 
        Rule0R1, 
    },
    Rule1R0: {
        symbols.NT_Rule, 1, 0, 
        symbols.Symbols{  
            symbols.NT_SyntaxRule,
        }, 
        Rule1R0, 
    },
    Rule1R1: {
        symbols.NT_Rule, 1, 1, 
        symbols.Symbols{  
            symbols.NT_SyntaxRule,
        }, 
        Rule1R1, 
    },
    Rules0R0: {
        symbols.NT_Rules, 0, 0, 
        symbols.Symbols{  
            symbols.NT_Rule,
        }, 
        Rules0R0, 
    },
    Rules0R1: {
        symbols.NT_Rules, 0, 1, 
        symbols.Symbols{  
            symbols.NT_Rule,
        }, 
        Rules0R1, 
    },
    Rules1R0: {
        symbols.NT_Rules, 1, 0, 
        symbols.Symbols{  
            symbols.NT_Rule, 
            symbols.NT_Rules,
        }, 
        Rules1R0, 
    },
    Rules1R1: {
        symbols.NT_Rules, 1, 1, 
        symbols.Symbols{  
            symbols.NT_Rule, 
            symbols.NT_Rules,
        }, 
        Rules1R1, 
    },
    Rules1R2: {
        symbols.NT_Rules, 1, 2, 
        symbols.Symbols{  
            symbols.NT_Rule, 
            symbols.NT_Rules,
        }, 
        Rules1R2, 
    },
    SyntaxAlternate0R0: {
        symbols.NT_SyntaxAlternate, 0, 0, 
        symbols.Symbols{  
            symbols.NT_SyntaxSymbols,
        }, 
        SyntaxAlternate0R0, 
    },
    SyntaxAlternate0R1: {
        symbols.NT_SyntaxAlternate, 0, 1, 
        symbols.Symbols{  
            symbols.NT_SyntaxSymbols,
        }, 
        SyntaxAlternate0R1, 
    },
    SyntaxAlternate1R0: {
        symbols.NT_SyntaxAlternate, 1, 0, 
        symbols.Symbols{  
            symbols.T_11,
        }, 
        SyntaxAlternate1R0, 
    },
    SyntaxAlternate1R1: {
        symbols.NT_SyntaxAlternate, 1, 1, 
        symbols.Symbols{  
            symbols.T_11,
        }, 
        SyntaxAlternate1R1, 
    },
    SyntaxAlternates0R0: {
        symbols.NT_SyntaxAlternates, 0, 0, 
        symbols.Symbols{  
            symbols.NT_SyntaxAlternate,
        }, 
        SyntaxAlternates0R0, 
    },
    SyntaxAlternates0R1: {
        symbols.NT_SyntaxAlternates, 0, 1, 
        symbols.Symbols{  
            symbols.NT_SyntaxAlternate,
        }, 
        SyntaxAlternates0R1, 
    },
    SyntaxAlternates1R0: {
        symbols.NT_SyntaxAlternates, 1, 0, 
        symbols.Symbols{  
            symbols.NT_SyntaxAlternate, 
            symbols.T_23, 
            symbols.NT_SyntaxAlternates,
        }, 
        SyntaxAlternates1R0, 
    },
    SyntaxAlternates1R1: {
        symbols.NT_SyntaxAlternates, 1, 1, 
        symbols.Symbols{  
            symbols.NT_SyntaxAlternate, 
            symbols.T_23, 
            symbols.NT_SyntaxAlternates,
        }, 
        SyntaxAlternates1R1, 
    },
    SyntaxAlternates1R2: {
        symbols.NT_SyntaxAlternates, 1, 2, 
        symbols.Symbols{  
            symbols.NT_SyntaxAlternate, 
            symbols.T_23, 
            symbols.NT_SyntaxAlternates,
        }, 
        SyntaxAlternates1R2, 
    },
    SyntaxAlternates1R3: {
        symbols.NT_SyntaxAlternates, 1, 3, 
        symbols.Symbols{  
            symbols.NT_SyntaxAlternate, 
            symbols.T_23, 
            symbols.NT_SyntaxAlternates,
        }, 
        SyntaxAlternates1R3, 
    },
    SyntaxRule0R0: {
        symbols.NT_SyntaxRule, 0, 0, 
        symbols.Symbols{  
            symbols.NT_NT, 
            symbols.T_3, 
            symbols.NT_SyntaxAlternates, 
            symbols.T_4,
        }, 
        SyntaxRule0R0, 
    },
    SyntaxRule0R1: {
        symbols.NT_SyntaxRule, 0, 1, 
        symbols.Symbols{  
            symbols.NT_NT, 
            symbols.T_3, 
            symbols.NT_SyntaxAlternates, 
            symbols.T_4,
        }, 
        SyntaxRule0R1, 
    },
    SyntaxRule0R2: {
        symbols.NT_SyntaxRule, 0, 2, 
        symbols.Symbols{  
            symbols.NT_NT, 
            symbols.T_3, 
            symbols.NT_SyntaxAlternates, 
            symbols.T_4,
        }, 
        SyntaxRule0R2, 
    },
    SyntaxRule0R3: {
        symbols.NT_SyntaxRule, 0, 3, 
        symbols.Symbols{  
            symbols.NT_NT, 
            symbols.T_3, 
            symbols.NT_SyntaxAlternates, 
            symbols.T_4,
        }, 
        SyntaxRule0R3, 
    },
    SyntaxRule0R4: {
        symbols.NT_SyntaxRule, 0, 4, 
        symbols.Symbols{  
            symbols.NT_NT, 
            symbols.T_3, 
            symbols.NT_SyntaxAlternates, 
            symbols.T_4,
        }, 
        SyntaxRule0R4, 
    },
    SyntaxSymbol0R0: {
        symbols.NT_SyntaxSymbol, 0, 0, 
        symbols.Symbols{  
            symbols.NT_NT,
        }, 
        SyntaxSymbol0R0, 
    },
    SyntaxSymbol0R1: {
        symbols.NT_SyntaxSymbol, 0, 1, 
        symbols.Symbols{  
            symbols.NT_NT,
        }, 
        SyntaxSymbol0R1, 
    },
    SyntaxSymbol1R0: {
        symbols.NT_SyntaxSymbol, 1, 0, 
        symbols.Symbols{  
            symbols.NT_TokID,
        }, 
        SyntaxSymbol1R0, 
    },
    SyntaxSymbol1R1: {
        symbols.NT_SyntaxSymbol, 1, 1, 
        symbols.Symbols{  
            symbols.NT_TokID,
        }, 
        SyntaxSymbol1R1, 
    },
    SyntaxSymbol2R0: {
        symbols.NT_SyntaxSymbol, 2, 0, 
        symbols.Symbols{  
            symbols.T_19,
        }, 
        SyntaxSymbol2R0, 
    },
    SyntaxSymbol2R1: {
        symbols.NT_SyntaxSymbol, 2, 1, 
        symbols.Symbols{  
            symbols.T_19,
        }, 
        SyntaxSymbol2R1, 
    },
    SyntaxSymbols0R0: {
        symbols.NT_SyntaxSymbols, 0, 0, 
        symbols.Symbols{  
            symbols.NT_SyntaxSymbol,
        }, 
        SyntaxSymbols0R0, 
    },
    SyntaxSymbols0R1: {
        symbols.NT_SyntaxSymbols, 0, 1, 
        symbols.Symbols{  
            symbols.NT_SyntaxSymbol,
        }, 
        SyntaxSymbols0R1, 
    },
    SyntaxSymbols1R0: {
        symbols.NT_SyntaxSymbols, 1, 0, 
        symbols.Symbols{  
            symbols.NT_SyntaxSymbol, 
            symbols.NT_SyntaxSymbols,
        }, 
        SyntaxSymbols1R0, 
    },
    SyntaxSymbols1R1: {
        symbols.NT_SyntaxSymbols, 1, 1, 
        symbols.Symbols{  
            symbols.NT_SyntaxSymbol, 
            symbols.NT_SyntaxSymbols,
        }, 
        SyntaxSymbols1R1, 
    },
    SyntaxSymbols1R2: {
        symbols.NT_SyntaxSymbols, 1, 2, 
        symbols.Symbols{  
            symbols.NT_SyntaxSymbol, 
            symbols.NT_SyntaxSymbols,
        }, 
        SyntaxSymbols1R2, 
    },
    TokID0R0: {
        symbols.NT_TokID, 0, 0, 
        symbols.Symbols{  
            symbols.T_20,
        }, 
        TokID0R0, 
    },
    TokID0R1: {
        symbols.NT_TokID, 0, 1, 
        symbols.Symbols{  
            symbols.T_20,
        }, 
        TokID0R1, 
    },
    UnicodeClass0R0: {
        symbols.NT_UnicodeClass, 0, 0, 
        symbols.Symbols{  
            symbols.T_12,
        }, 
        UnicodeClass0R0, 
    },
    UnicodeClass0R1: {
        symbols.NT_UnicodeClass, 0, 1, 
        symbols.Symbols{  
            symbols.T_12,
        }, 
        UnicodeClass0R1, 
    },
    UnicodeClass1R0: {
        symbols.NT_UnicodeClass, 1, 0, 
        symbols.Symbols{  
            symbols.T_21,
        }, 
        UnicodeClass1R0, 
    },
    UnicodeClass1R1: {
        symbols.NT_UnicodeClass, 1, 1, 
        symbols.Symbols{  
            symbols.T_21,
        }, 
        UnicodeClass1R1, 
    },
    UnicodeClass2R0: {
        symbols.NT_UnicodeClass, 2, 0, 
        symbols.Symbols{  
            symbols.T_13,
        }, 
        UnicodeClass2R0, 
    },
    UnicodeClass2R1: {
        symbols.NT_UnicodeClass, 2, 1, 
        symbols.Symbols{  
            symbols.T_13,
        }, 
        UnicodeClass2R1, 
    },
    UnicodeClass3R0: {
        symbols.NT_UnicodeClass, 3, 0, 
        symbols.Symbols{  
            symbols.T_16,
        }, 
        UnicodeClass3R0, 
    },
    UnicodeClass3R1: {
        symbols.NT_UnicodeClass, 3, 1, 
        symbols.Symbols{  
            symbols.T_16,
        }, 
        UnicodeClass3R1, 
    },
    UnicodeClass4R0: {
        symbols.NT_UnicodeClass, 4, 0, 
        symbols.Symbols{  
            symbols.T_18,
        }, 
        UnicodeClass4R0, 
    },
    UnicodeClass4R1: {
        symbols.NT_UnicodeClass, 4, 1, 
        symbols.Symbols{  
            symbols.T_18,
        }, 
        UnicodeClass4R1, 
    },
}

var slotIndex = map[Index]Label { 
	Index{ symbols.NT_GoGLL,0,0 }: GoGLL0R0,
	Index{ symbols.NT_GoGLL,0,1 }: GoGLL0R1,
	Index{ symbols.NT_GoGLL,0,2 }: GoGLL0R2,
	Index{ symbols.NT_LexAlternates,0,0 }: LexAlternates0R0,
	Index{ symbols.NT_LexAlternates,0,1 }: LexAlternates0R1,
	Index{ symbols.NT_LexAlternates,1,0 }: LexAlternates1R0,
	Index{ symbols.NT_LexAlternates,1,1 }: LexAlternates1R1,
	Index{ symbols.NT_LexAlternates,1,2 }: LexAlternates1R2,
	Index{ symbols.NT_LexAlternates,1,3 }: LexAlternates1R3,
	Index{ symbols.NT_LexBracket,0,0 }: LexBracket0R0,
	Index{ symbols.NT_LexBracket,0,1 }: LexBracket0R1,
	Index{ symbols.NT_LexBracket,1,0 }: LexBracket1R0,
	Index{ symbols.NT_LexBracket,1,1 }: LexBracket1R1,
	Index{ symbols.NT_LexBracket,2,0 }: LexBracket2R0,
	Index{ symbols.NT_LexBracket,2,1 }: LexBracket2R1,
	Index{ symbols.NT_LexBracket,3,0 }: LexBracket3R0,
	Index{ symbols.NT_LexBracket,3,1 }: LexBracket3R1,
	Index{ symbols.NT_LexGroup,0,0 }: LexGroup0R0,
	Index{ symbols.NT_LexGroup,0,1 }: LexGroup0R1,
	Index{ symbols.NT_LexGroup,0,2 }: LexGroup0R2,
	Index{ symbols.NT_LexGroup,0,3 }: LexGroup0R3,
	Index{ symbols.NT_LexOneOrMore,0,0 }: LexOneOrMore0R0,
	Index{ symbols.NT_LexOneOrMore,0,1 }: LexOneOrMore0R1,
	Index{ symbols.NT_LexOneOrMore,0,2 }: LexOneOrMore0R2,
	Index{ symbols.NT_LexOneOrMore,0,3 }: LexOneOrMore0R3,
	Index{ symbols.NT_LexOptional,0,0 }: LexOptional0R0,
	Index{ symbols.NT_LexOptional,0,1 }: LexOptional0R1,
	Index{ symbols.NT_LexOptional,0,2 }: LexOptional0R2,
	Index{ symbols.NT_LexOptional,0,3 }: LexOptional0R3,
	Index{ symbols.NT_LexRule,0,0 }: LexRule0R0,
	Index{ symbols.NT_LexRule,0,1 }: LexRule0R1,
	Index{ symbols.NT_LexRule,0,2 }: LexRule0R2,
	Index{ symbols.NT_LexRule,0,3 }: LexRule0R3,
	Index{ symbols.NT_LexRule,0,4 }: LexRule0R4,
	Index{ symbols.NT_LexSymbol,0,0 }: LexSymbol0R0,
	Index{ symbols.NT_LexSymbol,0,1 }: LexSymbol0R1,
	Index{ symbols.NT_LexSymbol,1,0 }: LexSymbol1R0,
	Index{ symbols.NT_LexSymbol,1,1 }: LexSymbol1R1,
	Index{ symbols.NT_LexSymbol,1,2 }: LexSymbol1R2,
	Index{ symbols.NT_LexSymbol,2,0 }: LexSymbol2R0,
	Index{ symbols.NT_LexSymbol,2,1 }: LexSymbol2R1,
	Index{ symbols.NT_LexSymbol,3,0 }: LexSymbol3R0,
	Index{ symbols.NT_LexSymbol,3,1 }: LexSymbol3R1,
	Index{ symbols.NT_LexSymbol,4,0 }: LexSymbol4R0,
	Index{ symbols.NT_LexSymbol,4,1 }: LexSymbol4R1,
	Index{ symbols.NT_LexSymbol,4,2 }: LexSymbol4R2,
	Index{ symbols.NT_LexSymbol,5,0 }: LexSymbol5R0,
	Index{ symbols.NT_LexSymbol,5,1 }: LexSymbol5R1,
	Index{ symbols.NT_LexZeroOrMore,0,0 }: LexZeroOrMore0R0,
	Index{ symbols.NT_LexZeroOrMore,0,1 }: LexZeroOrMore0R1,
	Index{ symbols.NT_LexZeroOrMore,0,2 }: LexZeroOrMore0R2,
	Index{ symbols.NT_LexZeroOrMore,0,3 }: LexZeroOrMore0R3,
	Index{ symbols.NT_NT,0,0 }: NT0R0,
	Index{ symbols.NT_NT,0,1 }: NT0R1,
	Index{ symbols.NT_Package,0,0 }: Package0R0,
	Index{ symbols.NT_Package,0,1 }: Package0R1,
	Index{ symbols.NT_Package,0,2 }: Package0R2,
	Index{ symbols.NT_RegExp,0,0 }: RegExp0R0,
	Index{ symbols.NT_RegExp,0,1 }: RegExp0R1,
	Index{ symbols.NT_RegExp,1,0 }: RegExp1R0,
	Index{ symbols.NT_RegExp,1,1 }: RegExp1R1,
	Index{ symbols.NT_RegExp,1,2 }: RegExp1R2,
	Index{ symbols.NT_Rule,0,0 }: Rule0R0,
	Index{ symbols.NT_Rule,0,1 }: Rule0R1,
	Index{ symbols.NT_Rule,1,0 }: Rule1R0,
	Index{ symbols.NT_Rule,1,1 }: Rule1R1,
	Index{ symbols.NT_Rules,0,0 }: Rules0R0,
	Index{ symbols.NT_Rules,0,1 }: Rules0R1,
	Index{ symbols.NT_Rules,1,0 }: Rules1R0,
	Index{ symbols.NT_Rules,1,1 }: Rules1R1,
	Index{ symbols.NT_Rules,1,2 }: Rules1R2,
	Index{ symbols.NT_SyntaxAlternate,0,0 }: SyntaxAlternate0R0,
	Index{ symbols.NT_SyntaxAlternate,0,1 }: SyntaxAlternate0R1,
	Index{ symbols.NT_SyntaxAlternate,1,0 }: SyntaxAlternate1R0,
	Index{ symbols.NT_SyntaxAlternate,1,1 }: SyntaxAlternate1R1,
	Index{ symbols.NT_SyntaxAlternates,0,0 }: SyntaxAlternates0R0,
	Index{ symbols.NT_SyntaxAlternates,0,1 }: SyntaxAlternates0R1,
	Index{ symbols.NT_SyntaxAlternates,1,0 }: SyntaxAlternates1R0,
	Index{ symbols.NT_SyntaxAlternates,1,1 }: SyntaxAlternates1R1,
	Index{ symbols.NT_SyntaxAlternates,1,2 }: SyntaxAlternates1R2,
	Index{ symbols.NT_SyntaxAlternates,1,3 }: SyntaxAlternates1R3,
	Index{ symbols.NT_SyntaxRule,0,0 }: SyntaxRule0R0,
	Index{ symbols.NT_SyntaxRule,0,1 }: SyntaxRule0R1,
	Index{ symbols.NT_SyntaxRule,0,2 }: SyntaxRule0R2,
	Index{ symbols.NT_SyntaxRule,0,3 }: SyntaxRule0R3,
	Index{ symbols.NT_SyntaxRule,0,4 }: SyntaxRule0R4,
	Index{ symbols.NT_SyntaxSymbol,0,0 }: SyntaxSymbol0R0,
	Index{ symbols.NT_SyntaxSymbol,0,1 }: SyntaxSymbol0R1,
	Index{ symbols.NT_SyntaxSymbol,1,0 }: SyntaxSymbol1R0,
	Index{ symbols.NT_SyntaxSymbol,1,1 }: SyntaxSymbol1R1,
	Index{ symbols.NT_SyntaxSymbol,2,0 }: SyntaxSymbol2R0,
	Index{ symbols.NT_SyntaxSymbol,2,1 }: SyntaxSymbol2R1,
	Index{ symbols.NT_SyntaxSymbols,0,0 }: SyntaxSymbols0R0,
	Index{ symbols.NT_SyntaxSymbols,0,1 }: SyntaxSymbols0R1,
	Index{ symbols.NT_SyntaxSymbols,1,0 }: SyntaxSymbols1R0,
	Index{ symbols.NT_SyntaxSymbols,1,1 }: SyntaxSymbols1R1,
	Index{ symbols.NT_SyntaxSymbols,1,2 }: SyntaxSymbols1R2,
	Index{ symbols.NT_TokID,0,0 }: TokID0R0,
	Index{ symbols.NT_TokID,0,1 }: TokID0R1,
	Index{ symbols.NT_UnicodeClass,0,0 }: UnicodeClass0R0,
	Index{ symbols.NT_UnicodeClass,0,1 }: UnicodeClass0R1,
	Index{ symbols.NT_UnicodeClass,1,0 }: UnicodeClass1R0,
	Index{ symbols.NT_UnicodeClass,1,1 }: UnicodeClass1R1,
	Index{ symbols.NT_UnicodeClass,2,0 }: UnicodeClass2R0,
	Index{ symbols.NT_UnicodeClass,2,1 }: UnicodeClass2R1,
	Index{ symbols.NT_UnicodeClass,3,0 }: UnicodeClass3R0,
	Index{ symbols.NT_UnicodeClass,3,1 }: UnicodeClass3R1,
	Index{ symbols.NT_UnicodeClass,4,0 }: UnicodeClass4R0,
	Index{ symbols.NT_UnicodeClass,4,1 }: UnicodeClass4R1,
}

var alternates = map[symbols.NT][]Label{ 
	symbols.NT_GoGLL:[]Label{ GoGLL0R0 },
	symbols.NT_Package:[]Label{ Package0R0 },
	symbols.NT_Rules:[]Label{ Rules0R0,Rules1R0 },
	symbols.NT_Rule:[]Label{ Rule0R0,Rule1R0 },
	symbols.NT_LexRule:[]Label{ LexRule0R0 },
	symbols.NT_RegExp:[]Label{ RegExp0R0,RegExp1R0 },
	symbols.NT_LexSymbol:[]Label{ LexSymbol0R0,LexSymbol1R0,LexSymbol2R0,LexSymbol3R0,LexSymbol4R0,LexSymbol5R0 },
	symbols.NT_LexBracket:[]Label{ LexBracket0R0,LexBracket1R0,LexBracket2R0,LexBracket3R0 },
	symbols.NT_LexGroup:[]Label{ LexGroup0R0 },
	symbols.NT_LexOptional:[]Label{ LexOptional0R0 },
	symbols.NT_LexZeroOrMore:[]Label{ LexZeroOrMore0R0 },
	symbols.NT_LexOneOrMore:[]Label{ LexOneOrMore0R0 },
	symbols.NT_LexAlternates:[]Label{ LexAlternates0R0,LexAlternates1R0 },
	symbols.NT_UnicodeClass:[]Label{ UnicodeClass0R0,UnicodeClass1R0,UnicodeClass2R0,UnicodeClass3R0,UnicodeClass4R0 },
	symbols.NT_SyntaxRule:[]Label{ SyntaxRule0R0 },
	symbols.NT_NT:[]Label{ NT0R0 },
	symbols.NT_SyntaxAlternates:[]Label{ SyntaxAlternates0R0,SyntaxAlternates1R0 },
	symbols.NT_SyntaxAlternate:[]Label{ SyntaxAlternate0R0,SyntaxAlternate1R0 },
	symbols.NT_SyntaxSymbols:[]Label{ SyntaxSymbols0R0,SyntaxSymbols1R0 },
	symbols.NT_SyntaxSymbol:[]Label{ SyntaxSymbol0R0,SyntaxSymbol1R0,SyntaxSymbol2R0 },
	symbols.NT_TokID:[]Label{ TokID0R0 },
}

