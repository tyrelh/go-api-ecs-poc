#! /usr/bin/env python3
import os, sys, argparse
# import the shared utils module
working_directory = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.dirname(working_directory))
import shared as utils

human_name = "ECS (Elastic Container Service) Task Definition"

parser = argparse.ArgumentParser(description=f"Deploy Go API POC {human_name} stack")
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
parser.add_argument(
    "--ci",
    action="store_true"
)
args = parser.parse_args()

profile = "infrastructure-admin"
if args.aws_account == "dev":
    profile = "infrastructure-admin-dev"

stack = "go-api-poc-ecs-task"
template_file = "ecs-task-stack.yml"
if args.ci:
    template_file = f"infrastructure/{template_file}"

print(f"AWS Account: {args.aws_account}")
print(f"AWS Region: {args.aws_region}")
print(f"AWS Profile: {profile}")
print(f"Template File: {template_file}")
print(f"Deploying {stack}...")
command = f'aws cloudformation deploy --template-file {template_file} --stack-name {stack} --parameter-overrides "AwsAccount={args.aws_account}" --profile {profile} --region {args.aws_region} --capabilities CAPABILITY_IAM'
utils.run_process(command)