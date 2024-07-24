serve:
	  @echo "Running the application..."
	  go run ./cmd/

generate:
	  @echo "Generating Templ code..."
	  templ generate

all: generate run