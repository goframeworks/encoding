package encoding

// Encoding 抽象了通用编解码的接口
//
// 用法:
//  + 通过工厂构造实例 (推荐)
//    encoding.NewFactory().JsonEncoding()
//
//  + 直接构造对应类型的实例
//    encoding.NewJsonEncoding()
type Encoding interface {
	Encode(v interface{}) ([]byte, error)
	Decode(b []byte, v interface{}) error
}
