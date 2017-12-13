package basics

type MemberOfTheFamily struct {
	Name     string
	Children []MemberOfTheFamily
}

func NumberOfDescendants(member MemberOfTheFamily) int {

	//nbOfDescendants = children
	nbOfDescendants := len(member.Children)
	for _, child := range member.Children {
		//children of children
		nbOfDescendants += NumberOfDescendants(child)
	}
	return nbOfDescendants
}
