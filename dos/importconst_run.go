// +build ignore

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// CSource is the filename of the temporary file *.c
const CSource = "makeconst.cpp"

// CC is the Compiler command
const CC = "gcc"

var GoSource = flag.String("o", "const.go", "go-source-filename to output constants")

var clean = flag.Bool("c", false, "clean output")
var debug = flag.Bool("d", false, "debug flag")
var packageName = os.Getenv("GOPACKAGE")

func makeCSource(csrcname string, headers []string, vars []string) {
	fd, err := os.Create(csrcname)
	if err != nil {
		fmt.Fprintf(fd, "%s: can not create %s\n", os.Args[0], csrcname)
		return
	}
	defer fd.Close()

	for _, header1 := range headers {
		fmt.Fprintf(fd, "#include %s\n", header1)
	}
	fmt.Fprint(fd, `
void p(const char *name,const char *s){
	printf("const %s=\"%s\"\n",name,s);
}
void p(const char *name,int n){
	printf("const %s=%d\n",name,n);
}
void p(const char *name,long n){
	printf("const %s=%ld\n",name,n);
}
void p(const char *name,unsigned long n){
	printf("const %s=%ld\n",name,n);
}
void p(const char *name,double n){
	printf("const %s=%lf\n",name,n);
}

int main()
{
`)
	fmt.Fprintln(fd, `    printf("package `+packageName+`\n\n");`)

	for _, name1 := range vars {
		fmt.Fprintf(fd, "    p(\"%s\",%s);\n", name1, name1)
	}
	fmt.Fprintln(fd, "    return 0;\n}\n")
}

func compile() error {
	var gcc exec.Cmd
	gcc.Args = []string{
		CC,
		CSource,
	}
	fn, err := exec.LookPath(CC)
	if err != nil {
		return err
	}
	gcc.Path = fn
	gcc.Stdout = os.Stdout
	gcc.Stderr = os.Stderr
	return gcc.Run()
}

func nameOfExecutable() string {
	if runtime.GOOS == "windows" {
		return "a.exe"
	} else {
		return "a.out"
	}
}

func aexe() (string, error) {
	constC, err := os.Create(*GoSource)
	if err != nil {
		return "", err
	}
	defer constC.Close()

	fname := nameOfExecutable()
	aexe := exec.Cmd{
		Args:   []string{fname},
		Path:   fname,
		Stdout: constC,
		Stderr: os.Stderr,
	}
	return fname, aexe.Run()
}

func gofmt() error {
	var gofmt exec.Cmd
	gofmt.Args = []string{
		"go",
		"fmt",
		*GoSource,
	}
	fn, err := exec.LookPath("go")
	if err != nil {
		return err
	}
	gofmt.Path = fn
	gofmt.Stdout = os.Stdout
	gofmt.Stderr = os.Stderr
	return gofmt.Run()
}

func readGoGenerateParameter() ([]string, error) {
	gofile := os.Getenv("GOFILE")
	if gofile == "" {
		return nil, errors.New("$GOFILE is not defined. Use `go generate`")
	}
	fd, err := os.Open(gofile)
	if err != nil {
		return nil, err
	}
	goline := os.Getenv("GOLINE")
	if goline == "" {
		return nil, errors.New("$GOLINE is not defined. Use `go generate`")
	}
	lnum, err := strconv.Atoi(goline)
	if err != nil {
		return nil, fmt.Errorf("$GOLINE: %s", err.Error())
	}
	sc := bufio.NewScanner(fd)
	var tokens []string
	for sc.Scan() {
		lnum--
		if lnum >= 0 {
			continue
		}
		text := sc.Text()
		if !strings.HasPrefix(text, "//") {
			break
		}
		fields := strings.Fields(text[2:])
		if len(fields) <= 0 {
			break
		}
		for _, arg1 := range fields {
			tokens = append(tokens, arg1)
		}
	}
	return tokens, nil
}

func main1() error {
	flag.Parse()

	if *clean {
		os.Remove(CSource)
		os.Remove(nameOfExecutable())
		os.Remove(*GoSource)
		return nil
	}
	goParams, err := readGoGenerateParameter()
	if err != nil {
		return err
	}
	headers := []string{"<cstdio>"}
	if runtime.GOOS == "windows" {
		headers = append(headers, "<windows.h>")
	}
	vars := make([]string, 0)
	for _, s := range goParams {
		if len(s) > 0 && s[0] == '<' {
			headers = append(headers, s)
		} else if strings.HasSuffix(s, ".h") {
			headers = append(headers, fmt.Sprintf(`"%s"`, s))
		} else {
			vars = append(vars, s)
		}
	}

	makeCSource(CSource, headers, vars)

	if !*debug {
		defer os.Remove(CSource)
	}

	err = compile()
	if err != nil {
		return err
	}
	fname, err := aexe()
	if err != nil {
		return err
	}
	os.Remove(fname)
	err = gofmt()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := main1(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
