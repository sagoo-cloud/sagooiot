package internal

type options struct {
	appKey    string
	appSecret string
	agentID   string
}

type Option func(c *options)

func AppKey(d string) Option {
	return func(opts *options) {
		opts.appKey = d
	}
}

func AppSecret(d string) Option {
	return func(opts *options) {
		opts.appSecret = d
	}
}
func AgentID(d string) Option {
	return func(opts *options) {
		opts.agentID = d
	}
}
