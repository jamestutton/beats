// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("auditbeat", "fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
