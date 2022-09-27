// Copyright 2021 Juan Pablo Tosso
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"regexp"
	"strings"

	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/operators"
	"github.com/gijsbers/go-pcre"
)

type rx struct {
	data     string
	compiled bool
	macro    *coraza.Macro
	re       pcre.Regexp
}

func (o *rx) Init(options coraza.RuleOperatorOptions) error {
	data := options.Arguments
	var err error
	re, err := regexp.Compile(`%{.*}`)
	if err != nil {
		return err
	}
	macros := re.FindAllString(data, -1)
	if len(macros) > 0 {
		o.macro, err = coraza.NewMacro(macros[0])
		if err != nil {
			return err
		}
	} else {
		o.compiled = true
		o.re, err = pcre.Compile(data, pcre.DOTALL|pcre.DOLLAR_ENDONLY)
	}
	o.data = data
	return err
}
func (o *rx) Evaluate(tx *coraza.Transaction, value string) bool {
	if !o.compiled && o.macro != nil {
		var err error
		o.re, err = pcre.Compile(strings.Replace(o.data, o.macro.String(), o.macro.Expand(tx), -1), pcre.DOTALL|pcre.DOLLAR_ENDONLY)
		if err != nil {
			tx.WAF.Logger.Error("@rx operator compile macro data error", err)
			return false
		}
		o.compiled = true

	}

	m := o.re.MatcherString(value, 0)
	if tx.Capture {
		for i := 0; i < m.Groups()+1; i++ {
			if i == 10 {
				return true
			}
			tx.CaptureField(i, m.GroupString(i))
		}
	}
	return m.Matches()
}

func init() {
	operators.Register("rx", func() coraza.RuleOperator { return &rx{} })
}

var _ coraza.RuleOperator = &rx{}
