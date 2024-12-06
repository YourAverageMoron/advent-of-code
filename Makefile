language=go
year=2024
run:
ifeq ($(language), go)
	make run_go
else ifeq ($(language), sql)
	make run_sql
endif


run_go: 
	go run ./$(language)/$(year)/day$(day)/ ./data/$(year)/day$(day)

run_sql:
	./sql/lib/run.sh $(year) day$(day) 
