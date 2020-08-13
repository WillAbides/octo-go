package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/willabides/octo-go/generator/internal/model"
	"github.com/willabides/octo-go/generator/internal/model/openapi"
)

func main() {
	var schemaPath string
	var outputPath string
	var pkgPath string
	var pkgName string
	flag.StringVar(&schemaPath, "schema", "", "path to openapi schema")
	flag.StringVar(&outputPath, "out", "", "directory to write all these files")
	flag.StringVar(&pkgPath, "pkgpath", "", "path for output package")
	flag.StringVar(&pkgName, "pkg", "", "name for output package")
	flag.Parse()
	err := run(schemaPath, outputPath, pkgPath, pkgName)
	if err != nil {
		log.Fatal(err)
	}
}

func mkDirs(outputPath string) error {
	for _, dirName := range []string{"components/examples", "internal"} {
		target := filepath.Join(outputPath, filepath.FromSlash(dirName))
		err := os.MkdirAll(target, 0o750)
		if err != nil {
			return err
		}
	}
	return nil
}

func run(schemaPath, outputPath, pkgPath, pkgName string) error {
	err := mkDirs(outputPath)
	if err != nil {
		return err
	}
	err = componentExampleFiles(schemaPath, filepath.Join(outputPath, "components", "examples"))
	if err != nil {
		return err
	}
	schemaFile, err := os.Open(schemaPath)
	if err != nil {
		return err
	}
	mdl, err := openapi.Openapi2Model(schemaFile)
	if err != nil {
		return err
	}
	endpoints := mdl.Endpoints

	sort.Slice(endpoints, func(i, j int) bool {
		return endpoints[i].ID < endpoints[j].ID
	})

	reqFiles := map[string]*jen.File{}
	for i := range endpoints {
		endpoint := endpoints[i].Clone()
		err = addEndpointToRequestFiles(endpoint, reqFiles, pkgPath)
		if err != nil {
			return err
		}
	}

	err = renderRequestFiles(reqFiles, outputPath)
	if err != nil {
		return err
	}

	pq := pkgQual(pkgPath)
	componentSchemas := mdl.ComponentSchemas
	if len(componentSchemas) > 0 {
		filename := "schemas_gen.go"
		var f *os.File
		f, err = os.Create(filepath.Join(outputPath, "components", filename))
		if err != nil {
			return err
		}
		cf := jen.NewFilePath(path.Join(pkgPath, "components"))
		cf.HeaderComment("Code generated by octo-go; DO NOT EDIT.")
		addComponentSchemas(cf, componentSchemas, pq)
		err = cf.Render(f)
		if err != nil {
			return err
		}
	}

	err = generateUnmarshalTests(outputPath, pq, endpoints)
	if err != nil {
		return err
	}

	ccf := concernClientFile(endpoints, pkgPath, pkgName)
	ccFile, err := os.Create(filepath.Join(outputPath, "client_gen.go"))
	if err != nil {
		return err
	}
	err = ccf.Render(ccFile)
	if err != nil {
		return err
	}

	return nil
}

func renderRequestFiles(reqFiles map[string]*jen.File, outputPath string) error {
	for pkg, reqFile := range reqFiles {
		dir := filepath.Join(outputPath, "requests", pkg)
		err := os.MkdirAll(dir, 0o750)
		if err != nil {
			return err
		}
		var f *os.File
		f, err = os.Create(filepath.Join(dir, pkg+"_gen.go"))
		if err != nil {
			return err
		}
		err = reqFile.Render(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func addEndpointToRequestFiles(endpoint *model.Endpoint, requestFiles map[string]*jen.File, pkgPath string) error {
	pq := pkgQual(pkgPath)
	err := applyEndpointOverrides(endpoint)
	if err != nil {
		return err
	}
	if endpoint.Legacy {
		return nil
	}
	reqPkg := concernPackage(endpoint.Concern)
	if requestFiles[reqPkg] == nil {
		reqPkgPath := path.Join(pkgPath, "requests", reqPkg)
		cf := jen.NewFilePathName(reqPkgPath, reqPkg)
		cf.HeaderComment("Code generated by octo-go; DO NOT EDIT.")
		cf.Comment("Client is a set of options to apply to requests")
		cf.Type().Id("Client").Op("[]").Qual(pq.pkgPath("requests"), "Option")
		cf.Comment("NewClient returns a new Client")
		cf.Func().Id("NewClient").Params(
			jen.Id("opt").Op("...").Qual(pq.pkgPath("requests"), "Option"),
		).Id("Client").Block(
			jen.Return(jen.Id("opt")),
		)
		requestFiles[reqPkg] = cf
	}
	file := requestFiles[reqPkg]
	addRequestFunc(file, pq, endpoint)
	addClientMethod(file, pq, endpoint)
	addRequestStruct(file, pq, endpoint)
	addRequestBody(file, pq, endpoint)
	addResponseBody(file, pq, endpoint)
	addResponse(file, pq, endpoint)
	return nil
}

func concernClientFile(endpoints []*model.Endpoint, pkgPath, pkgName string) *jen.File {
	file := jen.NewFilePathName(pkgPath, pkgName)
	file.HeaderComment("Code generated by octo-go; DO NOT EDIT.")
	pq := pkgQual(pkgPath)
	concernsMap := map[string]bool{}
	for _, endpoint := range endpoints {
		concernsMap[endpoint.Concern] = true
	}
	concerns := make([]string, 0, len(concernsMap))
	for k := range concernsMap {
		concerns = append(concerns, k)
	}
	sort.Strings(concerns)
	for _, c := range concerns {
		file.Add(concernClient(c, pq))
	}
	return file
}

func concernClient(concern string, pq pkgQual) *jen.Statement {
	concernPkg := concernPackage(concern)
	qualClient := &qualifiedType{
		pkg:  path.Join("requests", concernPkg),
		name: "Client",
	}
	qualNewClient := &qualifiedType{
		pkg:  path.Join("requests", concernPkg),
		name: "NewClient",
	}
	stmt := jen.Commentf("%s returns a %s", toExportedName(concern), fmt.Sprintf("%s.Client", concernPkg))
	stmt.Line()
	stmt.Id("func (c Client)").Id(toExportedName(concern)).Params().Add(qualClient.jenType(pq)).Block(
		jen.Return(qualNewClient.jenType(pq).Call(jen.Id("c..."))),
	)
	stmt.Line()
	return stmt
}

func concernPackage(concern string) string {
	return strings.ReplaceAll(concern, "-", "")
}
