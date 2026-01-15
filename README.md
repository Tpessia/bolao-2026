# Install Dependencies
go mod download

# Run
make run



go get -u github.com/swaggo/swag
go get -u github.com/swaggo/http-swagger
go install github.com/swaggo/swag/cmd/swag@latest
swag init
go mod tidy



// External Frameworks (ranked by usage in Coinbase):
// net/http
// github.com/gorilla/mux
// github.com/go-chi/chi
// github.com/gin-gonic/gin

// github.com/danielgtaylor/huma to auto-generate swagger docs, with net/http
// or others it's not possible because they interect directly with the socket,
// while in .NET for example the controller handler returns the Type directly