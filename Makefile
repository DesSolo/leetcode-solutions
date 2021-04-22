SOLUTIONS_PATH=$(PWD)/solutions/

new:
	python scripts/generator.py -l Go -e go -s _test -t $(PWD)/scripts -d ${SOLUTIONS_PATH} ${SLUG}

generate-readme:
	bash $(PWD)/scripts/md_gen.sh &> README.md
