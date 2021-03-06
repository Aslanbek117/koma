package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "github.com/mlabouardy/komiser/handlers/aws"
	. "github.com/mlabouardy/komiser/handlers/digitalocean"
	. "github.com/mlabouardy/komiser/handlers/gcp"
	. "github.com/mlabouardy/komiser/handlers/ovh"
	. "github.com/mlabouardy/komiser/services/cache"
	_ "github.com/mlabouardy/komiser/services/ini"
	"strconv"
	"time"
)

const (
	DEFAULT_PORT     = 3005
	DEFAULT_DURATION = 30
)


func startServer(port int, cache Cache, dataset string, multiple bool) {
	cache.Connect()

	digitaloceanHandler := NewDigitalOceanHandler(cache)
	gcpHandler := NewGCPHandler(cache, dataset)
	awsHandler := NewAWSHandler(cache, multiple)
	ovhHandler := NewOVHHandler(cache, "")

	r := mux.NewRouter()
	r.HandleFunc("/aws/profiles", awsHandler.ConfigProfilesHandler)
	r.HandleFunc("/aws/iam/users", awsHandler.IAMUsersHandler)
	r.HandleFunc("/aws/iam/account", awsHandler.IAMUserHandler)
	r.HandleFunc("/aws/cost/current", awsHandler.CurrentCostHandler)
	r.HandleFunc("/aws/cost/history", awsHandler.CostAndUsageHandler)
	r.HandleFunc("/aws/resources/regions", awsHandler.UsedRegionsHandler)
	r.HandleFunc("/aws/cloudwatch/alarms", awsHandler.CloudWatchAlarmsHandler)
	r.HandleFunc("/aws/ec2/regions", awsHandler.EC2InstancesHandler)
	r.HandleFunc("/aws/lambda/functions", awsHandler.LambdaFunctionHandler)
	r.HandleFunc("/aws/lambda/invocations", awsHandler.GetLambdaInvocationMetrics)
	r.HandleFunc("/aws/s3/buckets", awsHandler.S3BucketsHandler)
	r.HandleFunc("/aws/s3/size", awsHandler.S3BucketsSizeHandler)
	r.HandleFunc("/aws/s3/objects", awsHandler.S3BucketsObjectsHandler)
	r.HandleFunc("/aws/ebs", awsHandler.EBSHandler)
	r.HandleFunc("/aws/rds/instances", awsHandler.RDSInstanceHandler)
	r.HandleFunc("/aws/dynamodb/tables", awsHandler.DynamoDBTableHandler)
	r.HandleFunc("/aws/elasticache/clusters", awsHandler.ElasticacheClustersHandler)
	r.HandleFunc("/aws/vpc", awsHandler.VPCHandler)
	r.HandleFunc("/aws/acl", awsHandler.ACLHandler)
	r.HandleFunc("/aws/route_tables", awsHandler.RouteTableHandler)
	r.HandleFunc("/aws/cloudfront/requests", awsHandler.CloudFrontRequestsHandler)
	r.HandleFunc("/aws/cloudfront/distributions", awsHandler.CloudFrontDistributionsHandler)
	r.HandleFunc("/aws/apigateway/requests", awsHandler.APIGatewayRequestsHandler)
	r.HandleFunc("/aws/apigateway/apis", awsHandler.APIGatewayRestAPIsHandler)
	r.HandleFunc("/aws/elb/requests", awsHandler.ELBRequestsHandler)
	r.HandleFunc("/aws/elb/family", awsHandler.ElasticLoadBalancerHandler)
	r.HandleFunc("/aws/kms", awsHandler.KMSKeysHandler)
	r.HandleFunc("/aws/key_pairs", awsHandler.KeyPairHandler)
	r.HandleFunc("/aws/security_groups", awsHandler.SecurityGroupHandler)
	r.HandleFunc("/aws/security_groups/unrestricted", awsHandler.ListUnrestrictedSecurityGroups)
	r.HandleFunc("/aws/acm/certificates", awsHandler.APIGatewayListCertificatesHandler)
	r.HandleFunc("/aws/acm/expired", awsHandler.APIGatewayExpiredCertificatesHandler)
	r.HandleFunc("/aws/sqs/messages", awsHandler.GetNumberOfMessagesSentAndDeletedSQSHandler)
	r.HandleFunc("/aws/sqs/queues", awsHandler.SQSQueuesHandler)
	r.HandleFunc("/aws/sns/topics", awsHandler.SNSTopicsHandler)
	r.HandleFunc("/aws/mq/brokers", awsHandler.ActiveMQBrokersHandler)
	r.HandleFunc("/aws/kinesis/streams", awsHandler.KinesisListStreamsHandler)
	r.HandleFunc("/aws/kinesis/shards", awsHandler.KinesisListShardsHandler)
	r.HandleFunc("/aws/glue/crawlers", awsHandler.GlueGetCrawlersHandler)
	r.HandleFunc("/aws/glue/jobs", awsHandler.GlueGetJobsHandler)
	r.HandleFunc("/aws/datapipeline/pipelines", awsHandler.DataPipelineListPipelines)
	r.HandleFunc("/aws/es/domains", awsHandler.ESListDomainsHandler)
	r.HandleFunc("/aws/swf/domains", awsHandler.SWFListDomainsHandler)
	r.HandleFunc("/aws/support/open", awsHandler.SupportOpenTicketsHandler)
	r.HandleFunc("/aws/support/history", awsHandler.SupportTicketsInLastSixMonthsHandlers)
	r.HandleFunc("/aws/ecs", awsHandler.ECSHandler)
	r.HandleFunc("/aws/route53/zones", awsHandler.Route53HostedZonesHandler)
	r.HandleFunc("/aws/route53/records", awsHandler.Route53ARecordsHandler)
	r.HandleFunc("/aws/logs/volume", awsHandler.LogsVolumeHandler)
	r.HandleFunc("/aws/cloudtrail/sign_in_event", awsHandler.CloudTrailConsoleSignInEventsHandler)
	r.HandleFunc("/aws/cloudtrail/source_ip", awsHandler.CloudTrailConsoleSignInSourceIpEventsHandler)
	r.HandleFunc("/aws/lambda/errors", awsHandler.GetLambdaErrorsMetrics)
	r.HandleFunc("/aws/ec2/scheduled", awsHandler.ScheduledEC2Instances)
	r.HandleFunc("/aws/ec2/reserved", awsHandler.ReservedEC2Instances)
	r.HandleFunc("/aws/ec2/spot", awsHandler.SpotEC2Instances)
	r.HandleFunc("/aws/cost/instance_type", awsHandler.CostAndUsagePerInstanceTypeHandler)
	r.HandleFunc("/aws/eks/clusters", awsHandler.EKSClustersHandler)
	r.HandleFunc("/aws/logs/retention", awsHandler.MaximumLogsRetentionPeriodHandler)
	r.HandleFunc("/aws/nat/traffic", awsHandler.GetNatGatewayTrafficHandler)
	r.HandleFunc("/aws/iam/organization", awsHandler.DescribeOrganizationHandler)
	r.HandleFunc("/aws/service/limits", awsHandler.DescribeServiceLimitsChecks)
	r.HandleFunc("/aws/s3/empty", awsHandler.GetEmptyBucketsHandler)
	r.HandleFunc("/aws/eip/detached", awsHandler.ElasticIPHandler)
	r.HandleFunc("/aws/redshift/clusters", awsHandler.DescribeRedshiftClustersHandler)
	r.HandleFunc("/aws/vpc/subnets", awsHandler.DescribeSubnetsHandler)
	r.HandleFunc("/aws/cost/forecast", awsHandler.DescribeForecastPriceHandler)

	r.HandleFunc("/gcp/resourcemanager/projects", gcpHandler.ProjectsHandler)
	r.HandleFunc("/gcp/compute/instances", gcpHandler.ComputeInstancesHandler)
	r.HandleFunc("/gcp/iam/roles", gcpHandler.IAMRolesHandler)
	r.HandleFunc("/gcp/dns/zones", gcpHandler.DNSManagedZonesHandler)
	r.HandleFunc("/gcp/storage/buckets", gcpHandler.StorageBucketsHandler)
	r.HandleFunc("/gcp/cloudfunctions/functions", gcpHandler.CloudFunctionsHandler)
	r.HandleFunc("/gcp/compute/disks", gcpHandler.ComputeDisksHandler)
	r.HandleFunc("/gcp/pubsub/topics", gcpHandler.PubSubTopicsHandler)
	r.HandleFunc("/gcp/sql/instances", gcpHandler.SqlInstancesHandler)
	r.HandleFunc("/gcp/vpc/networks", gcpHandler.VpcNetworksHandler)
	r.HandleFunc("/gcp/vpc/firewalls", gcpHandler.VpcFirewallsHandler)
	r.HandleFunc("/gcp/vpc/routers", gcpHandler.VpcRoutersHandler)
	r.HandleFunc("/gcp/compute/snapshots", gcpHandler.DiskSnapshotsHandler)
	r.HandleFunc("/gcp/storage/size", gcpHandler.StorageBucketsSizeHandler)
	r.HandleFunc("/gcp/storage/objects", gcpHandler.StorageBucketsObjectsHandler)
	r.HandleFunc("/gcp/logging/bytes_ingested", gcpHandler.LoggingBillableReceivedBytesHandler)
	r.HandleFunc("/gcp/kubernetes/clusters", gcpHandler.KubernetesClustersHandler)
	r.HandleFunc("/gcp/compute/images", gcpHandler.ComputeImagesHandler)
	r.HandleFunc("/gcp/redis/instances", gcpHandler.RedisInstancesHandler)
	r.HandleFunc("/gcp/compute/cpu", gcpHandler.ComputeCPUUtilizationHandler)
	r.HandleFunc("/gcp/iam/users", gcpHandler.IAMUsersHandler)
	r.HandleFunc("/gcp/bigquery/statements", gcpHandler.BigQueryScannedStatementsHandler)
	r.HandleFunc("/gcp/bigquery/storage", gcpHandler.BigQueryStoredBytesHandler)
	r.HandleFunc("/gcp/bigquery/datasets", gcpHandler.BigQueryDatasetsHandler)
	r.HandleFunc("/gcp/bigquery/tables", gcpHandler.BigQueryTablesHandler)
	r.HandleFunc("/gcp/compute/quotas", gcpHandler.ComputeQuotasHandler)
	r.HandleFunc("/gcp/lb/requests", gcpHandler.LoadBalancersRequestsHandler)
	r.HandleFunc("/gcp/api/requests", gcpHandler.ConsumedAPIRequestsHandler)
	r.HandleFunc("/gcp/lb/total", gcpHandler.LoadBalancersTotalHandler)
	r.HandleFunc("/gcp/vpc/subnets", gcpHandler.VpcSubnetsHandler)
	r.HandleFunc("/gcp/vpc/addresses", gcpHandler.VpcExternalAddressesHandler)
	r.HandleFunc("/gcp/vpn/tunnels", gcpHandler.VpnTunnelsHandler)
	r.HandleFunc("/gcp/ssl/certificates", gcpHandler.SSLCertificatesHandler)
	r.HandleFunc("/gcp/ssl/policies", gcpHandler.SSLPoliciesHandler)
	r.HandleFunc("/gcp/security/policies", gcpHandler.SecurityPoliciesHandler)
	r.HandleFunc("/gcp/kms/cryptokeys", gcpHandler.KMSCryptoKeysHandler)
	r.HandleFunc("/gcp/gae/bandwidth", gcpHandler.AppEngineOutgoingBandwidthHandler)
	r.HandleFunc("/gcp/serviceusage/apis", gcpHandler.EnabledAPIsHandler)
	r.HandleFunc("/gcp/dataproc/jobs", gcpHandler.DataprocJobsHandler)
	r.HandleFunc("/gcp/dataproc/clusters", gcpHandler.DataprocClustersHandler)
	r.HandleFunc("/gcp/billing/history", gcpHandler.BillingLastSixMonthsHandler)
	r.HandleFunc("/gcp/billing/service", gcpHandler.BillingPerServiceHandler)
	r.HandleFunc("/gcp/dns/records", gcpHandler.DNSARecordsHandler)
	r.HandleFunc("/gcp/iam/service_accounts", gcpHandler.IAMServiceAccountsHandler)
	r.HandleFunc("/gcp/dataflow/jobs", gcpHandler.DataflowJobsHandler)
	r.HandleFunc("/gcp/nat/gateways", gcpHandler.NatGatewaysHandler)

	r.HandleFunc("/ovh/cloud/projects", ovhHandler.DescribeCloudProjectsHandler)
	r.HandleFunc("/ovh/cloud/instances", ovhHandler.DescribeCloudInstancesHandler)
	r.HandleFunc("/ovh/cloud/storage", ovhHandler.DescribeCloudStorageContainersHandler)
	r.HandleFunc("/ovh/cloud/users", ovhHandler.DescribeCloudUsersHandler)
	r.HandleFunc("/ovh/cloud/volumes", ovhHandler.DescribeCloudVolumesHandler)
	r.HandleFunc("/ovh/cloud/snapshots", ovhHandler.DescribeCloudSnapshotsHandler)
	r.HandleFunc("/ovh/cloud/alerts", ovhHandler.DescribeCloudAlertsandler)
	r.HandleFunc("/ovh/cloud/currenxt", ovhHandler.DescribeCurrentUsageHandler)
	r.HandleFunc("/ovh/cloud/images", ovhHandler.DescribeCloudImagesHandler)
	r.HandleFunc("/ovh/cloud/ip", ovhHandler.DescribeCloudIpsHandler)
	r.HandleFunc("/ovh/cloud/network/private", ovhHandler.DescribeCloudPrivateNetworksHandler)
	r.HandleFunc("/ovh/cloud/network/public", ovhHandler.DescribeCloudPublicNetworksHandler)
	r.HandleFunc("/ovh/cloud/failover/ip", ovhHandler.DescribeCloudFailoverIpsHandler)
	r.HandleFunc("/ovh/cloud/vrack", ovhHandler.DescribeCloudVRacksHandler)
	r.HandleFunc("/ovh/cloud/kube/clusters", ovhHandler.DescribeCloudKubeClustersHandler)
	r.HandleFunc("/ovh/cloud/kube/nodes", ovhHandler.DescribeCloudKubeNodesHandler)
	r.HandleFunc("/ovh/cloud/sshkeys", ovhHandler.DescribeCloudSSHKeysHandler)
	r.HandleFunc("/ovh/cloud/quotas", ovhHandler.DescribeCloudLimitsHandler)
	r.HandleFunc("/ovh/cloud/ssl/certificates", ovhHandler.DescribeSSLCertificatesHandler)
	r.HandleFunc("/ovh/cloud/ssl/gateways", ovhHandler.DescribeSSLGatewaysHandler)
	r.HandleFunc("/ovh/cloud/profile", ovhHandler.DescribeProfileHandler)
	r.HandleFunc("/ovh/cloud/tickets", ovhHandler.DescribeTicketsHandler)

	r.HandleFunc("/digitalocean/account", digitaloceanHandler.AccountProfileHandler)
	r.HandleFunc("/digitalocean/actions", digitaloceanHandler.ActionsHistoryHandler)
	r.HandleFunc("/digitalocean/cdns", digitaloceanHandler.ContentDeliveryNetworksHandler)
	r.HandleFunc("/digitalocean/certificates", digitaloceanHandler.CertificatesHandler)
	r.HandleFunc("/digitalocean/databases", digitaloceanHandler.DatabasesHandler)
	r.HandleFunc("/digitalocean/domains", digitaloceanHandler.DomainsHandler)
	r.HandleFunc("/digitalocean/droplets", digitaloceanHandler.DropletsHandler)
	r.HandleFunc("/digitalocean/firewalls/list", digitaloceanHandler.DescribeFirewallsHandler)
	r.HandleFunc("/digitalocean/firewalls/unsecure", digitaloceanHandler.DescribeUnsecureFirewallsHandler)
	r.HandleFunc("/digitalocean/floatingips", digitaloceanHandler.FloatingIpsHandler)
	r.HandleFunc("/digitalocean/k8s", digitaloceanHandler.KubernetesHandler)
	r.HandleFunc("/digitalocean/keys", digitaloceanHandler.SSHKeysHandler)
	r.HandleFunc("/digitalocean/loadbalancers", digitaloceanHandler.LoadBalancersHandler)
	r.HandleFunc("/digitalocean/projects", digitaloceanHandler.ProjectsHandler)
	r.HandleFunc("/digitalocean/records", digitaloceanHandler.RecordsHandler)
	r.HandleFunc("/digitalocean/snapshots", digitaloceanHandler.SnapshotsHandler)
	r.HandleFunc("/digitalocean/volumes", digitaloceanHandler.VolumesHandler)

	//r.PathPrefix("/").Handler(http.FileServer(assetFS()))

	headersOk := handlers.AllowedHeaders([]string{"profile"})
	loggedRouter := handlers.LoggingHandler(os.Stdout, handlers.CORS(headersOk)(r))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), loggedRouter)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Server started on port %d", port)
	}
}

func main() {

	e := godotenv.Load()

	if e != nil {
		log.Println(e)
	}

	applicationPortEnv := os.Getenv("APPLICATION_PORT")

	applicationPort, err := strconv.Atoi(applicationPortEnv)

	if err != nil {
		log.Println("Application port should be integer")
		return
	}

	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPort := os.Getenv("REDIS_PORT")

	multipleAwsAccountEnv := os.Getenv("MULTIPLE_AWS_ACCOUNT")
	multipleAwsAccount, err := strconv.ParseBool(multipleAwsAccountEnv)

	if err != nil {
		log.Println("multiple_aws_account should be bool")
	}

	redisUrl := redisAddress + ":" + redisPort

	cache := &Redis{
		Addr:       redisUrl,
		Expiration: time.Duration(DEFAULT_DURATION),
	}
	startServer(applicationPort, cache, "", multipleAwsAccount)

}

// need to expose docker host and port
//docker run -p 3972:3972 ....... etc
// host will be docker.host.internal to point to localhost
//
