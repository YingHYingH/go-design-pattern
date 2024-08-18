package _8_adapter

type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}
type AWS struct {
}

func (c *AWS) RunInstance() error {
	return nil
}

type AWSAdapter struct {
	Client AWS
}

func (a *AWSAdapter) CreateServer(cpu, mem float64) error {
	return a.Client.RunInstance()
}

type Aliyun struct {
}

func (c *Aliyun) CreateServer(cpu, mem float64) error {
	return nil
}

type AliyunAdapter struct {
	Client Aliyun
}

func (a *AliyunAdapter) CreateServer(cpu, mem float64) error {
	return a.Client.CreateServer(cpu, mem)
}
