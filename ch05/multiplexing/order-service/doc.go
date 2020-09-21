// +	Package productinfo for products
// +
// +*/
// +
// +package productinfo
// +
// +//go:generate echo "For Service"
// +//go:generate protoc -I proto/ proto/product_info.proto --go_out=plugins=grpc:go/server/ecommerce
// +//go:generate echo "For Client"
// +//go:generate protoc -I proto/ proto/product_info.proto --go_out=plugins=grpc:go/client/ecommerce

//	Package productinfo for products

package productinfo

//go:generate cd /Users/growth/drinnovations/mywork_jmd/gogrpc/gooreillygrpcuprunning/samples/ch05/multiplexing/order-service
//go:generate echo "For Service"
//go:generate protoc -I proto/ proto/health_check.proto --go_out=plugins=grpc:go/server/healthcheck
//go:generate echo "For Client"
//go:generate protoc -I proto/ proto/health_check.proto --go_out=plugins=grpc:go/client/healthcheck

