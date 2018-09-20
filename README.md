# A simple Slack bot

## Deployment to AWS

### Using serverless

1. Install [serverless](https://www.google.com) `npm install -g serverless`
2. Run `serverless deploy -v`

### Using Up

1. Install [up](https://up.docs.apex.sh/) `curl -sf https://up.apex.sh/install | sh`
2. Run `up`

### Deploy Manually

1. Build lite-bot by `make build` or `env GOOS=linux go build -ldflags="-s -w" -o bin/litebot main.go`
2. Compress the build file into zip
3. Create a new lambda function, choose runtime as `Go 1.x`
4. Upload zip file to AWS lambda function create at (3). Set Handler=`litebot`. Then `Save`
5. Create AWS API Gateway, named `lite-bot`
6. `Create Method` for Api gateway: Choose `POST` since Slack slash-command will only push as `POST`. Check `Use Lambda Proxy integration`. Set Lambda function to `lite-bot`
7. `Deploy API` for Api gateway: named `dev`, then click `Deploy`. Now the Invoke URL is created `https://xxxxxx.execute-api.ap-southeast-1.amazonaws.com/dev`
8. Config Invoke URL to Slack
9. Done

### Grant internet access to VPC Lambda function

(follow tutorial [video](https://youtu.be/JcRKdEP94jM) by Kien at [aws document](https://aws.amazon.com/premiumsupport/knowledge-center/internet-access-lambda-function/))

1. Best practise is create 2 seperate subnet for public and private-lambda function
2. Create NAT Gateway
3. Bind the NAT Gateway with the internet-subnet through Route gateway: set 0.0.0.0/0 - nat-...
4. Create Internet Gateway (if not existed)
5. Bind the Internet Gateway with the external-subnet through Route gateway: set 0.0.0.0/0 - igw-...
6. Go to lambda function console, choose VPC, `only` select the internal subnets, and security group. Remember to check if the security group is allowed outbound connection
