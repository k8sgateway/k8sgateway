module github.com/solo-io/solo-projects

go 1.18

require (
	github.com/Masterminds/goutils v1.1.1
	github.com/avast/retry-go v2.4.3+incompatible
	github.com/aws/aws-sdk-go v1.42.15
	github.com/deislabs/oras v0.11.1
	github.com/docker/distribution v2.7.1+incompatible
	github.com/envoyproxy/go-control-plane v0.10.2-0.20211130161932-f62def555c97
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/fgrosse/zaptest v1.1.0
	github.com/form3tech-oss/jwt-go v3.2.3+incompatible
	github.com/gertd/go-pluralize v0.1.1
	github.com/getkin/kin-openapi v0.80.0
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/gin-gonic/gin v1.7.7
	github.com/go-logr/logr v0.4.0
	github.com/go-logr/zapr v0.4.0
	github.com/go-openapi/inflect v0.19.0
	github.com/go-openapi/swag v0.19.15
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-test/deep v1.0.7
	github.com/gobuffalo/packr v1.30.1
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/google/wire v0.4.0
	github.com/graphql-go/graphql v0.8.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/iancoleman/strcase v0.1.3
	github.com/jhump/protoreflect v1.6.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/keybase/go-ps v0.0.0-20190827175125-91aafc93ba19
	github.com/mitchellh/hashstructure v1.0.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.18.1
	github.com/opencontainers/go-digest v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/prometheus v2.5.0+incompatible
	github.com/rotisserie/eris v0.5.0
	github.com/sirupsen/logrus v1.8.1
	github.com/solo-io/anyvendor v0.0.4
	github.com/solo-io/cli-kit v0.2.1
	github.com/solo-io/ext-auth-plugins v0.2.1
	github.com/solo-io/ext-auth-service v0.19.20
	github.com/solo-io/external-apis v0.1.11
	github.com/solo-io/gloo v1.12.0-beta11
	github.com/solo-io/go-list-licenses v0.1.4
	github.com/solo-io/go-utils v0.21.25
	github.com/solo-io/k8s-utils v0.1.0
	github.com/solo-io/licensing v0.2.0
	github.com/solo-io/protoc-gen-ext v0.0.16
	github.com/solo-io/protoc-gen-openapi v0.0.4
	github.com/solo-io/rate-limiter v0.6.2
	github.com/solo-io/skv2 v0.22.11
	github.com/solo-io/skv2-enterprise v0.2.8
	github.com/solo-io/solo-apis v0.0.0-20220506170656-df8d0234623b
	github.com/solo-io/solo-kit v0.27.0
	github.com/solo-io/wasm-kit v0.1.3
	github.com/solo-io/wasm/tools/wasme/pkg v0.0.0-20201021213306-77f82bdc3cc3
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	go.opencensus.io v0.23.0
	go.uber.org/zap v1.19.1
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3
	golang.org/x/net v0.0.0-20211205041911-012df41ee64c
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/tools v0.1.10
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/square/go-jose.v2 v2.6.0
	gopkg.in/yaml.v1 v1.0.0-20140924161607-9f9df34309c0
	helm.sh/helm/v3 v3.7.1
	istio.io/pkg v0.0.0-20211123161558-1e5d0c4ee827
	istio.io/tools v0.0.0-20210420211536-9c0f48df3262
	k8s.io/api v0.22.4
	k8s.io/apiextensions-apiserver v0.22.4
	k8s.io/apimachinery v0.22.4
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.22.4
	sigs.k8s.io/controller-runtime v0.10.3
	sigs.k8s.io/yaml v1.3.0
)

require (
	cloud.google.com/go v0.98.0 // indirect
	cloud.google.com/go/datastore v1.3.0 // indirect
	cloud.google.com/go/logging v1.4.2 // indirect
	cloud.google.com/go/storage v1.18.2 // indirect
	contrib.go.opencensus.io/exporter/prometheus v0.4.0 // indirect
	cuelang.org/go v0.3.2 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.20 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.15 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/MakeNowJust/heredoc v0.0.0-20171113091838-e9091a26100e // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/Masterminds/sprig/v3 v3.2.2 // indirect
	github.com/Masterminds/squirrel v1.5.2 // indirect
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/Microsoft/hcsshim v0.8.21 // indirect
	github.com/OneOfOne/xxhash v1.2.8 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20210428141323-04723f9f07d7 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/acomagu/bufpipe v1.0.3 // indirect
	github.com/armon/go-metrics v0.3.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200428143746-21a406dcc535 // indirect
	github.com/aws/aws-app-mesh-controller-for-k8s v1.1.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bshuster-repo/logrus-logstash-hook v1.0.2 // indirect
	github.com/bugsnag/bugsnag-go v1.5.0 // indirect
	github.com/bytecodealliance/wasmtime-go v0.27.0 // indirect
	github.com/census-instrumentation/opencensus-proto v0.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/chai2010/gettext-go v0.0.0-20160711120539-c6fed771bfd5 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/cncf/udpa/go v0.0.0-20210930031921-04548b0d99d4 // indirect
	github.com/cncf/xds/go v0.0.0-20211011173535-cb28da3451f1 // indirect
	github.com/cockroachdb/apd/v2 v2.0.1 // indirect
	github.com/containerd/cgroups v1.0.1 // indirect
	github.com/containerd/containerd v1.5.1 // indirect
	github.com/cyphar/filepath-securejoin v0.2.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/daviddengcn/go-colortext v0.0.0-20160507010035-511bcaf42ccd // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/docker/cli v20.10.10+incompatible // indirect
	github.com/docker/docker v20.10.10+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.6.4 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/emicklei/go-restful v2.11.1+incompatible // indirect
	github.com/emicklei/proto v1.6.15 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/evanphx/json-patch v4.11.0+incompatible // indirect
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fsnotify/fsnotify v1.5.3 // indirect
	github.com/fvbommel/sortorder v1.0.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.3.1 // indirect
	github.com/go-git/go-git/v5 v5.4.1 // indirect
	github.com/go-kit/log v0.1.0 // indirect
	github.com/go-ldap/ldap v3.0.3+incompatible // indirect
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-openapi/analysis v0.19.5 // indirect
	github.com/go-openapi/errors v0.19.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.5 // indirect
	github.com/go-openapi/loads v0.19.4 // indirect
	github.com/go-openapi/spec v0.19.6 // indirect
	github.com/go-openapi/strfmt v0.19.5 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/gobuffalo/envy v1.8.1 // indirect
	github.com/gobuffalo/packd v1.0.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gogo/googleapis v1.4.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/google/go-github/v32 v32.0.0 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/google/subcommands v1.0.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/gax-go v2.0.2+incompatible // indirect
	github.com/googleapis/gax-go/v2 v2.1.1 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/goph/emperror v0.17.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gosuri/uitable v0.0.4 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/hashicorp/consul/api v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-immutable-radix v1.1.0 // indirect
	github.com/hashicorp/go-retryablehttp v0.6.8 // indirect
	github.com/hashicorp/go-rootcerts v1.0.1 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/go-version v1.3.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/serf v0.8.5 // indirect
	github.com/hashicorp/vault/api v1.0.5-0.20191108163347-bdd38fca2cff // indirect
	github.com/hashicorp/vault/sdk v0.1.14-0.20191112033314-390e96e22eb2 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/ilackarms/protoc-gen-doc v1.0.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/imroc/req v0.3.0 // indirect
	github.com/inconshreveable/go-update v0.0.0-20160112193335-8152e7eb6ccf // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jackc/pgx v3.6.2+incompatible // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jmoiron/sqlx v1.3.1 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/juju/ansiterm v0.0.0-20180109212912-720a0952cc2a // indirect
	github.com/k0kubun/pp v3.0.1+incompatible // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lib/pq v1.10.0 // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/lithammer/dedent v1.1.0 // indirect
	github.com/lunixbochs/vtclean v0.0.0-20180621232353-2d01aacdc34a // indirect
	github.com/lyft/protoc-gen-star v0.6.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/manifoldco/promptui v0.7.0 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mattn/go-zglob v0.0.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/moby/term v0.0.0-20210610120745-9d4ed1856297 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/mpvl/unique v0.0.0-20150818121801-cbe035fff7de // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/open-policy-agent/opa v0.29.4 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/prometheus/statsd_exporter v0.21.0 // indirect
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/pseudomuto/protoc-gen-doc v1.4.1 // indirect
	github.com/pseudomuto/protokit v0.2.0 // indirect
	github.com/radovskyb/watcher v1.0.7 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/rubenv/sql-migrate v0.0.0-20210614095031-55d5740dbbcc // indirect
	github.com/russross/blackfriday v1.5.2 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/sahilm/fuzzy v0.1.0 // indirect
	github.com/schollz/progressbar/v3 v3.7.4 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/servicemeshinterface/smi-sdk-go v0.4.1 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/solo-io/cue v0.4.1 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.8.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/tredoe/osutil v0.0.0-20191018075336-e272fdda81c8 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/xlab/treeprint v0.0.0-20181112141820-a009c3971eca // indirect
	github.com/yashtewari/glob-intersection v0.0.0-20180916065949-5c77d914dd0b // indirect
	github.com/yvasiyarov/newrelic_platform_go v0.0.0-20160601141957-9c099fbc30e9 // indirect
	go.mongodb.org/mongo-driver v1.8.2 // indirect
	go.opentelemetry.io/proto/otlp v0.7.0 // indirect
	go.starlark.net v0.0.0-20211013185944-b0039bd2cfe3 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	gocloud.dev v0.18.0 // indirect
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sys v0.0.0-20220422013727-9388b58f7150 // indirect
	golang.org/x/term v0.0.0-20210220032956-6a3ed077a48d // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gomodules.xyz/jsonpatch/v2 v2.2.0 // indirect
	google.golang.org/api v0.61.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20211129164237-f09f9a12af12 // indirect
	gopkg.in/AlecAivazis/survey.v1 v1.8.7 // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/gorp.v1 v1.7.2 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	istio.io/api v0.0.0-20211130000055-e9d1de18c118 // indirect
	istio.io/client-go v1.12.0-alpha.5.0.20211124144007-2c98ed70c967 // indirect
	istio.io/gogo-genproto v0.0.0-20211111171331-7926fce33bb3 // indirect
	istio.io/istio v0.0.0-20211130191315-d9963eecf782 // indirect
	k8s.io/apiserver v0.22.4 // indirect
	k8s.io/cli-runtime v0.22.4 // indirect
	k8s.io/component-base v0.22.4 // indirect
	k8s.io/component-helpers v0.22.4 // indirect
	k8s.io/gengo v0.0.0-20211129171323-c02415ce4185 // indirect
	k8s.io/klog/v2 v2.10.0 // indirect
	k8s.io/kube-openapi v0.0.0-20211115234752-e816edb12b65 // indirect
	k8s.io/kubectl v0.22.4 // indirect
	k8s.io/metrics v0.22.4 // indirect
	k8s.io/utils v0.0.0-20211208161948-7d6a63dca704 // indirect
	knative.dev/pkg v0.0.0-20211206113427-18589ac7627e // indirect
	sigs.k8s.io/kustomize/api v0.8.11 // indirect
	sigs.k8s.io/kustomize/kustomize/v4 v4.2.0 // indirect
	sigs.k8s.io/kustomize/kyaml v0.11.0 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.1.2 // indirect
)

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/apache/thrift => github.com/apache/thrift v0.14.0
	// used by github.com/deislabs/oras v0.10.0
	github.com/containerd/containerd => github.com/containerd/containerd v1.4.13
	// pin to the jwt-go fork to fix CVE.
	// using the pseudo version of github.com/form3tech-oss/jwt-go@v3.2.3 instead of the version directly,
	// to avoid error about it being used for two different module paths
	github.com/dgrijalva/jwt-go => github.com/form3tech-oss/jwt-go v0.0.0-20210511163231-5b2d2b5f6c34
	github.com/docker/docker => github.com/moby/moby v20.10.14+incompatible
	github.com/golang/mock v1.4.4-0.20200406172829-6d816de489c1 => github.com/golang/mock v1.4.3

	github.com/opencontainers/go-digest => github.com/opencontainers/go-digest v1.0.0-rc1
	// skv2 uses a newer version than the imported solo-kit version which causes issues. Replaces the version with the solo-kit version
	github.com/pseudomuto/protoc-gen-doc => github.com/pseudomuto/protoc-gen-doc v1.0.0
	github.com/sclevine/agouti => github.com/yuval-k/agouti v0.0.0-20190109124522-0e71d6bad483

	// Gloo-fed code currently relies on older skv2: https://github.com/solo-io/solo-projects/issues/3540
	github.com/solo-io/skv2 => github.com/solo-io/skv2 v0.21.6
	go.opentelemetry.io/otel => go.opentelemetry.io/otel v0.11.0

	// Lock sys package to fix issue https://github.com/golang/go/issues/49219
	golang.org/x/sys => golang.org/x/sys v0.0.0-20220412211240-33da011f77ad
	golang.org/x/tools => golang.org/x/tools v0.0.0-20210106214847-113979e3529a
	// Pinning helm v3.6.0 because helm 3.7+ depends on containerd/containerd v1.5+, which introduces a breaking change
	// that breaks our other dependencies.
	helm.sh/helm/v3 => helm.sh/helm/v3 v3.6.0

	k8s.io/api => k8s.io/api v0.22.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.22.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.22.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.22.4
	k8s.io/client-go => k8s.io/client-go v0.22.4

	// klog is likely unused, but if it is we want to use this fork
	// see https://github.com/solo-io/gloo/pull/1880
	k8s.io/klog => github.com/stefanprodan/klog v0.0.0-20190418165334-9cbb78b20423
	k8s.io/kubectl => k8s.io/kubectl v0.22.4

	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.9.7

)

exclude (
	// Exclude pre-go-mod kubernetes tags, because they are older
	// than v0.x releases but are picked when updating dependencies.
	k8s.io/client-go v1.4.0
	k8s.io/client-go v1.5.0
	k8s.io/client-go v1.5.1
	k8s.io/client-go v1.5.2
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/client-go v2.0.0+incompatible
	k8s.io/client-go v2.0.0-alpha.1+incompatible
	k8s.io/client-go v3.0.0+incompatible
	k8s.io/client-go v3.0.0-beta.0+incompatible
	k8s.io/client-go v4.0.0+incompatible
	k8s.io/client-go v4.0.0-beta.0+incompatible
	k8s.io/client-go v5.0.0+incompatible
	k8s.io/client-go v5.0.1+incompatible
	k8s.io/client-go v6.0.0+incompatible
	k8s.io/client-go v7.0.0+incompatible
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/client-go v9.0.0+incompatible
	k8s.io/client-go v9.0.0-invalid+incompatible
)
