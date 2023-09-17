package entity

type ListOrMap interface{}

type ListOrString interface{}

type ListStringOrListObject interface{}

type StringOrMap interface{}

type StringOrList interface{}

type ByteValue interface{}

type Permission interface{}

type DockerComposeConfigLongSyntax struct {
	Source string     `yaml:"source,omitempty" json:"source,omitempty"`
	Target string     `yaml:"target,omitempty" json:"target,omitempty"`
	Uid    string     `yaml:"uid,omitempty" json:"uid,omitempty"`
	Gid    string     `yaml:"gid,omitempty" json:"gid,omitempty"`
	Mode   Permission `yaml:"mode,omitempty" json:"mode,omitempty"`
}
type DockerComposeConfigShortSyntax []string

// [Docker Compose Config Definition](https://github.com/compose-spec/compose-spec/blob/master/05-services.md#configs)
//
// # Need function to parse to short or long syntax
//
// Short syntax will be an array of string
type DockerComposeConfig interface{}

type DockerComposeSecretShortSyntax map[string]string
type DockerComposeSecretLongSyntax struct {
	Source string     `yaml:"source,omitempty" json:"source,omitempty"`
	Target string     `yaml:"targe,omitempty" json:"target,omitempty"`
	Uid    string     `yaml:"uid,omitempty" json:"uid,omitempty"`
	Gid    string     `yaml:"gid,omitempty" json:"gid,omitempty"`
	Mode   Permission `yaml:"mode,omitempty" json:"mode,omitempty"`
}

type TimeValue interface{}

// [Docker Compose Secret Definition](https://github.com/compose-spec/compose-spec/blob/master/05-services.md#secrets)
//
// Need functions to check and convert to Short syntax or Long syntax
// **Note: the structure of the syntax above is the item structure**
// If use sort syntax, the secret will be an object, in contrast the secret will be an array if use long syntax
type DockerComposeSecret interface{}

// [Docker Compose Build Configuration Definition](https://github.com/compose-spec/compose-spec/blob/master/build.md)
type DockerComposeServiceBuildObject struct {
	Context            string              `default:"." yaml:"context,omitempty" json:"context,omitempty"`
	Dockerfile         string              `default:"." yaml:"dockerfile,omitempty" json:"dockerfile,omitempty"`
	DockerfileInline   string              `yaml:"dockerfile_inline,omitempty" json:"dockerfile_inline,omitempty"`
	Args               ListOrMap           `yaml:"args,omitempty" json:"args,omitempty"`
	Ssh                []string            `yaml:"ssh,omitempty" json:"ssh,omitempty"`
	CacheFrom          []string            `yaml:"cache_from,omitempty" json:"cache_from,omitempty"`
	CacheTo            []string            `yaml:"cache_to,omitempty" json:"cache_to,omitempty"`
	AdditionalContexts ListOrMap           `yaml:"additional_contexts,omitempty" json:"additional_contexts,omitempty"`
	ExtraHosts         []string            `yaml:"extra_hosts,omitempty" json:"extra_hosts,omitempty"`
	Isolation          interface{}         `yaml:"isolation,omitempty" json:"isolation,omitempty"`
	Privileged         bool                `yaml:"privileged,omitempty" json:"privileged,omitempty"`
	Labels             ListOrMap           `yaml:"labels,omitempty" json:"labels,omitempty"`
	NoCache            interface{}         `yaml:"no_cache,omitempty" json:"no_cache,omitempty"`
	Pull               interface{}         `yaml:"pull,omitempty" json:"pull,omitempty"`
	ShmSize            ByteValue           `yaml:"shm_size,omitempty" json:"shm_size,omitempty"`
	Target             string              `yaml:"target,omitempty" json:"target,omitempty"`
	Secrets            DockerComposeSecret `yaml:"secrets,omitempty" json:"secrets,omitempty"`
	Tags               []string            `yaml:"tags,omitempty" json:"tags,omitempty"`
	Platforms          []string            `yaml:"platforms,omitempty" json:"platforms,omitempty"`
}

type DockerComposeBlkioConfig struct {
	Weight       int `yaml:"weight,omitempty" json:"weight,omitempty"`
	WeightDevice []struct {
		Path   string `yaml:"path,omitempty" json:"path,omitempty"`
		Weight int    `yaml:"weight,omitempty" json:"weight,omitempty"`
	} `yaml:"weight_device,omitempty" json:"weight_device,omitempty"`
	DeviceReadBps []struct {
		Path string    `yaml:"path,omitempty" json:"path,omitempty"`
		Rate ByteValue `yaml:"rate,omitempty" json:"rate,omitempty"`
	} `yaml:"device_read_bps,omitempty" json:"device_read_bps,omitempty"`
	DeviceReadIops []struct {
		Path string `yaml:"path,omitempty" json:"path,omitempty"`
		Rate int    `yaml:"rate,omitempty" json:"rate,omitempty"`
	} `yaml:"device_read_iops,omitempty" json:"device_read_iops,omitempty"`
	DeviceWriteBps []struct {
		Path string    `yaml:"path,omitempty" json:"path,omitempty"`
		Rate ByteValue `yaml:"rate,omitempty" json:"rate,omitempty"`
	} `yaml:"device_write_bps,omitempty" json:"device_write_bps,omitempty"`
	DeviceWriteIops []struct {
		Path string `yaml:"path,omitempty" json:"path,omitempty"`
		Rate int    `yaml:"rate,omitempty" json:"rate,omitempty"`
	} `yaml:"device_write_iops,omitempty" json:"device_write_iops,omitempty"`
}

type DockerComposeCredentialSpec struct {
	File     string `yaml:"file,omitempty" json:"file,omitempty"`
	Config   string `yaml:"config,omitempty" json:"config,omitempty"`
	Registry string `yaml:"registry,omitempty" json:"registry,omitempty"`
}

type DockerComposeServiceDependsOnLongSyntax struct {
	Restart   bool   `yaml:"restart,omitempty" json:"restart,omitempty"`
	Condition string `yaml:"condition,omitempty" json:"condition,omitempty"`
	Required  bool   `default:"true" yaml:"required,omitempty" json:"required,omitempty"`
}
type DockerComposeServiceDependsOnShortSyntax []string

// [Docker Compose Service Depends On Definition](https://github.com/compose-spec/compose-spec/blob/master/05-services.md#depends_on)
//
// Need functions to select Short syntax or long syntax
// Short syntax will be a list of string, long syntax is an object
type DockerComposeServiceDependsOn interface{}

type DockerComposeDeployResources struct {
	Limits struct {
		Cpu     string    `yaml:"cpu,omitempty" json:"cpu,omitempty"`
		Memory  ByteValue `yaml:"memory,omitempty" json:"memory,omitempty"`
		Pids    int       `yaml:"pids,omitempty" json:"pids,omitempty"`
		Devices []struct {
			Capabilities []string               `yaml:"capabilities,omitempty" json:"capabilities,omitempty"`
			Driver       string                 `yaml:"driver,omitempty" json:"driver,omitempty"`
			Count        int                    `yaml:"count,omitempty" json:"count,omitempty"`
			DeviceIds    []string               `yaml:"device_ids,omitempty" json:"device_ids,omitempty"`
			Options      map[string]interface{} `yaml:"options,omitempty" json:"options,omitempty"`
		} `yaml:"devices,omitempty" json:"devices,omitempty"`
	} `yaml:"limits,omitempty" json:"limits,omitempty"`
}

type DockerComposeDeployRestartPolicy struct {
	Condition   string    `yaml:"condition,omitempty" json:"condition,omitempty"`
	Delay       int       `yaml:"delay,omitempty" json:"delay,omitempty"`
	MaxAttempts int       `yaml:"max_attempts,omitempty" json:"max_attempts,omitempty"`
	Window      TimeValue `yaml:"window,omitempty" json:"window,omitempty"`
}

type DockerComposeDeployRollbackConfig struct {
	Parallelism     int       `yaml:"parallelism,omitempty" json:"parallelism,omitempty"`
	Delay           TimeValue `yaml:"delay,omitempty" json:"delay,omitempty"`
	FailureAction   string    `default:"pause" yaml:"failure_action,omitempty" json:"failure_action,omitempty"`
	Monitor         TimeValue `yaml:"monitor,omitempty" json:"monitor,omitempty"`
	MaxFailureRatio float32   `yaml:"max_failure_ratio,omitempty" json:"max_failure_ratio,omitempty"`
	Order           string    `default:"stop-first" yaml:"order,omitempty" json:"order,omitempty"`
}

type DockerComposeDeployUpdateConfig struct {
	Parallelism     int       `yaml:"parallelism,omitempty" json:"parallelism,omitempty"`
	Delay           TimeValue `yaml:"delay,omitempty" json:"delay,omitempty"`
	FailureAction   string    `default:"pause" yaml:"failure_action,omitempty" json:"failure_action,omitempty"`
	Monitor         TimeValue `yaml:"monitor,omitempty" json:"monitor,omitempty"`
	MaxFailureRatio float32   `yaml:"max_failure_ratio,omitempty" json:"max_failure_ratio,omitempty"`
	Order           string    `default:"stop-first" yaml:"order,omitempty" json:"order,omitempty"`
}

type DockerComposeDeploy struct {
	EndpointMode   string                            `yaml:"endpoint_mode,omitempty" json:"endpoint_mode,omitempty"`
	Labels         map[string]string                 `yaml:"labels,omitempty" json:"labels,omitempty"`
	Mode           string                            `yaml:"mode,omitempty" json:"mode,omitempty"`
	Placement      string                            `yaml:"placement,omitempty" json:"placement,omitempty"`
	Reference      ListOrMap                         `yaml:"reference,omitempty" json:"reference,omitempty"`
	Replicas       int                               `yaml:"replicas,omitempty" json:"replicas,omitempty"`
	Resources      DockerComposeDeployResources      `yaml:"resources,omitempty" json:"resources,omitempty"`
	RestartPolicy  DockerComposeDeployRestartPolicy  `yaml:"restart_policy,omitempty" json:"restart_policy,omitempty"`
	RollbackConfig DockerComposeDeployRollbackConfig `yaml:"rollback_config,omitempty" json:"rollback_config,omitempty"`
	UpdateConfig   DockerComposeDeployUpdateConfig
}

type DockerComposeServiceExtends struct {
	File    string `yaml:"file,omitempty" json:"file,omitempty"`
	Service string `yaml:"service,omitempty" json:"service,omitempty"`
}

type DockerComposeHealthCheck struct {
	Test          StringOrList `yaml:"test,omitempty" json:"test,omitempty"`
	Interval      TimeValue    `yaml:"interval,omitempty" json:"interval,omitempty"`
	Timeout       TimeValue    `yaml:"timeout,omitempty" json:"timeout,omitempty"`
	Retries       int          `yaml:"retries,omitempty" json:"retries,omitempty"`
	StartPeriod   TimeValue    `yaml:"start_period,omitempty" json:"start_period,omitempty"`
	StartInterval TimeValue    `yaml:"start_interval,omitempty" json:"start_interval,omitempty"`
}

type DockerComposeServiceLogging struct {
	Driver  string                 `yaml:"driver,omitempty"`
	Options map[string]interface{} `yaml:"options,omitempty" json:"options,omitempty"`
}

type DockerComposeServiceUlimits struct {
	Nproc  int `yaml:"nproc,omitempty" json:"nproc,omitempty"`
	Nofile struct {
		Soft int `yaml:"soft,omitempty" json:"soft,omitempty"`
		Hard int `yaml:"hard,omitempty" json:"hard,omitempty"`
	} `yaml:"nofile,omitempty" json:"nofile,omitempty"`
}

// [Docker Compose Service Definition](https://github.com/compose-spec/compose-spec/blob/master/05-services.md)
type DockerComposeService struct {
	Attach bool `default:"true" yaml:"attach,omitempty" json:"attach,omitempty"`

	Build       StringOrMap              `yaml:"build,omitempty" json:"build,omitempty"` // Build can be a string or a structured data defined as DockerComposeServiceBuildObject
	BlkioConfig DockerComposeBlkioConfig `yaml:"blkio_config,omitempty" json:"blkio_config,omitempty"`

	CpuCount     int         `yaml:"cpu_count,omitempty" json:"cpu_count,omitempty"`
	CpuPercent   int         `yaml:"cpu_percent,omitempty" json:"cpu_percent,omitempty"`
	CpuShares    int         `yaml:"cpu_shares,omitempty" json:"cpu_shares,omitempty"`
	CpuPeriod    interface{} `yaml:"cpu_period,omitempty" json:"cpu_period,omitempty"`
	CpuQuota     interface{} `yaml:"cpu_quota,omitempty" json:"cpu_quota,omitempty"`
	CpuRtRuntime TimeValue   `yaml:"cpu_rt_runtime,omitempty" json:"cpu_rt_runtime,omitempty"`
	CpuRtPeriod  TimeValue   `yaml:"cpu_rt_period,omitempty" json:"cpu_rt_period,omitempty"`
	CpuSet       interface{} `yaml:"cpu_set,omitempty" json:"cpu_set,omitempty"`

	CapAdd  []string `yaml:"cap_add,omitempty" json:"cap_add,omitempty"`
	CapDrop []string `yaml:"cap_drop,omitempty" json:"cap_drop,omitempty"`

	Cgroup       string `yaml:"cgroup,omitempty" json:"cgroup,omitempty"`
	CgroupParent string `yaml:"cgroup_parent,omitempty" json:"cgroup_parent,omitempty"`

	Command ListOrString `yaml:"command,omitempty" json:"command,omitempty"`

	Configs []string `yaml:"configs,omitempty" json:"configs,omitempty"`

	ContainerName string `yaml:"container_name,omitempty" json:"container_name,omitempty"`

	CredentialSpec DockerComposeCredentialSpec `yaml:"credential_spec,omitempty" json:"credential_spec,omitempty"`

	DependsOn DockerComposeServiceDependsOn `yaml:"depends_on,omitempty" json:"depends_on,omitempty"`

	Deploy DockerComposeDeploy `yaml:"deploy,omitempty" json:"deploy,omitempty"`

	DeviceCgroupRules []string `yaml:"device_cgroup_rules,omitempty" json:"device_cgroup_rules,omitempty"`
	Devices           []string `yaml:"devices,omitempty" json:"device,omitempty"`

	Dns        StringOrList `yaml:"dns,omitempty" json:"dns,omitempty"`
	DnsOpt     []string     `yaml:"dns_opt,omitempty" json:"dns_opt,omitempty"`
	DnsSear    StringOrList `yaml:"dns_search,omitempty" json:"dns_search,omitempty"`
	DomainName string       `yaml:"domainname,omitempty" json:"domainname,omitempty"`

	Entrypoint StringOrList `yaml:"entrypoint,omitempty" json:"entrypoint,omitempty"`

	EnvFile     StringOrList `yaml:"env_file,omitempty" json:"env_file,omitempty"`
	Environment ListOrMap    `yaml:"environment,omitempty" json:"environment,omitempty"`

	Expose []string `yaml:"expose,omitempty" json:"expose,omitempty"`

	Extends DockerComposeServiceExtends `yaml:"extends,omitempty" json:"extends,omitempty"`

	Annotations ListOrMap `yaml:"annotations,omitempty" json:"annotations,omitempty"`

	ExternalLinks []string `yaml:"external_links,omitempty" json:"external_links,omitempty"`

	ExtraHosts ListOrMap `yaml:"extra_hosts,omitempty" json:"extra_hosts,omitempty"`

	GroupAdd []string `yaml:"group_add,omitempty" json:"group_add,omitempty"`

	Healthcheck DockerComposeHealthCheck `yaml:"healthcheck,omitempty" json:"healthcheck,omitempty"`

	Hostname string `yaml:"hostname,omitempty" json:"hostname,omitempty"`

	Image string `yaml:"image,omitempty" json:"image,omitempty"`

	Init bool `yaml:"init,omitempty" json:"init,omitempty"`

	Ipc string `yaml:"ipc,omitempty" json:"ipc,omitempty"`

	Uts string `yaml:"uts,omitempty" json:"uts,omitempty"`

	Isolation interface{} `yaml:"isolation,omitempty" json:"isolation,omitempty"`

	Labels ListOrMap `yaml:"labels,omitempty" json:"labels,omitempty"`

	Links []string `yaml:"links,omitempty" json:"links,omitempty"`

	Logging DockerComposeServiceLogging `yaml:"logging,omitempty" json:"logging,omitempty"`

	NetworkMode string    `yaml:"network_mode,omitempty" json:"network_mode,omitempty"`
	Networks    ListOrMap `yaml:"networks,omitempty" json:"networks,omitempty"`

	MacAddress string `yaml:"mac_address,omitempty" json:"mac_address,omitempty"`

	MemSwappiness int    `yaml:"mem_swappiness,omitempty" json:"mem_swappiness,omitempty"`
	MemswapLimit  string `yaml:"memswap_limit,omitempty" json:"memswap_limit,omitempty"`

	OomKillDisable bool `yaml:"oom_kill_disable,omitempty" json:"oom_kill_disable,omitempty"`
	OomScoreAdj    int  `yaml:"oom_score_adj,omitempty" json:"oom_score_adj,omitempty"`

	Pid int `yaml:"pid,omitempty" json:"pid,omitempty"`

	Platform string `yaml:"platform,omitempty" json:"platform,omitempty"`

	Ports ListStringOrListObject `yaml:"ports,omitempty" json:"ports,omitempty"`

	Privileged string `yaml:"privileged,omitempty" json:"privileged,omitempty"`

	Profiles []string `yaml:"profiles,omitempty" json:"profiles,omitempty"`

	PullPolicy string `yaml:"pull_policy,omitempty" json:"pull_policy,omitempty"`

	ReadOnly interface{} `yaml:"read_only,omitempty" json:"read_only,omitempty"`

	Restart string `yaml:"restart,omitempty" json:"restart,omitempty"`

	Runtime string `yaml:"runtime,omitempty" json:"runtime,omitempty"`

	Secrets ListStringOrListObject `yaml:"secrets,omitempty" json:"secrets,omitempty"`

	SecurityOpt []string `yaml:"security_opt,omitempty" json:"security_opt,omitempty"`

	ShmSize ByteValue `yaml:"shm_size,omitempty" json:"shm_size,omitempty"`

	StdinOpen interface{} `yaml:"stdin_open,omitempty" json:"stdin_open,omitempty"`

	StopRatePeriod TimeValue `yaml:"stop_rate_period,omitempty" json:"stop_rate_period,omitempty"`
	StopSignal     string    `yaml:"stop_signal,omitempty" json:"stop_signal,omitempty"`

	StorageOpt map[string]interface{} `yaml:"storage_opt,omitempty" json:"storage_opt,omitempty"`

	Sysctls ListOrMap `yaml:"sysctls,omitempty" json:"sysctls,omitempty"`

	Tmpfs StringOrList `yaml:"tmpfs,omitempty" json:"tmpfs,omitempty"`

	Tty interface{} `yaml:"tty,omitempty" json:"tty,omitempty"`

	Ulimits DockerComposeServiceUlimits `yaml:"ulimits,omitempty" json:"ulimits,omitempty"`

	User       string `yaml:"user,omitempty" json:"user,omitempty"`
	UsernsMode string `yaml:"userns_mode,omitempty" json:"userns_mode,omitempty"`

	Volumes     ListStringOrListObject `yaml:"volumes,omitempty" json:"volumes,omitempty"`
	VolumesFrom []string               `yaml:"volumes_from,omitempty" json:"volumes_from,omitempty"`

	WorkingDir string `yaml:"working_dir,omitempty" json:"working_dir,omitempty"`
}
type DockerComposeServices map[string]DockerComposeService
type DockerComposeVolumes struct{}
type DockerComposeConfigs struct{}
type DockerComposeSecrets struct{}
type DockerComposeNetworks struct{}

type DockerComposeStruct struct {
	Services DockerComposeServices `yaml:"services,omitempty" json:"services,omitempty"`
	Volumes  DockerComposeVolumes  `yaml:"volumes,omitempty" json:"volumes,omitempty"`
	Configs  DockerComposeConfigs  `yaml:"configs,omitempty" json:"configs,omitempty"`
	Secrets  DockerComposeSecrets  `yaml:"secret,omitempty" json:"secrets,omitempty"`
	Networks DockerComposeNetworks `yaml:"networks,omitempty" json:"networks,omitempty"`
}
