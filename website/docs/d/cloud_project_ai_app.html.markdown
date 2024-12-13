---
subcategory : "Managed Databases"
---

# ovh_cloud_project_ai_app (Resource)

Use this resource to manage an [AI deploy](https://www.ovhcloud.com/fr/public-cloud/ai-deploy/) app

## Example Usage

```hcl
resource "ovh_cloud_project_ai_app" "my_test_app" {
  service_name = "2982d81c622c48228ebb42e48efb4337"
  region = "GRA"
  resources = {
    cpu = 1
    flavor = "ai1-1-cpu"
  }
  image = "ubuntu"
  default_http_port=8080
  grpc_port=8081
  env_vars = [
    {
      name = "ENV_VAR_1"
      value = "env var 1 value"
    },
  ]
}
```

## Argument Reference

* `service_name` - (Required) The id of the public cloud project. If omitted,
  the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.

* `region` - (Required) Host region of the app

* `image` - (Required) Docker image to use for the AI deploy app

* `command` - App command

* `name` - App name

* `partner_id` - Partner ID

* `default_http_port` - Default port to access http service inside the app

* `unsecure_http` - Whether if app api port can be accessed without any authentication token

* `grpc_port` - GRPC Port that we want to expose in case workload HTTP & gRPC servers cannot be multiplexed to listen on the same port

* `labels` - Key/Value map of labels that are used to scope tokens, labels prefixed by 'ovh/' are owned by the platform and overridden

* `probe` - App readiness probe
  * `failure_threshold` - Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
  * `initial_delay_seconds` - Number of seconds after the container has started before liveness probes are initiated.
  * `path` - Path to access to check for readiness
  * `period_seconds` - How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
  * `port` - Port to access to check for readiness
  * `success_threshold` - Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.
  * `timeout_seconds` - Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1.

* `liveness_probe` - App liveness probe
  * `failure_threshold` - Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
  * `initial_delay_seconds` - Number of seconds after the container has started before liveness probes are initiated.
  * `path` - Path to access to check for readiness
  * `period_seconds` - How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
  * `port` - Port to access to check for readiness
  * `success_threshold` - Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.
  * `timeout_seconds` - Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1.

* `deployment_strategy` - List of environment variable to be set inside the app
  * `max_surge` - Maximum number of replicas that can be created over the desired number of Pods (can be expressed as a percentage of the desired pods, suffixed with '%')
  * `max_unavailable` - Maximum number of replicas that can be unavailable during the update process (can be expressed as a percentage of the desired pods, suffixed with '%')
  * `progress_deadline_seconds` - Number of seconds you want to wait for your Deployment to progress before the system reports back that the Deployment has failed progressing

* `env_vars` - List of environment variable to be set inside the app
  * `name` - Name of the environment variable to set inside the job
  * `value` - Value of the environment variable to set inside the job

* `resources` -  App resources
  * `cpu` -  Number of vCPU resources requested
  * `gpu` -  Number of GPU resources requested
  * `ephemeral_storage` - The amount of ephemeral storage in bytes
  * `flavor` - Instance flavor
  * `gpu_brand ` - The GPU Brand
  * `gpu_memory` - The GPU Memory in bytes
  * `gpu_model` - The GPU Model
  * `memory` -  The amount of memory in bytes
  * `private_network` - The private network bandwidth in bits per seconds
  * `public_network` - The public network bandwidth in bits per seconds

* `scaling_strategy` - App scaling strategy
  * `automatic` - Strategy setting a variable number of replicas, based on an average resource usage threshold (conflicts with 'fixed' property when both are not null)
    * `average_usage_target` - The average resource usage threshold that the app upscale or downscale will be triggered from, in percent
    * `replicas_max` - Maximum number of replicas
    * `replicas_min` - Minimum number of replicas
    * `resource_type` - Type of the resource to base the automatic scaling on. Allowed: CPU┃RAM
  * `fixed` - Strategy setting a fix number of replicas (conflicts with 'automatic' property when both are not null)
    * `replicas` - Number of wanted replicas

* `volumes` - App Data linked
  * `cache` - Enable/disable volume caching
  * `container` - Public Cloud Storage container to attach
  * `mount_path` - (Required) Path where to mount the data inside the container
  * `permission` - (Required) Permissions to use on the mounted volume.  Allowed: RO┃RW┃RWD
  * `prefix` - Prefix to fetch only part of the volume
  * `region` - Public Cloud Storage Region
  * `volume_source` - Source volume details
    * `data_store` - Volume details for data store containers.
      * `alias` - (Required) Data store alias
      * `container` - (Required) Data store container to attach
      * `archive` - Name of the tar archive that needs to be saved
      * `internal` - True if data is stored on OVHcloud AI's internal storage
      * `prefix` - Prefix to fetch only part of the volume
    * `public_git` - Volume details for public git repositories.
      * `url` - (Required) URL of the public git repository
    * `public_swift` - Volume details for public swift containers.
      * `url` - (Required) URL of the public swift container
    * `standalone` - Volume details for volumes that do not have a datasource.
      * `name` - Name of the volume
  * `volume_target` - Target volume details
    * `target_data_store` - Volume details for data store containers
      * `alias` - (Required) Data store alias
      * `container` - (Required) Data store container to attach
      * `archive` - Name of the tar archive that needs to be saved
      * `internal` - True if data is stored on OVHcloud AI's internal storage
      * `prefix` - Prefix to fetch only part of the volume
  
## Attributes Reference

`id` App Id
`created_at` App creation date

`spec` - App spec
  * `service_name` - (Required) The id of the public cloud project. If omitted,
  the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.

  * `region` - (Required) Host region of the app

  * `image` - (Required) Docker image to use for the AI deploy app

  * `command` - App command

  * `name` - App name

  * `partner_id` - Partner ID

  * `default_http_port` - Default port to access http service inside the app

  * `unsecure_http` - Whether if app api port can be accessed without any authentication token

  * `grpc_port` - GRPC Port that we want to expose in case workload HTTP & gRPC servers cannot be multiplexed to listen on the same port

  * `labels` - Key/Value map of labels that are used to scope tokens, labels prefixed by 'ovh/' are owned by the platform and overridden

  * `probe` - App readiness probe
    * `failure_threshold` - Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
    * `initial_delay_seconds` - Number of seconds after the container has started before liveness probes are initiated.
    * `path` - Path to access to check for readiness
    * `period_seconds` - How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
    * `port` - Port to access to check for readiness
    * `success_threshold` - Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.
    * `timeout_seconds` - Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1.

  * `liveness_probe` - App liveness probe
    * `failure_threshold` - Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
    * `initial_delay_seconds` - Number of seconds after the container has started before liveness probes are initiated.
    * `path` - Path to access to check for readiness
    * `period_seconds` - How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
    * `port` - Port to access to check for readiness
    * `success_threshold` - Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.
    * `timeout_seconds` - Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1.

  * `deployment_strategy` - List of environment variable to be set inside the app
    * `max_surge` - Maximum number of replicas that can be created over the desired number of Pods (can be expressed as a percentage of the desired pods, suffixed with '%')
    * `max_unavailable` - Maximum number of replicas that can be unavailable during the update process (can be expressed as a percentage of the desired pods, suffixed with '%')
    * `progress_deadline_seconds` - Number of seconds you want to wait for your Deployment to progress before the system reports back that the Deployment has failed progressing

  * `env_vars` - List of environment variable to be set inside the app
    * `name` - Name of the environment variable to set inside the job
    * `value` - Value of the environment variable to set inside the job

  * `resources` -  App resources
    * `cpu` -  Number of vCPU resources requested
    * `gpu` -  Number of GPU resources requested
    * `ephemeral_storage` - The amount of ephemeral storage in bytes
    * `flavor` - Instance flavor
    * `gpu_brand ` - The GPU Brand
    * `gpu_memory` - The GPU Memory in bytes
    * `gpu_model` - The GPU Model
    * `memory` -  The amount of memory in bytes
    * `private_network` - The private network bandwidth in bits per seconds
    * `public_network` - The public network bandwidth in bits per seconds

  * `scaling_strategy` - App scaling strategy
    * `automatic` - Strategy setting a variable number of replicas, based on an average resource usage threshold (conflicts with 'fixed' property when both are not null)
      * `average_usage_target` - The average resource usage threshold that the app upscale or downscale will be triggered from, in percent
      * `replicas_max` - Maximum number of replicas
      * `replicas_min` - Minimum number of replicas
      * `resource_type` - Type of the resource to base the automatic scaling on. Allowed: CPU┃RAM
    * `fixed` - Strategy setting a fix number of replicas (conflicts with 'automatic' property when both are not null)
      * `replicas` - Number of wanted replicas

  * `volumes` - App Data linked
    * `cache` - Enable/disable volume caching
    * `container` - Public Cloud Storage container to attach
    * `mount_path` - (Required) Path where to mount the data inside the container
    * `permission` - (Required) Permissions to use on the mounted volume.  Allowed: RO┃RW┃RWD
    * `prefix` - Prefix to fetch only part of the volume
    * `region` - Public Cloud Storage Region
    * `volume_source` - Source volume details
      * `data_store` - Volume details for data store containers.
        * `alias` - (Required) Data store alias
        * `container` - (Required) Data store container to attach
        * `archive` - Name of the tar archive that needs to be saved
        * `internal` - True if data is stored on OVHcloud AI's internal storage
        * `prefix` - Prefix to fetch only part of the volume
      * `public_git` - Volume details for public git repositories.
        * `url` - (Required) URL of the public git repository
      * `public_swift` - Volume details for public swift containers.
        * `url` - (Required) URL of the public swift container
      * `standalone` - Volume details for volumes that do not have a datasource.
        * `name` - Name of the volume
    * `volume_target` - Target volume details
      * `target_data_store` - Volume details for data store containers
        * `alias` - (Required) Data store alias
        * `container` - (Required) Data store container to attach
        * `archive` - Name of the tar archive that needs to be saved
        * `internal` - True if data is stored on OVHcloud AI's internal storage
        * `prefix` - Prefix to fetch only part of the volume