# postgres-example

# Install Docker:
[https://docs.docker.com/desktop/install/mac-install](https://docs.docker.com/desktop/install/mac-install)

# Install postgres image:
docker pull postgres

# Start docker:
docker run --name test-postgres -e POSTGRES_USER=myusername -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres

# GORM:
[https://gorm.io/docs/associations.html](https://gorm.io/docs/associations.html)

# Build and run
go build
./postgres-example