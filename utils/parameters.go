package utils

// Constants with CPF parameters
const (
	CpfLength         = 11
	CpfValidatorRegex = `^([\d]{3})([\d]{3})([\d]{3})([\d]{2})$`
	CpfFormatterRegex = "$1.$2.$3-$4"
	CpfNumberOfDigits = 2
)

