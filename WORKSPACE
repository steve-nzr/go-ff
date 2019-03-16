workspace(name = "goff")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Load Go rules

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.16.5/rules_go-0.16.5.tar.gz"],
    sha256 = "7be7dc01f1e0afdba6c8eb2b43d2fa01c743be1b9273ab1eaf6c233df078d705",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

# Load Golang Docker rules

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "aed1c249d4ec8f703edddf35cbe9dfaca0b5f5ea6e4cd9e83e99f3b0d1136c3d",
    strip_prefix = "rules_docker-0.7.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.7.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

# Load K8s rules

http_archive(
    name = "io_bazel_rules_k8s",
    sha256 = "91fef3e6054096a8947289ba0b6da3cba559ecb11c851d7bdfc9ca395b46d8d8",
    strip_prefix = "rules_k8s-0.1",
    urls = ["https://github.com/bazelbuild/rules_k8s/archive/v0.1.tar.gz"],
)

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_repositories", "k8s_defaults")

k8s_repositories()

_CLUSTER = "docker-desktop"

_CONTEXT = _CLUSTER

k8s_defaults(
    name = "k8s_object",
    cluster = _CLUSTER,
    context = _CONTEXT,
    image_chroot = "localhost:5000/go-ff",
)

k8s_defaults(
    name = "k8s_deploy",
    cluster = _CLUSTER,
    context = _CONTEXT,
    image_chroot = "localhost:5000/go-ff",
    kind = "deployment",
)

# Load Go package discovery

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.16.0/bazel-gazelle-0.16.0.tar.gz"],
    sha256 = "7949fc6cc17b5b191103e97481cf8889217263acf52e00b560683413af204fcb",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "co_honnef_go_tools",
    commit = "3f1c8253044a",
    importpath = "honnef.co/go/tools",
)

go_repository(
    name = "com_github_ahmetb_go_linq",
    importpath = "github.com/ahmetb/go-linq",
    tag = "v3.0.0",
)

go_repository(
    name = "com_github_anmitsu_go_shlex",
    commit = "648efa622239",
    importpath = "github.com/anmitsu/go-shlex",
)

go_repository(
    name = "com_github_beorn7_perks",
    commit = "3a771d992973",
    importpath = "github.com/beorn7/perks",
)

go_repository(
    name = "com_github_bradfitz_go_smtpd",
    commit = "deb6d6237625",
    importpath = "github.com/bradfitz/go-smtpd",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    tag = "v0.3.1",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    tag = "v0.3.4",
)

go_repository(
    name = "com_github_coreos_go_systemd",
    commit = "c6f51f82210d",
    importpath = "github.com/coreos/go-systemd",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_denisenkom_go_mssqldb",
    commit = "041949b8d268",
    importpath = "github.com/denisenkom/go-mssqldb",
)

go_repository(
    name = "com_github_dustin_go_humanize",
    importpath = "github.com/dustin/go-humanize",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_erikstmartin_go_testdb",
    commit = "8d10e4a1bae5",
    importpath = "github.com/erikstmartin/go-testdb",
)

go_repository(
    name = "com_github_flynn_go_shlex",
    commit = "3f9db97f8568",
    importpath = "github.com/flynn/go-shlex",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    tag = "v1.4.7",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_gliderlabs_ssh",
    importpath = "github.com/gliderlabs/ssh",
    tag = "v0.1.1",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    tag = "v1.4.1",
)

go_repository(
    name = "com_github_gofrs_uuid",
    importpath = "github.com/gofrs/uuid",
    tag = "v3.2.0",
)

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_golang_geo",
    commit = "476085157cff",
    importpath = "github.com/golang/geo",
)

go_repository(
    name = "com_github_golang_glog",
    commit = "23def4e6c14b",
    importpath = "github.com/golang/glog",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_google_btree",
    commit = "4030bb1f1f0c",
    importpath = "github.com/google/btree",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    tag = "v0.2.0",
)

go_repository(
    name = "com_github_google_go_github",
    importpath = "github.com/google/go-github",
    tag = "v17.0.0",
)

go_repository(
    name = "com_github_google_go_querystring",
    importpath = "github.com/google/go-querystring",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_google_martian",
    importpath = "github.com/google/martian",
    tag = "v2.1.0",
)

go_repository(
    name = "com_github_google_pprof",
    commit = "3ea8567a2e57",
    importpath = "github.com/google/pprof",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    commit = "0766667cb4d1",
    importpath = "github.com/gopherjs/gopherjs",
)

go_repository(
    name = "com_github_gregjones_httpcache",
    commit = "9cad4c3443a7",
    importpath = "github.com/gregjones/httpcache",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    tag = "v1.5.0",
)

go_repository(
    name = "com_github_jellevandenhooff_dkim",
    commit = "f50fe3d243e1",
    importpath = "github.com/jellevandenhooff/dkim",
)

go_repository(
    name = "com_github_jinzhu_gorm",
    importpath = "github.com/jinzhu/gorm",
    tag = "v1.9.2",
)

go_repository(
    name = "com_github_jinzhu_inflection",
    commit = "04140366298a",
    importpath = "github.com/jinzhu/inflection",
)

go_repository(
    name = "com_github_jinzhu_now",
    importpath = "github.com/jinzhu/now",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_joho_godotenv",
    importpath = "github.com/joho/godotenv",
    tag = "v1.3.0",
)

go_repository(
    name = "com_github_jstemmer_go_junit_report",
    commit = "af01ea7f8024",
    importpath = "github.com/jstemmer/go-junit-report",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    tag = "v0.1.0",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    tag = "v1.1.3",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    tag = "v0.1.0",
)

go_repository(
    name = "com_github_lib_pq",
    importpath = "github.com/lib/pq",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    importpath = "github.com/mattn/go-sqlite3",
    tag = "v1.10.0",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    tag = "v1.0.1",
)

go_repository(
    name = "com_github_microcosm_cc_bluemonday",
    importpath = "github.com/microcosm-cc/bluemonday",
    tag = "v1.0.1",
)

go_repository(
    name = "com_github_neelance_astrewrite",
    commit = "99348263ae86",
    importpath = "github.com/neelance/astrewrite",
)

go_repository(
    name = "com_github_neelance_sourcemap",
    commit = "8c68805598ab",
    importpath = "github.com/neelance/sourcemap",
)

go_repository(
    name = "com_github_openzipkin_zipkin_go",
    importpath = "github.com/openzipkin/zipkin-go",
    tag = "v0.1.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    importpath = "github.com/prometheus/client_golang",
    tag = "v0.8.0",
)

go_repository(
    name = "com_github_prometheus_client_model",
    commit = "5c3871d89910",
    importpath = "github.com/prometheus/client_model",
)

go_repository(
    name = "com_github_prometheus_common",
    commit = "c7de2306084e",
    importpath = "github.com/prometheus/common",
)

go_repository(
    name = "com_github_prometheus_procfs",
    commit = "05ee40e3a273",
    importpath = "github.com/prometheus/procfs",
)

go_repository(
    name = "com_github_russross_blackfriday",
    importpath = "github.com/russross/blackfriday",
    tag = "v1.5.2",
)

go_repository(
    name = "com_github_sergi_go_diff",
    importpath = "github.com/sergi/go-diff",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_shurcool_component",
    commit = "f88ec8f54cc4",
    importpath = "github.com/shurcooL/component",
)

go_repository(
    name = "com_github_shurcool_events",
    commit = "410e4ca65f48",
    importpath = "github.com/shurcooL/events",
)

go_repository(
    name = "com_github_shurcool_github_flavored_markdown",
    commit = "2122de532470",
    importpath = "github.com/shurcooL/github_flavored_markdown",
)

go_repository(
    name = "com_github_shurcool_go",
    commit = "9e1955d9fb6e",
    importpath = "github.com/shurcooL/go",
)

go_repository(
    name = "com_github_shurcool_go_goon",
    commit = "37c2f522c041",
    importpath = "github.com/shurcooL/go-goon",
)

go_repository(
    name = "com_github_shurcool_gofontwoff",
    commit = "29b52fc0a18d",
    importpath = "github.com/shurcooL/gofontwoff",
)

go_repository(
    name = "com_github_shurcool_gopherjslib",
    commit = "feb6d3990c2c",
    importpath = "github.com/shurcooL/gopherjslib",
)

go_repository(
    name = "com_github_shurcool_highlight_diff",
    commit = "09bb4053de1b",
    importpath = "github.com/shurcooL/highlight_diff",
)

go_repository(
    name = "com_github_shurcool_highlight_go",
    commit = "98c3abbbae20",
    importpath = "github.com/shurcooL/highlight_go",
)

go_repository(
    name = "com_github_shurcool_home",
    commit = "80b7ffcb30f9",
    importpath = "github.com/shurcooL/home",
)

go_repository(
    name = "com_github_shurcool_htmlg",
    commit = "d01228ac9e50",
    importpath = "github.com/shurcooL/htmlg",
)

go_repository(
    name = "com_github_shurcool_httperror",
    commit = "86b7830d14cc",
    importpath = "github.com/shurcooL/httperror",
)

go_repository(
    name = "com_github_shurcool_httpfs",
    commit = "809beceb2371",
    importpath = "github.com/shurcooL/httpfs",
)

go_repository(
    name = "com_github_shurcool_httpgzip",
    commit = "b1c53ac65af9",
    importpath = "github.com/shurcooL/httpgzip",
)

go_repository(
    name = "com_github_shurcool_issues",
    commit = "6292fdc1e191",
    importpath = "github.com/shurcooL/issues",
)

go_repository(
    name = "com_github_shurcool_issuesapp",
    commit = "048589ce2241",
    importpath = "github.com/shurcooL/issuesapp",
)

go_repository(
    name = "com_github_shurcool_notifications",
    commit = "627ab5aea122",
    importpath = "github.com/shurcooL/notifications",
)

go_repository(
    name = "com_github_shurcool_octicon",
    commit = "fa4f57f9efb2",
    importpath = "github.com/shurcooL/octicon",
)

go_repository(
    name = "com_github_shurcool_reactions",
    commit = "f2e0b4ca5b82",
    importpath = "github.com/shurcooL/reactions",
)

go_repository(
    name = "com_github_shurcool_sanitized_anchor_name",
    commit = "86672fcb3f95",
    importpath = "github.com/shurcooL/sanitized_anchor_name",
)

go_repository(
    name = "com_github_shurcool_users",
    commit = "49c67e49c537",
    importpath = "github.com/shurcooL/users",
)

go_repository(
    name = "com_github_shurcool_webdavfs",
    commit = "18c3829fa133",
    importpath = "github.com/shurcooL/webdavfs",
)

go_repository(
    name = "com_github_sourcegraph_annotate",
    commit = "f4cad6c6324d",
    importpath = "github.com/sourcegraph/annotate",
)

go_repository(
    name = "com_github_sourcegraph_syntaxhighlight",
    commit = "bd320f5d308e",
    importpath = "github.com/sourcegraph/syntaxhighlight",
)

go_repository(
    name = "com_github_streadway_amqp",
    commit = "14f78b41ce6d",
    importpath = "github.com/streadway/amqp",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    tag = "v1.2.2",
)

go_repository(
    name = "com_github_tarm_serial",
    commit = "98f6abe2eb07",
    importpath = "github.com/tarm/serial",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    tag = "v0.37.0",
)

go_repository(
    name = "com_shuralyov_dmitri_app_changes",
    commit = "0a106ad413e3",
    importpath = "dmitri.shuralyov.com/app/changes",
)

go_repository(
    name = "com_shuralyov_dmitri_html_belt",
    commit = "f7d459c86be0",
    importpath = "dmitri.shuralyov.com/html/belt",
)

go_repository(
    name = "com_shuralyov_dmitri_service_change",
    commit = "a85b471d5412",
    importpath = "dmitri.shuralyov.com/service/change",
)

go_repository(
    name = "com_shuralyov_dmitri_state",
    commit = "28bcc343414c",
    importpath = "dmitri.shuralyov.com/state",
)

go_repository(
    name = "com_sourcegraph_sourcegraph_go_diff",
    importpath = "sourcegraph.com/sourcegraph/go-diff",
    tag = "v0.5.0",
)

go_repository(
    name = "com_sourcegraph_sqs_pbtypes",
    commit = "d3ebe8f20ae4",
    importpath = "sourcegraph.com/sqs/pbtypes",
)

go_repository(
    name = "in_gopkg_check_v1",
    commit = "20d25e280405",
    importpath = "gopkg.in/check.v1",
)

go_repository(
    name = "in_gopkg_inf_v0",
    importpath = "gopkg.in/inf.v0",
    tag = "v0.9.1",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    tag = "v2.2.1",
)

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    tag = "v0.18.0",
)

go_repository(
    name = "org_apache_git_thrift_git",
    commit = "2566ecd5d999",
    importpath = "git.apache.org/thrift.git",
)

go_repository(
    name = "org_go4",
    commit = "417644f6feb5",
    importpath = "go4.org",
)

go_repository(
    name = "org_go4_grpc",
    commit = "11d0a25b4919",
    importpath = "grpc.go4.org",
)

go_repository(
    name = "org_golang_google_api",
    importpath = "google.golang.org/api",
    tag = "v0.1.0",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    tag = "v1.4.0",
)

go_repository(
    name = "org_golang_google_genproto",
    commit = "b5d61aea6440",
    importpath = "google.golang.org/genproto",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    tag = "v1.19.0",
)

go_repository(
    name = "org_golang_x_build",
    commit = "041ab4dc3f9d",
    importpath = "golang.org/x/build",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "a1f597ede03a",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_exp",
    commit = "509febef88a4",
    importpath = "golang.org/x/exp",
)

go_repository(
    name = "org_golang_x_lint",
    commit = "5b3e6a55c961",
    importpath = "golang.org/x/lint",
)

go_repository(
    name = "org_golang_x_net",
    commit = "3a22650c66bd",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_oauth2",
    commit = "d668ce993890",
    importpath = "golang.org/x/oauth2",
)

go_repository(
    name = "org_golang_x_perf",
    commit = "6e6d33e29852",
    importpath = "golang.org/x/perf",
)

go_repository(
    name = "org_golang_x_sync",
    commit = "e225da77a7e6",
    importpath = "golang.org/x/sync",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "d0b11bdaac8a",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    tag = "v0.3.1-0.20180807135948-17ff2d5776d2",
)

go_repository(
    name = "org_golang_x_time",
    commit = "85acf8d2951c",
    importpath = "golang.org/x/time",
)

go_repository(
    name = "org_golang_x_tools",
    commit = "f727befe758c",
    importpath = "golang.org/x/tools",
)
