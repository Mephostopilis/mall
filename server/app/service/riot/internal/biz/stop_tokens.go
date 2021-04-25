package biz

import (
	"bufio"
	"os"

	"edu/pkg/ecode"

	"github.com/go-kratos/kratos/v2/log"
)

// StopTokens stop tokens map
type StopTokens struct {
	tokens map[string]bool
	log    *log.Helper
}

// New 从 stopTokenFile 中读入停用词，一个词一行
func NewStopTokens(logger log.Logger, stopTokenFile string) (st *StopTokens, err error) {
	if stopTokenFile == "" {
		err = ecode.ErrDeptId
		return
	}
	log := log.NewHelper("tokens", logger)
	st = &StopTokens{
		tokens: make(map[string]bool),
		log:    log,
	}
	file, err := os.Open(stopTokenFile)
	if err != nil {
		log.Errorf("err = %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			st.tokens[text] = true
		}
	}
	return
}

// IsStopToken to determine whether to stop token
func (st *StopTokens) IsStopToken(token string) bool {
	_, found := st.tokens[token]
	return found
}
