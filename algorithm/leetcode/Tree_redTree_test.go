package leetcode

const (
	Black = false
	Read  = true
)

//红黑树
type RBTree struct {
	Root *node
	size int
}

type node struct {
	key, value interface{}
	color      bool
	left       *node
	right      *node
	parent     *node
}

func NewRBTree() *RBTree {
	return &RBTree{}
}

func (tree *RBTree) Put(key, value interface{}) {
	var insertNode *node
	if tree.Root == nil {
		tree.Root = &node{key: key, value: value, color: Black}
		insertNode = tree.Root
	}
	tree.insert1(insertNode)
	tree.size++
}

func (tree *RBTree) Get(key interface{}) (value interface{}, found bool) {
	return
}

func (tree *RBTree) insert1(node *node) {
	if node.parent == nil {
		node.color = Black
	} else {

	}
}
