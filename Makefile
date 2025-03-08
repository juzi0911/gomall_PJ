.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --server_name frontend --type HTTP  --idl ../../idl/frontend/auth_page.proto -module github.com/juzi0911/gomall_PJ/app/frontend -I ../../idl
	@cd app/frontend && cwgo server --server_name frontend --type HTTP  --idl ../../idl/frontend/home.proto -module github.com/juzi0911/gomall_PJ/app/frontend -I ../../idl