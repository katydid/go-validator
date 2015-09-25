package lexer

import (

	// "fmt"
	// "github.com/katydid/katydid/expr/util"

	"github.com/katydid/katydid/expr/token"
	"io/ioutil"
	"unicode/utf8"
)

const (
	NoState    = -1
	NumStates  = 267
	NumSymbols = 194
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (this *Lexer) Scan() (tok *token.Token) {

	// fmt.Printf("Lexer.Scan() pos=%d\n", this.pos)

	tok = new(token.Token)
	if this.pos >= len(this.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = this.pos, this.line, this.column
		return
	}
	start, end := this.pos, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {

		// fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", this.pos, this.line, this.column, state)

		if this.pos >= len(this.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(this.src[this.pos:])
			this.pos += size
		}
		switch rune1 {
		case '\n':
			this.line++
			this.column = 1
		case '\r':
			this.column = 1
		case '\t':
			this.column += 4
		default:
			this.column++
		}

		// Production start
		if rune1 != -1 {
			state = TransTab[state](rune1)
		} else {
			state = -1
		}
		// Production end

		// Debug start
		// nextState := -1
		// if rune1 != -1 {
		// 	nextState = TransTab[state](rune1)
		// }
		// fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		// fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", this.pos, size, start, end)
		// if nextState != -1 {
		// 	fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		// }
		// state = nextState
		// Debug end

		if state != -1 {
			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				// fmt.Printf("\t Accept(%s), %s(%d)\n", string(act), token.TokMap.Id(tok), tok)
				end = this.pos
			case ActTab[state].Ignore != "":
				// fmt.Printf("\t Ignore(%s)\n", string(act))
				start = this.pos
				state = 0
				if start >= len(this.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = this.pos
			}
		}
	}
	if end > start {
		this.pos = end
		tok.Lit = this.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset = start
	tok.Pos.Column = this.column
	tok.Pos.Line = this.line
	return
}

func (this *Lexer) Reset() {
	this.pos = 0
}

/*
Lexer symbols:
0: '('
1: ')'
2: '('
3: ')'
4: '('
5: '-'
6: ')'
7: '$'
8: '$'
9: '$'
10: '$'
11: '$'
12: '$'
13: '{'
14: ','
15: '}'
16: '['
17: ']'
18: 'b'
19: 'o'
20: 'o'
21: 'l'
22: '['
23: ']'
24: 'i'
25: 'n'
26: 't'
27: '['
28: ']'
29: 'u'
30: 'i'
31: 'n'
32: 't'
33: '['
34: ']'
35: 'd'
36: 'o'
37: 'u'
38: 'b'
39: 'l'
40: 'e'
41: '['
42: ']'
43: 's'
44: 't'
45: 'r'
46: 'i'
47: 'n'
48: 'g'
49: '['
50: ']'
51: '['
52: ']'
53: 'b'
54: 'y'
55: 't'
56: 'e'
57: 't'
58: 'r'
59: 'u'
60: 'e'
61: 'f'
62: 'a'
63: 'l'
64: 's'
65: 'e'
66: '='
67: '('
68: ')'
69: '{'
70: '}'
71: ','
72: ';'
73: '#'
74: '&'
75: '|'
76: '['
77: ']'
78: ':'
79: '!'
80: '*'
81: '_'
82: '~'
83: '.'
84: '@'
85: '-'
86: '>'
87: '='
88: '='
89: '!'
90: '='
91: '<'
92: '>'
93: '<'
94: '='
95: '>'
96: '='
97: '~'
98: '='
99: '*'
100: '='
101: '^'
102: '='
103: '$'
104: '='
105: '/'
106: '/'
107: '\n'
108: '/'
109: '*'
110: '*'
111: '*'
112: '/'
113: ' '
114: '\t'
115: '\n'
116: '\r'
117: '0'
118: '0'
119: 'x'
120: 'X'
121: '-'
122: 'e'
123: 'E'
124: '+'
125: '-'
126: '.'
127: '.'
128: '.'
129: '_'
130: '_'
131: 'd'
132: 'o'
133: 'u'
134: 'b'
135: 'l'
136: 'e'
137: 'i'
138: 'n'
139: 't'
140: 'u'
141: 'i'
142: 'n'
143: 't'
144: '['
145: ']'
146: 'b'
147: 'y'
148: 't'
149: 'e'
150: 's'
151: 't'
152: 'r'
153: 'i'
154: 'n'
155: 'g'
156: 'b'
157: 'o'
158: 'o'
159: 'l'
160: '.'
161: '\'
162: 'U'
163: '\'
164: 'u'
165: '\'
166: 'x'
167: '\'
168: '`'
169: '`'
170: '\'
171: 'a'
172: 'b'
173: 'f'
174: 'n'
175: 'r'
176: 't'
177: 'v'
178: '\'
179: '''
180: '"'
181: '"'
182: '"'
183: '''
184: '''
185: '0'-'9'
186: '0'-'7'
187: '0'-'9'
188: 'A'-'F'
189: 'a'-'f'
190: '1'-'9'
191: 'A'-'Z'
192: 'a'-'z'
193: .

*/
