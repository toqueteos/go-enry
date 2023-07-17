package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"

	"github.com/go-enry/go-enry/v2/data/types"
	"gopkg.in/yaml.v2"
)

type languageInfo struct {
	FSName         string   `yaml:"fs_name"`
	Type           string   `yaml:"type,omitempty"`
	Color          string   `yaml:"color,omitempty"`
	Group          string   `yaml:"group,omitempty"`
	Aliases        []string `yaml:"aliases,omitempty"`
	Extensions     []string `yaml:"extensions,omitempty,flow"`
	Interpreters   []string `yaml:"interpreters,omitempty,flow"`
	Filenames      []string `yaml:"filenames,omitempty,flow"`
	MimeType       string   `yaml:"codemirror_mime_type,omitempty,flow"`
	TMScope        string   `yaml:"tm_scope"`
	AceMode        string   `yaml:"ace_mode"`
	CodeMirrorMode string   `yaml:"codemirror_mode"`
	Wrap           bool     `yaml:"wrap"`
	LanguageID     *int     `yaml:"language_id,omitempty"`
}

func getAlphabeticalOrderedKeys(languages map[string]*languageInfo) []string {
	keyList := make([]string, 0)
	for lang := range languages {
		keyList = append(keyList, lang)
	}

	sort.Strings(keyList)
	return keyList
}

// LanguageInfo generates maps in Go with language name -> LanguageInfo and language ID -> LanguageInfo.
// It is of generator.File type.
func LanguageInfo(fileToParse, samplesDir, outPath, tmplPath, tmplName, commit string) error {
	data, err := ioutil.ReadFile(fileToParse)
	if err != nil {
		return err
	}

	languages := make(map[string]*languageInfo)
	if err := yaml.Unmarshal(data, &languages); err != nil {
		return err
	}

	if err := writeLanguageInfoFile(languages, outPath); err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	if err := executeLanguageInfoTemplate(buf, languages, tmplPath, tmplName, commit); err != nil {
		return err
	}

	return formatedWrite(outPath, buf.Bytes())
}

func writeLanguageInfoFile(langInfo map[string]*languageInfo, filename string) error {
	datafilePath := filename + ".json"
	f, err := os.Create(datafilePath)
	if err != nil {
		return fmt.Errorf("could not create %q: %w", datafilePath, err)
	}
	defer f.Close()

	res := make(map[int]types.LanguageInfo)

	for language, info := range langInfo {
		res[*info.LanguageID] = types.LanguageInfo{
			Name:           language,
			FSName:         info.FSName,
			Type:           types.TypeForString(info.Type),
			Color:          info.Color,
			Group:          info.Group,
			Aliases:        info.Aliases,
			Extensions:     info.Extensions,
			Interpreters:   info.Interpreters,
			Filenames:      info.Filenames,
			MimeType:       info.MimeType,
			TMScope:        info.TMScope,
			AceMode:        info.AceMode,
			CodeMirrorMode: info.CodeMirrorMode,
			Wrap:           info.Wrap,
			LanguageID:     *info.LanguageID,
		}
	}

	enc := json.NewEncoder(f)
	return enc.Encode(res)
}

func executeLanguageInfoTemplate(out io.Writer, languages map[string]*languageInfo, tmplPath, tmplName, commit string) error {
	return executeTemplate(out, tmplName, tmplPath, commit, nil, languages)
}
