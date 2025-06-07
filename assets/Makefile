default:
	@echo "Usage: make <target>"
	@echo "Available targets:"
	@echo "  demo: generate the demo"
	@echo "  tui: generate the TUI promo image"
	@exit 1

demo: check_vhs_installed setup_local_env
	vhs demo.tape
	@make -s teardown_local_env

tui: check_vhs_installed setup_local_env
	vhs tui.tape
	@make -s teardown_local_env

check_vhs_installed:
	@if ! command -v vhs &> /dev/null; then \
		echo "VHS is not installed. Please install it to use this Makefile."; \
		exit 1; \
	fi

setup_local_env: teardown_local_env
	@cd .. && make install
	@git init > /dev/null 2> /dev/null 
	@echo "[x] Basic CLI" > .git/TODO
	@echo "[x] Refactor" >> .git/TODO
	@echo "[x] Automatable CLI (json output, tab-separated output and other stuff)" >> .git/TODO
	@echo "[x] TUI" >> .git/TODO
	@echo "[x] Git hooks" >> .git/TODO
	@echo "[ ] README" >> .git/TODO
	@echo "[x] Tests" >> .git/TODO
	@echo "[ ] Documentation" >> .git/TODO
	@echo "[ ] CI" >> .git/TODO

teardown_local_env:
	@[ -d .git ] && rm -rf .git || true

