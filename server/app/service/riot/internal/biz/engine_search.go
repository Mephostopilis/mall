package biz

import (
	"context"
	"sort"

	"edu/srevice/riot/internal/types"

	"github.com/blevesearch/bleve"
)

// Search find the document that satisfies the search criteria.
// This function is thread safe
// 查找满足搜索条件的文档，此函数线程安全
func (engine *Engine) Search(ctx context.Context, req *types.SearchReq) (output *types.SearchResp, err error) {
	output = &types.SearchResp{
		Status: types.SearchStatus{
			Total:      0,
			Successful: 0,
			Failed:     0,
		},
		Total:    0,
		From:     uint64(req.From),
		Took:     0,
		MaxScore: 0,
	}

	query := bleve.NewQueryStringQuery(req.Q)
	searchRequest := bleve.NewSearchRequestOptions(query, int(req.Size), 0, false)
	hits := make(types.ScoredDocArray, 0)
	var xret *bleve.SearchResult
	for i := 0; i < len(engine.indexer); i++ {
		ret, xerr := engine.indexer[i].SearchInContext(ctx, searchRequest)
		if xerr != nil {
			// err = ecode.SearchErr
			continue
		}
		if xret == nil {
			xret = ret
		} else {
			xret.Merge(ret)
		}
	}
	output.Status.Total = int64(xret.Total)
	output.Status.Successful = int64(xret.Status.Successful)
	output.Status.Failed = int64(xret.Status.Failed)
	output.Total = uint64(xret.Total)
	output.Took = xret.Took
	output.MaxScore = xret.MaxScore

	for i := 0; i < len(xret.Hits); i++ {
		hit := xret.Hits[i]
		hits = append(hits, &types.ScoredDoc{
			Index: hit.Index,
			ID:    hit.ID,
			Score: hit.Score,
		})
	}

	if len(hits) > 0 {
		sort.Sort(hits)
	}
	output.Hits = hits
	return
}
