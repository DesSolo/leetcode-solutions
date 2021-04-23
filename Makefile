SOLUTIONS_PATH=$(PWD)/solutions
SCRIPTS_PATH=$(PWD)/scripts

test:
	go test ${SOLUTIONS_PATH}/*.go

new:
	python scripts/generator.py -l Go -e go -s _test -t ${SCRIPTS_PATH} -d ${SOLUTIONS_PATH} ${SLUG}

generate-readme:
	bash ${SCRIPTS_PATH}/md_gen.sh &> README.md
