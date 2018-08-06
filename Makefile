.PHONY: dev runserver buildclient watchsrc help
.SILENT: dev help
.ONESHELL:

help:
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

dev: ## Nice and friendly to start the server and parcel, side by side
	if [[ $$TERM == *tmux* ]]; then
		tmux kill-pane -a -t $$TMUX_PANE
		tmux send-keys "make runserver" Enter
		tmux split-window -h -c "#{pane_current_path}"
		tmux send-keys "make watchsrc" Enter
		tmux select-pane -L
	else
		make runserver &
		make watchsrc &
	fi

runserver: boatsinker ## Just runs the server
	clear
	./boatsinker

boatsinker: server/**/*.go server/*.go ## Builds the server
	cd server
	go build -i -o boatsinker
	cd ..
	mv server/boatsinker .

buildclient: ## Build the client side of the application
	parcel build src/index.html

watchsrc: ## Watch the files and rebuilds the client as needed
	parcel watch src/index.html --no-hmr

