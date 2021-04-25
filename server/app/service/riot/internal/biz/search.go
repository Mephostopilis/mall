package biz

import (
	"context"
	"encoding/json"
	"strings"

	pb "edu/api/riot"

	"edu/riot/internal/types"
)

// Search 搜索数据
func (s *Engine) Search(ctx context.Context, req *pb.SearchReq) (reply *pb.SearchResp, err error) {
	out, xerr := s.engine.Search(ctx, &types.SearchReq{
		Q:    req.Query,
		From: req.From,
		Size: req.Count,
	})
	if xerr != nil {
		err = xerr
		return
	}
	reply = &pb.SearchResp{
		Status: &pb.SearchStatus{
			Total:      out.Status.Total,
			Successful: out.Status.Successful,
			Failed:     out.Status.Failed,
		},
		Total:    out.Total,
		Took:     out.Took.String(),
		MaxScore: out.MaxScore,
	}
	if out.Total > 0 {
		docs := make([]*pb.Doc, 0)
		for i := 0; i < len(out.Hits); i++ {
			hit := out.Hits[i]
			ctn := ""
			tags := []string{}
			sps := strings.Split(hit.ID, ":")
			if len(sps) >= 2 {
				if sps[0] == "block" {
					blk, err := s.dao.GetBlockByHash(ctx, sps[1])
					if err != nil {
						continue
					}
					x, err := json.Marshal(blk)
					if err != nil {
						continue
					}
					ctn = string(x[:])
					tags = append(tags, sps[0])
				} else if sps[0] == "normaltx" {
					tx, err := s.dao.GetNormalTxByHash(ctx, sps[1])
					if err != nil {
						continue
					}
					x, err := json.Marshal(tx)
					if err != nil {
						continue
					}
					ctn = string(x[:])
					tags = append(tags, sps[0])
				} else if sps[0] == "complextx" {
					tx, err := s.dao.GetComplexTxByHash(ctx, sps[1])
					if err != nil {
						continue
					}
					x, err := json.Marshal(tx)
					if err != nil {
						continue
					}
					ctn = string(x[:])
					tags = append(tags, sps[0])
				} else if sps[0] == "pledgetx" {
					tx, err := s.dao.GetPledgeTxByHash(ctx, sps[1])
					if err != nil {
						continue
					}
					x, err := json.Marshal(tx)
					if err != nil {
						continue
					}
					ctn = string(x[:])
					tags = append(tags, sps[0])
				} else if sps[0] == "contractdeploytx" {
					tx, err := s.dao.GetContractDeployTxByHash(ctx, sps[1])
					if err != nil {
						continue
					}
					x, err := json.Marshal(tx)
					if err != nil {
						continue
					}
					ctn = string(x[:])
					tags = append(tags, sps[0])
				} else if sps[0] == "contractinvoketx" {
					tx, err := s.dao.GetContractInvokeTxByHash(ctx, sps[1])
					if err != nil {
						continue
					}
					x, err := json.Marshal(tx)
					if err != nil {
						continue
					}
					ctn = string(x[:])
					tags = append(tags, sps[0])
				} else if sps[0] == "contractchangestatetx" {
					tx, err := s.dao.GetContractChangeStateTxByHash(ctx, sps[1])
					if err != nil {
						continue
					}
					x, err := json.Marshal(tx)
					if err != nil {
						continue
					}
					ctn = string(x[:])
					tags = append(tags, sps[0])
				} else if sps[0] == "authorizetx" {
					tx, err := s.dao.GetAuthorizeTxByHash(ctx, sps[1])
					if err != nil {
						continue
					}
					x, err := json.Marshal(tx)
					if err != nil {
						continue
					}
					ctn = string(x[:])
					tags = append(tags, sps[0])
				} else if sps[0] == "reporttx" {
					tx, err := s.dao.GetReportTxByHash(ctx, sps[1])
					if err != nil {
						continue
					}
					x, err := json.Marshal(tx)
					if err != nil {
						continue
					}
					ctn = string(x[:])
					tags = append(tags, sps[0])
				} else if sps[0] == "account" {
					blk, err := s.dao.GetAccountByAddress(ctx, sps[1])
					if err != nil {
						continue
					}
					x, err := json.Marshal(blk)
					if err != nil {
						continue
					}
					ctn = string(x[:])
					tags = append(tags, sps[0])
				} else if sps[0] == "contract" {
					blk, err := s.md.GetContract(ctx, sps[1])
					if err != nil {
						continue
					}
					x, err := json.Marshal(blk)
					if err != nil {
						continue
					}
					ctn = string(x[:])
					tags = append(tags, sps[0])
				}
				if len(tags) > 0 {
					docs = append(docs, &pb.Doc{
						Index:   hit.Index,
						ID:      hit.ID,
						Tags:    tags,
						Content: ctn,
						Score:   hit.Score,
					})
				}
			}
		}
		reply.Hits = docs
	}
	return
}
