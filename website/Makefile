MKDOCS   := squidfunk/mkdocs-material:9.6
ROOT_DIR := $(shell dirname $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

build: assets
	docker run --rm -t -v "$(ROOT_DIR):/mnt" -w /mnt/website $(MKDOCS) -- build

watch: assets
	docker run --rm -t -v "$(ROOT_DIR):/mnt" -w /mnt/website -p 8000:8000 $(MKDOCS) -- serve --watch-theme --dev-addr 0.0.0.0:8000

assets:
	cp $(ROOT_DIR)/assets/demo.gif $(ROOT_DIR)/website/docs/assets/demo.gif
	cp $(ROOT_DIR)/assets/tui.gif $(ROOT_DIR)/website/docs/assets/tui.gif
	cp $(ROOT_DIR)/assets/tui.png $(ROOT_DIR)/website/docs/assets/tui.png