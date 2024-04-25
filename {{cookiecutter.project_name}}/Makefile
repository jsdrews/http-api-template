RED:=\033[1;31m
GREEN:=\033[1;32m
NC:=\033[0m

bootstrap:
	cp .envrc.template .envrc
ifneq ("$(shell which task > /dev/null && echo 0)","0")
	@curl --location https://taskfile.dev/install.sh | sh -s -- -d
	@printf "${GREEN}task is installed: ./bin/task${NC}\n"
else
	@printf "${GREEN}task is installed${NC}\n"
endif
ifneq ("$(shell which direnv > /dev/null && echo 0)","0")
	@echo
	@printf "${RED}You will need to install direnv: https://direnv.net/docs/installation.html${NC}\n"
else
	@printf "${GREEN}direnv is installed (run: direnv allow)${NC}\n"
endif
ifneq ("$(shell which docker > /dev/null && echo 0)","0")
	@printf "${RED}You will need to install docker: https://docs.docker.com/engine/install/${NC}\n"
else
	@printf "${GREEN}docker is installed${NC}\n"
endif
ifneq ("$(shell which helm > /dev/null && echo 0)","0")
	@printf "${RED}You will need to install helm: https://helm.sh/docs/intro/install/\n"
else
	@printf "${GREEN}helm is installed${NC}\n"
endif
