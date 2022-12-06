package structure

type ProofOfWork struct {
	block *Block
}

func NewProofOfWork(block *Block) *ProofOfWork {
	return &ProofOfWork{block: block}
}

func (p *ProofOfWork) Run() ([]byte, int64) {
	return nil, 0
}
