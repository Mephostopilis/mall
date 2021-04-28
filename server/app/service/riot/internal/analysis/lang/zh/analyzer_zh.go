package zh

import (
	"edu/service/riot/internal/analysis/tokenizer/gse"

	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/analysis/token/lowercase"
	"github.com/blevesearch/bleve/registry"
)

const AnalyzerName = "zh"

func AnalyzerConstructor(config map[string]interface{}, cache *registry.Cache) (*analysis.Analyzer, error) {
	tokenizer, err := cache.TokenizerNamed(gse.Name)
	if err != nil {
		return nil, err
	}
	toLowerFilter, err := cache.TokenFilterNamed(lowercase.Name)
	if err != nil {
		return nil, err
	}
	stopEnFilter, err := cache.TokenFilterNamed("stop_en")
	if err != nil {
		return nil, err
	}
	stopZhFilter, err := cache.TokenFilterNamed(StopName)
	if err != nil {
		return nil, err
	}
	// stemmerEnFilter, err := cache.TokenFilterNamed(porter.Name)
	// if err != nil {
	// 	return nil, err
	// }
	rv := analysis.Analyzer{
		Tokenizer: tokenizer,
		TokenFilters: []analysis.TokenFilter{
			toLowerFilter,
			stopEnFilter,
			// stemmerEnFilter,
			stopZhFilter,
		},
	}
	return &rv, nil
}

func init() {
	registry.RegisterAnalyzer(AnalyzerName, AnalyzerConstructor)
}
