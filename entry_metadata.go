package puzzle

import (
	"strings"
	"unicode"
)

type EntryMetadata struct {
	Description    string
	FlagName       string
	ShortFlagName  string
	EnvName        string
	Format         string
	SliceSeparator string
	IsConfigFile   bool
}

type MetadataOption = func(*EntryMetadata)

func keyToFlagNameMapping(r rune) rune {
	// return r if r is alphanumeric
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return r
	}
	// otherwise, return '-'
	return '-'
}

func GenerateEnvName(key string) string {
	return strings.ToUpper(
		strings.ReplaceAll(
			strings.Map(keyToFlagNameMapping, key), "-", "_"))
}

func GenerateFlagName(key string) string {
	return strings.ToLower(
		strings.Trim(
			strings.Map(keyToFlagNameMapping, key), "-"))
}

func newMetadataFromEntry(key string) *EntryMetadata {
	return &EntryMetadata{
		Description:    "",
		FlagName:       GenerateFlagName(key),
		ShortFlagName:  "",
		EnvName:        GenerateEnvName(key),
		Format:         "",
		SliceSeparator: ",",
		IsConfigFile:   false,
	}
}

func WithDescription(description string) MetadataOption {
	return func(metadata *EntryMetadata) {
		metadata.Description = description
	}
}

func WithEnvName(name string) MetadataOption {
	return func(metadata *EntryMetadata) {
		metadata.EnvName = name
	}
}

func WithoutEnv() MetadataOption {
	return WithEnvName("")
}

func WithFlagName(name string) MetadataOption {
	return func(metadata *EntryMetadata) {
		metadata.FlagName = name
	}
}

func WithoutFlagName() MetadataOption {
	return WithFlagName("")
}

func WithShortFlagName(name string) MetadataOption {
	return func(metadata *EntryMetadata) {
		metadata.ShortFlagName = name
	}
}

func WithSliceSeparator(sep string) MetadataOption {
	return func(metadata *EntryMetadata) {
		metadata.SliceSeparator = sep
	}
}

func WithFormat(format string) MetadataOption {
	return func(metadata *EntryMetadata) {
		metadata.Format = format
	}
}
