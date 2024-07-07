.DEFAULT_GOAL := skeleton
.PHONY: skeleton

skeleton:
	@ if [[ -n $$day && -n $$year ]]; then \
		go run skeleton/main.go --year $(year) --day $(day); \
	elif [[ -n $$year ]]; then \
		go run skeleton/main.go --year $(year); \
	elif [[ -n $$day ]]; then \
		go run skeleton/main.go --day $(day); \
	else \
		go run skeleton/main.go; \
	fi ||:


