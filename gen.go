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

type expfunc struct {
	pkgName  string
	funcName string
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	var defined []expfunc
	var mustImport []string
	// walk all directories
	for i := 2015; i <= time.Now().Year(); i++ {
		dir := path.Join(cwd, strconv.Itoa(i))
		if _, err := os.Stat(dir); err != nil {
			continue
		}
		for day := 1; day < 26; day++ {
			fname := path.Join(dir, "d"+strconv.Itoa(day))
			fname = path.Join(fname, strconv.Itoa(day)+".go")
			if _, err := os.Stat(fname); err != nil {
				continue
			}
			var pos token.Position
			pos.Filename = fname

			file, err := os.Open(pos.Filename)
			if err != nil {
				log.Fatalln(err)
			}
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "", file, 0)
			if err != nil {
				log.Fatalln(err)
			}
			var add bool
			for _, dec := range f.Decls {
				if fdec, ok := dec.(*ast.FuncDecl); ok && fdec != nil {
					name := fdec.Name.String()
					if !(name == "Part1" || name == "Part2") {
						continue
					}
					defined = append(defined, expfunc{
						pkgName:  f.Name.String(),
						funcName: name,
					})
					add = true
				}
			}
			if add {
				mustImport = append(mustImport, f.Name.String())
			}
			file.Close()
		}
		if len(defined) < 1 || len(mustImport) < 1 {
			continue
		}

		sfile, err := os.Create(path.Join(dir, "solver.go"))
		if err != nil {
			log.Fatalln(err)
		}
		if _, err = sfile.WriteString(writefile(i, defined, mustImport)); err != nil {
			log.Fatalln(err)
		}
		sfile.Close()
	}
}

func writefile(year int, defined []expfunc, imports []string) string {
	pname := strconv.Itoa(year)
	pname = pname[2:]
	pname = "y" + pname
	var bldr strings.Builder
	bldr.WriteString("package " + pname + "\n\n")
	bldr.WriteString("// code generated by ../gen.go\n\n")
	bldr.WriteString("import(\n\t\"aoc\"\n\n")

	for _, imp := range imports {
		bldr.WriteString("\t\"aoc/" + strconv.Itoa(year) + "/" + imp + "\"\n")
	}
	bldr.WriteString(")\n\n")
	bldr.WriteString("type solution func(string) (string, error)\n\n")
	bldr.WriteString("var solutions [25][2]solution\n\n")
	bldr.WriteString("func init() {\n")
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
		bldr.WriteString("\tsolutions[" + strconv.Itoa(day) + "][" + strconv.Itoa(part) + "] = " + def.pkgName + "." + def.funcName + "\n")
	}

	bldr.WriteString("}\n\n")

	bldr.WriteString(`func Solve(day, part int) (string, error) {
	if part != 1 && part != 2 {
		return "", aoc.ErrUndefined
	}
	if day-1 >= len(solutions) {
		return "", aoc.ErrUndefined
	}
	if sol := solutions[day-1][part-1]; sol != nil {
		return sol("input.txt")
	}
	return "", aoc.ErrUndefined
}

`)

	return bldr.String()
}
