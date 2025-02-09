#! /usr/bin/env python3
import os, sys, argparse
# import the shared utils module
working_directory = os.path.dirname(os.path.realpath(__file__))
sys.path.append(working_directory)
import shared as utils

parser = argparse.ArgumentParser(description="Deploy Go API POC")
parser.add_argument(
    "--aws-account",
    type=str,
    choices=["dev", "production"],
    required=True,
    help="The AWS Account to act on"
)
parser.add_argument(
    "--aws-region",
    type=str,
    required=False,
    default="ca-west-1",
    help="The AWS Region to deploy the stack to"
)
args = parser.parse_args()

profile = "infrastructure-admin"
if args.aws_account == "dev":
    profile = "infrastructure-admin-dev"

print(f"AWS Account: {args.aws_account}")
print(f"AWS Region: {args.aws_region}")
print(f"AWS Profile: {profile}")

input("🚸 This will terminate all Go API POC resources for a region, excluding any centralized databases. Hit return to continue...")

print(f"Terminating Go API POC in ${args.aws_region}...")

stacks = [
    {"template": "ecs-service-stack.yml", "name": "go-api-poc-ecs-service"},
    {"template": "networking-stack.yml", "name": "go-api-poc-networking"},
    {"template": "alb-stack.yml", "name": "go-api-poc-alb"},
    {"template": "ecr-stack.yml", "name": "go-api-poc-ecr"},
    {"template": "ecs-task-stack.yml", "name": "go-api-poc-ecs-task"},
]
ecr_repository_name = "go-api-poc-ecr-repository"

for stack in stacks:
    if stack["name"] == "go-api-poc-ecr":
        # https://stackoverflow.com/questions/73305551/how-to-delete-all-images-in-an-ecr-repository
        command = f'aws ecr list-images --region ${args.aws_region} --repository-name ${ecr_repository_name} --query "imageIds[*]" --output json'
        print(f"Running: {command}")
        image_ids = utils.run_process_sync_in_background_return_result(command)
        print(f"Image IDs: {image_ids}")
        print("Deleting all images in the ECR repository...")
        command = f'aws ecr batch-delete-image --region ${args.aws_region} --repository-name ${ecr_repository_name} --image-ids "${image_ids}"'
        code = utils.run_process(command)
        if code != 0:
            print(f"❌ Failed to delete images from the ECR repository")
            sys.exit(code)
    command = f'aws cloudformation delete-stack --stack-name {stack["name"]} --profile {profile} --region {args.aws_region}'
    code = utils.run_process(command)
    if code != 0:
        print(f"❌ Failed to delete {stack["name"]}")
        sys.exit(code)

print(f"✅ Go API POC in {args.aws_region} terminated successfully")
utils.write_deployment_status(args.aws_region, "inactive")
