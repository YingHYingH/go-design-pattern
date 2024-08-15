package _23_abstract_factory

type IRuleConfigParser interface {
	Parser(data []byte)
}

type jsonRuleConfigParser struct {
}

func (j jsonRuleConfigParser) Parser(data []byte) {
	panic("")
}

type ISystemConfigParser interface {
	ParseSystem(data []byte)
}

type jsonSystemConfigParser struct {
}

func (j jsonSystemConfigParser) ParseSystem(data []byte) {
	panic("")
}

type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

type jsonConfigParserFactory struct {
}

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}
