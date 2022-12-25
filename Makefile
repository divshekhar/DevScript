# Go Compiler
PROGRAM := DevScript
COMPILER := go

# Src path
SRC := ./src

LEXER := $(SRC)/lexer
PARSER := $(SRC)/parser
EVAL := $(SRC)/eval

# Bin path
BIN := ./bin

# Run test
test_all:
	$(COMPILER) test $(LEXER)
	$(COMPILER) test $(PARSER)
	$(COMPILER) test $(EVAL)

test_lexer:
	$(COMPILER) test $(LEXER)

test_parser:
	$(COMPILER) test $(PARSER)

test_eval:
	$(COMPILER) test $(EVAL)

# Build
build:
	bash ./build.sh $(PROGRAM) && $(COMPILER) build -o $(BIN)/$(PROGRAM).exe