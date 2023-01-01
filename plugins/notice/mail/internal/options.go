package internal

type options struct {
	mailHost string
	mailPort int
	mailUser string // 发件人
	mailPass string // 发件人密码
}

type Option func(c *options)

func MailHost(d string) Option {
	return func(opts *options) {
		opts.mailHost = d
	}
}

func MailPort(d int) Option {
	return func(opts *options) {
		opts.mailPort = d
	}
}
func MailUser(d string) Option {
	return func(opts *options) {
		opts.mailUser = d
	}
}

func MailPass(d string) Option {
	return func(opts *options) {
		opts.mailPass = d
	}
}
