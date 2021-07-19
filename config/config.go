package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
}

type System struct {
	ApiMode       bool   `mapstructure:"api-mode" json:"api-mode" yaml:"api-mode"`                  // api模式
	Port          string `mapstructure:"port" json:"port" yaml:"port"`                              // api服务端口值
	TempFile      string `mapstructure:"temp-file" json:"temp-file" yaml:"temp-file"`               // 临时文件
	GenerateDir   string `mapstructure:"generate-dir" json:"generate-dir" yaml:"generate-dir"`      // 接口生成文件夹路径
	ConfigTxPath  string `mapstructure:"configtx-path" json:"configtx-path" yaml:"configtx-path"`   //生成配置文件路径
	ConfigTxFile  string `mapstructure:"configtx-file" json:"configtx-file" yaml:"configtx-file"`   // 生成配置文件名字
	CryptoConfig  bool   `mapstructure:"crypto-config" json:"crypto-config" yaml:"crypto-config"`   // 是否需要crypto-config的yaml文件
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                 // 环境值
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                              // 端口值
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                      // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`                   // Oss类型
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // 多点登录拦截
}

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                           // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                        // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                        // 日志前缀
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                 // 日志文件夹
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`                // 软链接名称
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`                 // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`  // 输出控制台
}
