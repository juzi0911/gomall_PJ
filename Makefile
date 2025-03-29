export ROOT_MOD=github.com/juzi0911/gomall_PJ

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --server_name frontend --type HTTP  --idl ../../idl/frontend/auth_page.proto -module ${ROOT_MOD}/app/frontend -I ../../idl
	@cd app/frontend && cwgo server --server_name frontend --type HTTP  --idl ../../idl/frontend/home.proto -module ${ROOT_MOD}/app/frontend -I ../../idl

.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module ${ROOT_MOD}/rpc_gen --I ../idl  --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module ${ROOT_MOD}/app/user --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/user.proto

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --service product --module ${ROOT_MOD}/rpc_gen --I ../idl  --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module ${ROOT_MOD}/app/product --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/product.proto