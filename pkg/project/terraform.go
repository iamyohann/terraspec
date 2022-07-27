package project

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
)

const (
	QueryTypeNone     = iota
	QueryTypeResource = iota
)

var rootSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "terraform",
			LabelNames: nil,
		},
		{
			Type:       "variable",
			LabelNames: []string{"name"},
		},
		{
			Type:       "output",
			LabelNames: []string{"name"},
		},
		{
			Type:       "provider",
			LabelNames: []string{"name"},
		},
		{
			Type:       "resource",
			LabelNames: []string{"type", "name"},
		},
		{
			Type:       "data",
			LabelNames: []string{"type", "name"},
		},
		{
			Type:       "module",
			LabelNames: []string{"name"},
		},
	},
}

func parseHCLFile(path string) *hcl.File {
	filecontents, _ := ioutil.ReadFile(path)

	parser := hclparse.NewParser()
	hclfile, _ := parser.ParseHCL(filecontents, filepath.Base(path))
	return hclfile
}

func FromDirectory(dir string) (TerraformProject, error) {
	var hclFiles []*hcl.File
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if strings.HasSuffix(path, ".tf") {
			hclFiles = append(hclFiles, parseHCLFile(path))
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	return TerraformProject{
		files:     hclFiles,
		queryType: QueryTypeNone,
	}, nil
}

type TerraformProject struct {
	queryType int
	files     []*hcl.File
}

type TerraformResource struct {
	Name       string
	Attributes map[string]interface{}
}

func (p *TerraformProject) FindResources(expression string) []TerraformResource {

	var resources []TerraformResource

	for _, file := range p.files {
		content, _, _ := file.Body.PartialContent(rootSchema)

		for _, block := range content.Blocks {
			if block.Type == "resource" {
				attrs, _ := block.Body.JustAttributes()
				resource := TerraformResource{
					Name:       "something",
					Attributes: getAttributes(file, attrs),
				}
				resources = append(resources, resource)
			}
		}
	}
	return resources
}

func getAttributes(someFile *hcl.File, attrs hcl.Attributes) map[string]interface{} {
	var res map[string]interface{} = make(map[string]interface{})

	for key, value := range attrs {
		rr := value.Expr.Range()
		value := someFile.Bytes[rr.Start.Byte:rr.End.Byte]
		res[key] = string(value)
	}

	return res
}
