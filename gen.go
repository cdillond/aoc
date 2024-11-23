//go:build ignore

package main

/*
	This file is used to build the solver.go file for each year, based on the functions
	defined in the year's subpackages.
*/

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type solutionFunc struct {
	pkgName  string
	funcName string
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	// walk all directories
	for year := 2015; year <= time.Now().Year(); year++ {
		var (
			defined    []solutionFunc
			mustImport []string
		)
		dir := path.Join(cwd, strconv.Itoa(year))
		if _, err := os.Stat(dir); err != nil {
			continue
		}
		for day := 1; day < 26; day++ {
			dayStr := strconv.Itoa(day)
			srcFilePath := path.Join(dir, "d"+dayStr)
			srcFilePath = path.Join(srcFilePath, dayStr+".go")
			if _, err := os.Stat(srcFilePath); err != nil {
				continue
			}
			var pos token.Position
			pos.Filename = srcFilePath

			srcFile, err := os.Open(pos.Filename)
			if err != nil {
				log.Fatalln(err)
			}
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "", srcFile, 0)
			if err != nil {
				log.Fatalln(err)
			}
			var include bool
			for _, dec := range f.Decls {
				if funcDec, ok := dec.(*ast.FuncDecl); ok && funcDec != nil {
					name := funcDec.Name.String()
					if !(name == "Part1" || name == "Part2") {
						continue
					}
					defined = append(defined, solutionFunc{
						pkgName:  f.Name.String(),
						funcName: name,
					})
					include = true
				}
			}
			if include {
				mustImport = append(mustImport, f.Name.String())
			}
			srcFile.Close()
		}
		if len(defined) < 1 || len(mustImport) < 1 {
			continue
		}

		genFile, err := os.Create(path.Join(dir, "solver.go"))
		if err != nil {
			log.Fatalln(err)
		}
		if _, err = genFile.WriteString(buildContent(year, defined, mustImport)); err != nil {
			log.Fatalln(err)
		}
		genFile.Close()
	}
}

func buildContent(year int, defined []solutionFunc, imports []string) string {
	pkgName := strconv.Itoa(year)
	pkgName = pkgName[2:]
	pkgName = "y" + pkgName
	var bldr strings.Builder
	bldr.WriteString("package " + pkgName + "\n\n")
	bldr.WriteString("// code generated by ../gen.go\n\n")
	bldr.WriteString("import (\n\t\"aoc\"\n\n")

	for _, imp := range imports {
		bldr.WriteString("\t\"aoc/" + strconv.Itoa(year) + "/" + imp + "\"\n")
	}
	bldr.WriteString(")\n\n")
	bldr.WriteString("type solution func(string) (string, error)\n\n")

	bldr.WriteString(`func Solve(day, part int, path string) (string, error) {
	if part != 1 && part != 2 {
		return "", aoc.ErrUndefined
	}

	day--
	part--`)
	bldr.WriteString("\n\n\tsolutions := [...]solution{\n")
	for _, def := range defined {
		day, err := strconv.Atoi(def.pkgName[1:])
		if err != nil {
			continue
		}
		day--
		var part int
		if def.funcName == "Part2" {
			part = 1
		}
		bldr.WriteString("\t\t" + strconv.Itoa(day*2+part) + ":\t" + def.pkgName + "." + def.funcName + ",\n")
	}
	bldr.WriteString("\t}\n")
	bldr.WriteString(`
	i := (2 * day) + part
	if i < 0 || i >= len(solutions) {
		return "", aoc.ErrUndefined
	}
	if solve := solutions[i]; solve != nil {
		return solve(path)
	}
	return "", aoc.ErrUndefined
}`)
	bldr.WriteString("\n")

	return bldr.String()
}
