package sub

// Constants are sorted to the top of the code documentation.
//
// Formatting
//
// Because it is declared on its own and not in a const() block, Applename gets
// full formatting support.
const AppleName = "apple"

// This text is a doc comment on the const block as a whole.  Unlike the
// comments inside the block, this text has full formatting.
//
// This can be a nice way to write a simgle comment about a group of variables.
const (
	// PearName is an exported constant used to create pear foods. Docs on
	// constants and variables enclosed in a block like this one get printed out
	// as-is.
	PearName = "pear"

	// PineappleName is the name of a pineapple.
	PineappleName = "pineapple"
)

// Apple is a global variable in the package.
//
// Because apple is declared on its own, rather in a block of var (), its
// documentation gets full formatting support.
var Apple = NewFood(AppleName)

// This is a block of variables with a formatted comment.
//
// Pretty nice.
var (
	// Pear is another global variable in the package.
	Pear = NewFood(PearName)
)

// Food is a type to show how type documentation works.
//
// Formatting
//
// You always get full formatting on type, function, and method docs.
//
// Below is the auto-embedded example from the function ExampleFood in
// more_examples_test.go.
type Food struct {
	// Name holds the common name of the food.
	//
	// Alas, no formatting here.
	Name string

	// weight is an unexported field that won't show up in documentation.
	weight int
}

// NewFood creates a new Food with the given name and a calculated weight.
//
// Godoc sorts "constructor" style functions like this, that return a type
// defined in the package, under the type itself.
//
// Note that you do not need to (and should not) make separate lines describing
// each parameter to the function.  Simply refer to the parameters by their name
// when you need to reference them in the rest of the documentation.
func NewFood(name string) Food {
	return Food{name, len(name)}
}

// Weight reports the weight of the food in letters.
//
// When decribing a function without parameters that simply returns data, it is
// traditional to write the doc saying the function reports foo, rather than
// *returns* foo.
//
// The example below intentionally fails during go test, to show you what that
// looks like.  The test will report:
//
//   --- FAIL: ExampleFood_Weight (4.933us)
//   got:
//   9
//   want:
//   10
//   FAIL
//   exit status 1
func (f Food) Weight() int {
	return f.weight
}

// Zap consumes the food and prints out a nasty message.
//
// This is just another top level function to show how the alphabetical sorting
// works.
func Zap(f Food) {
	fmt.Printf("I hate %ss!", f.Name)

}

// Eat consumes the food and prints out a nice message.
//
// This is an example of a function that isn't a "constructor" style function,
// so it will be sorted at the top level of the documentation.
//
// Functions get sorted before types.
func Eat(f Food) {
	fmt.Printf("I love %ss!", f.Name)
}
