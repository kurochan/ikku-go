.PHONY: test
test:
	go test -v ./...

.PHONY: test/all
test/all: test test/fuzz

.PHONY: test/fuzz
test/fuzz:
	go test -fuzz=FuzzTestJudgeHiraganaKatakana -fuzztime 10s
	go test -fuzz=FuzzTestJudgeAllRunes -fuzztime 5s

.PHONY: lint
lint:
	gofmt -l -w .
	golangci-lint run --fix

.PHONY: clean
clean:
	go clean -fuzzcache
