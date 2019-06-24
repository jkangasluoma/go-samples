# Requires ENV variables PROFILE
# Run like PROFILE=XXX S3_BUCKET=YYY ./build.sh

make build
sam package \
    --output-template-file packaged.yaml \
    --s3-bucket $S3_BUCKET \
    --s3-prefix serverless/go-sample --profile $PROFILE
