#!/usr/bin/env bash

# Function to write deployment status to a file
write_deployment_status() {
    local region=$1
    local status=$2
    local file="deployment_status.txt"

    if [ -f "$file" ]; then
        if grep -q "^$region:" "$file"; then
            sed -i '' "s/^$region:.*/$region: $status/" "$file"
        else
            echo "$region: $status" >> "$file"
        fi
    else
        echo "$region: $status" > "$file"
    fi
}

# Function to run a command and print its output
run_process() {
    local command=$1
    echo "Running: $command"
    eval "$command"
    local code=$?
    if [ $code -ne 0 ]; then
        echo "❌ Command failed with code $code"
    else
        echo "🟢 Command succeeded"
    fi
    return $code
}

# Parse arguments
AWS_ACCOUNT=""
AWS_REGION="ca-west-1"

while [[ $# -gt 0 ]]; do
    case $1 in
        --aws-account)
            AWS_ACCOUNT="$2"
            shift
            shift
            ;;
        --aws-region)
            AWS_REGION="$2"
            shift
            shift
            ;;
        *)
            shift
            ;;
    esac
done

if [ -z "$AWS_ACCOUNT" ]; then
    echo "Error: --aws-account is required"
    exit 1
fi

PROFILE="infrastructure-admin"
if [ "$AWS_ACCOUNT" == "dev" ]; then
    PROFILE="infrastructure-admin-dev"
fi

echo "AWS Account: $AWS_ACCOUNT"
echo "AWS Region: $AWS_REGION"
echo "AWS Profile: $PROFILE"

read -p "🚸 This will deploy all infra and build and deploy the current code on your working branch. Press enter to continue..."

echo "Deploying Go API POC..."

STACKS=(
    "networking-stack.yml:go-api-poc-networking"
    "alb-stack.yml:go-api-poc-alb"
    "ecr-stack.yml:go-api-poc-ecr"
    "ecs-task-stack.yml:go-api-poc-ecs-task"
    "ecs-service-stack.yml:go-api-poc-ecs-service"
)

for stack in "${STACKS[@]}"; do
    TEMPLATE="${stack%%:*}"
    NAME="${stack##*:}"

    command="aws cloudformation deploy --template-file ../infrastructure/$TEMPLATE --stack-name $NAME --parameter-overrides AwsAccount=$AWS_ACCOUNT --profile $PROFILE --region $AWS_REGION --capabilities CAPABILITY_IAM"
    run_process "$command"
    if [ $? -ne 0 ]; then
        echo "❌ Failed to deploy $TEMPLATE"
        write_deployment_status "$AWS_REGION" "failed"
        exit 1
    fi
    echo "🟢 $TEMPLATE deployed successfully"

    if [ "$TEMPLATE" == "ecr-stack.yml" ]; then
        echo "Pushing the image to the new ECR..."
        command="make -C .. push AWS_REGION=$AWS_REGION"
        run_process "$command"
        if [ $? -ne 0 ]; then
            echo "❌ Failed to push the image to the new ECR"
            write_deployment_status "$AWS_REGION" "failed"
            exit 1
        fi
    fi
done

echo "✅ Go API POC deployed successfully"
write_deployment_status "$AWS_REGION" "active"
echo "🚀 Access the API at the following URL: https://goapi.giftbitdev.com/go/system/health"
