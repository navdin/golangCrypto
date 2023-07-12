package blockchain

type Block struct {
	data     string
	hash     []byte
	prevHash []byte
}

type Blockchain struct {
	chain []Block
}
