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

input("🚸 This will deploy all infra and build and deploy the current code on your working branch. Press enter to continue...")

print("Deploying Go API POC...")

stacks = [
    {"template": "networking-stack.yml", "name": "go-api-poc-networking"},
    {"template": "alb-stack.yml", "name": "go-api-poc-alb"},
    {"template": "ecr-stack.yml", "name": "go-api-poc-ecr"},
    {"template": "ecs-task-stack.yml", "name": "go-api-poc-ecs-task"},
    {"template": "ecs-service-stack.yml", "name": "go-api-poc-ecs-service"},
]

for stack in stacks:
    # command = f"python3 {working_directory}/infrastructure/{script} --aws-account {args.aws_account} --aws-region {args.aws_region} --ci"
    command = f'aws cloudformation deploy --template-file infrastructure/{stack["template"]} --stack-name {stack["name"]} --parameter-overrides "AwsAccount={args.aws_account}" --profile {profile} --region {args.aws_region} --capabilities CAPABILITY_IAM'
    print(f"Running: {command}")
    code = utils.run_process(command)
    if code != 0:
        print(f"❌ Failed to deploy {stack["template"]}")
        utils.write_deployment_status(args.aws_region, "failed")
        sys.exit(code)
    print(f"🟢 {stack["template"]} deployed successfully")
    if stack["template"] == "ecr-stack.yml":
        print("Pushing the image to the new ECR...")
        command = f"make push AWS_REGION={args.aws_region}"
        print(f"Running: {command}")
        code = utils.run_process(command)
        if code != 0:
            print(f"❌ Failed to push the image to the new ECR")
            utils.write_deployment_status(args.aws_region, "failed")
            sys.exit(code)
        # print(" Image pushed successfully")
print("✅ Go API POC deployed successfully")
utils.write_deployment_status(args.aws_region, "active")
print("🚀 Access the API at the following URL: https://goapi.giftbitdev.com/go/system/health")