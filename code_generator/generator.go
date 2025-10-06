package codegenerator


type IVisitor interface{
	Visit()string
}

type BodyVisitor struct{}
type BoldVisitor struct{}
type EmphasisVisitor struct{}
type TextVisitor struct{}
type SentenceVisitor struct{
	visitor_map map[string]IVisitor{
	"BOLD":BoldVisitor,

}
}

// func (bv BodyVisitor) Visit() (string, error) {
//
// }
