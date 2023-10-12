Init:
	make submodule_init

gen_abi:
	abigen -abi - -pkg $(PACKAGE) -type $(NAME) -out $(OUT)

submodule_init:
	git submodule init
	git submodule update

#MODULES_PATH = $(shell find ./contracts/omni_swap/repo/ethereum/export/abi -type f)
#CONTRACT_PATH = "./contracts/omni_swap/"
#PACKAGE_NAME=omni_swap
gen_file_abis:
	@for i in $(MODULES_PATH); do \
	  filename=$$(basename $$i); \
      filename_no_ext=$${filename%.*}; \
      abi_content=$$(cat $$i | jq '.abi?'); \
	  if [ "$$abi_content" = "" ]; then \
      	abi_content=$$(cat $$i) ; \
	  fi; \
	  echo $$abi_content | make gen_abi PACKAGE=$(PACKAGE_NAME) NAME=$${filename_no_ext} OUT=$(CONTRACT_PATH)$${filename_no_ext}.go; \
	done

gen_omni_abis:
	make gen_file_abis MODULES_PATH=./contracts/omni_swap/repo/ethereum/export/abi/*.json PACKAGE_NAME=omni_swap CONTRACT_PATH="./contracts/omni_swap/"
gen_other_abi:
	make gen_file_abis MODULES_PATH=./contracts/*.json PACKAGE_NAME=contracts CONTRACT_PATH="./contracts/"