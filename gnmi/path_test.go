// Copyright (c) 2017 Arista Networks, Inc.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package gnmi

import (
	"fmt"
	"testing"

	"github.com/aristanetworks/goarista/test"

	pb "github.com/openconfig/gnmi/proto/gnmi"
)

func p(s ...string) []string {
	return s
}

func TestSplitPath(t *testing.T) {
	for i, tc := range []struct {
		in  string
		exp []string
	}{{
		in:  "/foo/bar",
		exp: p("foo", "bar"),
	}, {
		in:  "/foo/bar/",
		exp: p("foo", "bar"),
	}, {
		in:  "//foo//bar//",
		exp: p("", "foo", "", "bar", ""),
	}, {
		in:  "/foo[name=///]/bar",
		exp: p("foo[name=///]", "bar"),
	}, {
		in:  `/foo[name=[\\\]/]/bar`,
		exp: p(`foo[name=[\\\]/]`, "bar"),
	}, {
		in:  `/foo[name=[\\]/bar`,
		exp: p(`foo[name=[\\]`, "bar"),
	}, {
		in:  "/foo[a=1][b=2]/bar",
		exp: p("foo[a=1][b=2]", "bar"),
	}, {
		in:  "/foo[a=1\\]2][b=2]/bar",
		exp: p("foo[a=1\\]2][b=2]", "bar"),
	}, {
		in:  "/foo[a=1][b=2]/bar\\baz",
		exp: p("foo[a=1][b=2]", "bar\\baz"),
	}} {
		got := SplitPath(tc.in)
		if !test.DeepEqual(tc.exp, got) {
			t.Errorf("[%d] unexpect split for %q. Expected: %v, Got: %v",
				i, tc.in, tc.exp, got)
		}
	}
}

func TestStrPath(t *testing.T) {
	for i, tc := range []struct {
		path string
	}{{
		path: "/",
	}, {
		path: "/foo/bar",
	}, {
		path: "/foo[name=a]/bar",
	}, {
		path: "/foo[a=1][b=2]/bar",
	}, {
		path: "/foo[a=1\\]2][b=2]/bar",
	}, {
		path: "/foo[a=1][b=2]/bar\\/baz",
	}, {
		path: `/foo[a\==1]`,
	}, {
		path: `/f\/o\[o[a\==1]`,
	}} {
		sElms := SplitPath(tc.path)
		pbPath, err := ParseGNMIElements(sElms)
		if err != nil {
			t.Errorf("failed to parse %s: %s", sElms, err)
		}
		s := StrPath(pbPath)
		if tc.path != s {
			t.Errorf("[%d] want %s, got %s", i, tc.path, s)
		}
	}
}

func TestStrPathBackwardsCompat(t *testing.T) {
	for i, tc := range []struct {
		path *pb.Path
		str  string
	}{{
		path: &pb.Path{
			Element: p("foo[a=1][b=2]", "bar"),
		},
		str: "/foo[a=1][b=2]/bar",
	}} {
		got := StrPath(tc.path)
		if got != tc.str {
			t.Errorf("[%d] want %q, got %q", i, tc.str, got)
		}
	}
}

func TestParseElement(t *testing.T) {
	// test cases
	cases := []struct {
		// name is the name of the test useful if you want to run a single test
		// from the command line -run TestParseElement/<name>
		name string
		// in is the path element to be parsed
		in string
		// fieldName is field name (YANG node name) expected to be parsed from the path element.
		// Normally this is simply the path element, or if the path element contains keys this is
		// the text before the first [
		fieldName string
		// keys is a map of the expected key value pairs from within the []s in the
		// `path element.
		//
		// For example prefix[ip-prefix=10.0.0.0/24][masklength-range=26..28]
		// fieldName would be "prefix"
		// keys would be {"ip-prefix": "10.0.0.0/24", "masklength-range": "26..28"}
		keys map[string]string
		// expectedError is the exact error we expect.
		expectedError error
	}{{
		name:      "no_elms",
		in:        "hello",
		fieldName: "hello",
	}, {
		name:          "single_open",
		in:            "[",
		expectedError: fmt.Errorf("failed to find element name in %q", "["),
	}, {
		name:          "no_equal_no_close",
		in:            "hello[there",
		expectedError: fmt.Errorf("failed to find '=' in %q", "[there"),
	}, {
		name:          "no_equals",
		in:            "hello[there]",
		expectedError: fmt.Errorf("failed to find '=' in %q", "[there]"),
	}, {
		name:      "no_left_side",
		in:        "hello[=there]",
		fieldName: "hello",
		keys:      map[string]string{"": "there"},
	}, {
		name:      "no_right_side",
		in:        "hello[there=]",
		fieldName: "hello",
		keys:      map[string]string{"there": ""},
	}, {
		name:          "hanging_escape",
		in:            "hello[there\\",
		expectedError: fmt.Errorf("failed to find '=' in %q", "[there\\"),
	}, {
		name:      "single_name_value",
		in:        "hello[there=where]",
		fieldName: "hello",
		keys:      map[string]string{"there": "where"},
	}, {
		name:      "single_value_with=",
		in:        "hello[there=whe=r=e]",
		fieldName: "hello",
		keys:      map[string]string{"there": "whe=r=e"},
	}, {
		name:      "single_value_with=_and_escaped_]",
		in:        `hello[there=whe=\]r=e]`,
		fieldName: "hello",
		keys:      map[string]string{"there": `whe=]r=e`},
	}, {
		name:      "single_value_with[",
		in:        "hello[there=w[[here]",
		fieldName: "hello",
		keys:      map[string]string{"there": "w[[here"},
	}, {
		name:          "value_single_open",
		in:            "hello[first=value][",
		expectedError: fmt.Errorf("failed to find '=' in %q", "["),
	}, {
		name:          "value_no_close",
		in:            "hello[there=where][somename",
		expectedError: fmt.Errorf("failed to find '=' in %q", "[somename"),
	}, {
		name:          "value_no_equals",
		in:            "hello[there=where][somename]",
		expectedError: fmt.Errorf("failed to find '=' in %q", "[somename]"),
	}, {
		name:      "no_left_side",
		in:        "hello[there=where][=somevalue]",
		fieldName: "hello",
		keys:      map[string]string{"there": "where", "": "somevalue"},
	}, {
		name:      "no_right_side",
		in:        "hello[there=where][somename=]",
		fieldName: "hello",
		keys:      map[string]string{"there": "where", "somename": ""},
	}, {
		name:      "two_name_values",
		in:        "hello[there=where][somename=somevalue]",
		fieldName: "hello",
		keys:      map[string]string{"there": "where", "somename": "somevalue"},
	}, {
		name:      "three_name_values",
		in:        "hello[there=where][somename=somevalue][anothername=value]",
		fieldName: "hello",
		keys: map[string]string{"there": "where", "somename": "somevalue",
			"anothername": "value"},
	}, {
		name:      "aserisk_value",
		in:        "hello[there=*][somename=somevalue][anothername=value]",
		fieldName: "hello",
		keys: map[string]string{"there": "*", "somename": "somevalue",
			"anothername": "value"},
	}, {
		name:      "escaped =",
		in:        `hello[foo\==bar]`,
		fieldName: "hello",
		keys:      map[string]string{"foo=": "bar"},
	}, {
		name:      "escaped [",
		in:        `hell\[o[foo=bar]`,
		fieldName: "hell[o",
		keys:      map[string]string{"foo": "bar"},
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fieldName, keys, err := parseElement(tc.in)
			if !test.DeepEqual(tc.expectedError, err) {
				t.Fatalf("[%s] expected err %#v, got %#v", tc.name, tc.expectedError, err)
			}
			if !test.DeepEqual(tc.keys, keys) {
				t.Fatalf("[%s] expected output %#v, got %#v", tc.name, tc.keys, keys)
			}
			if tc.fieldName != fieldName {
				t.Fatalf("[%s] expected field name %s, got %s", tc.name, tc.fieldName, fieldName)
			}
		})
	}
}

func strToPath(pathStr string) *pb.Path {
	splitPath := SplitPath(pathStr)
	path, _ := ParseGNMIElements(splitPath)
	path.Element = nil
	return path
}

func strsToPaths(pathStrs []string) []*pb.Path {
	var paths []*pb.Path
	for _, splitPath := range SplitPaths(pathStrs) {
		path, _ := ParseGNMIElements(splitPath)
		path.Element = nil
		paths = append(paths, path)
	}
	return paths
}

func TestJoinPath(t *testing.T) {
	cases := []struct {
		paths []*pb.Path
		exp   string
	}{{
		paths: strsToPaths([]string{"/foo/bar", "/baz/qux"}),
		exp:   "/foo/bar/baz/qux",
	},
		{
			paths: strsToPaths([]string{
				"/foo/bar[somekey=someval][otherkey=otherval]", "/baz/qux"}),
			exp: "/foo/bar[otherkey=otherval][somekey=someval]/baz/qux",
		},
		{
			paths: strsToPaths([]string{
				"/foo/bar[somekey=someval][otherkey=otherval]",
				"/baz/qux[somekey=someval][otherkey=otherval]"}),
			exp: "/foo/bar[otherkey=otherval][somekey=someval]/" +
				"baz/qux[otherkey=otherval][somekey=someval]",
		},
		{
			paths: []*pb.Path{
				&pb.Path{Element: []string{"foo", "bar[somekey=someval][otherkey=otherval]"}},
				&pb.Path{Element: []string{"baz", "qux[somekey=someval][otherkey=otherval]"}}},
			exp: "/foo/bar[somekey=someval][otherkey=otherval]/" +
				"baz/qux[somekey=someval][otherkey=otherval]",
		},
	}

	for _, tc := range cases {
		got := JoinPaths(tc.paths...)
		exp := strToPath(tc.exp)
		exp.Element = nil
		if !test.DeepEqual(got, exp) {
			t.Fatalf("ERROR!\n Got: %s,\n Want %s\n", got, exp)
		}
	}
}

func BenchmarkPathElementToSigleElementName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = parseElement("hello")
	}
}

func BenchmarkPathElementTwoKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = parseElement("hello[hello=world][bye=moon]")
	}
}

func BenchmarkPathElementBadKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = parseElement("hello[hello=world][byemoon]")
	}
}

func BenchmarkPathElementMaxKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = parseElement("hello[name=firstName][name=secondName][name=thirdName]" +
			"[name=fourthName][name=fifthName][name=sixthName]")
	}
}

func BenchmarkKeyToString(b *testing.B) {
	for _, bench := range []struct {
		name string
		key  map[string]string
	}{{
		name: "singlekey",
		key:  map[string]string{"abcdefghijkm": "nopqrstuvwxyz"},
	}, {
		name: "fivekeys",
		key: map[string]string{
			"one":   "one",
			"two":   "two",
			"three": "three",
			"four":  "four",
			"five":  "five",
		},
	}, {
		name: "escaped",
		key: map[string]string{
			"foo=": "bar",
			"foo":  "ba]r]",
		},
	}} {
		b.Run(bench.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				KeyToString(bench.key)
			}
		})
	}
}
