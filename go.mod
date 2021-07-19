module crypto_webtool

go 1.16

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/Shopify/sarama v1.29.0 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20210208195552-ff826a37aa15 // indirect
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-gonic/gin v1.7.2
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/go-version v1.3.0 // indirect
	github.com/hyperledger/fabric-amcl v0.0.0-20210603140002-2670f91851c8
	github.com/hyperledger/fabric-config v0.1.0 // indirect
	github.com/hyperledger/fabric-protos-go v0.0.0-20210528200356-82833ecdac31
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/miekg/pkcs11 v1.0.3
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v0.0.0-20150908122457-1967d93db724
	github.com/spf13/viper180 v1.8.0
	github.com/stretchr/testify v1.7.0
	github.com/sykesm/zap-logfmt v0.0.4
	go.uber.org/zap v1.17.0
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b
	google.golang.org/grpc v1.38.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/cheggaaa/pb.v1 v1.0.28
	gopkg.in/yaml.v2 v2.4.0
)

//replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.3
replace github.com/mitchellh/mapstructure v1.4.1 => github.com/mitchellh/mapstructure v1.2.2
replace github.com/spf13/viper180 => github.com/spf13/viper v1.8.0
replace github.com/spf13/viper => github.com/spf13/viper v0.0.0-20150908122457-1967d93db724
