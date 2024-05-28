// This code was AUTOGENERATED using the library.
// Please DO NOT EDIT THIS FILE.

package associated_token_account

type ProgramError interface {
	Code() int
	Error() string
}

var codeToErrorMap = make(map[int]ProgramError)
var nameToErrorMap = make(map[string]ProgramError)

func init() {
	codeToErrorMap[0] = new(InvalidOwnerError)
	nameToErrorMap["InvalidOwner"] = new(InvalidOwnerError)
}

func GetAssociatedTokenAccountErrorFromCode(code int) ProgramError {
	return codeToErrorMap[code]
}

func GetAssociatedTokenAccountErrorFromName(name string) ProgramError {
	return nameToErrorMap[name]
}

// InvalidOwnerError Error: 0 `Associated token account owner does not match address derivation`
type InvalidOwnerError struct{}

func (e InvalidOwnerError) Code() int {
	return 0
}
func (e InvalidOwnerError) Error() string {
	return "Associated token account owner does not match address derivation"
}