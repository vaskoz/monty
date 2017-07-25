package main

import (
	"bytes"
	"flag"
	"os"
	"testing"
)

var testcases = []struct {
	in  []string
	out string
}{
	{[]string{"monty program name", "-doors", "3", "-reveals", "1", "-games", "1"}, ""},
	{[]string{"monty program name", "-doors", "3", "-reveals", "1", "-games", "1", "too much"}, ""},
}

func TestMain(t *testing.T) {
	realExit := exit
	exit = func(code int) {}
	t.Parallel()
	for _, c := range testcases {
		os.Args = c.in
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		buff := new(bytes.Buffer)
		stdout = buff
		main()
		out := buff.String()
		if out != c.out {
			t.Errorf("%v does not equal %v", out, c.out)
		}
	}
	exit = realExit
}

func BenchmarkMain(b *testing.B) {
	realExit := exit
	exit = func(code int) {}
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			os.Args = c.in
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			buff := new(bytes.Buffer)
			stdout = buff
			main()
		}
	}
	exit = realExit
}
