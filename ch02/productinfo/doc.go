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

//go:generate echo "For Service"
//go:generate protoc -I proto/ proto/product_info.proto --go_out=plugins=grpc:go/server/ecommerce
//go:generate echo "For Client"
//go:generate protoc -I proto/ proto/product_info.proto --go_out=plugins=grpc:go/client/ecommerce

