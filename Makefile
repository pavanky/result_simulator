all: simulate

simulate: src/simulate.go
	@go build $<

clean:
	rm -f simulate

run: simulate
	@./simulate ./data/IPL/2014
