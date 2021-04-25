// Copyright 2013 Hui Chen
// Copyright 2016 ego authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package types

import "time"

type SearchStatus struct {
	Total      int64
	Successful int64
	Failed     int64
}

type ScoredDocAttri struct {
}

// ScoredDoc scored the document
type ScoredDoc struct {
	Index   string
	ID      string
	Tags    []string
	Content string
	Score   float64
	// new 返回文档属性 Attri
	Attri interface{}
	// new 返回评分字段
	Fields map[string]interface{}
}

type ScoredDocArray []*ScoredDoc

//Len()
func (s ScoredDocArray) Len() int {
	return len(s)
}

//Less():成绩将有低到高排序
func (s ScoredDocArray) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}

//Swap()
func (s ScoredDocArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// SearchResp search response options
type SearchResp struct {
	Status   SearchStatus // 索引结果
	Total    uint64       // 总共
	From     uint64
	Hits     ScoredDocArray
	Took     time.Duration
	MaxScore float64
}
