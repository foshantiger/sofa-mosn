/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package types

import (
	"reflect"
	"testing"
)

func TestInitSet(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want SortedStringSetType
	}{
		{
			name: "testcase1",
			args: args{
				input: []string{"ac", "ab", "cd"},
			},
			want: SortedStringSetType{
				keys: []string{
					"ab", "ac", "cd",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitSet(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitSortedMap(t *testing.T) {
	type args struct {
		input map[string]string
	}

	inmap := map[string]string{
		"hello": "yes",
		"bb":    "yes",
		"aa":    "no",
	}

	want := []SortedPair{
		{"aa","no"},
		{"bb", "yes"},
		{"hello","yes"},
	}
	
	tests := []struct {
		name string
		args args
		want SortedMap
	}{
		{
			name: "case1",
			args: args{
				input: inmap,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitSortedMap(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitSortedMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
