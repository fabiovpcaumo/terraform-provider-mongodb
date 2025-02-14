package mongodb

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MONGO_HOST", "127.0.0.1"),
				Description: "The mongodb server address",
			},
			"port": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MONGO_PORT", "27017"),
				Description: "The mongodb server port",
			},
			"certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MONGODB_CERT", ""),
				Description: "PEM-encoded content of Mongodb host CA certificate",
			},

			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MONGO_USR", nil),
				Description: "The mongodb user",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MONGO_PWD", nil),
				Description: "The mongodb password",
			},
			"auth_database": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "admin",
				Description: "The mongodb auth database",
			},
			"replica_set": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The mongodb replica set",
			},
			"insecure_skip_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "ignore hostname verification",
			},
			"ssl": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "ssl activation",
			},
			"direct": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "enforces a direct connection instead of discovery",
			},
			"retrywrites": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Retryable Writes",
			},
			"proxy": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"ALL_PROXY",
					"all_proxy",
				}, nil),
				ValidateDiagFunc: validateDiagFunc(validation.StringMatch(regexp.MustCompile("^socks5h?://.*:\\d+$"), "The proxy URL is not a valid socks url.")),
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     10000,
				Description: "Specifies the number of milliseconds that a single operation run on the Client can take before returning a timeout error. Operations honor this setting only if there is no deadline on the operation Context.",
			},
			"connect_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     30000,
				Description: "Specifies the time in milliseconds to attempt a connection before timing out.",
			},
			"server_selection_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the time in milliseconds to wait to find an available, suitable server to execute an operation.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"mongodb_db_user": resourceDatabaseUser(),
			"mongodb_db_role": resourceDatabaseRole(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	clientConfig := ClientConfig{
		Host:                   d.Get("host").(string),
		Port:                   d.Get("port").(string),
		Username:               d.Get("username").(string),
		Password:               d.Get("password").(string),
		DB:                     d.Get("auth_database").(string),
		Ssl:                    d.Get("ssl").(bool),
		ReplicaSet:             d.Get("replica_set").(string),
		Certificate:            d.Get("certificate").(string),
		InsecureSkipVerify:     d.Get("insecure_skip_verify").(bool),
		Direct:                 d.Get("direct").(bool),
		RetryWrites:            d.Get("retrywrites").(bool),
		Proxy:                  d.Get("proxy").(string),
		Timeout:                d.Get("timeout").(int),
		ConnectTimeout:         d.Get("connect_timeout").(int),
		ServerSelectionTimeout: d.Get("server_selection_timeout").(int),
	}

	return &clientConfig, diags

}
