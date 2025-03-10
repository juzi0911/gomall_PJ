.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --server_name frontend --type HTTP  --idl ../../idl/frontend/auth_page.proto -module github.com/juzi0911/gomall_PJ/app/frontend -I ../../idl
	@cd app/frontend && cwgo server --server_name frontend --type HTTP  --idl ../../idl/frontend/home.proto -module github.com/juzi0911/gomall_PJ/app/frontend -I ../../idl

.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/juzi0911/gomall_PJ/rpc_gen --I ../idl  --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module github.com/juzi0911/gomall_PJ/app/user --pass "-use github.com/juzi0911/gomall_PJ/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/user.proto