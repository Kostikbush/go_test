package main

import "testing"

func TestMaskUrlInMessage(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "http url masked after // until space",
			in:   "See http://example.com now",
			out:  "See http://*********** now",
		},
		{
			name: "https url",
			in:   "x https://a.b/c y",
			out:  "x https://***** y",
		},
		{
			name: "no url",
			in:   "hello world",
			out:  "hello world",
		},
		{
			name: "ends with url",
			in:   "go to http://ex.com",
			out:  "go to http://******",
		},
		{
			name: "double slash not url path a/b",
			in:   "a/b c",
			out:  "a/b c",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := maskUrlInMessage(tc.in)
			if got != tc.out {
				t.Fatalf("want %q, got %q", tc.out, got)
			}
		})
	}
}