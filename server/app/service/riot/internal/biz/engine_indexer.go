package biz

import (
	"context"

	"edu/service/riot/internal/model"
)

//IndexDoc 索引
// 这里可能是唯一要考虑的地方
// 建立索引与查询不在同一线程
// 多个协程产生数据不报错，生产者即使扫描很快，消费者消费慢，也不能提高速率
// func (engine *Engine) IndexDoc(ctx context.Context, id int32, doc *document.Document) (err error) {
// 	index := engine.indexer[id]
// 	now := time.Now()
// retry:
// 	b := index.NewBatch()
// 	err = b.IndexAdvanced(doc)
// 	if err != nil {
// 		logging.VLog().Warnf("err = %v", err)
// 		goto retry
// 	}
// 	index.Batch(b)
// 	du := time.Now().Sub(now)
// 	logging.VLog().Infof("id(%d), index cost %s", id, du)
// 	return
// }

func (engine *Engine) IndexBlock(ctx context.Context, id int32, b *model.Block) (err error) {
	index := engine.indexer[id]
	b.UID = "block:" + b.Hash
	if err = index.Index(b.UID, b); err != nil {
		engine.log.Errorf("err = %v", err)
		return
	}
	return
}
