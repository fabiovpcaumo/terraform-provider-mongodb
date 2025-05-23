# MongoDB Provider

The MongoDB provider is used to interact with the resources supported by [MongoDB](https://www.mongodb.com/). The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available provider resources.

You may want to consider pinning the [provider version](https://www.terraform.io/docs/configuration/providers.html#provider-versions) to ensure you have a chance to review and prepare for changes.

## Example Usage

```hcl
# Configure the MongoDB Provider
provider "mongodb" {
  host = "127.0.0.1"
  port = "27017"
  username = "root"
  password = "root"
  auth_database = "admin"
  ssl = true
  replica_set = "replica-set" #optional
  retrywrites = false # default true
  direct = true // default false
  proxy = "socks5://myproxy:8080" // Optional
}
```

## Example Usage with ssl

```hcl
# Configure the MongoDB Provider
provider "mongodb" {
  insecure_skip_verify = true  # default false (set to true to ignore hostname verification)
  # -> specify certificate path
  certificate = file(pathexpand("path/to/certificate/ca.pem"))
}
```

### Environment variables

You can also provide your credentials via the environment variables, MONGO_HOST, MONGO_PORT, MONGO_USR, and MONGO_PWD respectively:

```hcl
provider "mongodb" {
  auth_database = "admin"
}
```

Usage (prefix the export commands with a space to avoid the keys being recorded in OS history):

```shell
$  export MONGO_HOST="xxxx"
$  export MONGO_PORT="xxxx"
$  export MONGO_USR="xxxx"
$  export MONGO_PWD="xxxx"
$ terraform plan
```

## Certificate information :

Specify certificate information either with a directory or directly with the content of the files for connecting to the Mongodb host via TLS.

```hcl
provider "mongodb" {
  host = "127.0.0.1"
  port = "27017"
  username = "root"
  password = "root"
  auth_database = "admin"
  ssl = true
  # -> specify either
  certificate = pathexpand("~/.mongodb/ca.pem")
  }
```

## Argument Reference

In addition to [generic `provider`
arguments](https://www.terraform.io/docs/configuration/providers.html) (e.g.
`alias` and `version`), the following arguments are supported in the MongoDB
`provider` block:

* `host` - (Optional) This is the host your MongoDB Server. It must be
  provided, but it can also be sourced from the `MONGO_HOST`
  environment variable.
* `port` - (Optional) This is the port that your MongoDB Server uses. It must be
  provided, but it can also be sourced from the `MONGO_PORT`
  environment variable.

* `certificate` - (Optional) Path to a directory with certificate files for connecting to the Docker host via TLS. I. If the path is blank, the MONGODB_CERT will also be checked.

* `username ` - (Optional) Specifies a username with which to authenticate to the MongoDB database. It must be
  provided, but it can also be sourced from the `MONGO_USR`
  environment variable.
* `password  ` - (Optional) Specifies a password with which to authenticate to the MongoDB database. It must be
  provided, but it can also be sourced from the `MONGO_PWD`
  environment variable.
* `auth_database   ` - (Required) Specifies the authentication database where the specified `username` has been created.
* `ssl   ` - (Optional) `default = false `set it to true to connect to a deployment using TLS/SSL with SCRAM authentication.
* `retrywrites   ` - (Optional) `default = true `Retryable writes allow MongoDB drivers to automatically retry certain write operations a single time if they encounter network errors, or if they cannot find a healthy primary in the replica sets or sharded cluster.
* `direct   ` - (Optional) `default = false ` determine if a direct connection is needed..
* `proxy   ` - (Optional) `default = "" ` determine if connecting via a SOCKS5 proxy is needed, it can also be sourced from the `ALL_PROXY` or `all_proxy` environment variable.
* `timeout` - (Optional) `default = 10000 ` Specifies the number of milliseconds that a single operation run on the Client can take before returning a timeout error.
* `connect_timeout` - (Optional) `default = 30000 ` Specifies the time in milliseconds to attempt a connection before timing out.
* `server_selection_timeout` - (Optional) Specifies the time in milliseconds to wait to find an available, suitable server to execute an operation.
* `replica_set` - (Optional) Specifies the name of the replica set to connect to.
* `replica_set_hosts` - (Optional) Comma separated list of hosts for the replica set.
* `read_preference` - (Optional) Specifies the default read preference for the client (excluding tags). See [read preference](https://www.mongodb.com/docs/manual/core/read-preference/#std-label-read-preference) for more information.
* `max_pool_size` - (Optional) Specifies the maximum number of clients or connections the driver can create in its connection pool.
* `max_connecting` - (Optional) Specifies the maximum number of connections a driver's connection pool may be establishing concurrently.
  Optional) Specifies the maximum number of connections a driver's connection pool may be establishing concurrently.
