package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var typeToTypeConst = map[string]int{
	"data":        1,
	"programming": 2,
	"markup":      3,
	"prose":       4,
}

// Types reads from fileToParse and builds source file from tmplPath. It complies with type File signature.
func Types(fileToParse, samplesDir, outPath, tmplPath, tmplName, commit string) error {
	data, err := ioutil.ReadFile(fileToParse)
	if err != nil {
		return err
	}

	languages := make(map[string]*languageInfo)
	if err := yaml.Unmarshal(data, &languages); err != nil {
		return err
	}

	langTypeMap := buildLanguageTypeMap(languages)

	if err := writeTypesFile(langTypeMap, outPath); err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	if err := executeTypesTemplate(buf, tmplPath, tmplName, commit); err != nil {
		return err
	}

	return formatedWrite(outPath, buf.Bytes())
}

func buildLanguageTypeMap(languages map[string]*languageInfo) map[string]int {
	langTypeMap := make(map[string]int)
	for lang, info := range languages {
		langTypeMap[lang] = typeToTypeConst[info.Type]
	}

	return langTypeMap
}

func writeTypesFile(types map[string]int, filename string) error {
	datafilePath := filename + ".json"
	f, err := os.Create(datafilePath)
	if err != nil {
		return fmt.Errorf("could not create %q: %w", datafilePath, err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	return enc.Encode(types)
}

func executeTypesTemplate(out io.Writer, tmplPath, tmplName, commit string) error {
	return executeTemplate(out, tmplName, tmplPath, commit, nil, nil)
}
