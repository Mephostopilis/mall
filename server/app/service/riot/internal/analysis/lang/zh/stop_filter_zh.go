package zh

import (
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/analysis/token/stop"
	"github.com/blevesearch/bleve/registry"
)

//StopTokenFilterConstructor
func StopTokenFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	tokenMap, err := cache.TokenMapNamed(StopName)
	if err != nil {
		return nil, err
	}
	return stop.NewStopTokensFilter(tokenMap), nil
}

func init() {
	registry.RegisterTokenFilter(StopName, StopTokenFilterConstructor)
}
