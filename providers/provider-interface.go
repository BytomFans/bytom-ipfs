package providers

type ProviderInterface interface {
	SendRequest(params interface{}) error
	Close() error
}
