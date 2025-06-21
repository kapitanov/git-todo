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
	@echo "items:" > .git/TODO
	@echo "  - id: 478e1212" >> .git/TODO
	@echo "    done: true" >> .git/TODO
	@echo "    title: Basic CLI" >> .git/TODO
	@echo "  - id: fa1e7a89" >> .git/TODO
	@echo "    done: true" >> .git/TODO
	@echo "    title: Refactor" >> .git/TODO
	@echo "  - id: 66f3902f" >> .git/TODO
	@echo "    done: true" >> .git/TODO
	@echo "    title: Automatable CLI (json output, tab-separated output and other stuff)" >> .git/TODO
	@echo "  - id: 57ed61f0" >> .git/TODO
	@echo "    done: true" >> .git/TODO
	@echo "    title: TUI" >> .git/TODO
	@echo "  - id: f9eb2530" >> .git/TODO
	@echo "    done: true" >> .git/TODO
	@echo "    title: Git hooks" >> .git/TODO
	@echo "  - id: 69e27356" >> .git/TODO
	@echo "    title: README" >> .git/TODO
	@echo "  - id: 39fdec11" >> .git/TODO
	@echo "    done: true" >> .git/TODO
	@echo "    title: Tests" >> .git/TODO
	@echo "  - id: 9e9cf322" >> .git/TODO
	@echo "    title: Documentation" >> .git/TODO
	@echo "  - id: '13844228'" >> .git/TODO
	@echo "    title: CI" >> .git/TODO
	@echo "  - id: 32fa9037" >> .git/TODO
	@echo "    title: New TODO Item A" >> .git/TODO
	@echo "" >> .git/TODO

items:
    - id: 478e1212
      done: true
      title: Basic CLI
    - id: fa1e7a89
      done: true
      title: Refactor
    - id: 66f3902f
      done: true
      title: Automatable CLI (json output, tab-separated output and other stuff)
    - id: 57ed61f0
      done: true
      title: TUI
    - id: f9eb2530
      done: true
      title: Git hooks
    - id: 69e27356
      title: README
    - id: 39fdec11
      done: true
      title: Tests
    - id: 9e9cf322
      title: Documentation
    - id: "13844228"
      title: CI
    - id: 32fa9037
      title: New TODO Item A

teardown_local_env:
	@[ -d .git ] && rm -rf .git || true

