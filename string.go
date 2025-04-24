package puzzle

var StringConverter = newConverter(
	func(entry *Entry[string], stringValue string) error {
		*entry.ValueP = stringValue
		entry.Value = stringValue
		return nil
	},
)
