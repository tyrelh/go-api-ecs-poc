#! /usr/bin/env bash

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

read -p "🚸 This will terminate all Go API POC resources for a region, excluding any centralized databases. Hit return to continue..."

echo "Terminating Go API POC in $AWS_REGION..."

STACKS=(
    "ecs-service-stack.yml:go-api-poc-ecs-service"
    "ecs-task-stack.yml:go-api-poc-ecs-task"
    "alb-stack.yml:go-api-poc-alb"
    "ecr-stack.yml:go-api-poc-ecr"
    "networking-stack.yml:go-api-poc-networking"
)
ECR_REPOSITORY_NAME="go-api-poc-ecr-repository"

for stack in "${STACKS[@]}"; do
    TEMPLATE="${stack%%:*}"
    NAME="${stack##*:}"

    if [ "$NAME" == "go-api-poc-ecr" ]; then
        # Delete all images in the ECR repository
        command="aws ecr list-images --region $AWS_REGION --repository-name $ECR_REPOSITORY_NAME --query 'imageIds[*]' --output json --profile $PROFILE"
        echo "Running: $command"
        IMAGE_IDS=$(eval "$command")
        if [ -n "$IMAGE_IDS" ] && [ "$IMAGE_IDS" != "[]" ]; then
            command="aws ecr batch-delete-image --region $AWS_REGION --repository-name $ECR_REPOSITORY_NAME --image-ids '$IMAGE_IDS' --profile $PROFILE"
            echo "Running: $command"
            result=$(eval "$command")
            failures=$(echo "$result" | jq -r '.failures')
            if [ -n "$failures" ] && [ "$failures" != "[]" ]; then
                echo "❌ Some images failed to delete: $failures"
                echo "Reattempting to delete images..."
                command="aws ecr list-images --region $AWS_REGION --repository-name $ECR_REPOSITORY_NAME --query 'imageIds[*]' --output json --profile $PROFILE"
                echo "Running: $command"
                IMAGE_IDS=$(eval "$command")
                command="aws ecr batch-delete-image --region $AWS_REGION --repository-name $ECR_REPOSITORY_NAME --image-ids '$IMAGE_IDS' --profile $PROFILE"
                echo "Running: $command"
                result=$(eval "$command")
                if [ $? -ne 0 ]; then
                    echo "❌ Failed to delete images from the ECR repository"
                    write_deployment_status "$AWS_REGION" "degraded"
                    exit 1
                fi
            fi
            if [ $? -ne 0 ]; then
                echo "❌ Failed to delete images from the ECR repository"
                write_deployment_status "$AWS_REGION" "degraded"
                exit 1
            fi
        fi
    fi

    # Delete the CloudFormation stack
    command="aws cloudformation delete-stack --stack-name $NAME --profile $PROFILE --region $AWS_REGION"
    echo "Running: $command"
    eval "$command"
    if [ $? -ne 0 ]; then
        echo "❌ Failed to delete $NAME"
        write_deployment_status "$AWS_REGION" "degraded"
        exit 1
    fi
    stack_status="CREATE_COMPLETE"
    while [ "$stack_status" != "DELETE_COMPLETE" ]; do
        command="aws cloudformation describe-stacks --stack-name $NAME --profile $PROFILE --region $AWS_REGION"
        # echo "Running: $command"
        resut=$(eval "$command")
        if [ $? -ne 0 ]; then
            echo "Fail to delete: ${result}"
            break
        fi
        stack_status=$(eval "$command" | jq -r '.Stacks[0].StackStatus')
        echo "Stack status: $stack_status"
        if [ "$stack_status" == "DELETE_FAILED" ]; then
            echo "❌ Failed to delete $NAME"
            write_deployment_status "$AWS_REGION" "degraded"
            exit 1
        fi
        sleep 2
    done
done

echo "✅ Go API POC in $AWS_REGION terminated successfully"
write_deployment_status "$AWS_REGION" "inactive"