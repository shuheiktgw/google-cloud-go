// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package bigquery

import (
	"fmt"
	"testing"
	"time"

	"cloud.google.com/go/internal/testutil"
	bq "google.golang.org/api/bigquery/v2"
)

func testRoutineConversion(t *testing.T, conversion string, in interface{}, want interface{}) {
	var got interface{}
	var err error
	switch conversion {
	case "ToRoutineMetadata":
		input, ok := in.(*bq.Routine)
		if !ok {
			t.Fatalf("failed input type conversion (bq.Routine): %v", in)
		}
		got, err = bqToRoutineMetadata(input)
	case "FromRoutineMetadataToUpdate":
		input, ok := in.(*RoutineMetadataToUpdate)
		if !ok {
			t.Fatalf("failed input type conversion: %v", in)
		}
		got, err = input.toBQ()
	case "ToRoutineArgument":
		input, ok := in.(*bq.Argument)
		if !ok {
			t.Fatalf("failed input type conversion: %v", in)
		}
		got, err = bqToRoutineArgument(input)
	case "FromRoutineArgument":
		input, ok := in.(*RoutineArgument)
		if !ok {
			t.Fatalf("failed input type conversion: %v", in)
		}
		got, err = input.toBQ()
	default:
		t.Fatalf("invalid comparison: %s", conversion)
	}

	if err != nil {
		t.Fatalf("failed conversion function for %q", conversion)
	}
	if diff := testutil.Diff(got, want); diff != "" {
		t.Fatalf("%+v: -got, +want:\n%s", in, diff)
	}
}

func TestRoutineTypeConversions(t *testing.T) {
	aTime := time.Date(2019, 3, 14, 0, 0, 0, 0, time.Local)
	aTimeMillis := aTime.UnixNano() / 1e6

	tests := []struct {
		name       string
		conversion string
		in         interface{}
		want       interface{}
	}{
		{"empty", "ToRoutineMetadata", &bq.Routine{}, &RoutineMetadata{}},
		{"basic", "ToRoutineMetadata",
			&bq.Routine{
				CreationTime:     aTimeMillis,
				LastModifiedTime: aTimeMillis,
				DefinitionBody:   "body",
				Description:      "desc",
				Etag:             "etag",
				DeterminismLevel: "DETERMINISTIC",
				RoutineType:      "type",
				Language:         "lang",
				ReturnType:       &bq.StandardSqlDataType{TypeKind: "INT64"},
				ReturnTableType: &bq.StandardSqlTableType{
					Columns: []*bq.StandardSqlField{
						{Name: "field", Type: &bq.StandardSqlDataType{TypeKind: "FLOAT64"}},
					},
				},
			},
			&RoutineMetadata{
				CreationTime:     aTime,
				LastModifiedTime: aTime,
				Description:      "desc",
				DeterminismLevel: Deterministic,
				Body:             "body",
				ETag:             "etag",
				Type:             "type",
				Language:         "lang",
				ReturnType:       &StandardSQLDataType{TypeKind: "INT64"},
				ReturnTableType: &StandardSQLTableType{
					Columns: []*StandardSQLField{
						{Name: "field", Type: &StandardSQLDataType{TypeKind: "FLOAT64"}},
					},
				},
			}},
		{"body_and_libs", "FromRoutineMetadataToUpdate",
			&RoutineMetadataToUpdate{
				Body:              "body",
				ImportedLibraries: []string{"foo", "bar"},
				ReturnType:        &StandardSQLDataType{TypeKind: "FOO"},
			},
			&bq.Routine{
				DefinitionBody:    "body",
				ImportedLibraries: []string{"foo", "bar"},
				ReturnType:        &bq.StandardSqlDataType{TypeKind: "FOO"},
				ForceSendFields:   []string{"DefinitionBody", "ImportedLibraries", "ReturnType"},
			}},
		{"null_fields", "FromRoutineMetadataToUpdate",
			&RoutineMetadataToUpdate{
				Type:              "type",
				Arguments:         []*RoutineArgument{},
				ImportedLibraries: []string{},
			},
			&bq.Routine{
				RoutineType:     "type",
				ForceSendFields: []string{"RoutineType"},
				NullFields:      []string{"Arguments", "ImportedLibraries"},
			}},
		{"empty", "ToRoutineArgument",
			&bq.Argument{},
			&RoutineArgument{}},
		{"basic", "ToRoutineArgument",
			&bq.Argument{
				Name:         "foo",
				ArgumentKind: "bar",
				Mode:         "baz",
			},
			&RoutineArgument{
				Name: "foo",
				Kind: "bar",
				Mode: "baz",
			}},
		{"empty", "FromRoutineArgument",
			&RoutineArgument{},
			&bq.Argument{},
		},
		{"basic", "FromRoutineArgument",
			&RoutineArgument{
				Name: "foo",
				Kind: "bar",
				Mode: "baz",
			},
			&bq.Argument{
				Name:         "foo",
				ArgumentKind: "bar",
				Mode:         "baz",
			}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s/%s", test.conversion, test.name), func(t *testing.T) {
			t.Parallel()
			testRoutineConversion(t, test.conversion, test.in, test.want)
		})
	}
}

func TestRoutineIdentifiers(t *testing.T) {
	testRoutine := &Routine{
		ProjectID: "p",
		DatasetID: "d",
		RoutineID: "r",
		c:         nil,
	}
	for _, tc := range []struct {
		description string
		in          *Routine
		format      IdentifierFormat
		want        string
		wantErr     bool
	}{
		{
			description: "empty format string",
			in:          testRoutine,
			format:      "",
			wantErr:     true,
		},
		{
			description: "legacy",
			in:          testRoutine,
			wantErr:     true,
		},
		{
			description: "standard unquoted",
			in:          testRoutine,
			format:      StandardSQLID,
			want:        "p.d.r",
		},
		{
			description: "standard w/dash",
			in:          &Routine{ProjectID: "p-p", DatasetID: "d", RoutineID: "r"},
			format:      StandardSQLID,
			want:        "`p-p`.d.r",
		},
		{
			description: "api resource",
			in:          testRoutine,
			format:      StorageAPIResourceID,
			wantErr:     true,
		},
	} {
		got, err := tc.in.Identifier(tc.format)
		if tc.wantErr && err == nil {
			t.Errorf("case %q: wanted err, was success", tc.description)
		}
		if !tc.wantErr {
			if err != nil {
				t.Errorf("case %q: wanted success, got err: %v", tc.description, err)
			} else {
				if got != tc.want {
					t.Errorf("case %q:  got %s, want %s", tc.description, got, tc.want)
				}
			}
		}
	}
}
