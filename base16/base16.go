package base16

import (
	"io"

	"github.com/topi314/chroma/v2"
	"gopkg.in/yaml.v3"
)

type Base16 struct {
	// Scheme is the name of the scheme.
	Scheme string `yaml:"scheme"`
	// Author is the name of the author.
	Author string `yaml:"author"`
	// Theme is either "light" or "dark".
	Theme string `yaml:"theme"`
	// Base00 Default Background
	Base00 string `yaml:"base00"`
	// Base01 Lighter Background (Used for status bars, line number and folding marks)
	Base01 string `yaml:"base01"`
	// Base02 Selection Background
	Base02 string `yaml:"base02"`
	// Base03 Comments, Invisibles, Line Highlighting
	Base03 string `yaml:"base03"`
	// Base04 Dark Foreground (Used for status bars)
	Base04 string `yaml:"base04"`
	// Base05 Default Foreground, Caret, Delimiters, Operators
	Base05 string `yaml:"base05"`
	// Base06 Light Foreground (Not often used)
	Base06 string `yaml:"base06"`
	// Base07 Light Background (Not often used)
	Base07 string `yaml:"base07"`
	// Base08 Variables, XML Tags, Markup Link Text, Markup Lists, Diff Deleted
	Base08 string `yaml:"base08"`
	// Base09 Integers, Boolean, Constants, XML Attributes, Markup Link Url
	Base09 string `yaml:"base09"`
	// Base0A Classes, Markup Bold, Search Text Background
	Base0A string `yaml:"base0A"`
	// Base0B Strings, Inherited Class, Markup Code, Diff Inserted
	Base0B string `yaml:"base0B"`
	// Base0C Support, Regular Expressions, Escape Characters, Markup Quotes
	Base0C string `yaml:"base0C"`
	// Base0D Functions, Methods, Attribute IDs, Headings
	Base0D string `yaml:"base0D"`
	// Base0E Keywords, Storage, Selector, Markup Italic, Diff Changed
	Base0E string `yaml:"base0E"`
	// Base0F Deprecated, Opening/Closing Embedded Language Tags, e.g. `<?php ?>`
	Base0F string `yaml:"base0F"`
}

func (b *Base16) toStyle() (*chroma.Style, error) {
	entries := chroma.StyleEntries{
		chroma.Other:                    "#" + b.Base05,
		chroma.Error:                    "#" + b.Base08,
		chroma.Background:               "bg:" + "#" + b.Base00,
		chroma.Keyword:                  "#" + b.Base0E,
		chroma.KeywordConstant:          "#" + b.Base0E,
		chroma.KeywordDeclaration:       "#" + b.Base08,
		chroma.KeywordNamespace:         "#" + b.Base0E,
		chroma.KeywordPseudo:            "#" + b.Base0E,
		chroma.KeywordReserved:          "#" + b.Base0E,
		chroma.KeywordType:              "#" + b.Base0C,
		chroma.Name:                     "#" + b.Base05,
		chroma.NameAttribute:            "#" + b.Base0D,
		chroma.NameBuiltin:              "#" + b.Base08,
		chroma.NameBuiltinPseudo:        "#" + b.Base05,
		chroma.NameClass:                "#" + b.Base0A,
		chroma.NameConstant:             "#" + b.Base09,
		chroma.NameDecorator:            "#" + b.Base09,
		chroma.NameEntity:               "#" + b.Base05,
		chroma.NameException:            "#" + b.Base05,
		chroma.NameFunction:             "#" + b.Base0D,
		chroma.NameLabel:                "#" + b.Base08,
		chroma.NameNamespace:            "#" + b.Base05,
		chroma.NameOther:                "#" + b.Base05,
		chroma.NameTag:                  "#" + b.Base0E,
		chroma.NameVariable:             "#" + b.Base08,
		chroma.NameVariableClass:        "#" + b.Base08,
		chroma.NameVariableGlobal:       "#" + b.Base08,
		chroma.NameVariableInstance:     "#" + b.Base08,
		chroma.Literal:                  "#" + b.Base05,
		chroma.LiteralDate:              "#" + b.Base05,
		chroma.LiteralString:            "#" + b.Base0B,
		chroma.LiteralStringBacktick:    "#" + b.Base0B,
		chroma.LiteralStringChar:        "#" + b.Base0B,
		chroma.LiteralStringDoc:         "#" + b.Base0B,
		chroma.LiteralStringDouble:      "#" + b.Base0B,
		chroma.LiteralStringEscape:      "#" + b.Base0B,
		chroma.LiteralStringHeredoc:     "#" + b.Base0B,
		chroma.LiteralStringInterpol:    "#" + b.Base0B,
		chroma.LiteralStringOther:       "#" + b.Base0B,
		chroma.LiteralStringRegex:       "#" + b.Base0B,
		chroma.LiteralStringSingle:      "#" + b.Base0B,
		chroma.LiteralStringSymbol:      "#" + b.Base0B,
		chroma.LiteralNumber:            "#" + b.Base09,
		chroma.LiteralNumberBin:         "#" + b.Base09,
		chroma.LiteralNumberFloat:       "#" + b.Base09,
		chroma.LiteralNumberHex:         "#" + b.Base09,
		chroma.LiteralNumberInteger:     "#" + b.Base09,
		chroma.LiteralNumberIntegerLong: "#" + b.Base09,
		chroma.LiteralNumberOct:         "#" + b.Base09,
		chroma.Operator:                 "#" + b.Base0E,
		chroma.OperatorWord:             "#" + b.Base0E,
		chroma.Punctuation:              "#" + b.Base05,
		chroma.Comment:                  "#" + b.Base03,
		chroma.CommentHashbang:          "#" + b.Base03,
		chroma.CommentMultiline:         "#" + b.Base03,
		chroma.CommentSingle:            "#" + b.Base03,
		chroma.CommentSpecial:           "#" + b.Base03,
		chroma.CommentPreproc:           "#" + b.Base03,
		chroma.Generic:                  "#" + b.Base05,
		chroma.GenericDeleted:           "#" + b.Base08,
		chroma.GenericEmph:              "underline #" + b.Base05,
		chroma.GenericError:             "#" + b.Base08,
		chroma.GenericHeading:           "bold #" + b.Base05,
		chroma.GenericInserted:          "bold #" + b.Base05,
		chroma.GenericOutput:            "#" + b.Base02,
		chroma.GenericPrompt:            "#" + b.Base05,
		chroma.GenericStrong:            "italic #" + b.Base05,
		chroma.GenericSubheading:        "bold #" + b.Base05,
		chroma.GenericTraceback:         "#" + b.Base05,
		chroma.GenericUnderline:         "underline",
		chroma.Text:                     "#" + b.Base05,
		chroma.TextWhitespace:           "#" + b.Base05,
	}

	return chroma.NewStyle(b.Scheme, b.Theme, entries)
}

func NewStyle(r io.Reader) (*chroma.Style, error) {
	dec := yaml.NewDecoder(r)
	b16 := &Base16{}
	if err := dec.Decode(b16); err != nil {
		return nil, err
	}

	return b16.toStyle()
}
