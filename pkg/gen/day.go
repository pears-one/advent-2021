package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	usage         = "usage: %s <day int>\n"
	errDayExists  = "day exists, creation stopping"
	errFileCreate = "failed to create directory: %s"
	solutionsPath = "pkg/solutions"
	pkgFormat     = "day%d"
	pkgDeclarationFormat = "package %s\n\n"
	importString = `import "github.com/evanfpearson/advent-2021/pkg/advent"`
)

var (
	genFiles = []AOCFile{
		{
			name: "a.go",
			fn:   "A",
		},
		{
			name: "b.go",
			fn:   "B",
		},
		{
			name: "utils.go",
			fn:   "",
		},
	}
	fTemplate = "func %s(input *advent.Input) (int, error) {\n\treturn 0, nil\n}\n"
)


type AOCFile struct {
	name string
	fn string
}

func main() {
	if len(os.Args) != 2 {
		fail()
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fail()
	}
	pkgName := fmt.Sprintf(pkgFormat, n)
	pkgPath := strings.Join([]string{solutionsPath, pkgName}, string(os.PathSeparator))
	_, err = os.Stat(pkgPath)
	if err == nil {
		fail(errDayExists)
	}
	err = os.Mkdir(pkgPath, 0755)
	if err != nil {
		fail(fmt.Sprintf(errFileCreate, pkgPath))
	}
	for _, f := range genFiles {
		filePath := strings.Join([]string{pkgPath, f.name}, string(os.PathSeparator))
		file, err := os.Create(filePath)
		if err != nil {
			fail(err.Error())
		}
		pkgDeclaration := fmt.Sprintf(pkgDeclarationFormat, pkgName)
		var fn string
		var imports string
		if len(f.fn) > 0 {
			fn = fmt.Sprintf(fTemplate, f.fn)
			imports = fmt.Sprintf("%s\n\n", importString)
		}
		_, err = file.WriteString(fmt.Sprintf("%s%s%s", pkgDeclaration, imports, fn))
		if err != nil {
			fail(err.Error())
		}
	}
}

func fail(messages ...string) {
	for _, msg := range messages {
		fmt.Println(msg)
	}
	fmt.Printf(usage, os.Args[0])
	os.Exit(1)
}