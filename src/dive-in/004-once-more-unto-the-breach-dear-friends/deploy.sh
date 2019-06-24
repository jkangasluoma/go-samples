# Run like PROFILE=XXX ./deploy.sh

sam deploy \
    --template-file packaged.yaml \
    --stack-name go-sample \
    --capabilities CAPABILITY_IAM \
    --region eu-west-1 --profile $PROFILE
