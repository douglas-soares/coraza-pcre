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
	return nil
}
func (o *rx) Evaluate(tx *coraza.Transaction, value string) bool {
	return false
}

func init() {
	operators.Register("rx", func() coraza.RuleOperator { return &rx{} })
}

var _ coraza.RuleOperator = &rx{}
