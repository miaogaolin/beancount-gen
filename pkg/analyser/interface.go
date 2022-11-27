package analyser

import (
	"fmt"

	"github.com/miaogaolin/beancount-gen/pkg/analyser/alipay"
	"github.com/miaogaolin/beancount-gen/pkg/analyser/htsec"
	"github.com/miaogaolin/beancount-gen/pkg/analyser/huobi"
	"github.com/miaogaolin/beancount-gen/pkg/analyser/wechat"
	"github.com/miaogaolin/beancount-gen/pkg/config"
	"github.com/miaogaolin/beancount-gen/pkg/consts"
	"github.com/miaogaolin/beancount-gen/pkg/ir"
)

// Interface is the interface of analyser.
type Interface interface {
	GetAllCandidateAccounts(cfg *config.Config) map[string]bool
	GetAccountsAndTags(o *ir.Order, cfg *config.Config, target, provider string) (string, string, map[ir.Account]string, []string)
}

// New creates a new analyser.
func New(providerName string) (Interface, error) {
	switch providerName {
	case consts.ProviderAlipay:
		return alipay.Alipay{}, nil
	case consts.ProviderWechat:
		return wechat.Wechat{}, nil
	case consts.ProviderHuobi:
		return huobi.Huobi{}, nil
	case consts.ProviderHtsec:
		return htsec.Htsec{}, nil
	default:
		return nil, fmt.Errorf("Fail to create the analyser for the given name %s", providerName)
	}
}
