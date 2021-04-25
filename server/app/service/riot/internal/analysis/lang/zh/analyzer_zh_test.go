package zh

import (
	"reflect"
	"testing"

	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/registry"
)

func TestEnglishAnalyzer(t *testing.T) {
	tests := []struct {
		input  []byte
		output analysis.TokenStream
	}{
		// stemming
		{
			input: []byte("books"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("book"),
					Position: 1,
					Start:    0,
					End:      5,
				},
			},
		},
		{
			input: []byte("book"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("book"),
					Position: 1,
					Start:    0,
					End:      4,
				},
			},
		},
		// stop word removal
		{
			input:  []byte("the"),
			output: analysis.TokenStream{},
		},
		// possessive removal
		{
			input: []byte("steven's"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("steven"),
					Position: 1,
					Start:    0,
					End:      8,
				},
			},
		},
		{
			input: []byte("steven\u2019s"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("steven"),
					Position: 1,
					Start:    0,
					End:      10,
				},
			},
		},
		{
			input: []byte("steven\uFF07s"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:     []byte("steven"),
					Position: 1,
					Start:    0,
					End:      10,
				},
			},
		},
	}

	cache := registry.NewCache()
	analyzer, err := cache.AnalyzerNamed(AnalyzerName)
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		actual := analyzer.Analyze(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf("expected %v, got %v", test.output, actual)
		}
	}
}
