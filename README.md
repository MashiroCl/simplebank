## Getting Started

### Build and Run
* prerequisite
```shell
brew install golang-migrate
brew install sqlc
```

* Setup Infrastructure
```shell
make network
make postgres
make createdb
make migrateup
```

* How to run
```shell
make 
make build_image
make run_image
```