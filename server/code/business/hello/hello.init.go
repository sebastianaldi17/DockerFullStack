package hello

type Business struct {
	helloData helloData
}

func New(helloData helloData) *Business {
	hello := Business{
		helloData: helloData,
	}
	return &hello
}
