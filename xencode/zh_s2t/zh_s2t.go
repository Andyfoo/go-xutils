// Copyright 2019 Andyfoo
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

package zh_s2t

var Dict *dictS

type dictS struct {
	T2S map[rune]rune
	S2T map[rune]rune
}

func init() {
	Dict = &dictS{}

	Dict.T2S = make(map[rune]rune)
	for k := range Trad {
		Dict.T2S[Trad[k]] = Simp[k]
	}

	Dict.S2T = make(map[rune]rune)
	for k := range Simp {
		Dict.S2T[Simp[k]] = Trad[k]
	}
}

func ToSimp(s string) string {
	r := []rune(s)
	for k := range r {
		_, exists := Dict.T2S[r[k]]
		if exists {
			r[k] = Dict.T2S[r[k]]
		}
	}
	return string(r)
}

func ToTrad(s string) string {
	r := []rune(s)
	for k := range r {
		_, exists := Dict.S2T[r[k]]
		if exists {
			r[k] = Dict.S2T[r[k]]
		}
	}
	return string(r)
}
