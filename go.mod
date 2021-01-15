module github.com/vmware-tanzu-private/core

go 1.13

require (
	cloud.google.com/go/storage v1.0.0
	github.com/AlecAivazis/survey/v2 v2.1.1
	github.com/Jeffail/gabs v1.4.0
	github.com/MakeNowJust/heredoc v1.0.0
	github.com/Netflix/go-expect v0.0.0-20190729225929-0e00d9168667 // indirect
	github.com/adrg/xdg v0.2.1
	github.com/amenzhinsky/go-memexec v0.3.0
	github.com/aunum/log v0.0.0-20200821225356-38d2e2c8b489
	github.com/caarlos0/spin v1.1.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fabriziopandini/capi-conditions v0.0.0-20201102133039-7eb142d1b6d6
	github.com/fatih/color v1.10.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v0.1.0
	github.com/gobwas/glob v0.2.3
	github.com/golang/protobuf v1.4.2
	github.com/gosuri/uitable v0.0.4
	github.com/google/go-containerregistry v0.0.0-20190412005658-1d38b9cfdb9d
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/hinshun/vt10x v0.0.0-20180809195222-d55458df857c // indirect
	github.com/imdario/mergo v0.3.9
	github.com/jedib0t/go-pretty v4.3.0+incompatible
	github.com/jeremywohl/flatten v1.0.1
	github.com/k14s/imgpkg v0.2.0
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/olekukonko/tablewriter v0.0.4
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.3
	github.com/pkg/errors v0.9.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	github.com/vmware-tanzu-private/tkg-cli v1.3.0-pre-alpha-1.0.20210114003033-285a8c9131d4
	github.com/vmware-tanzu-private/tkg-providers v1.3.0-pre-alpha-1.0.20210113202657-eb07b4e0558d
	go.opencensus.io v0.22.2 // indirect
	go.uber.org/multierr v1.1.0
	golang.org/x/mod v0.3.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/api v0.13.0
	google.golang.org/grpc v1.26.0
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.17.11
	k8s.io/apimachinery v0.17.11
	k8s.io/client-go v0.17.11
	sigs.k8s.io/cluster-api v0.3.10
	sigs.k8s.io/controller-runtime v0.5.11
)