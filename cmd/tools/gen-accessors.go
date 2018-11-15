//+build ignore

/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2017 by authors and contributors.

 gen-accessors generates GetXx, GetXxOk, HasXx and SetXx methods for struct field types with pointers.

 It should be run by the go-datadog-api authors to ge-generate methods before pushing changes to GitHub.
*/

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"
)

const (
	fileSuffix = "-accessors.go"
)

var (
	verbose    = flag.Bool("v", false, "Print verbose log messages")
	sourceTmpl = template.Must(template.New("source").Parse(source))
)

func logf(fmt string, args ...interface{}) {
	if *verbose {
		log.Printf(fmt, args...)
	}
}

func main() {
	flag.Parse()
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, ".", sourceFilter, 0)
	if err != nil {
		log.Fatal(err)
		return
	}

	for pkgName, pkg := range pkgs {
		t := &templateData{
			Year:     time.Now().Year(),
			Filename: pkgName + fileSuffix,
			Package:  pkgName,
			Imports:  map[string]string{},
		}
		for filename, f := range pkg.Files {
			logf("Processing %v...", filename)
			t.sourceFile = filename
			if err := t.processAST(f); err != nil {
				log.Fatal(err)
			}
		}
		if err := t.dump(); err != nil {
			log.Fatal(err)
		}
	}
	logf("Done.")
}

// processAST walks source code files, looks for struct definitions, and calls appropriate helper functions
// to add generate accessors for type fields are star expressions -ie: pointers
func (t *templateData) processAST(f *ast.File) error {
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gd.Specs {
			// Types only
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			// Structs only
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}

			for _, field := range st.Fields.List {
				// Pointers only
				se, ok := field.Type.(*ast.StarExpr)
				if len(field.Names) == 0 || !ok {
					continue
				}

				// TODO: handle more than one types, is that a thing?
				fieldName := field.Names[0]

				switch x := se.X.(type) {
				// An array or slice type
				case *ast.ArrayType:
					t.addArrayType(x, ts.Name.String(), fieldName.String())
				// An identifier. Most of our fields will match this
				case *ast.Ident:
					t.addIdent(x, ts.Name.String(), fieldName.String())
				// A map
				case *ast.MapType:
					t.addMapType(x, ts.Name.String(), fieldName.String())
				// A selector expression, ie json.Number
				case *ast.SelectorExpr:
					t.addSelectorExpr(x, ts.Name.String(), fieldName.String())
				default:
					logf("processAST: type %q, field %q, unknown %T: %+v", ts.Name, fieldName, x, x)
				}
			}
		}
	}
	return nil
}

// sourceFilter is a filter to ignore blacklisted file patterns
func sourceFilter(fi os.FileInfo) bool {
	return !strings.HasSuffix(fi.Name(), "_test.go") && !strings.HasSuffix(fi.Name(), fileSuffix)
}

// dump sorts the results, and writes the generated code to a file
func (t *templateData) dump() error {
	if len(t.Accessors) == 0 {
		logf("No accessors for %v; skipping.", t.Filename)
		return nil
	}

	// Sort accessors by ReceiverType.FieldName
	sort.Sort(byName(t.Accessors))

	var buf bytes.Buffer
	if err := sourceTmpl.Execute(&buf, t); err != nil {
		return err
	}
	clean, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	outFile := filepath.Join(filepath.Dir(t.sourceFile), t.Filename)
	logf("Writing %v...", outFile)
	return ioutil.WriteFile(outFile, clean, 0644)
}

// newAccessor returns a new accessor to be processed later
func newAccessor(receiverType, fieldName, fieldType, zeroValue string) *accessor {
	return &accessor{
		sortVal:      strings.ToLower(receiverType) + "." + strings.ToLower(fieldName),
		ReceiverVar:  strings.ToLower(receiverType[:1]),
		ReceiverType: receiverType,
		FieldName:    fieldName,
		FieldType:    fieldType,
		ZeroValue:    zeroValue,
	}
}

// addIdent adds an accessor for an identifier, for a given receiver and field
func (t *templateData) addIdent(x *ast.Ident, receiverType, fieldName string) {
	var zeroValue string
	switch x.String() {
	case "int":
		zeroValue = "0"
	case "string":
		zeroValue = `""`
	case "bool":
		zeroValue = "false"
	case "float64":
		zeroValue = "0"
	case "Status":
		zeroValue = "0"
	case "PrecisionT":
		zeroValue = `""`
	default:
		zeroValue = fmt.Sprintf("%s{}", x.String())
	}

	t.Accessors = append(t.Accessors, newAccessor(receiverType, fieldName, x.String(), zeroValue))
}

// addSelectorExpr adds an accessor for an identifier selector expression, for a given receiver and field
func (t *templateData) addSelectorExpr(x *ast.SelectorExpr, receiverType, fieldName string) {
	if strings.ToLower(fieldName[:1]) == fieldName[:1] { // non-exported field
		return
	}

	var xX string
	if xx, ok := x.X.(*ast.Ident); ok {
		xX = xx.String()
	}

	switch xX {
	case "time", "json":
		if xX == "json" {
			t.Imports["encoding/json"] = "encoding/json"
		} else {
			t.Imports[xX] = xX
		}
		fieldType := fmt.Sprintf("%v.%v", xX, x.Sel.Name)
		zeroValue := fmt.Sprintf("%v.%v{}", xX, x.Sel.Name)
		if xX == "json" && x.Sel.Name == "Number" {
			zeroValue = `""`
		}
		if xX == "time" && x.Sel.Name == "Duration" {
			zeroValue = "0"
		}
		t.Accessors = append(t.Accessors, newAccessor(receiverType, fieldName, fieldType, zeroValue))
	default:
		logf("addSelectorExpr: xX %q, type %q, field %q: unknown x=%+v; skipping.", xX, receiverType, fieldName, x)
	}
}

// addMypType adds an accessor for a map type, for a given receiver and field
func (t *templateData) addMapType(x *ast.MapType, receiverType, fieldName string) {
	// TODO: should we make this dynamic? Could handle more cases than string only
	var keyType string
	switch key := x.Key.(type) {
	case *ast.Ident:
		keyType = key.String()
	default:
		logf("addMapType: type %q, field %q: unknown key type: %T %+v; skipping.", receiverType, fieldName, key, key)
		return
	}

	var valueType string
	switch value := x.Value.(type) {
	case *ast.Ident:
		valueType = value.String()
	default:
		logf("addMapType: type %q, field %q: unknown value type: %T %+v; skipping.", receiverType, fieldName, value, value)
		return
	}

	fieldType := fmt.Sprintf("map[%v]%v", keyType, valueType)
	zeroValue := fmt.Sprintf("map[%v]%v{}", keyType, valueType)
	t.Accessors = append(t.Accessors, newAccessor(receiverType, fieldName, fieldType, zeroValue))
}

// addArrayType adds an accessor for a array type for a given receiver and field
func (t *templateData) addArrayType(x *ast.ArrayType, receiverType, fieldName string) {
	// TODO: should we make this dynamic? Could handle more cases than string only
	var eltType string
	switch elt := x.Elt.(type) {
	case *ast.Ident:
		eltType = elt.String()
	default:
		logf("addArrayType: type %q, field %q: unknown element type: %T %+v; skipping.", receiverType, fieldName, elt, elt)
		return
	}

	t.Accessors = append(t.Accessors, newAccessor(receiverType, fieldName, "[]"+eltType, "nil"))
}

// hold data used by our template
type templateData struct {
	sourceFile string
	Year       int
	Filename   string
	Package    string
	Imports    map[string]string
	Accessors  []*accessor
}

// our accessor used to render templates
type accessor struct {
	sortVal      string // lower-case version of "ReceiverType.FieldName"
	ReceiverVar  string // the one-letter variable name to match the ReceiverType
	ReceiverType string
	FieldName    string
	FieldType    string
	ZeroValue    string
}

// some helpers to sort
type byName []*accessor

func (b byName) Len() int           { return len(b) }
func (b byName) Less(i, j int) bool { return b[i].sortVal < b[j].sortVal }
func (b byName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

const source = `/*
 THIS FILE IS AUTOMATICALLY GENERATED BY create-accessors; DO NOT EDIT.
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright {{.Year}} by authors and contributors.
 */

package {{.Package}}
{{if .Imports}}
import (
  {{range .Imports}}
  "{{.}}"{{end}}
)
{{end}}
{{range .Accessors}}
// Get{{.FieldName}} returns the {{.FieldName}} field if non-nil, zero value otherwise.
func ({{.ReceiverVar}} *{{.ReceiverType}}) Get{{.FieldName}}() {{.FieldType}} {
  if {{.ReceiverVar}} == nil || {{.ReceiverVar}}.{{.FieldName}} == nil {
    return {{.ZeroValue}}
  }
  return *{{.ReceiverVar}}.{{.FieldName}}
}

// Get{{.FieldName}}Ok returns a tuple with the {{.FieldName}} field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func ({{.ReceiverVar}} *{{.ReceiverType}}) Get{{.FieldName}}Ok() ({{.FieldType}}, bool){
  if {{.ReceiverVar}} == nil || {{.ReceiverVar}}.{{.FieldName}} == nil {
    return {{.ZeroValue}}, false
  }
  return *{{.ReceiverVar}}.{{.FieldName}}, true
}

// Has{{.FieldName}} returns a boolean if a field has been set.
func ({{.ReceiverVar}} *{{.ReceiverType}}) Has{{.FieldName}}() bool {
   if {{.ReceiverVar}} != nil && {{.ReceiverVar}}.{{.FieldName}} != nil {
    return true
   }

   return false
}

// Set{{.FieldName}} allocates a new {{.ReceiverVar}}.{{.FieldName}} and returns the pointer to it.
func ({{.ReceiverVar}} *{{.ReceiverType}}) Set{{.FieldName}}(v {{.FieldType}}) {
  {{.ReceiverVar}}.{{.FieldName}} = &v
}

{{end}}
`
