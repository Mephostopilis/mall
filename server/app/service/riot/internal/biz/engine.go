package biz

import (
	"edu/pkg/constants"
	"edu/pkg/utils"
	"edu/riot/internal/conf"
	"edu/riot/internal/dao"
	"edu/riot/internal/mapping"

	"csac.pro/csacio/go-csac/util/logging"
	"github.com/blevesearch/bleve"
	"github.com/go-kratos/kratos/v2/log"
)

// Engine initialize the engine
type Engine struct {
	indexer []bleve.Index
	dao     dao.Dao
	log     *log.Helper
}

// New initialize the engine
func New(logger log.Logger, d dao.Dao) (engine *Engine, cf func(), err error) {
	cfg := conf.BleveConfig{
		Indexs: 1,
		Path:   []string{"bleve/0"},
	}

	// 初始索引携程
	engine = &Engine{
		indexer: make([]bleve.Index, cfg.Indexs),
		dao:     d,
	}
	for i := int32(0); i < cfg.Indexs; i++ {
		exists, xerr := utils.PathExists(cfg.Path[i])
		if xerr != nil {
			err = xerr
			return
		}
		if exists {
			index, xerr := bleve.OpenUsing(cfg.Path[i], nil)
			if xerr != nil {
				logging.VLog().Errorf("err = %v", xerr)
				err = xerr
				return
			}
			engine.indexer[i] = index
		} else {
			// 初始索引
			indexMapping := bleve.NewIndexMapping()
			indexMapping.AddDocumentMapping(constants.Block, mapping.NewBlockDocumentMapping())

			indexMapping.AddDocumentMapping(constants.Transaction, mapping.NewNormalTxDocumentMapping())
			indexMapping.AddDocumentMapping(constants.NormalTransaction, mapping.NewNormalTxDocumentMapping())
			indexMapping.AddDocumentMapping(constants.ComplexTransaction, mapping.NewNormalTxDocumentMapping())
			indexMapping.AddDocumentMapping(constants.PledgeTransaction, mapping.NewNormalTxDocumentMapping())
			indexMapping.AddDocumentMapping(constants.ContractDeployTransaction, mapping.NewContractDeployTxDocumentMapping())
			indexMapping.AddDocumentMapping(constants.ContractInvokeTransaction, mapping.NewContractInvokeTxDocumentMapping())
			indexMapping.AddDocumentMapping(constants.ContractChangeStateTransaction, mapping.NewNormalTxDocumentMapping())
			indexMapping.AddDocumentMapping(constants.AuthorizeTransaction, mapping.NewNormalTxDocumentMapping())
			indexMapping.AddDocumentMapping(constants.ReportTransaction, mapping.NewNormalTxDocumentMapping())

			indexMapping.AddDocumentMapping(constants.Account, mapping.NewAccountDocumentMapping())
			indexMapping.AddDocumentMapping(constants.Contract, mapping.NewContractDocumentMapping())
			index, xerr := bleve.NewUsing(cfg.Path[i], indexMapping, bleve.Config.DefaultIndexType, bleve.Config.DefaultKVStore, nil)
			if xerr != nil {
				logging.VLog().Errorf("err = %v", xerr)
				err = xerr
				return
			}
			engine.indexer[i] = index
		}
	}

	cf = engine.Close
	return
}

// IndexerCount get count of indexer
func (engine *Engine) IndexerCount() (cnt int) {
	cnt = len(engine.indexer)
	return
}

// DocCount get count of doc
func (engine *Engine) DocCount() (cnt uint64, err error) {
	cnt = 0
	for i := 0; i < len(engine.indexer); i++ {
		x, xerr := engine.indexer[i].DocCount()
		if xerr != nil {
			logging.VLog().Errorf("err = %v", xerr)
			err = xerr
			return
		}
		cnt = cnt + x
	}
	return
}

// Close close the engine
func (engine *Engine) Close() {
	// for i := 0; i < len(engine.indexer); i++ {
	// 	close(engine.addDocCh[i])
	// }
	// engine.cf()
	// engine.wg.Wait()
	for i := 0; i < len(engine.indexer); i++ {
		engine.indexer[i].Close()
	}
	logging.VLog().Infof("engine closed")
}
