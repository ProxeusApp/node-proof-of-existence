# node-proof-of-existence
An external node implementation for Proxeus core. Returns social posts information given an id or url 

## Implementation

Current implementation only supports Twitter posts. 

## Usage

It is recommended to start it using docker.

The latest image is available at `proxeus/node-proof-of-existence:latest`

See the configuration paragraph for more information on what environments variables can be overridden

## Configuration

The following parameters can be set via environment variables. 
Please note a few of them are required


| Environmentvariable | Required | Default value
--- | --- | --- 
TWITTER_CONSUMER_KEY | X | 
TWITTER_CONSUMER_SECRET | X | 
PROXEUS_INSTANCE_URL |  | http://127.0.0.1:1323
SERVICE_NAME |  | proof-of-existence
SERVICE_URL |  | 127.0.0.1
SERVICE_PORT |  | 8012
AUTH_KEY |  | auth

## Deployment

The node is available as docker image and can be used within a typical Proxeus Platform setup by including the following docker-compose service:

```
  node-proof-of-existence:
    image: proxeus/node-proof-of-existence:latest
    container_name: xes_node-proof-of-existence
    restart: unless-stopped
    environment:
      PROXEUS_INSTANCE_URL: "http://xes-platform:1323"
      SERVICE_SECRET: yourSecret
      SERVICE_PORT: 8014
      SERVICE_URL: "http://node-proof-of-existence:8014"
      TWITTER_CONSUMER_KEY: "yourTwitterConsumerKey"
      TWITTER_CONSUMER_SECRET: "yourTwitterConsumerSecret"
    ports:
      - "8014:8014"
```
