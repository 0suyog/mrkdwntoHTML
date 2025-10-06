package parsers

import (
	"testing"

	"github.com/0suyog/mrkdwntoHTML/parser/nodes"
	"github.com/0suyog/mrkdwntoHTML/tokenizer"
)

func Test_match_first(t *testing.T) {
	t.Skip()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		tokens  tokenizer.TokenList
		parsers []IParser
		want    nodes.INode
		want2   int
	}{
		{
			name: "Bold test Asterik Variant",
			tokens: *tokenizer.NewTokenList(
				[]tokenizer.IToken{
					tokenizer.NewToken(tokenizer.ASTERIK, "*"),
					tokenizer.NewToken(tokenizer.ASTERIK, "*"),
					tokenizer.NewToken(tokenizer.TEXT, "meow"),
					tokenizer.NewToken(tokenizer.ASTERIK, "*"),
					tokenizer.NewToken(tokenizer.ASTERIK, "*"),
				},
			),
			parsers: []IParser{BoldParser{}},
			want:    nodes.NewNode("BOLD", "meow", 5),
			want2:   5,
		},
		{
			name: "Bold test Underscore Variant",
			tokens: *tokenizer.NewTokenList(
				[]tokenizer.IToken{
					tokenizer.NewToken(tokenizer.UNDERSCORE, "_"),
					tokenizer.NewToken(tokenizer.UNDERSCORE, "_"),
					tokenizer.NewToken(tokenizer.TEXT, "meow"),
					tokenizer.NewToken(tokenizer.UNDERSCORE, "_"),
					tokenizer.NewToken(tokenizer.UNDERSCORE, "_"),
				},
			),
			parsers: []IParser{BoldParser{}},
			want:    nodes.NewNode("BOLD", "meow", 5),
			want2:   5,
		}, {
			name: "Bold test with wrong parser",
			tokens: *tokenizer.NewTokenList(
				[]tokenizer.IToken{
					tokenizer.NewToken(tokenizer.UNDERSCORE, "_"),
					tokenizer.NewToken(tokenizer.UNDERSCORE, "_"),
					tokenizer.NewToken(tokenizer.TEXT, "meow"),
					tokenizer.NewToken(tokenizer.UNDERSCORE, "_"),
					tokenizer.NewToken(tokenizer.UNDERSCORE, "_"),
				},
			),
			parsers: []IParser{EmphasisParser{}},
			want:    nodes.NullNode{},
			want2:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := match_first(tt.tokens, tt.parsers...)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("match_first() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("match_first() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_match_star(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		tokens tokenizer.TokenList
		p      IParser
		want   []nodes.INode
		want2  int
	}{
		{
			name: "Parser cant parse the provided token so should return empty array",
			tokens: *tokenizer.NewTokenList(
				[]tokenizer.IToken{
					tokenizer.NewToken(tokenizer.UNDERSCORE, "_"),
					tokenizer.NewToken(tokenizer.UNDERSCORE, "_"),
					tokenizer.NewToken(tokenizer.TEXT, "meow"),
					tokenizer.NewToken(tokenizer.UNDERSCORE, "_"),
					tokenizer.NewToken(tokenizer.UNDERSCORE, "_"),
				},
			),
			p:     EmphasisParser{},
			want:  []nodes.INode{},
			want2: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := match_star(tt.tokens, tt.p)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("match_star() = %v, want %v", got, tt.want)
			}
			if true {
				t.Errorf("match_star() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_match_plus(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		tokens  tokenizer.TokenList
		p       IParser
		want    []nodes.INode
		want2   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2, gotErr := match_plus(tt.tokens, tt.p)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("match_plus() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("match_plus() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("match_plus() = %v, want %v", got, tt.want)
			}
			if true {
				t.Errorf("match_plus() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
