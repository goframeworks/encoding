package encoding

// EncodingFactory 是 Encoding 接口的工厂接口
type EncodingFactory interface {
	// JsonEncoding 构造基于json的encoding实例
	JsonEncoding() Encoding
}

func NewFactory() EncodingFactory {
	return newEncodingFactory()
}

// encodingFactory 实现了 EncodingFactory接口
type encodingFactory struct{}

func newEncodingFactory() *encodingFactory {
	return &encodingFactory{}
}

func (f *encodingFactory) JsonEncoding() Encoding {
	return NewJsonEncoding()
}
