package config

import (
	"github.com/miaogaolin/beancount-gen/pkg/provider/alipay"
	"github.com/miaogaolin/beancount-gen/pkg/provider/htsec"
	"github.com/miaogaolin/beancount-gen/pkg/provider/huobi"
	"github.com/miaogaolin/beancount-gen/pkg/provider/wechat"
)

// Config is the global configuration.
type Config struct {
	Title                    string         `yaml:"title,omitempty"`
	DefaultMinusAccount      string         `yaml:"defaultMinusAccount,omitempty"`
	DefaultPlusAccount       string         `yaml:"defaultPlusAccount,omitempty"`
	DefaultCashAccount       string         `yaml:"defaultCashAccount,omitempty"`
	DefaultPositionAccount   string         `yaml:"defaultPositionAccount,omitempty"`
	DefaultCommissionAccount string         `yaml:"defaultCommissionAccount,omitempty"`
	DefaultPnlAccount        string         `yaml:"defaultPnlAccount,omitempty"`
	DefaultCurrency          string         `yaml:"defaultCurrency,omitempty"`
	Alipay                   *alipay.Config `yaml:"alipay,omitempty"`
	Wechat                   *wechat.Config `yaml:"wechat,omitempty"`
	Huobi                    *huobi.Config  `yaml:"huobi,omitempty"`
	Htsec                    *htsec.Config  `yaml:"htsec,omitempty"`
}
