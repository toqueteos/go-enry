package types

// LanguageInfo exposes the data for a language's Linguist YAML entry as a Go struct.
// See https://github.com/github/linguist/blob/master/lib/linguist/languages.yml
type LanguageInfo struct {
	// Name is the language name. May contain symbols not safe for use in some filesystems (e.g., `F*`).
	Name string
	// FSName is the filesystem safe name. Will only be set if Name is not safe for use in all filesystems.
	FSName string
	// Type is the language Type. See data.Type for values.
	Type Type
	// Color is the CSS hex color to represent the language. Only used if type is "programming" or "markup".
	Color string
	// Group is the name of the parent language. Languages in a group are counted in the statistics as the parent language.
	Group string
	// Aliases is a slice of additional aliases (implicitly includes name.downcase)
	Aliases []string
	// Extensions is a slice of associated extensions (the first one is considered the primary extension).
	Extensions []string
	// A slice of associated interpreters
	Interpreters []string
	// Filenames is a slice of filenames commonly associated with the language.
	Filenames []string
	// MimeType (maps to codemirror_mime_type in linguist.yaml) is the string name of the file mime type used for highlighting whenever a file is edited.
	MimeType string
	// TMScope is the TextMate scope that represents this programming language.
	TMScope string
	// AceMode is the name of the Ace Mode used for highlighting whenever a file is edited.
	AceMode string
	// CodeMirrorMode is the name of the CodeMirror Mode used for highlighting whenever a file is edited.
	CodeMirrorMode string
	// Wrap is a boolean flag to enable line wrapping in an editor.
	Wrap bool
	// LanguageID is the Linguist-assigned numeric ID for the language.
	LanguageID int
}
