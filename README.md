# Upload file to S3

A restful API that can upload files to s3.

## How to use?

Download the latest release version, and copy `.env.example` to `.env`.  
Complete the .env file, and run the binary file to start the server.

For the docker usage guide, It will release soon.

## API Endpoints

HTTP POST /api/v1/upload

## Deploy

Edit environment in the [docker-compose.yaml](docker-compose.yaml), and run it.

```
docker-compose up -d
```